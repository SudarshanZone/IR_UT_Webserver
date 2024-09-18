// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: fnosquoff/fnoSquareOff.proto

package __

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

// SquareOffServiceClient is the client API for SquareOffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SquareOffServiceClient interface {
	SquareOffOrder(ctx context.Context, in *SquareOffRequest, opts ...grpc.CallOption) (*SquareOffResponse, error)
}

type squareOffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSquareOffServiceClient(cc grpc.ClientConnInterface) SquareOffServiceClient {
	return &squareOffServiceClient{cc}
}

func (c *squareOffServiceClient) SquareOffOrder(ctx context.Context, in *SquareOffRequest, opts ...grpc.CallOption) (*SquareOffResponse, error) {
	out := new(SquareOffResponse)
	err := c.cc.Invoke(ctx, "/fnosquoff.SquareOffService/SquareOffOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareOffServiceServer is the server API for SquareOffService service.
// All implementations must embed UnimplementedSquareOffServiceServer
// for forward compatibility
type SquareOffServiceServer interface {
	SquareOffOrder(context.Context, *SquareOffRequest) (*SquareOffResponse, error)
	mustEmbedUnimplementedSquareOffServiceServer()
}

// UnimplementedSquareOffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSquareOffServiceServer struct {
}

func (UnimplementedSquareOffServiceServer) SquareOffOrder(context.Context, *SquareOffRequest) (*SquareOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareOffOrder not implemented")
}
func (UnimplementedSquareOffServiceServer) mustEmbedUnimplementedSquareOffServiceServer() {}

// UnsafeSquareOffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SquareOffServiceServer will
// result in compilation errors.
type UnsafeSquareOffServiceServer interface {
	mustEmbedUnimplementedSquareOffServiceServer()
}

func RegisterSquareOffServiceServer(s grpc.ServiceRegistrar, srv SquareOffServiceServer) {
	s.RegisterService(&SquareOffService_ServiceDesc, srv)
}

func _SquareOffService_SquareOffOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareOffServiceServer).SquareOffOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fnosquoff.SquareOffService/SquareOffOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareOffServiceServer).SquareOffOrder(ctx, req.(*SquareOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SquareOffService_ServiceDesc is the grpc.ServiceDesc for SquareOffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SquareOffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fnosquoff.SquareOffService",
	HandlerType: (*SquareOffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareOffOrder",
			Handler:    _SquareOffService_SquareOffOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fnosquoff/fnoSquareOff.proto",
}
