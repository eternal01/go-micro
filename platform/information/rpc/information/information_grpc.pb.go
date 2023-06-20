// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: information.proto

package information

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
	InformationService_Ping_FullMethodName = "/information.InformationService/Ping"
)

// InformationServiceClient is the client API for InformationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformationServiceClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type informationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInformationServiceClient(cc grpc.ClientConnInterface) InformationServiceClient {
	return &informationServiceClient{cc}
}

func (c *informationServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, InformationService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformationServiceServer is the server API for InformationService service.
// All implementations must embed UnimplementedInformationServiceServer
// for forward compatibility
type InformationServiceServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedInformationServiceServer()
}

// UnimplementedInformationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInformationServiceServer struct {
}

func (UnimplementedInformationServiceServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInformationServiceServer) mustEmbedUnimplementedInformationServiceServer() {}

// UnsafeInformationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformationServiceServer will
// result in compilation errors.
type UnsafeInformationServiceServer interface {
	mustEmbedUnimplementedInformationServiceServer()
}

func RegisterInformationServiceServer(s grpc.ServiceRegistrar, srv InformationServiceServer) {
	s.RegisterService(&InformationService_ServiceDesc, srv)
}

func _InformationService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// InformationService_ServiceDesc is the grpc.ServiceDesc for InformationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "information.InformationService",
	HandlerType: (*InformationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _InformationService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "information.proto",
}

const (
	InformationTopicService_AddTopic_FullMethodName     = "/information.InformationTopicService/AddTopic"
	InformationTopicService_GetTopic_FullMethodName     = "/information.InformationTopicService/GetTopic"
	InformationTopicService_GetTopicList_FullMethodName = "/information.InformationTopicService/GetTopicList"
	InformationTopicService_UpdateTopic_FullMethodName  = "/information.InformationTopicService/UpdateTopic"
	InformationTopicService_DeleteTopic_FullMethodName  = "/information.InformationTopicService/DeleteTopic"
)

// InformationTopicServiceClient is the client API for InformationTopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformationTopicServiceClient interface {
	// 添加文章
	AddTopic(ctx context.Context, in *AddTopicRequest, opts ...grpc.CallOption) (*AddTopicResponse, error)
	// 获取文章
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error)
	// 获取文章列表
	GetTopicList(ctx context.Context, in *GetTopicsRequest, opts ...grpc.CallOption) (*GetTopicsResponse, error)
	// 编辑文章
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*UpdateTopicResponse, error)
	// 删除文章
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error)
}

type informationTopicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInformationTopicServiceClient(cc grpc.ClientConnInterface) InformationTopicServiceClient {
	return &informationTopicServiceClient{cc}
}

func (c *informationTopicServiceClient) AddTopic(ctx context.Context, in *AddTopicRequest, opts ...grpc.CallOption) (*AddTopicResponse, error) {
	out := new(AddTopicResponse)
	err := c.cc.Invoke(ctx, InformationTopicService_AddTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationTopicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error) {
	out := new(GetTopicResponse)
	err := c.cc.Invoke(ctx, InformationTopicService_GetTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationTopicServiceClient) GetTopicList(ctx context.Context, in *GetTopicsRequest, opts ...grpc.CallOption) (*GetTopicsResponse, error) {
	out := new(GetTopicsResponse)
	err := c.cc.Invoke(ctx, InformationTopicService_GetTopicList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationTopicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*UpdateTopicResponse, error) {
	out := new(UpdateTopicResponse)
	err := c.cc.Invoke(ctx, InformationTopicService_UpdateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationTopicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := c.cc.Invoke(ctx, InformationTopicService_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformationTopicServiceServer is the server API for InformationTopicService service.
// All implementations must embed UnimplementedInformationTopicServiceServer
// for forward compatibility
type InformationTopicServiceServer interface {
	// 添加文章
	AddTopic(context.Context, *AddTopicRequest) (*AddTopicResponse, error)
	// 获取文章
	GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error)
	// 获取文章列表
	GetTopicList(context.Context, *GetTopicsRequest) (*GetTopicsResponse, error)
	// 编辑文章
	UpdateTopic(context.Context, *UpdateTopicRequest) (*UpdateTopicResponse, error)
	// 删除文章
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error)
	mustEmbedUnimplementedInformationTopicServiceServer()
}

// UnimplementedInformationTopicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInformationTopicServiceServer struct {
}

func (UnimplementedInformationTopicServiceServer) AddTopic(context.Context, *AddTopicRequest) (*AddTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopic not implemented")
}
func (UnimplementedInformationTopicServiceServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedInformationTopicServiceServer) GetTopicList(context.Context, *GetTopicsRequest) (*GetTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicList not implemented")
}
func (UnimplementedInformationTopicServiceServer) UpdateTopic(context.Context, *UpdateTopicRequest) (*UpdateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (UnimplementedInformationTopicServiceServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedInformationTopicServiceServer) mustEmbedUnimplementedInformationTopicServiceServer() {
}

// UnsafeInformationTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformationTopicServiceServer will
// result in compilation errors.
type UnsafeInformationTopicServiceServer interface {
	mustEmbedUnimplementedInformationTopicServiceServer()
}

func RegisterInformationTopicServiceServer(s grpc.ServiceRegistrar, srv InformationTopicServiceServer) {
	s.RegisterService(&InformationTopicService_ServiceDesc, srv)
}

func _InformationTopicService_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicServiceServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicService_AddTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicServiceServer).AddTopic(ctx, req.(*AddTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformationTopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicService_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformationTopicService_GetTopicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicServiceServer).GetTopicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicService_GetTopicList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicServiceServer).GetTopicList(ctx, req.(*GetTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformationTopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicService_UpdateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformationTopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicService_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InformationTopicService_ServiceDesc is the grpc.ServiceDesc for InformationTopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformationTopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "information.InformationTopicService",
	HandlerType: (*InformationTopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTopic",
			Handler:    _InformationTopicService_AddTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _InformationTopicService_GetTopic_Handler,
		},
		{
			MethodName: "GetTopicList",
			Handler:    _InformationTopicService_GetTopicList_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _InformationTopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _InformationTopicService_DeleteTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "information.proto",
}

const (
	InformationTopicAuditRecordService_AddTopicAuditRecord_FullMethodName = "/information.InformationTopicAuditRecordService/AddTopicAuditRecord"
	InformationTopicAuditRecordService_GetTopicAuditRecord_FullMethodName = "/information.InformationTopicAuditRecordService/GetTopicAuditRecord"
)

// InformationTopicAuditRecordServiceClient is the client API for InformationTopicAuditRecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformationTopicAuditRecordServiceClient interface {
	// 添加文章审核记录
	AddTopicAuditRecord(ctx context.Context, in *AddTopicAuditRecordRequest, opts ...grpc.CallOption) (*AddTopicAuditRecordResponse, error)
	// 获取文章审核记录
	GetTopicAuditRecord(ctx context.Context, in *GetTopicAuditRecordRequest, opts ...grpc.CallOption) (*GetTopicAuditRecordResponse, error)
}

type informationTopicAuditRecordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInformationTopicAuditRecordServiceClient(cc grpc.ClientConnInterface) InformationTopicAuditRecordServiceClient {
	return &informationTopicAuditRecordServiceClient{cc}
}

func (c *informationTopicAuditRecordServiceClient) AddTopicAuditRecord(ctx context.Context, in *AddTopicAuditRecordRequest, opts ...grpc.CallOption) (*AddTopicAuditRecordResponse, error) {
	out := new(AddTopicAuditRecordResponse)
	err := c.cc.Invoke(ctx, InformationTopicAuditRecordService_AddTopicAuditRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationTopicAuditRecordServiceClient) GetTopicAuditRecord(ctx context.Context, in *GetTopicAuditRecordRequest, opts ...grpc.CallOption) (*GetTopicAuditRecordResponse, error) {
	out := new(GetTopicAuditRecordResponse)
	err := c.cc.Invoke(ctx, InformationTopicAuditRecordService_GetTopicAuditRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformationTopicAuditRecordServiceServer is the server API for InformationTopicAuditRecordService service.
// All implementations must embed UnimplementedInformationTopicAuditRecordServiceServer
// for forward compatibility
type InformationTopicAuditRecordServiceServer interface {
	// 添加文章审核记录
	AddTopicAuditRecord(context.Context, *AddTopicAuditRecordRequest) (*AddTopicAuditRecordResponse, error)
	// 获取文章审核记录
	GetTopicAuditRecord(context.Context, *GetTopicAuditRecordRequest) (*GetTopicAuditRecordResponse, error)
	mustEmbedUnimplementedInformationTopicAuditRecordServiceServer()
}

// UnimplementedInformationTopicAuditRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInformationTopicAuditRecordServiceServer struct {
}

func (UnimplementedInformationTopicAuditRecordServiceServer) AddTopicAuditRecord(context.Context, *AddTopicAuditRecordRequest) (*AddTopicAuditRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopicAuditRecord not implemented")
}
func (UnimplementedInformationTopicAuditRecordServiceServer) GetTopicAuditRecord(context.Context, *GetTopicAuditRecordRequest) (*GetTopicAuditRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicAuditRecord not implemented")
}
func (UnimplementedInformationTopicAuditRecordServiceServer) mustEmbedUnimplementedInformationTopicAuditRecordServiceServer() {
}

// UnsafeInformationTopicAuditRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformationTopicAuditRecordServiceServer will
// result in compilation errors.
type UnsafeInformationTopicAuditRecordServiceServer interface {
	mustEmbedUnimplementedInformationTopicAuditRecordServiceServer()
}

func RegisterInformationTopicAuditRecordServiceServer(s grpc.ServiceRegistrar, srv InformationTopicAuditRecordServiceServer) {
	s.RegisterService(&InformationTopicAuditRecordService_ServiceDesc, srv)
}

func _InformationTopicAuditRecordService_AddTopicAuditRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicAuditRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicAuditRecordServiceServer).AddTopicAuditRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicAuditRecordService_AddTopicAuditRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicAuditRecordServiceServer).AddTopicAuditRecord(ctx, req.(*AddTopicAuditRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InformationTopicAuditRecordService_GetTopicAuditRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicAuditRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationTopicAuditRecordServiceServer).GetTopicAuditRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InformationTopicAuditRecordService_GetTopicAuditRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationTopicAuditRecordServiceServer).GetTopicAuditRecord(ctx, req.(*GetTopicAuditRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InformationTopicAuditRecordService_ServiceDesc is the grpc.ServiceDesc for InformationTopicAuditRecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformationTopicAuditRecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "information.InformationTopicAuditRecordService",
	HandlerType: (*InformationTopicAuditRecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTopicAuditRecord",
			Handler:    _InformationTopicAuditRecordService_AddTopicAuditRecord_Handler,
		},
		{
			MethodName: "GetTopicAuditRecord",
			Handler:    _InformationTopicAuditRecordService_GetTopicAuditRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "information.proto",
}
