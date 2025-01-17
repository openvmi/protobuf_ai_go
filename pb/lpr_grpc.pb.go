// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: lpr.proto

package pb

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

// LprInferenceServiceClient is the client API for LprInferenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LprInferenceServiceClient interface {
	Infer(ctx context.Context, in *LprInferRequest, opts ...grpc.CallOption) (*LprInferResponse, error)
}

type lprInferenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLprInferenceServiceClient(cc grpc.ClientConnInterface) LprInferenceServiceClient {
	return &lprInferenceServiceClient{cc}
}

func (c *lprInferenceServiceClient) Infer(ctx context.Context, in *LprInferRequest, opts ...grpc.CallOption) (*LprInferResponse, error) {
	out := new(LprInferResponse)
	err := c.cc.Invoke(ctx, "/LprInferenceService/infer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LprInferenceServiceServer is the server API for LprInferenceService service.
// All implementations must embed UnimplementedLprInferenceServiceServer
// for forward compatibility
type LprInferenceServiceServer interface {
	Infer(context.Context, *LprInferRequest) (*LprInferResponse, error)
	mustEmbedUnimplementedLprInferenceServiceServer()
}

// UnimplementedLprInferenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLprInferenceServiceServer struct {
}

func (UnimplementedLprInferenceServiceServer) Infer(context.Context, *LprInferRequest) (*LprInferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Infer not implemented")
}
func (UnimplementedLprInferenceServiceServer) mustEmbedUnimplementedLprInferenceServiceServer() {}

// UnsafeLprInferenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LprInferenceServiceServer will
// result in compilation errors.
type UnsafeLprInferenceServiceServer interface {
	mustEmbedUnimplementedLprInferenceServiceServer()
}

func RegisterLprInferenceServiceServer(s grpc.ServiceRegistrar, srv LprInferenceServiceServer) {
	s.RegisterService(&LprInferenceService_ServiceDesc, srv)
}

func _LprInferenceService_Infer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LprInferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LprInferenceServiceServer).Infer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LprInferenceService/infer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LprInferenceServiceServer).Infer(ctx, req.(*LprInferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LprInferenceService_ServiceDesc is the grpc.ServiceDesc for LprInferenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LprInferenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LprInferenceService",
	HandlerType: (*LprInferenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "infer",
			Handler:    _LprInferenceService_Infer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lpr.proto",
}
