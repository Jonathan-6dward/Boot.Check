// Package provider define a abstração de "Analysis provider" do BootCheck.
//
// Objetivo: permitir que o mesmo EvidencePackage seja analisado por um
// provedor LOCAL (LLM rodando na máquina do usuário, sem sair da rede)
// ou por um provedor CLOUD (conta própria do usuário, camada gratuita
// ou paga, chave de API dele, nunca da infraestrutura do BootCheck).
//
// Nenhum adapter aqui decide política de consentimento — isso é
// responsabilidade da camada de orquestração (Previa -> Consentimento
// no state diagram do README). O provider só recebe um EvidencePackage
// já normalizado/redigido e devolve um Verdict ou um erro.
package provider

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

// ErrSchemaInvalid é retornado quando a resposta do modelo não bate
// com o schema de veredito. A camada acima DEVE tratar isso como
// Inconclusivo, nunca como erro fatal de UI.
var ErrSchemaInvalid = errors.New("provider: resposta fora do schema de veredito")

// ErrUnavailable é retornado quando o provedor não está acessível
// (Ollama não rodando, endpoint cloud sem rede, etc).
var ErrUnavailable = errors.New("provider: indisponível")

// ErrQuotaExceeded é retornado quando o provedor cloud reporta limite
// de camada gratuita/paga excedido. Também deve cair em Inconclusivo,
// com mensagem específica para o usuário.
var ErrQuotaExceeded = errors.New("provider: cota do provedor excedida")

// Kind identifica a categoria do provider, usada só para telemetria
// local e para a tela de consentimento saber o que exibir.
type Kind string

const (
	KindLocal Kind = "local" // roda na máquina do usuário, sem rede
	KindCloud Kind = "cloud" // conta própria do usuário, requer consentimento de rede
)

// EvidenceLevel espelha os valores do schema de evidência do README:
// OBSERVED, INFERRED, UNAVAILABLE, WITHHELD, UNKNOWN.
type EvidenceLevel string

const (
	LevelObserved    EvidenceLevel = "OBSERVED"
	LevelInferred    EvidenceLevel = "INFERRED"
	LevelUnavailable EvidenceLevel = "UNAVAILABLE"
	LevelWithheld    EvidenceLevel = "WITHHELD"
	LevelUnknown     EvidenceLevel = "UNKNOWN"
)

// Evidence é um item individual do pacote de evidências. Os campos de
// texto livre (Value) são a superfície de ataque de prompt injection:
// vêm de cmdline, paths, nomes de tarefa etc. do sistema investigado,
// e por isso NUNCA devem ser interpolados como instrução — só como
// dado dentro de um bloco delimitado e citado por ID.
type Evidence struct {
	ID    string        `json:"evidence_id"`
	Kind  string        `json:"kind"` // process | run_key | scheduled_task | service | wmi | winlogon | network | defender
	Level EvidenceLevel `json:"level"`
	Value string        `json:"value"`
}

// EvidencePackage é a entrada padrão para qualquer Provider.
type EvidencePackage struct {
	SchemaVersion string     `json:"schema_version"`
	CollectedAt   time.Time  `json:"collected_at"`
	Items         []Evidence `json:"items"`
}

// VerdictState espelha os três estados definidos no README.
type VerdictState string

const (
	StateLikelySafe   VerdictState = "likely_safe"
	StateSuspicious   VerdictState = "suspicious"
	StateInconclusive VerdictState = "inconclusive"
)

// Verdict é a saída estruturada 5W2H + MITRE que o relatório consome.
// Cada Claim deve apontar para EvidenceIDs existentes no pacote — a
// camada de report/renderer é responsável por rejeitar claims que
// citam evidence_id inexistente.
type Verdict struct {
	State        VerdictState `json:"state"`
	Summary      string       `json:"summary"`      // linguagem leiga
	Claims       []Claim      `json:"claims"`       // apêndice técnico
	MitreATTACK  []string     `json:"mitre_attack"` // vazio se não sustentado por evidência
	ProviderKind Kind         `json:"provider_kind"`
	ModelName    string       `json:"model_name"`
}

// Claim é uma afirmação decisiva do apêndice técnico, sempre rastreável.
type Claim struct {
	What        string   `json:"what"`
	Who         string   `json:"who"`
	When        string   `json:"when"`
	Where       string   `json:"where"`
	Why         string   `json:"why"`
	How         string   `json:"how"`
	Impact      string   `json:"impact"`
	EvidenceIDs []string `json:"evidence_ids"`
}

// Provider é a interface que Local e Cloud implementam.
type Provider interface {
	// Kind identifica local vs cloud, usado pela UI de consentimento.
	Kind() Kind

	// Name é o nome legível exibido ao usuário (ex.: "Ollama (qwen2.5:7b)",
	// "Anthropic (claude-haiku-4-5)").
	Name() string

	// Available faz um health-check rápido (sem enviar evidência real).
	// Deve retornar rapidamente (< 2s) para não travar a UI de setup.
	Available(ctx context.Context) error

	// Analyze recebe o pacote já normalizado/redigido e devolve um
	// Verdict validado contra o schema, ou um erro dos tipos Err*
	// acima. Implementações NÃO devem logar o conteúdo de Evidence.Value
	// (pode conter dado pessoal) — só evidence_id e nível.
	Analyze(ctx context.Context, pkg EvidencePackage) (Verdict, error)
}

// ValidateVerdict faz a checagem mínima de schema antes de qualquer
// Provider devolver um Verdict como válido. Adapters devem chamar isso
// internamente e mapear falha para ErrSchemaInvalid.
func ValidateVerdict(raw []byte, pkg EvidencePackage) (Verdict, error) {
	var v Verdict
	if err := json.Unmarshal(raw, &v); err != nil {
		return Verdict{}, ErrSchemaInvalid
	}
	switch v.State {
	case StateLikelySafe, StateSuspicious, StateInconclusive:
	default:
		return Verdict{}, ErrSchemaInvalid
	}
	known := make(map[string]bool, len(pkg.Items))
	for _, item := range pkg.Items {
		known[item.ID] = true
	}
	for _, c := range v.Claims {
		for _, id := range c.EvidenceIDs {
			if !known[id] {
				// claim aponta pra evidence_id que não existe no pacote:
				// trata como schema inválido, não como veredito "quase bom".
				return Verdict{}, ErrSchemaInvalid
			}
		}
	}
	return v, nil
}
