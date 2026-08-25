# BootCheck — Responsabilidade legal, uso e privacidade

> **ISTO NÃO É ACONSELHAMENTO JURÍDICO.** Sou uma IA, não advogado. Este documento é um rascunho de produto e privacidade, baseado nas fontes listadas ao final, e deve ser revisado por advogado brasileiro com experiência em proteção de dados, software e responsabilidade civil **antes de qualquer publicação, distribuição ou venda**. A revisão deve confirmar a base legal, os papéis de controlador/operador, retenção, transferências internacionais, contratos com provedores e texto final do termo.

## 1. Finalidade e limites jurídicos do documento

Este arquivo organiza salvaguardas de responsabilidade para uma ferramenta defensiva de **triagem**. Ele não certifica conformidade com a LGPD, não substitui análise de risco, registro das operações, avaliação de impacto, contratos ou políticas internas e não garante que o produto seja lícito em qualquer cenário de uso. A Lei nº 13.709/2018 disciplina o tratamento de dados pessoais inclusive em meios digitais [1], e a ANPD publica guias orientativos e não-vinculantes sobre agentes de tratamento e segurança [2] [3].

A publicação deve ser bloqueada até que o responsável legal aprove: a versão do termo; aviso de privacidade; mecanismo de consentimento; canal de atendimento ao titular; contrato com provedor de API; política de retenção e eliminação; segurança do produto; e processo para incidentes.

## 2. Rascunho de termo de uso

### 2.1 Aceitação

Ao instalar, abrir ou utilizar o BootCheck, o usuário declara que leu este termo, o aviso de privacidade e a tela de consentimento, e que utilizará a ferramenta apenas em máquina própria ou para a qual possua autorização válida. Se o usuário estiver atuando por uma empresa, declara possuir autoridade para autorizar a triagem e o eventual envio de dados ao provedor indicado na interface.

### 2.2 Objeto do serviço

O BootCheck coleta evidências técnicas locais para investigar o sintoma de janelas de `cmd.exe` ou PowerShell aparecendo durante a inicialização do Windows. O serviço é limitado à observação do estado disponível no momento da coleta e à elaboração de uma avaliação automatizada explicada em linguagem simples.

### 2.3 Somente-leitura

O coletor foi projetado para ler dados. Ele não deve executar, encerrar ou modificar processos; alterar registro, serviços, tarefas agendadas, arquivos do usuário ou assinaturas WMI; limpar eventos; alterar políticas; desativar proteções; remover ou quarentenar artefatos; nem fazer remediação automática. O usuário não deve interpretar o relatório como autorização para praticar qualquer uma dessas ações.

### 2.4 Uso autorizado e responsabilidade do usuário

O usuário é responsável por obter autorização para examinar o equipamento, respeitar políticas de sua organização e manter cópia do relatório em local adequado. O usuário deve evitar compartilhar um pacote que contenha caminhos, nomes de conta, hostname, endereços IP, argumentos de processo, hashes ou outros dados que possam identificar pessoas, dispositivos ou sistemas, salvo quando isso for necessário e autorizado.

### 2.5 Natureza indicativa do resultado

O resultado é uma hipótese de triagem produzida a partir de evidências limitadas. Ele pode estar incorreto, desatualizado, incompleto, sujeito a falso positivo ou falso negativo e não constitui certificação de segurança, diagnóstico definitivo, prova de comprometimento, prova de autoria ou garantia de inocência. O BootCheck não substitui antivírus, EDR, perícia, investigação forense, resposta a incidentes, suporte do fabricante ou aconselhamento profissional.

### 2.6 Proibições

É proibido usar o BootCheck para obter acesso não autorizado, investigar equipamento de terceiro sem autorização, contornar controles de segurança, ocultar atividade, desenvolver malware, praticar vigilância indevida ou realizar ação ofensiva. O produto não fornece e não autoriza técnicas de evasão de antivírus/EDR, exploração, persistência, exfiltração ou comprometimento.

### 2.7 Disponibilidade e alterações

O produto pode apresentar falhas de permissão, incompatibilidades de edição/versão do Windows, indisponibilidade do provedor de API, erros de rede ou resultado inconclusivo. A organização pode atualizar o coletor, schema, prompt, provedor, preço, limites ou termo; mudanças relevantes devem ser comunicadas conforme orientação jurídica.

### 2.8 Limitação de responsabilidade — rascunho

Na extensão permitida pela legislação aplicável, o fornecedor não promete que a triagem detectará toda atividade maliciosa, nem que o relatório evitará danos. Nenhuma limitação deve ser publicada sem revisão jurídica das regras aplicáveis ao consumidor, à contratação empresarial e a situações de vulnerabilidade. Nada neste termo pretende excluir responsabilidade que a lei não permita excluir.

### 2.9 Contato

O produto deve informar um canal real para dúvidas de privacidade, solicitações de titular e incidentes: **[inserir e-mail/endereço do controlador]**. O nome empresarial, CNPJ, endereço, responsável por privacidade/encarregado quando aplicável e provedor do serviço devem ser preenchidos antes da publicação.

## 3. Disclaimer de falsos positivos e falsos negativos

> **Aviso importante sobre o resultado:** o BootCheck faz uma triagem automatizada com as informações acessíveis no momento da coleta. Um resultado “provavelmente seguro” não garante que o computador esteja livre de ameaça; um resultado “suspeito” não prova infecção; e um resultado “inconclusivo” significa que os dados não bastaram para uma conclusão responsável. Falsos positivos e falsos negativos são possíveis. Não execute, encerre, remova, bloqueie, desative ou altere nada com base apenas no relatório. Salve o relatório e procure um profissional autorizado quando a situação for relevante.

O aviso deve aparecer antes do envio, no cabeçalho do relatório e antes de qualquer ação de exportação/compartilhamento. A UI deve evitar verbos absolutos como “confirmado”, “limpo”, “seguro” ou “infectado”. As três classes permitidas são `likely_safe`, `suspicious` e `inconclusive`, com a definição operacional documentada em `/docs/ESPECIFICACAO_PRODUTO.md`.

## 4. LGPD e fluxo de dados

### 4.1 Princípios de desenho

A implementação deve ser alinhada, após validação jurídica, aos princípios de finalidade, adequação, necessidade, transparência, segurança, prevenção, responsabilização e prestação de contas previstos na LGPD [1]. A coleta local deve ter finalidade declarada e limitada ao sintoma da v1; módulos fora de escopo não devem ser coletados “por precaução”. A ANPD disponibiliza guia sobre agentes de tratamento e encarregado e guia de segurança para agentes de pequeno porte [2] [3]; ambos devem ser usados como material de planejamento, não como certificado.

### 4.2 Papéis a confirmar

Como hipótese de trabalho, a empresa que define a finalidade do BootCheck provavelmente terá papel de **controladora** do tratamento que decide realizar no serviço; o provedor de API pode atuar como operador ou assumir outro papel conforme contrato, instruções, finalidade própria e configuração efetiva; o usuário pode ser titular dos dados ou atuar em nome de uma organização. Essa qualificação não pode ser fixada apenas por este documento: o advogado deve avaliar a operação concreta, inclusive coleta para uso exclusivamente pessoal, uso empresarial, subcontratados, hospedagem, suporte e transferências internacionais.

### 4.3 O que é coletado localmente

A v1 pode ler as seguintes categorias técnicas, sempre com minimização e possibilidade de estado “não acessível”: versão/build/arquitetura do Windows; hostname e usuário interativo, se necessários; processos, PID/PPID, nome, caminho, linha de comando, horário, integridade, assinatura e hash opcional; chaves `Run`/`RunOnce`; definições de Task Scheduler; metadados de serviços; subscriptions WMI; valores autorizados de Winlogon; endpoints de rede ativos com PID/imagem quando disponível; e eventos recentes do Defender.

A coleta local, por si só, não deve enviar esses dados a terceiros. O pacote deve ser salvo localmente somente pelo tempo necessário para visualizar, gerar e, se o usuário decidir, enviar o relatório. O produto deve permitir cancelar e excluir seus temporários sem excluir dados do sistema que apenas foram lidos.

### 4.4 O que sai da máquina quando o usuário consente

Antes do primeiro envio, a tela deve listar exatamente o conteúdo efetivo do payload, com contagem e prévia. O payload mínimo proposto para o veredito é:

| Categoria enviada | Campos incluídos | Tratamento padrão |
| --- | --- | --- |
| Metadados | Versão/build, arquitetura, identificador de coleta, timestamps, versão do coletor, estados de categoria e limitações. | Enviar; hostname/usuário devem ser redigidos ou omitidos por padrão. |
| Processos | PID/PPID, nome, caminho, linha de comando, integridade, assinatura, hash e horário quando necessários à correlação. | Redigir nome de usuário, perfil e segmentos de caminho; preservar identificador estável local quando necessário. |
| Persistência | Hive/escopo, caminho da chave, nome/valor e origem. | Redigir perfil do usuário e valor que não tenha relação com o sintoma; mostrar a prévia. |
| Tarefas/serviços/WMI/Winlogon | Nome/caminho, ações declaradas, contexto, conta, estado e metadados relevantes. | Enviar apenas campos necessários; valores extensos e identificadores pessoais devem ser redigidos. |
| Rede | Protocolo, endereços/portas, estado, PID e imagem proprietária quando disponível. | Redigir IP privado/hostname se não forem necessários; não fazer resolução ativa. |
| Defender | Provedor, ID, timestamp, categoria/ação e mensagem minimizada. | Não enviar conteúdo além dos eventos necessários; evitar identificadores pessoais. |
| Integridade | Hash do JSON canônico e versão do schema/prompt. | Enviar para rastreabilidade, sem segredo. |

O aplicativo também transmite metadados inevitáveis do transporte, como endereço IP de rede, data/hora, user-agent ou identificadores técnicos do provedor, conforme a operação do endpoint. A tela de consentimento deve informar isso com o provedor real e sua política de privacidade; esta lista não substitui inventário técnico do tráfego.

É proibido enviar automaticamente: tokens ou chaves de API, credenciais, cookies, conteúdo de arquivos não relacionados, memória de processo, documentos, e-mails, histórico de navegação, senhas ou dados deliberadamente coletados para outra finalidade. Segredos nunca devem entrar no JSON de evidências.

### 4.5 Consentimento e retirada

A primeira tela de envio deve ter caixas desmarcadas por padrão e explicar finalidade, categorias, provedor, retenção, possível transferência internacional, modo `redacted`/`full`, canal do titular e consequência de não consentir. A ação deve registrar versão do texto, timestamp, identificador da coleta, modo de dados e escolha do usuário. O cancelamento deve ser fácil antes da transmissão e o produto deve informar como retirar consentimento para operações futuras. A validade jurídica dessa base e a necessidade de outra base legal devem ser confirmadas pelo advogado.

Texto sugerido para a tela:

> “A coleta abaixo ocorrerá somente nesta máquina e em modo somente-leitura. Para gerar o veredito automatizado, o BootCheck enviará ao provedor **[nome e país do provedor]** os campos exibidos na prévia: **[contagem e categorias]**. O envio tem a finalidade exclusiva de investigar o sintoma de janelas de cmd.exe/PowerShell no boot. O resultado pode conter erros. Nenhum arquivo será executado, encerrado, alterado ou removido pelo BootCheck. Ao marcar ‘Concordo e enviar’, autorizo o envio descrito para esta análise, ciente da política de retenção **[prazo/política]**. Posso cancelar e gerar apenas o relatório local.”

## 5. Segurança e governança mínima

A organização deve manter inventário do tratamento, controle de acesso, minimização, criptografia em trânsito, gestão de segredos fora do repositório, logs sem dados sensíveis, retenção limitada, revisão de dependências, resposta a incidentes, procedimento para exclusão e mecanismo de auditoria do consentimento. O guia da ANPD para agentes de pequeno porte e seu checklist podem apoiar a priorização de medidas [3].

O cliente deve impedir logar o payload completo, linha de comando completa, hash de token, resposta bruta com dados pessoais ou chave de API. O servidor deve validar tamanho, schema e autenticação, limitar retry, remover dados temporários conforme política e não usar o pacote para treinamento ou finalidade secundária sem base e transparência adequadas.

## 6. Aprovações bloqueantes antes do lançamento

| Decisão | Responsável sugerido | Evidência necessária |
| --- | --- | --- |
| Identidade do controlador e papéis do provedor | Jurídico/privacidade | Parecer ou registro interno aprovado. |
| Base legal e texto de consentimento | Jurídico/privacidade e produto | Tela versionada e teste de entendimento. |
| Países, suboperadores e retenção do LLM | Compras/jurídico/segurança | Contrato, DPA e aviso de privacidade. |
| Segurança do binário e cadeia de assinatura | Engenharia/segurança | Build reproduzível, assinatura e verificação. |
| Falsos positivos/negativos e suporte | Produto/engenharia/suporte | Avaliação com fixtures e processo de escalonamento. |
| Termo de uso e limitações | Advogado | Aprovação escrita antes de publicação. |

## Referências

[1]: https://www.planalto.gov.br/ccivil_03/_ato2015-2018/2018/lei/l13709.htm "Planalto — Lei nº 13.709/2018 (LGPD)"
[2]: https://www.gov.br/anpd/pt-br/centrais-de-conteudo/materiais-educativos-e-publicacoes/guia-orientativo-para-definicoes-dos-agentes-de-tratamento-de-dados-pessoais-e-do-encarregado "ANPD — Guia orientativo para definições dos agentes de tratamento de dados pessoais e do encarregado"
[3]: https://www.gov.br/anpd/pt-br/centrais-de-conteudo/materiais-educativos-e-publicacoes/guia-orientativo-sobre-seguranca-da-informacao-para-agentes-de-tratamento-de-pequeno-porte "ANPD — Guia orientativo sobre segurança da informação para agentes de tratamento de pequeno porte"
