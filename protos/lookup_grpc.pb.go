// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: protos/lookup.proto

package protos

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
	LookupService_Lookup_FullMethodName = "/protos.LookupService/Lookup"
)

// LookupServiceClient is the client API for LookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LookupServiceClient interface {
	Lookup(ctx context.Context, in *LookupReq, opts ...grpc.CallOption) (*LookupRes, error)
}

type lookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupServiceClient(cc grpc.ClientConnInterface) LookupServiceClient {
	return &lookupServiceClient{cc}
}

func (c *lookupServiceClient) Lookup(ctx context.Context, in *LookupReq, opts ...grpc.CallOption) (*LookupRes, error) {
	out := new(LookupRes)
	err := c.cc.Invoke(ctx, LookupService_Lookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LookupServiceServer is the server API for LookupService service.
// All implementations must embed UnimplementedLookupServiceServer
// for forward compatibility
type LookupServiceServer interface {
	Lookup(context.Context, *LookupReq) (*LookupRes, error)
	mustEmbedUnimplementedLookupServiceServer()
}

// UnimplementedLookupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLookupServiceServer struct {
}

func (UnimplementedLookupServiceServer) Lookup(context.Context, *LookupReq) (*LookupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedLookupServiceServer) mustEmbedUnimplementedLookupServiceServer() {}

// UnsafeLookupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LookupServiceServer will
// result in compilation errors.
type UnsafeLookupServiceServer interface {
	mustEmbedUnimplementedLookupServiceServer()
}

func RegisterLookupServiceServer(s grpc.ServiceRegistrar, srv LookupServiceServer) {
	s.RegisterService(&LookupService_ServiceDesc, srv)
}

func _LookupService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LookupService_Lookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).Lookup(ctx, req.(*LookupReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LookupService_ServiceDesc is the grpc.ServiceDesc for LookupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LookupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.LookupService",
	HandlerType: (*LookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _LookupService_Lookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/lookup.proto",
}
