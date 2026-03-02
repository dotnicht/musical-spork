# Go gRPC API (DDD + CQRS + Modular Monolith) + Postgres (GORM)

Scaffold showcasing:
- **Modular monolith** (bounded-context module under `internal/modules/users`)
- **DDD** layers (domain/application/infrastructure/interfaces)
- **CQRS** (commands vs queries)
- **gRPC** API
- **Postgres** via **GORM**
- Simple migration via AutoMigrate (swap for goose/atlas in real projects)

## Quick start
```bash
docker compose up -d
go mod tidy
go run ./cmd/server
```

## Try with grpcurl
```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext -d '{"email":"a@b.com","name":"Alice"}' localhost:50051 api.users.v1.UsersService/CreateUser
```
