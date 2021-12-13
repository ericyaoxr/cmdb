// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/rds/pb/rds.proto

package rds

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
	return file_app_rds_pb_rds_proto_enumTypes[0].Descriptor()
}

func (UpdateMode) Type() protoreflect.EnumType {
	return &file_app_rds_pb_rds_proto_enumTypes[0]
}

func (x UpdateMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateMode.Descriptor instead.
func (UpdateMode) EnumDescriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_PENDING   Status = 0
	Status_RUNNING   Status = 1
	Status_ISOLATING Status = 2
	Status_ISOLATED  Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "RUNNING",
		2: "ISOLATING",
		3: "ISOLATED",
	}
	Status_value = map[string]int32{
		"PENDING":   0,
		"RUNNING":   1,
		"ISOLATING": 2,
		"ISOLATED":  3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_app_rds_pb_rds_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_app_rds_pb_rds_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{1}
}

type PayType int32

const (
	PayType_PREPAID          PayType = 0
	PayType_POSTPAID_BY_HOUR PayType = 1
)

// Enum value maps for PayType.
var (
	PayType_name = map[int32]string{
		0: "PREPAID",
		1: "POSTPAID_BY_HOUR",
	}
	PayType_value = map[string]int32{
		"PREPAID":          0,
		"POSTPAID_BY_HOUR": 1,
	}
)

func (x PayType) Enum() *PayType {
	p := new(PayType)
	*p = x
	return p
}

func (x PayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_rds_pb_rds_proto_enumTypes[2].Descriptor()
}

func (PayType) Type() protoreflect.EnumType {
	return &file_app_rds_pb_rds_proto_enumTypes[2]
}

func (x PayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayType.Descriptor instead.
func (PayType) EnumDescriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{2}
}

type DescribeBy int32

const (
	DescribeBy_RDS_ID      DescribeBy = 0
	DescribeBy_INSTANCE_ID DescribeBy = 1
)

// Enum value maps for DescribeBy.
var (
	DescribeBy_name = map[int32]string{
		0: "RDS_ID",
		1: "INSTANCE_ID",
	}
	DescribeBy_value = map[string]int32{
		"RDS_ID":      0,
		"INSTANCE_ID": 1,
	}
)

func (x DescribeBy) Enum() *DescribeBy {
	p := new(DescribeBy)
	*p = x
	return p
}

func (x DescribeBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeBy) Descriptor() protoreflect.EnumDescriptor {
	return file_app_rds_pb_rds_proto_enumTypes[3].Descriptor()
}

func (DescribeBy) Type() protoreflect.EnumType {
	return &file_app_rds_pb_rds_proto_enumTypes[3]
}

func (x DescribeBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeBy.Descriptor instead.
func (DescribeBy) EnumDescriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{3}
}

type Rds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        *resource.Base        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information,omitempty"`
	Describe    *Describe             `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
}

func (x *Rds) Reset() {
	*x = Rds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rds) ProtoMessage() {}

func (x *Rds) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rds.ProtoReflect.Descriptor instead.
func (*Rds) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{0}
}

func (x *Rds) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Rds) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *Rds) GetDescribe() *Describe {
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
	// @gotags: json:"resource_id"
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// 核数
	// @gotags: json:"cpu"
	Cpu int64 `protobuf:"varint,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// 内存
	// @gotags: json:"memory"
	Memory int64 `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	// 硬盘容量，单位为 GB
	// @gotags: json:"volume"
	Volume int64 `protobuf:"varint,4,opt,name=volume,proto3" json:"volume,omitempty"`
	// 实例引擎版本
	// @gotags: json:"engine_version"
	EngineVersion string `protobuf:"bytes,5,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	// 节点数
	// @gotags: json:"instance_nodes"
	InstanceNodes int64 `protobuf:"varint,6,opt,name=instance_nodes,json=instanceNodes,proto3" json:"instance_nodes,omitempty"`
	// 实例类型，可能的返回值：1-主实例；2-灾备实例；3-只读实例
	// @gotags: json:"instance_type"
	InstanceType int64 `protobuf:"varint,7,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// 内网端口号
	// @gotags: json:"port"
	Port int64 `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	// Master
	// @gotags: json:"master"
	Master string `protobuf:"bytes,9,opt,name=master,proto3" json:"master,omitempty"`
	// Slave
	// @gotags: json:"slave"
	Slave string `protobuf:"bytes,10,opt,name=slave,proto3" json:"slave,omitempty"`
	// QPS 每秒查询数量
	// @gotags: json:"qps"
	Qps int64 `protobuf:"varint,11,opt,name=qps,proto3" json:"qps,omitempty"`
	// 可用区部署方式。可能的值为：0 - 单可用区；1 - 多可用区
	// @gotags: json:"deploy_mode"
	DeployMode int64 `protobuf:"varint,12,opt,name=deploy_mode,json=deployMode,proto3" json:"deploy_mode,omitempty"`
	// 外网端口
	// @gotags: json:"wan_port"
	WanPort int64 `protobuf:"varint,13,opt,name=wan_port,json=wanPort,proto3" json:"wan_port,omitempty"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[1]
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
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Describe) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Describe) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Describe) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Describe) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *Describe) GetInstanceNodes() int64 {
	if x != nil {
		return x.InstanceNodes
	}
	return 0
}

func (x *Describe) GetInstanceType() int64 {
	if x != nil {
		return x.InstanceType
	}
	return 0
}

func (x *Describe) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Describe) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *Describe) GetSlave() string {
	if x != nil {
		return x.Slave
	}
	return ""
}

func (x *Describe) GetQps() int64 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *Describe) GetDeployMode() int64 {
	if x != nil {
		return x.DeployMode
	}
	return 0
}

func (x *Describe) GetWanPort() int64 {
	if x != nil {
		return x.WanPort
	}
	return 0
}

type QueryRdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Keywords   string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryRdsRequest) Reset() {
	*x = QueryRdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRdsRequest) ProtoMessage() {}

func (x *QueryRdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRdsRequest.ProtoReflect.Descriptor instead.
func (*QueryRdsRequest) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRdsRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryRdsRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QueryRdsRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type RdsSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Rds `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *RdsSet) Reset() {
	*x = RdsSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RdsSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RdsSet) ProtoMessage() {}

func (x *RdsSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RdsSet.ProtoReflect.Descriptor instead.
func (*RdsSet) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{3}
}

func (x *RdsSet) GetItems() []*Rds {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RdsSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateRdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: json:"update_mode"
	UpdateMode UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=ericyaoxr.cmdb.rds.UpdateMode" json:"update_mode,omitempty"`
	// @gotags: json:"data" validate:"required"
	UpdateRdsData *UpdateRdsData `protobuf:"bytes,3,opt,name=update_rds_data,json=updateRdsData,proto3" json:"update_rds_data,omitempty"`
}

func (x *UpdateRdsRequest) Reset() {
	*x = UpdateRdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRdsRequest) ProtoMessage() {}

func (x *UpdateRdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRdsRequest.ProtoReflect.Descriptor instead.
func (*UpdateRdsRequest) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRdsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRdsRequest) GetUpdateMode() UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return UpdateMode_PUT
}

func (x *UpdateRdsRequest) GetUpdateRdsData() *UpdateRdsData {
	if x != nil {
		return x.UpdateRdsData
	}
	return nil
}

type UpdateRdsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Information *resource.Information `protobuf:"bytes,1,opt,name=information,proto3" json:"information,omitempty"`
	Describe    *Describe             `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe,omitempty"`
}

func (x *UpdateRdsData) Reset() {
	*x = UpdateRdsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRdsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRdsData) ProtoMessage() {}

func (x *UpdateRdsData) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRdsData.ProtoReflect.Descriptor instead.
func (*UpdateRdsData) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRdsData) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *UpdateRdsData) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type DescribeRdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"describe_by"
	DescribeBy DescribeBy `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=ericyaoxr.cmdb.rds.DescribeBy" json:"describe_by,omitempty"`
	// @gotags: json:"value" validate:"required"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribeRdsRequest) Reset() {
	*x = DescribeRdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRdsRequest) ProtoMessage() {}

func (x *DescribeRdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRdsRequest.ProtoReflect.Descriptor instead.
func (*DescribeRdsRequest) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeRdsRequest) GetDescribeBy() DescribeBy {
	if x != nil {
		return x.DescribeBy
	}
	return DescribeBy_RDS_ID
}

func (x *DescribeRdsRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeleteRdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRdsRequest) Reset() {
	*x = DeleteRdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_rds_pb_rds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRdsRequest) ProtoMessage() {}

func (x *DeleteRdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_rds_pb_rds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRdsRequest.ProtoReflect.Descriptor instead.
func (*DeleteRdsRequest) Descriptor() ([]byte, []int) {
	return file_app_rds_pb_rds_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRdsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_app_rds_pb_rds_proto protoreflect.FileDescriptor

var file_app_rds_pb_rds_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78,
	0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x1a, 0x1e, 0x61, 0x70, 0x70, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x03, 0x52,
	0x64, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72, 0x69,
	0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x70, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x77, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x77, 0x61, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x6b, 0x0a, 0x0f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x06, 0x52, 0x64, 0x73, 0x53, 0x65,
	0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xae, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78,
	0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x6b, 0x0a, 0x12, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f,
	0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x20, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55,
	0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x2a, 0x3f,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a,
	0x2c, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52,
	0x45, 0x50, 0x41, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4f, 0x53, 0x54, 0x50,
	0x41, 0x49, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x01, 0x2a, 0x29, 0x0a,
	0x0a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x44, 0x53, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x53, 0x54, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x32, 0x85, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x52, 0x64, 0x73, 0x12,
	0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79,
	0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64,
	0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x64, 0x73, 0x12,
	0x23, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73, 0x53, 0x65, 0x74,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x64, 0x73, 0x12,
	0x24, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78,
	0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x64, 0x73, 0x12,
	0x26, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61,
	0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x64, 0x73, 0x12,
	0x24, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78,
	0x72, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x64, 0x73, 0x22, 0x00,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x72, 0x69, 0x63, 0x79, 0x61, 0x6f, 0x78, 0x72, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_rds_pb_rds_proto_rawDescOnce sync.Once
	file_app_rds_pb_rds_proto_rawDescData = file_app_rds_pb_rds_proto_rawDesc
)

func file_app_rds_pb_rds_proto_rawDescGZIP() []byte {
	file_app_rds_pb_rds_proto_rawDescOnce.Do(func() {
		file_app_rds_pb_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_rds_pb_rds_proto_rawDescData)
	})
	return file_app_rds_pb_rds_proto_rawDescData
}

var file_app_rds_pb_rds_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_app_rds_pb_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_rds_pb_rds_proto_goTypes = []interface{}{
	(UpdateMode)(0),              // 0: ericyaoxr.cmdb.rds.UpdateMode
	(Status)(0),                  // 1: ericyaoxr.cmdb.rds.Status
	(PayType)(0),                 // 2: ericyaoxr.cmdb.rds.PayType
	(DescribeBy)(0),              // 3: ericyaoxr.cmdb.rds.DescribeBy
	(*Rds)(nil),                  // 4: ericyaoxr.cmdb.rds.Rds
	(*Describe)(nil),             // 5: ericyaoxr.cmdb.rds.Describe
	(*QueryRdsRequest)(nil),      // 6: ericyaoxr.cmdb.rds.QueryRdsRequest
	(*RdsSet)(nil),               // 7: ericyaoxr.cmdb.rds.RdsSet
	(*UpdateRdsRequest)(nil),     // 8: ericyaoxr.cmdb.rds.UpdateRdsRequest
	(*UpdateRdsData)(nil),        // 9: ericyaoxr.cmdb.rds.UpdateRdsData
	(*DescribeRdsRequest)(nil),   // 10: ericyaoxr.cmdb.rds.DescribeRdsRequest
	(*DeleteRdsRequest)(nil),     // 11: ericyaoxr.cmdb.rds.DeleteRdsRequest
	(*resource.Base)(nil),        // 12: ericyaoxr.cmdb.resource.Base
	(*resource.Information)(nil), // 13: ericyaoxr.cmdb.resource.Information
}
var file_app_rds_pb_rds_proto_depIdxs = []int32{
	12, // 0: ericyaoxr.cmdb.rds.Rds.base:type_name -> ericyaoxr.cmdb.resource.Base
	13, // 1: ericyaoxr.cmdb.rds.Rds.information:type_name -> ericyaoxr.cmdb.resource.Information
	5,  // 2: ericyaoxr.cmdb.rds.Rds.describe:type_name -> ericyaoxr.cmdb.rds.Describe
	4,  // 3: ericyaoxr.cmdb.rds.RdsSet.items:type_name -> ericyaoxr.cmdb.rds.Rds
	0,  // 4: ericyaoxr.cmdb.rds.UpdateRdsRequest.update_mode:type_name -> ericyaoxr.cmdb.rds.UpdateMode
	9,  // 5: ericyaoxr.cmdb.rds.UpdateRdsRequest.update_rds_data:type_name -> ericyaoxr.cmdb.rds.UpdateRdsData
	13, // 6: ericyaoxr.cmdb.rds.UpdateRdsData.information:type_name -> ericyaoxr.cmdb.resource.Information
	5,  // 7: ericyaoxr.cmdb.rds.UpdateRdsData.describe:type_name -> ericyaoxr.cmdb.rds.Describe
	3,  // 8: ericyaoxr.cmdb.rds.DescribeRdsRequest.describe_by:type_name -> ericyaoxr.cmdb.rds.DescribeBy
	4,  // 9: ericyaoxr.cmdb.rds.Service.SaveRds:input_type -> ericyaoxr.cmdb.rds.Rds
	6,  // 10: ericyaoxr.cmdb.rds.Service.QueryRds:input_type -> ericyaoxr.cmdb.rds.QueryRdsRequest
	8,  // 11: ericyaoxr.cmdb.rds.Service.UpdateRds:input_type -> ericyaoxr.cmdb.rds.UpdateRdsRequest
	10, // 12: ericyaoxr.cmdb.rds.Service.DescribeRds:input_type -> ericyaoxr.cmdb.rds.DescribeRdsRequest
	11, // 13: ericyaoxr.cmdb.rds.Service.DeleteRds:input_type -> ericyaoxr.cmdb.rds.DeleteRdsRequest
	4,  // 14: ericyaoxr.cmdb.rds.Service.SaveRds:output_type -> ericyaoxr.cmdb.rds.Rds
	7,  // 15: ericyaoxr.cmdb.rds.Service.QueryRds:output_type -> ericyaoxr.cmdb.rds.RdsSet
	4,  // 16: ericyaoxr.cmdb.rds.Service.UpdateRds:output_type -> ericyaoxr.cmdb.rds.Rds
	4,  // 17: ericyaoxr.cmdb.rds.Service.DescribeRds:output_type -> ericyaoxr.cmdb.rds.Rds
	4,  // 18: ericyaoxr.cmdb.rds.Service.DeleteRds:output_type -> ericyaoxr.cmdb.rds.Rds
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_app_rds_pb_rds_proto_init() }
func file_app_rds_pb_rds_proto_init() {
	if File_app_rds_pb_rds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_rds_pb_rds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rds); i {
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
		file_app_rds_pb_rds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_app_rds_pb_rds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRdsRequest); i {
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
		file_app_rds_pb_rds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RdsSet); i {
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
		file_app_rds_pb_rds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRdsRequest); i {
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
		file_app_rds_pb_rds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRdsData); i {
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
		file_app_rds_pb_rds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRdsRequest); i {
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
		file_app_rds_pb_rds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRdsRequest); i {
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
			RawDescriptor: file_app_rds_pb_rds_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_rds_pb_rds_proto_goTypes,
		DependencyIndexes: file_app_rds_pb_rds_proto_depIdxs,
		EnumInfos:         file_app_rds_pb_rds_proto_enumTypes,
		MessageInfos:      file_app_rds_pb_rds_proto_msgTypes,
	}.Build()
	File_app_rds_pb_rds_proto = out.File
	file_app_rds_pb_rds_proto_rawDesc = nil
	file_app_rds_pb_rds_proto_goTypes = nil
	file_app_rds_pb_rds_proto_depIdxs = nil
}