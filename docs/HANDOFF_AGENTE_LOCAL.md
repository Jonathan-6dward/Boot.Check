# Handoff do BootCheck para agente local

> **Instrução de segurança:** este handoff é para finalizar uma ferramenta defensiva de triagem, usada pelo dono da máquina ou por profissional autorizado. Todo código deve permanecer estritamente **somente-leitura**. Nenhuma tarefa abaixo pode introduzir execução de payload, encerramento de processo, alteração de registro, serviço, tarefa, WMI, arquivo do usuário, política do Defender, firewall ou rede. Não gerar técnicas ofensivas nem evasão de antivírus/EDR. Se uma tarefa parecer exigir escrita, execução, elevação, persistência ou remediação, pare e peça aprovação humana explícita.

## Prompt para colar em um agente de codificação local

Você é o agente de implementação local do projeto **BootCheck**, uma ferramenta comercial de triagem forense defensiva para Windows. Trabalhe no repositório que contém esta pasta. Não tente executar coleta real dentro de ambiente não-Windows. Faça alterações pequenas, revisáveis e acompanhadas de testes. Não invente dados, não ignore limitações e não transforme a saída do modelo em diagnóstico definitivo.

### 1. Sintoma e contrato do produto

A v1 responde somente: **“Por que janelas de `cmd.exe` ou PowerShell aparecem sozinhas durante a inicialização do Windows, e isso é malicioso ou legítimo?”**. A coleta é local, somente-leitura, e percorre processos/árvore pai-filho, `Run`/`RunOnce`, Task Scheduler, serviços, subscriptions WMI, chaves de Winlogon, conexões de rede ativas e eventos recentes do Defender. O fluxo é:

```text
coleta local → pacote JSON versionado → prévia/redaction → consentimento explícito → API LLM → veredito estruturado → HTML leigo + apêndice técnico
```

O produto não executa comandos, não corrige a máquina e não oferece resposta a incidentes. Análise de e-mail/phishing, Linux/macOS, remoção/quarentena, agente residente e monitoramento contínuo ficam fora da v1.

### 2. O que já está pronto

| Arquivo | Conteúdo pronto |
| --- | --- |
| `docs/ESPECIFICACAO_PRODUTO.md` | MVP, persona, jornada, métricas, mensagens, fora de escopo, riscos e decisões abertas. |
| `docs/ARQUITETURA_TECNICA.md` | Fluxo, decisão proposta por Go, comparação com PowerShell/Rust, limites de segurança, schema de referência e prompt 5W2H/MITRE com saída leiga/técnica. |
| `docs/RESPONSABILIDADE_LEGAL.md` | Rascunho de termo de uso, disclaimer de falsos positivos/negativos, inventário de dados, tela de consentimento e questões LGPD. **Revisão de advogado é obrigatória antes de publicar.** |
| `docs/MODELO_DE_NEGOCIO.md` | Três opções sem decisão: compra única, assinatura e freemium, incluindo cálculo unitário de API e premissas. |
| `scaffold/collector/types.go` | Tipos Go para `EvidencePackage` e categorias. |
| `scaffold/collector/collector.go` | Ponto de entrada e contrato de somente-leitura; retorna TODO deliberado. |
| `scaffold/collector/collect_*.go` | Stubs detalhados para processos, persistência, tarefas, serviços, WMI, Winlogon, rede e Defender. |
| `scaffold/collector/evidence.schema.json` | JSON Schema 2020-12 do pacote de evidências. É a fonte de verdade para validação. |
| `scaffold/collector/privacy.go` | Contrato-placeholder de redaction e validação de consentimento. |
| `scaffold/collector/integrity.go` | Contrato-placeholder de canonicalização, SHA-256 e validação. |
| `scaffold/api/prompt_template.go` | Prompt de sistema/usuário com boundary de dados não confiáveis, 5W2H, MITRE e saída leiga. |
| `scaffold/api/client.go` | Cliente HTTP-placeholder com timeout, retry limitado, limite de resposta e validação básica. Não loga payload. |
| `scaffold/api/schema.go` e `verdict.schema.json` | Contrato estruturado de `VerdictResponse`. |
| `scaffold/report/report.go` | Gerador HTML com seção leiga, apêndice técnico, aviso e escaping de template. |
| `scaffold/tests/*.json` | Fixtures sintéticas benign, maliciosa e ambígua; não são dados de endpoint real. |
| `scaffold/tests/*.go` | Testes de forma das fixtures. Há testes adicionais no pacote `api` e `report`. |

### 3. Decisões humanas ainda pendentes

Antes de fechar o release, peça aprovação para: nome final (BootCheck ou alternativa); linguagem do coletor (Go é a recomendação atual, não decisão irrevogável); provedor/modelo da API; política de retenção; texto de consentimento; papel jurídico de cada participante; e modelo de negócio entre compra única, assinatura e freemium. Não escolha pelo proprietário.

## 4. Tarefas restantes, em sessões pequenas

### Tarefa 1 — Fixar decisões e baseline de build

Leia os cinco documentos em `docs/`, confirme decisões com o proprietário e registre uma decisão curta em `docs/DECISOES_APROVADAS.md`. Não altere o escopo do sintoma sem aprovação.

**Aceite:** nome, modelo/provedor, modelo de negócio, retenção e modo padrão (`redacted`) estão registrados; `go test ./...` é executado a partir de `scaffold/`; nenhuma chave ou dado real entra no repositório.

**Dependências:** nenhuma. Esta tarefa bloqueia as demais que dependem do provedor ou da política de retenção.

### Tarefa 2 — Validar o schema do pacote

Implemente a validação de `scaffold/collector/evidence.schema.json` com biblioteca aprovada e testes positivos/negativos. Corrija qualquer divergência entre `types.go`, o schema e os fixtures.

**Aceite:** os três fixtures sintéticos têm JSON válido; um fixture com campo obrigatório faltando, hash inválido ou `read_only_assertion=false` falha; falhas de categoria são representadas sem apagar `limitations`.

**Dependências:** Tarefa 1. Bloqueia a orquestração do coletor e o cliente de API.

### Tarefa 3 — Implementar redaction e prévia local

Substitua o TODO de `privacy.go`. Redija hostname, usuário, perfil, IP privado e argumentos conforme política; preserve correlação determinística apenas se necessária. Mantenha o original apenas em memória/local protegido e nunca em log.

**Aceite:** modo padrão é `redacted`; saída redigida não contém valores proibidos; aplicar a política duas vezes é idempotente; consentimento `full` sem confirmação explícita é rejeitado; nenhuma chamada de rede é feita.

**Dependências:** Tarefa 2. Bloqueia a Tarefa 12.

### Tarefa 4 — Implementar processos e árvore pai/filho

Implemente `collect_processes.go` usando APIs Windows de consulta. Capture PID/PPID, imagem, caminho, linha de comando, horário, integridade, usuário, assinatura e hash opcional. Não solicite direitos de escrita e não carregue/executa imagem.

**Aceite:** em máquina de teste, a enumeração é pontual, não cria filho, não abre handle com permissão de escrita, preserva campos inacessíveis como `null` e gera `parent_evidence_id` somente para PID observado. Snapshot antes/depois não mostra mutação.

**Dependências:** Tarefa 2. Pode ocorrer em paralelo apenas com tarefas que não compartilhem o mesmo wrapper Windows; integração depende da Tarefa 13.

### Tarefa 5 — Implementar `Run` e `RunOnce`

Implemente `collect_persistence.go` para HKCU/HKLM e visões 32/64-bit necessárias. Abra chaves em consulta, registre escopo e preserve valores com redaction.

**Aceite:** as chaves allow-listed são lidas; nenhuma chave é criada, alterada ou removida; falha de uma hive não esconde as demais; teste antes/depois do registro permanece idêntico; valores e caminhos têm `evidence_id`.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 6 — Implementar Task Scheduler 2.0

Implemente `collect_tasks.go` com API Task Scheduler 2.0. Enumere nome/caminho, autor, descrição, estado, gatilhos, ações declaradas, principal, contexto e nível.

**Aceite:** nenhuma tarefa é iniciada, parada, registrada, atualizada ou excluída; gatilhos e ações são dados, não executados; caminhos inacessíveis viram `limitations`; teste de snapshot e monitoramento não detecta escrita.

**Dependências:** Tarefa 2. Pode começar após baseline, mas integração depende da Tarefa 13.

### Tarefa 7 — Implementar serviços

Implemente `collect_services.go` com enumeração somente-leitura do Service Control Manager. Capture estado, tipo de início, conta, caminho, descrição e assinatura.

**Aceite:** não chama APIs de iniciar/parar/configurar/excluir; não carrega o binário do serviço; serviço inacessível não é “seguro”; categoria e limitações são válidas contra o schema.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 8 — Implementar WMI subscriptions permanentes

Implemente `collect_wmi.go` para leitura de filtros, consumers e bindings relevantes no namespace aprovado. Não crie, altere, associe, desassocie ou remova objetos.

**Aceite:** coleta é limitada e documentada; query timeout funciona; repositório inacessível aparece com impacto; nenhum objeto WMI muda antes/depois; campos de creator/query respeitam redaction.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 9 — Implementar Winlogon

Implemente `collect_winlogon.go` com allow-list definida nos testes para os valores que podem explicar inicialização. Leia registro em modo consulta.

**Aceite:** allow-list explícita; nenhum valor fora do escopo é coletado; nenhuma escrita ocorre; resultado distingue “valor não encontrado” de “acesso negado”; nenhuma conclusão de malícia é criada no coletor.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 10 — Implementar rede local

Implemente `collect_network.go` usando API de enumeração de endpoints, incluindo IPv4/IPv6 e PID quando disponível. A implementação não fará DNS ativo, scan, socket, bloqueio ou tentativa de conexão.

**Aceite:** conexões existentes são listadas sem tráfego adicional; ausência de owner PID é representada; nenhum firewall/rota/socket é alterado; a categoria valida contra o schema e registra limitações.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 11 — Implementar eventos recentes do Defender

Implemente `collect_defender.go` para ler canal e janela temporal aprovados. Minimize mensagem e preserve ID/timestamp/ação/categoria.

**Aceite:** limite de quantidade/tempo é aplicado; nenhum log é limpo; nenhuma política é alterada; nenhuma proteção é desativada; erro de canal vira limitação; conteúdo é redigido antes da prévia.

**Dependências:** Tarefa 2. Integração na Tarefa 13.

### Tarefa 12 — Consentimento e contrato de envio

Conecte `ValidateSubmissionPolicy` à UI/CLI local. Mostre provedor, país, finalidade, campos, modo, retenção, possíveis metadados de transporte e o disclaimer. Não use checkbox pré-marcado.

**Aceite:** sem consentimento local não há coleta; sem consentimento LLM não há HTTP; a prévia corresponde byte a byte ao payload enviado; cancelamento não chama a rede; registro guarda versão do texto, timestamp, modo e provedor; logs não contêm segredo/payload.

**Dependências:** Tarefa 3 e decisão jurídica. Bloqueia Tarefa 14.

### Tarefa 13 — Orquestração, integridade e pacote canônico

Substitua o TODO do ponto de entrada. Execute categorias independentes em sequência controlada ou concorrência segura, sem processos-filhos; normalize arrays; calcule hash excluindo o próprio campo de hash; valide schema; e preserve limitações.

**Aceite:** todas as categorias exigidas aparecem; falha parcial não vira pacote “seguro”; `read_only_assertion=true` só é emitido se o contrato puder ser sustentado; hash é determinístico para mesma entrada; coleta encerra no timeout; nenhum request externo sai do pacote `collector`.

**Dependências:** Tarefas 2–11. Tarefa 3 é necessária antes de submeter.

### Tarefa 14 — Adaptador do provedor LLM e erros

Adapte `client.go` ao provedor aprovado. Extraia a resposta estruturada correta do envelope do provedor, use schema estrito se suportado, valide todos os `evidence_id` contra o pacote e mantenha retry apenas para 429/5xx/erros transitórios.

**Aceite:** endpoint/modelo são configuráveis; timeout, limite de resposta e retry máximo são testados; não há log de segredo/resposta bruta; erro permanente não é repetido; resposta inválida gera `inconclusive` técnico sem fabricar veredito; conteúdo do pacote é tratado como dado não confiável.

**Dependências:** Tarefas 2, 3, 12 e 13; decisão de provedor/modelo.

### Tarefa 15 — Gerador de relatório e proteção de conteúdo

Complete `report.go` para incluir collection ID, versão de schema/prompt, hash do pacote, limitações e referências. Mantenha duas seções: “Para você” e “Apêndice técnico”.

**Aceite:** HTML escapa texto vindo do modelo; headline e resumo aparecem sem jargão; apêndice inclui caminhos/hashes apenas de acordo com consentimento; o aviso de falso positivo/negativo aparece no topo e rodapé; nenhuma recomendação destrutiva ou execução automática é criada.

**Dependências:** Tarefa 14 e `VerdictResponse` validado.

### Tarefa 16 — Testes de prompt com mock provider

Execute os três fixtures contra mock local. Adicione testes de prompt injection em campos de evidência, referências inexistentes, categorias inacessíveis e conflito de sinais.

**Aceite:** benigno tende a `likely_safe`, malicioso tende a `suspicious` e ambíguo tende a `inconclusive`; cada afirmação decisiva cita evidência existente; o agente ignora instruções inseridas nos valores; a avaliação não envia dados a provedor real.

**Dependências:** Tarefas 2, 3, 14 e 15.

### Tarefa 17 — Verificação de somente-leitura em Windows

Crie harness de teste em máquina descartável com snapshots de registro, serviços, tarefas, WMI, arquivos e estado de rede; combine revisão estática, logs de chamadas e comparação antes/depois. Não usar máquina de produção ou infectada como primeiro teste.

**Aceite:** diferenças não autorizadas fazem o pipeline falhar; nenhuma chamada a execução, alteração ou remediação é encontrada; relatório de teste é anexado ao release; exceções são tratadas como bloqueantes até revisão humana.

**Dependências:** Tarefas 4–13.

### Tarefa 18 — Build, assinatura e distribuição

Prepare build reprodutível Windows amd64/arm64, versionamento, hash, artefato e assinatura Authenticode com certificado da organização. Documente verificação. Não tente contornar antivírus/EDR.

**Aceite:** o binário assinado é verificável; pipeline não contém segredo em texto; hash publicado corresponde ao artefato; instruções de instalação não executam scripts ocultos; alertas de antivírus são tratados por revisão e contato com fornecedores, nunca por evasão.

**Dependências:** Tarefas 13 e 17; decisão de linguagem.

### Tarefa 19 — Revisão de privacidade, legal e release gate

Leve `docs/RESPONSABILIDADE_LEGAL.md` a advogado, atualize controlador/operador, aviso, retenção, suboperadores, canal do titular e procedimento de incidente. Não declarar conformidade sem aprovação.

**Aceite:** aprovação escrita ou pendências explícitas; tela de consentimento corresponde ao tráfego real; política de retenção implementada; release bloqueado se qualquer ponto jurídico essencial estiver em aberto.

**Dependências:** Tarefa 12 e contrato do provedor.

## 5. Ordem recomendada

A ordem mínima é **1 → 2 → 3 → 4–11 → 13 → 12 → 14 → 15 → 16 → 17 → 18 → 19**. As tarefas 4–11 podem ser feitas em sessões separadas, mas cada uma deve entregar testes e não deve ser integrada silenciosamente. Tarefa 12 pode começar em paralelo com coleta após o schema, porém nenhuma chamada remota deve existir antes de seu aceite. As tarefas 17–19 são gates de release e não devem ser puladas para “testar com usuário”.

## 6. Checklist de segurança em toda sessão

Antes de aceitar uma alteração, confirme que ela não chama shell, não cria processo-filho, não carrega payload, não usa handle de escrita, não modifica registro/serviço/tarefa/WMI, não limpa eventos, não altera Defender/firewall, não faz DNS/scan/conexão ativa e não transmite dados sem consentimento. Rejeite qualquer código que trate “não acessível” como “seguro”, que invente publisher/hash ou que instrua o usuário a remover algo automaticamente.

Se a implementação precisar de privilégio administrativo para ver uma categoria, registre a limitação e avalie se o usuário pode autorizar a leitura. Não inclua elevação silenciosa nem desative controles. Se houver conflito entre funcionalidade e somente-leitura, a regra de somente-leitura vence e a decisão sobe para o proprietário.

## 7. Resultado esperado do agente local

Ao concluir, entregue um pull request local ou conjunto de commits pequenos contendo código, testes, relatório do harness, binários assinados apenas no pipeline aprovado e atualização da documentação. Inclua uma matriz que associe cada afirmação do relatório a `evidence_id`, limitações conhecidas e versão do prompt/modelo. O agente local não deve declarar que o produto detecta “toda” ameaça ou que o resultado é conclusivo.

**Fim do prompt de handoff.**
