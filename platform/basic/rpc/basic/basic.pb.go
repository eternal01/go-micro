// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: basic.proto

package basic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type GetRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetRegionRequest) Reset() {
	*x = GetRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionRequest) ProtoMessage() {}

func (x *GetRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionRequest.ProtoReflect.Descriptor instead.
func (*GetRegionRequest) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{2}
}

func (x *GetRegionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ParentId int64  `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetRegionResponse) Reset() {
	*x = GetRegionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionResponse) ProtoMessage() {}

func (x *GetRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionResponse.ProtoReflect.Descriptor instead.
func (*GetRegionResponse) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{3}
}

func (x *GetRegionResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRegionResponse) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetRegionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRegionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId int64 `protobuf:"varint,1,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
}

func (x *GetRegionsRequest) Reset() {
	*x = GetRegionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionsRequest) ProtoMessage() {}

func (x *GetRegionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionsRequest.ProtoReflect.Descriptor instead.
func (*GetRegionsRequest) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{4}
}

func (x *GetRegionsRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

type GetRegionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GetRegionChild `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *GetRegionsResponse) Reset() {
	*x = GetRegionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionsResponse) ProtoMessage() {}

func (x *GetRegionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionsResponse.ProtoReflect.Descriptor instead.
func (*GetRegionsResponse) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{5}
}

func (x *GetRegionsResponse) GetList() []*GetRegionChild {
	if x != nil {
		return x.List
	}
	return nil
}

type GetRegionChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ParentId int64  `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetRegionChild) Reset() {
	*x = GetRegionChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionChild) ProtoMessage() {}

func (x *GetRegionChild) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionChild.ProtoReflect.Descriptor instead.
func (*GetRegionChild) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{6}
}

func (x *GetRegionChild) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRegionChild) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetRegionChild) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetIndustryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndustryId string `protobuf:"bytes,1,opt,name=IndustryId,proto3" json:"IndustryId,omitempty"`
}

func (x *GetIndustryRequest) Reset() {
	*x = GetIndustryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIndustryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIndustryRequest) ProtoMessage() {}

func (x *GetIndustryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIndustryRequest.ProtoReflect.Descriptor instead.
func (*GetIndustryRequest) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{7}
}

func (x *GetIndustryRequest) GetIndustryId() string {
	if x != nil {
		return x.IndustryId
	}
	return ""
}

type GetIndustryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IndustryId  string `protobuf:"bytes,2,opt,name=IndustryId,proto3" json:"IndustryId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ParentId    string `protobuf:"bytes,4,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	LevelType   int64  `protobuf:"varint,5,opt,name=LevelType,proto3" json:"LevelType,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *GetIndustryResponse) Reset() {
	*x = GetIndustryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIndustryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIndustryResponse) ProtoMessage() {}

func (x *GetIndustryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIndustryResponse.ProtoReflect.Descriptor instead.
func (*GetIndustryResponse) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{8}
}

func (x *GetIndustryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetIndustryResponse) GetIndustryId() string {
	if x != nil {
		return x.IndustryId
	}
	return ""
}

func (x *GetIndustryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetIndustryResponse) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetIndustryResponse) GetLevelType() int64 {
	if x != nil {
		return x.LevelType
	}
	return 0
}

func (x *GetIndustryResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetIndustriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
}

func (x *GetIndustriesRequest) Reset() {
	*x = GetIndustriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIndustriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIndustriesRequest) ProtoMessage() {}

func (x *GetIndustriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIndustriesRequest.ProtoReflect.Descriptor instead.
func (*GetIndustriesRequest) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{9}
}

func (x *GetIndustriesRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type GetIndustriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GetIndustryChild `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *GetIndustriesResponse) Reset() {
	*x = GetIndustriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIndustriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIndustriesResponse) ProtoMessage() {}

func (x *GetIndustriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIndustriesResponse.ProtoReflect.Descriptor instead.
func (*GetIndustriesResponse) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{10}
}

func (x *GetIndustriesResponse) GetList() []*GetIndustryChild {
	if x != nil {
		return x.List
	}
	return nil
}

type GetIndustryChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IndustryId  string `protobuf:"bytes,2,opt,name=IndustryId,proto3" json:"IndustryId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ParentId    string `protobuf:"bytes,4,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	LevelType   int64  `protobuf:"varint,5,opt,name=LevelType,proto3" json:"LevelType,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *GetIndustryChild) Reset() {
	*x = GetIndustryChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIndustryChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIndustryChild) ProtoMessage() {}

func (x *GetIndustryChild) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIndustryChild.ProtoReflect.Descriptor instead.
func (*GetIndustryChild) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{11}
}

func (x *GetIndustryChild) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetIndustryChild) GetIndustryId() string {
	if x != nil {
		return x.IndustryId
	}
	return ""
}

func (x *GetIndustryChild) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetIndustryChild) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetIndustryChild) GetLevelType() int64 {
	if x != nil {
		return x.LevelType
	}
	return 0
}

func (x *GetIndustryChild) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId int64  `protobuf:"varint,1,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *AddRegionRequest) Reset() {
	*x = AddRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRegionRequest) ProtoMessage() {}

func (x *AddRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRegionRequest.ProtoReflect.Descriptor instead.
func (*AddRegionRequest) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{12}
}

func (x *AddRegionRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *AddRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ParentId int64  `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *AddRegionResponse) Reset() {
	*x = AddRegionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRegionResponse) ProtoMessage() {}

func (x *AddRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRegionResponse.ProtoReflect.Descriptor instead.
func (*AddRegionResponse) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{13}
}

func (x *AddRegionResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddRegionResponse) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *AddRegionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_basic_proto protoreflect.FileDescriptor

var file_basic_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x6e, 0x67, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x50,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x53, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x85, 0x03, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x12, 0x27, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_basic_proto_rawDescOnce sync.Once
	file_basic_proto_rawDescData = file_basic_proto_rawDesc
)

func file_basic_proto_rawDescGZIP() []byte {
	file_basic_proto_rawDescOnce.Do(func() {
		file_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_proto_rawDescData)
	})
	return file_basic_proto_rawDescData
}

var file_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_basic_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: basic.Request
	(*Response)(nil),              // 1: basic.Response
	(*GetRegionRequest)(nil),      // 2: basic.GetRegionRequest
	(*GetRegionResponse)(nil),     // 3: basic.GetRegionResponse
	(*GetRegionsRequest)(nil),     // 4: basic.GetRegionsRequest
	(*GetRegionsResponse)(nil),    // 5: basic.GetRegionsResponse
	(*GetRegionChild)(nil),        // 6: basic.GetRegionChild
	(*GetIndustryRequest)(nil),    // 7: basic.GetIndustryRequest
	(*GetIndustryResponse)(nil),   // 8: basic.GetIndustryResponse
	(*GetIndustriesRequest)(nil),  // 9: basic.GetIndustriesRequest
	(*GetIndustriesResponse)(nil), // 10: basic.GetIndustriesResponse
	(*GetIndustryChild)(nil),      // 11: basic.GetIndustryChild
	(*AddRegionRequest)(nil),      // 12: basic.AddRegionRequest
	(*AddRegionResponse)(nil),     // 13: basic.AddRegionResponse
}
var file_basic_proto_depIdxs = []int32{
	6,  // 0: basic.GetRegionsResponse.List:type_name -> basic.GetRegionChild
	11, // 1: basic.GetIndustriesResponse.List:type_name -> basic.GetIndustryChild
	0,  // 2: basic.Basic.Ping:input_type -> basic.Request
	2,  // 3: basic.Basic.getRegion:input_type -> basic.GetRegionRequest
	4,  // 4: basic.Basic.getRegions:input_type -> basic.GetRegionsRequest
	7,  // 5: basic.Basic.getIndustry:input_type -> basic.GetIndustryRequest
	9,  // 6: basic.Basic.getIndustries:input_type -> basic.GetIndustriesRequest
	12, // 7: basic.Basic.addRegion:input_type -> basic.AddRegionRequest
	1,  // 8: basic.Basic.Ping:output_type -> basic.Response
	3,  // 9: basic.Basic.getRegion:output_type -> basic.GetRegionResponse
	5,  // 10: basic.Basic.getRegions:output_type -> basic.GetRegionsResponse
	8,  // 11: basic.Basic.getIndustry:output_type -> basic.GetIndustryResponse
	10, // 12: basic.Basic.getIndustries:output_type -> basic.GetIndustriesResponse
	13, // 13: basic.Basic.addRegion:output_type -> basic.AddRegionResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_basic_proto_init() }
func file_basic_proto_init() {
	if File_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionChild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIndustryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIndustryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIndustriesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIndustriesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIndustryChild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRegionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_basic_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRegionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_proto_goTypes,
		DependencyIndexes: file_basic_proto_depIdxs,
		MessageInfos:      file_basic_proto_msgTypes,
	}.Build()
	File_basic_proto = out.File
	file_basic_proto_rawDesc = nil
	file_basic_proto_goTypes = nil
	file_basic_proto_depIdxs = nil
}
