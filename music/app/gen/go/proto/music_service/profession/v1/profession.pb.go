// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/music_service/profession/v1/profession.proto

package pb_music_album

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

type Profession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfessionId      uint32 `protobuf:"varint,1,opt,name=profession_id,json=professionId,proto3" json:"profession_id,omitempty"`
	ProfessionalTitle string `protobuf:"bytes,2,opt,name=professional_title,json=professionalTitle,proto3" json:"professional_title,omitempty"`
}

func (x *Profession) Reset() {
	*x = Profession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_profession_v1_profession_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profession) ProtoMessage() {}

func (x *Profession) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_profession_v1_profession_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profession.ProtoReflect.Descriptor instead.
func (*Profession) Descriptor() ([]byte, []int) {
	return file_proto_music_service_profession_v1_profession_proto_rawDescGZIP(), []int{0}
}

func (x *Profession) GetProfessionId() uint32 {
	if x != nil {
		return x.ProfessionId
	}
	return 0
}

func (x *Profession) GetProfessionalTitle() string {
	if x != nil {
		return x.ProfessionalTitle
	}
	return ""
}

var File_proto_music_service_profession_v1_profession_proto protoreflect.FileDescriptor

var file_proto_music_service_profession_v1_profession_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x60, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x72, 0x4d, 0x6f, 0x6c, 0x6f, 0x64, 0x79,
	0x61, 0x6b, 0x6f, 0x76, 0x2f, 0x76, 0x67, 0x6d, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_music_service_profession_v1_profession_proto_rawDescOnce sync.Once
	file_proto_music_service_profession_v1_profession_proto_rawDescData = file_proto_music_service_profession_v1_profession_proto_rawDesc
)

func file_proto_music_service_profession_v1_profession_proto_rawDescGZIP() []byte {
	file_proto_music_service_profession_v1_profession_proto_rawDescOnce.Do(func() {
		file_proto_music_service_profession_v1_profession_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_music_service_profession_v1_profession_proto_rawDescData)
	})
	return file_proto_music_service_profession_v1_profession_proto_rawDescData
}

var file_proto_music_service_profession_v1_profession_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_music_service_profession_v1_profession_proto_goTypes = []interface{}{
	(*Profession)(nil), // 0: proto.music_service.profession.v1.Profession
}
var file_proto_music_service_profession_v1_profession_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_music_service_profession_v1_profession_proto_init() }
func file_proto_music_service_profession_v1_profession_proto_init() {
	if File_proto_music_service_profession_v1_profession_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_music_service_profession_v1_profession_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profession); i {
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
			RawDescriptor: file_proto_music_service_profession_v1_profession_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_music_service_profession_v1_profession_proto_goTypes,
		DependencyIndexes: file_proto_music_service_profession_v1_profession_proto_depIdxs,
		MessageInfos:      file_proto_music_service_profession_v1_profession_proto_msgTypes,
	}.Build()
	File_proto_music_service_profession_v1_profession_proto = out.File
	file_proto_music_service_profession_v1_profession_proto_rawDesc = nil
	file_proto_music_service_profession_v1_profession_proto_goTypes = nil
	file_proto_music_service_profession_v1_profession_proto_depIdxs = nil
}
