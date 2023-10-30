// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: languagemodeling/v1/languagemodeling.proto

package languagemodelingv1

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
	LanguageModelingService_Predict_FullMethodName = "/languagemodeling.v1.LanguageModelingService/Predict"
)

// LanguageModelingServiceClient is the client API for LanguageModelingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LanguageModelingServiceClient interface {
	Predict(ctx context.Context, in *LanguageModelingRequest, opts ...grpc.CallOption) (*LanguageModelingResponse, error)
}

type languageModelingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLanguageModelingServiceClient(cc grpc.ClientConnInterface) LanguageModelingServiceClient {
	return &languageModelingServiceClient{cc}
}

func (c *languageModelingServiceClient) Predict(ctx context.Context, in *LanguageModelingRequest, opts ...grpc.CallOption) (*LanguageModelingResponse, error) {
	out := new(LanguageModelingResponse)
	err := c.cc.Invoke(ctx, LanguageModelingService_Predict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageModelingServiceServer is the server API for LanguageModelingService service.
// All implementations must embed UnimplementedLanguageModelingServiceServer
// for forward compatibility
type LanguageModelingServiceServer interface {
	Predict(context.Context, *LanguageModelingRequest) (*LanguageModelingResponse, error)
	mustEmbedUnimplementedLanguageModelingServiceServer()
}

// UnimplementedLanguageModelingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLanguageModelingServiceServer struct {
}

func (UnimplementedLanguageModelingServiceServer) Predict(context.Context, *LanguageModelingRequest) (*LanguageModelingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedLanguageModelingServiceServer) mustEmbedUnimplementedLanguageModelingServiceServer() {
}

// UnsafeLanguageModelingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LanguageModelingServiceServer will
// result in compilation errors.
type UnsafeLanguageModelingServiceServer interface {
	mustEmbedUnimplementedLanguageModelingServiceServer()
}

func RegisterLanguageModelingServiceServer(s grpc.ServiceRegistrar, srv LanguageModelingServiceServer) {
	s.RegisterService(&LanguageModelingService_ServiceDesc, srv)
}

func _LanguageModelingService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageModelingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageModelingServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LanguageModelingService_Predict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageModelingServiceServer).Predict(ctx, req.(*LanguageModelingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LanguageModelingService_ServiceDesc is the grpc.ServiceDesc for LanguageModelingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LanguageModelingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "languagemodeling.v1.LanguageModelingService",
	HandlerType: (*LanguageModelingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _LanguageModelingService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "languagemodeling/v1/languagemodeling.proto",
}