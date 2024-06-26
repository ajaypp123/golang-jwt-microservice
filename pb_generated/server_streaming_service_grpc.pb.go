// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: proto/server_streaming_service.proto

package pb_generated

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

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	CountNumbers(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (StreamingService_CountNumbersClient, error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) CountNumbers(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (StreamingService_CountNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], "/pb_generated.StreamingService/CountNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceCountNumbersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingService_CountNumbersClient interface {
	Recv() (*CountResponse, error)
	grpc.ClientStream
}

type streamingServiceCountNumbersClient struct {
	grpc.ClientStream
}

func (x *streamingServiceCountNumbersClient) Recv() (*CountResponse, error) {
	m := new(CountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility
type StreamingServiceServer interface {
	CountNumbers(*CountRequest, StreamingService_CountNumbersServer) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServiceServer struct {
}

func (UnimplementedStreamingServiceServer) CountNumbers(*CountRequest, StreamingService_CountNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method CountNumbers not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_CountNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).CountNumbers(m, &streamingServiceCountNumbersServer{stream})
}

type StreamingService_CountNumbersServer interface {
	Send(*CountResponse) error
	grpc.ServerStream
}

type streamingServiceCountNumbersServer struct {
	grpc.ServerStream
}

func (x *streamingServiceCountNumbersServer) Send(m *CountResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_generated.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CountNumbers",
			Handler:       _StreamingService_CountNumbers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/server_streaming_service.proto",
}
