// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: common/proto/email_service.proto

package email_service_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscribtionCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailLogin string `protobuf:"bytes,1,opt,name=EmailLogin,proto3" json:"EmailLogin,omitempty"`
	EmailPass  string `protobuf:"bytes,2,opt,name=EmailPass,proto3" json:"EmailPass,omitempty"`
}

func (x *SubscribtionCmd) Reset() {
	*x = SubscribtionCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_email_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribtionCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribtionCmd) ProtoMessage() {}

func (x *SubscribtionCmd) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_email_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribtionCmd.ProtoReflect.Descriptor instead.
func (*SubscribtionCmd) Descriptor() ([]byte, []int) {
	return file_common_proto_email_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribtionCmd) GetEmailLogin() string {
	if x != nil {
		return x.EmailLogin
	}
	return ""
}

func (x *SubscribtionCmd) GetEmailPass() string {
	if x != nil {
		return x.EmailPass
	}
	return ""
}

type SubCmdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *SubCmdResponse) Reset() {
	*x = SubCmdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_email_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmdResponse) ProtoMessage() {}

func (x *SubCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_email_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmdResponse.ProtoReflect.Descriptor instead.
func (*SubCmdResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_email_service_proto_rawDescGZIP(), []int{1}
}

func (x *SubCmdResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UnreadCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailLogin string `protobuf:"bytes,1,opt,name=EmailLogin,proto3" json:"EmailLogin,omitempty"`
	EmailPass  string `protobuf:"bytes,2,opt,name=EmailPass,proto3" json:"EmailPass,omitempty"`
}

func (x *UnreadCountRequest) Reset() {
	*x = UnreadCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_email_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnreadCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnreadCountRequest) ProtoMessage() {}

func (x *UnreadCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_email_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnreadCountRequest.ProtoReflect.Descriptor instead.
func (*UnreadCountRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_email_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnreadCountRequest) GetEmailLogin() string {
	if x != nil {
		return x.EmailLogin
	}
	return ""
}

func (x *UnreadCountRequest) GetEmailPass() string {
	if x != nil {
		return x.EmailPass
	}
	return ""
}

type UnreadCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageCount int64  `protobuf:"varint,1,opt,name=MessageCount,proto3" json:"MessageCount,omitempty"`
	Error        string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *UnreadCountResponse) Reset() {
	*x = UnreadCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_email_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnreadCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnreadCountResponse) ProtoMessage() {}

func (x *UnreadCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_email_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnreadCountResponse.ProtoReflect.Descriptor instead.
func (*UnreadCountResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_email_service_proto_rawDescGZIP(), []int{3}
}

func (x *UnreadCountResponse) GetMessageCount() int64 {
	if x != nil {
		return x.MessageCount
	}
	return 0
}

func (x *UnreadCountResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type NewEmailCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	From    string `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *NewEmailCmd) Reset() {
	*x = NewEmailCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_email_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEmailCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEmailCmd) ProtoMessage() {}

func (x *NewEmailCmd) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_email_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEmailCmd.ProtoReflect.Descriptor instead.
func (*NewEmailCmd) Descriptor() ([]byte, []int) {
	return file_common_proto_email_service_proto_rawDescGZIP(), []int{4}
}

func (x *NewEmailCmd) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewEmailCmd) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *NewEmailCmd) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_proto_email_service_proto protoreflect.FileDescriptor

var file_common_proto_email_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0f,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x22, 0x26, 0x0a,
	0x0e, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x52, 0x0a, 0x12, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x13, 0x55, 0x6e, 0x72,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4b, 0x0a, 0x0b, 0x4e, 0x65,
	0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6d, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb7, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x72, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e,
	0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x6f, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x1c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6d, 0x64, 0x1a, 0x1b, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x4b, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x12, 0x3a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x18, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6d, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x75, 0x69,
	0x73, 0x65, 0x33, 0x32, 0x32, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x2d, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_email_service_proto_rawDescOnce sync.Once
	file_common_proto_email_service_proto_rawDescData = file_common_proto_email_service_proto_rawDesc
)

func file_common_proto_email_service_proto_rawDescGZIP() []byte {
	file_common_proto_email_service_proto_rawDescOnce.Do(func() {
		file_common_proto_email_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_email_service_proto_rawDescData)
	})
	return file_common_proto_email_service_proto_rawDescData
}

var file_common_proto_email_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_proto_email_service_proto_goTypes = []interface{}{
	(*SubscribtionCmd)(nil),     // 0: email_proto.SubscribtionCmd
	(*SubCmdResponse)(nil),      // 1: email_proto.SubCmdResponse
	(*UnreadCountRequest)(nil),  // 2: email_proto.UnreadCountRequest
	(*UnreadCountResponse)(nil), // 3: email_proto.UnreadCountResponse
	(*NewEmailCmd)(nil),         // 4: email_proto.NewEmailCmd
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_common_proto_email_service_proto_depIdxs = []int32{
	2, // 0: email_proto.EmailService.GetUnreadEmailCount:input_type -> email_proto.UnreadCountRequest
	0, // 1: email_proto.EmailService.SubscribeToInbox:input_type -> email_proto.SubscribtionCmd
	4, // 2: email_proto.NewEmailNotif.Notify:input_type -> email_proto.NewEmailCmd
	3, // 3: email_proto.EmailService.GetUnreadEmailCount:output_type -> email_proto.UnreadCountResponse
	1, // 4: email_proto.EmailService.SubscribeToInbox:output_type -> email_proto.SubCmdResponse
	5, // 5: email_proto.NewEmailNotif.Notify:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_email_service_proto_init() }
func file_common_proto_email_service_proto_init() {
	if File_common_proto_email_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_email_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribtionCmd); i {
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
		file_common_proto_email_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmdResponse); i {
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
		file_common_proto_email_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnreadCountRequest); i {
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
		file_common_proto_email_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnreadCountResponse); i {
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
		file_common_proto_email_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEmailCmd); i {
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
			RawDescriptor: file_common_proto_email_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_common_proto_email_service_proto_goTypes,
		DependencyIndexes: file_common_proto_email_service_proto_depIdxs,
		MessageInfos:      file_common_proto_email_service_proto_msgTypes,
	}.Build()
	File_common_proto_email_service_proto = out.File
	file_common_proto_email_service_proto_rawDesc = nil
	file_common_proto_email_service_proto_goTypes = nil
	file_common_proto_email_service_proto_depIdxs = nil
}
