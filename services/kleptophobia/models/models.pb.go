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
	return file_models_models_proto_enumTypes[0].Descriptor()
}

func (RegisterReply_Status) Type() protoreflect.EnumType {
	return &file_models_models_proto_enumTypes[0]
}

func (x RegisterReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterReply_Status.Descriptor instead.
func (RegisterReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{1, 0}
}

// The request message containing the user's name.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *PrivatePerson `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetPerson() *PrivatePerson {
	if x != nil {
		return x.Person
	}
	return nil
}

// The response message containing the greetings
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
		mi := &file_models_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{1}
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

	FirstName          string  `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName         *string `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	SecondName         string  `protobuf:"bytes,3,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Username           string  `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password           string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	EncryptedDiagnosis string  `protobuf:"bytes,6,opt,name=encrypted_diagnosis,json=encryptedDiagnosis,proto3" json:"encrypted_diagnosis,omitempty"`
}

func (x *PrivatePerson) Reset() {
	*x = PrivatePerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivatePerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivatePerson) ProtoMessage() {}

func (x *PrivatePerson) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PrivatePerson.ProtoReflect.Descriptor instead.
func (*PrivatePerson) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{2}
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

func (x *PrivatePerson) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PrivatePerson) GetEncryptedDiagnosis() string {
	if x != nil {
		return x.EncryptedDiagnosis
	}
	return ""
}

type PublicPerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName          string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName         string `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	EncryptedDiagnosis string `protobuf:"bytes,3,opt,name=encrypted_diagnosis,json=encryptedDiagnosis,proto3" json:"encrypted_diagnosis,omitempty"`
}

func (x *PublicPerson) Reset() {
	*x = PublicPerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicPerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicPerson) ProtoMessage() {}

func (x *PublicPerson) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublicPerson.ProtoReflect.Descriptor instead.
func (*PublicPerson) Descriptor() ([]byte, []int) {
	return file_models_models_proto_rawDescGZIP(), []int{3}
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

func (x *PublicPerson) GetEncryptedDiagnosis() string {
	if x != nil {
		return x.EncryptedDiagnosis
	}
	return ""
}

var File_models_models_proto protoreflect.FileDescriptor

var file_models_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x40, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22,
	0x8c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c,
	0x10, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xee,
	0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2f,
	0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x69, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x7f, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73,
	0x32, 0x4c, 0x0a, 0x0c, 0x4b, 0x6c, 0x65, 0x70, 0x74, 0x6f, 0x70, 0x68, 0x6f, 0x62, 0x69, 0x61,
	0x12, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x16,
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

var file_models_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_models_proto_goTypes = []interface{}{
	(RegisterReply_Status)(0), // 0: models.RegisterReply.Status
	(*RegisterRequest)(nil),   // 1: models.RegisterRequest
	(*RegisterReply)(nil),     // 2: models.RegisterReply
	(*PrivatePerson)(nil),     // 3: models.PrivatePerson
	(*PublicPerson)(nil),      // 4: models.PublicPerson
}
var file_models_models_proto_depIdxs = []int32{
	3, // 0: models.RegisterRequest.person:type_name -> models.PrivatePerson
	0, // 1: models.RegisterReply.status:type_name -> models.RegisterReply.Status
	1, // 2: models.Kleptophobia.Register:input_type -> models.RegisterRequest
	2, // 3: models.Kleptophobia.Register:output_type -> models.RegisterReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_models_proto_init() }
func file_models_models_proto_init() {
	if File_models_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_models_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_models_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_models_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	file_models_models_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_models_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
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
