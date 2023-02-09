// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/music_service/person/v1/person_service.proto

package pb_music_album

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	FindAllPersons(ctx context.Context, in *FindAllPersonsRequest, opts ...grpc.CallOption) (*FindAllPersonsResponse, error)
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) FindAllPersons(ctx context.Context, in *FindAllPersonsRequest, opts ...grpc.CallOption) (*FindAllPersonsResponse, error) {
	out := new(FindAllPersonsResponse)
	err := c.cc.Invoke(ctx, "/proto.music_service.person.v1.PersonService/FindAllPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error) {
	out := new(CreatePersonResponse)
	err := c.cc.Invoke(ctx, "/proto.music_service.person.v1.PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	FindAllPersons(context.Context, *FindAllPersonsRequest) (*FindAllPersonsResponse, error)
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error)
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) FindAllPersons(context.Context, *FindAllPersonsRequest) (*FindAllPersonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllPersons not implemented")
}
func (UnimplementedPersonServiceServer) CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_FindAllPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).FindAllPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.music_service.person.v1.PersonService/FindAllPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).FindAllPersons(ctx, req.(*FindAllPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.music_service.person.v1.PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.music_service.person.v1.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllPersons",
			Handler:    _PersonService_FindAllPersons_Handler,
		},
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/music_service/person/v1/person_service.proto",
}