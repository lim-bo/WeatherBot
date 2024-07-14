// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: weatherApi.proto

package weatherApi

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_weatherApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_weatherApi_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WeatherCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32              `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Main       map[string]float64 `protobuf:"bytes,2,rep,name=main,proto3" json:"main,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Wind       map[string]float64 `protobuf:"bytes,3,rep,name=wind,proto3" json:"wind,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *WeatherCast) Reset() {
	*x = WeatherCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherCast) ProtoMessage() {}

func (x *WeatherCast) ProtoReflect() protoreflect.Message {
	mi := &file_weatherApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherCast.ProtoReflect.Descriptor instead.
func (*WeatherCast) Descriptor() ([]byte, []int) {
	return file_weatherApi_proto_rawDescGZIP(), []int{1}
}

func (x *WeatherCast) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *WeatherCast) GetMain() map[string]float64 {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *WeatherCast) GetWind() map[string]float64 {
	if x != nil {
		return x.Wind
	}
	return nil
}

var File_weatherApi_proto protoreflect.FileDescriptor

var file_weatherApi_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x70, 0x69, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8d, 0x02, 0x0a, 0x0b, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x43, 0x61, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x41, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x61, 0x73, 0x74, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x35, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x43, 0x61, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x77, 0x69, 0x6e, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x37, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x70, 0x0a, 0x12, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x43, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x69, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41,
	0x70, 0x69, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x61, 0x73, 0x74, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x15, 0x5a, 0x13, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weatherApi_proto_rawDescOnce sync.Once
	file_weatherApi_proto_rawDescData = file_weatherApi_proto_rawDesc
)

func file_weatherApi_proto_rawDescGZIP() []byte {
	file_weatherApi_proto_rawDescOnce.Do(func() {
		file_weatherApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_weatherApi_proto_rawDescData)
	})
	return file_weatherApi_proto_rawDescData
}

var file_weatherApi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weatherApi_proto_goTypes = []any{
	(*City)(nil),        // 0: weatherApi.City
	(*WeatherCast)(nil), // 1: weatherApi.WeatherCast
	nil,                 // 2: weatherApi.WeatherCast.MainEntry
	nil,                 // 3: weatherApi.WeatherCast.WindEntry
}
var file_weatherApi_proto_depIdxs = []int32{
	2, // 0: weatherApi.WeatherCast.main:type_name -> weatherApi.WeatherCast.MainEntry
	3, // 1: weatherApi.WeatherCast.wind:type_name -> weatherApi.WeatherCast.WindEntry
	0, // 2: weatherApi.WeatherCastService.GetCurrentWeather:input_type -> weatherApi.City
	1, // 3: weatherApi.WeatherCastService.GetCurrentWeather:output_type -> weatherApi.WeatherCast
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weatherApi_proto_init() }
func file_weatherApi_proto_init() {
	if File_weatherApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weatherApi_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*City); i {
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
		file_weatherApi_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WeatherCast); i {
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
			RawDescriptor: file_weatherApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weatherApi_proto_goTypes,
		DependencyIndexes: file_weatherApi_proto_depIdxs,
		MessageInfos:      file_weatherApi_proto_msgTypes,
	}.Build()
	File_weatherApi_proto = out.File
	file_weatherApi_proto_rawDesc = nil
	file_weatherApi_proto_goTypes = nil
	file_weatherApi_proto_depIdxs = nil
}
