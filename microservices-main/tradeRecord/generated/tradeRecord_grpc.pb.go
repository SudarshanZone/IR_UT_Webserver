// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: trdrec/tradeRecord.proto

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	GetCSV(ctx context.Context, in *CSVDataRequest, opts ...grpc.CallOption) (DataService_GetCSVClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) GetCSV(ctx context.Context, in *CSVDataRequest, opts ...grpc.CallOption) (DataService_GetCSVClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], "/trdrec.DataService/GetCSV", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceGetCSVClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_GetCSVClient interface {
	Recv() (*CSVDataResponse, error)
	grpc.ClientStream
}

type dataServiceGetCSVClient struct {
	grpc.ClientStream
}

func (x *dataServiceGetCSVClient) Recv() (*CSVDataResponse, error) {
	m := new(CSVDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	GetCSV(*CSVDataRequest, DataService_GetCSVServer) error
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) GetCSV(*CSVDataRequest, DataService_GetCSVServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCSV not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_GetCSV_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CSVDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).GetCSV(m, &dataServiceGetCSVServer{stream})
}

type DataService_GetCSVServer interface {
	Send(*CSVDataResponse) error
	grpc.ServerStream
}

type dataServiceGetCSVServer struct {
	grpc.ServerStream
}

func (x *dataServiceGetCSVServer) Send(m *CSVDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trdrec.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCSV",
			Handler:       _DataService_GetCSV_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "trdrec/tradeRecord.proto",
}
