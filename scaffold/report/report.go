package report

import (
	"encoding/json"
	"html/template"
	"io"
	"time"

	"github.com/Jonathan-6dward/Boot.Check/scaffold/api"
)

type ReportData struct {
	Title       string
	GeneratedAt time.Time
	Result      api.VerdictResponse
}

// RenderHTML renders a local, self-contained report. It is intentionally a
// pure formatting function: it does not call the collector, the network or a
// remediation command.
// TODO(local-agent): add report provenance, collection_id and package hash to
// the header once the collector contract is wired in.
func RenderHTML(w io.Writer, result api.VerdictResponse) error {
	if err := api.ValidateVerdict(result); err != nil {
		return err
	}
	tmpl, err := template.New("bootcheck-report").Funcs(template.FuncMap{
		"json": func(value any) string {
			encoded, marshalErr := json.MarshalIndent(value, "", "  ")
			if marshalErr != nil {
				return "<dados técnicos indisponíveis>"
			}
			return string(encoded)
		},
		"mul100": func(value float64) float64 { return value * 100 },
	}).Parse(htmlTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, ReportData{Title: "BootCheck — Relatório de triagem", GeneratedAt: time.Now().UTC(), Result: result})
}

const htmlTemplate = `<!doctype html>
<html lang="pt-BR">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>
    :root { color-scheme: light; font-family: system-ui, sans-serif; }
    body { max-width: 960px; margin: 0 auto; padding: 2rem; line-height: 1.5; color: #172033; }
    header { border-bottom: 2px solid #d8dee9; margin-bottom: 2rem; }
    .badge { display: inline-block; padding: .35rem .65rem; border-radius: .45rem; font-weight: 700; background: #e8eef9; }
    .notice { padding: 1rem; background: #fff4d6; border-left: 4px solid #c98a00; }
    section { margin: 2rem 0; }
    pre { overflow-x: auto; padding: 1rem; background: #f5f7fa; border-radius: .4rem; }
    dt { font-weight: 700; margin-top: .6rem; }
    dd { margin-left: 0; }
  </style>
</head>
<body>
  <header>
    <h1>{{.Title}}</h1>
    <p>Gerado em {{.GeneratedAt}}</p>
    <p><span class="badge">{{.Result.Verdict}}</span> Confiança de triagem: {{printf "%.0f" (mul100 .Result.Confidence)}}%</p>
  </header>

  <div class="notice"><strong>Aviso:</strong> esta é uma triagem automatizada e pode conter falsos positivos, falsos negativos ou ficar inconclusiva. Não execute, encerre, remova, bloqueie, desative ou altere nada com base apenas neste relatório. Isto não é aconselhamento jurídico.</div>

  <section id="leigo">
    <h2>Para você</h2>
    <h3>{{.Result.HeadlinePlain}}</h3>
    <p>{{.Result.PlainLanguageSummary}}</p>
    <h3>O que fazer agora</h3>
    <ul>{{range .Result.RecommendedNextSteps}}<li>{{.}}</li>{{end}}</ul>
  </section>

  <section id="tecnico">
    <h2>Apêndice técnico</h2>
    <h3>5W2H</h3>
    <dl>
      <dt>O quê</dt><dd>{{.Result.FiveWTwoH.What}}</dd>
      <dt>Quem</dt><dd>{{.Result.FiveWTwoH.Who}}</dd>
      <dt>Quando</dt><dd>{{.Result.FiveWTwoH.When}}</dd>
      <dt>Onde</dt><dd>{{.Result.FiveWTwoH.Where}}</dd>
      <dt>Por quê</dt><dd>{{.Result.FiveWTwoH.Why}}</dd>
      <dt>Como</dt><dd>{{.Result.FiveWTwoH.How}}</dd>
      <dt>Impacto observado</dt><dd>{{.Result.FiveWTwoH.Impact}}</dd>
    </dl>
    <h3>Mapeamentos MITRE ATT&amp;CK</h3>
    {{if .Result.MITRE}}<ul>{{range .Result.MITRE}}<li><strong>{{.TechniqueID}} — {{.Name}}</strong>: {{.Rationale}} (evidências: {{.EvidenceIDs}})</li>{{end}}</ul>{{else}}<p>Nenhum mapeamento foi atribuído.</p>{{end}}
    <h3>Evidências e limitações</h3>
    <ul>{{range .Result.SupportingEvidence}}<li>{{.EvidenceID}} — {{.Role}}: {{.Explanation}}</li>{{end}}</ul>
    <p>Limitações: {{.Result.Limitations}}</p>
    <h3>Dados técnicos</h3>
    <pre>{{json .Result.TechnicalAppendix}}</pre>
  </section>

  <footer><p>{{.Result.SafetyNotice}}</p></footer>
</body>
</html>`
