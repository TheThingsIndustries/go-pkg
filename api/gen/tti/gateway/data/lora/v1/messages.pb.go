// Copyright © 2024 The Things Industries B.V.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: tti/gateway/data/lora/v1/messages.proto

package lorav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerHelloNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerHelloNotification) Reset() {
	*x = ServerHelloNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerHelloNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerHelloNotification) ProtoMessage() {}

func (x *ServerHelloNotification) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerHelloNotification.ProtoReflect.Descriptor instead.
func (*ServerHelloNotification) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{0}
}

type ClientHelloNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime *durationpb.Duration `protobuf:"bytes,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// Organizationally Unique Identifier (OUI) assigned by IEEE.
	DeviceManufacturer uint32 `protobuf:"varint,2,opt,name=device_manufacturer,json=deviceManufacturer,proto3" json:"device_manufacturer,omitempty"`
	DeviceModel        string `protobuf:"bytes,3,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	HardwareVersion    string `protobuf:"bytes,4,opt,name=hardware_version,json=hardwareVersion,proto3" json:"hardware_version,omitempty"`
	RuntimeVersion     string `protobuf:"bytes,5,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	FirmwareVersion    string `protobuf:"bytes,6,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
}

func (x *ClientHelloNotification) Reset() {
	*x = ClientHelloNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientHelloNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientHelloNotification) ProtoMessage() {}

func (x *ClientHelloNotification) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientHelloNotification.ProtoReflect.Descriptor instead.
func (*ClientHelloNotification) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ClientHelloNotification) GetUptime() *durationpb.Duration {
	if x != nil {
		return x.Uptime
	}
	return nil
}

func (x *ClientHelloNotification) GetDeviceManufacturer() uint32 {
	if x != nil {
		return x.DeviceManufacturer
	}
	return 0
}

func (x *ClientHelloNotification) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *ClientHelloNotification) GetHardwareVersion() string {
	if x != nil {
		return x.HardwareVersion
	}
	return ""
}

func (x *ClientHelloNotification) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *ClientHelloNotification) GetFirmwareVersion() string {
	if x != nil {
		return x.FirmwareVersion
	}
	return ""
}

type StatusNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusNotification) Reset() {
	*x = StatusNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusNotification) ProtoMessage() {}

func (x *StatusNotification) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusNotification.ProtoReflect.Descriptor instead.
func (*StatusNotification) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{2}
}

type ConfigureLoraGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *GatewayConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConfigureLoraGatewayRequest) Reset() {
	*x = ConfigureLoraGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureLoraGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureLoraGatewayRequest) ProtoMessage() {}

func (x *ConfigureLoraGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureLoraGatewayRequest.ProtoReflect.Descriptor instead.
func (*ConfigureLoraGatewayRequest) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigureLoraGatewayRequest) GetConfig() *GatewayConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ConfigureLoraGatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigureLoraGatewayResponse) Reset() {
	*x = ConfigureLoraGatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureLoraGatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureLoraGatewayResponse) ProtoMessage() {}

func (x *ConfigureLoraGatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureLoraGatewayResponse.ProtoReflect.Descriptor instead.
func (*ConfigureLoraGatewayResponse) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{4}
}

type UplinkMessagesNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*UplinkMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *UplinkMessagesNotification) Reset() {
	*x = UplinkMessagesNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UplinkMessagesNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UplinkMessagesNotification) ProtoMessage() {}

func (x *UplinkMessagesNotification) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UplinkMessagesNotification.ProtoReflect.Descriptor instead.
func (*UplinkMessagesNotification) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *UplinkMessagesNotification) GetMessages() []*UplinkMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type NetworkServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId uint32 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Types that are assignable to Message:
	//
	//	*NetworkServerMessage_ErrorNotification
	//	*NetworkServerMessage_ServerHelloNotification
	//	*NetworkServerMessage_ConfigureLoraGatewayRequest
	Message isNetworkServerMessage_Message `protobuf_oneof:"message"`
}

func (x *NetworkServerMessage) Reset() {
	*x = NetworkServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServerMessage) ProtoMessage() {}

func (x *NetworkServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServerMessage.ProtoReflect.Descriptor instead.
func (*NetworkServerMessage) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *NetworkServerMessage) GetTransactionId() uint32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (m *NetworkServerMessage) GetMessage() isNetworkServerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *NetworkServerMessage) GetErrorNotification() *ErrorNotification {
	if x, ok := x.GetMessage().(*NetworkServerMessage_ErrorNotification); ok {
		return x.ErrorNotification
	}
	return nil
}

func (x *NetworkServerMessage) GetServerHelloNotification() *ServerHelloNotification {
	if x, ok := x.GetMessage().(*NetworkServerMessage_ServerHelloNotification); ok {
		return x.ServerHelloNotification
	}
	return nil
}

func (x *NetworkServerMessage) GetConfigureLoraGatewayRequest() *ConfigureLoraGatewayRequest {
	if x, ok := x.GetMessage().(*NetworkServerMessage_ConfigureLoraGatewayRequest); ok {
		return x.ConfigureLoraGatewayRequest
	}
	return nil
}

type isNetworkServerMessage_Message interface {
	isNetworkServerMessage_Message()
}

type NetworkServerMessage_ErrorNotification struct {
	ErrorNotification *ErrorNotification `protobuf:"bytes,2,opt,name=error_notification,json=errorNotification,proto3,oneof"`
}

type NetworkServerMessage_ServerHelloNotification struct {
	ServerHelloNotification *ServerHelloNotification `protobuf:"bytes,3,opt,name=server_hello_notification,json=serverHelloNotification,proto3,oneof"`
}

type NetworkServerMessage_ConfigureLoraGatewayRequest struct {
	ConfigureLoraGatewayRequest *ConfigureLoraGatewayRequest `protobuf:"bytes,4,opt,name=configure_lora_gateway_request,json=configureLoraGatewayRequest,proto3,oneof"`
}

func (*NetworkServerMessage_ErrorNotification) isNetworkServerMessage_Message() {}

func (*NetworkServerMessage_ServerHelloNotification) isNetworkServerMessage_Message() {}

func (*NetworkServerMessage_ConfigureLoraGatewayRequest) isNetworkServerMessage_Message() {}

type GatewayMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId uint32 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Types that are assignable to Message:
	//
	//	*GatewayMessage_ErrorNotification
	//	*GatewayMessage_ClientHelloNotification
	//	*GatewayMessage_StatusNotification
	//	*GatewayMessage_ConfigureLoraGatewayResponse
	//	*GatewayMessage_UplinkMessagesNotification
	Message isGatewayMessage_Message `protobuf_oneof:"message"`
}

func (x *GatewayMessage) Reset() {
	*x = GatewayMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMessage) ProtoMessage() {}

func (x *GatewayMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tti_gateway_data_lora_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMessage.ProtoReflect.Descriptor instead.
func (*GatewayMessage) Descriptor() ([]byte, []int) {
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *GatewayMessage) GetTransactionId() uint32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (m *GatewayMessage) GetMessage() isGatewayMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *GatewayMessage) GetErrorNotification() *ErrorNotification {
	if x, ok := x.GetMessage().(*GatewayMessage_ErrorNotification); ok {
		return x.ErrorNotification
	}
	return nil
}

func (x *GatewayMessage) GetClientHelloNotification() *ClientHelloNotification {
	if x, ok := x.GetMessage().(*GatewayMessage_ClientHelloNotification); ok {
		return x.ClientHelloNotification
	}
	return nil
}

func (x *GatewayMessage) GetStatusNotification() *StatusNotification {
	if x, ok := x.GetMessage().(*GatewayMessage_StatusNotification); ok {
		return x.StatusNotification
	}
	return nil
}

func (x *GatewayMessage) GetConfigureLoraGatewayResponse() *ConfigureLoraGatewayResponse {
	if x, ok := x.GetMessage().(*GatewayMessage_ConfigureLoraGatewayResponse); ok {
		return x.ConfigureLoraGatewayResponse
	}
	return nil
}

func (x *GatewayMessage) GetUplinkMessagesNotification() *UplinkMessagesNotification {
	if x, ok := x.GetMessage().(*GatewayMessage_UplinkMessagesNotification); ok {
		return x.UplinkMessagesNotification
	}
	return nil
}

type isGatewayMessage_Message interface {
	isGatewayMessage_Message()
}

type GatewayMessage_ErrorNotification struct {
	ErrorNotification *ErrorNotification `protobuf:"bytes,2,opt,name=error_notification,json=errorNotification,proto3,oneof"`
}

type GatewayMessage_ClientHelloNotification struct {
	ClientHelloNotification *ClientHelloNotification `protobuf:"bytes,3,opt,name=client_hello_notification,json=clientHelloNotification,proto3,oneof"`
}

type GatewayMessage_StatusNotification struct {
	StatusNotification *StatusNotification `protobuf:"bytes,4,opt,name=status_notification,json=statusNotification,proto3,oneof"`
}

type GatewayMessage_ConfigureLoraGatewayResponse struct {
	ConfigureLoraGatewayResponse *ConfigureLoraGatewayResponse `protobuf:"bytes,5,opt,name=configure_lora_gateway_response,json=configureLoraGatewayResponse,proto3,oneof"`
}

type GatewayMessage_UplinkMessagesNotification struct {
	UplinkMessagesNotification *UplinkMessagesNotification `protobuf:"bytes,6,opt,name=uplink_messages_notification,json=uplinkMessagesNotification,proto3,oneof"`
}

func (*GatewayMessage_ErrorNotification) isGatewayMessage_Message() {}

func (*GatewayMessage_ClientHelloNotification) isGatewayMessage_Message() {}

func (*GatewayMessage_StatusNotification) isGatewayMessage_Message() {}

func (*GatewayMessage_ConfigureLoraGatewayResponse) isGatewayMessage_Message() {}

func (*GatewayMessage_UplinkMessagesNotification) isGatewayMessage_Message() {}

var File_tti_gateway_data_lora_v1_messages_proto protoreflect.FileDescriptor

var file_tti_gateway_data_lora_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x74, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x74, 0x69, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x74, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x74, 0x69, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x6f, 0x72, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19,
	0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9f, 0x02, 0x0a, 0x17, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x29, 0x0a, 0x10,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5e, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4c, 0x6f,
	0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4c, 0x6f,
	0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x61, 0x0a, 0x1a, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x43, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x95, 0x03, 0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x12, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x5f, 0x6c, 0x6f, 0x72, 0x61, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x74, 0x74,
	0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x4c, 0x6f, 0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x1b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4c,
	0x6f, 0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xed, 0x04, 0x0a,
	0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x12, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x17, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x72, 0x61, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x4c, 0x6f, 0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x1c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x4c, 0x6f, 0x72, 0x61, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x1c, 0x75, 0x70, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1a, 0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xf3, 0x01, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x74, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x69, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x74, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x72, 0x61, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x54, 0x47, 0x44, 0x4c, 0xaa, 0x02, 0x18, 0x54, 0x74, 0x69, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x72, 0x61, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x18, 0x54, 0x74, 0x69, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c,
	0x44, 0x61, 0x74, 0x61, 0x5c, 0x4c, 0x6f, 0x72, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x54,
	0x74, 0x69, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c,
	0x4c, 0x6f, 0x72, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x54, 0x74, 0x69, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x4c, 0x6f, 0x72, 0x61, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tti_gateway_data_lora_v1_messages_proto_rawDescOnce sync.Once
	file_tti_gateway_data_lora_v1_messages_proto_rawDescData = file_tti_gateway_data_lora_v1_messages_proto_rawDesc
)

func file_tti_gateway_data_lora_v1_messages_proto_rawDescGZIP() []byte {
	file_tti_gateway_data_lora_v1_messages_proto_rawDescOnce.Do(func() {
		file_tti_gateway_data_lora_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_tti_gateway_data_lora_v1_messages_proto_rawDescData)
	})
	return file_tti_gateway_data_lora_v1_messages_proto_rawDescData
}

var file_tti_gateway_data_lora_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tti_gateway_data_lora_v1_messages_proto_goTypes = []interface{}{
	(*ServerHelloNotification)(nil),      // 0: tti.gateway.data.lora.v1.ServerHelloNotification
	(*ClientHelloNotification)(nil),      // 1: tti.gateway.data.lora.v1.ClientHelloNotification
	(*StatusNotification)(nil),           // 2: tti.gateway.data.lora.v1.StatusNotification
	(*ConfigureLoraGatewayRequest)(nil),  // 3: tti.gateway.data.lora.v1.ConfigureLoraGatewayRequest
	(*ConfigureLoraGatewayResponse)(nil), // 4: tti.gateway.data.lora.v1.ConfigureLoraGatewayResponse
	(*UplinkMessagesNotification)(nil),   // 5: tti.gateway.data.lora.v1.UplinkMessagesNotification
	(*NetworkServerMessage)(nil),         // 6: tti.gateway.data.lora.v1.NetworkServerMessage
	(*GatewayMessage)(nil),               // 7: tti.gateway.data.lora.v1.GatewayMessage
	(*durationpb.Duration)(nil),          // 8: google.protobuf.Duration
	(*GatewayConfig)(nil),                // 9: tti.gateway.data.lora.v1.GatewayConfig
	(*UplinkMessage)(nil),                // 10: tti.gateway.data.lora.v1.UplinkMessage
	(*ErrorNotification)(nil),            // 11: tti.gateway.data.lora.v1.ErrorNotification
}
var file_tti_gateway_data_lora_v1_messages_proto_depIdxs = []int32{
	8,  // 0: tti.gateway.data.lora.v1.ClientHelloNotification.uptime:type_name -> google.protobuf.Duration
	9,  // 1: tti.gateway.data.lora.v1.ConfigureLoraGatewayRequest.config:type_name -> tti.gateway.data.lora.v1.GatewayConfig
	10, // 2: tti.gateway.data.lora.v1.UplinkMessagesNotification.messages:type_name -> tti.gateway.data.lora.v1.UplinkMessage
	11, // 3: tti.gateway.data.lora.v1.NetworkServerMessage.error_notification:type_name -> tti.gateway.data.lora.v1.ErrorNotification
	0,  // 4: tti.gateway.data.lora.v1.NetworkServerMessage.server_hello_notification:type_name -> tti.gateway.data.lora.v1.ServerHelloNotification
	3,  // 5: tti.gateway.data.lora.v1.NetworkServerMessage.configure_lora_gateway_request:type_name -> tti.gateway.data.lora.v1.ConfigureLoraGatewayRequest
	11, // 6: tti.gateway.data.lora.v1.GatewayMessage.error_notification:type_name -> tti.gateway.data.lora.v1.ErrorNotification
	1,  // 7: tti.gateway.data.lora.v1.GatewayMessage.client_hello_notification:type_name -> tti.gateway.data.lora.v1.ClientHelloNotification
	2,  // 8: tti.gateway.data.lora.v1.GatewayMessage.status_notification:type_name -> tti.gateway.data.lora.v1.StatusNotification
	4,  // 9: tti.gateway.data.lora.v1.GatewayMessage.configure_lora_gateway_response:type_name -> tti.gateway.data.lora.v1.ConfigureLoraGatewayResponse
	5,  // 10: tti.gateway.data.lora.v1.GatewayMessage.uplink_messages_notification:type_name -> tti.gateway.data.lora.v1.UplinkMessagesNotification
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tti_gateway_data_lora_v1_messages_proto_init() }
func file_tti_gateway_data_lora_v1_messages_proto_init() {
	if File_tti_gateway_data_lora_v1_messages_proto != nil {
		return
	}
	file_tti_gateway_data_lora_v1_error_proto_init()
	file_tti_gateway_data_lora_v1_lora_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerHelloNotification); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientHelloNotification); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusNotification); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureLoraGatewayRequest); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureLoraGatewayResponse); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UplinkMessagesNotification); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkServerMessage); i {
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
		file_tti_gateway_data_lora_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMessage); i {
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
	file_tti_gateway_data_lora_v1_messages_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*NetworkServerMessage_ErrorNotification)(nil),
		(*NetworkServerMessage_ServerHelloNotification)(nil),
		(*NetworkServerMessage_ConfigureLoraGatewayRequest)(nil),
	}
	file_tti_gateway_data_lora_v1_messages_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*GatewayMessage_ErrorNotification)(nil),
		(*GatewayMessage_ClientHelloNotification)(nil),
		(*GatewayMessage_StatusNotification)(nil),
		(*GatewayMessage_ConfigureLoraGatewayResponse)(nil),
		(*GatewayMessage_UplinkMessagesNotification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tti_gateway_data_lora_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tti_gateway_data_lora_v1_messages_proto_goTypes,
		DependencyIndexes: file_tti_gateway_data_lora_v1_messages_proto_depIdxs,
		MessageInfos:      file_tti_gateway_data_lora_v1_messages_proto_msgTypes,
	}.Build()
	File_tti_gateway_data_lora_v1_messages_proto = out.File
	file_tti_gateway_data_lora_v1_messages_proto_rawDesc = nil
	file_tti_gateway_data_lora_v1_messages_proto_goTypes = nil
	file_tti_gateway_data_lora_v1_messages_proto_depIdxs = nil
}