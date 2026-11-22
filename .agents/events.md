# RabbitMQ events

One topic exchange: `ticketflow.events`. Routing key = event name. Each consumer has its own durable queue bound to the keys it needs.

| Event | Published by | Consumed by | Effect |
| --- | --- | --- | --- |
| `order.created` | orders | notifications | Tell the customer the reservation is held |
| `payment.succeeded` | payments | orders, notifications | Orders confirms the order |
| `payment.failed` | payments | orders, notifications | Orders releases seats in Events |
| `order.confirmed` | orders | tickets | Tickets issues the tickets |
| `order.cancelled` | orders | payments, notifications | Payments refunds |
| `tickets.issued` | tickets | notifications | Send tickets to the customer |
| `event.updated` | events | notifications | Tell ticket holders the event changed |

## Message envelope

Every message uses the same JSON envelope:

```json
{
  "id": "uuid",
  "type": "order.confirmed",
  "occurredAt": "2026-12-10T18:30:00Z",
  "correlationId": "uuid from the original HTTP request",
  "data": { "orderId": "...", "userId": "...", "items": [] }
}
```

## Rules

- Consumers must be **idempotent**: the same message may arrive twice. Use `id` to skip duplicates.
- Ack only after the work is saved. On error, nack and let it go to a dead-letter queue (`<queue>.dlq`).
- Always forward `correlationId` into logs and into any message you publish as a result.
- Adding a new event: add a row to the table above first.
