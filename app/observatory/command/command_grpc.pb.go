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

// ObservatoryServiceClient is the client API for ObservatoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObservatoryServiceClient interface {
	GetOutboundStatus(ctx context.Context, in *GetOutboundStatusRequest, opts ...grpc.CallOption) (*GetOutboundStatusResponse, error)
}

type observatoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObservatoryServiceClient(cc grpc.ClientConnInterface) ObservatoryServiceClient {
	return &observatoryServiceClient{cc}
}

func (c *observatoryServiceClient) GetOutboundStatus(ctx context.Context, in *GetOutboundStatusRequest, opts ...grpc.CallOption) (*GetOutboundStatusResponse, error) {
	out := new(GetOutboundStatusResponse)
	err := c.cc.Invoke(ctx, "/v2fly.core.app.observatory.command.ObservatoryService/GetOutboundStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObservatoryServiceServer is the server API for ObservatoryService service.
// All implementations must embed UnimplementedObservatoryServiceServer
// for forward compatibility
type ObservatoryServiceServer interface {
	GetOutboundStatus(context.Context, *GetOutboundStatusRequest) (*GetOutboundStatusResponse, error)
	mustEmbedUnimplementedObservatoryServiceServer()
}

// UnimplementedObservatoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObservatoryServiceServer struct {
}

func (UnimplementedObservatoryServiceServer) GetOutboundStatus(context.Context, *GetOutboundStatusRequest) (*GetOutboundStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutboundStatus not implemented")
}
func (UnimplementedObservatoryServiceServer) mustEmbedUnimplementedObservatoryServiceServer() {}

// UnsafeObservatoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObservatoryServiceServer will
// result in compilation errors.
type UnsafeObservatoryServiceServer interface {
	mustEmbedUnimplementedObservatoryServiceServer()
}

func RegisterObservatoryServiceServer(s grpc.ServiceRegistrar, srv ObservatoryServiceServer) {
	s.RegisterService(&ObservatoryService_ServiceDesc, srv)
}

func _ObservatoryService_GetOutboundStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutboundStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservatoryServiceServer).GetOutboundStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2fly.core.app.observatory.command.ObservatoryService/GetOutboundStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservatoryServiceServer).GetOutboundStatus(ctx, req.(*GetOutboundStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObservatoryService_ServiceDesc is the grpc.ServiceDesc for ObservatoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObservatoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2fly.core.app.observatory.command.ObservatoryService",
	HandlerType: (*ObservatoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOutboundStatus",
			Handler:    _ObservatoryService_GetOutboundStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/observatory/command/command.proto",
}
