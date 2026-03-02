package accountsv1

import (
	"context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	AccountsService_CreateAccount_FullMethodName    = "/api.accounts.v1.AccountsService/CreateAccount"
	AccountsService_GetAccount_FullMethodName       = "/api.accounts.v1.AccountsService/GetAccount"
	AccountsService_ListAccountsByUser_FullMethodName = "/api.accounts.v1.AccountsService/ListAccountsByUser"
	AccountsService_UpdateAccount_FullMethodName    = "/api.accounts.v1.AccountsService/UpdateAccount"
	AccountsService_DeleteAccount_FullMethodName    = "/api.accounts.v1.AccountsService/DeleteAccount"
)

type AccountsServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	ListAccountsByUser(context.Context, *ListAccountsByUserRequest) (*ListAccountsByUserResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
}

type UnimplementedAccountsServiceServer struct{}

func (UnimplementedAccountsServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountsServiceServer) ListAccountsByUser(context.Context, *ListAccountsByUserRequest) (*ListAccountsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountsByUser not implemented")
}
func (UnimplementedAccountsServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountsServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

func RegisterAccountsServiceServer(s grpc.ServiceRegistrar, srv AccountsServiceServer) {
	s.RegisterService(&AccountsService_ServiceDesc, srv)
}

func _AccountsService_CreateAccount_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(AccountsServiceServer).CreateAccount(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: AccountsService_CreateAccount_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(AccountsServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest)) }
	return interceptor(ctx, in, info, h)
}
func _AccountsService_GetAccount_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(AccountsServiceServer).GetAccount(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: AccountsService_GetAccount_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(AccountsServiceServer).GetAccount(ctx, req.(*GetAccountRequest)) }
	return interceptor(ctx, in, info, h)
}
func _AccountsService_ListAccountsByUser_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ListAccountsByUserRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(AccountsServiceServer).ListAccountsByUser(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: AccountsService_ListAccountsByUser_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(AccountsServiceServer).ListAccountsByUser(ctx, req.(*ListAccountsByUserRequest)) }
	return interceptor(ctx, in, info, h)
}
func _AccountsService_UpdateAccount_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(AccountsServiceServer).UpdateAccount(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: AccountsService_UpdateAccount_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(AccountsServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest)) }
	return interceptor(ctx, in, info, h)
}
func _AccountsService_DeleteAccount_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil { return nil, err }
	if interceptor == nil { return srv.(AccountsServiceServer).DeleteAccount(ctx, in) }
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: AccountsService_DeleteAccount_FullMethodName}
	h := func(ctx context.Context, req any) (any, error) { return srv.(AccountsServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest)) }
	return interceptor(ctx, in, info, h)
}

var AccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.accounts.v1.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateAccount", Handler: _AccountsService_CreateAccount_Handler},
		{MethodName: "GetAccount", Handler: _AccountsService_GetAccount_Handler},
		{MethodName: "ListAccountsByUser", Handler: _AccountsService_ListAccountsByUser_Handler},
		{MethodName: "UpdateAccount", Handler: _AccountsService_UpdateAccount_Handler},
		{MethodName: "DeleteAccount", Handler: _AccountsService_DeleteAccount_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/v1/accounts.proto",
}
