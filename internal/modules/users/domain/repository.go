package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id UserID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	List(ctx context.Context, limit, offset int) ([]*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id UserID) error
}
