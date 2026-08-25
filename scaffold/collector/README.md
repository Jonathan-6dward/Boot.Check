# Collector scaffold

Este diretório contém o esqueleto do coletor Windows do BootCheck. O scaffold **não faz coleta real**: cada categoria retorna um erro `TODO` deliberado para impedir que o artefato seja confundido com uma ferramenta pronta.

A implementação local deve permanecer somente-leitura e deve seguir o contrato em `/docs/ARQUITETURA_TECNICA.md`. O coletor não pode criar processos-filhos, executar comandos, carregar payloads, alterar registro, iniciar/parar/configurar serviços, iniciar/parar/registrar/excluir tarefas, criar/remover WMI subscriptions, limpar eventos, modificar firewall ou abrir conexões de rede.

## Implementação recomendada

O agente local deve implementar uma categoria por sessão, começando por processos e persistência, e adicionar testes antes de conectar cada categoria à orquestração. Falhas de permissão precisam ser representadas em `limitations`. Um pacote parcialmente coletado é válido quando o schema e a integridade estão corretos; o veredito deve considerar as limitações.

A assinatura Authenticode, o build reprodutível, a verificação de hash e os testes em Windows descartável são etapas de release. Não implementar técnicas de evasão de antivírus/EDR. Se qualquer requisito parecer exigir escrita, execução ou remediação, parar e solicitar aprovação humana explícita.
