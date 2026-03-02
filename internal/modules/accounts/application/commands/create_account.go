package commands

import "context"

type CreateAccountCommand struct {
	UserID string
	Label  string
}

type CreateAccountHandler interface {
	Handle(ctx context.Context, cmd CreateAccountCommand) (string, error)
}
