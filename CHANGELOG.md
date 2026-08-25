# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/),
e este projeto adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [0.1.0-scaffold] - 2026-08-25

### Adicionado
- Especificação completa do produto (MVP, persona, jornada, métricas) em `docs/ESPECIFICACAO_PRODUTO.md`
- Arquitetura técnica com decisão proposta Go, schema e prompt 5W2H/MITRE em `docs/ARQUITETURA_TECNICA.md`
- Rascunho de termo de uso, disclaimer e questões LGPD em `docs/RESPONSABILIDADE_LEGAL.md`
- Comparação de três modelos comerciais (compra única, assinatura, freemium) em `docs/MODELO_DE_NEGOCIO.md`
- Handoff para agente de implementação local em `docs/HANDOFF_AGENTE_LOCAL.md`
- Scaffold Go do coletor com stubs para 8 categorias (processos, persistência, tarefas, serviços, WMI, Winlogon, rede, Defender)
- JSON Schema 2020-12 canônico do pacote de evidências em `scaffold/collector/evidence.schema.json`
- Contratos-placeholder de redaction e validação de consentimento em `scaffold/collector/privacy.go`
- Contratos-placeholder de canonicalização SHA-256 em `scaffold/collector/integrity.go`
- Cliente HTTP-placeholder de API LLM com timeout, retry limitado e validação em `scaffold/api/client.go`
- Prompt de sistema/usuário com boundary de dados não confiáveis em `scaffold/api/prompt_template.go`
- JSON Schema do contrato `VerdictResponse` em `scaffold/api/verdict.schema.json`
- Gerador HTML de relatório com seção leiga e apêndice técnico em `scaffold/report/report.go`
- Três fixtures sintéticas (benign, malicious, ambiguous) e testes de forma em `scaffold/tests/`
- Script de validação de schemas em `validate_schemas.py`
- Decisões aprovadas registradas em `docs/DECISOES_APROVADAS.md`
- ADR de provider Local/Cloud + interface `Provider` + adapters Ollama/Anthropic/OpenAI/Gemini em `scaffold/api/provider/`

### Segurança
- Contrato estrito de somente-leitura documentado em `collector.go` e `collector/README.md`
- Guardrails em todo o handoff: nenhuma execução, escrita, elevação ou remediação
- Aviso de "isto não é aconselhamento jurídico" em todos os documentos relevantes

### Pendente (próximas versões)
- [ ] Tarefa 3: Implementar redaction real (substituir TODO em `privacy.go`)
- [ ] Tarefas 4-11: Implementar coleta real por categoria
- [ ] Tarefa 12: Tela de consentimento e UI local
- [ ] Tarefa 13: Orquestração, integridade e pacote canônico
- [ ] Tarefa 14: Adaptador de provedor LLM
- [ ] Tarefa 15: Completar gerador de relatório
- [ ] Tarefa 16: Testes de prompt com mock provider
- [ ] Tarefa 17: Verificação de somente-leitura em Windows
- [ ] Tarefa 18: Build reprodutível + assinatura Authenticode
- [ ] Tarefa 19: Revisão jurídica + release gate

[0.1.0-scaffold]: https://github.com/Jonathan-6dward/Boot.Check/releases/tag/v0.1.0-scaffold
