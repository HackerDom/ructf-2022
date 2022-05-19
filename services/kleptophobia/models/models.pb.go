// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/models.proto

package models

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

type GetPublicInfoReply_Status int32

const (
	GetPublicInfoReply_OK   GetPublicInfoReply_Status = 0
	GetPublicInfoReply_FAIL GetPublicInfoReply_Status = 1
)

// Enum value maps for GetPublicInfoReply_Status.
var (
	GetPublicInfoReply_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	GetPublicInfoReply_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x GetPublicInfoReply_Status) Enum() *GetPublicInfoReply_Status {
	p := new(GetPublicInfoReply_Status)
	*p = x
	return p
}

func (x GetPublicInfoReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPublicInfoReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_models_models_proto_enumTypes[0].Descriptor()
}

func (GetPublicInfoReply_Status) Type() protoreflect.EnumType {
	return &file_models_models_proto_enumTypes[0]
}

func (x GetPublicInfoReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPublicInfoReply_Status.Descriptor instead.
func (GetPublicInfoReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{1, 0}
}

type GetEncryptedFullInfoReply_Status int32

const (
	GetEncryptedFullInfoReply_OK   GetEncryptedFullInfoReply_Status = 0
	GetEncryptedFullInfoReply_FAIL GetEncryptedFullInfoReply_Status = 1
)

// Enum value maps for GetEncryptedFullInfoReply_Status.
var (
	GetEncryptedFullInfoReply_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	GetEncryptedFullInfoReply_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x GetEncryptedFullInfoReply_Status) Enum() *GetEncryptedFullInfoReply_Status {
	p := new(GetEncryptedFullInfoReply_Status)
	*p = x
	return p
}

func (x GetEncryptedFullInfoReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetEncryptedFullInfoReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_models_models_proto_enumTypes[1].Descriptor()
}

func (GetEncryptedFullInfoReply_Status) Type() protoreflect.EnumType {
	return &file_models_models_proto_enumTypes[1]
}

func (x GetEncryptedFullInfoReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetEncryptedFullInfoReply_Status.Descriptor instead.
func (GetEncryptedFullInfoReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{2, 0}
}

type RegisterReply_Status int32

const (
	RegisterReply_OK   RegisterReply_Status = 0
	RegisterReply_FAIL RegisterReply_Status = 1
)

// Enum value maps for RegisterReply_Status.
var (
	RegisterReply_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	RegisterReply_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x RegisterReply_Status) Enum() *RegisterReply_Status {
	p := new(RegisterReply_Status)
	*p = x
	return p
}

func (x RegisterReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_models_models_proto_enumTypes[2].Descriptor()
}

func (RegisterReply_Status) Type() protoreflect.EnumType {
	return &file_models_models_proto_enumTypes[2]
}

func (x RegisterReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterReply_Status.Descriptor instead.
func (RegisterReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{4, 0}
}

type GetByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetByUsernameRequest) Reset() {
	*x = GetByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUsernameRequest) ProtoMessage() {}

func (x *GetByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{0}
}

func (x *GetByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPublicInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetPublicInfoReply_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.GetPublicInfoReply_Status" json:"status,omitempty"`
	Message *string                   `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Person  *PublicPerson             `protobuf:"bytes,3,opt,name=person,proto3,oneof" json:"person,omitempty"`
}

func (x *GetPublicInfoReply) Reset() {
	*x = GetPublicInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicInfoReply) ProtoMessage() {}

func (x *GetPublicInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicInfoReply.ProtoReflect.Descriptor instead.
func (*GetPublicInfoReply) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{1}
}

func (x *GetPublicInfoReply) GetStatus() GetPublicInfoReply_Status {
	if x != nil {
		return x.Status
	}
	return GetPublicInfoReply_OK
}

func (x *GetPublicInfoReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetPublicInfoReply) GetPerson() *PublicPerson {
	if x != nil {
		return x.Person
	}
	return nil
}

type GetEncryptedFullInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            GetEncryptedFullInfoReply_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.GetEncryptedFullInfoReply_Status" json:"status,omitempty"`
	Message           *string                          `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	EncryptedFullInfo []byte                           `protobuf:"bytes,3,opt,name=encryptedFullInfo,proto3,oneof" json:"encryptedFullInfo,omitempty"`
}

func (x *GetEncryptedFullInfoReply) Reset() {
	*x = GetEncryptedFullInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEncryptedFullInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEncryptedFullInfoReply) ProtoMessage() {}

func (x *GetEncryptedFullInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEncryptedFullInfoReply.ProtoReflect.Descriptor instead.
func (*GetEncryptedFullInfoReply) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{2}
}

func (x *GetEncryptedFullInfoReply) GetStatus() GetEncryptedFullInfoReply_Status {
	if x != nil {
		return x.Status
	}
	return GetEncryptedFullInfoReply_OK
}

func (x *GetEncryptedFullInfoReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetEncryptedFullInfoReply) GetEncryptedFullInfo() []byte {
	if x != nil {
		return x.EncryptedFullInfo
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person   *PrivatePerson `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	Password string         `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetPerson() *PrivatePerson {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  RegisterReply_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.RegisterReply_Status" json:"status,omitempty"`
	Message *string              `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterReply) GetStatus() RegisterReply_Status {
	if x != nil {
		return x.Status
	}
	return RegisterReply_OK
}

func (x *RegisterReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type PrivatePerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string  `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName *string `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	SecondName string  `protobuf:"bytes,3,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Username   string  `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Room       int32   `protobuf:"varint,5,opt,name=room,proto3" json:"room,omitempty"`
	Diagnosis  string  `protobuf:"bytes,6,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
}

func (x *PrivatePerson) Reset() {
	*x = PrivatePerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivatePerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivatePerson) ProtoMessage() {}

func (x *PrivatePerson) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivatePerson.ProtoReflect.Descriptor instead.
func (*PrivatePerson) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{5}
}

func (x *PrivatePerson) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PrivatePerson) GetMiddleName() string {
	if x != nil && x.MiddleName != nil {
		return *x.MiddleName
	}
	return ""
}

func (x *PrivatePerson) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *PrivatePerson) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PrivatePerson) GetRoom() int32 {
	if x != nil {
		return x.Room
	}
	return 0
}

func (x *PrivatePerson) GetDiagnosis() string {
	if x != nil {
		return x.Diagnosis
	}
	return ""
}

type PublicPerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName string `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Username   string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Room       int32  `protobuf:"varint,4,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *PublicPerson) Reset() {
	*x = PublicPerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicPerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicPerson) ProtoMessage() {}

func (x *PublicPerson) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicPerson.ProtoReflect.Descriptor instead.
func (*PublicPerson) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{6}
}

func (x *PublicPerson) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PublicPerson) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *PublicPerson) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PublicPerson) GetRoom() int32 {
	if x != nil {
		return x.Room
	}
	return 0
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpcHost string `protobuf:"bytes,1,opt,name=grpc_host,json=grpcHost,proto3" json:"grpc_host,omitempty"`
	GrpcPort int32  `protobuf:"varint,2,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{7}
}

func (x *ClientConfig) GetGrpcHost() string {
	if x != nil {
		return x.GrpcHost
	}
	return ""
}

func (x *ClientConfig) GetGrpcPort() int32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

type PGConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	DbName   string `protobuf:"bytes,5,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *PGConfig) Reset() {
	*x = PGConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PGConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PGConfig) ProtoMessage() {}

func (x *PGConfig) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PGConfig.ProtoReflect.Descriptor instead.
func (*PGConfig) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{8}
}

func (x *PGConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PGConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PGConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PGConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PGConfig) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpcPort int32     `protobuf:"varint,1,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	PgConfig *PGConfig `protobuf:"bytes,2,opt,name=pg_config,json=pgConfig,proto3" json:"pg_config,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_models_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{9}
}

func (x *ServerConfig) GetGrpcPort() int32 {
	if x != nil {
		return x.GrpcPort
	}
	return 0
}

func (x *ServerConfig) GetPgConfig() *PGConfig {
	if x != nil {
		return x.PgConfig
	}
	return nil
}

var File_models_models_proto protoreflect.FileDescriptor

var file_models_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x32, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xed, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x01, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x22, 0x1a, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x0c, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x48, 0x0a, 0x0c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x72, 0x70, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x70,
	0x63, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x50, 0x47, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x67, 0x72, 0x70, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x67, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x47, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x70,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0xf4, 0x01, 0x0a, 0x0c, 0x4b, 0x6c, 0x65, 0x70,
	0x74, 0x6f, 0x70, 0x68, 0x6f, 0x62, 0x69, 0x61, 0x12, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x68, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x64, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_models_proto_rawDescOnce sync.Once
	file_models_models_proto_rawDescData = file_models_models_proto_rawDesc
)

func file_models_models_proto_rawDescGZIP() []byte {
	file_models_models_proto_rawDescOnce.Do(func() {
		file_models_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_models_proto_rawDescData)
	})
	return file_models_models_proto_rawDescData
}

var file_models_models_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_models_models_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_models_models_proto_goTypes = []interface{}{
	(GetPublicInfoReply_Status)(0),        // 0: models.GetPublicInfoReply.Status
	(GetEncryptedFullInfoReply_Status)(0), // 1: models.GetEncryptedFullInfoReply.Status
	(RegisterReply_Status)(0),             // 2: models.RegisterReply.Status
	(*GetByUsernameRequest)(nil),          // 3: models.GetByUsernameRequest
	(*GetPublicInfoReply)(nil),            // 4: models.GetPublicInfoReply
	(*GetEncryptedFullInfoReply)(nil),     // 5: models.GetEncryptedFullInfoReply
	(*RegisterRequest)(nil),               // 6: models.RegisterRequest
	(*RegisterReply)(nil),                 // 7: models.RegisterReply
	(*PrivatePerson)(nil),                 // 8: models.PrivatePerson
	(*PublicPerson)(nil),                  // 9: models.PublicPerson
	(*ClientConfig)(nil),                  // 10: models.ClientConfig
	(*PGConfig)(nil),                      // 11: models.PGConfig
	(*ServerConfig)(nil),                  // 12: models.ServerConfig
}
var file_models_models_proto_depIdxs = []int32{
	0,  // 0: models.GetPublicInfoReply.status:type_name -> models.GetPublicInfoReply.Status
	9,  // 1: models.GetPublicInfoReply.person:type_name -> models.PublicPerson
	1,  // 2: models.GetEncryptedFullInfoReply.status:type_name -> models.GetEncryptedFullInfoReply.Status
	8,  // 3: models.RegisterRequest.person:type_name -> models.PrivatePerson
	2,  // 4: models.RegisterReply.status:type_name -> models.RegisterReply.Status
	11, // 5: models.ServerConfig.pg_config:type_name -> models.PGConfig
	6,  // 6: models.Kleptophobia.Register:input_type -> models.RegisterRequest
	3,  // 7: models.Kleptophobia.GetPublicInfo:input_type -> models.GetByUsernameRequest
	3,  // 8: models.Kleptophobia.GetEncryptedFullInfo:input_type -> models.GetByUsernameRequest
	7,  // 9: models.Kleptophobia.Register:output_type -> models.RegisterReply
	4,  // 10: models.Kleptophobia.GetPublicInfo:output_type -> models.GetPublicInfoReply
	5,  // 11: models.Kleptophobia.GetEncryptedFullInfo:output_type -> models.GetEncryptedFullInfoReply
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_models_models_proto_init() }
func file_models_models_proto_init() {
	if File_models_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUsernameRequest); i {
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
		file_models_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicInfoReply); i {
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
		file_models_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEncryptedFullInfoReply); i {
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
		file_models_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_models_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_models_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivatePerson); i {
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
		file_models_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicPerson); i {
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
		file_models_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_models_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PGConfig); i {
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
		file_models_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
	file_models_models_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_models_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_models_models_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_models_models_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_models_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_models_proto_goTypes,
		DependencyIndexes: file_models_models_proto_depIdxs,
		EnumInfos:         file_models_models_proto_enumTypes,
		MessageInfos:      file_models_models_proto_msgTypes,
	}.Build()
	File_models_models_proto = out.File
	file_models_models_proto_rawDesc = nil
	file_models_models_proto_goTypes = nil
	file_models_models_proto_depIdxs = nil
}
