// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/ConfigService.proto

package ConfigService

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

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_proto_ConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string      `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Data    []*Property `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_ConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Config) GetData() []*Property {
	if x != nil {
		return x.Data
	}
	return nil
}

type Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *Configs) Reset() {
	*x = Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ConfigService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ConfigService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_proto_ConfigService_proto_rawDescGZIP(), []int{2}
}

func (x *Configs) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

type ConfigNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ConfigNameRequest) Reset() {
	*x = ConfigNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ConfigService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigNameRequest) ProtoMessage() {}

func (x *ConfigNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ConfigService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigNameRequest.ProtoReflect.Descriptor instead.
func (*ConfigNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_ConfigService_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigNameRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_proto_ConfigService_proto protoreflect.FileDescriptor

var file_proto_ConfigService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4f, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x20, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x52, 0x0a, 0x16, 0x67, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3f,
	0x0a, 0x0d, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x3c, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a,
	0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x20, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x30, 0x01, 0x12, 0x4f,
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x73, 0x65, 0x46,
	0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x13, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x2f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ConfigService_proto_rawDescOnce sync.Once
	file_proto_ConfigService_proto_rawDescData = file_proto_ConfigService_proto_rawDesc
)

func file_proto_ConfigService_proto_rawDescGZIP() []byte {
	file_proto_ConfigService_proto_rawDescOnce.Do(func() {
		file_proto_ConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ConfigService_proto_rawDescData)
	})
	return file_proto_ConfigService_proto_rawDescData
}

var file_proto_ConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_ConfigService_proto_goTypes = []interface{}{
	(*Property)(nil),          // 0: ConfigService.Property
	(*Config)(nil),            // 1: ConfigService.Config
	(*Configs)(nil),           // 2: ConfigService.Configs
	(*ConfigNameRequest)(nil), // 3: ConfigService.ConfigNameRequest
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_proto_ConfigService_proto_depIdxs = []int32{
	0,  // 0: ConfigService.Config.data:type_name -> ConfigService.Property
	1,  // 1: ConfigService.Configs.configs:type_name -> ConfigService.Config
	1,  // 2: ConfigService.ConfigService.addConfig:input_type -> ConfigService.Config
	3,  // 3: ConfigService.ConfigService.getConfig:input_type -> ConfigService.ConfigNameRequest
	3,  // 4: ConfigService.ConfigService.getAllVersionsOfConfig:input_type -> ConfigService.ConfigNameRequest
	4,  // 5: ConfigService.ConfigService.getAllConfigs:input_type -> google.protobuf.Empty
	1,  // 6: ConfigService.ConfigService.updateConfig:input_type -> ConfigService.Config
	3,  // 7: ConfigService.ConfigService.deleteConfig:input_type -> ConfigService.ConfigNameRequest
	3,  // 8: ConfigService.ConfigService.useConfig:input_type -> ConfigService.ConfigNameRequest
	3,  // 9: ConfigService.ConfigService.stopConfigUseForAll:input_type -> ConfigService.ConfigNameRequest
	1,  // 10: ConfigService.ConfigService.addConfig:output_type -> ConfigService.Config
	1,  // 11: ConfigService.ConfigService.getConfig:output_type -> ConfigService.Config
	2,  // 12: ConfigService.ConfigService.getAllVersionsOfConfig:output_type -> ConfigService.Configs
	2,  // 13: ConfigService.ConfigService.getAllConfigs:output_type -> ConfigService.Configs
	1,  // 14: ConfigService.ConfigService.updateConfig:output_type -> ConfigService.Config
	1,  // 15: ConfigService.ConfigService.deleteConfig:output_type -> ConfigService.Config
	1,  // 16: ConfigService.ConfigService.useConfig:output_type -> ConfigService.Config
	4,  // 17: ConfigService.ConfigService.stopConfigUseForAll:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ConfigService_proto_init() }
func file_proto_ConfigService_proto_init() {
	if File_proto_ConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_proto_ConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_ConfigService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configs); i {
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
		file_proto_ConfigService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigNameRequest); i {
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
			RawDescriptor: file_proto_ConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ConfigService_proto_goTypes,
		DependencyIndexes: file_proto_ConfigService_proto_depIdxs,
		MessageInfos:      file_proto_ConfigService_proto_msgTypes,
	}.Build()
	File_proto_ConfigService_proto = out.File
	file_proto_ConfigService_proto_rawDesc = nil
	file_proto_ConfigService_proto_goTypes = nil
	file_proto_ConfigService_proto_depIdxs = nil
}