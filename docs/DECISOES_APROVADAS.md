# Decisões aprovadas

> Este arquivo registra decisões humanas aprovadas para o projeto BootCheck.
> Atualize este documento sempre que uma decisão for tomada, mantendo o
> histórico (não apague decisões anteriores; use tachado ou mova para
> "histórico" quando revertidas).

## Estado atual (v0.1.0-scaffold)

| Decisão | Valor aprovado | Data | Aprovado por | Observação |
| --- | --- | --- | --- | --- |
| **Nome do produto** | **BootCheck** (placeholder mantido) | 2026-08-25 | proprietário | Pesquisa de marca pendente; alternativas: BootLens, Startup Sentinel, InitTrace |
| **Linguagem do coletor** | **Go** (1.22+) | 2026-08-25 | proprietário | Recomendação técnica aceita; build com flags reprodutíveis |
| **Provedor LLM (v0.1)** | **mock provider** (sem provedor real) | 2026-08-25 | proprietário | Cliente aceita interface plugável; provedor real será decidido após Tarefa 14 |
| **Modelo de negócio** | **compra única** (a confirmar) | 2026-08-25 | proprietário | Três opções comparadas em `MODELO_DE_NEGOCIO.md`; decisão final antes do release |
| **Modo padrão de dados** | **redacted** | 2026-08-25 | proprietário | `full` exige opt-in explícito |
| **Política de retenção** | **mínima** (sem retenção em nuvem) | 2026-08-25 | proprietário | Pacote local; nenhum histórico em servidor |
| **Modo de operação** | **local** (sem agente residente) | 2026-08-25 | proprietário | Monitoramento contínuo explicitamente fora de escopo da v1 |
| **Visibilidade do repositório** | **público** | 2026-08-25 | proprietário | Licença MIT; revogável antes do release se necessário |
| **Licença** | **MIT** | 2026-08-25 | proprietário | Veja `LICENSE` |
| **Base legal (LGPD)** | **consentimento explícito** | 2026-08-25 | proprietário | Rascunho em `RESPONSABILIDADE_LEGAL.md`; revisão de advogado antes de publicar |
| **Janela de eventos Defender** | últimos 24h, máximo 1000 eventos | 2026-08-25 | proprietário | Ajuste empírico antes do release |

## Decisões ainda pendentes

| Decisão | Opções | Bloqueia |
| --- | --- | --- |
| Pesquisa de marca do nome | confirmar disponibilidade de "BootCheck" e domínios | release público |
| Provedor/Modelo LLM real | mock / Anthropic / OpenAI / outro | Tarefa 14 |
| Tabela de preços unitários de API | confirmar valores vigentes | `MODELO_DE_NEGOCIO.md` |
| Texto final de consentimento | advogado | Tarefa 12 e release |
| Política de retenção detalhada | advogado + decisão comercial | release |
| Identidade de controlador/operador LGPD | jurídico | release |

## Histórico

_(nenhuma decisão revertida até o momento)_

## Como registrar uma nova decisão

1. Adicione uma linha na tabela **Estado atual** com data e aprovador
2. Se a decisão for **pendente** e ainda não tomada, adicione em **Decisões pendentes**
3. Se a decisão **reverter** uma anterior, mova a linha antiga para **Histórico**
4. **Não apague** histórico de decisões — isso é auditoria
