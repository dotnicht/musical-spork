package queries

import "context"

type GetUserQuery struct{ ID string }

type GetUserHandler interface {
	Handle(ctx context.Context, q GetUserQuery) (*UserDTO, error)
}
