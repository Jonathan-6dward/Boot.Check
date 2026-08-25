package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const PromptVersion = "1.0.0"

// Config contains provider-specific values supplied at runtime, never stored
// in source control or in the evidence package.
type Config struct {
	Endpoint         string
	APIKey           string
	Model            string
	Timeout          time.Duration
	MaxResponseBytes int64
	MaxRetries       int
}

type chatRequest struct {
	Model          string         `json:"model"`
	Messages       []chatMessage  `json:"messages"`
	Temperature    float64        `json:"temperature,omitempty"`
	ResponseFormat responseFormat `json:"response_format"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type responseFormat struct {
	Type       string         `json:"type"`
	JSONSchema map[string]any `json:"json_schema,omitempty"`
}

// VerdictResponse is the validated semantic result consumed by the report
// generator. Keep it deliberately smaller than the raw provider response.
type VerdictResponse struct {
	Verdict              string              `json:"verdict"`
	Confidence           float64             `json:"confidence"`
	HeadlinePlain        string              `json:"headline_plain"`
	PlainLanguageSummary string              `json:"plain_language_summary"`
	FiveWTwoH            FiveWTwoH           `json:"five_w_two_h"`
	MITRE                []MITREMapping      `json:"mitre"`
	SupportingEvidence   []EvidenceReference `json:"supporting_evidence"`
	Limitations          []string            `json:"limitations"`
	RecommendedNextSteps []string            `json:"recommended_next_steps"`
	TechnicalAppendix    TechnicalAppendix   `json:"technical_appendix"`
	SafetyNotice         string              `json:"safety_notice"`
}

type FiveWTwoH struct {
	What        string   `json:"what"`
	Who         string   `json:"who"`
	When        string   `json:"when"`
	Where       string   `json:"where"`
	Why         string   `json:"why"`
	How         string   `json:"how"`
	Impact      string   `json:"impact"`
	EvidenceIDs []string `json:"evidence_ids"`
}

type MITREMapping struct {
	TechniqueID string   `json:"technique_id"`
	Name        string   `json:"name"`
	Rationale   string   `json:"rationale"`
	EvidenceIDs []string `json:"evidence_ids"`
}

type EvidenceReference struct {
	EvidenceID  string `json:"evidence_id"`
	Role        string `json:"role"`
	Explanation string `json:"explanation"`
}

type TechnicalAppendix struct {
	Processes   []map[string]any `json:"processes"`
	Persistence []map[string]any `json:"persistence"`
	Network     []map[string]any `json:"network"`
	Notes       []string         `json:"notes"`
}

// Analyze sends only the caller-approved, already-redacted JSON to the
// configured provider. This client does not collect endpoint data and does
// not perform any local remediation.
// TODO(local-agent): bind ValidateEvidencePackage to the collector schema
// before calling this method, and pass the exact consent record through the
// call boundary. Never call Analyze without an explicit UI confirmation.
func Analyze(ctx context.Context, cfg Config, schemaVersion, dataMode string, evidenceJSON []byte) (VerdictResponse, error) {
	if cfg.Endpoint == "" || cfg.APIKey == "" || cfg.Model == "" {
		return VerdictResponse{}, errors.New("LLM endpoint, API key and model are required")
	}
	if dataMode != "redacted" && dataMode != "full" {
		return VerdictResponse{}, errors.New("invalid data mode")
	}
	if len(evidenceJSON) == 0 {
		return VerdictResponse{}, errors.New("evidence package is empty")
	}

	system, user := BuildPrompts(schemaVersion, PromptVersion, dataMode, evidenceJSON)
	body, err := json.Marshal(chatRequest{
		Model:       cfg.Model,
		Messages:    []chatMessage{{Role: "system", Content: system}, {Role: "user", Content: user}},
		Temperature: 0,
		ResponseFormat: responseFormat{Type: "json_schema", JSONSchema: map[string]any{
			"name":   "bootcheck_verdict",
			"strict": true,
			"schema": VerdictJSONSchema(),
		}},
	})
	if err != nil {
		return VerdictResponse{}, fmt.Errorf("marshal request: %w", err)
	}

	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 45 * time.Second
	}
	maxBytes := cfg.MaxResponseBytes
	if maxBytes <= 0 {
		maxBytes = 1024 * 1024
	}
	maxRetries := cfg.MaxRetries
	if maxRetries < 0 || maxRetries > 2 {
		return VerdictResponse{}, errors.New("MaxRetries must be between 0 and 2")
	}

	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		result, retry, callErr := doRequest(ctx, cfg, timeout, maxBytes, body)
		if callErr == nil {
			if err := ValidateVerdict(result); err != nil {
				return VerdictResponse{}, fmt.Errorf("invalid verdict: %w", err)
			}
			return result, nil
		}
		lastErr = callErr
		if !retry || attempt == maxRetries {
			break
		}
		time.Sleep(time.Duration(attempt+1) * 250 * time.Millisecond)
	}
	return VerdictResponse{}, fmt.Errorf("LLM request failed after bounded retry: %w", lastErr)
}

func doRequest(ctx context.Context, cfg Config, timeout time.Duration, maxBytes int64, body []byte) (VerdictResponse, bool, error) {
	requestCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(requestCtx, http.MethodPost, cfg.Endpoint, bytes.NewReader(body))
	if err != nil {
		return VerdictResponse{}, false, err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return VerdictResponse{}, true, err
	}
	defer resp.Body.Close()
	limited := io.LimitReader(resp.Body, maxBytes)
	responseBody, err := io.ReadAll(limited)
	if err != nil {
		return VerdictResponse{}, true, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Do not include responseBody in the error: provider responses may
		// contain personal data or secrets. Keep only status metadata.
		return VerdictResponse{}, resp.StatusCode >= 500 || resp.StatusCode == 429, fmt.Errorf("provider returned HTTP %d", resp.StatusCode)
	}

	// TODO(local-agent): adapt this envelope to the selected provider without
	// logging the raw response. Extract the structured JSON content and reject
	// tool calls, markdown wrappers and unrecognized fields.
	var result VerdictResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return VerdictResponse{}, false, fmt.Errorf("decode structured verdict: %w", err)
	}
	return result, false, nil
}

func ValidateVerdict(result VerdictResponse) error {
	if result.Verdict != "likely_safe" && result.Verdict != "suspicious" && result.Verdict != "inconclusive" {
		return errors.New("verdict must be likely_safe, suspicious or inconclusive")
	}
	if result.Confidence < 0 || result.Confidence > 1 {
		return errors.New("confidence must be between 0 and 1")
	}
	if result.HeadlinePlain == "" || result.PlainLanguageSummary == "" {
		return errors.New("plain-language fields are required")
	}
	if result.SafetyNotice == "" {
		return errors.New("safety notice is required")
	}
	for _, item := range result.MITRE {
		if item.TechniqueID == "" || item.Name == "" || item.Rationale == "" || len(item.EvidenceIDs) == 0 {
			return errors.New("each MITRE mapping needs rationale and evidence IDs")
		}
	}
	for _, ref := range result.SupportingEvidence {
		if ref.EvidenceID == "" || (ref.Role != "supports" && ref.Role != "contradicts" && ref.Role != "context") {
			return errors.New("invalid supporting evidence reference")
		}
	}
	return nil
}
