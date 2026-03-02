package service

import (
	"context"
	"errors"

	"example.com/modmonolith/internal/modules/accounts/application/commands"
	"example.com/modmonolith/internal/modules/accounts/application/queries"
	"example.com/modmonolith/internal/modules/accounts/domain"
	userspublic "example.com/modmonolith/internal/modules/users/public"
)

type Handlers struct {
	Create commands.CreateAccountHandler
	Update commands.UpdateAccountHandler
	Delete commands.DeleteAccountHandler

	Get  queries.GetAccountHandler
	ListByUser queries.ListAccountsByUserHandler
}

type createAccountHandler struct {
	repo domain.AccountRepository
	users userspublic.UserReader // <-- boundary dependency (contract)
}

func (h createAccountHandler) Handle(ctx context.Context, cmd commands.CreateAccountCommand) (string, error) {
	exists, err := h.users.Exists(ctx, cmd.UserID)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", domain.ErrUserDoesNotExist
	}

	a, err := domain.CreateNewAccount(cmd.UserID, cmd.Label)
	if err != nil {
		return "", err
	}
	if err := h.repo.Create(ctx, a); err != nil {
		return "", err
	}
	return string(a.ID()), nil
}

type updateAccountHandler struct { repo domain.AccountRepository }

func (h updateAccountHandler) Handle(ctx context.Context, cmd commands.UpdateAccountCommand) error {
	a, err := h.repo.GetByID(ctx, domain.AccountID(cmd.ID))
	if err != nil {
		return err
	}
	if a == nil {
		return domain.ErrAccountNotFound
	}
	if cmd.Label != nil {
		if err := a.Relabel(*cmd.Label); err != nil {
			return err
		}
	}
	return h.repo.Update(ctx, a)
}

type deleteAccountHandler struct { repo domain.AccountRepository }

func (h deleteAccountHandler) Handle(ctx context.Context, cmd commands.DeleteAccountCommand) error {
	return h.repo.Delete(ctx, domain.AccountID(cmd.ID))
}

type getAccountHandler struct { repo domain.AccountRepository }

func (h getAccountHandler) Handle(ctx context.Context, q queries.GetAccountQuery) (*queries.AccountDTO, error) {
	a, err := h.repo.GetByID(ctx, domain.AccountID(q.ID))
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, domain.ErrAccountNotFound
	}
	return &queries.AccountDTO{
		ID: string(a.ID()),
		UserID: a.UserID(),
		Label: a.Label(),
		CreatedAt: a.CreatedAt(),
		UpdatedAt: a.UpdatedAt(),
	}, nil
}

type listAccountsByUserHandler struct { repo domain.AccountRepository }

func (h listAccountsByUserHandler) Handle(ctx context.Context, q queries.ListAccountsByUserQuery) ([]*queries.AccountDTO, error) {
	limit := int(q.Limit)
	offset := int(q.Offset)
	if limit <= 0 || limit > 200 { limit = 50 }
	if offset < 0 { offset = 0 }

	accs, err := h.repo.ListByUser(ctx, q.UserID, limit, offset)
	if err != nil {
		return nil, err
	}
	out := make([]*queries.AccountDTO, 0, len(accs))
	for _, a := range accs {
		out = append(out, &queries.AccountDTO{
			ID: string(a.ID()),
			UserID: a.UserID(),
			Label: a.Label(),
			CreatedAt: a.CreatedAt(),
			UpdatedAt: a.UpdatedAt(),
		})
	}
	return out, nil
}

func NewHandlers(repo domain.AccountRepository, users userspublic.UserReader) Handlers {
	if repo == nil { panic(errors.New("nil repo")) }
	if users == nil { panic(errors.New("nil users reader")) }
	return Handlers{
		Create: createAccountHandler{repo: repo, users: users},
		Update: updateAccountHandler{repo: repo},
		Delete: deleteAccountHandler{repo: repo},
		Get:    getAccountHandler{repo: repo},
		ListByUser: listAccountsByUserHandler{repo: repo},
	}
}
