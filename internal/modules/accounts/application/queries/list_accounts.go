package queries

import "context"

type ListAccountsByUserQuery struct {
	UserID string
	Limit  int32
	Offset int32
}

type ListAccountsByUserHandler interface {
	Handle(ctx context.Context, q ListAccountsByUserQuery) ([]*AccountDTO, error)
}
