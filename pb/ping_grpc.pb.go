// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongClient interface {
	PingMessage(ctx context.Context, opts ...grpc.CallOption) (PingPong_PingMessageClient, error)
}

type pingPongClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongClient(cc grpc.ClientConnInterface) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) PingMessage(ctx context.Context, opts ...grpc.CallOption) (PingPong_PingMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingPong_ServiceDesc.Streams[0], "/pb.PingPong/PingMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingPongPingMessageClient{stream}
	return x, nil
}

type PingPong_PingMessageClient interface {
	Send(*PingPongRequest) error
	Recv() (*PingPongResponse, error)
	grpc.ClientStream
}

type pingPongPingMessageClient struct {
	grpc.ClientStream
}

func (x *pingPongPingMessageClient) Send(m *PingPongRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingPongPingMessageClient) Recv() (*PingPongResponse, error) {
	m := new(PingPongResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongServer is the server API for PingPong service.
// All implementations must embed UnimplementedPingPongServer
// for forward compatibility
type PingPongServer interface {
	PingMessage(PingPong_PingMessageServer) error
	mustEmbedUnimplementedPingPongServer()
}

// UnimplementedPingPongServer must be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (UnimplementedPingPongServer) PingMessage(PingPong_PingMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method PingMessage not implemented")
}
func (UnimplementedPingPongServer) mustEmbedUnimplementedPingPongServer() {}

// UnsafePingPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServer will
// result in compilation errors.
type UnsafePingPongServer interface {
	mustEmbedUnimplementedPingPongServer()
}

func RegisterPingPongServer(s grpc.ServiceRegistrar, srv PingPongServer) {
	s.RegisterService(&PingPong_ServiceDesc, srv)
}

func _PingPong_PingMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingPongServer).PingMessage(&pingPongPingMessageServer{stream})
}

type PingPong_PingMessageServer interface {
	Send(*PingPongResponse) error
	Recv() (*PingPongRequest, error)
	grpc.ServerStream
}

type pingPongPingMessageServer struct {
	grpc.ServerStream
}

func (x *pingPongPingMessageServer) Send(m *PingPongResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingPongPingMessageServer) Recv() (*PingPongRequest, error) {
	m := new(PingPongRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPong_ServiceDesc is the grpc.ServiceDesc for PingPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingMessage",
			Handler:       _PingPong_PingMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/proto/ping.proto",
}