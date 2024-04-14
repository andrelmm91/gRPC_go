// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CalculatorServicesClient is the client API for CalculatorServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServicesClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type calculatorServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServicesClient(cc grpc.ClientConnInterface) CalculatorServicesClient {
	return &calculatorServicesClient{cc}
}

func (c *calculatorServicesClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorServices/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServicesServer is the server API for CalculatorServices service.
// All implementations must embed UnimplementedCalculatorServicesServer
// for forward compatibility
type CalculatorServicesServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	mustEmbedUnimplementedCalculatorServicesServer()
}

// UnimplementedCalculatorServicesServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServicesServer struct {
}

func (UnimplementedCalculatorServicesServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServicesServer) mustEmbedUnimplementedCalculatorServicesServer() {}

// UnsafeCalculatorServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServicesServer will
// result in compilation errors.
type UnsafeCalculatorServicesServer interface {
	mustEmbedUnimplementedCalculatorServicesServer()
}

func RegisterCalculatorServicesServer(s grpc.ServiceRegistrar, srv CalculatorServicesServer) {
	s.RegisterService(&CalculatorServices_ServiceDesc, srv)
}

func _CalculatorServices_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServicesServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorServices/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServicesServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorServices_ServiceDesc is the grpc.ServiceDesc for CalculatorServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorServices",
	HandlerType: (*CalculatorServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorServices_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
