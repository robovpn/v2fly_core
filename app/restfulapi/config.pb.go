package restfulapi

import (
	_ "../v2fly_core/common/protoext"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenAddr string `protobuf:"bytes,1,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	ListenPort int32  `protobuf:"varint,2,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	AuthToken  string `protobuf:"bytes,3,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_restfulapi_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_restfulapi_config_proto_msgTypes[0]
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
	return file_app_restfulapi_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Config) GetListenPort() int32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Config) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

var File_app_restfulapi_config_proto protoreflect.FileDescriptor

var file_app_restfulapi_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c,
	0x61, 0x70, 0x69, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x3a, 0x1d, 0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x82, 0xb5, 0x18, 0x0c, 0x12, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x61, 0x70, 0x69,
	0x42, 0x61, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x50, 0x01,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66,
	0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x61, 0x70, 0x69, 0xaa,
	0x02, 0x11, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_restfulapi_config_proto_rawDescOnce sync.Once
	file_app_restfulapi_config_proto_rawDescData = file_app_restfulapi_config_proto_rawDesc
)

func file_app_restfulapi_config_proto_rawDescGZIP() []byte {
	file_app_restfulapi_config_proto_rawDescOnce.Do(func() {
		file_app_restfulapi_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_restfulapi_config_proto_rawDescData)
	})
	return file_app_restfulapi_config_proto_rawDescData
}

var file_app_restfulapi_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_restfulapi_config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: v2fly.app.restfulapi.Config
}
var file_app_restfulapi_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_restfulapi_config_proto_init() }
func file_app_restfulapi_config_proto_init() {
	if File_app_restfulapi_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_restfulapi_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_restfulapi_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_restfulapi_config_proto_goTypes,
		DependencyIndexes: file_app_restfulapi_config_proto_depIdxs,
		MessageInfos:      file_app_restfulapi_config_proto_msgTypes,
	}.Build()
	File_app_restfulapi_config_proto = out.File
	file_app_restfulapi_config_proto_rawDesc = nil
	file_app_restfulapi_config_proto_goTypes = nil
	file_app_restfulapi_config_proto_depIdxs = nil
}
