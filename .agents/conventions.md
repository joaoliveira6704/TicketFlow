# Conventions

## Language

English everywhere: identifiers, DB tables and columns (`snake_case`), JSON fields (`camelCase`), comments, commits, docs.

## Layers

| Layer | Node (`src/`) | Go (`internal/`) | Does | Never does |
| --- | --- | --- | --- | --- |
| Routes | `routes/` | `routes/` | Maps URI + method to a controller, attaches middlewares | Business logic |
| Controllers | `controllers/` | `handlers/` | Parses/validates input, calls a service, sets status code | SQL, calling other services |
| Services | `services/` | `services/` | Business rules, calls models, other services, publishes events | Touch `req`/`res` |
| Models | `models/` | `models/` | DB access | Business rules |

Cross-cutting code goes in `middlewares/` (JWT, correlation id, error handler), `config/` (env vars, DB connection) and `messaging/` (RabbitMQ).

## Configuration

Read only from environment variables: `PORT`, `DATABASE_URL` or `MONGO_URL`, `RABBITMQ_URL`, `JWT_SECRET`, `LOG_LEVEL`. Document each one in `.env.example`.

## Errors

All services return the same error body:

```json
{ "error": { "code": "SOLD_OUT", "message": "Not enough seats available" } }
```

## Logging

- JSON logs to stdout. Node: Pino. Go: zerolog.
- Every log line has: `timestamp`, `level`, `service`, `correlationId`, `message`.
- Read `X-Correlation-Id` from the request (create one if missing), return it in the response, pass it to every outgoing HTTP call and RabbitMQ message.
- Levels: `info` for requests and business events, `warn` for handled problems (sold out, invalid token), `error` for failures. Never log passwords or tokens.

## API docs

- Node: `swagger-jsdoc` + `swagger-ui-express`, served at `/docs`.
- Go: `swaggo`, served at `/docs`.
- Spec must match the real endpoints. Update it in the same commit.

## Tests

- Node: Jest + Supertest. Go: standard `testing` package; run with `-race`.
- Unit tests for services layer; integration tests for routes against the CI databases.
- CI env vars: `DATABASE_URL`, `MONGO_URL`, `RABBITMQ_URL`, `JWT_SECRET` (see `.github/workflows/tests.yml`).

## Git

- Branches: `feat/<service>-<topic>`, `fix/<service>-<topic>`.
- Commits: Conventional Commits, e.g. `feat(orders): expire pending reservations`.
- Every PR reviewed by another team member; CI must be green.
