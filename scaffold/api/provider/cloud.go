package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CloudVendor identifica qual API externa o usuário conectou.
// Todos oferecem camada gratuita, o que era o requisito original.
type CloudVendor string

const (
	VendorAnthropic CloudVendor = "anthropic"
	VendorOpenAI    CloudVendor = "openai"
	VendorGemini    CloudVendor = "gemini"
)

// CloudCredentials guarda só o necessário para uma chamada. O campo
// APIKey NUNCA deve ser persistido em texto plano — a camada de
// configuração é responsável por criptografar com DPAPI (Windows)
// antes de gravar em disco, e por não logar este struct.
type CloudCredentials struct {
	Vendor CloudVendor
	APIKey string
	Model  string // ex.: "claude-haiku-4-5", "gpt-4o-mini", "gemini-2.0-flash"
}

// CloudProvider é o adapter genérico — a lógica de request/response
// varia por vendor internamente, mas a interface externa é única,
// então trocar de provedor cloud não exige mudança na orquestração.
type CloudProvider struct {
	Creds        CloudCredentials
	SystemPrompt string
	httpClient   *http.Client
}

func NewCloudProvider(creds CloudCredentials, systemPrompt string) *CloudProvider {
	return &CloudProvider{
		Creds:        creds,
		SystemPrompt: systemPrompt,
		httpClient:   &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *CloudProvider) Kind() Kind { return KindCloud }

func (p *CloudProvider) Name() string {
	return fmt.Sprintf("%s (%s) — conta do usuário", p.Creds.Vendor, p.Creds.Model)
}

// Available faz uma checagem leve de credencial. Para a maioria dos
// vendors isso significa uma chamada mínima (ou apenas validar formato
// da chave) — evitar gastar quota do usuário só para health-check.
func (p *CloudProvider) Available(ctx context.Context) error {
	if p.Creds.APIKey == "" {
		return fmt.Errorf("%w: chave de API não configurada para %s", ErrUnavailable, p.Creds.Vendor)
	}
	return nil
}

// Analyze despacha para o vendor correto. Cada branch é isolado para
// que adicionar um quarto provedor não exija tocar nos outros.
func (p *CloudProvider) Analyze(ctx context.Context, pkg EvidencePackage) (Verdict, error) {
	userPrompt, err := buildEvidencePrompt(pkg)
	if err != nil {
		return Verdict{}, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}

	var raw []byte
	switch p.Creds.Vendor {
	case VendorAnthropic:
		raw, err = p.callAnthropic(ctx, userPrompt)
	case VendorOpenAI:
		raw, err = p.callOpenAI(ctx, userPrompt)
	case VendorGemini:
		raw, err = p.callGemini(ctx, userPrompt)
	default:
		return Verdict{}, fmt.Errorf("%w: vendor desconhecido %q", ErrUnavailable, p.Creds.Vendor)
	}
	if err != nil {
		return Verdict{}, err
	}

	verdict, err := ValidateVerdict(raw, pkg)
	if err != nil {
		return Verdict{}, err
	}
	verdict.ProviderKind = KindCloud
	verdict.ModelName = p.Creds.Model
	return verdict, nil
}

// --- Anthropic --------------------------------------------------------

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system"`
	Messages  []anthropicMessage `json:"messages"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Error *struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

func (p *CloudProvider) callAnthropic(ctx context.Context, userPrompt string) ([]byte, error) {
	reqBody := anthropicRequest{
		Model:     p.Creds.Model,
		MaxTokens: 2048,
		System:    p.SystemPrompt + "\n\nResponda APENAS com JSON válido, sem markdown, sem preâmbulo.",
		Messages: []anthropicMessage{
			{Role: "user", Content: userPrompt},
		},
	}
	buf, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.Creds.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	body, status, err := p.doRequest(req)
	if err != nil {
		return nil, err
	}

	var out anthropicResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}
	if out.Error != nil {
		if status == http.StatusTooManyRequests || out.Error.Type == "rate_limit_error" {
			return nil, fmt.Errorf("%w: %s", ErrQuotaExceeded, out.Error.Message)
		}
		return nil, fmt.Errorf("%w: %s", ErrUnavailable, out.Error.Message)
	}
	if len(out.Content) == 0 {
		return nil, ErrSchemaInvalid
	}
	return []byte(out.Content[0].Text), nil
}

// --- OpenAI -------------------------------------------------------------

type openAIRequest struct {
	Model          string          `json:"model"`
	Messages       []openAIMessage `json:"messages"`
	ResponseFormat struct {
		Type string `json:"type"`
	} `json:"response_format"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

func (p *CloudProvider) callOpenAI(ctx context.Context, userPrompt string) ([]byte, error) {
	reqBody := openAIRequest{
		Model: p.Creds.Model,
		Messages: []openAIMessage{
			{Role: "system", Content: p.SystemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}
	reqBody.ResponseFormat.Type = "json_object"
	buf, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.Creds.APIKey)

	body, status, err := p.doRequest(req)
	if err != nil {
		return nil, err
	}

	var out openAIResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}
	if out.Error != nil {
		if status == http.StatusTooManyRequests || out.Error.Type == "insufficient_quota" {
			return nil, fmt.Errorf("%w: %s", ErrQuotaExceeded, out.Error.Message)
		}
		return nil, fmt.Errorf("%w: %s", ErrUnavailable, out.Error.Message)
	}
	if len(out.Choices) == 0 {
		return nil, ErrSchemaInvalid
	}
	return []byte(out.Choices[0].Message.Content), nil
}

// --- Gemini ---------------------------------------------------------------

type geminiRequest struct {
	Contents          []geminiContent `json:"contents"`
	SystemInstruction *geminiContent  `json:"systemInstruction,omitempty"`
	GenerationConfig  struct {
		ResponseMimeType string `json:"responseMimeType"`
	} `json:"generationConfig"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content geminiContent `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

func (p *CloudProvider) callGemini(ctx context.Context, userPrompt string) ([]byte, error) {
	reqBody := geminiRequest{
		Contents: []geminiContent{
			{Parts: []geminiPart{{Text: userPrompt}}},
		},
		SystemInstruction: &geminiContent{Parts: []geminiPart{{Text: p.SystemPrompt}}},
	}
	reqBody.GenerationConfig.ResponseMimeType = "application/json"
	buf, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		p.Creds.Model, p.Creds.APIKey,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, status, err := p.doRequest(req)
	if err != nil {
		return nil, err
	}

	var out geminiResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}
	if out.Error != nil {
		if status == http.StatusTooManyRequests || out.Error.Status == "RESOURCE_EXHAUSTED" {
			return nil, fmt.Errorf("%w: %s", ErrQuotaExceeded, out.Error.Message)
		}
		return nil, fmt.Errorf("%w: %s", ErrUnavailable, out.Error.Message)
	}
	if len(out.Candidates) == 0 || len(out.Candidates[0].Content.Parts) == 0 {
		return nil, ErrSchemaInvalid
	}
	return []byte(out.Candidates[0].Content.Parts[0].Text), nil
}

// --- helper compartilhado -------------------------------------------------

func (p *CloudProvider) doRequest(req *http.Request) ([]byte, int, error) {
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
