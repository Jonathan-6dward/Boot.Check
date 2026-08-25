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

// LocalOllamaProvider fala com um runtime Ollama rodando em localhost.
// Nenhum dado sai da máquina do usuário — por isso este provider NÃO
// exige a tela de consentimento de rede do README (só o consentimento
// de execução da análise em si, que é local a qualquer forma).
type LocalOllamaProvider struct {
	// BaseURL normalmente "http://localhost:11434"
	BaseURL string
	// Model é o nome do modelo já puxado no Ollama, ex. "qwen2.5:7b-instruct"
	Model string
	// SystemPrompt é o template 5W2H/MITRE já existente em
	// scaffold/api/prompt_template.go — reaproveitado tal qual.
	SystemPrompt string

	httpClient *http.Client
}

// NewLocalOllamaProvider cria o adapter com timeout padrão de 90s
// (modelos locais em CPU podem ser lentos; cloud costuma ser mais rápido).
func NewLocalOllamaProvider(baseURL, model, systemPrompt string) *LocalOllamaProvider {
	return &LocalOllamaProvider{
		BaseURL:      baseURL,
		Model:        model,
		SystemPrompt: systemPrompt,
		httpClient:   &http.Client{Timeout: 90 * time.Second},
	}
}

func (p *LocalOllamaProvider) Kind() Kind { return KindLocal }

func (p *LocalOllamaProvider) Name() string {
	return fmt.Sprintf("Ollama local (%s)", p.Model)
}

// Available faz um GET /api/tags — se o daemon não responder em 2s,
// consideramos indisponível e a UI oferece instrução de instalação
// ou cai para o fluxo sem veredito de IA (RelatorioLocal).
func (p *LocalOllamaProvider) Available(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.BaseURL+"/api/tags", nil)
	if err != nil {
		return err
	}
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: status %d", ErrUnavailable, resp.StatusCode)
	}
	return nil
}

type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Format   string          `json:"format"` // "json" força saída estruturada
	Stream   bool            `json:"stream"`
}

type ollamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ollamaChatResponse struct {
	Message ollamaMessage `json:"message"`
	Done    bool          `json:"done"`
}

// Analyze monta o prompt isolando cada Evidence.Value como dado citado
// por evidence_id (nunca concatenado como instrução), chama
// /api/chat com format:"json" e valida a resposta contra o schema.
func (p *LocalOllamaProvider) Analyze(ctx context.Context, pkg EvidencePackage) (Verdict, error) {
	userPrompt, err := buildEvidencePrompt(pkg)
	if err != nil {
		return Verdict{}, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}

	reqBody := ollamaChatRequest{
		Model: p.Model,
		Messages: []ollamaMessage{
			{Role: "system", Content: p.SystemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Format: "json",
		Stream: false,
	}
	buf, err := json.Marshal(reqBody)
	if err != nil {
		return Verdict{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.BaseURL+"/api/chat", bytes.NewReader(buf))
	if err != nil {
		return Verdict{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return Verdict{}, fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Verdict{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Verdict{}, fmt.Errorf("%w: status %d body=%s", ErrUnavailable, resp.StatusCode, string(body))
	}

	var out ollamaChatResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return Verdict{}, fmt.Errorf("%w: %v", ErrSchemaInvalid, err)
	}

	verdict, err := ValidateVerdict([]byte(out.Message.Content), pkg)
	if err != nil {
		// Falha de schema em modelo local é esperada com mais frequência
		// que em cloud — a camada de orquestração deve tratar isso como
		// Inconclusivo, não como crash.
		return Verdict{}, err
	}
	verdict.ProviderKind = KindLocal
	verdict.ModelName = p.Model
	return verdict, nil
}
