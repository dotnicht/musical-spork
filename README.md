# Go gRPC API (DDD + CQRS + Modular Monolith) + Postgres (GORM)

This scaffold demonstrates:
- **Modular monolith** with multiple modules (bounded contexts): `users` and `accounts`
- **DDD** layers per module: `domain / application / infrastructure / interfaces`
- **CQRS** separation: commands vs queries handlers
- **In-process module communication** via **public contracts** (`internal/modules/users/public`)
- **gRPC** API endpoints for both modules
- **Postgres** persistence via **GORM**

## Quick start
```bash
docker compose up -d
go mod tidy
go run ./cmd/server
```

## Users (grpcurl)
```bash
grpcurl -plaintext localhost:50051 api.users.v1.UsersService/CreateUser -d '{"email":"a@b.com","name":"Alice"}'
grpcurl -plaintext localhost:50051 api.users.v1.UsersService/ListUsers -d '{"limit":50,"offset":0}'
```

## Accounts (grpcurl)
Accounts are created **for an existing user** (accounts module calls users module in-process via an interface).
```bash
# create user, grab id
grpcurl -plaintext localhost:50051 api.users.v1.UsersService/CreateUser -d '{"email":"x@y.com","name":"X"}'

# create account for that user id
grpcurl -plaintext localhost:50051 api.accounts.v1.AccountsService/CreateAccount -d '{"user_id":"<user-id>","label":"Primary"}'
```

## Module boundaries (important idea)
- `accounts` depends on `users/public` **only** (a narrow contract).
- `accounts` does NOT import `users/domain` or `users/infrastructure` or `users/interfaces`.

