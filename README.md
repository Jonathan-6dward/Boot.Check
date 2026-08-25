# BootCheck

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-v0.1.0--scaffold-blue.svg)](CHANGELOG.md)
[![Status](https://img.shields.io/badge/status-scaffold-orange.svg)](CHANGELOG.md)
[![Read-Only](https://img.shields.io/badge/contrato-somente--leituda-green.svg)](scaffold/collector/README.md)

> **Ferramenta defensiva e somente-leitura de triagem forense para Windows.**
>
> v0.1.0-scaffold: especificação + scaffold Go. **Nenhuma coleta real é executada**
> nesta versão. Os stubs retornam `TODO` deliberadamente para impedir uso indevido.

## 🎯 O que é o BootCheck?

O BootCheck responde a **uma única pergunta**:

> *"Por que janelas de `cmd.exe` ou PowerShell aparecem sozinhas durante a
> inicialização do Windows, e isso é malicioso ou legítimo?"*

É uma ferramenta de **triagem**, não de diagnóstico. O relatório final separa:

- 🟢 **Provavelmente seguro** — sinais compatíveis com atividade legítima
- 🟡 **Suspeito** — sinais que merecem investigação
- ⚪ **Inconclusivo** — dados insuficientes ou conflitantes

## 🛡️ Contrato de somente-leitura

Esta ferramenta **NUNCA**:

- ❌ Executa comandos ou payloads
- ❌ Encerra ou modifica processos
- ❌ Altera registro, serviços, tarefas agendadas
- ❌ Cria ou remove WMI subscriptions
- ❌ Limpa logs de eventos
- ❌ Modifica firewall ou Defender
- ❌ Envia dados sem consentimento explícito

Permitido:

- ✅ Ler processos, registro, tarefas, serviços, WMI, Winlogon, rede, eventos Defender
- ✅ Gerar pacote JSON local com evidências
- ✅ Enviar pacote a LLM **somente** com consentimento explícito e modo `redacted` por padrão
- ✅ Gerar relatório HTML local com seção leiga e apêndice técnico

## 📦 O que esta versão entrega

Esta é a versão **v0.1.0-scaffold**. Contém:

| Camada | Conteúdo | Estado |
| ------ | -------- | ------ |
| **Documentação** | Especificação, arquitetura, modelo de negócio, LGPD, handoff | ✅ Completo |
| **Scaffold Go** | 8 stubs de coleta, cliente LLM, gerador de relatório, ADR + adapters de provider | ✅ Pronto |
| **Schemas JSON** | EvidencePackage, VerdictResponse (Draft 2020-12) | ✅ Validados |
| **Fixtures** | benign, malicious, ambiguous + testes de forma | ✅ 100% válidos |
| **CI** | GitHub Actions (validate-schemas, go test, gofmt) | ✅ Configurado |
| **Decisões** | Nome, linguagem, licença, modo, provedor LLM | ✅ Registradas |

**Não** contém (próximas versões):

- ⏳ Coleta real em Windows (Tarefas 4-11)
- ⏳ Tela de consentimento e UI (Tarefa 12)
- ⏳ Build assinado (Tarefa 18)

## 🚀 Início rápido

### Pré-requisitos

- Go 1.22+
- Python 3.8+ com `jsonschema`

### Validação de schemas

```bash
pip install jsonschema
python validate_schemas.py
```

Saída esperada:

```text
evidence schema OK: ambiguous.json
evidence schema OK: benign_obvious.json
evidence schema OK: malicious_obvious.json
verdict schema OK
```

### Testes Go

```bash
cd scaffold
go test ./...
```

## 📁 Estrutura

```
bootcheck/
├── README.md                        ← este arquivo
├── CHANGELOG.md                     ← histórico de versões
├── CONTRIBUTING.md                  ← como contribuir
├── LICENSE                          ← MIT
├── COMO_SUBIR_NO_GITHUB.md          ← guia de publicação
├── docs/                            ← 6 documentos de especificação
│   ├── ESPECIFICACAO_PRODUTO.md
│   ├── ARQUITETURA_TECNICA.md
│   ├── RESPONSABILIDADE_LEGAL.md    ← ⚖️ RASCUNHO, precisa de advogado
│   ├── MODELO_DE_NEGOCIO.md
│   ├── HANDOFF_AGENTE_LOCAL.md      ← prompt para agente de implementação
│   └── DECISOES_APROVADAS.md
├── scaffold/                        ← código Go
│   ├── go.mod
│   ├── collector/                   ← coletor Windows (somente-leitura)
│   │   ├── types.go
│   │   ├── collector.go             ← contrato de segurança
│   │   ├── collect_*.go             ← 8 stubs
│   │   ├── privacy.go               ← redaction
│   │   ├── integrity.go             ← SHA-256 + canonicalization
│   │   └── evidence.schema.json
│   ├── api/                         ← cliente LLM
│   │   ├── client.go
│   │   ├── prompt_template.go       ← prompt 5W2H/MITRE
│   │   ├── schema.go
│   │   ├── verdict.schema.json
│   │   ├── client_test.go
│   │   └── provider/                ← ADR + adapters Local (Ollama) e Cloud
│   │       ├── provider.go          ← interface Provider
│   │       ├── local_ollama.go
│   │       ├── cloud.go             ← Anthropic / OpenAI / Gemini
│   │       ├── prompt.go            ← isolamento de evidência como dado
│   │       ├── provider_test.go
│   │       └── ADR_PROVIDER_LOCAL_E_CLOUD.md
│   ├── report/                      ← gerador HTML
│   │   ├── report.go
│   │   └── report_test.go
│   └── tests/                       ← 3 fixtures sintéticas
│       ├── benign_obvious.json
│       ├── malicious_obvious.json
│       ├── ambiguous.json
│       └── fixtures_test.go
├── validate_schemas.py              ← validador
├── .github/
│   ├── workflows/ci.yml             ← GitHub Actions
│   ├── ISSUE_TEMPLATE/              ← bug_report.md, feature_request.md
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── SECURITY.md
└── .gitignore
```

## 🛣️ Roadmap

| Versão | Escopo | Status |
| ------ | ------ | ------ |
| **v0.1.0-scaffold** | Especificação + scaffold + schemas + fixtures | ✅ **Atual** |
| v0.2.0 | Redaction determinística + Tarefa 3 | ⏳ |
| v0.3.0 | Coleta de processos e árvore pai/filho | ⏳ |
| v0.4.0 | Coleta de persistência (Run/RunOnce) | ⏳ |
| v0.5.0 | Coleta de Task Scheduler | ⏳ |
| v0.6.0 | Coleta de serviços | ⏳ |
| v0.7.0 | Coleta de WMI | ⏳ |
| v0.8.0 | Coleta de Winlogon | ⏳ |
| v0.9.0 | Coleta de rede | ⏳ |
| v0.10.0 | Coleta de eventos Defender | ⏳ |
| v0.11.0 | Orquestração + integridade | ⏳ |
| v0.12.0 | Tela de consentimento + UI | ⏳ |
| v1.0.0-rc1 | LLM real + relatório + build + assinatura | ⏳ |
| v1.0.0 | **Release público** (após revisão jurídica) | ⏳ |

## 🤝 Como contribuir

Veja [CONTRIBUTING.md](CONTRIBUTING.md).

**Regra de ouro:** se uma contribuição parecer exigir escrita, execução,
elevação, persistência ou remediação, **PARE** e peça aprovação humana.

## 🔒 Segurança

Vulnerabilidades devem ser reportadas de forma privada. Veja
[SECURITY.md](.github/SECURITY.md).

**Não** abra issue pública para falhas de segurança.

## ⚖️ Aviso legal

> Este software é uma ferramenta de triagem defensiva e somente-leitura.
> O relatório é uma avaliação automatizada baseada nas evidências disponíveis
> naquele momento; pode conter falsos positivos, falsos negativos ou ficar
> inconclusivo. **Não execute, remova ou altere nada com base apenas neste
> relatório.** Em caso de suspeita relevante, procure um profissional
> autorizado. **Isto não é aconselhamento jurídico**; o produto e seus
> textos devem ser revisados por advogado antes da publicação.

O arquivo `docs/RESPONSABILIDADE_LEGAL.md` está marcado como **RASCUNHO**.

## 📜 Licença

[MIT](LICENSE) — veja também o disclaimer adicional ao final do arquivo.

## 🙏 Agradecimentos

- Microsoft Learn (Task Scheduler, iphlpapi)
- ANPD (Guia LGPD para pequenos agentes)
- Comunidade open source de segurança defensiva

---

**v0.1.0-scaffold** · Pronto para subir no GitHub · Aguardando implementação real
