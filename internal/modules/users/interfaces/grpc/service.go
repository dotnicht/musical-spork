package grpc

import (
	"context"
	"errors"
	"time"

	usersv1 "example.com/modmonolith/api/gen/users/v1"
	"example.com/modmonolith/internal/modules/users/application/commands"
	"example.com/modmonolith/internal/modules/users/application/queries"
	"example.com/modmonolith/internal/modules/users/application/service"
	"example.com/modmonolith/internal/modules/users/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	usersv1.UnimplementedUsersServiceServer
	h service.Handlers
}

func New(h service.Handlers) *Service { return &Service{h: h} }

func (s *Service) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	id, err := s.h.Create.Handle(ctx, commands.CreateUserCommand{Email: req.Email, Name: req.Name})
	if err != nil {
		return nil, mapErr(err)
	}
	return &usersv1.CreateUserResponse{Id: id}, nil
}

func (s *Service) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	dto, err := s.h.Get.Handle(ctx, queries.GetUserQuery{ID: req.Id})
	if err != nil {
		return nil, mapErr(err)
	}
	return &usersv1.GetUserResponse{User: toProto(dto)}, nil
}

func (s *Service) ListUsers(ctx context.Context, req *usersv1.ListUsersRequest) (*usersv1.ListUsersResponse, error) {
	dtos, err := s.h.List.Handle(ctx, queries.ListUsersQuery{Limit: req.Limit, Offset: req.Offset})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]*usersv1.User, 0, len(dtos))
	for _, d := range dtos {
		out = append(out, toProto(d))
	}
	return &usersv1.ListUsersResponse{Users: out}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *usersv1.UpdateUserRequest) (*usersv1.UpdateUserResponse, error) {
	cmd := commands.UpdateUserCommand{ID: req.Id, Email: req.Email, Name: req.Name}
	if err := s.h.Update.Handle(ctx, cmd); err != nil {
		return nil, mapErr(err)
	}
	return &usersv1.UpdateUserResponse{}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	if err := s.h.Delete.Handle(ctx, commands.DeleteUserCommand{ID: req.Id}); err != nil {
		return nil, mapErr(err)
	}
	return &usersv1.DeleteUserResponse{}, nil
}

func toProto(d *queries.UserDTO) *usersv1.User {
	return &usersv1.User{
		Id: d.ID,
		Email: d.Email,
		Name: d.Name,
		CreatedAtRfc3339: d.CreatedAt.Format(time.RFC3339),
		UpdatedAtRfc3339: d.UpdatedAt.Format(time.RFC3339),
	}
}

func mapErr(err error) error {
	switch {
	case errors.Is(err, domain.ErrInvalidEmail), errors.Is(err, domain.ErrInvalidUserName):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrEmailAlreadyUsed):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domain.ErrUserNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
