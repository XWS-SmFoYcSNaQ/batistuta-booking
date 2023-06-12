// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: auth/auth-service.proto

package auth

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserRole int32

const (
	UserRole_Host  UserRole = 0
	UserRole_Guest UserRole = 1
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "Host",
		1: "Guest",
	}
	UserRole_value = map[string]int32{
		"Host":  0,
		"Guest": 1,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_auth_service_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_auth_auth_service_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{0}
}

type Empty_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty_Request) Reset() {
	*x = Empty_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty_Request) ProtoMessage() {}

func (x *Empty_Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty_Request.ProtoReflect.Descriptor instead.
func (*Empty_Request) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{0}
}

type Verify_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verified     bool     `protobuf:"varint,1,opt,name=Verified,proto3" json:"Verified,omitempty"`
	ErrorMessage string   `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	UserId       string   `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserRole     UserRole `protobuf:"varint,4,opt,name=UserRole,proto3,enum=auth.UserRole" json:"UserRole,omitempty"`
}

func (x *Verify_Response) Reset() {
	*x = Verify_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verify_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verify_Response) ProtoMessage() {}

func (x *Verify_Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verify_Response.ProtoReflect.Descriptor instead.
func (*Verify_Response) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *Verify_Response) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Verify_Response) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Verify_Response) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Verify_Response) GetUserRole() UserRole {
	if x != nil {
		return x.UserRole
	}
	return UserRole_Host
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Role        UserRole `protobuf:"varint,2,opt,name=Role,proto3,enum=auth.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName   string   `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,7,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[2]
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
	return file_auth_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyName string `protobuf:"bytes,1,opt,name=PropertyName,proto3" json:"PropertyName,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Authentication_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Authentication_Request) Reset() {
	*x = Authentication_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication_Request) ProtoMessage() {}

func (x *Authentication_Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication_Request.ProtoReflect.Descriptor instead.
func (*Authentication_Request) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *Authentication_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Authentication_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Authentication_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Success      bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	User         *User  `protobuf:"bytes,4,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *Authentication_Response) Reset() {
	*x = Authentication_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication_Response) ProtoMessage() {}

func (x *Authentication_Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication_Response.ProtoReflect.Descriptor instead.
func (*Authentication_Response) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *Authentication_Response) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Authentication_Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Authentication_Response) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Authentication_Response) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Register_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        UserRole `protobuf:"varint,1,opt,name=Role,proto3,enum=auth.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	FirstName   string   `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,7,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *Register_Request) Reset() {
	*x = Register_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register_Request) ProtoMessage() {}

func (x *Register_Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register_Request.ProtoReflect.Descriptor instead.
func (*Register_Request) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *Register_Request) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *Register_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Register_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Register_Request) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Register_Request) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Register_Request) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Register_Request) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

type Register_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Token   string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Errors  []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	Message string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	User    *User    `protobuf:"bytes,5,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *Register_Response) Reset() {
	*x = Register_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register_Response) ProtoMessage() {}

func (x *Register_Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register_Response.ProtoReflect.Descriptor instead.
func (*Register_Response) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *Register_Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Register_Response) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Register_Response) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Register_Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Register_Response) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_auth_auth_service_proto protoreflect.FileDescriptor

var file_auth_auth_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x95,
	0x01, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x22, 0x4f, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x2a, 0x1f, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x74, 0x10, 0x01, 0x32, 0x9e, 0x02,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x12, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x0c,
	0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_service_proto_rawDescOnce sync.Once
	file_auth_auth_service_proto_rawDescData = file_auth_auth_service_proto_rawDesc
)

func file_auth_auth_service_proto_rawDescGZIP() []byte {
	file_auth_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_service_proto_rawDescData)
	})
	return file_auth_auth_service_proto_rawDescData
}

var file_auth_auth_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_auth_service_proto_goTypes = []interface{}{
	(UserRole)(0),                   // 0: auth.UserRole
	(*Empty_Request)(nil),           // 1: auth.Empty_Request
	(*Verify_Response)(nil),         // 2: auth.Verify_Response
	(*User)(nil),                    // 3: auth.User
	(*Error)(nil),                   // 4: auth.Error
	(*Authentication_Request)(nil),  // 5: auth.Authentication_Request
	(*Authentication_Response)(nil), // 6: auth.Authentication_Response
	(*Register_Request)(nil),        // 7: auth.Register_Request
	(*Register_Response)(nil),       // 8: auth.Register_Response
}
var file_auth_auth_service_proto_depIdxs = []int32{
	0, // 0: auth.Verify_Response.UserRole:type_name -> auth.UserRole
	0, // 1: auth.User.Role:type_name -> auth.UserRole
	3, // 2: auth.Authentication_Response.User:type_name -> auth.User
	0, // 3: auth.Register_Request.Role:type_name -> auth.UserRole
	4, // 4: auth.Register_Response.errors:type_name -> auth.Error
	3, // 5: auth.Register_Response.User:type_name -> auth.User
	1, // 6: auth.AuthService.Verify:input_type -> auth.Empty_Request
	7, // 7: auth.AuthService.Register:input_type -> auth.Register_Request
	5, // 8: auth.AuthService.Login:input_type -> auth.Authentication_Request
	2, // 9: auth.AuthService.Verify:output_type -> auth.Verify_Response
	8, // 10: auth.AuthService.Register:output_type -> auth.Register_Response
	6, // 11: auth.AuthService.Login:output_type -> auth.Authentication_Response
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_auth_auth_service_proto_init() }
func file_auth_auth_service_proto_init() {
	if File_auth_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty_Request); i {
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
		file_auth_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verify_Response); i {
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
		file_auth_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_auth_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_auth_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication_Request); i {
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
		file_auth_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication_Response); i {
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
		file_auth_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register_Request); i {
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
		file_auth_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register_Response); i {
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
			RawDescriptor: file_auth_auth_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_auth_service_proto_depIdxs,
		EnumInfos:         file_auth_auth_service_proto_enumTypes,
		MessageInfos:      file_auth_auth_service_proto_msgTypes,
	}.Build()
	File_auth_auth_service_proto = out.File
	file_auth_auth_service_proto_rawDesc = nil
	file_auth_auth_service_proto_goTypes = nil
	file_auth_auth_service_proto_depIdxs = nil
}
