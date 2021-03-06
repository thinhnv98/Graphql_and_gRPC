// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// BookServicesClient is the client API for BookServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServicesClient interface {
	CreateBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error)
}

type bookServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServicesClient(cc grpc.ClientConnInterface) BookServicesClient {
	return &bookServicesClient{cc}
}

func (c *bookServicesClient) CreateBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error) {
	out := new(NewBookResponse)
	err := c.cc.Invoke(ctx, "/common.BookServices/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServicesServer is the server API for BookServices service.
// All implementations must embed UnimplementedBookServicesServer
// for forward compatibility
type BookServicesServer interface {
	CreateBook(context.Context, *NewBookRequest) (*NewBookResponse, error)
	mustEmbedUnimplementedBookServicesServer()
}

// UnimplementedBookServicesServer must be embedded to have forward compatible implementations.
type UnimplementedBookServicesServer struct {
}

func (UnimplementedBookServicesServer) CreateBook(context.Context, *NewBookRequest) (*NewBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServicesServer) mustEmbedUnimplementedBookServicesServer() {}

// UnsafeBookServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServicesServer will
// result in compilation errors.
type UnsafeBookServicesServer interface {
	mustEmbedUnimplementedBookServicesServer()
}

func RegisterBookServicesServer(s grpc.ServiceRegistrar, srv BookServicesServer) {
	s.RegisterService(&BookServices_ServiceDesc, srv)
}

func _BookServices_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServicesServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.BookServices/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServicesServer).CreateBook(ctx, req.(*NewBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookServices_ServiceDesc is the grpc.ServiceDesc for BookServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.BookServices",
	HandlerType: (*BookServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookServices_CreateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/common.proto",
}
