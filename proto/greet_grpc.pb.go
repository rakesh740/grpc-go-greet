// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: greet.proto

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

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Sum(ctx context.Context, in *Operand, opts ...grpc.CallOption) (*Response, error)
	GreetManyResponse(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (Calculator_GreetManyResponseClient, error)
	GetAllPrimes(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (Calculator_GetAllPrimesClient, error)
	GetAverage(ctx context.Context, opts ...grpc.CallOption) (Calculator_GetAverageClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Sum(ctx context.Context, in *Operand, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/greet.Calculator/sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) GreetManyResponse(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (Calculator_GreetManyResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], "/greet.Calculator/greetManyResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorGreetManyResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_GreetManyResponseClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type calculatorGreetManyResponseClient struct {
	grpc.ClientStream
}

func (x *calculatorGreetManyResponseClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) GetAllPrimes(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (Calculator_GetAllPrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], "/greet.Calculator/getAllPrimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorGetAllPrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_GetAllPrimesClient interface {
	Recv() (*NumberResponse, error)
	grpc.ClientStream
}

type calculatorGetAllPrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorGetAllPrimesClient) Recv() (*NumberResponse, error) {
	m := new(NumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) GetAverage(ctx context.Context, opts ...grpc.CallOption) (Calculator_GetAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[2], "/greet.Calculator/getAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorGetAverageClient{stream}
	return x, nil
}

type Calculator_GetAverageClient interface {
	Send(*NumberRequest) error
	CloseAndRecv() (*CalcResponse, error)
	grpc.ClientStream
}

type calculatorGetAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorGetAverageClient) Send(m *NumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorGetAverageClient) CloseAndRecv() (*CalcResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalcResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Sum(context.Context, *Operand) (*Response, error)
	GreetManyResponse(*GreetRequest, Calculator_GreetManyResponseServer) error
	GetAllPrimes(*NumberRequest, Calculator_GetAllPrimesServer) error
	GetAverage(Calculator_GetAverageServer) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Sum(context.Context, *Operand) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServer) GreetManyResponse(*GreetRequest, Calculator_GreetManyResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyResponse not implemented")
}
func (UnimplementedCalculatorServer) GetAllPrimes(*NumberRequest, Calculator_GetAllPrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllPrimes not implemented")
}
func (UnimplementedCalculatorServer) GetAverage(Calculator_GetAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAverage not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.Calculator/sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sum(ctx, req.(*Operand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_GreetManyResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).GreetManyResponse(m, &calculatorGreetManyResponseServer{stream})
}

type Calculator_GreetManyResponseServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type calculatorGreetManyResponseServer struct {
	grpc.ServerStream
}

func (x *calculatorGreetManyResponseServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_GetAllPrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).GetAllPrimes(m, &calculatorGetAllPrimesServer{stream})
}

type Calculator_GetAllPrimesServer interface {
	Send(*NumberResponse) error
	grpc.ServerStream
}

type calculatorGetAllPrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorGetAllPrimesServer) Send(m *NumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_GetAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).GetAverage(&calculatorGetAverageServer{stream})
}

type Calculator_GetAverageServer interface {
	SendAndClose(*CalcResponse) error
	Recv() (*NumberRequest, error)
	grpc.ServerStream
}

type calculatorGetAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorGetAverageServer) SendAndClose(m *CalcResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorGetAverageServer) Recv() (*NumberRequest, error) {
	m := new(NumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sum",
			Handler:    _Calculator_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "greetManyResponse",
			Handler:       _Calculator_GreetManyResponse_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getAllPrimes",
			Handler:       _Calculator_GetAllPrimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getAverage",
			Handler:       _Calculator_GetAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
