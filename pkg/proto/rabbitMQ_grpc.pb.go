// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: rabbitMQ.proto

package proto

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

// SideCarClient is the client API for SideCar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SideCarClient interface {
	// Sends a greeting
	StartService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceReply, error)
}

type sideCarClient struct {
	cc grpc.ClientConnInterface
}

func NewSideCarClient(cc grpc.ClientConnInterface) SideCarClient {
	return &sideCarClient{cc}
}

func (c *sideCarClient) StartService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceReply, error) {
	out := new(ServiceReply)
	err := c.cc.Invoke(ctx, "/proto.SideCar/StartService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SideCarServer is the server API for SideCar service.
// All implementations must embed UnimplementedSideCarServer
// for forward compatibility
type SideCarServer interface {
	// Sends a greeting
	StartService(context.Context, *ServiceRequest) (*ServiceReply, error)
	mustEmbedUnimplementedSideCarServer()
}

// UnimplementedSideCarServer must be embedded to have forward compatible implementations.
type UnimplementedSideCarServer struct {
}

func (UnimplementedSideCarServer) StartService(context.Context, *ServiceRequest) (*ServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedSideCarServer) mustEmbedUnimplementedSideCarServer() {}

// UnsafeSideCarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SideCarServer will
// result in compilation errors.
type UnsafeSideCarServer interface {
	mustEmbedUnimplementedSideCarServer()
}

func RegisterSideCarServer(s grpc.ServiceRegistrar, srv SideCarServer) {
	s.RegisterService(&SideCar_ServiceDesc, srv)
}

func _SideCar_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SideCarServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SideCar/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SideCarServer).StartService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SideCar_ServiceDesc is the grpc.ServiceDesc for SideCar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SideCar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SideCar",
	HandlerType: (*SideCarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartService",
			Handler:    _SideCar_StartService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rabbitMQ.proto",
}
