// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: models/models.proto

package models

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

// KleptophobiaClient is the client API for Kleptophobia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KleptophobiaClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	GetPublicInfo(ctx context.Context, in *GetPublicInfoRequest, opts ...grpc.CallOption) (*GetPublicInfoReply, error)
}

type kleptophobiaClient struct {
	cc grpc.ClientConnInterface
}

func NewKleptophobiaClient(cc grpc.ClientConnInterface) KleptophobiaClient {
	return &kleptophobiaClient{cc}
}

func (c *kleptophobiaClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kleptophobiaClient) GetPublicInfo(ctx context.Context, in *GetPublicInfoRequest, opts ...grpc.CallOption) (*GetPublicInfoReply, error) {
	out := new(GetPublicInfoReply)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/GetPublicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KleptophobiaServer is the server API for Kleptophobia service.
// All implementations must embed UnimplementedKleptophobiaServer
// for forward compatibility
type KleptophobiaServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	GetPublicInfo(context.Context, *GetPublicInfoRequest) (*GetPublicInfoReply, error)
	mustEmbedUnimplementedKleptophobiaServer()
}

// UnimplementedKleptophobiaServer must be embedded to have forward compatible implementations.
type UnimplementedKleptophobiaServer struct {
}

func (UnimplementedKleptophobiaServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedKleptophobiaServer) GetPublicInfo(context.Context, *GetPublicInfoRequest) (*GetPublicInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicInfo not implemented")
}
func (UnimplementedKleptophobiaServer) mustEmbedUnimplementedKleptophobiaServer() {}

// UnsafeKleptophobiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KleptophobiaServer will
// result in compilation errors.
type UnsafeKleptophobiaServer interface {
	mustEmbedUnimplementedKleptophobiaServer()
}

func RegisterKleptophobiaServer(s grpc.ServiceRegistrar, srv KleptophobiaServer) {
	s.RegisterService(&Kleptophobia_ServiceDesc, srv)
}

func _Kleptophobia_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KleptophobiaServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Kleptophobia/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KleptophobiaServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kleptophobia_GetPublicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KleptophobiaServer).GetPublicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Kleptophobia/GetPublicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KleptophobiaServer).GetPublicInfo(ctx, req.(*GetPublicInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kleptophobia_ServiceDesc is the grpc.ServiceDesc for Kleptophobia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kleptophobia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.Kleptophobia",
	HandlerType: (*KleptophobiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Kleptophobia_Register_Handler,
		},
		{
			MethodName: "GetPublicInfo",
			Handler:    _Kleptophobia_GetPublicInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models/models.proto",
}
