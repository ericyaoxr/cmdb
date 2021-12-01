// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/secret/pb/secret.proto

package secret

import (
	resource "github.com/ericyaoxr/cmdb/app/resource"
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

type Type int32

const (
	Type_API_KEY  Type = 0
	Type_PASSWORD Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "API_KEY",
		1: "PASSWORD",
	}
	Type_value = map[string]int32{
		"API_KEY":  0,
		"PASSWORD": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_app_secret_pb_secret_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_app_secret_pb_secret_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{0}
}

type CreateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 描述
	// @gotags: json:"description" validate:"required,lte=100"
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description" validate:"required,lte=100"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.Vendor `protobuf:"varint,2,opt,name=vendor,proto3,enum=ericyaoxr.cmdb.resource.Vendor" json:"vendor"`
	// 允许同步的区域
	// @gotags: json:"allow_regions"
	AllowRegions []string `protobuf:"bytes,3,rep,name=allow_regions,json=allowRegions,proto3" json:"allow_regions"`
	// 凭证类型
	// @gotags: json:"crendential_type"
	CrendentialType Type `protobuf:"varint,4,opt,name=crendential_type,json=crendentialType,proto3,enum=ericyaoxr.cmdb.secret.Type" json:"crendential_type"`
	// 服务地址, 云商不用填写
	// @gotags: json:"address"
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	// key
	// @gotags: json:"api_key" validate:"required,lte=100"
	ApiKey string `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3" json:"api_key" validate:"required,lte=100"`
	// secrete
	// @gotags: json:"api_secret" validate:"required,lte=100"
	ApiSecret string `protobuf:"bytes,7,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret" validate:"required,lte=100"`
	// 请求速率限制, 默认1秒5个
	// @gotags: json:"request_rate"
	RequestRate int32 `protobuf:"varint,8,opt,name=request_rate,json=requestRate,proto3" json:"request_rate"`
}

func (x *CreateSecretRequest) Reset() {
	*x = CreateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRequest) ProtoMessage() {}

func (x *CreateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSecretRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSecretRequest) GetVendor() resource.Vendor {
	if x != nil {
		return x.Vendor
	}
	return resource.Vendor(0)
}

func (x *CreateSecretRequest) GetAllowRegions() []string {
	if x != nil {
		return x.AllowRegions
	}
	return nil
}

func (x *CreateSecretRequest) GetCrendentialType() Type {
	if x != nil {
		return x.CrendentialType
	}
	return Type_API_KEY
}

func (x *CreateSecretRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateSecretRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *CreateSecretRequest) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *CreateSecretRequest) GetRequestRate() int32 {
	if x != nil {
		return x.RequestRate
	}
	return 0
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 全局唯一Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 创建时间
	// @gotags: json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at"`
	// 描述
	// @gotags: json:"description" validate:"required,lte=100"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" validate:"required,lte=100"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.Vendor `protobuf:"varint,4,opt,name=vendor,proto3,enum=ericyaoxr.cmdb.resource.Vendor" json:"vendor"`
	// 允许同步的区域
	// @gotags: json:"allow_regions"
	AllowRegions []string `protobuf:"bytes,5,rep,name=allow_regions,json=allowRegions,proto3" json:"allow_regions"`
	// 凭证类型
	// @gotags: json:"crendential_type"
	CrendentialType Type `protobuf:"varint,6,opt,name=crendential_type,json=crendentialType,proto3,enum=ericyaoxr.cmdb.secret.Type" json:"crendential_type"`
	// 服务地址, 云商不用填写
	// @gotags: json:"address"
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address"`
	// key
	// @gotags: json:"api_key" validate:"required,lte=100"
	ApiKey string `protobuf:"bytes,8,opt,name=api_key,json=apiKey,proto3" json:"api_key" validate:"required,lte=100"`
	// secrete
	// @gotags: json:"api_secret" validate:"required,lte=100"
	ApiSecret string `protobuf:"bytes,9,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret" validate:"required,lte=100"`
	// 请求速率限制, 默认1秒5个
	// @gotags: json:"request_rate"
	RequestRate int32 `protobuf:"varint,10,opt,name=request_rate,json=requestRate,proto3" json:"request_rate"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{1}
}

func (x *Secret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Secret) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Secret) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Secret) GetVendor() resource.Vendor {
	if x != nil {
		return x.Vendor
	}
	return resource.Vendor(0)
}

func (x *Secret) GetAllowRegions() []string {
	if x != nil {
		return x.AllowRegions
	}
	return nil
}

func (x *Secret) GetCrendentialType() Type {
	if x != nil {
		return x.CrendentialType
	}
	return Type_API_KEY
}

func (x *Secret) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Secret) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *Secret) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *Secret) GetRequestRate() int32 {
	if x != nil {
		return x.RequestRate
	}
	return 0
}

type QuerySecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page_size"
	PageSize uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	// @gotags: json:"page_number"
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number"`
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QuerySecretRequest) Reset() {
	*x = QuerySecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySecretRequest) ProtoMessage() {}

func (x *QuerySecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySecretRequest.ProtoReflect.Descriptor instead.
func (*QuerySecretRequest) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySecretRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QuerySecretRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QuerySecretRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type SecretSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"items"
	Items []*Secret `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *SecretSet) Reset() {
	*x = SecretSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretSet) ProtoMessage() {}

func (x *SecretSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretSet.ProtoReflect.Descriptor instead.
func (*SecretSet) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{3}
}

func (x *SecretSet) GetItems() []*Secret {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SecretSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DescribeSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeSecretRequest) Reset() {
	*x = DescribeSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSecretRequest) ProtoMessage() {}

func (x *DescribeSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSecretRequest.ProtoReflect.Descriptor instead.
func (*DescribeSecretRequest) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeSecretRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSecretRequest) Reset() {
	*x = DeleteSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_secret_pb_secret_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretRequest) ProtoMessage() {}

func (x *DeleteSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_secret_pb_secret_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return file_app_secret_pb_secret_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSecretRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_app_secret_pb_secret_proto protoreflect.FileDescriptor

var file_app_secret_pb_secret_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x72,
	0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0x1e, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x63,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78,
	0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x22, 0xf2, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x46, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x72, 0x69,
	0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x22, 0x6e, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x56, 0x0a,
	0x09, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x32, 0xfa, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x5a, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x29,
	0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x5d, 0x0a, 0x0e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2c, 0x2e,
	0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x72,
	0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x65, 0x72, 0x69,
	0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f,
	0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2f, 0x63, 0x6d,
	0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_secret_pb_secret_proto_rawDescOnce sync.Once
	file_app_secret_pb_secret_proto_rawDescData = file_app_secret_pb_secret_proto_rawDesc
)

func file_app_secret_pb_secret_proto_rawDescGZIP() []byte {
	file_app_secret_pb_secret_proto_rawDescOnce.Do(func() {
		file_app_secret_pb_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_secret_pb_secret_proto_rawDescData)
	})
	return file_app_secret_pb_secret_proto_rawDescData
}

var file_app_secret_pb_secret_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_secret_pb_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_app_secret_pb_secret_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: ericyaoxr.cmdb.secret.Type
	(*CreateSecretRequest)(nil),   // 1: ericyaoxr.cmdb.secret.CreateSecretRequest
	(*Secret)(nil),                // 2: ericyaoxr.cmdb.secret.Secret
	(*QuerySecretRequest)(nil),    // 3: ericyaoxr.cmdb.secret.QuerySecretRequest
	(*SecretSet)(nil),             // 4: ericyaoxr.cmdb.secret.SecretSet
	(*DescribeSecretRequest)(nil), // 5: ericyaoxr.cmdb.secret.DescribeSecretRequest
	(*DeleteSecretRequest)(nil),   // 6: ericyaoxr.cmdb.secret.DeleteSecretRequest
	(resource.Vendor)(0),          // 7: ericyaoxr.cmdb.resource.Vendor
}
var file_app_secret_pb_secret_proto_depIdxs = []int32{
	7, // 0: ericyaoxr.cmdb.secret.CreateSecretRequest.vendor:type_name -> ericyaoxr.cmdb.resource.Vendor
	0, // 1: ericyaoxr.cmdb.secret.CreateSecretRequest.crendential_type:type_name -> ericyaoxr.cmdb.secret.Type
	7, // 2: ericyaoxr.cmdb.secret.Secret.vendor:type_name -> ericyaoxr.cmdb.resource.Vendor
	0, // 3: ericyaoxr.cmdb.secret.Secret.crendential_type:type_name -> ericyaoxr.cmdb.secret.Type
	2, // 4: ericyaoxr.cmdb.secret.SecretSet.items:type_name -> ericyaoxr.cmdb.secret.Secret
	1, // 5: ericyaoxr.cmdb.secret.Service.CreateSecret:input_type -> ericyaoxr.cmdb.secret.CreateSecretRequest
	3, // 6: ericyaoxr.cmdb.secret.Service.QuerySecret:input_type -> ericyaoxr.cmdb.secret.QuerySecretRequest
	5, // 7: ericyaoxr.cmdb.secret.Service.DescribeSecret:input_type -> ericyaoxr.cmdb.secret.DescribeSecretRequest
	6, // 8: ericyaoxr.cmdb.secret.Service.DeleteSecret:input_type -> ericyaoxr.cmdb.secret.DeleteSecretRequest
	2, // 9: ericyaoxr.cmdb.secret.Service.CreateSecret:output_type -> ericyaoxr.cmdb.secret.Secret
	4, // 10: ericyaoxr.cmdb.secret.Service.QuerySecret:output_type -> ericyaoxr.cmdb.secret.SecretSet
	2, // 11: ericyaoxr.cmdb.secret.Service.DescribeSecret:output_type -> ericyaoxr.cmdb.secret.Secret
	2, // 12: ericyaoxr.cmdb.secret.Service.DeleteSecret:output_type -> ericyaoxr.cmdb.secret.Secret
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_app_secret_pb_secret_proto_init() }
func file_app_secret_pb_secret_proto_init() {
	if File_app_secret_pb_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_secret_pb_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretRequest); i {
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
		file_app_secret_pb_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_app_secret_pb_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySecretRequest); i {
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
		file_app_secret_pb_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretSet); i {
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
		file_app_secret_pb_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSecretRequest); i {
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
		file_app_secret_pb_secret_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretRequest); i {
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
			RawDescriptor: file_app_secret_pb_secret_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_secret_pb_secret_proto_goTypes,
		DependencyIndexes: file_app_secret_pb_secret_proto_depIdxs,
		EnumInfos:         file_app_secret_pb_secret_proto_enumTypes,
		MessageInfos:      file_app_secret_pb_secret_proto_msgTypes,
	}.Build()
	File_app_secret_pb_secret_proto = out.File
	file_app_secret_pb_secret_proto_rawDesc = nil
	file_app_secret_pb_secret_proto_goTypes = nil
	file_app_secret_pb_secret_proto_depIdxs = nil
}
