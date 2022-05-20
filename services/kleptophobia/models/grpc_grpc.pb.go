// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/grpc.proto

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
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error)
	GetPublicInfo(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetPublicInfoRsp, error)
	GetEncryptedFullInfo(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetEncryptedFullInfoRsp, error)
	Ping(ctx context.Context, in *PingBody, opts ...grpc.CallOption) (*PingBody, error)
}

type kleptophobiaClient struct {
	cc grpc.ClientConnInterface
}

func NewKleptophobiaClient(cc grpc.ClientConnInterface) KleptophobiaClient {
	return &kleptophobiaClient{cc}
}

func (c *kleptophobiaClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kleptophobiaClient) GetPublicInfo(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetPublicInfoRsp, error) {
	out := new(GetPublicInfoRsp)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/GetPublicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kleptophobiaClient) GetEncryptedFullInfo(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetEncryptedFullInfoRsp, error) {
	out := new(GetEncryptedFullInfoRsp)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/GetEncryptedFullInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kleptophobiaClient) Ping(ctx context.Context, in *PingBody, opts ...grpc.CallOption) (*PingBody, error) {
	out := new(PingBody)
	err := c.cc.Invoke(ctx, "/models.Kleptophobia/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KleptophobiaServer is the server API for Kleptophobia service.
// All implementations must embed UnimplementedKleptophobiaServer
// for forward compatibility
type KleptophobiaServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRsp, error)
	GetPublicInfo(context.Context, *GetByUsernameReq) (*GetPublicInfoRsp, error)
	GetEncryptedFullInfo(context.Context, *GetByUsernameReq) (*GetEncryptedFullInfoRsp, error)
	Ping(context.Context, *PingBody) (*PingBody, error)
	mustEmbedUnimplementedKleptophobiaServer()
}

// UnimplementedKleptophobiaServer must be embedded to have forward compatible implementations.
type UnimplementedKleptophobiaServer struct {
}

func (UnimplementedKleptophobiaServer) Register(context.Context, *RegisterReq) (*RegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedKleptophobiaServer) GetPublicInfo(context.Context, *GetByUsernameReq) (*GetPublicInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicInfo not implemented")
}
func (UnimplementedKleptophobiaServer) GetEncryptedFullInfo(context.Context, *GetByUsernameReq) (*GetEncryptedFullInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncryptedFullInfo not implemented")
}
func (UnimplementedKleptophobiaServer) Ping(context.Context, *PingBody) (*PingBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
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
	in := new(RegisterReq)
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
		return srv.(KleptophobiaServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kleptophobia_GetPublicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameReq)
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
		return srv.(KleptophobiaServer).GetPublicInfo(ctx, req.(*GetByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kleptophobia_GetEncryptedFullInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KleptophobiaServer).GetEncryptedFullInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Kleptophobia/GetEncryptedFullInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KleptophobiaServer).GetEncryptedFullInfo(ctx, req.(*GetByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kleptophobia_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KleptophobiaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Kleptophobia/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KleptophobiaServer).Ping(ctx, req.(*PingBody))
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
		{
			MethodName: "GetEncryptedFullInfo",
			Handler:    _Kleptophobia_GetEncryptedFullInfo_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Kleptophobia_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc.proto",
}