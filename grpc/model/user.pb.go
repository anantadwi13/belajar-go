// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: grpc/model/user.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User_UserType int32

const (
	User_ADMIN User_UserType = 0
	User_USER  User_UserType = 1
)

// Enum value maps for User_UserType.
var (
	User_UserType_name = map[int32]string{
		0: "ADMIN",
		1: "USER",
	}
	User_UserType_value = map[string]int32{
		"ADMIN": 0,
		"USER":  1,
	}
)

func (x User_UserType) Enum() *User_UserType {
	p := new(User_UserType)
	*p = x
	return p
}

func (x User_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_model_user_proto_enumTypes[0].Descriptor()
}

func (User_UserType) Type() protoreflect.EnumType {
	return &file_grpc_model_user_proto_enumTypes[0]
}

func (x User_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_UserType.Descriptor instead.
func (User_UserType) EnumDescriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{0, 0}
}

type User_PhoneType int32

const (
	User_WORK User_PhoneType = 0
	User_HOME User_PhoneType = 1
)

// Enum value maps for User_PhoneType.
var (
	User_PhoneType_name = map[int32]string{
		0: "WORK",
		1: "HOME",
	}
	User_PhoneType_value = map[string]int32{
		"WORK": 0,
		"HOME": 1,
	}
)

func (x User_PhoneType) Enum() *User_PhoneType {
	p := new(User_PhoneType)
	*p = x
	return p
}

func (x User_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_model_user_proto_enumTypes[1].Descriptor()
}

func (User_PhoneType) Type() protoreflect.EnumType {
	return &file_grpc_model_user_proto_enumTypes[1]
}

func (x User_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_PhoneType.Descriptor instead.
func (User_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{0, 1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones      []*User_PhoneNumber    `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	UserType    User_UserType          `protobuf:"varint,5,opt,name=user_type,json=userType,proto3,enum=model.User_UserType" json:"user_type,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhones() []*User_PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *User) GetUserType() User_UserType {
	if x != nil {
		return x.UserType
	}
	return User_ADMIN
}

func (x *User) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type GetUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserByIdReq) Reset() {
	*x = GetUserByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdReq) ProtoMessage() {}

func (x *GetUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{1}
}

type GetUserByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUserByIdRes) Reset() {
	*x = GetUserByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRes) ProtoMessage() {}

func (x *GetUserByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRes.ProtoReflect.Descriptor instead.
func (*GetUserByIdRes) Descriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdRes) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string         `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   User_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=model.User_PhoneType" json:"type,omitempty"`
}

func (x *User_PhoneNumber) Reset() {
	*x = User_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_PhoneNumber) ProtoMessage() {}

func (x *User_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_PhoneNumber.ProtoReflect.Descriptor instead.
func (*User_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_grpc_model_user_proto_rawDescGZIP(), []int{0, 0}
}

func (x *User_PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *User_PhoneNumber) GetType() User_PhoneType {
	if x != nil {
		return x.Type
	}
	return User_WORK
}

var File_grpc_model_user_proto protoreflect.FileDescriptor

var file_grpc_model_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x50, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x22, 0x1f, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x33, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x32, 0x47, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x61, 0x64, 0x77,
	0x69, 0x31, 0x33, 0x2f, 0x62, 0x65, 0x6c, 0x61, 0x6a, 0x61, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_model_user_proto_rawDescOnce sync.Once
	file_grpc_model_user_proto_rawDescData = file_grpc_model_user_proto_rawDesc
)

func file_grpc_model_user_proto_rawDescGZIP() []byte {
	file_grpc_model_user_proto_rawDescOnce.Do(func() {
		file_grpc_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_model_user_proto_rawDescData)
	})
	return file_grpc_model_user_proto_rawDescData
}

var file_grpc_model_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_grpc_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_model_user_proto_goTypes = []interface{}{
	(User_UserType)(0),            // 0: model.User.UserType
	(User_PhoneType)(0),           // 1: model.User.PhoneType
	(*User)(nil),                  // 2: model.User
	(*GetUserByIdReq)(nil),        // 3: model.GetUserByIdReq
	(*GetUserByIdRes)(nil),        // 4: model.GetUserByIdRes
	(*User_PhoneNumber)(nil),      // 5: model.User.PhoneNumber
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_grpc_model_user_proto_depIdxs = []int32{
	5, // 0: model.User.phones:type_name -> model.User.PhoneNumber
	0, // 1: model.User.user_type:type_name -> model.User.UserType
	6, // 2: model.User.last_updated:type_name -> google.protobuf.Timestamp
	2, // 3: model.GetUserByIdRes.users:type_name -> model.User
	1, // 4: model.User.PhoneNumber.type:type_name -> model.User.PhoneType
	3, // 5: model.UserService.GetAll:input_type -> model.GetUserByIdReq
	4, // 6: model.UserService.GetAll:output_type -> model.GetUserByIdRes
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_model_user_proto_init() }
func file_grpc_model_user_proto_init() {
	if File_grpc_model_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_model_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_grpc_model_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdReq); i {
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
		file_grpc_model_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdRes); i {
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
		file_grpc_model_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_PhoneNumber); i {
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
			RawDescriptor: file_grpc_model_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_model_user_proto_goTypes,
		DependencyIndexes: file_grpc_model_user_proto_depIdxs,
		EnumInfos:         file_grpc_model_user_proto_enumTypes,
		MessageInfos:      file_grpc_model_user_proto_msgTypes,
	}.Build()
	File_grpc_model_user_proto = out.File
	file_grpc_model_user_proto_rawDesc = nil
	file_grpc_model_user_proto_goTypes = nil
	file_grpc_model_user_proto_depIdxs = nil
}
