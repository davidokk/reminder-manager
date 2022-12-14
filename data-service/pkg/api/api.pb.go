// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api.proto

package api

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reminder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Text string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Id   uint64                 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Reminder) Reset() {
	*x = Reminder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reminder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reminder) ProtoMessage() {}

func (x *Reminder) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reminder.ProtoReflect.Descriptor instead.
func (*Reminder) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Reminder) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Reminder) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Reminder) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReminderGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReminderGetRequest) Reset() {
	*x = ReminderGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderGetRequest) ProtoMessage() {}

func (x *ReminderGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderGetRequest.ProtoReflect.Descriptor instead.
func (*ReminderGetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *ReminderGetRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReminderGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reminder *Reminder `protobuf:"bytes,1,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *ReminderGetResponse) Reset() {
	*x = ReminderGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderGetResponse) ProtoMessage() {}

func (x *ReminderGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderGetResponse.ProtoReflect.Descriptor instead.
func (*ReminderGetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReminderGetResponse) GetReminder() *Reminder {
	if x != nil {
		return x.Reminder
	}
	return nil
}

type ReminderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReminderListRequest) Reset() {
	*x = ReminderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderListRequest) ProtoMessage() {}

func (x *ReminderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderListRequest.ProtoReflect.Descriptor instead.
func (*ReminderListRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

type ReminderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reminders []*Reminder `protobuf:"bytes,1,rep,name=reminders,proto3" json:"reminders,omitempty"`
}

func (x *ReminderListResponse) Reset() {
	*x = ReminderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderListResponse) ProtoMessage() {}

func (x *ReminderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderListResponse.ProtoReflect.Descriptor instead.
func (*ReminderListResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *ReminderListResponse) GetReminders() []*Reminder {
	if x != nil {
		return x.Reminders
	}
	return nil
}

type ReminderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Text string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ReminderCreateRequest) Reset() {
	*x = ReminderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderCreateRequest) ProtoMessage() {}

func (x *ReminderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderCreateRequest.ProtoReflect.Descriptor instead.
func (*ReminderCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ReminderCreateRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ReminderCreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ReminderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReminderCreateResponse) Reset() {
	*x = ReminderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderCreateResponse) ProtoMessage() {}

func (x *ReminderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderCreateResponse.ProtoReflect.Descriptor instead.
func (*ReminderCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *ReminderCreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReminderUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ReminderUpdateRequest) Reset() {
	*x = ReminderUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderUpdateRequest) ProtoMessage() {}

func (x *ReminderUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderUpdateRequest.ProtoReflect.Descriptor instead.
func (*ReminderUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *ReminderUpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReminderUpdateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ReminderUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReminderUpdateResponse) Reset() {
	*x = ReminderUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderUpdateResponse) ProtoMessage() {}

func (x *ReminderUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderUpdateResponse.ProtoReflect.Descriptor instead.
func (*ReminderUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

type ReminderRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReminderRemoveRequest) Reset() {
	*x = ReminderRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderRemoveRequest) ProtoMessage() {}

func (x *ReminderRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderRemoveRequest.ProtoReflect.Descriptor instead.
func (*ReminderRemoveRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *ReminderRemoveRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReminderRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReminderRemoveResponse) Reset() {
	*x = ReminderRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReminderRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReminderRemoveResponse) ProtoMessage() {}

func (x *ReminderRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReminderRemoveResponse.ProtoReflect.Descriptor instead.
func (*ReminderRemoveResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x08,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x5b, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3b, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x18, 0x0a, 0x16,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x18, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd0, 0x02, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x12, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x64, 0x61, 0x76, 0x69, 0x64, 0x6f, 0x6b, 0x6b, 0x2f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_goTypes = []interface{}{
	(*Reminder)(nil),               // 0: Reminder
	(*ReminderGetRequest)(nil),     // 1: ReminderGetRequest
	(*ReminderGetResponse)(nil),    // 2: ReminderGetResponse
	(*ReminderListRequest)(nil),    // 3: ReminderListRequest
	(*ReminderListResponse)(nil),   // 4: ReminderListResponse
	(*ReminderCreateRequest)(nil),  // 5: ReminderCreateRequest
	(*ReminderCreateResponse)(nil), // 6: ReminderCreateResponse
	(*ReminderUpdateRequest)(nil),  // 7: ReminderUpdateRequest
	(*ReminderUpdateResponse)(nil), // 8: ReminderUpdateResponse
	(*ReminderRemoveRequest)(nil),  // 9: ReminderRemoveRequest
	(*ReminderRemoveResponse)(nil), // 10: ReminderRemoveResponse
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_api_proto_depIdxs = []int32{
	11, // 0: Reminder.date:type_name -> google.protobuf.Timestamp
	0,  // 1: ReminderGetResponse.reminder:type_name -> Reminder
	0,  // 2: ReminderListResponse.reminders:type_name -> Reminder
	11, // 3: ReminderCreateRequest.date:type_name -> google.protobuf.Timestamp
	3,  // 4: Data.ReminderList:input_type -> ReminderListRequest
	1,  // 5: Data.ReminderGet:input_type -> ReminderGetRequest
	5,  // 6: Data.ReminderCreate:input_type -> ReminderCreateRequest
	7,  // 7: Data.ReminderUpdate:input_type -> ReminderUpdateRequest
	9,  // 8: Data.ReminderRemove:input_type -> ReminderRemoveRequest
	4,  // 9: Data.ReminderList:output_type -> ReminderListResponse
	2,  // 10: Data.ReminderGet:output_type -> ReminderGetResponse
	6,  // 11: Data.ReminderCreate:output_type -> ReminderCreateResponse
	8,  // 12: Data.ReminderUpdate:output_type -> ReminderUpdateResponse
	10, // 13: Data.ReminderRemove:output_type -> ReminderRemoveResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reminder); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderGetRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderGetResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderListRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderListResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderCreateRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderCreateResponse); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderUpdateRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderUpdateResponse); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderRemoveRequest); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReminderRemoveResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
