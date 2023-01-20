// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/music_service/credit/v1/credit_service.proto

package pb_music_credit

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

type CreateCreditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credits []*Credit `protobuf:"bytes,1,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *CreateCreditRequest) Reset() {
	*x = CreateCreditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_credit_v1_credit_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCreditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCreditRequest) ProtoMessage() {}

func (x *CreateCreditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_credit_v1_credit_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCreditRequest.ProtoReflect.Descriptor instead.
func (*CreateCreditRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_service_credit_v1_credit_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCreditRequest) GetCredits() []*Credit {
	if x != nil {
		return x.Credits
	}
	return nil
}

type CreateCreditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateCreditResponse) Reset() {
	*x = CreateCreditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_credit_v1_credit_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCreditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCreditResponse) ProtoMessage() {}

func (x *CreateCreditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_credit_v1_credit_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCreditResponse.ProtoReflect.Descriptor instead.
func (*CreateCreditResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_service_credit_v1_credit_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCreditResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_proto_music_service_credit_v1_credit_service_proto protoreflect.FileDescriptor

var file_proto_music_service_credit_v1_credit_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x56, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x56, 0x72, 0x4d, 0x6f, 0x6c, 0x6f, 0x64, 0x79, 0x61, 0x6b, 0x6f, 0x76, 0x2f, 0x76, 0x67,
	0x6d, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x62, 0x5f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_music_service_credit_v1_credit_service_proto_rawDescOnce sync.Once
	file_proto_music_service_credit_v1_credit_service_proto_rawDescData = file_proto_music_service_credit_v1_credit_service_proto_rawDesc
)

func file_proto_music_service_credit_v1_credit_service_proto_rawDescGZIP() []byte {
	file_proto_music_service_credit_v1_credit_service_proto_rawDescOnce.Do(func() {
		file_proto_music_service_credit_v1_credit_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_music_service_credit_v1_credit_service_proto_rawDescData)
	})
	return file_proto_music_service_credit_v1_credit_service_proto_rawDescData
}

var file_proto_music_service_credit_v1_credit_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_music_service_credit_v1_credit_service_proto_goTypes = []interface{}{
	(*CreateCreditRequest)(nil),  // 0: proto.music_service.credit.v1.CreateCreditRequest
	(*CreateCreditResponse)(nil), // 1: proto.music_service.credit.v1.CreateCreditResponse
	(*Credit)(nil),               // 2: proto.music_service.credit.v1.Credit
}
var file_proto_music_service_credit_v1_credit_service_proto_depIdxs = []int32{
	2, // 0: proto.music_service.credit.v1.CreateCreditRequest.credits:type_name -> proto.music_service.credit.v1.Credit
	0, // 1: proto.music_service.credit.v1.CreditService.CreateCredit:input_type -> proto.music_service.credit.v1.CreateCreditRequest
	1, // 2: proto.music_service.credit.v1.CreditService.CreateCredit:output_type -> proto.music_service.credit.v1.CreateCreditResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_music_service_credit_v1_credit_service_proto_init() }
func file_proto_music_service_credit_v1_credit_service_proto_init() {
	if File_proto_music_service_credit_v1_credit_service_proto != nil {
		return
	}
	file_proto_music_service_credit_v1_credit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_music_service_credit_v1_credit_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCreditRequest); i {
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
		file_proto_music_service_credit_v1_credit_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCreditResponse); i {
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
			RawDescriptor: file_proto_music_service_credit_v1_credit_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_music_service_credit_v1_credit_service_proto_goTypes,
		DependencyIndexes: file_proto_music_service_credit_v1_credit_service_proto_depIdxs,
		MessageInfos:      file_proto_music_service_credit_v1_credit_service_proto_msgTypes,
	}.Build()
	File_proto_music_service_credit_v1_credit_service_proto = out.File
	file_proto_music_service_credit_v1_credit_service_proto_rawDesc = nil
	file_proto_music_service_credit_v1_credit_service_proto_goTypes = nil
	file_proto_music_service_credit_v1_credit_service_proto_depIdxs = nil
}