// Copyright © 2024 The Things Industries B.V.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: tti/gateway/data/lora/v1/error.proto

package lorav1

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

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNSPECIFIED  ErrorCode = 0
	ErrorCode_ERROR_CODE_INTERNAL     ErrorCode = 1
	ErrorCode_ERROR_CODE_TIMEOUT      ErrorCode = 2
	ErrorCode_ERROR_CODE_ARGUMENT     ErrorCode = 3
	ErrorCode_ERROR_CODE_TX_TOO_LATE  ErrorCode = 4
	ErrorCode_ERROR_CODE_TX_TOO_EARLY ErrorCode = 5
	ErrorCode_ERROR_CODE_TX_FREQUENCY ErrorCode = 6
	ErrorCode_ERROR_CODE_TX_POWER     ErrorCode = 7
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_UNSPECIFIED",
		1: "ERROR_CODE_INTERNAL",
		2: "ERROR_CODE_TIMEOUT",
		3: "ERROR_CODE_ARGUMENT",
		4: "ERROR_CODE_TX_TOO_LATE",
		5: "ERROR_CODE_TX_TOO_EARLY",
		6: "ERROR_CODE_TX_FREQUENCY",
		7: "ERROR_CODE_TX_POWER",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNSPECIFIED":  0,
		"ERROR_CODE_INTERNAL":     1,
		"ERROR_CODE_TIMEOUT":      2,
		"ERROR_CODE_ARGUMENT":     3,
		"ERROR_CODE_TX_TOO_LATE":  4,
		"ERROR_CODE_TX_TOO_EARLY": 5,
		"ERROR_CODE_TX_FREQUENCY": 6,
		"ERROR_CODE_TX_POWER":     7,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_tti_gateway_data_lora_v1_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_tti_gateway_data_lora_v1_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_error_proto_rawDescGZIP(), []int{0}
}

type ErrorNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=tti.gateway.data.lora.v1.ErrorCode" json:"code,omitempty"`
	Details string    `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *ErrorNotification) Reset() {
	*x = ErrorNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorNotification) ProtoMessage() {}

func (x *ErrorNotification) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorNotification.ProtoReflect.Descriptor instead.
func (*ErrorNotification) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorNotification) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *ErrorNotification) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_tti_gateway_data_lora_v1_error_proto protoreflect.FileDescriptor

var file_tti_gateway_data_lora_v1_error_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x74, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x22, 0x66, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0xe0, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x58, 0x5f, 0x54, 0x4f,
	0x4f, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x58, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x45, 0x41,
	0x52, 0x4c, 0x59, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x54, 0x58, 0x5f, 0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59,
	0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x54, 0x58, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x07, 0x42, 0xf0, 0x01, 0x0a, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x6f, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x74, 0x69, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x6f, 0x72,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x72, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x54, 0x47,
	0x44, 0x4c, 0xaa, 0x02, 0x18, 0x54, 0x74, 0x69, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x72, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18,
	0x54, 0x74, 0x69, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x44, 0x61, 0x74, 0x61,
	0x5c, 0x4c, 0x6f, 0x72, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x54, 0x74, 0x69, 0x5c, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x4c, 0x6f, 0x72, 0x61,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1c, 0x54, 0x74, 0x69, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0x3a,
	0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x4c, 0x6f, 0x72, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tti_gateway_data_lora_v1_error_proto_rawDescOnce sync.Once
	file_tti_gateway_data_lora_v1_error_proto_rawDescData = file_tti_gateway_data_lora_v1_error_proto_rawDesc
)

func file_tti_gateway_data_lora_v1_error_proto_rawDescGZIP() []byte {
	file_tti_gateway_data_lora_v1_error_proto_rawDescOnce.Do(func() {
		file_tti_gateway_data_lora_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_tti_gateway_data_lora_v1_error_proto_rawDescData)
	})
	return file_tti_gateway_data_lora_v1_error_proto_rawDescData
}

var file_tti_gateway_data_lora_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tti_gateway_data_lora_v1_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tti_gateway_data_lora_v1_error_proto_goTypes = []interface{}{
	(ErrorCode)(0),            // 0: tti.gateway.data.lora.v1.ErrorCode
	(*ErrorNotification)(nil), // 1: tti.gateway.data.lora.v1.ErrorNotification
}
var file_tti_gateway_data_lora_v1_error_proto_depIdxs = []int32{
	0, // 0: tti.gateway.data.lora.v1.ErrorNotification.code:type_name -> tti.gateway.data.lora.v1.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tti_gateway_data_lora_v1_error_proto_init() }
func file_tti_gateway_data_lora_v1_error_proto_init() {
	if File_tti_gateway_data_lora_v1_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tti_gateway_data_lora_v1_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorNotification); i {
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
			RawDescriptor: file_tti_gateway_data_lora_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tti_gateway_data_lora_v1_error_proto_goTypes,
		DependencyIndexes: file_tti_gateway_data_lora_v1_error_proto_depIdxs,
		EnumInfos:         file_tti_gateway_data_lora_v1_error_proto_enumTypes,
		MessageInfos:      file_tti_gateway_data_lora_v1_error_proto_msgTypes,
	}.Build()
	File_tti_gateway_data_lora_v1_error_proto = out.File
	file_tti_gateway_data_lora_v1_error_proto_rawDesc = nil
	file_tti_gateway_data_lora_v1_error_proto_goTypes = nil
	file_tti_gateway_data_lora_v1_error_proto_depIdxs = nil
}
