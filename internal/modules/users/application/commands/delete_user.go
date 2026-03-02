package commands

import "context"

type DeleteUserCommand struct {
	ID string
}

type DeleteUserHandler interface {
	Handle(ctx context.Context, cmd DeleteUserCommand) error
}
