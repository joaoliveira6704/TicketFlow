## Summary

<!-- What does this PR change, and why? One or two sentences. -->

## Service(s)

- [ ] users
- [ ] events
- [ ] orders
- [ ] tickets
- [ ] payments
- [ ] notifications
- [ ] gateway / infrastructure (Traefik, RabbitMQ, compose, stack)
- [ ] docs / CI

## Changes

-
-

## How to test

<!-- Commands, requests or Postman steps a reviewer can run. -->

```bash

```

## Checklist

- [ ] Branch named `feat/<service>-<topic>` or `fix/<service>-<topic>`
- [ ] Commits follow Conventional Commits (`feat(orders): ...`)
- [ ] Layers respected: `routes → handlers/controllers → services → models`, no DB access outside `models/`
- [ ] No access to another service's database (REST or RabbitMQ only)
- [ ] New/changed endpoints added to Swagger/OpenAPI and the Postman collection
- [ ] New RabbitMQ events documented in `.agents/events.md`
- [ ] Tests added or updated, and CI is green
- [ ] New env variables added to `.env.example`, no secrets committed
- [ ] Stack/architecture changes recorded in `.agents/decisions.md` and agreed with the team

## Notes for the reviewer

<!-- Open questions, follow-ups, anything out of scope. -->
