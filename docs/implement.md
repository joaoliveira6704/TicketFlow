# TicketFlow — Plano de Implementação

Grupo de 3 · de 22/11/2026 até à defesa (12–13/01/2027)

| Marco | Data |
| --- | --- |
| Primeira entrega | 04/12/2026 |
| Entrega final | 31/12/2026 |
| Defesa e demonstração | 12 e 13/01/2027 |

## Divisão de responsabilidades

Cada membro é dono de dois serviços e de uma parte transversal. A carga fica equilibrada: o **A** tem a autenticação e o GraphQL, o **B** tem os dois serviços em Go e a parte de concorrência, e o **C** tem os dois serviços mais simples, mas fica com a mensageria e a infraestrutura do Swarm.

| Membro | Serviços | Parte transversal | Alíneas principais |
| --- | --- | --- | --- |
| **A** | Users (Node), Events (Node + GraphQL) | API Gateway (Traefik), middleware JWT partilhado | 9, 10, 11 |
| **B** | Orders (Go), Tickets (Go) | Logging e `correlationId`, concorrência na BD | 4, 12 |
| **C** | Payments (Node), Notifications (Node) | RabbitMQ, `docker-stack.yml`, alta disponibilidade e escala | 14, 15, 16 |

Regras para todos:

- Todo o código em inglês: variáveis, funções, tabelas e colunas, campos JSON, comentários e mensagens de commit. O relatório e o README ficam em português.
- Cada membro faz, nos **seus** serviços: arquitetura em camadas (5), base de dados (6), CRUD REST (7), Swagger (8), pedidos na coleção Postman (13) e testes (bónus).
- No relatório, cada um escreve as secções da sua parte transversal e os diagramas de sequência que lhe são atribuídos.
- Nada entra na `main` sem um *pull request* revisto por outro membro.

---

## Fase 0 — Arranque (22/11 a 24/11)

Objetivo: todos conseguem correr o projeto localmente com a mesma estrutura.

| A | B | C |
| --- | --- | --- |
| Criar o repositório e a estrutura de pastas; proteger a `main` | Definir o formato dos logs em JSON e o cabeçalho `X-Correlation-Id` | Criar o `docker-compose.yml` de desenvolvimento (BDs + RabbitMQ) |
| Definir convenções: nomes de rotas, códigos de estado, formato de erros | Preparar o template de serviço Go (camadas, `cmd/server/main.go`) | Preparar o template de serviço Node (camadas, `app.js`) |

Entregável: `README.md` provisório com "como correr em local".

---

## Fase 1 — Primeira entrega (25/11 a 04/12)

Objetivo: o código base de todos os serviços e o relatório inicial.

| A | B | C |
| --- | --- | --- |
| Esqueleto Users e Events: camadas, Dockerfile, `GET /health` | Esqueleto Orders e Tickets: camadas, Dockerfile, `GET /health` | Esqueleto Payments e Notifications: camadas, Dockerfile, `GET /health` |
| Relatório: contexto e operações principais (1) | Relatório: divisão em serviços e tecnologias (2, 4) | Relatório: diagrama de arquitetura e estrutura do repositório (2, 3) |
| Rever o relatório completo | Rever o relatório completo | Juntar tudo e submeter |

Checklist da entrega:

- [ ] Os 6 serviços arrancam com `docker compose up` e respondem em `/health`
- [ ] Relatório com contexto, operações, requisitos funcionais e diagrama
- [ ] Link do repositório no relatório

---

## Fase 2 — Serviços e dados (05/12 a 15/12)

Objetivo: cada serviço funciona sozinho, com CRUD, BD e documentação.

| A | B | C |
| --- | --- | --- |
| Users: registo, login, emissão de JWT, perfis, papéis (RF01–RF04) | Orders: CRUD de encomendas, reserva com expiração de 10 min (RF10–RF12, RF20) | Payments: pagamento simulado e reembolso (RF13, RF14) |
| Events: CRUD de recintos, eventos e tipos de bilhete, pesquisa (RF05–RF08) | Tickets: emissão com QR, consulta e validação única (RF15–RF17) | Notifications: CRUD e histórico de notificações (RF19) |
| Middleware JWT partilhado (Node) e verificação de papéis | Middleware JWT em Go; atualizações atómicas contra *overselling* | Swagger nos seus serviços |
| Swagger nos seus serviços | Swagger nos seus serviços (`swaggo`) | Pedidos Postman dos seus serviços |
| Pedidos Postman dos seus serviços | Pedidos Postman dos seus serviços | |

Ponto de controlo (15/12): demonstração interna de cada serviço isolado, via Swagger.

---

## Fase 3 — Integração (16/12 a 23/12)

Objetivo: os serviços falam entre si e o fluxo de compra funciona de ponta a ponta.

| A | B | C |
| --- | --- | --- |
| Traefik como ponto único de entrada, com rotas para todos os serviços | Chamada síncrona Orders → Events (reservar lugares) | Configurar o RabbitMQ: *exchanges*, filas duráveis, *acks* |
| Autenticação no gateway (*forwardAuth* para o Users) | Publicar `order.created`, `order.confirmed`, `order.cancelled`; consumir `payment.*` | Payments: chamada Payments → Orders, publicar `payment.*`, consumir `order.cancelled` |
| GraphQL no Events: evento + recinto + disponibilidade (RF09) | Tickets: consumir `order.confirmed`, publicar `tickets.issued` | Notifications: consumir todos os eventos (RF18) |
| Events: publicar `event.updated` | Logging com `correlationId` em todos os serviços (ajudar A e C) | |

Ponto de controlo (23/12): uma compra completa, do login ao bilhete emitido, feita pelo Postman.

---

## Fase 4 — Infraestrutura e entrega final (24/12 a 31/12)

Objetivo: tudo corre em Docker Swarm, sem pontos únicos de falha, e a documentação está fechada.

| A | B | C |
| --- | --- | --- |
| Gateway com 2+ réplicas; testar falha de uma réplica | Testes unitários e de integração do Orders e do Tickets (bónus) | `docker-stack.yml` com réplicas, *healthchecks* e *restart policy* |
| Testes dos seus serviços (bónus) | Teste de carga simples à reserva para provar que não há *overselling* | Alta disponibilidade das BDs e do RabbitMQ; análise de SPOF (15) |
| Coleção Postman final, organizada por fluxo (13) | Relatório: logging, observabilidade, requisitos não funcionais (12, 19) | Escalar com `docker service scale` sem mexer na configuração (16) |
| Relatório: padrões de microserviços e comunicação (17) | Diagramas de sequência: reserva e pagamento, validação à entrada (18) | README com instruções de execução completas (21) |
| Diagramas de sequência: registo e login, publicação de evento (18) | | Relatório: benefícios e desafios vs. monólito, uso de IA (20) |

Os três reveem e juntam o relatório final (22) até 30/12, e submetem a 31/12.

Checklist da entrega final:

- [ ] `docker stack deploy` põe tudo a funcionar a partir do README
- [ ] Swagger acessível em todos os serviços através do gateway
- [ ] Coleção Postman cobre todos os pedidos
- [ ] Relatório com todos os pontos da alínea 22

---

## Fase 5 — Preparação da defesa (01/01 a 11/01)

Objetivo: uma demonstração sem surpresas, com todos capazes de responder sobre todo o sistema.

| A | B | C |
| --- | --- | --- |
| Guião da demo: chamadas à API e Swagger | Guião da demo: compra concorrente e logs com `correlationId` | Guião da demo: escalar serviços e simular uma falha |
| Explicar aos outros: JWT, gateway, GraphQL | Explicar aos outros: Go, concorrência, logging | Explicar aos outros: Swarm, RabbitMQ, alta disponibilidade |

Dois ensaios completos da defesa (por exemplo, 06/01 e 10/01), cronometrados, com perguntas cruzadas entre membros.

---

## Riscos e como os reduzir

| Risco | Mitigação |
| --- | --- |
| Atrasos no Go por ser uma linguagem nova | B começa pelo template na Fase 0; A e C ajudam a rever o código |
| A integração só falha no fim | Ponto de controlo a 23/12 com o fluxo completo; eventos documentados desde a Fase 1 |
| O Swarm comporta-se de forma diferente do Compose | C testa a stack no Swarm já na Fase 3, mesmo incompleta |
| Período de festas (24/12 a 01/01) | Fechar o essencial até 23/12; a Fase 4 fica para infraestrutura e documentação |
