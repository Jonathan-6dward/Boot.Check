BOOTCHECK

<p align="center">
<strong>WINDOWS BOOT ANOMALY TRIAGE</strong>  

  <sub>Observe first. Modify nothing.</sub>
</p> <p align="center">
  <a href="#validar-em-5-minutos"><img src="https://img.shields.io/badge/VALIDATION-READY-2ea043?style=flat-square&labelColor=111827" alt="Validation ready"></a>
  <a href="#modelo-de-segurança"><img src="https://img.shields.io/badge/MODE-READ--ONLY-2ea043?style=flat-square&labelColor=111827" alt="Read-only"></a>
  <a href="#status-real-do-repositório"><img src="https://img.shields.io/badge/TARGET-WINDOWS-2563eb?style=flat-square&labelColor=111827" alt="Windows"></a>
  <a href="#o-que-esta-pronto-e-o-que-nao-esta"><img src="https://img.shields.io/badge/PRODUCT-SCAFFOLD-f59e0b?style=flat-square&labelColor=111827" alt="Scaffold"></a>
</p> <p align="center">
  <a href="#o-problema">Problema</a> ·
  <a href="#como-funciona">Como funciona</a> ·
  <a href="#validar-em-5-minutos">Validar agora</a> ·
  <a href="#arquitetura">Arquitetura</a> ·
  <a href="#segurança-e-privacidade">Segurança</a> ·
  <a href="#roadmap">Roadmap</a>
</p>




Product signal

Plain Text


┌──────────────────────────────────────────────────────────────────┐
│  BOOTCHECK                                                       │
│  WINDOWS BOOT ANOMALY TRIAGE                                     │
│                                                                  │
│  READ-ONLY  ·  LOCAL-FIRST  ·  EVIDENCE-DRIVEN                  │
│                                                                  │
│  OBSERVE  →  CORRELATE  →  CLASSIFY  →  EXPLAIN                 │
└──────────────────────────────────────────────────────────────────┘



O BootCheck é uma camada de triagem defensiva para investigar um sintoma específico: janelas de cmd.exe ou PowerShell que aparecem inesperadamente durante a inicialização do Windows.

A proposta não é prometer “encontrar vírus”. A proposta é mais precisa e mais defensável: coletar sinais observáveis, preservar as limitações, correlacionar a origem provável e produzir um veredito explicável para o usuário final — com profundidade técnica disponível quando necessária.


Evidence before action. O BootCheck observa primeiro, modifica nada e explica o que os dados realmente sustentam.

O problema

Uma janela de terminal no boot pode estar relacionada a uma atualização legítima, um agente de suporte, uma tarefa agendada, um serviço, uma configuração de inicialização ou uma atividade potencialmente abusiva. O sintoma isolado não é suficiente para concluir que existe malware.

O BootCheck transforma a pergunta vaga do usuário em uma investigação estruturada:


Qual processo abriu o interpretador, de onde ele veio, qual mecanismo o iniciou e quais evidências sustentam uma leitura segura, suspeita ou inconclusiva?

O escopo da v1 permanece intencionalmente limitado. O projeto não pretende substituir antivírus, EDR, sandbox, perícia digital ou resposta profissional a incidentes.

O que o cliente consegue validar agora

A entrega atual foi organizada para que um cliente consiga avaliar a qualidade do produto sem precisar executar uma coleta em uma máquina comprometida. O repositório valida contratos, fixtures, prompt, renderização e compilação do scaffold; a coleta Windows real está explicitamente marcada como próxima etapa, não escondida atrás de uma promessa de funcionamento.

Status real do repositório

Camada
Estado atual
O que isso significa
Especificação do produto
READY
MVP, persona, critérios de sucesso e limites estão documentados.
Arquitetura e modelo de evidências
READY
Fluxo, schema JSON, privacidade, veredito e responsabilidades estão definidos.
Fixtures sintéticas
VALIDATED
Casos provavelmente seguro, suspeito e inconclusivo passam pelo schema.
Prompt de veredito
SCAFFOLDED
Template 5W2H/MITRE e saída leiga/técnica estão embutidos.
Gerador de relatório
TESTED
HTML com seção leiga, apêndice técnico, escaping e aviso de segurança.
Coletor Windows
NEXT
Stubs seguros e detalhados; coleta real ainda precisa ser implementada no agente local.
API do provedor
NEXT
Cliente, timeout, retry e contrato existem; envelope do provedor escolhido ainda precisa ser conectado.
Assinatura de código
RELEASE GATE
Requer pipeline Windows, certificado e verificação Authenticode.




<details>
<summary><strong>O que significa “scaffold” neste projeto?</strong></summary>

Scaffold significa que as fronteiras, contratos e testes de integração foram preparados, mas as chamadas específicas de coleta Windows permanecem como TODOs deliberados. Essa escolha impede que um protótipo incompleto seja confundido com um coletor comercial pronto e permite que cada categoria seja implementada, testada e revisada de forma independente.

O scaffold compila, os contratos JSON são verificáveis e os testes são executáveis localmente. Ele não coleta dados do sistema atual, não inicia processos e não faz chamadas a um provedor LLM por padrão.

</details>

Validar em 5 minutos

Esta é a sequência recomendada para um cliente que quer verificar se o repositório está organizado, se os contratos são válidos e se o pipeline básico funciona.

1. Obter o projeto

Bash


git clone <URL_DO_REPOSITORIO>
cd bootcheck



2. Validar o scaffold Go

Bash


cd scaffold
gofmt -w collector/*.go api/*.go report/*.go tests/*.go
go test ./...



Resultado esperado: os pacotes api, report e tests terminam com ok. O pacote collector não possui testes de coleta real porque sua implementação Windows ainda está no roadmap.

3. Validar schemas e fixtures

Bash


cd ..
python3 -m pip install jsonschema
python3 validate_schemas.py



Resultado esperado:

Plain Text


evidence schema OK: ambiguous.json
evidence schema OK: benign_obvious.json
evidence schema OK: malicious_obvious.json
verdict schema OK



4. Confirmar compilação para Windows

Bash


cd scaffold
GOOS=windows GOARCH=amd64 go build ./...



Esse comando confirma a compilação cruzada do scaffold. Ele não executa a coleta no Linux e não substitui a validação dinâmica em uma máquina Windows descartável.

5. Avaliar a experiência do produto

Abra os seguintes arquivos nesta ordem:

Ordem
Arquivo
Pergunta de validação
01
docs/ESPECIFICACAO_PRODUTO.md
O sintoma resolvido e o cliente-alvo estão claros?
02
docs/ARQUITETURA_TECNICA.md
O fluxo coleta → veredito → relatório é auditável?
03
scaffold/tests/README.md
Os três cenários cobrem seguro, suspeito e ambíguo?
04
scaffold/api/prompt_template.go
O veredito é produzido com linguagem leiga e evidências rastreáveis?
05
docs/HANDOFF_AGENTE_LOCAL.md
A próxima fase de implementação tem tarefas e critérios de aceite?




<details>
<summary><strong>Checklist visual do cliente</strong></summary>

Plain Text


[ PASS ] O problema está restrito a Windows boot + cmd.exe/PowerShell
[ PASS ] O produto não promete ser antivírus ou EDR
[ PASS ] O coletor é somente-leitura por contrato
[ PASS ] A transmissão depende de consentimento explícito
[ PASS ] O pacote de evidências tem schema versionado
[ PASS ] Existem casos sintéticos para validação sem endpoint infectado
[ PASS ] O relatório separa linguagem leiga de apêndice técnico
[ NEXT ] Implementar e verificar coleta Windows real
[ NEXT ] Conectar e homologar o provedor LLM escolhido
[ GATE ] Assinar o binário e concluir revisão jurídica



</details>

Como funciona

mermaid

Fonte



O processo tem quatro fronteiras operacionais: o coletor lê o endpoint; o empacotador normaliza e redige; o analisador interpreta somente o JSON aprovado; e o relatório apresenta o resultado. Nenhuma dessas fronteiras possui uma ação automática de limpeza ou remediação.

O pipeline em uma tela

Plain Text


┌───────────────┐     ┌─────────────────┐     ┌────────────────┐
│ WINDOWS HOST  │────▶│ LOCAL OBSERVER  │────▶│ EVIDENCE JSON  │
│ startup       │     │ read-only       │     │ schema + hash  │
└───────────────┘     └─────────────────┘     └───────┬────────┘
                                                       │
                                                       ▼
                                             ┌──────────────────┐
                                             │ PREVIEW + REDACT │
                                             │ user-controlled  │
                                             └────────┬─────────┘
                                                      │ consent
                                                      ▼
                                             ┌──────────────────┐
                                             │ ANALYSIS API     │
                                             │ structured JSON  │
                                             └────────┬─────────┘
                                                      ▼
                                             ┌──────────────────┐
                                             │ VERDICT + REPORT │
                                             │ plain + technical│
                                             └──────────────────┘



Arquitetura

mermaid

Fonte



Para a implementação final, o coletor considera APIs Windows de leitura. A documentação da Microsoft recomenda a API Task Scheduler 2.0 para novo desenvolvimento . Para endpoints TCP, a referência arquitetural considera GetExtendedTcpTable, que enumera endpoints disponíveis ao aplicativo . Essas referências orientam a leitura do sistema; não autorizam iniciar tarefas, abrir conexões, alterar serviços ou modificar o host.

Modelo de segurança

Plain Text


MODE        READ_ONLY
TARGET      WINDOWS
TELEMETRY   LOCAL_FIRST
EXECUTION   DISABLED
REMEDIATION DISABLED
NETWORK     CONSENT_REQUIRED



<details>
<summary><strong>O que o BootCheck pode ler?</strong></summary>

Superfície
Evidência esperada
Processos
PID, PPID, nome, caminho, linha de comando, integridade, assinatura e hash opcional.
Inicialização
Chaves Run e RunOnce por escopo.
Agendamento
Nome, caminho, gatilhos, ações declaradas e contexto da tarefa.
Serviços
Estado, tipo de início, conta, caminho e assinatura.
WMI
Filtros, consumers e bindings permanentes acessíveis.
Winlogon
Valores de uma allow-list definida para o sintoma.
Rede
Endpoints TCP/UDP ativos e PID/imagem quando disponível.
Defender
Eventos recentes relevantes, sem limpar ou alterar o log.




</details> <details>
<summary><strong>O que o BootCheck nunca deve fazer?</strong></summary>

Plain Text


✗ executar cmd.exe, PowerShell, scripts ou payloads
✗ criar, encerrar, suspender, injetar ou alterar processos
✗ criar, alterar ou excluir chaves e valores do registro
✗ iniciar, parar, configurar ou remover serviços
✗ iniciar, registrar, atualizar ou excluir tarefas
✗ criar, alterar, associar ou remover WMI subscriptions
✗ limpar eventos ou desativar o Microsoft Defender
✗ fazer port scan, DNS ativo ou tentativa de conexão
✗ enviar dados antes da tela de consentimento
✗ remover, bloquear, quarentenar ou corrigir artefatos automaticamente



</details>

Veredito e experiência final

A resposta do BootCheck não é uma sentença binária. O modelo precisa explicar o que foi observado, o que é inferência, o que contradiz a hipótese e qual informação ficou indisponível.

Estado
Mensagem para o cliente
likely_safe
“Os sinais encontrados são compatíveis com atividade legítima, mas isso não é uma garantia.”
suspicious
“Há sinais que merecem investigação; isso não confirma uma infecção.”
inconclusive
“Os dados não bastam para decidir com segurança.”




O veredito usa uma estrutura 5W2H — o quê, quem, quando, onde, por quê, como e impacto observado — e um mapeamento MITRE ATT&CK somente quando sustentado por evidências. Cada afirmação decisiva deve apontar para um evidence_id existente.

Plain Text


┌─ TRIAGE STATUS ──────────────────────────────────────────────┐
│                                                               │
│  Boot anomaly detected                                       │
│                                                               │
│  Process        powershell.exe                               │
│  Parent         explorer.exe                                 │
│  Trigger        Startup                                       │
│  Evidence       14 artifacts                                  │
│  Verdict        INCONCLUSIVE                                  │
│                                                               │
│  Next step      Review evidence — do not modify automatically │
└───────────────────────────────────────────────────────────────┘



O bloco acima é uma representação visual do relatório, não uma captura de endpoint real. As fixtures do repositório são sintéticas e servem para validar o contrato, não para simular uma promessa de detecção em produção.

Modelo de evidências

O contrato canônico está em scaffold/collector/evidence.schema.json. O pacote distingue explicitamente entre:

Plain Text


OBSERVED       valor lido do sistema
INFERRED       hipótese explicada pelo analisador
UNAVAILABLE    categoria sem permissão ou suporte
WITHHELD       campo retido por redaction/consentimento
UNKNOWN        dado que não pôde ser confirmado



A estrutura foi desenhada para preservar contexto e não transformar ausência de acesso em “seguro”. O envio padrão deve usar o modo redacted, com prévia dos campos e confirmação do usuário.

Segurança e privacidade

O BootCheck é local-first, mas a análise remota pode tratar caminhos, nomes de usuário, hostname, argumentos, endereços IP ou outros dados técnicos que possam ser pessoais. A tela de consentimento deve mostrar finalidade, provedor, país, categorias, modo de dados, retenção e metadados inevitáveis do transporte.

A implementação deve ser avaliada à luz da LGPD  e das orientações da ANPD para agentes de tratamento e segurança de pequeno porte . O documento docs/RESPONSABILIDADE_LEGAL.md é um rascunho de trabalho e deve ser revisado por advogado brasileiro antes de publicação, distribuição ou venda.


Aviso: o BootCheck pode gerar falsos positivos, falsos negativos ou resultado inconclusivo. Não execute, encerre, remova, bloqueie, desative ou altere nada com base apenas no relatório. Em caso de suspeita relevante, procure um profissional autorizado.

O que está pronto e o que não está

<details>
<summary><strong>Pronto para revisão do cliente</strong></summary>

A especificação do MVP, a arquitetura, o modelo de evidências, o prompt de veredito, os contratos JSON, a camada de relatório, a matriz de cenários e o plano de handoff estão disponíveis e referenciados no repositório. O cliente consegue validar a coerência do produto, a disciplina de segurança e o pipeline de contratos com os comandos do Quick start.

</details> <details>
<summary><strong>Ainda depende do agente local</strong></summary>

A implementação das APIs Windows de coleta real, a verificação dinâmica de que nenhum estado foi alterado, a adaptação do envelope do provedor LLM, a política de retenção em produção, a assinatura Authenticode e a revisão jurídica ainda são gates de implementação/release. Cada etapa possui tarefa e critério de aceite em docs/HANDOFF_AGENTE_LOCAL.md.

</details>

Estrutura do repositório

Plain Text


bootcheck/
├── docs/                         Product, architecture and governance
│   ├── ESPECIFICACAO_PRODUTO.md  MVP, persona, metrics and scope
│   ├── ARQUITETURA_TECNICA.md    Flow, schema, API and security boundary
│   ├── RESPONSABILIDADE_LEGAL.md Privacy, consent and legal draft
│   ├── MODELO_DE_NEGOCIO.md      Three monetization options
│   ├── HANDOFF_AGENTE_LOCAL.md   Implementation prompt and acceptance gates
│   └── DECISOES_APROVADAS.md     Pending decisions register
│
├── scaffold/                     Reference implementation scaffold
│   ├── collector/                Read-only collection contracts and stubs
│   ├── api/                      Prompt, client and verdict schema
│   ├── report/                   HTML report renderer and tests
│   └── tests/                    Synthetic fixtures and validation tests
│
├── validate_schemas.py            Evidence/verdict schema validator
├── VALIDACAO.md                   Validation record
└── SUMARIO_EXECUTIVO.md           One-page executive summary



Roadmap

Fase
Entrega
Estado
01
Especificação, arquitetura, schemas e experiência do repositório
Concluído
02
Validação dos três cenários com mock provider
Próximo passo
03
Processos e Run/RunOnce somente-leitura
Implementação local
04
Task Scheduler, serviços, WMI, Winlogon, rede e Defender
Implementação local
05
Redaction, consentimento, integridade e API real
Implementação local
06
Harness Windows antes/depois, build e assinatura
Gate de release
07
Histórico, console multiusuário ou resposta a incidentes
Fora da v1




Expansões para Linux/macOS, phishing, remediação, monitoramento contínuo ou resposta a incidentes exigem nova especificação e nova análise de risco. Elas não devem ser adicionadas silenciosamente ao MVP.

Desenvolvimento e handoff

O próximo agente deve começar pelo handoff de implementação, não por uma coleta improvisada. A ordem recomendada é validar schema e fixtures, implementar uma categoria por vez, conectar a orquestração, testar redaction/consentimento, adaptar o provedor, verificar o relatório e somente então executar os gates Windows de release.

Se uma funcionalidade entrar em conflito com o princípio de somente-leitura, o princípio de somente-leitura vence. A tarefa deve parar e pedir aprovação humana explícita.

Divulgação responsável

Não publique neste repositório dados reais de endpoint, caminhos, hashes, linhas de comando, relatórios de clientes, tokens ou credenciais. Para relatar uma vulnerabilidade, use um canal privado configurado pelo mantenedor; não envie detalhes sensíveis em issue pública.

Canal de contato: [inserir endereço de disclosure antes da publicação]

Licença

A licença do repositório ainda precisa ser escolhida. Antes da distribuição pública, defina a licença, os avisos de terceiros, a política de segurança e os termos comerciais aplicáveis.

Referências

[1] Microsoft Learn — About the Task Scheduler
[2] Microsoft Learn — GetExtendedTcpTable function
[3] Planalto — Lei nº 13.709/2018 (LGPD )
[4] ANPD — Guia orientativo sobre segurança da informação para agentes de tratamento de pequeno porte
