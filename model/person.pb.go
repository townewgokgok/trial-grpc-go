// Code generated by protoc-gen-go.
// source: person.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	IdRequest
	Person
	AddressBook
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type IdRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *IdRequest) Reset()                    { *m = IdRequest{} }
func (m *IdRequest) String() string            { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()               {}
func (*IdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Person struct {
	Name  string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phone" json:"phone,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhone() []*Person_PhoneNumber {
	if m != nil {
		return m.Phone
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=model.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*IdRequest)(nil), "model.IdRequest")
	proto.RegisterType((*Person)(nil), "model.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "model.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "model.AddressBook")
	proto.RegisterEnum("model.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrialGrpc service

type TrialGrpcClient interface {
	// Get a person
	GetPerson(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Person, error)
}

type trialGrpcClient struct {
	cc *grpc.ClientConn
}

func NewTrialGrpcClient(cc *grpc.ClientConn) TrialGrpcClient {
	return &trialGrpcClient{cc}
}

func (c *trialGrpcClient) GetPerson(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := grpc.Invoke(ctx, "/model.TrialGrpc/GetPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrialGrpc service

type TrialGrpcServer interface {
	// Get a person
	GetPerson(context.Context, *IdRequest) (*Person, error)
}

func RegisterTrialGrpcServer(s *grpc.Server, srv TrialGrpcServer) {
	s.RegisterService(&_TrialGrpc_serviceDesc, srv)
}

func _TrialGrpc_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialGrpcServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.TrialGrpc/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialGrpcServer).GetPerson(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrialGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.TrialGrpc",
	HandlerType: (*TrialGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _TrialGrpc_GetPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}

func init() { proto.RegisterFile("person.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0xd2, 0x24, 0x98, 0xa9, 0x96, 0x30, 0x88, 0xc6, 0x7a, 0x29, 0x01, 0xa1, 0x50, 0x88,
	0x10, 0xbd, 0x79, 0xb2, 0x50, 0x6a, 0xd1, 0x9a, 0xb2, 0x14, 0x3c, 0xa7, 0xee, 0x80, 0xc1, 0x24,
	0xbb, 0x6e, 0xd2, 0x43, 0xff, 0xdb, 0x0f, 0x90, 0x6c, 0xd6, 0x62, 0xc1, 0xdb, 0x9b, 0x79, 0x6f,
	0xde, 0x63, 0x66, 0xe0, 0x54, 0x92, 0xaa, 0x45, 0x15, 0x4b, 0x25, 0x1a, 0x81, 0x6e, 0x29, 0x38,
	0x15, 0xd1, 0x35, 0xf8, 0x4b, 0xce, 0xe8, 0x6b, 0x47, 0x75, 0x83, 0x43, 0xb0, 0x73, 0x1e, 0x5a,
	0x63, 0x6b, 0xe2, 0x32, 0x3b, 0xe7, 0xd1, 0xb7, 0x05, 0xde, 0x5a, 0x0f, 0x21, 0x82, 0x53, 0x65,
	0x25, 0x69, 0xd2, 0x67, 0x1a, 0x1b, 0xb9, 0xfd, 0x2b, 0xc7, 0x73, 0x70, 0xa9, 0xcc, 0xf2, 0x22,
	0xec, 0x6b, 0x51, 0x57, 0xe0, 0x2d, 0xb8, 0xf2, 0x43, 0x54, 0x14, 0x3a, 0xe3, 0xfe, 0x64, 0x90,
	0x5c, 0xc5, 0x3a, 0x38, 0xee, 0x7c, 0xe3, 0x75, 0x4b, 0xbd, 0xee, 0xca, 0x2d, 0x29, 0xd6, 0xe9,
	0x46, 0x0c, 0x06, 0x7f, 0xba, 0x78, 0x01, 0x5e, 0xa5, 0x91, 0xc9, 0x36, 0x15, 0x4e, 0xc1, 0x69,
	0xf6, 0x92, 0x74, 0xfe, 0x30, 0xb9, 0xfc, 0xc7, 0x76, 0xb3, 0x97, 0xc4, 0xb4, 0x28, 0x9a, 0x82,
	0x7f, 0x68, 0x21, 0x80, 0xb7, 0x4a, 0x67, 0xcb, 0x97, 0x79, 0xd0, 0xc3, 0x13, 0x70, 0x9e, 0xd2,
	0xd5, 0x3c, 0xb0, 0x5a, 0xf4, 0x96, 0xb2, 0xe7, 0xc0, 0x8e, 0xee, 0x61, 0xf0, 0xc8, 0xb9, 0xa2,
	0xba, 0x9e, 0x09, 0xf1, 0x89, 0x37, 0xe0, 0x49, 0x12, 0xb2, 0x68, 0x97, 0x6f, 0x37, 0x38, 0x3b,
	0x8a, 0x62, 0x86, 0x4c, 0x1e, 0xc0, 0xdf, 0xa8, 0x3c, 0x2b, 0x16, 0x4a, 0xbe, 0x63, 0x0c, 0xfe,
	0x82, 0x1a, 0x73, 0xbb, 0xc0, 0x0c, 0x1c, 0x0e, 0x3d, 0x3a, 0xb6, 0x88, 0x7a, 0x5b, 0x4f, 0x3f,
	0xe5, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x82, 0x7b, 0x75, 0xa4, 0x01, 0x00, 0x00,
}
