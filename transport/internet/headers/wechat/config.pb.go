package wechat

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

type VideoConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoConfig) Reset() {
	*x = VideoConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_headers_wechat_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoConfig) ProtoMessage() {}

func (x *VideoConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_wechat_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoConfig.ProtoReflect.Descriptor instead.
func (*VideoConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_wechat_config_proto_rawDescGZIP(), []int{0}
}

var File_transport_internet_headers_wechat_config_proto protoreflect.FileDescriptor

var file_transport_internet_headers_wechat_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2c, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x22, 0x0d,
	0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xa5, 0x01,
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x77, 0x65, 0x63, 0x68,
	0x61, 0x74, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x35, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0xaa, 0x02, 0x2c, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x57,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_headers_wechat_config_proto_rawDescOnce sync.Once
	file_transport_internet_headers_wechat_config_proto_rawDescData = file_transport_internet_headers_wechat_config_proto_rawDesc
)

func file_transport_internet_headers_wechat_config_proto_rawDescGZIP() []byte {
	file_transport_internet_headers_wechat_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_headers_wechat_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_headers_wechat_config_proto_rawDescData)
	})
	return file_transport_internet_headers_wechat_config_proto_rawDescData
}

var file_transport_internet_headers_wechat_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_headers_wechat_config_proto_goTypes = []interface{}{
	(*VideoConfig)(nil), // 0: v2fly.core.transport.internet.headers.wechat.VideoConfig
}
var file_transport_internet_headers_wechat_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_headers_wechat_config_proto_init() }
func file_transport_internet_headers_wechat_config_proto_init() {
	if File_transport_internet_headers_wechat_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_headers_wechat_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoConfig); i {
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
			RawDescriptor: file_transport_internet_headers_wechat_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_headers_wechat_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_headers_wechat_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_headers_wechat_config_proto_msgTypes,
	}.Build()
	File_transport_internet_headers_wechat_config_proto = out.File
	file_transport_internet_headers_wechat_config_proto_rawDesc = nil
	file_transport_internet_headers_wechat_config_proto_goTypes = nil
	file_transport_internet_headers_wechat_config_proto_depIdxs = nil
}
