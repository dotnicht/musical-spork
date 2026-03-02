package commands

import "context"

type UpdateAccountCommand struct {
	ID    string
	Label *string
}

type UpdateAccountHandler interface {
	Handle(ctx context.Context, cmd UpdateAccountCommand) error
}
