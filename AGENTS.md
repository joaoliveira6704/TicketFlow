# AGENTS.md

Instructions for AI coding agents (Claude Code, Copilot, Cursor, Codex) working in this repository.

## Read first

Before changing code, read the files in `.agents/` that apply to the task:

| File | Read it when |
| --- | --- |
| `.agents/project.md` | Always: what TicketFlow is and what the course requires |
| `.agents/architecture.md` | Touching more than one service, the gateway, or infrastructure |
| `.agents/services.md` | Working on any service: responsibilities, stack, endpoints |
| `.agents/events.md` | Publishing or consuming RabbitMQ messages |
| `.agents/conventions.md` | Always: code style, layers, API rules, logging, tests |
| `.agents/decisions.md` | Before proposing a change to the stack or architecture |

## Hard rules

- Everything in English: code, comments, identifiers, commit messages, docs in `.agents/`.
- Respect the layered architecture: `routes → controllers/handlers → services → models`. No DB access outside `models/`.
- A service never reads another service's database. Use its REST API or a RabbitMQ event.
- Every new endpoint must be added to the service's Swagger/OpenAPI spec and to the Postman collection in `postman/`.
- Every new feature comes with tests. CI (`.github/workflows/tests.yml`) must pass.
- Do not change the stack (languages, databases, gateway, broker) without updating `.agents/decisions.md` and asking the team.
- Never commit secrets. Use `.env.example` for new variables.

## Commands

```bash
docker compose up -d                          # local dev: all services + DBs + RabbitMQ
docker stack deploy -c docker-stack.yml tf    # Docker Swarm
cd services/<node-service> && npm test        # Node tests
cd services/<go-service> && go test -race ./...  # Go tests
```
