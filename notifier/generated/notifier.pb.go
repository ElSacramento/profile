// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: notifier/generated/notifier.proto

package generated

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_generated_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_generated_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_notifier_generated_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NotificationRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

type NotificationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationReply) Reset() {
	*x = NotificationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifier_generated_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationReply) ProtoMessage() {}

func (x *NotificationReply) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_generated_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationReply.ProtoReflect.Descriptor instead.
func (*NotificationReply) Descriptor() ([]byte, []int) {
	return file_notifier_generated_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_notifier_generated_notifier_proto protoreflect.FileDescriptor

var file_notifier_generated_notifier_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x43, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x50, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifier_generated_notifier_proto_rawDescOnce sync.Once
	file_notifier_generated_notifier_proto_rawDescData = file_notifier_generated_notifier_proto_rawDesc
)

func file_notifier_generated_notifier_proto_rawDescGZIP() []byte {
	file_notifier_generated_notifier_proto_rawDescOnce.Do(func() {
		file_notifier_generated_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifier_generated_notifier_proto_rawDescData)
	})
	return file_notifier_generated_notifier_proto_rawDescData
}

var file_notifier_generated_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notifier_generated_notifier_proto_goTypes = []interface{}{
	(*NotificationRequest)(nil), // 0: notifier.NotificationRequest
	(*NotificationReply)(nil),   // 1: notifier.NotificationReply
}
var file_notifier_generated_notifier_proto_depIdxs = []int32{
	0, // 0: notifier.Notifier.Push:input_type -> notifier.NotificationRequest
	1, // 1: notifier.Notifier.Push:output_type -> notifier.NotificationReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notifier_generated_notifier_proto_init() }
func file_notifier_generated_notifier_proto_init() {
	if File_notifier_generated_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifier_generated_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationRequest); i {
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
		file_notifier_generated_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationReply); i {
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
			RawDescriptor: file_notifier_generated_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifier_generated_notifier_proto_goTypes,
		DependencyIndexes: file_notifier_generated_notifier_proto_depIdxs,
		MessageInfos:      file_notifier_generated_notifier_proto_msgTypes,
	}.Build()
	File_notifier_generated_notifier_proto = out.File
	file_notifier_generated_notifier_proto_rawDesc = nil
	file_notifier_generated_notifier_proto_goTypes = nil
	file_notifier_generated_notifier_proto_depIdxs = nil
}
