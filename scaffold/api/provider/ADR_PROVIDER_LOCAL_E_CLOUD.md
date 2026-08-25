# ADR: Suporte a dois Analysis Providers (Local e Cloud)

## Status
Proposto — depende de revisão antes de entrar na Fase 05 do roadmap.

## Contexto
O README define o "Analysis provider" como uma fronteira única que recebe
o `EvidencePackage` redigido e devolve um veredito estruturado. A tabela de
status classifica essa camada como `NEXT`: cliente, timeout e retry
existem no scaffold, mas o envelope de um provedor específico ainda não
foi conectado.

Foi pedido suporte a **duas opções de análise**, para o usuário escolher:

1. **Local** — LLM rodando na própria máquina, sem dado saindo da rede.
2. **Browser/Cloud** — o usuário conecta a própria conta (Anthropic,
   OpenAI ou Gemini), inclusive camada gratuita. O BootCheck nunca
   intermedeia ou armazena a chave em serviço próprio.

## Decisão
Introduzir uma interface `Provider` única (`scaffold/api/provider`) com
dois adapters:

- `LocalOllamaProvider` — fala com Ollama em `localhost:11434`, usa
  `format: "json"` para saída estruturada, timeout de 90s (CPU local
  pode ser lento).
- `CloudProvider` — multi-vendor (Anthropic / OpenAI / Gemini), chave de
  API do próprio usuário, timeout de 30s.

Ambos implementam:

```go
type Provider interface {
    Kind() Kind
    Name() string
    Available(ctx context.Context) error
    Analyze(ctx context.Context, pkg EvidencePackage) (Verdict, error)
}
```

A orquestração (`Previa -> Consentimento -> AnaliseRemota` no state
diagram existente) não muda de forma — só passa a escolher qual
`Provider` invocar com base na seleção do usuário.

## Consequências

**Consentimento diferenciado por Kind.**
`KindLocal` não aciona a tela de consentimento de rede do
`docs/RESPONSABILIDADE_LEGAL.md`, porque nenhum dado sai da máquina.
`KindCloud` aciona a tela normalmente, exibindo o vendor escolhido,
o modelo e o fato de que o processamento segue a política de dados
daquele provedor — não a do BootCheck.

**Chave de API nunca passa pela infraestrutura do BootCheck.**
`CloudCredentials.APIKey` é fornecida pelo usuário e usada só na
chamada HTTPS direta ao vendor. A camada de configuração local deve
criptografar a chave em repouso (DPAPI no Windows) — isso é tarefa
separada, não coberta por este ADR.

**Falha de schema é tratada como estado, não como exceção.**
`ValidateVerdict` rejeita: estado fora de
`likely_safe|suspicious|inconclusive`, e qualquer `evidence_id` citado
em `claims` que não exista no pacote original. Ambos os providers devem
mapear essa falha para `ErrSchemaInvalid`, e a orquestração deve tratar
isso como veredito `Inconclusivo` — nunca travar a UI.

**Modelos locais erram schema com mais frequência.**
Isso é esperado e coberto pelo teste `TestValidateVerdict_EvidenceIdOrfao`
e pelo estado `Inconclusivo` já existente no diagrama — não é motivo
para adicionar um caminho de erro novo.

**Mitigação de prompt injection é compartilhada.**
`buildEvidencePrompt` isola cada `Evidence.Value` dentro de um bloco
`<evidence_package>` delimitado, com instrução explícita de que texto
ali dentro é dado, não comando. Isso vale igualmente para local e cloud,
porque os valores vêm de cmdline/paths do host investigado e, num
cenário real de comprometimento, podem ter sido manipulados de propósito
por quem os criou.

**Cota excedida em camada gratuita tem erro dedicado.**
`ErrQuotaExceeded` é distinto de `ErrUnavailable` para que a UI mostre
uma mensagem específica ("limite do provedor atingido") em vez de um
erro genérico de rede.

## Alternativas consideradas

- **Só cloud (sem local):** rejeitado porque o pedido explícito era
  permitir análise sem dependência de rede/conta externa, além de
  reduzir custo para o cliente.
- **Só um vendor cloud (ex. só Anthropic):** rejeitado porque o
  requisito era permitir que o usuário use a própria conta em serviços
  com camada gratuita — travar em um único vendor reduz essa opção.
- **Abstração via LangChain/framework externo:** rejeitado por ora para
  manter o scaffold sem dependências externas além da stdlib de Go,
  consistente com o restante do projeto.

## Itens ainda em aberto (não cobertos por este ADR)
- Criptografia da chave de API em disco (DPAPI).
- Tela de setup/onboarding para escolher provider e inserir chave.
- Lista de modelos recomendados por vendor e por Kind, com trade-off de
  custo/velocidade/qualidade de schema.
- Grammar-constrained decoding (GBNF) para reduzir ainda mais falha de
  schema em modelos locais pequenos.
