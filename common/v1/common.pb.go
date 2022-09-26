// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: opensergo/common/v1/common.proto

package v1

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

type TimeUnit int32

const (
	// The time unit is not known.
	TimeUnit_UNKNOWN     TimeUnit = 0
	TimeUnit_MILLISECOND TimeUnit = 1
	TimeUnit_SECOND      TimeUnit = 2
	TimeUnit_MINUTE      TimeUnit = 3
	TimeUnit_HOUR        TimeUnit = 4
	TimeUnit_DAY         TimeUnit = 5
)

// Enum value maps for TimeUnit.
var (
	TimeUnit_name = map[int32]string{
		0: "UNKNOWN",
		1: "MILLISECOND",
		2: "SECOND",
		3: "MINUTE",
		4: "HOUR",
		5: "DAY",
	}
	TimeUnit_value = map[string]int32{
		"UNKNOWN":     0,
		"MILLISECOND": 1,
		"SECOND":      2,
		"MINUTE":      3,
		"HOUR":        4,
		"DAY":         5,
	}
)

func (x TimeUnit) Enum() *TimeUnit {
	p := new(TimeUnit)
	*p = x
	return p
}

func (x TimeUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_opensergo_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (TimeUnit) Type() protoreflect.EnumType {
	return &file_opensergo_common_v1_common_proto_enumTypes[0]
}

func (x TimeUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeUnit.Descriptor instead.
func (TimeUnit) EnumDescriptor() ([]byte, []int) {
	return file_opensergo_common_v1_common_proto_rawDescGZIP(), []int{0}
}

var File_opensergo_common_v1_common_proto protoreflect.FileDescriptor

var file_opensergo_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x53, 0x0a, 0x08, 0x54, 0x69,
	0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x48, 0x4f, 0x55, 0x52, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x05, 0x42,
	0x6e, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x6e, 0x61, 0x6e, 0x38, 0x30, 0x36, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x72, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opensergo_common_v1_common_proto_rawDescOnce sync.Once
	file_opensergo_common_v1_common_proto_rawDescData = file_opensergo_common_v1_common_proto_rawDesc
)

func file_opensergo_common_v1_common_proto_rawDescGZIP() []byte {
	file_opensergo_common_v1_common_proto_rawDescOnce.Do(func() {
		file_opensergo_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_opensergo_common_v1_common_proto_rawDescData)
	})
	return file_opensergo_common_v1_common_proto_rawDescData
}

var file_opensergo_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_opensergo_common_v1_common_proto_goTypes = []interface{}{
	(TimeUnit)(0), // 0: io.opensergo.common.v1.TimeUnit
}
var file_opensergo_common_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_opensergo_common_v1_common_proto_init() }
func file_opensergo_common_v1_common_proto_init() {
	if File_opensergo_common_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opensergo_common_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opensergo_common_v1_common_proto_goTypes,
		DependencyIndexes: file_opensergo_common_v1_common_proto_depIdxs,
		EnumInfos:         file_opensergo_common_v1_common_proto_enumTypes,
	}.Build()
	File_opensergo_common_v1_common_proto = out.File
	file_opensergo_common_v1_common_proto_rawDesc = nil
	file_opensergo_common_v1_common_proto_goTypes = nil
	file_opensergo_common_v1_common_proto_depIdxs = nil
}
