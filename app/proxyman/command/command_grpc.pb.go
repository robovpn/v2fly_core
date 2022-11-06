package command

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

// HandlerServiceClient is the client API for HandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandlerServiceClient interface {
	AddInbound(ctx context.Context, in *AddInboundRequest, opts ...grpc.CallOption) (*AddInboundResponse, error)
	RemoveInbound(ctx context.Context, in *RemoveInboundRequest, opts ...grpc.CallOption) (*RemoveInboundResponse, error)
	AlterInbound(ctx context.Context, in *AlterInboundRequest, opts ...grpc.CallOption) (*AlterInboundResponse, error)
	AddOutbound(ctx context.Context, in *AddOutboundRequest, opts ...grpc.CallOption) (*AddOutboundResponse, error)
	RemoveOutbound(ctx context.Context, in *RemoveOutboundRequest, opts ...grpc.CallOption) (*RemoveOutboundResponse, error)
	AlterOutbound(ctx context.Context, in *AlterOutboundRequest, opts ...grpc.CallOption) (*AlterOutboundResponse, error)
}

type handlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandlerServiceClient(cc grpc.ClientConnInterface) HandlerServiceClient {
	return &handlerServiceClient{cc}
}

func (c *handlerServiceClient) AddInbound(ctx context.Context, in *AddInboundRequest, opts ...grpc.CallOption) (*AddInboundResponse, error) {
	out := new(AddInboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/AddInbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerServiceClient) RemoveInbound(ctx context.Context, in *RemoveInboundRequest, opts ...grpc.CallOption) (*RemoveInboundResponse, error) {
	out := new(RemoveInboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/RemoveInbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerServiceClient) AlterInbound(ctx context.Context, in *AlterInboundRequest, opts ...grpc.CallOption) (*AlterInboundResponse, error) {
	out := new(AlterInboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/AlterInbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerServiceClient) AddOutbound(ctx context.Context, in *AddOutboundRequest, opts ...grpc.CallOption) (*AddOutboundResponse, error) {
	out := new(AddOutboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/AddOutbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerServiceClient) RemoveOutbound(ctx context.Context, in *RemoveOutboundRequest, opts ...grpc.CallOption) (*RemoveOutboundResponse, error) {
	out := new(RemoveOutboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/RemoveOutbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerServiceClient) AlterOutbound(ctx context.Context, in *AlterOutboundRequest, opts ...grpc.CallOption) (*AlterOutboundResponse, error) {
	out := new(AlterOutboundResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.proxyman.command.HandlerService/AlterOutbound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlerServiceServer is the server API for HandlerService service.
// All implementations must embed UnimplementedHandlerServiceServer
// for forward compatibility
type HandlerServiceServer interface {
	AddInbound(context.Context, *AddInboundRequest) (*AddInboundResponse, error)
	RemoveInbound(context.Context, *RemoveInboundRequest) (*RemoveInboundResponse, error)
	AlterInbound(context.Context, *AlterInboundRequest) (*AlterInboundResponse, error)
	AddOutbound(context.Context, *AddOutboundRequest) (*AddOutboundResponse, error)
	RemoveOutbound(context.Context, *RemoveOutboundRequest) (*RemoveOutboundResponse, error)
	AlterOutbound(context.Context, *AlterOutboundRequest) (*AlterOutboundResponse, error)
	mustEmbedUnimplementedHandlerServiceServer()
}

// UnimplementedHandlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHandlerServiceServer struct {
}

func (UnimplementedHandlerServiceServer) AddInbound(context.Context, *AddInboundRequest) (*AddInboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInbound not implemented")
}
func (UnimplementedHandlerServiceServer) RemoveInbound(context.Context, *RemoveInboundRequest) (*RemoveInboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInbound not implemented")
}
func (UnimplementedHandlerServiceServer) AlterInbound(context.Context, *AlterInboundRequest) (*AlterInboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterInbound not implemented")
}
func (UnimplementedHandlerServiceServer) AddOutbound(context.Context, *AddOutboundRequest) (*AddOutboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOutbound not implemented")
}
func (UnimplementedHandlerServiceServer) RemoveOutbound(context.Context, *RemoveOutboundRequest) (*RemoveOutboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOutbound not implemented")
}
func (UnimplementedHandlerServiceServer) AlterOutbound(context.Context, *AlterOutboundRequest) (*AlterOutboundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterOutbound not implemented")
}
func (UnimplementedHandlerServiceServer) mustEmbedUnimplementedHandlerServiceServer() {}

// UnsafeHandlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandlerServiceServer will
// result in compilation errors.
type UnsafeHandlerServiceServer interface {
	mustEmbedUnimplementedHandlerServiceServer()
}

func RegisterHandlerServiceServer(s grpc.ServiceRegistrar, srv HandlerServiceServer) {
	s.RegisterService(&HandlerService_ServiceDesc, srv)
}

func _HandlerService_AddInbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).AddInbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/AddInbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).AddInbound(ctx, req.(*AddInboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerService_RemoveInbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).RemoveInbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/RemoveInbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).RemoveInbound(ctx, req.(*RemoveInboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerService_AlterInbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterInboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).AlterInbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/AlterInbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).AlterInbound(ctx, req.(*AlterInboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerService_AddOutbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOutboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).AddOutbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/AddOutbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).AddOutbound(ctx, req.(*AddOutboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerService_RemoveOutbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveOutboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).RemoveOutbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/RemoveOutbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).RemoveOutbound(ctx, req.(*RemoveOutboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerService_AlterOutbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterOutboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServiceServer).AlterOutbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.proxyman.command.HandlerService/AlterOutbound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServiceServer).AlterOutbound(ctx, req.(*AlterOutboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HandlerService_ServiceDesc is the grpc.ServiceDesc for HandlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2fly.core.app.proxyman.command.HandlerService",
	HandlerType: (*HandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInbound",
			Handler:    _HandlerService_AddInbound_Handler,
		},
		{
			MethodName: "RemoveInbound",
			Handler:    _HandlerService_RemoveInbound_Handler,
		},
		{
			MethodName: "AlterInbound",
			Handler:    _HandlerService_AlterInbound_Handler,
		},
		{
			MethodName: "AddOutbound",
			Handler:    _HandlerService_AddOutbound_Handler,
		},
		{
			MethodName: "RemoveOutbound",
			Handler:    _HandlerService_RemoveOutbound_Handler,
		},
		{
			MethodName: "AlterOutbound",
			Handler:    _HandlerService_AlterOutbound_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/proxyman/command/command.proto",
}
