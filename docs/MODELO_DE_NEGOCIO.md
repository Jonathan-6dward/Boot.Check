# BootCheck — Modelo de negócio

> **Aviso de produto:** os números abaixo são uma simulação para decisão, não uma promessa de preço ou margem. O custo real inclui mais do que tokens: infraestrutura, distribuição, assinatura de código, suporte, impostos, chargebacks, armazenamento, observabilidade, atendimento e eventuais taxas do provedor. O produto continua sendo uma triagem defensiva, e não um serviço de resposta a incidentes.

## 1. Premissas da simulação

Para tornar as alternativas comparáveis, cada análise considera **20.000 tokens de entrada** (prompt, schema e pacote JSON já minimizado) e **1.500 tokens de saída** (veredito estruturado). Foi aplicado um buffer de **30%** para retries limitados, variação de tamanho e overhead de chamada. O preço de referência da API é o preço padrão por 1 milhão de tokens exibido na página oficial consultada em **25/08/2026** [1]. Para conversão ilustrativa, usa-se **US$ 1 = R$ 5,50**, valor que deve ser substituído pelo câmbio e encargos reais na modelagem financeira.

| Modelo de referência | Entrada / 1M tokens | Saída / 1M tokens | Custo-base por análise | Custo com buffer de 30% | Aproximação em BRL com câmbio ilustrativo |
| --- | ---: | ---: | ---: | ---: | ---: |
| `gpt-5.6-luna` | US$ 0,20 | US$ 1,20 | US$ 0,0058 | **US$ 0,00754** | **R$ 0,0415** |
| `gpt-5.6-terra` | US$ 2,00 | US$ 12,00 | US$ 0,0580 | **US$ 0,0754** | **R$ 0,4147** |
| `gpt-5.6-sol` | US$ 4,00 | US$ 20,00 | US$ 0,1100 | **US$ 0,1430** | **R$ 0,7865** |

A tabela não deve ser interpretada como cotação fixa. Preços, modelos, limites, impostos e políticas podem mudar; o cliente deve consultar a tabela vigente do provedor na hora de contratar. O cálculo também pressupõe que o pacote foi minimizado: enviar linhas de comando extensas, eventos numerosos ou dados `full` pode elevar o custo significativamente.

## 2. Opção A — Compra única de baixo custo

Nesta opção, o cliente compra uma licença, versão ou pacote de análises por um valor único. A proposta combina com um sintoma ocasional: o usuário percebe a janela no boot, roda a coleta, recebe o relatório e não precisa de histórico contínuo. A arquitetura deve separar o direito de usar o executável do custo variável da chamada de API.

| Elemento | Desenho possível, sujeito a aprovação |
| --- | --- |
| Entrega | Executável assinado + relatório local + número limitado de vereditos remotos, por exemplo 1 ou 3 análises. |
| API | Embutida no preço até um limite; acima do limite, compra adicional ou modo local. `luna` reduz custo; `terra` pode elevar qualidade/robustez após validação. |
| Preço ilustrativo | R$ 19,90 por análise única ou R$ 39,90 por pacote de 3; estes valores não são recomendação final. |
| Ponto de equilíbrio só da API | A R$ 19,90, cobre aproximadamente 479 análises `luna`, 48 análises `terra` ou 25 análises `sol` antes de outros custos. |

**Prós.** A cobrança é simples, alinha pagamento com um problema concreto, evita compromisso mensal e limita a exposição do fornecedor a custo recorrente de API. Também pode ser adequada para o usuário doméstico que só precisa de uma triagem pontual.

**Contras.** A receita é menos previsível, aumenta a fricção para suporte e atualizações e pode incentivar o cliente a compartilhar uma licença. O produto tem pouco espaço para histórico, telemetria consentida e acompanhamento. Se o usuário comprar uma vez e usar muitas análises, o limite de custo precisa ser técnico e transparente.

**Ponto de decisão.** Definir se a compra inclui chamadas remotas, quantas, por quanto tempo a licença recebe atualizações e qual é a política quando o provedor fica indisponível.

## 3. Opção B — Assinatura mensal

A assinatura oferece uso continuado, histórico de relatórios e atualização do coletor, desde que isso seja desenhado com retenção e privacidade adequadas. Pode atender melhor prestadores de suporte e pequenas empresas que fazem triagens recorrentes, mas cria expectativa de disponibilidade e suporte.

| Plano ilustrativo | Inclui | Limite sugerido para simulação | Custo variável de API usando `terra` |
| --- | --- | ---: | ---: |
| Individual | Coleta, relatório leigo, até 5 vereditos/mês | 5 | R$ 2,07/mês |
| Profissional | Histórico local/conta, até 25 vereditos/mês, apêndice técnico | 25 | R$ 10,37/mês |
| Pequena equipe | Até 100 vereditos/mês, controles de acesso e exportação | 100 | R$ 41,47/mês |

Os valores de API acima são apenas `quantidade × R$ 0,4147` e não incluem custos de plataforma. Se `luna` for suficiente após validação dos casos sintéticos e humanos, o custo variável de 25 análises seria aproximadamente R$ 1,04; se `sol` for escolhido, seria aproximadamente R$ 19,66. A escolha de modelo deve ser validada com métricas de qualidade e não apenas com preço.

**Prós.** Receita previsível, financiamento de atualizações/assinatura de código/suporte, possibilidade de histórico e experiência melhor para prestadores recorrentes. Limites mensais tornam o gasto da API previsível e permitem pausar análise em vez de gerar custo sem controle.

**Contras.** Maior barreira de compra para um sintoma ocasional, obrigação de cumprir expectativas de disponibilidade e necessidade de governar histórico e contas. Também há risco de abandono se o produto não tiver uso recorrente suficiente.

**Ponto de decisão.** Definir limites, preço, política de excedente, retenção do histórico, número de dispositivos, suporte, cancelamento e se relatórios armazenados na nuvem serão opcionais. Uma assinatura sem necessidade frequente pode ser percebida como desproporcional.

## 4. Opção C — Freemium

No freemium, a coleta local e um veredito básico são gratuitos, enquanto o relatório técnico completo, histórico e recursos de exportação são pagos. A opção pode reduzir a fricção inicial e permitir validação com usuário leigo, mas o fornecedor precisa controlar abuso e custo da camada grátis.

| Camada | Inclui | Exemplo de limite | Modelo/custo de referência |
| --- | --- | ---: | --- |
| Grátis | Coleta somente-leitura, prévia local e resumo básico | 3 análises por dispositivo/mês | `luna`: cerca de R$ 0,12/mês de API no limite. |
| Individual pago | Apêndice técnico, histórico local/conta e exportação | 20 análises/mês | `terra`: cerca de R$ 8,29/mês no limite. |
| Profissional pago | Histórico de equipe, retenção configurável e suporte | 100 análises/mês | `terra`: cerca de R$ 41,47/mês no limite. |

**Prós.** Permite que a pessoa teste a proposta sem pagar, favorece a compreensão do usuário leigo e cria funil para suporte profissional. O relatório básico pode demonstrar valor antes de solicitar dados para a camada técnica.

**Contras.** A camada grátis pode atrair automação abusiva, gerar custo desproporcional e exigir rate limiting, quotas por dispositivo/conta e monitoramento. Diferenciar “básico” de “técnico” sem degradar a segurança é difícil; o usuário deve saber antes do envio o que será processado e armazenado. O free tier não deve enviar mais dados do que o necessário para vender o plano pago.

**Ponto de decisão.** Definir se o veredito básico é local ou remoto, quantos usos gratuitos são sustentáveis, como tratar dispositivos compartilhados, quais dados entram no histórico e como permitir exclusão/portabilidade.

## 5. Comparação executiva

| Critério | Compra única | Assinatura mensal | Freemium |
| --- | --- | --- | --- |
| Complexidade comercial | Baixa | Média/alta | Alta |
| Previsibilidade de receita | Baixa | Alta | Média, dependente de conversão |
| Adequação ao sintoma ocasional | Alta | Média/baixa | Alta |
| Controle de custo de API | Por pacote/licença | Por quota mensal | Requer quota grátis + paga |
| Necessidade de histórico | Opcional e local | Mais provável | Diferencial pago |
| Risco de suporte | Pico após incidentes | Contínuo | Alto volume de usuários gratuitos |
| Hipótese de público | Doméstico avançado | Suporte TI/pequena empresa | Aquisição ampla e validação |

## 6. Guardrails econômicos que valem para as três opções

O cliente deve definir um teto de custo por análise e um teto mensal global, com contador local/servidor e bloqueio seguro quando o limite for atingido. Retry deve ser limitado a erros transitórios e nunca repetir uma solicitação por falha de validação causada por conteúdo potencialmente sensível sem revisão. O modo local deve continuar disponível quando não houver consentimento, quota ou conectividade.

A página de compra deve informar que o custo de API é incorporado ao preço ou ao limite de uso, que o relatório é indicativo, que o fornecedor pode alterar modelos/preços com aviso e que nenhum plano autoriza uso em máquinas de terceiros sem permissão. O custo não inclui uma promessa de acurácia, de detecção total ou de resposta a incidentes.

## 7. Decisão pendente

Não decidir por este documento. O proprietário deve escolher uma das três opções após validar: frequência real do sintoma; disposição a pagar; custo efetivo de API com fixtures; custo de suporte; necessidade de histórico; exigências de privacidade; e capacidade de operar quotas. A escolha deve ser registrada em uma decisão de produto e refletida no cliente da API, no aviso de privacidade e no handoff local.

## Referências

[1]: https://developers.openai.com/api/docs/pricing "OpenAI Developers — Pricing"
[2]: https://www.gov.br/anpd/pt-br/centrais-de-conteudo/materiais-educativos-e-publicacoes/guia-orientativo-sobre-seguranca-da-informacao-para-agentes-de-tratamento-de-pequeno-porte "ANPD — Guia orientativo sobre segurança da informação para agentes de tratamento de pequeno porte"
