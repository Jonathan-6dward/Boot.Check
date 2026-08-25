# Contributing to BootCheck

Obrigado por considerar contribuir com o **BootCheck**! Este projeto é uma
ferramenta defensiva e somente-leitura de triagem forense para Windows. Toda
contribuição deve respeitar o **contrato de somente-leitura** descrito em
`scaffold/collector/README.md` e `docs/ARQUITETURA_TECNICA.md`.

## 🚨 Regra de ouro

> Se uma contribuição parecer exigir **escrita**, **execução**, **elevação**,
> **persistência** ou **remediação**, **PARE** e peça aprovação humana explícita
> antes de submeter.

Isto não é negociável.

## Como contribuir

1. **Faça um fork** do repositório
2. **Crie uma branch** a partir de `main`:
   ```bash
   git checkout -b feat/sua-contribuicao
   ```
3. **Implemente mudanças pequenas, revisáveis** com testes
4. **Valide os schemas**:
   ```bash
   pip install jsonschema
   python validate_schemas.py
   ```
5. **Execute os testes Go**:
   ```bash
   cd scaffold
   go test ./...
   ```
6. **Abra um Pull Request** com descrição clara

## Padrões de código

- **Go:** siga `gofmt`, `go vet`, `staticcheck`
- **JSON:** indentação de 2 espaços, UTF-8, sem BOM
- **Markdown:** uma linha em branco entre seções
- **Commits:** mensagens claras,Conventional Commits quando possível
  (`feat:`, `fix:`, `docs:`, `test:`, `chore:`)

## Categorias de contribuição

| Tipo | Onde |
|------|------|
| Implementar coleta real | `scaffold/collector/collect_*.go` |
| Corrigir schema | `scaffold/collector/evidence.schema.json` |
| Adicionar fixture de teste | `scaffold/tests/` |
| Melhorar prompt LLM | `scaffold/api/prompt_template.go` |
| Melhorar relatório HTML | `scaffold/report/report.go` |
| Documentação | `docs/` (não altere `RESPONSABILIDADE_LEGAL.md` sem advogado) |

## O que NÃO fazer

- ❌ Adicionar execução de payload, mesmo em testes
- ❌ Implementar quarentena, remoção ou remediação automática
- ❌ Tentar contornar antivírus/EDR
- ❌ Coletar dados de máquina real em PRs públicos (use fixtures)
- ❌ Commitar chaves, tokens ou dados pessoais
- ❌ Alterar `docs/RESPONSABILIDADE_LEGAL.md` sem revisão de advogado
- ❌ Enviar PRs que tratem "não_accessible" como "seguro"

## Reportar vulnerabilidades

Abra uma **issue privada** com a tag `security` ou entre em contato diretamente
com os mantenedores. **Não** divulgue publicamente até correção.

## Código de Conduta

Este projeto adota um código de conduta baseado em respeito mútuo. Comportamento
abusivo não é tolerado.

## Licença

Ao contribuir, você concorda que suas contribuições serão licenciadas sob a
**MIT License** (veja `LICENSE`).
