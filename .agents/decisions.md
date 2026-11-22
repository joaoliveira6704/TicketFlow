# Decisions

Short log of architecture decisions. Newest first. Add a row before changing the stack or architecture.

| Date | Decision | Why |
| --- | --- | --- |
| 2026-11-22 | Orders and Tickets in Go; other services in Node.js | Go handles high concurrency well (sales peaks, entrance validation) and satisfies the "different stack" requirement with Go + PostgreSQL |
| 2026-11-22 | Overselling prevented in the database with atomic `UPDATE ... WHERE` | Goroutines give throughput, not correctness across replicas |
| 2026-11-22 | Postman for the API collection | The course asks for a Postman collection |
| 2026-11-22 | RabbitMQ with one topic exchange | Async bonus requirement; decouples Notifications and Tickets from the purchase path |
| 2026-11-22 | Traefik as API Gateway | Discovers Swarm services automatically, so scaling needs no config change |
| 2026-11-22 | Monorepo, one folder per service | Shared `docker-stack.yml`, Postman collection and CI in one place |
| 2026-11-22 | Database per service (PostgreSQL / MongoDB) | Services stay independent and can scale and fail separately |
