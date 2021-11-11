// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/config/pb/config.proto

package config

import (
	reflect "reflect"
	sync "sync"

	application "github.com/ericyaoxr/cmdb/app/application"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateMode int32

const (
	UpdateMode_PUT   UpdateMode = 0
	UpdateMode_PATCH UpdateMode = 1
)

// Enum value maps for UpdateMode.
var (
	UpdateMode_name = map[int32]string{
		0: "PUT",
		1: "PATCH",
	}
	UpdateMode_value = map[string]int32{
		"PUT":   0,
		"PATCH": 1,
	}
)

func (x UpdateMode) Enum() *UpdateMode {
	p := new(UpdateMode)
	*p = x
	return p
}

func (x UpdateMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateMode) Descriptor() protoreflect.EnumDescriptor {
	return file_app_config_pb_config_proto_enumTypes[0].Descriptor()
}

func (UpdateMode) Type() protoreflect.EnumType {
	return &file_app_config_pb_config_proto_enumTypes[0]
}

func (x UpdateMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateMode.Descriptor instead.
func (UpdateMode) EnumDescriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *application.Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Describe    *Describe                `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetApplication() *application.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *Config) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 关联Resource
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 关联Application
	// @gotags: json:"application_id"
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// 配置的key名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 主机地址
	// @gotags: json:"host"
	Host string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	// 端口
	// @gotags: json:"port"
	Port int32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	// 环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,6,opt,name=env,proto3" json:"env,omitempty"`
	// 配置类型
	// @gotags: json:"type"
	Type string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	// 来源
	// @gotags: json:"source"
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	// 创建时间
	// @gotags: json:"create_at"
	CreateAt int64 `protobuf:"varint,9,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	// 创建时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,10,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Describe) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Describe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Describe) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Describe) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Describe) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Describe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Describe) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Describe) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Describe) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type QueryConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Keywords   string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryConfigRequest) Reset() {
	*x = QueryConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryConfigRequest) ProtoMessage() {}

func (x *QueryConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryConfigRequest.ProtoReflect.Descriptor instead.
func (*QueryConfigRequest) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{2}
}

func (x *QueryConfigRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryConfigRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QueryConfigRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type ConfigSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Config `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ConfigSet) Reset() {
	*x = ConfigSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSet) ProtoMessage() {}

func (x *ConfigSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSet.ProtoReflect.Descriptor instead.
func (*ConfigSet) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigSet) GetItems() []*Config {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ConfigSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: json:"update_mode"
	UpdateMode UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=ericyaoxr.cmdb.config.UpdateMode" json:"update_mode,omitempty"`
	// @gotags: json:"data" validate:"required"
	UpdateConfigData *UpdateConfigData `protobuf:"bytes,3,opt,name=update_config_data,json=updateConfigData,proto3" json:"update_config_data,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateConfigRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateConfigRequest) GetUpdateMode() UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return UpdateMode_PUT
}

func (x *UpdateConfigRequest) GetUpdateConfigData() *UpdateConfigData {
	if x != nil {
		return x.UpdateConfigData
	}
	return nil
}

type UpdateConfigData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateConfigData) Reset() {
	*x = UpdateConfigData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigData) ProtoMessage() {}

func (x *UpdateConfigData) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigData.ProtoReflect.Descriptor instead.
func (*UpdateConfigData) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConfigData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DescribeConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeConfigRequest) Reset() {
	*x = DescribeConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeConfigRequest) ProtoMessage() {}

func (x *DescribeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeConfigRequest.ProtoReflect.Descriptor instead.
func (*DescribeConfigRequest) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeConfigRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteConfigRequest) Reset() {
	*x = DeleteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_pb_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigRequest) ProtoMessage() {}

func (x *DeleteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_pb_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) {
	return file_app_config_pb_config_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteConfigRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_app_config_pb_config_proto protoreflect.FileDescriptor

var file_app_config_pb_config_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x72,
	0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0xf5, 0x01, 0x0a,
	0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x74, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc0, 0x01, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x22, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x2a, 0x20, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x10, 0x01, 0x32, 0xd0, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1d,
	0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29,
	0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x22, 0x00, 0x12, 0x5b, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x2e,
	0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63,
	0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x2e, 0x65,
	0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x72, 0x69,
	0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x2e, 0x65, 0x72,
	0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61,
	0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72,
	0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_config_pb_config_proto_rawDescOnce sync.Once
	file_app_config_pb_config_proto_rawDescData = file_app_config_pb_config_proto_rawDesc
)

func file_app_config_pb_config_proto_rawDescGZIP() []byte {
	file_app_config_pb_config_proto_rawDescOnce.Do(func() {
		file_app_config_pb_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_config_pb_config_proto_rawDescData)
	})
	return file_app_config_pb_config_proto_rawDescData
}

var file_app_config_pb_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_config_pb_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_config_pb_config_proto_goTypes = []interface{}{
	(UpdateMode)(0),                 // 0: ericyaoxr.cmdb.config.UpdateMode
	(*Config)(nil),                  // 1: ericyaoxr.cmdb.config.Config
	(*Describe)(nil),                // 2: ericyaoxr.cmdb.config.Describe
	(*QueryConfigRequest)(nil),      // 3: ericyaoxr.cmdb.config.QueryConfigRequest
	(*ConfigSet)(nil),               // 4: ericyaoxr.cmdb.config.ConfigSet
	(*UpdateConfigRequest)(nil),     // 5: ericyaoxr.cmdb.config.UpdateConfigRequest
	(*UpdateConfigData)(nil),        // 6: ericyaoxr.cmdb.config.UpdateConfigData
	(*DescribeConfigRequest)(nil),   // 7: ericyaoxr.cmdb.config.DescribeConfigRequest
	(*DeleteConfigRequest)(nil),     // 8: ericyaoxr.cmdb.config.DeleteConfigRequest
	(*application.Application)(nil), // 9: ericyaoxr.cmdb.application.Application
}
var file_app_config_pb_config_proto_depIdxs = []int32{
	9,  // 0: ericyaoxr.cmdb.config.Config.application:type_name -> ericyaoxr.cmdb.application.Application
	2,  // 1: ericyaoxr.cmdb.config.Config.describe:type_name -> ericyaoxr.cmdb.config.Describe
	1,  // 2: ericyaoxr.cmdb.config.ConfigSet.items:type_name -> ericyaoxr.cmdb.config.Config
	0,  // 3: ericyaoxr.cmdb.config.UpdateConfigRequest.update_mode:type_name -> ericyaoxr.cmdb.config.UpdateMode
	6,  // 4: ericyaoxr.cmdb.config.UpdateConfigRequest.update_config_data:type_name -> ericyaoxr.cmdb.config.UpdateConfigData
	1,  // 5: ericyaoxr.cmdb.config.Service.SaveConfig:input_type -> ericyaoxr.cmdb.config.Config
	3,  // 6: ericyaoxr.cmdb.config.Service.QueryConfig:input_type -> ericyaoxr.cmdb.config.QueryConfigRequest
	5,  // 7: ericyaoxr.cmdb.config.Service.UpdateConfig:input_type -> ericyaoxr.cmdb.config.UpdateConfigRequest
	7,  // 8: ericyaoxr.cmdb.config.Service.DescribeConfig:input_type -> ericyaoxr.cmdb.config.DescribeConfigRequest
	8,  // 9: ericyaoxr.cmdb.config.Service.DeleteConfig:input_type -> ericyaoxr.cmdb.config.DeleteConfigRequest
	1,  // 10: ericyaoxr.cmdb.config.Service.SaveConfig:output_type -> ericyaoxr.cmdb.config.Config
	4,  // 11: ericyaoxr.cmdb.config.Service.QueryConfig:output_type -> ericyaoxr.cmdb.config.ConfigSet
	1,  // 12: ericyaoxr.cmdb.config.Service.UpdateConfig:output_type -> ericyaoxr.cmdb.config.Config
	1,  // 13: ericyaoxr.cmdb.config.Service.DescribeConfig:output_type -> ericyaoxr.cmdb.config.Config
	1,  // 14: ericyaoxr.cmdb.config.Service.DeleteConfig:output_type -> ericyaoxr.cmdb.config.Config
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_app_config_pb_config_proto_init() }
func file_app_config_pb_config_proto_init() {
	if File_app_config_pb_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_config_pb_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_app_config_pb_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
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
		file_app_config_pb_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryConfigRequest); i {
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
		file_app_config_pb_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSet); i {
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
		file_app_config_pb_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_app_config_pb_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigData); i {
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
		file_app_config_pb_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeConfigRequest); i {
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
		file_app_config_pb_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigRequest); i {
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
			RawDescriptor: file_app_config_pb_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_config_pb_config_proto_goTypes,
		DependencyIndexes: file_app_config_pb_config_proto_depIdxs,
		EnumInfos:         file_app_config_pb_config_proto_enumTypes,
		MessageInfos:      file_app_config_pb_config_proto_msgTypes,
	}.Build()
	File_app_config_pb_config_proto = out.File
	file_app_config_pb_config_proto_rawDesc = nil
	file_app_config_pb_config_proto_goTypes = nil
	file_app_config_pb_config_proto_depIdxs = nil
}
