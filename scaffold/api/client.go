package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Jonathan-6dward/Boot.Check/scaffold/api/provider"
)

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

// Analyze calls the selected provider to evaluate the evidence package.
// It handles retry logic for transient errors (like quotas or timeouts)
// and maps schema failures to Inconclusive verdicts.
func Analyze(ctx context.Context, prov provider.Provider, pkg provider.EvidencePackage, maxRetries int) (VerdictResponse, error) {
	if prov == nil {
		return VerdictResponse{}, errors.New("provider is required")
	}
	if maxRetries < 0 || maxRetries > 2 {
		return VerdictResponse{}, errors.New("MaxRetries must be between 0 and 2")
	}

	var lastErr error
	var provVerdict provider.Verdict
	success := false

	for attempt := 0; attempt <= maxRetries; attempt++ {
		result, callErr := prov.Analyze(ctx, pkg)
		if callErr == nil {
			provVerdict = result
			success = true
			break
		}
		
		// If it's a schema validation error from the provider, we don't retry,
		// we just map it to Inconclusive as per ADR.
		if errors.Is(callErr, provider.ErrSchemaInvalid) {
			provVerdict = provider.Verdict{
				State:        provider.StateInconclusive,
				Summary:      "Os dados retornados pelo modelo de IA não estavam no formato esperado.",
				ProviderKind: prov.Kind(),
				ModelName:    prov.Name(),
			}
			success = true
			break
		}
		
		lastErr = callErr
		
		if attempt == maxRetries {
			break
		}
		
		time.Sleep(time.Duration(attempt+1) * 500 * time.Millisecond)
	}

	if !success {
		return VerdictResponse{}, fmt.Errorf("LLM request failed after bounded retry: %w", lastErr)
	}

	// Map provider.Verdict back to api.VerdictResponse for the report generator
	vr := VerdictResponse{
		Verdict:              string(provVerdict.State),
		Confidence:           0.8, // placeholder since new provider doesn't output confidence
		HeadlinePlain:        provVerdict.Summary,
		PlainLanguageSummary: provVerdict.Summary,
		SafetyNotice:         "Triagem gerada pelo BootCheck utilizando " + provVerdict.ModelName,
	}

	if len(provVerdict.Claims) > 0 {
		c := provVerdict.Claims[0]
		vr.FiveWTwoH = FiveWTwoH{
			What: c.What, Who: c.Who, When: c.When, Where: c.Where,
			Why: c.Why, How: c.How, Impact: c.Impact, EvidenceIDs: c.EvidenceIDs,
		}
	}
	
	for _, m := range provVerdict.MitreATTACK {
		vr.MITRE = append(vr.MITRE, MITREMapping{
			TechniqueID: m,
			Name:        "Técnica MITRE",
			Rationale:   "Mapeado pelo provedor IA",
			EvidenceIDs: []string{},
		})
	}
	
	if err := ValidateVerdict(vr); err != nil {
		return vr, fmt.Errorf("invalid verdict mapped: %w", err)
	}
	
	return vr, nil
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
	return nil
}


