# Arquitetura Hexagonal

### Pontos importantes sobre arquitetura

- Crescimento sustentável;
- Software precisa se pagar ao passar do tempo;
- Peças precisam se encaixar e eventualmente substituídas;
- Software deve ser desenhado por você e não pelo seu framework;

### Lembre-se!
> Arquitetura é respeitar o futuro do seu software. **Crud, qualquer um faz!**


### Ciclo de vida de muitos projetos

- Fase 1:
  - Banco de dados;
  - Cadastros;
  - Validações;
  - Servidor web;
  - Controllers;
  - Views;
  - Autenticação; 

- Fase 2:
  - Novas regras de negócio;
  - Criação de APIs;
  - Consumo de APIs;
  - Autorização (ACL);
  - Relatórios;
  - Logs;

- Fase 3:
  - Mais acessos;
  - Upgrades de hardware;
  - Cache;
  - Consumo de API Parceiros;
  - Regras de parceiros;

- Fase 4:
  - Mais acessos;
  - Mais upgrade de hardware;
  - BD relatórios;
  - Comandos;
  - V2 da API;

- Fase 5:
  - Escala horizontal (+ máquinas);
  - Sessões;
  - Uploads;
  - Refatoração;
  - Autoscaling;
  - CI/CD;

- Fase 6:
  - GraphQL;
  - Bugs Constantes;
  - Logs;
  - Integração CRM;
  - Migração para React;

- Fase 7:
  - Inconsistência CRM;
  - Containers;
  - CI/CD;
  - Memória;

- Fase 8:
  - Microsserviços;
  - DB Compartilhado;
  - Problemas com tracing;
  - Lentidão;
  - Custo elevado;

- Fase 9:
  - K8S;
  - CI/CD;
  - Mensageria;
  - Perda de mensagens;
  - Consultoria;

