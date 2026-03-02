package commands

import "context"

type UpdateUserCommand struct {
	ID    string
	Email *string
	Name  *string
}

type UpdateUserHandler interface {
	Handle(ctx context.Context, cmd UpdateUserCommand) error
}
