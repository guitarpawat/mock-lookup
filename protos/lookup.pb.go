// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.4
// source: protos/lookup.proto

package __

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

type LookupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
}

func (x *LookupReq) Reset() {
	*x = LookupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lookup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupReq) ProtoMessage() {}

func (x *LookupReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lookup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupReq.ProtoReflect.Descriptor instead.
func (*LookupReq) Descriptor() ([]byte, []int) {
	return file_protos_lookup_proto_rawDescGZIP(), []int{0}
}

func (x *LookupReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type LookupRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *string `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Name  *string `protobuf:"bytes,2,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
}

func (x *LookupRes) Reset() {
	*x = LookupRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_lookup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRes) ProtoMessage() {}

func (x *LookupRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_lookup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRes.ProtoReflect.Descriptor instead.
func (*LookupRes) Descriptor() ([]byte, []int) {
	return file_protos_lookup_proto_rawDescGZIP(), []int{1}
}

func (x *LookupRes) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *LookupRes) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_protos_lookup_proto protoreflect.FileDescriptor

var file_protos_lookup_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x2d, 0x0a,
	0x09, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x09,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0x3f, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_lookup_proto_rawDescOnce sync.Once
	file_protos_lookup_proto_rawDescData = file_protos_lookup_proto_rawDesc
)

func file_protos_lookup_proto_rawDescGZIP() []byte {
	file_protos_lookup_proto_rawDescOnce.Do(func() {
		file_protos_lookup_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_lookup_proto_rawDescData)
	})
	return file_protos_lookup_proto_rawDescData
}

var file_protos_lookup_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_lookup_proto_goTypes = []interface{}{
	(*LookupReq)(nil), // 0: protos.LookupReq
	(*LookupRes)(nil), // 1: protos.LookupRes
}
var file_protos_lookup_proto_depIdxs = []int32{
	0, // 0: protos.LookupService.Lookup:input_type -> protos.LookupReq
	1, // 1: protos.LookupService.Lookup:output_type -> protos.LookupRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_lookup_proto_init() }
func file_protos_lookup_proto_init() {
	if File_protos_lookup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_lookup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupReq); i {
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
		file_protos_lookup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRes); i {
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
	file_protos_lookup_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_lookup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_lookup_proto_goTypes,
		DependencyIndexes: file_protos_lookup_proto_depIdxs,
		MessageInfos:      file_protos_lookup_proto_msgTypes,
	}.Build()
	File_protos_lookup_proto = out.File
	file_protos_lookup_proto_rawDesc = nil
	file_protos_lookup_proto_goTypes = nil
	file_protos_lookup_proto_depIdxs = nil
}
