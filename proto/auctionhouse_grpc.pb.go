//
//Generate-files:
//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.
//--go-grpc_opt=paths=source_relative auctionhouse.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: proto/auctionhouse.proto

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

const (
	AuctionhouseService_Result_FullMethodName            = "/proto.AuctionhouseService/Result"
	AuctionhouseService_Bid_FullMethodName               = "/proto.AuctionhouseService/Bid"
	AuctionhouseService_SendData_FullMethodName          = "/proto.AuctionhouseService/SendData"
	AuctionhouseService_SendDataToReplica_FullMethodName = "/proto.AuctionhouseService/SendDataToReplica"
)

// AuctionhouseServiceClient is the client API for AuctionhouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionhouseServiceClient interface {
	// result: send stream of qmgs bc client can ask at any point of time
	// what the highest bid in the round is.
	// prints winner when time runs out to all clients.
	Result(ctx context.Context, in *QueryResult, opts ...grpc.CallOption) (*ResponseToQuery, error)
	// First call to Bid registers the auctioners
	// Bidders can bid several times, but a bid must be higher than the previous
	// one(s)
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*StatusOfBid, error)
	// Send data to other server
	SendData(ctx context.Context, opts ...grpc.CallOption) (AuctionhouseService_SendDataClient, error)
	SendDataToReplica(ctx context.Context, in *GetDataRequestToReplica, opts ...grpc.CallOption) (*SendDataResponseToReplica, error)
}

type auctionhouseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionhouseServiceClient(cc grpc.ClientConnInterface) AuctionhouseServiceClient {
	return &auctionhouseServiceClient{cc}
}

func (c *auctionhouseServiceClient) Result(ctx context.Context, in *QueryResult, opts ...grpc.CallOption) (*ResponseToQuery, error) {
	out := new(ResponseToQuery)
	err := c.cc.Invoke(ctx, AuctionhouseService_Result_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionhouseServiceClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*StatusOfBid, error) {
	out := new(StatusOfBid)
	err := c.cc.Invoke(ctx, AuctionhouseService_Bid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionhouseServiceClient) SendData(ctx context.Context, opts ...grpc.CallOption) (AuctionhouseService_SendDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuctionhouseService_ServiceDesc.Streams[0], AuctionhouseService_SendData_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &auctionhouseServiceSendDataClient{stream}
	return x, nil
}

type AuctionhouseService_SendDataClient interface {
	Send(*GetDataRequest) error
	Recv() (*SendDataResponse, error)
	grpc.ClientStream
}

type auctionhouseServiceSendDataClient struct {
	grpc.ClientStream
}

func (x *auctionhouseServiceSendDataClient) Send(m *GetDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auctionhouseServiceSendDataClient) Recv() (*SendDataResponse, error) {
	m := new(SendDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *auctionhouseServiceClient) SendDataToReplica(ctx context.Context, in *GetDataRequestToReplica, opts ...grpc.CallOption) (*SendDataResponseToReplica, error) {
	out := new(SendDataResponseToReplica)
	err := c.cc.Invoke(ctx, AuctionhouseService_SendDataToReplica_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionhouseServiceServer is the server API for AuctionhouseService service.
// All implementations must embed UnimplementedAuctionhouseServiceServer
// for forward compatibility
type AuctionhouseServiceServer interface {
	// result: send stream of qmgs bc client can ask at any point of time
	// what the highest bid in the round is.
	// prints winner when time runs out to all clients.
	Result(context.Context, *QueryResult) (*ResponseToQuery, error)
	// First call to Bid registers the auctioners
	// Bidders can bid several times, but a bid must be higher than the previous
	// one(s)
	Bid(context.Context, *BidRequest) (*StatusOfBid, error)
	// Send data to other server
	SendData(AuctionhouseService_SendDataServer) error
	SendDataToReplica(context.Context, *GetDataRequestToReplica) (*SendDataResponseToReplica, error)
	mustEmbedUnimplementedAuctionhouseServiceServer()
}

// UnimplementedAuctionhouseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionhouseServiceServer struct {
}

func (UnimplementedAuctionhouseServiceServer) Result(context.Context, *QueryResult) (*ResponseToQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionhouseServiceServer) Bid(context.Context, *BidRequest) (*StatusOfBid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionhouseServiceServer) SendData(AuctionhouseService_SendDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedAuctionhouseServiceServer) SendDataToReplica(context.Context, *GetDataRequestToReplica) (*SendDataResponseToReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDataToReplica not implemented")
}
func (UnimplementedAuctionhouseServiceServer) mustEmbedUnimplementedAuctionhouseServiceServer() {}

// UnsafeAuctionhouseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionhouseServiceServer will
// result in compilation errors.
type UnsafeAuctionhouseServiceServer interface {
	mustEmbedUnimplementedAuctionhouseServiceServer()
}

func RegisterAuctionhouseServiceServer(s grpc.ServiceRegistrar, srv AuctionhouseServiceServer) {
	s.RegisterService(&AuctionhouseService_ServiceDesc, srv)
}

func _AuctionhouseService_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionhouseServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionhouseService_Result_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionhouseServiceServer).Result(ctx, req.(*QueryResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionhouseService_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionhouseServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionhouseService_Bid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionhouseServiceServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionhouseService_SendData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuctionhouseServiceServer).SendData(&auctionhouseServiceSendDataServer{stream})
}

type AuctionhouseService_SendDataServer interface {
	Send(*SendDataResponse) error
	Recv() (*GetDataRequest, error)
	grpc.ServerStream
}

type auctionhouseServiceSendDataServer struct {
	grpc.ServerStream
}

func (x *auctionhouseServiceSendDataServer) Send(m *SendDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auctionhouseServiceSendDataServer) Recv() (*GetDataRequest, error) {
	m := new(GetDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AuctionhouseService_SendDataToReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequestToReplica)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionhouseServiceServer).SendDataToReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionhouseService_SendDataToReplica_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionhouseServiceServer).SendDataToReplica(ctx, req.(*GetDataRequestToReplica))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionhouseService_ServiceDesc is the grpc.ServiceDesc for AuctionhouseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionhouseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuctionhouseService",
	HandlerType: (*AuctionhouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Result",
			Handler:    _AuctionhouseService_Result_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _AuctionhouseService_Bid_Handler,
		},
		{
			MethodName: "SendDataToReplica",
			Handler:    _AuctionhouseService_SendDataToReplica_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendData",
			Handler:       _AuctionhouseService_SendData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/auctionhouse.proto",
}
