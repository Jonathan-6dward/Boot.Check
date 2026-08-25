# BootCheck — Especificação do produto

> **Status:** rascunho de produto para validação humana. O BootCheck é uma ferramenta defensiva de triagem; não substitui investigação ou resposta a incidente conduzida por profissional qualificado. Este documento não constitui aconselhamento jurídico.

## 1. Definição do MVP

**O BootCheck coleta, de forma local e somente-leitura, evidências relacionadas a processos `cmd.exe` e PowerShell que surgem durante a inicialização do Windows e entrega uma avaliação explicada em linguagem simples, acompanhada de um apêndice técnico auditável, sobre a hipótese de atividade legítima, suspeita ou inconclusiva.**

O nome **BootCheck** é um placeholder. Alternativas a considerar, sem compromisso de disponibilidade de marca ou domínio, são **BootLens**, **Startup Sentinel** e **InitTrace**. A decisão de nome permanece pendente de aprovação e de pesquisa de marca.

O produto resolve um único sintoma na v1: “Por que janelas de `cmd.exe` ou PowerShell estão aparecendo sozinhas durante a inicialização do Windows, e isso é malicioso ou legítimo?”. A ferramenta deve evitar ampliar o diagnóstico para uma promessa geral de antivírus, EDR ou resposta a incidentes.

## 2. Usuário-alvo e contexto de uso

A persona primária é o **usuário doméstico avançado** ou o responsável por uma **pequena empresa sem SOC**, que percebe uma janela de terminal durante o boot, não possui ferramentas forenses especializadas e quer uma explicação compreensível antes de procurar ajuda profissional. Esse usuário consegue baixar um executável, aceitar uma tela de consentimento, salvar um relatório e compartilhar o arquivo com um técnico, mas não deve ser obrigado a interpretar chaves de registro, GUIDs, hashes ou técnicas MITRE ATT&CK.

A persona secundária é o prestador de suporte de TI que atende pequenas empresas e precisa de um pacote padronizado, somente-leitura e reproduzível para uma triagem inicial. O produto não deve induzir esse profissional a tratar a saída do LLM como prova conclusiva; o relatório deve separar observação, inferência e recomendação de próximos passos.

### Jornada principal

| Etapa | Experiência esperada | Resultado verificável |
| --- | --- | --- |
| 1. Início | O usuário abre o executável assinado e vê o sintoma coberto, os limites e a tela de consentimento. | Nenhuma coleta ocorre antes da confirmação explícita. |
| 2. Coleta | O coletor enumera evidências locais sem executar, encerrar ou modificar processos, registro, serviços, tarefas, arquivos ou WMI. | Pacote JSON versionado; erros parciais identificados por categoria. |
| 3. Revisão do envio | A interface mostra o que será enviado à API, campos redigidos e finalidade do tratamento. | O usuário pode cancelar ou escolher relatório local sem LLM. |
| 4. Veredito | A API retorna uma estrutura validada, com classificação, confiança, evidências citadas e linguagem leiga. | Nenhum veredito sem evidência correspondente; incerteza explícita. |
| 5. Relatório | O usuário recebe resumo “seguro”, “suspeito” ou “inconclusivo”, razões e próximos passos não destrutivos. | HTML salvo localmente e opcionalmente compartilhável. |

## 3. Escopo funcional da v1

A v1 deve capturar somente as categorias necessárias para investigar a origem de `cmd.exe`/PowerShell no boot: metadados do host e da coleta; processos em execução e árvore pai/filho; mecanismos de persistência `Run`/`RunOnce`; tarefas agendadas; serviços; assinaturas WMI permanentes; chaves de inicialização associadas a Winlogon; conexões de rede TCP ativas associadas a PID quando disponível; e eventos recentes do Microsoft Defender que estejam acessíveis ao usuário.

Cada observação deve preservar a distinção entre “não encontrado”, “não acessível” e “não coletado”. Caminhos, argumentos e identificadores que possam conter dados pessoais devem ser marcados no pacote para redaction/revisão antes do envio. O relatório deve exibir os dados técnicos completos somente na seção técnica e somente quando o usuário tiver optado por isso.

## 4. Critérios de sucesso mensuráveis

Os números abaixo são metas iniciais para validação do MVP e não alegações de desempenho já comprovado. Devem ser medidos em uma matriz de máquinas Windows representativa antes da publicação.

| Dimensão | Meta de aceitação da v1 | Como medir |
| --- | --- | --- |
| Tempo de coleta | P95 inferior a **120 segundos** em uma máquina de referência sem espera interativa. | 30 execuções em três perfis de hardware, contando do clique em “Iniciar” até o JSON fechado. |
| Confiabilidade | Pelo menos **95%** das execuções geram um pacote estruturalmente válido, mesmo quando uma categoria exige privilégio ou não está disponível. | Teste automatizado + matriz de permissões; falhas parciais não invalidam categorias independentes. |
| Integridade somente-leitura | **100%** dos testes de segurança não observam mutações em processos, serviços, tarefas, registro, arquivos de evidência ou WMI. | Snapshot antes/depois, monitor de eventos e revisão manual do binário. |
| Completude do sintoma | **100%** dos casos sintéticos com uma origem explícita em persistência incluem essa origem no conjunto de evidências. | Fixtures benignas, maliciosas e ambíguas em `/scaffold/tests/`. |
| Qualidade do veredito | Em um conjunto rotulado por dois revisores, pelo menos **90%** dos casos óbvios recebem a classe correta; divergências devem ser registradas. | Avaliação cega de fixtures; a meta não autoriza uso como decisão automática. |
| Falso positivo | Alvo inicial de no máximo **10%** nos casos benignos rotulados; o relatório deve preferir “inconclusivo” a acusar sem evidência. | Conjunto benigno ampliado com softwares legítimos de atualização, suporte e administração. |
| Rastreabilidade | **100%** das afirmações decisivas do relatório apontam para um `evidence_id` ou para a ausência documentada de evidência. | Validador que percorre referências no JSON de saída. |
| Tamanho do artefato | Relatório HTML leigo de até **500 KB** sem anexos; JSON técnico comprimido de até **10 MB** na maioria dos hosts. | Medição em 30 máquinas; excedentes devem gerar aviso, não descarte silencioso. |
| Privacidade | **100%** dos envios à API precedidos de consentimento registrado e de prévia dos campos. | Testes de interface, logs de auditoria locais e inspeção de tráfego em ambiente de teste. |

A meta de acurácia não transforma o produto em antivírus nem elimina falsos positivos/negativos. A classificação deve ser tratada como triagem assistida por evidências, e não como prova de comprometimento ou de inocência.

## 5. Experiência e mensagens mínimas

Antes da coleta, a interface deve apresentar: o sintoma coberto; o que será lido; o que não será modificado; se o resultado será enviado a um provedor de IA; quais campos poderão sair da máquina; período de retenção; como cancelar; e o aviso de que o resultado é indicativo. O botão de início deve estar desmarcado por padrão e exigir ação afirmativa.

Para pessoas leigas, o relatório deve usar uma das três classes controladas:

| Classe | Mensagem-base | Regra de uso |
| --- | --- | --- |
| Provavelmente seguro | “Os sinais encontrados são compatíveis com atividade legítima, mas isso não é garantia.” | Só quando houver evidência positiva de legitimidade e nenhum sinal forte de abuso. |
| Suspeito | “Há sinais que merecem investigação; isso não confirma uma infecção.” | Exigir pelo menos uma evidência forte ou combinação de sinais moderados. |
| Inconclusivo | “Os dados não bastam para decidir com segurança.” | Usar quando houver conflito, lacunas de acesso, pouca evidência ou baixa confiança. |

O relatório nunca deve recomendar executar um arquivo, desativar o Defender, remover chave, matar processo, apagar tarefa ou alterar configuração automaticamente. Próximos passos podem incluir salvar o relatório, consultar o fornecedor do software, desconectar a máquina conforme procedimento interno da organização e buscar um profissional autorizado; qualquer ação operacional fica fora da v1.

## 6. Fora de escopo da v1 e roadmap

Os itens a seguir são deliberadamente excluídos e devem aparecer como roadmap, não como implementação oculta:

| Fora de escopo na v1 | Possível fase futura, sujeita a nova análise |
| --- | --- |
| Análise de e-mail ou phishing | Módulo independente, com modelo de ameaça e governança próprios. |
| Resposta a incidente completa | Integração com playbooks conduzidos por profissional autorizado. |
| Remoção, quarentena ou remediação de artefatos | Nunca automática; apenas eventual orientação manual aprovada. |
| Suporte a Linux ou macOS | Novo coletor e nova matriz de evidências. |
| Qualquer execução, modificação ou remediação no sistema do usuário | Não é roadmap da v1; exigir decisão de segurança separada. |
| Monitoramento contínuo, agente residente ou varredura periódica | Produto diferente, com novos riscos de privacidade e operação. |
| Veredito local sem API | Pode ser explorado para modo offline, mas não faz parte do primeiro corte. |
| Histórico em nuvem, console multiusuário ou SOC | Fase comercial posterior, dependente de modelo de negócio e governança de dados. |

## 7. Riscos de produto e decisões abertas

Os riscos principais são a interpretação excessiva de um veredito probabilístico, a exposição de caminhos/nomes de usuário no envio, o comportamento de antivírus diante de um coletor novo, indisponibilidade de APIs Windows em edições diferentes e variação de custo do LLM. As decisões que precisam de aprovação antes da implementação final são: nome e marca; linguagem do coletor; provedor/modelo de LLM; política de retenção; definição do conjunto de eventos do Defender; e modelo de negócio.

## 8. Aviso de responsabilidade para a interface

> O BootCheck é uma ferramenta de triagem defensiva e somente-leitura. O relatório é uma avaliação automatizada baseada nas evidências disponíveis naquele momento; pode conter falsos positivos, falsos negativos ou ficar inconclusivo. Não execute, remova ou altere nada com base apenas neste relatório. Em caso de suspeita relevante, procure um profissional autorizado. **Isto não é aconselhamento jurídico; o produto e seus textos devem ser revisados por advogado antes da publicação.**

## Referências

[1]: https://www.planalto.gov.br/ccivil_03/_ato2015-2018/2018/lei/l13709.htm "Lei nº 13.709/2018 — Lei Geral de Proteção de Dados Pessoais"
[2]: https://www.gov.br/anpd/pt-br/centrais-de-conteudo/materiais-educativos-e-publicacoes/guia-orientativo-para-definicoes-dos-agentes-de-tratamento-de-dados-pessoais-e-do-encarregado "ANPD — Guia orientativo para definições dos agentes de tratamento de dados pessoais e do encarregado"
