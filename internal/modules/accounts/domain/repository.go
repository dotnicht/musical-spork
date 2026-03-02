package domain

import "context"

type AccountRepository interface {
	Create(ctx context.Context, a *Account) error
	GetByID(ctx context.Context, id AccountID) (*Account, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*Account, error)
	Update(ctx context.Context, a *Account) error
	Delete(ctx context.Context, id AccountID) error
}
