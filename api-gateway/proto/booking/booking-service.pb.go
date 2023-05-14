// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: booking/booking-service.proto

package booking

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

type AM_GetAllBookingRequests_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AM_GetAllBookingRequests_Request) Reset() {
	*x = AM_GetAllBookingRequests_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequests_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequests_Request) ProtoMessage() {}

func (x *AM_GetAllBookingRequests_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequests_Request.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequests_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{0}
}

type AM_GetAllBookingRequests_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BookingRequestsDTO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AM_GetAllBookingRequests_Response) Reset() {
	*x = AM_GetAllBookingRequests_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequests_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequests_Response) ProtoMessage() {}

func (x *AM_GetAllBookingRequests_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequests_Response.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequests_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{1}
}

func (x *AM_GetAllBookingRequests_Response) GetData() []*BookingRequestsDTO {
	if x != nil {
		return x.Data
	}
	return nil
}

type AM_BookingRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,1,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	StartDate       string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	NumberOfGuests  int32  `protobuf:"varint,4,opt,name=numberOfGuests,proto3" json:"numberOfGuests,omitempty"`
	UserId          string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AM_BookingRequest_Request) Reset() {
	*x = AM_BookingRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_BookingRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_BookingRequest_Request) ProtoMessage() {}

func (x *AM_BookingRequest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_BookingRequest_Request.ProtoReflect.Descriptor instead.
func (*AM_BookingRequest_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{2}
}

func (x *AM_BookingRequest_Request) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

func (x *AM_BookingRequest_Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AM_CreateBookingRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_CreateBookingRequest_Response) Reset() {
	*x = AM_CreateBookingRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_CreateBookingRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_CreateBookingRequest_Response) ProtoMessage() {}

func (x *AM_CreateBookingRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_CreateBookingRequest_Response.ProtoReflect.Descriptor instead.
func (*AM_CreateBookingRequest_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{3}
}

func (x *AM_CreateBookingRequest_Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BookingRequestsDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccommodationId string `protobuf:"bytes,2,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	StartDate       string `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	NumberOfGuests  int32  `protobuf:"varint,5,opt,name=numberOfGuests,proto3" json:"numberOfGuests,omitempty"`
	UserId          string `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *BookingRequestsDTO) Reset() {
	*x = BookingRequestsDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequestsDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequestsDTO) ProtoMessage() {}

func (x *BookingRequestsDTO) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequestsDTO.ProtoReflect.Descriptor instead.
func (*BookingRequestsDTO) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{4}
}

func (x *BookingRequestsDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingRequestsDTO) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *BookingRequestsDTO) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BookingRequestsDTO) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *BookingRequestsDTO) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

func (x *BookingRequestsDTO) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AM_DeleteBookingRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_DeleteBookingRequest_Request) Reset() {
	*x = AM_DeleteBookingRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_DeleteBookingRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_DeleteBookingRequest_Request) ProtoMessage() {}

func (x *AM_DeleteBookingRequest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_DeleteBookingRequest_Request.ProtoReflect.Descriptor instead.
func (*AM_DeleteBookingRequest_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{5}
}

func (x *AM_DeleteBookingRequest_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AM_DeleteBookingRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AM_DeleteBookingRequest_Response) Reset() {
	*x = AM_DeleteBookingRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_DeleteBookingRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_DeleteBookingRequest_Response) ProtoMessage() {}

func (x *AM_DeleteBookingRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_DeleteBookingRequest_Response.ProtoReflect.Descriptor instead.
func (*AM_DeleteBookingRequest_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{6}
}

var File_booking_booking_service_proto protoreflect.FileDescriptor

var file_booking_booking_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x20, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4c, 0x0a, 0x21, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x44, 0x54, 0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xbd, 0x01, 0x0a, 0x19, 0x41, 0x4d, 0x5f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x20, 0x41, 0x4d, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x1f,
	0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x22, 0x0a, 0x20, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xe1, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x21, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x70, 0x0a, 0x12, 0x4d, 0x61, 0x6b,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x41, 0x4d, 0x5f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41, 0x4d,
	0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x2a, 0x15, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_booking_service_proto_rawDescOnce sync.Once
	file_booking_booking_service_proto_rawDescData = file_booking_booking_service_proto_rawDesc
)

func file_booking_booking_service_proto_rawDescGZIP() []byte {
	file_booking_booking_service_proto_rawDescOnce.Do(func() {
		file_booking_booking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_booking_service_proto_rawDescData)
	})
	return file_booking_booking_service_proto_rawDescData
}

var file_booking_booking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_booking_booking_service_proto_goTypes = []interface{}{
	(*AM_GetAllBookingRequests_Request)(nil),  // 0: AM_GetAllBookingRequests_Request
	(*AM_GetAllBookingRequests_Response)(nil), // 1: AM_GetAllBookingRequests_Response
	(*AM_BookingRequest_Request)(nil),         // 2: AM_BookingRequest_Request
	(*AM_CreateBookingRequest_Response)(nil),  // 3: AM_CreateBookingRequest_Response
	(*BookingRequestsDTO)(nil),                // 4: BookingRequestsDTO
	(*AM_DeleteBookingRequest_Request)(nil),   // 5: AM_DeleteBookingRequest_Request
	(*AM_DeleteBookingRequest_Response)(nil),  // 6: AM_DeleteBookingRequest_Response
}
var file_booking_booking_service_proto_depIdxs = []int32{
	4, // 0: AM_GetAllBookingRequests_Response.data:type_name -> BookingRequestsDTO
	0, // 1: BookingService.GetAll:input_type -> AM_GetAllBookingRequests_Request
	2, // 2: BookingService.MakeBookingRequest:input_type -> AM_BookingRequest_Request
	5, // 3: BookingService.DeleteBookingRequest:input_type -> AM_DeleteBookingRequest_Request
	1, // 4: BookingService.GetAll:output_type -> AM_GetAllBookingRequests_Response
	3, // 5: BookingService.MakeBookingRequest:output_type -> AM_CreateBookingRequest_Response
	6, // 6: BookingService.DeleteBookingRequest:output_type -> AM_DeleteBookingRequest_Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_booking_booking_service_proto_init() }
func file_booking_booking_service_proto_init() {
	if File_booking_booking_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_booking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequests_Request); i {
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
		file_booking_booking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequests_Response); i {
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
		file_booking_booking_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_BookingRequest_Request); i {
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
		file_booking_booking_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_CreateBookingRequest_Response); i {
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
		file_booking_booking_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequestsDTO); i {
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
		file_booking_booking_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_DeleteBookingRequest_Request); i {
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
		file_booking_booking_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_DeleteBookingRequest_Response); i {
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
			RawDescriptor: file_booking_booking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_booking_service_proto_goTypes,
		DependencyIndexes: file_booking_booking_service_proto_depIdxs,
		MessageInfos:      file_booking_booking_service_proto_msgTypes,
	}.Build()
	File_booking_booking_service_proto = out.File
	file_booking_booking_service_proto_rawDesc = nil
	file_booking_booking_service_proto_goTypes = nil
	file_booking_booking_service_proto_depIdxs = nil
}
