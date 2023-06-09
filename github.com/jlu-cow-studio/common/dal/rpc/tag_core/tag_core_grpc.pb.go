// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: tag_core.proto

package tag_core

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

// TagCoreServiceClient is the client API for TagCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagCoreServiceClient interface {
	GetTagListByScene(ctx context.Context, in *GetTagListBySceneRequest, opts ...grpc.CallOption) (*GetTagListBySceneResponse, error)
	GetTagListByItem(ctx context.Context, in *GetTagListByItemRequest, opts ...grpc.CallOption) (*GetTagListByItemResponse, error)
	GetTagListByUser(ctx context.Context, in *GetTagListByUserRequest, opts ...grpc.CallOption) (*GetTagListByUserResponse, error)
	UpdateItemTags(ctx context.Context, in *UpdateItemTagsRequest, opts ...grpc.CallOption) (*UpdateItemTagsResponse, error)
	UpdateUserTags(ctx context.Context, in *UpdateUserTagsRequest, opts ...grpc.CallOption) (*UpdateUserTagsResponse, error)
}

type tagCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagCoreServiceClient(cc grpc.ClientConnInterface) TagCoreServiceClient {
	return &tagCoreServiceClient{cc}
}

func (c *tagCoreServiceClient) GetTagListByScene(ctx context.Context, in *GetTagListBySceneRequest, opts ...grpc.CallOption) (*GetTagListBySceneResponse, error) {
	out := new(GetTagListBySceneResponse)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByScene", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagCoreServiceClient) GetTagListByItem(ctx context.Context, in *GetTagListByItemRequest, opts ...grpc.CallOption) (*GetTagListByItemResponse, error) {
	out := new(GetTagListByItemResponse)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagCoreServiceClient) GetTagListByUser(ctx context.Context, in *GetTagListByUserRequest, opts ...grpc.CallOption) (*GetTagListByUserResponse, error) {
	out := new(GetTagListByUserResponse)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagCoreServiceClient) UpdateItemTags(ctx context.Context, in *UpdateItemTagsRequest, opts ...grpc.CallOption) (*UpdateItemTagsResponse, error) {
	out := new(UpdateItemTagsResponse)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.tag_core.TagCoreService/UpdateItemTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagCoreServiceClient) UpdateUserTags(ctx context.Context, in *UpdateUserTagsRequest, opts ...grpc.CallOption) (*UpdateUserTagsResponse, error) {
	out := new(UpdateUserTagsResponse)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.tag_core.TagCoreService/UpdateUserTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagCoreServiceServer is the server API for TagCoreService service.
// All implementations must embed UnimplementedTagCoreServiceServer
// for forward compatibility
type TagCoreServiceServer interface {
	GetTagListByScene(context.Context, *GetTagListBySceneRequest) (*GetTagListBySceneResponse, error)
	GetTagListByItem(context.Context, *GetTagListByItemRequest) (*GetTagListByItemResponse, error)
	GetTagListByUser(context.Context, *GetTagListByUserRequest) (*GetTagListByUserResponse, error)
	UpdateItemTags(context.Context, *UpdateItemTagsRequest) (*UpdateItemTagsResponse, error)
	UpdateUserTags(context.Context, *UpdateUserTagsRequest) (*UpdateUserTagsResponse, error)
	mustEmbedUnimplementedTagCoreServiceServer()
}

// UnimplementedTagCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagCoreServiceServer struct {
}

func (UnimplementedTagCoreServiceServer) GetTagListByScene(context.Context, *GetTagListBySceneRequest) (*GetTagListBySceneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagListByScene not implemented")
}
func (UnimplementedTagCoreServiceServer) GetTagListByItem(context.Context, *GetTagListByItemRequest) (*GetTagListByItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagListByItem not implemented")
}
func (UnimplementedTagCoreServiceServer) GetTagListByUser(context.Context, *GetTagListByUserRequest) (*GetTagListByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagListByUser not implemented")
}
func (UnimplementedTagCoreServiceServer) UpdateItemTags(context.Context, *UpdateItemTagsRequest) (*UpdateItemTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItemTags not implemented")
}
func (UnimplementedTagCoreServiceServer) UpdateUserTags(context.Context, *UpdateUserTagsRequest) (*UpdateUserTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserTags not implemented")
}
func (UnimplementedTagCoreServiceServer) mustEmbedUnimplementedTagCoreServiceServer() {}

// UnsafeTagCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagCoreServiceServer will
// result in compilation errors.
type UnsafeTagCoreServiceServer interface {
	mustEmbedUnimplementedTagCoreServiceServer()
}

func RegisterTagCoreServiceServer(s grpc.ServiceRegistrar, srv TagCoreServiceServer) {
	s.RegisterService(&TagCoreService_ServiceDesc, srv)
}

func _TagCoreService_GetTagListByScene_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListBySceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagCoreServiceServer).GetTagListByScene(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByScene",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagCoreServiceServer).GetTagListByScene(ctx, req.(*GetTagListBySceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagCoreService_GetTagListByItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListByItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagCoreServiceServer).GetTagListByItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagCoreServiceServer).GetTagListByItem(ctx, req.(*GetTagListByItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagCoreService_GetTagListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagCoreServiceServer).GetTagListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.tag_core.TagCoreService/GetTagListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagCoreServiceServer).GetTagListByUser(ctx, req.(*GetTagListByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagCoreService_UpdateItemTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagCoreServiceServer).UpdateItemTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.tag_core.TagCoreService/UpdateItemTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagCoreServiceServer).UpdateItemTags(ctx, req.(*UpdateItemTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagCoreService_UpdateUserTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagCoreServiceServer).UpdateUserTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.tag_core.TagCoreService/UpdateUserTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagCoreServiceServer).UpdateUserTags(ctx, req.(*UpdateUserTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagCoreService_ServiceDesc is the grpc.ServiceDesc for TagCoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagCoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jlu_cow_studio.tag_core.TagCoreService",
	HandlerType: (*TagCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagListByScene",
			Handler:    _TagCoreService_GetTagListByScene_Handler,
		},
		{
			MethodName: "GetTagListByItem",
			Handler:    _TagCoreService_GetTagListByItem_Handler,
		},
		{
			MethodName: "GetTagListByUser",
			Handler:    _TagCoreService_GetTagListByUser_Handler,
		},
		{
			MethodName: "UpdateItemTags",
			Handler:    _TagCoreService_UpdateItemTags_Handler,
		},
		{
			MethodName: "UpdateUserTags",
			Handler:    _TagCoreService_UpdateUserTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tag_core.proto",
}
