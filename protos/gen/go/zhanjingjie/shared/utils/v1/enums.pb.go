// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: zhanjingjie/shared/utils/v1/enums.proto

package utilsv1

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

type Month_Enum int32

const (
	Month_ENUM_UNSPECIFIED Month_Enum = 0
	Month_ENUM_JANUARY     Month_Enum = 1
	Month_ENUM_FEBRUARY    Month_Enum = 2
)

// Enum value maps for Month_Enum.
var (
	Month_Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_JANUARY",
		2: "ENUM_FEBRUARY",
	}
	Month_Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_JANUARY":     1,
		"ENUM_FEBRUARY":    2,
	}
)

func (x Month_Enum) Enum() *Month_Enum {
	p := new(Month_Enum)
	*p = x
	return p
}

func (x Month_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Month_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_zhanjingjie_shared_utils_v1_enums_proto_enumTypes[0].Descriptor()
}

func (Month_Enum) Type() protoreflect.EnumType {
	return &file_zhanjingjie_shared_utils_v1_enums_proto_enumTypes[0]
}

func (x Month_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Month_Enum.Descriptor instead.
func (Month_Enum) EnumDescriptor() ([]byte, []int) {
	return file_zhanjingjie_shared_utils_v1_enums_proto_rawDescGZIP(), []int{0, 0}
}

type Month struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Month) Reset() {
	*x = Month{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zhanjingjie_shared_utils_v1_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Month) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Month) ProtoMessage() {}

func (x *Month) ProtoReflect() protoreflect.Message {
	mi := &file_zhanjingjie_shared_utils_v1_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Month.ProtoReflect.Descriptor instead.
func (*Month) Descriptor() ([]byte, []int) {
	return file_zhanjingjie_shared_utils_v1_enums_proto_rawDescGZIP(), []int{0}
}

var File_zhanjingjie_shared_utils_v1_enums_proto protoreflect.FileDescriptor

var file_zhanjingjie_shared_utils_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x27, 0x7a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x7a, 0x68, 0x61, 0x6e, 0x6a,
	0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x22,
	0x41, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4a, 0x41, 0x4e, 0x55, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x46, 0x45, 0x42, 0x52, 0x55, 0x41, 0x52, 0x59,
	0x10, 0x02, 0x42, 0x92, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x68, 0x61, 0x6e, 0x6a,
	0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x2f, 0x62, 0x75, 0x66,
	0x2d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x7a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a,
	0x69, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x5a, 0x53, 0x55,
	0xaa, 0x02, 0x1b, 0x5a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1b, 0x5a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x5c, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5c, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x5a,
	0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x6a, 0x69, 0x65, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5c, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x5a, 0x68, 0x61, 0x6e, 0x6a, 0x69, 0x6e,
	0x67, 0x6a, 0x69, 0x65, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x3a, 0x3a, 0x55, 0x74,
	0x69, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zhanjingjie_shared_utils_v1_enums_proto_rawDescOnce sync.Once
	file_zhanjingjie_shared_utils_v1_enums_proto_rawDescData = file_zhanjingjie_shared_utils_v1_enums_proto_rawDesc
)

func file_zhanjingjie_shared_utils_v1_enums_proto_rawDescGZIP() []byte {
	file_zhanjingjie_shared_utils_v1_enums_proto_rawDescOnce.Do(func() {
		file_zhanjingjie_shared_utils_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_zhanjingjie_shared_utils_v1_enums_proto_rawDescData)
	})
	return file_zhanjingjie_shared_utils_v1_enums_proto_rawDescData
}

var file_zhanjingjie_shared_utils_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zhanjingjie_shared_utils_v1_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zhanjingjie_shared_utils_v1_enums_proto_goTypes = []interface{}{
	(Month_Enum)(0), // 0: zhanjingjie.shared.utils.v1.Month.Enum
	(*Month)(nil),   // 1: zhanjingjie.shared.utils.v1.Month
}
var file_zhanjingjie_shared_utils_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zhanjingjie_shared_utils_v1_enums_proto_init() }
func file_zhanjingjie_shared_utils_v1_enums_proto_init() {
	if File_zhanjingjie_shared_utils_v1_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zhanjingjie_shared_utils_v1_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Month); i {
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
			RawDescriptor: file_zhanjingjie_shared_utils_v1_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zhanjingjie_shared_utils_v1_enums_proto_goTypes,
		DependencyIndexes: file_zhanjingjie_shared_utils_v1_enums_proto_depIdxs,
		EnumInfos:         file_zhanjingjie_shared_utils_v1_enums_proto_enumTypes,
		MessageInfos:      file_zhanjingjie_shared_utils_v1_enums_proto_msgTypes,
	}.Build()
	File_zhanjingjie_shared_utils_v1_enums_proto = out.File
	file_zhanjingjie_shared_utils_v1_enums_proto_rawDesc = nil
	file_zhanjingjie_shared_utils_v1_enums_proto_goTypes = nil
	file_zhanjingjie_shared_utils_v1_enums_proto_depIdxs = nil
}
