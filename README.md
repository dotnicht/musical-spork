# Go API: gRPC + REST(JSON) side-by-side (DDD + CQRS + Modular Monolith) + Postgres (GORM)

This scaffold demonstrates:
- **Modular monolith** with multiple bounded contexts: `users` and `accounts`
- **DDD** per module: `domain / application / infrastructure / interfaces`
- **CQRS**: commands and queries handlers
- **Two transports side-by-side**:
  - gRPC (for internal/high-performance clients)
  - REST/JSON over HTTP (for non-gRPC clients)
- Transport adapters are **thin**: they only map requests -> application commands/queries and map errors.

## Quick start
```bash
docker compose up -d
go mod tidy
go run ./cmd/server
```

### gRPC
- listens on `:50051` by default

### REST/JSON
- listens on `:8080` by default

## REST endpoints (Users)
- `POST /v1/users`
- `GET /v1/users/{id}`
- `GET /v1/users?limit=50&offset=0`
- `PATCH /v1/users/{id}`
- `DELETE /v1/users/{id}`

Example:
```bash
curl -s -X POST http://localhost:8080/v1/users \
  -H 'content-type: application/json' \
  -d '{"email":"a@b.com","name":"Alice"}'

curl -s http://localhost:8080/v1/users?limit=50&offset=0
```

## Accounts module boundary
`accounts` calls `users` **in-process via contract** `internal/modules/users/public` only.
