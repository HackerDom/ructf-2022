// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: proto/models.proto

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

type GetPublicInfoRsp_Status int32

const (
	GetPublicInfoRsp_OK   GetPublicInfoRsp_Status = 0
	GetPublicInfoRsp_FAIL GetPublicInfoRsp_Status = 1
)

// Enum value maps for GetPublicInfoRsp_Status.
var (
	GetPublicInfoRsp_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	GetPublicInfoRsp_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x GetPublicInfoRsp_Status) Enum() *GetPublicInfoRsp_Status {
	p := new(GetPublicInfoRsp_Status)
	*p = x
	return p
}

func (x GetPublicInfoRsp_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPublicInfoRsp_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_models_proto_enumTypes[0].Descriptor()
}

func (GetPublicInfoRsp_Status) Type() protoreflect.EnumType {
	return &file_proto_models_proto_enumTypes[0]
}

func (x GetPublicInfoRsp_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPublicInfoRsp_Status.Descriptor instead.
func (GetPublicInfoRsp_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{2, 0}
}

type GetEncryptedFullInfoRsp_Status int32

const (
	GetEncryptedFullInfoRsp_OK   GetEncryptedFullInfoRsp_Status = 0
	GetEncryptedFullInfoRsp_FAIL GetEncryptedFullInfoRsp_Status = 1
)

// Enum value maps for GetEncryptedFullInfoRsp_Status.
var (
	GetEncryptedFullInfoRsp_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	GetEncryptedFullInfoRsp_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x GetEncryptedFullInfoRsp_Status) Enum() *GetEncryptedFullInfoRsp_Status {
	p := new(GetEncryptedFullInfoRsp_Status)
	*p = x
	return p
}

func (x GetEncryptedFullInfoRsp_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetEncryptedFullInfoRsp_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_models_proto_enumTypes[1].Descriptor()
}

func (GetEncryptedFullInfoRsp_Status) Type() protoreflect.EnumType {
	return &file_proto_models_proto_enumTypes[1]
}

func (x GetEncryptedFullInfoRsp_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetEncryptedFullInfoRsp_Status.Descriptor instead.
func (GetEncryptedFullInfoRsp_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{3, 0}
}

type RegisterRsp_Status int32

const (
	RegisterRsp_OK   RegisterRsp_Status = 0
	RegisterRsp_FAIL RegisterRsp_Status = 1
)

// Enum value maps for RegisterRsp_Status.
var (
	RegisterRsp_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	RegisterRsp_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x RegisterRsp_Status) Enum() *RegisterRsp_Status {
	p := new(RegisterRsp_Status)
	*p = x
	return p
}

func (x RegisterRsp_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterRsp_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_models_proto_enumTypes[2].Descriptor()
}

func (RegisterRsp_Status) Type() protoreflect.EnumType {
	return &file_proto_models_proto_enumTypes[2]
}

func (x RegisterRsp_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterRsp_Status.Descriptor instead.
func (RegisterRsp_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{5, 0}
}

type PingBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingBody) Reset() {
	*x = PingBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingBody) ProtoMessage() {}

func (x *PingBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingBody.ProtoReflect.Descriptor instead.
func (*PingBody) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{0}
}

func (x *PingBody) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetByUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetByUsernameReq) Reset() {
	*x = GetByUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUsernameReq) ProtoMessage() {}

func (x *GetByUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUsernameReq.ProtoReflect.Descriptor instead.
func (*GetByUsernameReq) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{1}
}

func (x *GetByUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPublicInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  GetPublicInfoRsp_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.GetPublicInfoRsp_Status" json:"status,omitempty"`
	Message *string                 `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Person  *PublicPerson           `protobuf:"bytes,3,opt,name=person,proto3,oneof" json:"person,omitempty"`
}

func (x *GetPublicInfoRsp) Reset() {
	*x = GetPublicInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicInfoRsp) ProtoMessage() {}

func (x *GetPublicInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicInfoRsp.ProtoReflect.Descriptor instead.
func (*GetPublicInfoRsp) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{2}
}

func (x *GetPublicInfoRsp) GetStatus() GetPublicInfoRsp_Status {
	if x != nil {
		return x.Status
	}
	return GetPublicInfoRsp_OK
}

func (x *GetPublicInfoRsp) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetPublicInfoRsp) GetPerson() *PublicPerson {
	if x != nil {
		return x.Person
	}
	return nil
}

type GetEncryptedFullInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            GetEncryptedFullInfoRsp_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.GetEncryptedFullInfoRsp_Status" json:"status,omitempty"`
	Message           *string                        `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	EncryptedFullInfo []byte                         `protobuf:"bytes,3,opt,name=encryptedFullInfo,proto3,oneof" json:"encryptedFullInfo,omitempty"`
}

func (x *GetEncryptedFullInfoRsp) Reset() {
	*x = GetEncryptedFullInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEncryptedFullInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEncryptedFullInfoRsp) ProtoMessage() {}

func (x *GetEncryptedFullInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEncryptedFullInfoRsp.ProtoReflect.Descriptor instead.
func (*GetEncryptedFullInfoRsp) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{3}
}

func (x *GetEncryptedFullInfoRsp) GetStatus() GetEncryptedFullInfoRsp_Status {
	if x != nil {
		return x.Status
	}
	return GetEncryptedFullInfoRsp_OK
}

func (x *GetEncryptedFullInfoRsp) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetEncryptedFullInfoRsp) GetEncryptedFullInfo() []byte {
	if x != nil {
		return x.EncryptedFullInfo
	}
	return nil
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person   *PrivatePerson `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	Password string         `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterReq) GetPerson() *PrivatePerson {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  RegisterRsp_Status `protobuf:"varint,1,opt,name=status,proto3,enum=models.RegisterRsp_Status" json:"status,omitempty"`
	Message *string            `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *RegisterRsp) Reset() {
	*x = RegisterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRsp) ProtoMessage() {}

func (x *RegisterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRsp.ProtoReflect.Descriptor instead.
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return file_proto_models_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterRsp) GetStatus() RegisterRsp_Status {
	if x != nil {
		return x.Status
	}
	return RegisterRsp_OK
}

func (x *RegisterRsp) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type PrivatePerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	SecondName string `protobuf:"bytes,3,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Username   string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Room       uint32 `protobuf:"varint,5,opt,name=room,proto3" json:"room,omitempty"`
	Diagnosis  string `protobuf:"bytes,6,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
}

func (x *PrivatePerson) Reset() {
	*x = PrivatePerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivatePerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivatePerson) ProtoMessage() {}

func (x *PrivatePerson) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[6]
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
	return file_proto_models_proto_rawDescGZIP(), []int{6}
}

func (x *PrivatePerson) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PrivatePerson) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
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

func (x *PrivatePerson) GetRoom() uint32 {
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
	Room       uint32 `protobuf:"varint,4,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *PublicPerson) Reset() {
	*x = PublicPerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicPerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicPerson) ProtoMessage() {}

func (x *PublicPerson) ProtoReflect() protoreflect.Message {
	mi := &file_proto_models_proto_msgTypes[7]
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
	return file_proto_models_proto_rawDescGZIP(), []int{7}
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

func (x *PublicPerson) GetRoom() uint32 {
	if x != nil {
		return x.Room
	}
	return 0
}

var File_proto_models_proto protoreflect.FileDescriptor

var file_proto_models_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x24, 0x0a, 0x08,
	0x50, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73,
	0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x31, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73,
	0x70, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x31, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x11, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x88, 0x01, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x2d, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x22, 0x1a,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x22, 0x7e, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_models_proto_rawDescOnce sync.Once
	file_proto_models_proto_rawDescData = file_proto_models_proto_rawDesc
)

func file_proto_models_proto_rawDescGZIP() []byte {
	file_proto_models_proto_rawDescOnce.Do(func() {
		file_proto_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_models_proto_rawDescData)
	})
	return file_proto_models_proto_rawDescData
}

var file_proto_models_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_models_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_models_proto_goTypes = []interface{}{
	(GetPublicInfoRsp_Status)(0),        // 0: models.GetPublicInfoRsp.Status
	(GetEncryptedFullInfoRsp_Status)(0), // 1: models.GetEncryptedFullInfoRsp.Status
	(RegisterRsp_Status)(0),             // 2: models.RegisterRsp.Status
	(*PingBody)(nil),                    // 3: models.PingBody
	(*GetByUsernameReq)(nil),            // 4: models.GetByUsernameReq
	(*GetPublicInfoRsp)(nil),            // 5: models.GetPublicInfoRsp
	(*GetEncryptedFullInfoRsp)(nil),     // 6: models.GetEncryptedFullInfoRsp
	(*RegisterReq)(nil),                 // 7: models.RegisterReq
	(*RegisterRsp)(nil),                 // 8: models.RegisterRsp
	(*PrivatePerson)(nil),               // 9: models.PrivatePerson
	(*PublicPerson)(nil),                // 10: models.PublicPerson
}
var file_proto_models_proto_depIdxs = []int32{
	0,  // 0: models.GetPublicInfoRsp.status:type_name -> models.GetPublicInfoRsp.Status
	10, // 1: models.GetPublicInfoRsp.person:type_name -> models.PublicPerson
	1,  // 2: models.GetEncryptedFullInfoRsp.status:type_name -> models.GetEncryptedFullInfoRsp.Status
	9,  // 3: models.RegisterReq.person:type_name -> models.PrivatePerson
	2,  // 4: models.RegisterRsp.status:type_name -> models.RegisterRsp.Status
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_models_proto_init() }
func file_proto_models_proto_init() {
	if File_proto_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingBody); i {
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
		file_proto_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUsernameReq); i {
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
		file_proto_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicInfoRsp); i {
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
		file_proto_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEncryptedFullInfoRsp); i {
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
		file_proto_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_proto_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRsp); i {
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
		file_proto_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_proto_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_models_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_proto_models_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_models_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_models_proto_goTypes,
		DependencyIndexes: file_proto_models_proto_depIdxs,
		EnumInfos:         file_proto_models_proto_enumTypes,
		MessageInfos:      file_proto_models_proto_msgTypes,
	}.Build()
	File_proto_models_proto = out.File
	file_proto_models_proto_rawDesc = nil
	file_proto_models_proto_goTypes = nil
	file_proto_models_proto_depIdxs = nil
}
