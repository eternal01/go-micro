// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: basic.proto

package basic

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
	Basic_Ping_FullMethodName       = "/basic.Basic/Ping"
	Basic_GetRegion_FullMethodName  = "/basic.Basic/getRegion"
	Basic_GetRegions_FullMethodName = "/basic.Basic/getRegions"
	Basic_AddRegion_FullMethodName  = "/basic.Basic/addRegion"
)

// BasicClient is the client API for Basic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error)
	GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error)
	AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error)
}

type basicClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicClient(cc grpc.ClientConnInterface) BasicClient {
	return &basicClient{cc}
}

func (c *basicClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Basic_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error) {
	out := new(GetRegionResponse)
	err := c.cc.Invoke(ctx, Basic_GetRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error) {
	out := new(GetRegionsResponse)
	err := c.cc.Invoke(ctx, Basic_GetRegions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error) {
	out := new(AddRegionResponse)
	err := c.cc.Invoke(ctx, Basic_AddRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServer is the server API for Basic service.
// All implementations must embed UnimplementedBasicServer
// for forward compatibility
type BasicServer interface {
	Ping(context.Context, *Request) (*Response, error)
	GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error)
	GetRegions(context.Context, *GetRegionsRequest) (*GetRegionsResponse, error)
	AddRegion(context.Context, *AddRegionRequest) (*AddRegionResponse, error)
	mustEmbedUnimplementedBasicServer()
}

// UnimplementedBasicServer must be embedded to have forward compatible implementations.
type UnimplementedBasicServer struct {
}

func (UnimplementedBasicServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBasicServer) GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegion not implemented")
}
func (UnimplementedBasicServer) GetRegions(context.Context, *GetRegionsRequest) (*GetRegionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegions not implemented")
}
func (UnimplementedBasicServer) AddRegion(context.Context, *AddRegionRequest) (*AddRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRegion not implemented")
}
func (UnimplementedBasicServer) mustEmbedUnimplementedBasicServer() {}

// UnsafeBasicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicServer will
// result in compilation errors.
type UnsafeBasicServer interface {
	mustEmbedUnimplementedBasicServer()
}

func RegisterBasicServer(s grpc.ServiceRegistrar, srv BasicServer) {
	s.RegisterService(&Basic_ServiceDesc, srv)
}

func _Basic_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Basic_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Basic_GetRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_GetRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).GetRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Basic_GetRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).GetRegions(ctx, req.(*GetRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_AddRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).AddRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Basic_AddRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).AddRegion(ctx, req.(*AddRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Basic_ServiceDesc is the grpc.ServiceDesc for Basic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Basic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basic.Basic",
	HandlerType: (*BasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Basic_Ping_Handler,
		},
		{
			MethodName: "getRegion",
			Handler:    _Basic_GetRegion_Handler,
		},
		{
			MethodName: "getRegions",
			Handler:    _Basic_GetRegions_Handler,
		},
		{
			MethodName: "addRegion",
			Handler:    _Basic_AddRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic.proto",
}
