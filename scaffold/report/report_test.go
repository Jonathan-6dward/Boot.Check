package report

import (
	"strings"
	"testing"

	"github.com/Jonathan-6dward/Boot.Check/scaffold/api"
	"github.com/Jonathan-6dward/Boot.Check/scaffold/collector"
)

func TestRenderHTMLContainsLaypersonAndTechnicalSections(t *testing.T) {
	result := api.VerdictResponse{
		Verdict:              "inconclusive",
		Confidence:           0.4,
		HeadlinePlain:        "Os dados não bastam.",
		PlainLanguageSummary: "Resumo sintético.",
		SafetyNotice:         "Triagem indicativa.",
	}
	var builder strings.Builder
	if err := RenderHTML(&builder, result, collector.EvidencePackage{}); err != nil {
		t.Fatal(err)
	}
	html := builder.String()
	for _, fragment := range []string{"id=\"leigo\"", "id=\"tecnico\"", "falsos positivos", "Resumo sintético."} {
		if !strings.Contains(html, fragment) {
			t.Errorf("report does not contain %q", fragment)
		}
	}
}

func TestRenderHTMLEscapesModelText(t *testing.T) {
	result := api.VerdictResponse{
		Verdict:              "inconclusive",
		Confidence:           0.4,
		HeadlinePlain:        "<script>fixture</script>",
		PlainLanguageSummary: "Texto de teste.",
		SafetyNotice:         "Triagem indicativa.",
	}
	var builder strings.Builder
	if err := RenderHTML(&builder, result, collector.EvidencePackage{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(builder.String(), "<script>fixture</script>") {
		t.Fatal("model text was not escaped")
	}
}
