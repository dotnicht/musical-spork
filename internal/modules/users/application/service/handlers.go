package service

import (
	"context"
	"errors"

	"example.com/modmonolith/internal/modules/users/application/commands"
	"example.com/modmonolith/internal/modules/users/application/queries"
	"example.com/modmonolith/internal/modules/users/domain"
)

type Handlers struct {
	Create commands.CreateUserHandler
	Update commands.UpdateUserHandler
	Delete commands.DeleteUserHandler

	Get  queries.GetUserHandler
	List queries.ListUsersHandler
}

type createUserHandler struct{ repo domain.UserRepository }

func (h createUserHandler) Handle(ctx context.Context, cmd commands.CreateUserCommand) (string, error) {
	u, err := domain.CreateNewUser(cmd.Email, cmd.Name)
	if err != nil {
		return "", err
	}
	if existing, _ := h.repo.GetByEmail(ctx, u.Email()); existing != nil {
		return "", domain.ErrEmailAlreadyUsed
	}
	if err := h.repo.Create(ctx, u); err != nil {
		return "", err
	}
	return string(u.ID()), nil
}

type updateUserHandler struct{ repo domain.UserRepository }

func (h updateUserHandler) Handle(ctx context.Context, cmd commands.UpdateUserCommand) error {
	u, err := h.repo.GetByID(ctx, domain.UserID(cmd.ID))
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrUserNotFound
	}
	if cmd.Email != nil {
		if err := u.ChangeEmail(*cmd.Email); err != nil {
			return err
		}
		if existing, _ := h.repo.GetByEmail(ctx, u.Email()); existing != nil && existing.ID() != u.ID() {
			return domain.ErrEmailAlreadyUsed
		}
	}
	if cmd.Name != nil {
		if err := u.Rename(*cmd.Name); err != nil {
			return err
		}
	}
	return h.repo.Update(ctx, u)
}

type deleteUserHandler struct{ repo domain.UserRepository }

func (h deleteUserHandler) Handle(ctx context.Context, cmd commands.DeleteUserCommand) error {
	return h.repo.Delete(ctx, domain.UserID(cmd.ID))
}

type getUserHandler struct{ repo domain.UserRepository }

func (h getUserHandler) Handle(ctx context.Context, q queries.GetUserQuery) (*queries.UserDTO, error) {
	u, err := h.repo.GetByID(ctx, domain.UserID(q.ID))
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, domain.ErrUserNotFound
	}
	return &queries.UserDTO{
		ID:        string(u.ID()),
		Email:     u.Email(),
		Name:      u.Name(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}, nil
}

type listUsersHandler struct{ repo domain.UserRepository }

func (h listUsersHandler) Handle(ctx context.Context, q queries.ListUsersQuery) ([]*queries.UserDTO, error) {
	limit := int(q.Limit)
	offset := int(q.Offset)
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	users, err := h.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	dtos := make([]*queries.UserDTO, 0, len(users))
	for _, u := range users {
		dtos = append(dtos, &queries.UserDTO{
			ID:        string(u.ID()),
			Email:     u.Email(),
			Name:      u.Name(),
			CreatedAt: u.CreatedAt(),
			UpdatedAt: u.UpdatedAt(),
		})
	}
	return dtos, nil
}

func NewHandlers(repo domain.UserRepository) Handlers {
	if repo == nil {
		panic(errors.New("nil repo"))
	}
	return Handlers{
		Create: createUserHandler{repo: repo},
		Update: updateUserHandler{repo: repo},
		Delete: deleteUserHandler{repo: repo},
		Get:    getUserHandler{repo: repo},
		List:   listUsersHandler{repo: repo},
	}
}
