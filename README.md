# TicketFlow

Event ticketing platform built with microservices. Organizers publish events, customers buy digital tickets with a QR code, and staff validate them at the venue entrance.

Final project for **Serviços e Interfaces para a Cloud** — Licenciatura em Tecnologias e Desenvolvimento Web, 2026/2027.

## Architecture

```mermaid
flowchart LR
    C[Client / Postman] -->|HTTPS + JWT| GW[API Gateway - Traefik]
    GW --> U[Users]
    GW --> E[Events]
    GW --> O[Orders]
    GW --> P[Payments]
    GW --> T[Tickets]
    GW --> N[Notifications]
    O -->|REST| E
    P -->|REST| O
    E -->|pub| MQ[(RabbitMQ)]
    O <-->|pub/sub| MQ
    P <-->|pub/sub| MQ
    T <-->|pub/sub| MQ
    MQ -->|sub| N
```

All requests go through the Traefik gateway, which validates the JWT. Each service owns its database. Services talk through REST only in two places (Orders → Events, Payments → Orders); everything else goes through RabbitMQ.

| Service | Stack | Database | Interfaces |
| --- | --- | --- | --- |
| users | Node.js + Express | PostgreSQL | REST |
| events | Node.js + Express + Apollo Server | PostgreSQL | REST + GraphQL |
| orders | Go + Gin | PostgreSQL | REST |
| tickets | Go + Gin | PostgreSQL | REST |
| payments | Node.js + Express | MongoDB | REST |
| notifications | Node.js + Express | MongoDB | REST + RabbitMQ consumer |

## Prerequisites

- Docker 24+ with Docker Compose v2
- Git
- Postman (to run the API collection)
- Optional, only to run tests outside Docker: Node.js 20+ and Go 1.27+

## Quick start (local, Docker Compose)

```bash
git clone https://github.com/<org>/ticketflow.git
cd ticketflow

# create the .env files from the examples
for s in services/*; do cp "$s/.env.example" "$s/.env"; done

docker compose up -d --build
docker compose ps        # all services should be "healthy"
```

The app is available at **http://localhost**.

## Deploy on Docker Swarm

```bash
# 1. start a Swarm (single node is enough for the demo)
docker swarm init

# 2. build the images
for s in users events orders tickets payments notifications; do
  docker build -t ticketflow/$s:latest services/$s
done

# 3. deploy the stack
docker stack deploy -c docker-stack.yml ticketflow

# 4. check it
docker stack services ticketflow
```

> On a multi-node Swarm, push the images to a registry (e.g. Docker Hub) and update the image names in `docker-stack.yml`, so every node can pull them.

To remove the stack: `docker stack rm ticketflow`.

## Build and run a single service

Every service has its own `Dockerfile` in `services/<service>/`. To build and run one on its own, for example `orders`:

```bash
# build the image (run from the repository root)
docker build -t ticketflow/orders:latest services/orders

# run it, publishing the container port 8080 on localhost:8081
docker run --rm -p 8081:8080 --env-file services/orders/.env ticketflow/orders:latest

# check it
curl http://localhost:8081/api/health
```

The orders image is a multi-stage build: it compiles a static Go binary with `golang:1.27-alpine` and runs it on an empty `scratch` image as a non-root user, with `GIN_MODE=release`.

## URLs

| What | URL |
| --- | --- |
| API (gateway) | http://localhost/api/... |
| Swagger — users | http://localhost/api/users/docs |
| Swagger — events | http://localhost/api/events/docs |
| Swagger — orders | http://localhost/api/orders/docs |
| Swagger — tickets | http://localhost/api/tickets/docs |
| Swagger — payments | http://localhost/api/payments/docs |
| Swagger — notifications | http://localhost/api/notifications/docs |
| GraphQL (events) | http://localhost/api/events/graphql |
| Traefik dashboard | http://localhost:8080 |
| RabbitMQ management | http://localhost:15672 (guest / guest) |

## Try it

```bash
# register
curl -X POST http://localhost/api/users/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Ana","email":"ana@example.com","password":"secret123"}'

# login and keep the token
TOKEN=$(curl -s -X POST http://localhost/api/users/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"ana@example.com","password":"secret123"}' | jq -r .token)

# list events
curl http://localhost/api/events -H "Authorization: Bearer $TOKEN"
```

### Postman

1. Import `postman/ticketflow.postman_collection.json`.
2. Import an environment from `postman/environments/` (`local` or `swarm`) and select it.
3. Run **Users → Login** first; the token is saved to `{{token}}` automatically.
4. To see the whole purchase flow, run the **Flows** folder with the Collection Runner.

## Demo: scaling and fault tolerance

```bash
# scale a service up and down (no config changes needed; Traefik picks up new replicas)
docker service scale ticketflow_orders=5
docker service ps ticketflow_orders
docker service scale ticketflow_orders=2

# simulate a failure: kill one container and watch Swarm replace it
docker ps --filter name=ticketflow_events
docker kill <container-id>
docker service ps ticketflow_events

# requests keep working during the failure
curl http://localhost/api/events
```

If Notifications is stopped, purchases still work: its messages wait in RabbitMQ and are processed when it comes back.

## Configuration

Each service reads its config from environment variables. See `services/<service>/.env.example`.

| Variable | Used by | Example |
| --- | --- | --- |
| `PORT` | all | `3000` |
| `DATABASE_URL` | users, events, orders, tickets | `postgres://user:pass@users-db:5432/users` |
| `MONGO_URL` | payments, notifications | `mongodb://payments-db:27017/payments` |
| `RABBITMQ_URL` | events, orders, tickets, payments, notifications | `amqp://guest:guest@rabbitmq:5672` |
| `JWT_SECRET` | all | `change-me` |
| `LOG_LEVEL` | all | `info` |

## Logging

All services write JSON logs to stdout with a shared `correlationId`, so one request can be followed across services:

```bash
docker service logs ticketflow_orders --follow
docker service logs ticketflow_orders 2>&1 | grep <correlationId>
```

## Tests

Tests run on every push and pull request (`.github/workflows/tests.yml`). To run them locally:

```bash
# Node services
cd services/users && npm install && npm test

# Go services (with race detector)
cd services/orders && go test -race ./...
```

## Project structure

```
ticketflow/
├── .agents/            # context for AI coding agents (see AGENTS.md)
├── .github/workflows/  # CI
├── gateway/            # Traefik configuration
├── docs/               # report and diagrams
├── postman/            # API collection and environments
├── services/
│   ├── users/          # Node.js
│   ├── events/         # Node.js + GraphQL
│   ├── orders/         # Go
│   ├── tickets/        # Go
│   ├── payments/       # Node.js
│   └── notifications/  # Node.js
├── docker-compose.yml  # local development
├── docker-stack.yml    # Docker Swarm
└── README.md
```

Node services use `src/{routes,controllers,services,models}`; Go services use `internal/{routes,handlers,services,models}`.

## Team

| Name | Responsibilities |
| --- | --- |
| A | users, events, API gateway, GraphQL |
| B | orders, tickets, logging |
| C | payments, notifications, RabbitMQ, Docker Swarm |

## Documentation

- Report: `docs/`
- Architecture decisions and conventions: `.agents/`
