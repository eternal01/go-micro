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
	Information_Ping_FullMethodName                = "/information.Information/Ping"
	Information_AddTopic_FullMethodName            = "/information.Information/AddTopic"
	Information_GetTopic_FullMethodName            = "/information.Information/GetTopic"
	Information_GetTopicList_FullMethodName        = "/information.Information/GetTopicList"
	Information_UpdateTopic_FullMethodName         = "/information.Information/UpdateTopic"
	Information_DeleteTopic_FullMethodName         = "/information.Information/DeleteTopic"
	Information_AddTopicAuditRecord_FullMethodName = "/information.Information/AddTopicAuditRecord"
	Information_GetTopicAuditRecord_FullMethodName = "/information.Information/GetTopicAuditRecord"
)

// InformationClient is the client API for Information service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformationClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
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
	// 添加文章审核记录
	AddTopicAuditRecord(ctx context.Context, in *AddTopicAuditRecordRequest, opts ...grpc.CallOption) (*AddTopicAuditRecordResponse, error)
	// 获取文章审核记录
	GetTopicAuditRecord(ctx context.Context, in *GetTopicAuditRecordRequest, opts ...grpc.CallOption) (*GetTopicAuditRecordResponse, error)
}

type informationClient struct {
	cc grpc.ClientConnInterface
}

func NewInformationClient(cc grpc.ClientConnInterface) InformationClient {
	return &informationClient{cc}
}

func (c *informationClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Information_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) AddTopic(ctx context.Context, in *AddTopicRequest, opts ...grpc.CallOption) (*AddTopicResponse, error) {
	out := new(AddTopicResponse)
	err := c.cc.Invoke(ctx, Information_AddTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error) {
	out := new(GetTopicResponse)
	err := c.cc.Invoke(ctx, Information_GetTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) GetTopicList(ctx context.Context, in *GetTopicsRequest, opts ...grpc.CallOption) (*GetTopicsResponse, error) {
	out := new(GetTopicsResponse)
	err := c.cc.Invoke(ctx, Information_GetTopicList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*UpdateTopicResponse, error) {
	out := new(UpdateTopicResponse)
	err := c.cc.Invoke(ctx, Information_UpdateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicResponse, error) {
	out := new(DeleteTopicResponse)
	err := c.cc.Invoke(ctx, Information_DeleteTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) AddTopicAuditRecord(ctx context.Context, in *AddTopicAuditRecordRequest, opts ...grpc.CallOption) (*AddTopicAuditRecordResponse, error) {
	out := new(AddTopicAuditRecordResponse)
	err := c.cc.Invoke(ctx, Information_AddTopicAuditRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) GetTopicAuditRecord(ctx context.Context, in *GetTopicAuditRecordRequest, opts ...grpc.CallOption) (*GetTopicAuditRecordResponse, error) {
	out := new(GetTopicAuditRecordResponse)
	err := c.cc.Invoke(ctx, Information_GetTopicAuditRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformationServer is the server API for Information service.
// All implementations must embed UnimplementedInformationServer
// for forward compatibility
type InformationServer interface {
	Ping(context.Context, *Request) (*Response, error)
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
	// 添加文章审核记录
	AddTopicAuditRecord(context.Context, *AddTopicAuditRecordRequest) (*AddTopicAuditRecordResponse, error)
	// 获取文章审核记录
	GetTopicAuditRecord(context.Context, *GetTopicAuditRecordRequest) (*GetTopicAuditRecordResponse, error)
	mustEmbedUnimplementedInformationServer()
}

// UnimplementedInformationServer must be embedded to have forward compatible implementations.
type UnimplementedInformationServer struct {
}

func (UnimplementedInformationServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInformationServer) AddTopic(context.Context, *AddTopicRequest) (*AddTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopic not implemented")
}
func (UnimplementedInformationServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedInformationServer) GetTopicList(context.Context, *GetTopicsRequest) (*GetTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicList not implemented")
}
func (UnimplementedInformationServer) UpdateTopic(context.Context, *UpdateTopicRequest) (*UpdateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (UnimplementedInformationServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedInformationServer) AddTopicAuditRecord(context.Context, *AddTopicAuditRecordRequest) (*AddTopicAuditRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopicAuditRecord not implemented")
}
func (UnimplementedInformationServer) GetTopicAuditRecord(context.Context, *GetTopicAuditRecordRequest) (*GetTopicAuditRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicAuditRecord not implemented")
}
func (UnimplementedInformationServer) mustEmbedUnimplementedInformationServer() {}

// UnsafeInformationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformationServer will
// result in compilation errors.
type UnsafeInformationServer interface {
	mustEmbedUnimplementedInformationServer()
}

func RegisterInformationServer(s grpc.ServiceRegistrar, srv InformationServer) {
	s.RegisterService(&Information_ServiceDesc, srv)
}

func _Information_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_AddTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).AddTopic(ctx, req.(*AddTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_GetTopicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).GetTopicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_GetTopicList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).GetTopicList(ctx, req.(*GetTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_UpdateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_AddTopicAuditRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicAuditRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).AddTopicAuditRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_AddTopicAuditRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).AddTopicAuditRecord(ctx, req.(*AddTopicAuditRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_GetTopicAuditRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicAuditRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).GetTopicAuditRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_GetTopicAuditRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).GetTopicAuditRecord(ctx, req.(*GetTopicAuditRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Information_ServiceDesc is the grpc.ServiceDesc for Information service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Information_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "information.Information",
	HandlerType: (*InformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Information_Ping_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _Information_AddTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _Information_GetTopic_Handler,
		},
		{
			MethodName: "GetTopicList",
			Handler:    _Information_GetTopicList_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _Information_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Information_DeleteTopic_Handler,
		},
		{
			MethodName: "AddTopicAuditRecord",
			Handler:    _Information_AddTopicAuditRecord_Handler,
		},
		{
			MethodName: "GetTopicAuditRecord",
			Handler:    _Information_GetTopicAuditRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "information.proto",
}
