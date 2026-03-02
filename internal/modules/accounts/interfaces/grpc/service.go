package grpc

import (
	"context"
	"errors"
	"time"

	accountsv1 "example.com/modmonolith/api/gen/accounts/v1"
	"example.com/modmonolith/internal/modules/accounts/application/commands"
	"example.com/modmonolith/internal/modules/accounts/application/queries"
	"example.com/modmonolith/internal/modules/accounts/application/service"
	"example.com/modmonolith/internal/modules/accounts/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	accountsv1.UnimplementedAccountsServiceServer
	h service.Handlers
}

func New(h service.Handlers) *Service { return &Service{h: h} }

func (s *Service) CreateAccount(ctx context.Context, req *accountsv1.CreateAccountRequest) (*accountsv1.CreateAccountResponse, error) {
	id, err := s.h.Create.Handle(ctx, commands.CreateAccountCommand{UserID: req.UserId, Label: req.Label})
	if err != nil { return nil, mapErr(err) }
	return &accountsv1.CreateAccountResponse{Id: id}, nil
}

func (s *Service) GetAccount(ctx context.Context, req *accountsv1.GetAccountRequest) (*accountsv1.GetAccountResponse, error) {
	dto, err := s.h.Get.Handle(ctx, queries.GetAccountQuery{ID: req.Id})
	if err != nil { return nil, mapErr(err) }
	return &accountsv1.GetAccountResponse{Account: toProto(dto)}, nil
}

func (s *Service) ListAccountsByUser(ctx context.Context, req *accountsv1.ListAccountsByUserRequest) (*accountsv1.ListAccountsByUserResponse, error) {
	dtos, err := s.h.ListByUser.Handle(ctx, queries.ListAccountsByUserQuery{UserID: req.UserId, Limit: req.Limit, Offset: req.Offset})
	if err != nil { return nil, mapErr(err) }
	out := make([]*accountsv1.Account, 0, len(dtos))
	for _, d := range dtos { out = append(out, toProto(d)) }
	return &accountsv1.ListAccountsByUserResponse{Accounts: out}, nil
}

func (s *Service) UpdateAccount(ctx context.Context, req *accountsv1.UpdateAccountRequest) (*accountsv1.UpdateAccountResponse, error) {
	cmd := commands.UpdateAccountCommand{ID: req.Id, Label: req.Label}
	if err := s.h.Update.Handle(ctx, cmd); err != nil { return nil, mapErr(err) }
	return &accountsv1.UpdateAccountResponse{}, nil
}

func (s *Service) DeleteAccount(ctx context.Context, req *accountsv1.DeleteAccountRequest) (*accountsv1.DeleteAccountResponse, error) {
	if err := s.h.Delete.Handle(ctx, commands.DeleteAccountCommand{ID: req.Id}); err != nil { return nil, mapErr(err) }
	return &accountsv1.DeleteAccountResponse{}, nil
}

func toProto(d *queries.AccountDTO) *accountsv1.Account {
	return &accountsv1.Account{
		Id: d.ID,
		UserId: d.UserID,
		Label: d.Label,
		CreatedAtRfc3339: d.CreatedAt.Format(time.RFC3339),
		UpdatedAtRfc3339: d.UpdatedAt.Format(time.RFC3339),
	}
}

func mapErr(err error) error {
	switch {
	case errors.Is(err, domain.ErrInvalidAccountLabel):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrUserDoesNotExist):
		return status.Error(codes.FailedPrecondition, err.Error())
	case errors.Is(err, domain.ErrAccountNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
