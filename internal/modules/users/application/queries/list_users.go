package queries

import "context"

type ListUsersQuery struct {
	Limit  int32
	Offset int32
}

type ListUsersHandler interface {
	Handle(ctx context.Context, q ListUsersQuery) ([]*UserDTO, error)
}
