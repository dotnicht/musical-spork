package commands

import "context"

type DeleteAccountCommand struct {
	ID string
}

type DeleteAccountHandler interface {
	Handle(ctx context.Context, cmd DeleteAccountCommand) error
}
