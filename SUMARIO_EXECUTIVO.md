# BootCheck — Sumário executivo

> **Status:** pacote de projeto quase pronto para implementação local. Esta entrega é uma especificação e scaffold; não executa coleta Windows real. A ferramenta projetada é defensiva, somente-leitura e destinada ao dono da máquina ou a profissional autorizado. **Isto não é aconselhamento jurídico; o texto legal deve ser revisado por advogado antes da publicação.**

## O que foi entregue

Foi criado um projeto único em `/home/ubuntu/bootcheck` contendo a especificação do MVP e persona em `docs/ESPECIFICACAO_PRODUTO.md`, a arquitetura com fluxo, decisão de linguagem, schema e prompt de veredito em `docs/ARQUITETURA_TECNICA.md`, o rascunho de termo/disclaimer/LGPD em `docs/RESPONSABILIDADE_LEGAL.md`, e a comparação das três opções comerciais com custos unitários de API em `docs/MODELO_DE_NEGOCIO.md`.

O diretório `scaffold/` contém o esqueleto Go do coletor com stubs para processos, Run/RunOnce, Task Scheduler, serviços, WMI, Winlogon, rede e Defender; o JSON Schema canônico do pacote de evidências; contratos de redaction e integridade; cliente de API com prompt 5W2H/MITRE embutido; schema do veredito; gerador HTML com seção leiga e apêndice técnico; e três fixtures sintéticas — benign, maliciosa e ambígua — acompanhadas de testes. Os stubs não fazem coleta real e retornam TODOs deliberados para evitar uso indevido.

O arquivo mais importante para continuidade, `docs/HANDOFF_AGENTE_LOCAL.md`, é um prompt autocontido para Claude Code/Cursor. Ele aponta para os artefatos existentes, ordena as tarefas restantes em sessões pequenas, define critérios de aceite e repete o guardrail de somente-leitura em cada etapa relevante.

## Decisões pendentes de aprovação

| Decisão | Estado atual |
| --- | --- |
| Nome do produto | “BootCheck” é placeholder; alternativas foram registradas, sem pesquisa de marca. |
| Linguagem do coletor | Go é a recomendação técnica atual; ainda precisa de aprovação do proprietário. |
| Provedor/modelo do LLM | Não fixado no código. O modelo de negócio usa preços de referência para `gpt-5.6-luna`, `gpt-5.6-terra` e `gpt-5.6-sol`, consultados em 25/08/2026; confirmar tabela vigente antes da contratação. |
| Modelo de negócio | Não escolhido: compra única, assinatura mensal e freemium foram comparados. |
| LGPD e responsabilidade | Rascunho sujeito a advogado: controlador/operador, base legal, consentimento, retenção, transferências e contrato do provedor precisam de validação. |
| Retenção e histórico | Proposta de retenção mínima e modo local, mas prazo e operação final ainda não foram aprovados. |

## Próximo passo imediato recomendado

Escolher e registrar as decisões de **nome, linguagem do coletor, provedor/modelo e modelo de negócio**, sem expandir o sintoma da v1. Em seguida, abrir `docs/HANDOFF_AGENTE_LOCAL.md` em uma sessão nova do Claude Code/Cursor e executar primeiro a validação de schema e os três casos sintéticos, usando mock provider e sem coletar uma máquina real. Paralelamente, encaminhar `docs/RESPONSABILIDADE_LEGAL.md` para revisão jurídica antes de qualquer venda, distribuição pública ou envio de dados a um provedor externo.
