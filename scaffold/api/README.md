# API — Cliente LLM (scaffold)

Este pacote contém o **cliente placeholder** de API para o provedor LLM. Na
v0.1.0-scaffold, nenhum provedor real é chamado; um **mock provider** é usado
para validar o contrato de prompt e o `VerdictResponse`.

## Componentes

| Arquivo | Função |
| ------- | ------ |
| `client.go` | Cliente HTTP com timeout, retry limitado, validação de resposta |
| `prompt_template.go` | System prompt + user prompt com boundary de dados |
| `schema.go` | Validação de `VerdictResponse` |
| `verdict.schema.json` | JSON Schema 2020-12 do contrato de saída |
| `client_test.go` | Testes de forma e timeout |

## Princípios

1. **Dados não confiáveis**: o conteúdo de `EvidencePackage` é **dado**, nunca
   instrução. O prompt tem regra explícita para ignorar tentativas de override.
2. **Sem log de payload**: o cliente nunca registra o JSON completo, tokens ou
   linhas de comando.
3. **Timeout e retry limitado**: configuráveis, com teto documentado.
4. **Validação estrita**: resposta fora do schema → `inconclusive` técnico, sem
   fabricar veredito.
5. **Sem execução**: este pacote **não** executa código da resposta.

## Provedor

A v0.1.0-scaffold usa um **mock provider** para que:

- Os testes de prompt não precisem de chave de API
- Nenhum dado saia da máquina durante validação
- O contrato possa ser testado repetidamente

Quando o provedor real for aprovado, o `client.go` será adaptado na **Tarefa 14**
do handoff, mantendo a mesma interface `LLMClient`.

## Como rodar (após Tarefa 14)

```bash
# Configurar provedor
export BOOTCHECK_LLM_PROVIDER=anthropic
export BOOTCHECK_LLM_API_KEY=...
export BOOTCHECK_LLM_MODEL=...

# Executar
cd scaffold
go test ./api/...
```

## Aviso

> O uso de provedor LLM real **exige** consentimento explícito do usuário e
> revisão jurídica de `docs/RESPONSABILIDADE_LEGAL.md`. Nunca transmita o
> pacote sem o consentimento registrado.
