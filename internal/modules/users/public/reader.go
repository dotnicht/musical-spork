package public

import (
	"context"
	"errors"

	"example.com/modmonolith/internal/modules/users/application/queries"
	"example.com/modmonolith/internal/modules/users/domain"
)

// reader adapts users CQRS query handler to the public UserReader contract.
type reader struct {
	get queries.GetUserHandler
}

func NewUserReader(get queries.GetUserHandler) UserReader {
	return reader{get: get}
}

func (r reader) Exists(ctx context.Context, userID string) (bool, error) {
	_, err := r.get.Handle(ctx, queries.GetUserQuery{ID: userID})
	if err == nil {
		return true, nil
	}
	if errors.Is(err, domain.ErrUserNotFound) {
		return false, nil
	}
	return false, err
}
