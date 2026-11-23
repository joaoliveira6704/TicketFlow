# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.27-alpine AS builder
WORKDIR /src

# Download dependencies first so this layer is cached between code changes
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Static binary (no cgo) so it runs on the empty scratch image
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/orders ./cmd/server

# Runtime stage
FROM scratch
# CA certificates for outbound TLS (e.g. calls to other services or managed DBs)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/orders /orders

ENV GIN_MODE=release
EXPOSE 8080
# Run as an unprivileged user
USER 65532:65532

ENTRYPOINT ["/orders"]
