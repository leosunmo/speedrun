// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: portal/portal.proto

package portal

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

type State int32

const (
	State_UNKNOWN   State = 0
	State_CHANGED   State = 1
	State_UNCHANGED State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "UNKNOWN",
		1: "CHANGED",
		2: "UNCHANGED",
	}
	State_value = map[string]int32{
		"UNKNOWN":   0,
		"CHANGED":   1,
		"UNCHANGED": 2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_portal_portal_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_portal_portal_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{0}
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{0}
}

func (x *CommandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   State  `protobuf:"varint,1,opt,name=state,proto3,enum=portal.State" json:"state,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{1}
}

func (x *CommandResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_UNKNOWN
}

func (x *CommandResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   State  `protobuf:"varint,1,opt,name=state,proto3,enum=portal.State" json:"state,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_UNKNOWN
}

func (x *ServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServiceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       State  `protobuf:"varint,1,opt,name=state,proto3,enum=portal.State" json:"state,omitempty"`
	Loadstate   string `protobuf:"bytes,2,opt,name=loadstate,proto3" json:"loadstate,omitempty"`
	Activestate string `protobuf:"bytes,3,opt,name=activestate,proto3" json:"activestate,omitempty"`
	Substate    string `protobuf:"bytes,4,opt,name=substate,proto3" json:"substate,omitempty"`
}

func (x *ServiceStatusResponse) Reset() {
	*x = ServiceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatusResponse) ProtoMessage() {}

func (x *ServiceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatusResponse.ProtoReflect.Descriptor instead.
func (*ServiceStatusResponse) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceStatusResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_UNKNOWN
}

func (x *ServiceStatusResponse) GetLoadstate() string {
	if x != nil {
		return x.Loadstate
	}
	return ""
}

func (x *ServiceStatusResponse) GetActivestate() string {
	if x != nil {
		return x.Activestate
	}
	return ""
}

func (x *ServiceStatusResponse) GetSubstate() string {
	if x != nil {
		return x.Substate
	}
	return ""
}

type CPUusageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CPUusageRequest) Reset() {
	*x = CPUusageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPUusageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUusageRequest) ProtoMessage() {}

func (x *CPUusageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUusageRequest.ProtoReflect.Descriptor instead.
func (*CPUusageRequest) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{5}
}

type CPUusageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loadavg1  int32 `protobuf:"varint,1,opt,name=loadavg1,proto3" json:"loadavg1,omitempty"`
	Loadavg5  int32 `protobuf:"varint,2,opt,name=loadavg5,proto3" json:"loadavg5,omitempty"`
	Loadavg15 int32 `protobuf:"varint,3,opt,name=loadavg15,proto3" json:"loadavg15,omitempty"`
}

func (x *CPUusageResponse) Reset() {
	*x = CPUusageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPUusageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUusageResponse) ProtoMessage() {}

func (x *CPUusageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUusageResponse.ProtoReflect.Descriptor instead.
func (*CPUusageResponse) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{6}
}

func (x *CPUusageResponse) GetLoadavg1() int32 {
	if x != nil {
		return x.Loadavg1
	}
	return 0
}

func (x *CPUusageResponse) GetLoadavg5() int32 {
	if x != nil {
		return x.Loadavg5
	}
	return 0
}

func (x *CPUusageResponse) GetLoadavg15() int32 {
	if x != nil {
		return x.Loadavg15
	}
	return 0
}

type FileReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileReadRequest) Reset() {
	*x = FileReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileReadRequest) ProtoMessage() {}

func (x *FileReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileReadRequest.ProtoReflect.Descriptor instead.
func (*FileReadRequest) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{7}
}

func (x *FileReadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FileReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   State  `protobuf:"varint,1,opt,name=state,proto3,enum=portal.State" json:"state,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileReadResponse) Reset() {
	*x = FileReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_portal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileReadResponse) ProtoMessage() {}

func (x *FileReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_portal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileReadResponse.ProtoReflect.Descriptor instead.
func (*FileReadResponse) Descriptor() ([]byte, []int) {
	return file_portal_portal_proto_rawDescGZIP(), []int{8}
}

func (x *FileReadResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_UNKNOWN
}

func (x *FileReadResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_portal_portal_proto protoreflect.FileDescriptor

var file_portal_portal_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x22, 0x38, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x50, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x50, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x43, 0x50, 0x55, 0x75, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x68, 0x0a, 0x10, 0x43, 0x50, 0x55, 0x75, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x31, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x35, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x31, 0x35, 0x22, 0x25, 0x0a, 0x0f, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x51, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2a, 0x30, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x44, 0x10, 0x02, 0x32, 0xdf, 0x03, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x12, 0x43, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x43, 0x50, 0x55, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x43, 0x50, 0x55, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x43, 0x50, 0x55, 0x75, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x70, 0x6f, 0x67, 0x6f, 0x72, 0x7a, 0x65, 0x6c, 0x73,
	0x6b, 0x69, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x72, 0x75, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portal_portal_proto_rawDescOnce sync.Once
	file_portal_portal_proto_rawDescData = file_portal_portal_proto_rawDesc
)

func file_portal_portal_proto_rawDescGZIP() []byte {
	file_portal_portal_proto_rawDescOnce.Do(func() {
		file_portal_portal_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_portal_proto_rawDescData)
	})
	return file_portal_portal_proto_rawDescData
}

var file_portal_portal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_portal_portal_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_portal_portal_proto_goTypes = []interface{}{
	(State)(0),                    // 0: portal.State
	(*CommandRequest)(nil),        // 1: portal.CommandRequest
	(*CommandResponse)(nil),       // 2: portal.CommandResponse
	(*ServiceRequest)(nil),        // 3: portal.ServiceRequest
	(*ServiceResponse)(nil),       // 4: portal.ServiceResponse
	(*ServiceStatusResponse)(nil), // 5: portal.ServiceStatusResponse
	(*CPUusageRequest)(nil),       // 6: portal.CPUusageRequest
	(*CPUusageResponse)(nil),      // 7: portal.CPUusageResponse
	(*FileReadRequest)(nil),       // 8: portal.FileReadRequest
	(*FileReadResponse)(nil),      // 9: portal.FileReadResponse
}
var file_portal_portal_proto_depIdxs = []int32{
	0,  // 0: portal.CommandResponse.state:type_name -> portal.State
	0,  // 1: portal.ServiceResponse.state:type_name -> portal.State
	0,  // 2: portal.ServiceStatusResponse.state:type_name -> portal.State
	0,  // 3: portal.FileReadResponse.state:type_name -> portal.State
	3,  // 4: portal.Portal.ServiceRestart:input_type -> portal.ServiceRequest
	3,  // 5: portal.Portal.ServiceStart:input_type -> portal.ServiceRequest
	3,  // 6: portal.Portal.ServiceStop:input_type -> portal.ServiceRequest
	3,  // 7: portal.Portal.ServiceStatus:input_type -> portal.ServiceRequest
	1,  // 8: portal.Portal.RunCommand:input_type -> portal.CommandRequest
	6,  // 9: portal.Portal.CPUusage:input_type -> portal.CPUusageRequest
	8,  // 10: portal.Portal.FileRead:input_type -> portal.FileReadRequest
	4,  // 11: portal.Portal.ServiceRestart:output_type -> portal.ServiceResponse
	4,  // 12: portal.Portal.ServiceStart:output_type -> portal.ServiceResponse
	4,  // 13: portal.Portal.ServiceStop:output_type -> portal.ServiceResponse
	5,  // 14: portal.Portal.ServiceStatus:output_type -> portal.ServiceStatusResponse
	2,  // 15: portal.Portal.RunCommand:output_type -> portal.CommandResponse
	7,  // 16: portal.Portal.CPUusage:output_type -> portal.CPUusageResponse
	9,  // 17: portal.Portal.FileRead:output_type -> portal.FileReadResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_portal_portal_proto_init() }
func file_portal_portal_proto_init() {
	if File_portal_portal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portal_portal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_portal_portal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_portal_portal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRequest); i {
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
		file_portal_portal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_portal_portal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatusResponse); i {
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
		file_portal_portal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPUusageRequest); i {
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
		file_portal_portal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPUusageResponse); i {
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
		file_portal_portal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileReadRequest); i {
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
		file_portal_portal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileReadResponse); i {
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
			RawDescriptor: file_portal_portal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portal_portal_proto_goTypes,
		DependencyIndexes: file_portal_portal_proto_depIdxs,
		EnumInfos:         file_portal_portal_proto_enumTypes,
		MessageInfos:      file_portal_portal_proto_msgTypes,
	}.Build()
	File_portal_portal_proto = out.File
	file_portal_portal_proto_rawDesc = nil
	file_portal_portal_proto_goTypes = nil
	file_portal_portal_proto_depIdxs = nil
}
