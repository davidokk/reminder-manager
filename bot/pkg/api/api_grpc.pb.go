// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api.proto

package api

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

// InterfaceClient is the client API for Interface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterfaceClient interface {
	ReminderList(ctx context.Context, in *ReminderListRequest, opts ...grpc.CallOption) (*ReminderListResponse, error)
	ReminderGet(ctx context.Context, in *ReminderGetRequest, opts ...grpc.CallOption) (*ReminderGetResponse, error)
	ReminderCreate(ctx context.Context, in *ReminderCreateRequest, opts ...grpc.CallOption) (*ReminderCreateResponse, error)
	ReminderUpdate(ctx context.Context, in *ReminderUpdateRequest, opts ...grpc.CallOption) (*ReminderUpdateResponse, error)
	ReminderRemove(ctx context.Context, in *ReminderRemoveRequest, opts ...grpc.CallOption) (*ReminderRemoveResponse, error)
}

type interfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterfaceClient(cc grpc.ClientConnInterface) InterfaceClient {
	return &interfaceClient{cc}
}

func (c *interfaceClient) ReminderList(ctx context.Context, in *ReminderListRequest, opts ...grpc.CallOption) (*ReminderListResponse, error) {
	out := new(ReminderListResponse)
	err := c.cc.Invoke(ctx, "/Interface/ReminderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ReminderGet(ctx context.Context, in *ReminderGetRequest, opts ...grpc.CallOption) (*ReminderGetResponse, error) {
	out := new(ReminderGetResponse)
	err := c.cc.Invoke(ctx, "/Interface/ReminderGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ReminderCreate(ctx context.Context, in *ReminderCreateRequest, opts ...grpc.CallOption) (*ReminderCreateResponse, error) {
	out := new(ReminderCreateResponse)
	err := c.cc.Invoke(ctx, "/Interface/ReminderCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ReminderUpdate(ctx context.Context, in *ReminderUpdateRequest, opts ...grpc.CallOption) (*ReminderUpdateResponse, error) {
	out := new(ReminderUpdateResponse)
	err := c.cc.Invoke(ctx, "/Interface/ReminderUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ReminderRemove(ctx context.Context, in *ReminderRemoveRequest, opts ...grpc.CallOption) (*ReminderRemoveResponse, error) {
	out := new(ReminderRemoveResponse)
	err := c.cc.Invoke(ctx, "/Interface/ReminderRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfaceServer is the server API for Interface service.
// All implementations must embed UnimplementedInterfaceServer
// for forward compatibility
type InterfaceServer interface {
	ReminderList(context.Context, *ReminderListRequest) (*ReminderListResponse, error)
	ReminderGet(context.Context, *ReminderGetRequest) (*ReminderGetResponse, error)
	ReminderCreate(context.Context, *ReminderCreateRequest) (*ReminderCreateResponse, error)
	ReminderUpdate(context.Context, *ReminderUpdateRequest) (*ReminderUpdateResponse, error)
	ReminderRemove(context.Context, *ReminderRemoveRequest) (*ReminderRemoveResponse, error)
	mustEmbedUnimplementedInterfaceServer()
}

// UnimplementedInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedInterfaceServer struct {
}

func (UnimplementedInterfaceServer) ReminderList(context.Context, *ReminderListRequest) (*ReminderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReminderList not implemented")
}
func (UnimplementedInterfaceServer) ReminderGet(context.Context, *ReminderGetRequest) (*ReminderGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReminderGet not implemented")
}
func (UnimplementedInterfaceServer) ReminderCreate(context.Context, *ReminderCreateRequest) (*ReminderCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReminderCreate not implemented")
}
func (UnimplementedInterfaceServer) ReminderUpdate(context.Context, *ReminderUpdateRequest) (*ReminderUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReminderUpdate not implemented")
}
func (UnimplementedInterfaceServer) ReminderRemove(context.Context, *ReminderRemoveRequest) (*ReminderRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReminderRemove not implemented")
}
func (UnimplementedInterfaceServer) mustEmbedUnimplementedInterfaceServer() {}

// UnsafeInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterfaceServer will
// result in compilation errors.
type UnsafeInterfaceServer interface {
	mustEmbedUnimplementedInterfaceServer()
}

func RegisterInterfaceServer(s grpc.ServiceRegistrar, srv InterfaceServer) {
	s.RegisterService(&Interface_ServiceDesc, srv)
}

func _Interface_ReminderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReminderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ReminderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interface/ReminderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ReminderList(ctx, req.(*ReminderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ReminderGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReminderGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ReminderGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interface/ReminderGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ReminderGet(ctx, req.(*ReminderGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ReminderCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReminderCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ReminderCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interface/ReminderCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ReminderCreate(ctx, req.(*ReminderCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ReminderUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReminderUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ReminderUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interface/ReminderUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ReminderUpdate(ctx, req.(*ReminderUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ReminderRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReminderRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ReminderRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interface/ReminderRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ReminderRemove(ctx, req.(*ReminderRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Interface_ServiceDesc is the grpc.ServiceDesc for Interface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Interface",
	HandlerType: (*InterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReminderList",
			Handler:    _Interface_ReminderList_Handler,
		},
		{
			MethodName: "ReminderGet",
			Handler:    _Interface_ReminderGet_Handler,
		},
		{
			MethodName: "ReminderCreate",
			Handler:    _Interface_ReminderCreate_Handler,
		},
		{
			MethodName: "ReminderUpdate",
			Handler:    _Interface_ReminderUpdate_Handler,
		},
		{
			MethodName: "ReminderRemove",
			Handler:    _Interface_ReminderRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
