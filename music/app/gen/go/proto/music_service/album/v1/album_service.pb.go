// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/music_service/album/v1/album_service.proto

package pb_music_album

import (
	v13 "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/filter/v1"
	v11 "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/credit/v1"
	v1 "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/info/v1"
	v12 "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/track/v1"
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

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAlbumRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *CreateAlbumResponse) Reset() {
	*x = CreateAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumResponse) ProtoMessage() {}

func (x *CreateAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumResponse.ProtoReflect.Descriptor instead.
func (*CreateAlbumResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type FindAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *FindAlbumRequest) Reset() {
	*x = FindAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAlbumRequest) ProtoMessage() {}

func (x *FindAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAlbumRequest.ProtoReflect.Descriptor instead.
func (*FindAlbumRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{2}
}

func (x *FindAlbumRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FullAlbum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album     *Album        `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
	Info      *v1.Info      `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Credits   []*v11.Credit `protobuf:"bytes,2,rep,name=credits,proto3" json:"credits,omitempty"`
	Tracklist []*v12.Track  `protobuf:"bytes,3,rep,name=tracklist,proto3" json:"tracklist,omitempty"`
}

func (x *FullAlbum) Reset() {
	*x = FullAlbum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullAlbum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullAlbum) ProtoMessage() {}

func (x *FullAlbum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullAlbum.ProtoReflect.Descriptor instead.
func (*FullAlbum) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{3}
}

func (x *FullAlbum) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

func (x *FullAlbum) GetInfo() *v1.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FullAlbum) GetCredits() []*v11.Credit {
	if x != nil {
		return x.Credits
	}
	return nil
}

func (x *FullAlbum) GetTracklist() []*v12.Track {
	if x != nil {
		return x.Tracklist
	}
	return nil
}

type FindAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *FindAlbumResponse) Reset() {
	*x = FindAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAlbumResponse) ProtoMessage() {}

func (x *FindAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAlbumResponse.ProtoReflect.Descriptor instead.
func (*FindAlbumResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{4}
}

func (x *FindAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type FindFullAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *FindFullAlbumRequest) Reset() {
	*x = FindFullAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindFullAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindFullAlbumRequest) ProtoMessage() {}

func (x *FindFullAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindFullAlbumRequest.ProtoReflect.Descriptor instead.
func (*FindFullAlbumRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{5}
}

func (x *FindFullAlbumRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FindFullAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *FullAlbum `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *FindFullAlbumResponse) Reset() {
	*x = FindFullAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindFullAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindFullAlbumResponse) ProtoMessage() {}

func (x *FindFullAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindFullAlbumResponse.ProtoReflect.Descriptor instead.
func (*FindFullAlbumResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{6}
}

func (x *FindFullAlbumResponse) GetAlbum() *FullAlbum {
	if x != nil {
		return x.Album
	}
	return nil
}

type FindAllAlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination  *v13.Pagination        `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Name        *v13.StringFieldFilter `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Person      *v13.StringFieldFilter `protobuf:"bytes,3,opt,name=person,proto3" json:"person,omitempty"`
	PublishedAt *v13.IntFieldFilter    `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	Sort        *v13.Sort              `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *FindAllAlbumsRequest) Reset() {
	*x = FindAllAlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAlbumsRequest) ProtoMessage() {}

func (x *FindAllAlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAlbumsRequest.ProtoReflect.Descriptor instead.
func (*FindAllAlbumsRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{7}
}

func (x *FindAllAlbumsRequest) GetPagination() *v13.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *FindAllAlbumsRequest) GetName() *v13.StringFieldFilter {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FindAllAlbumsRequest) GetPerson() *v13.StringFieldFilter {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *FindAllAlbumsRequest) GetPublishedAt() *v13.IntFieldFilter {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *FindAllAlbumsRequest) GetSort() *v13.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type FindAllAlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *FindAllAlbumsResponse) Reset() {
	*x = FindAllAlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAlbumsResponse) ProtoMessage() {}

func (x *FindAllAlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_service_album_v1_album_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAlbumsResponse.ProtoReflect.Descriptor instead.
func (*FindAllAlbumsResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_service_album_v1_album_service_proto_rawDescGZIP(), []int{8}
}

func (x *FindAllAlbumsResponse) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

var File_proto_music_service_album_v1_album_service_proto protoreflect.FileDescriptor

var file_proto_music_service_album_v1_album_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x50, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x81,
	0x02, 0x0a, 0x09, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x39, 0x0a, 0x05,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3f,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12,
	0x41, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x4e, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x22, 0x2c, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x56, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0xb6, 0x02, 0x0a, 0x14, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x22, 0x54, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x32, 0xec, 0x03, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e,
	0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x2e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a,
	0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0d, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x72, 0x4d, 0x6f, 0x6c, 0x6f, 0x64, 0x79, 0x61, 0x6b, 0x6f,
	0x76, 0x2f, 0x76, 0x67, 0x6d, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_music_service_album_v1_album_service_proto_rawDescOnce sync.Once
	file_proto_music_service_album_v1_album_service_proto_rawDescData = file_proto_music_service_album_v1_album_service_proto_rawDesc
)

func file_proto_music_service_album_v1_album_service_proto_rawDescGZIP() []byte {
	file_proto_music_service_album_v1_album_service_proto_rawDescOnce.Do(func() {
		file_proto_music_service_album_v1_album_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_music_service_album_v1_album_service_proto_rawDescData)
	})
	return file_proto_music_service_album_v1_album_service_proto_rawDescData
}

var file_proto_music_service_album_v1_album_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_music_service_album_v1_album_service_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil),    // 0: proto.music_service.album.v1.CreateAlbumRequest
	(*CreateAlbumResponse)(nil),   // 1: proto.music_service.album.v1.CreateAlbumResponse
	(*FindAlbumRequest)(nil),      // 2: proto.music_service.album.v1.FindAlbumRequest
	(*FullAlbum)(nil),             // 3: proto.music_service.album.v1.FullAlbum
	(*FindAlbumResponse)(nil),     // 4: proto.music_service.album.v1.FindAlbumResponse
	(*FindFullAlbumRequest)(nil),  // 5: proto.music_service.album.v1.FindFullAlbumRequest
	(*FindFullAlbumResponse)(nil), // 6: proto.music_service.album.v1.FindFullAlbumResponse
	(*FindAllAlbumsRequest)(nil),  // 7: proto.music_service.album.v1.FindAllAlbumsRequest
	(*FindAllAlbumsResponse)(nil), // 8: proto.music_service.album.v1.FindAllAlbumsResponse
	(*Album)(nil),                 // 9: proto.music_service.album.v1.Album
	(*v1.Info)(nil),               // 10: proto.music_service.info.v1.Info
	(*v11.Credit)(nil),            // 11: proto.music_service.credit.v1.Credit
	(*v12.Track)(nil),             // 12: proto.music_service.track.v1.Track
	(*v13.Pagination)(nil),        // 13: proto.filter.v1.Pagination
	(*v13.StringFieldFilter)(nil), // 14: proto.filter.v1.StringFieldFilter
	(*v13.IntFieldFilter)(nil),    // 15: proto.filter.v1.IntFieldFilter
	(*v13.Sort)(nil),              // 16: proto.filter.v1.Sort
}
var file_proto_music_service_album_v1_album_service_proto_depIdxs = []int32{
	9,  // 0: proto.music_service.album.v1.CreateAlbumResponse.album:type_name -> proto.music_service.album.v1.Album
	9,  // 1: proto.music_service.album.v1.FullAlbum.album:type_name -> proto.music_service.album.v1.Album
	10, // 2: proto.music_service.album.v1.FullAlbum.info:type_name -> proto.music_service.info.v1.Info
	11, // 3: proto.music_service.album.v1.FullAlbum.credits:type_name -> proto.music_service.credit.v1.Credit
	12, // 4: proto.music_service.album.v1.FullAlbum.tracklist:type_name -> proto.music_service.track.v1.Track
	9,  // 5: proto.music_service.album.v1.FindAlbumResponse.album:type_name -> proto.music_service.album.v1.Album
	3,  // 6: proto.music_service.album.v1.FindFullAlbumResponse.album:type_name -> proto.music_service.album.v1.FullAlbum
	13, // 7: proto.music_service.album.v1.FindAllAlbumsRequest.pagination:type_name -> proto.filter.v1.Pagination
	14, // 8: proto.music_service.album.v1.FindAllAlbumsRequest.name:type_name -> proto.filter.v1.StringFieldFilter
	14, // 9: proto.music_service.album.v1.FindAllAlbumsRequest.person:type_name -> proto.filter.v1.StringFieldFilter
	15, // 10: proto.music_service.album.v1.FindAllAlbumsRequest.published_at:type_name -> proto.filter.v1.IntFieldFilter
	16, // 11: proto.music_service.album.v1.FindAllAlbumsRequest.sort:type_name -> proto.filter.v1.Sort
	9,  // 12: proto.music_service.album.v1.FindAllAlbumsResponse.albums:type_name -> proto.music_service.album.v1.Album
	0,  // 13: proto.music_service.album.v1.AlbumService.CreateAlbum:input_type -> proto.music_service.album.v1.CreateAlbumRequest
	2,  // 14: proto.music_service.album.v1.AlbumService.FindAlbum:input_type -> proto.music_service.album.v1.FindAlbumRequest
	5,  // 15: proto.music_service.album.v1.AlbumService.FindFullAlbum:input_type -> proto.music_service.album.v1.FindFullAlbumRequest
	7,  // 16: proto.music_service.album.v1.AlbumService.FindAllAlbums:input_type -> proto.music_service.album.v1.FindAllAlbumsRequest
	1,  // 17: proto.music_service.album.v1.AlbumService.CreateAlbum:output_type -> proto.music_service.album.v1.CreateAlbumResponse
	4,  // 18: proto.music_service.album.v1.AlbumService.FindAlbum:output_type -> proto.music_service.album.v1.FindAlbumResponse
	6,  // 19: proto.music_service.album.v1.AlbumService.FindFullAlbum:output_type -> proto.music_service.album.v1.FindFullAlbumResponse
	8,  // 20: proto.music_service.album.v1.AlbumService.FindAllAlbums:output_type -> proto.music_service.album.v1.FindAllAlbumsResponse
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_music_service_album_v1_album_service_proto_init() }
func file_proto_music_service_album_v1_album_service_proto_init() {
	if File_proto_music_service_album_v1_album_service_proto != nil {
		return
	}
	file_proto_music_service_album_v1_album_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_music_service_album_v1_album_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumResponse); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAlbumRequest); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullAlbum); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAlbumResponse); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindFullAlbumRequest); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindFullAlbumResponse); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAlbumsRequest); i {
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
		file_proto_music_service_album_v1_album_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAlbumsResponse); i {
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
			RawDescriptor: file_proto_music_service_album_v1_album_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_music_service_album_v1_album_service_proto_goTypes,
		DependencyIndexes: file_proto_music_service_album_v1_album_service_proto_depIdxs,
		MessageInfos:      file_proto_music_service_album_v1_album_service_proto_msgTypes,
	}.Build()
	File_proto_music_service_album_v1_album_service_proto = out.File
	file_proto_music_service_album_v1_album_service_proto_rawDesc = nil
	file_proto_music_service_album_v1_album_service_proto_goTypes = nil
	file_proto_music_service_album_v1_album_service_proto_depIdxs = nil
}
