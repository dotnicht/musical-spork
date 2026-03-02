package usersv1

import (
	"context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	UsersService_CreateUser_FullMethodName = "/api.users.v1.UsersService/CreateUser"
	UsersService_GetUser_FullMethodName    = "/api.users.v1.UsersService/GetUser"
	UsersService_ListUsers_FullMethodName  = "/api.users.v1.UsersService/ListUsers"
	UsersService_UpdateUser_FullMethodName = "/api.users.v1.UsersService/UpdateUser"
	UsersService_DeleteUser_FullMethodName = "/api.users.v1.UsersService/DeleteUser"
)

type UsersServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

type UnimplementedUsersServiceServer struct{}

func (UnimplementedUsersServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsersServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUsersServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_CreateUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(UsersServiceServer).CreateUser(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: UsersService_CreateUser_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest)) }
	return interceptor(ctx, in, info, h)
}
func _UsersService_GetUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(UsersServiceServer).GetUser(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: UsersService_GetUser_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(UsersServiceServer).GetUser(ctx, req.(*GetUserRequest)) }
	return interceptor(ctx, in, info, h)
}
func _UsersService_ListUsers_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(UsersServiceServer).ListUsers(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: UsersService_ListUsers_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(UsersServiceServer).ListUsers(ctx, req.(*ListUsersRequest)) }
	return interceptor(ctx, in, info, h)
}
func _UsersService_UpdateUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(UsersServiceServer).UpdateUser(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: UsersService_UpdateUser_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(UsersServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest)) }
	return interceptor(ctx, in, info, h)
}
func _UsersService_DeleteUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(UsersServiceServer).DeleteUser(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: UsersService_DeleteUser_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(UsersServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest)) }
	return interceptor(ctx, in, info, h)
}

var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.users.v1.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateUser", Handler: _UsersService_CreateUser_Handler},
		{MethodName: "GetUser", Handler: _UsersService_GetUser_Handler},
		{MethodName: "ListUsers", Handler: _UsersService_ListUsers_Handler},
		{MethodName: "UpdateUser", Handler: _UsersService_UpdateUser_Handler},
		{MethodName: "DeleteUser", Handler: _UsersService_DeleteUser_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}
