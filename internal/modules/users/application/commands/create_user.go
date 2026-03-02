package commands

import "context"

type CreateUserCommand struct {
	Email string
	Name  string
}

type CreateUserHandler interface {
	Handle(ctx context.Context, cmd CreateUserCommand) (string, error)
}
