// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: user_core.proto

package user_core

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

const (
	UserCoreService_UserLogin_FullMethodName    = "/jlu_cow_studio.user_core.UserCoreService/UserLogin"
	UserCoreService_UserRegister_FullMethodName = "/jlu_cow_studio.user_core.UserCoreService/UserRegister"
	UserCoreService_UserInfo_FullMethodName     = "/jlu_cow_studio.user_core.UserCoreService/UserInfo"
	UserCoreService_UserAuth_FullMethodName     = "/jlu_cow_studio.user_core.UserCoreService/UserAuth"
)

// UserCoreServiceClient is the client API for UserCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCoreServiceClient interface {
	// 用户登录方法
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
	// 用户注册方法
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRes, error)
	// 获取用户信息方法
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error)
	// 用户权限认证方法
	UserAuth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthRes, error)
}

type userCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCoreServiceClient(cc grpc.ClientConnInterface) UserCoreServiceClient {
	return &userCoreServiceClient{cc}
}

func (c *userCoreServiceClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	out := new(UserLoginRes)
	err := c.cc.Invoke(ctx, UserCoreService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCoreServiceClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRes, error) {
	out := new(UserRegisterRes)
	err := c.cc.Invoke(ctx, UserCoreService_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCoreServiceClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, UserCoreService_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCoreServiceClient) UserAuth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*UserAuthRes, error) {
	out := new(UserAuthRes)
	err := c.cc.Invoke(ctx, UserCoreService_UserAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCoreServiceServer is the server API for UserCoreService service.
// All implementations must embed UnimplementedUserCoreServiceServer
// for forward compatibility
type UserCoreServiceServer interface {
	// 用户登录方法
	UserLogin(context.Context, *UserLoginReq) (*UserLoginRes, error)
	// 用户注册方法
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterRes, error)
	// 获取用户信息方法
	UserInfo(context.Context, *UserInfoReq) (*UserInfoRes, error)
	// 用户权限认证方法
	UserAuth(context.Context, *UserAuthReq) (*UserAuthRes, error)
	mustEmbedUnimplementedUserCoreServiceServer()
}

// UnimplementedUserCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserCoreServiceServer struct {
}

func (UnimplementedUserCoreServiceServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserCoreServiceServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserCoreServiceServer) UserInfo(context.Context, *UserInfoReq) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserCoreServiceServer) UserAuth(context.Context, *UserAuthReq) (*UserAuthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAuth not implemented")
}
func (UnimplementedUserCoreServiceServer) mustEmbedUnimplementedUserCoreServiceServer() {}

// UnsafeUserCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCoreServiceServer will
// result in compilation errors.
type UnsafeUserCoreServiceServer interface {
	mustEmbedUnimplementedUserCoreServiceServer()
}

func RegisterUserCoreServiceServer(s grpc.ServiceRegistrar, srv UserCoreServiceServer) {
	s.RegisterService(&UserCoreService_ServiceDesc, srv)
}

func _UserCoreService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCoreServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCoreService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCoreServiceServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCoreService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCoreServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCoreService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCoreServiceServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCoreService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCoreServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCoreService_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCoreServiceServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCoreService_UserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCoreServiceServer).UserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCoreService_UserAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCoreServiceServer).UserAuth(ctx, req.(*UserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCoreService_ServiceDesc is the grpc.ServiceDesc for UserCoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jlu_cow_studio.user_core.UserCoreService",
	HandlerType: (*UserCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserCoreService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserCoreService_UserRegister_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserCoreService_UserInfo_Handler,
		},
		{
			MethodName: "UserAuth",
			Handler:    _UserCoreService_UserAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_core.proto",
}