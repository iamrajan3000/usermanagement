// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: users.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsersByIds(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUsersByIdsClient, error)
	SearchUsers(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (UserService_SearchUsersClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersByIds(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUsersByIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/grpc.UserService/GetUsersByIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUsersByIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUsersByIdsClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceGetUsersByIdsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUsersByIdsClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) SearchUsers(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (UserService_SearchUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/grpc.UserService/SearchUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceSearchUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_SearchUsersClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceSearchUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceSearchUsersClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserById(context.Context, *UserRequest) (*UserResponse, error)
	GetUsersByIds(*UsersRequest, UserService_GetUsersByIdsServer) error
	SearchUsers(*SearchRequest, UserService_SearchUsersServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUsersByIds(*UsersRequest, UserService_GetUsersByIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsersByIds not implemented")
}
func (UnimplementedUserServiceServer) SearchUsers(*SearchRequest, UserService_SearchUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersByIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUsersByIds(m, &userServiceGetUsersByIdsServer{stream})
}

type UserService_GetUsersByIdsServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userServiceGetUsersByIdsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUsersByIdsServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_SearchUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).SearchUsers(m, &userServiceSearchUsersServer{stream})
}

type UserService_SearchUsersServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userServiceSearchUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceSearchUsersServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsersByIds",
			Handler:       _UserService_GetUsersByIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchUsers",
			Handler:       _UserService_SearchUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users.proto",
}
