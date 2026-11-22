# Services

| Service | Owner | Stack | Database | Interfaces |
| --- | --- | --- | --- | --- |
| users | A | Node.js + Express | PostgreSQL | REST |
| events | A | Node.js + Express + Apollo Server | PostgreSQL | REST + GraphQL |
| orders | B | Go + Gin | PostgreSQL | REST |
| tickets | B | Go + Gin | PostgreSQL | REST |
| payments | C | Node.js + Express | MongoDB | REST |
| notifications | C | Node.js + Express | MongoDB | REST + RabbitMQ consumer |

Replace A / B / C with team member names.

## users

Registration, login, JWT issuing and verification, profiles, roles (`customer`, `organizer`, `admin`).

- Passwords hashed with bcrypt. Email unique (DB constraint).
- Exposes `GET /auth/verify` for Traefik `forwardAuth`: returns 200 + user headers (`X-User-Id`, `X-User-Role`) or 401.

## events

Venues, events, ticket types (name, price, capacity, available seats), search.

- GraphQL endpoint `/graphql`: event with venue and ticket availability in one query.
- Internal endpoint used by Orders to reserve and release seats. Seat changes must be atomic:
  `UPDATE ticket_types SET available_seats = available_seats - $n WHERE id = $id AND available_seats >= $n`.
- Publishes `event.updated` when an event changes or is cancelled.

## orders (Go)

Reservations, orders, reservation expiry (10 minutes), sales stats per event.

- Order states: `pending → confirmed | cancelled | expired`.
- High concurrency during sales peaks. Never check-then-update; rely on the atomic update in Events.
- A background goroutine expires `pending` orders older than 10 minutes and releases seats.

## tickets (Go)

Ticket issuing with a unique code (QR), "my tickets", entrance validation.

- Ticket states: `valid → used | cancelled`.
- Validation must succeed only once:
  `UPDATE tickets SET status = 'used', used_at = now() WHERE code = $code AND status = 'valid'` (0 rows = reject).

## payments

Simulated payments and refunds.

- `POST /payments` with `orderId`; fetches the amount from Orders; result is random or driven by a test card value.
- Stores the full simulated provider response as a document.

## notifications

Consumes all domain events and records/sends notifications (email simulated or logged). History per user.

## Endpoint conventions (all services)

- Prefix through the gateway: `/api/<service>/...`
- Plural nouns, no verbs: `GET /api/events/{id}`, `POST /api/orders`, `PATCH /api/tickets/{code}/validate` is the only allowed action-style route.
- Status codes: 200 read/update, 201 create (+ `Location`), 204 delete, 400 validation, 401 no/invalid token, 403 wrong role, 404 not found, 409 conflict (sold out, already used).
