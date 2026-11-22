# Project

TicketFlow is an event ticketing platform built with microservices. Organizers publish events (concerts, festivals, theatre, conferences); customers buy digital tickets with a QR code; staff validate tickets at the venue entrance.

This is the final project of the course **Serviços e Interfaces para a Cloud** (Licenciatura em Tecnologias e Desenvolvimento Web, 2026/2027). It is built by a group of 3.

## Users

- **Customer**: searches events, reserves and buys tickets, sees their tickets.
- **Organizer**: manages venues, events and ticket types; sees sales; validates tickets at the entrance.
- **Admin**: manages users and roles, moderates events.

## Main flow

1. Customer logs in and gets a JWT.
2. Customer reserves N tickets of a ticket type. Seats are held for 10 minutes.
3. Customer pays (simulated). On success the order is confirmed; on failure or timeout seats are released.
4. Tickets are issued with a unique QR code and the customer is notified.
5. At the venue, the QR code is validated once and the ticket becomes `used`.

## What the course requires (summary)

- At least 4 services with clear responsibilities (we have 6).
- At least one service with a different language and/or database.
- Layered architecture in every service: routes, controllers, services, models.
- RESTful CRUD in every service, documented with Swagger/OpenAPI.
- User registration and login with JWT.
- GraphQL in at least one service (Events).
- Single entry point for the whole app (API Gateway).
- Logging strategy.
- Postman collection with all requests.
- Docker containers + Docker Swarm, high availability (no single point of failure), scaling without manual config changes.
- Bonus: async messaging (RabbitMQ) and tests.

## Deadlines

| Milestone | Date |
| --- | --- |
| First delivery (report + code skeleton) | 2026-12-04 |
| Final delivery (code + README + report) | 2026-12-31 |
| Defense and demo | 2027-01-12 / 2027-01-13 |

The demo must show: API calls, Swagger pages, scaling with Docker Swarm, and a simulated service failure.
