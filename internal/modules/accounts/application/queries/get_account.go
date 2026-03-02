package queries

import "context"

type GetAccountQuery struct{ ID string }

type GetAccountHandler interface {
	Handle(ctx context.Context, q GetAccountQuery) (*AccountDTO, error)
}
