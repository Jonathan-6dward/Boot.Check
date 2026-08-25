# BootCheck — Registro de validação do scaffold

## Verificações concluídas

| Verificação | Resultado |
| --- | --- |
| `gofmt` nos arquivos Go | OK. |
| `go test ./...` em Linux | OK nos pacotes `collector`, `api`, `report` e `tests`. |
| `GOOS=windows GOARCH=amd64 go build ./...` | OK; o scaffold compila para Windows. |
| Compilação dos binários de teste para Windows | OK com `go test -c`; os binários não foram executados no Linux. |
| Sintaxe JSON | OK para os schemas e três fixtures. |
| Validação formal das três fixtures contra `evidence.schema.json` | OK. |
| Validação formal de fixture de `VerdictResponse` contra `verdict.schema.json` | OK. |

## O que não foi executado

Não houve coleta em endpoint Windows, chamada a provedor LLM real, assinatura Authenticode, teste de permissões, snapshot antes/depois de registro/serviços/tarefas/WMI, nem verificação dinâmica de somente-leitura. Essas etapas pertencem ao agente local e estão descritas no handoff; devem ocorrer em máquina descartável e com aprovação humana antes de qualquer release.

A validação jurídica também não foi realizada. `docs/RESPONSABILIDADE_LEGAL.md` é um rascunho e deve ser revisado por advogado antes de publicação, distribuição ou venda.
