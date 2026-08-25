# Testes — BootCheck

Esta pasta contém **fixtures sintéticas** e testes de forma. **Não use dados
reais** de endpoint ou máquina de produção.

## Estrutura

| Arquivo | Conteúdo | Veredito esperado |
| ------- | -------- | ----------------- |
| `benign_obvious.json` | Vendor legítimo com tarefa agendada conhecida | `likely_safe` |
| `malicious_obvious.json` | PowerShell não-assinado em `C:\Users\Public` + Run + WMI + Defender | `suspicious` |
| `ambiguous.json` | Processo com assinatura desconhecida + acesso negado em tarefas/WMI | `inconclusive` |

## Como rodar

### Validação de schema

```bash
pip install jsonschema
python ../../validate_schemas.py
```

Saída esperada (a partir de v0.1.0-scaffold):

```text
evidence schema OK: ambiguous.json
evidence schema OK: benign_obvious.json
evidence schema OK: malicious_obvious.json
verdict schema OK
```

### Testes Go

```bash
cd ../
go test ./...
```

## Adicionando novos fixtures

1. Crie um novo `*.json` na raiz desta pasta
2. Use o envelope completo (todos os campos obrigatórios do schema)
3. Use `evidence_id` no formato `^ev-[A-Za-z0-9_-]{8,64}$`
4. Defina `redaction_status: "redacted"` como padrão
5. Use `HOST-REDACTED` e `USER-REDACTED` para hostname/usuário
6. Não use IPs de produção reais — use `192.0.2.0/24` (TEST-NET-1) ou `198.51.100.0/24` (TEST-NET-2)
7. Não use hashes reais de binários
8. Adicione um teste em `fixtures_test.go` se for um caso novo

## Aviso

> Estes fixtures são **fictícios** e não foram capturados de endpoints reais.
> Usar dados reais em testes pode expor informações pessoais e ferir LGPD.
