// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: transport/v1/protocol.proto

package v1

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

// OpenSergoUniversalTransportServiceClient is the client API for OpenSergoUniversalTransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenSergoUniversalTransportServiceClient interface {
	SubscribeConfig(ctx context.Context, opts ...grpc.CallOption) (OpenSergoUniversalTransportService_SubscribeConfigClient, error)
}

type openSergoUniversalTransportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenSergoUniversalTransportServiceClient(cc grpc.ClientConnInterface) OpenSergoUniversalTransportServiceClient {
	return &openSergoUniversalTransportServiceClient{cc}
}

func (c *openSergoUniversalTransportServiceClient) SubscribeConfig(ctx context.Context, opts ...grpc.CallOption) (OpenSergoUniversalTransportService_SubscribeConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &OpenSergoUniversalTransportService_ServiceDesc.Streams[0], "/io.opensergo.transport.v1.OpenSergoUniversalTransportService/SubscribeConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &openSergoUniversalTransportServiceSubscribeConfigClient{stream}
	return x, nil
}

type OpenSergoUniversalTransportService_SubscribeConfigClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type openSergoUniversalTransportServiceSubscribeConfigClient struct {
	grpc.ClientStream
}

func (x *openSergoUniversalTransportServiceSubscribeConfigClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *openSergoUniversalTransportServiceSubscribeConfigClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenSergoUniversalTransportServiceServer is the server API for OpenSergoUniversalTransportService service.
// All implementations must embed UnimplementedOpenSergoUniversalTransportServiceServer
// for forward compatibility
type OpenSergoUniversalTransportServiceServer interface {
	SubscribeConfig(OpenSergoUniversalTransportService_SubscribeConfigServer) error
	mustEmbedUnimplementedOpenSergoUniversalTransportServiceServer()
}

// UnimplementedOpenSergoUniversalTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpenSergoUniversalTransportServiceServer struct {
}

func (UnimplementedOpenSergoUniversalTransportServiceServer) SubscribeConfig(OpenSergoUniversalTransportService_SubscribeConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeConfig not implemented")
}
func (UnimplementedOpenSergoUniversalTransportServiceServer) mustEmbedUnimplementedOpenSergoUniversalTransportServiceServer() {
}

// UnsafeOpenSergoUniversalTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenSergoUniversalTransportServiceServer will
// result in compilation errors.
type UnsafeOpenSergoUniversalTransportServiceServer interface {
	mustEmbedUnimplementedOpenSergoUniversalTransportServiceServer()
}

func RegisterOpenSergoUniversalTransportServiceServer(s grpc.ServiceRegistrar, srv OpenSergoUniversalTransportServiceServer) {
	s.RegisterService(&OpenSergoUniversalTransportService_ServiceDesc, srv)
}

func _OpenSergoUniversalTransportService_SubscribeConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OpenSergoUniversalTransportServiceServer).SubscribeConfig(&openSergoUniversalTransportServiceSubscribeConfigServer{stream})
}

type OpenSergoUniversalTransportService_SubscribeConfigServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type openSergoUniversalTransportServiceSubscribeConfigServer struct {
	grpc.ServerStream
}

func (x *openSergoUniversalTransportServiceSubscribeConfigServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *openSergoUniversalTransportServiceSubscribeConfigServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenSergoUniversalTransportService_ServiceDesc is the grpc.ServiceDesc for OpenSergoUniversalTransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenSergoUniversalTransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.opensergo.transport.v1.OpenSergoUniversalTransportService",
	HandlerType: (*OpenSergoUniversalTransportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeConfig",
			Handler:       _OpenSergoUniversalTransportService_SubscribeConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "transport/v1/protocol.proto",
}