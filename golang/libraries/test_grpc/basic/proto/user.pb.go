// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package tencent_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	Session
	Result
	User
	AddUserRequest
	AddUserResponse
*/
package tencent_user

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

// The request session
type Session struct {
	Seq string `protobuf:"bytes,1,opt,name=seq" json:"seq,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Session) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

// The response result
type Result struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Sex  int32  `protobuf:"varint,2,opt,name=sex" json:"sex,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type AddUserRequest struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	User    *User    `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddUserRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AddUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type AddUserResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Result  *Result  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *AddUserResponse) Reset()                    { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()               {}
func (*AddUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddUserResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AddUserResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "tencent.user.Session")
	proto.RegisterType((*Result)(nil), "tencent.user.Result")
	proto.RegisterType((*User)(nil), "tencent.user.User")
	proto.RegisterType((*AddUserRequest)(nil), "tencent.user.AddUserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "tencent.user.AddUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	// Sends a greeting
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := grpc.Invoke(ctx, "/tencent.user.UserService/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	// Sends a greeting
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencent.user.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tencent.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xad, 0xa6, 0x09, 0x4e, 0x44, 0x65, 0x50, 0x08, 0xfe, 0x01, 0xd9, 0x83, 0x78, 0x28,
	0x11, 0x22, 0x78, 0xf7, 0xe6, 0x79, 0x8b, 0x78, 0xae, 0xc9, 0x20, 0x01, 0xbb, 0x9b, 0xee, 0x6c,
	0xa4, 0x1f, 0x5f, 0x76, 0xb2, 0x15, 0xf7, 0xe0, 0xa1, 0xb7, 0x49, 0xde, 0x7b, 0xf3, 0x7b, 0xcc,
	0x02, 0x8c, 0x4c, 0xae, 0x1e, 0x9c, 0xf5, 0x16, 0x4f, 0x3c, 0x99, 0x96, 0x8c, 0xaf, 0xc3, 0x3f,
	0x75, 0x0d, 0xc5, 0x92, 0x98, 0x7b, 0x6b, 0xf0, 0x1c, 0x8e, 0x98, 0x36, 0xd5, 0xec, 0x6e, 0xf6,
	0x70, 0xac, 0xc3, 0xa8, 0x9e, 0x21, 0xd7, 0xc4, 0xe3, 0x97, 0x47, 0x84, 0xac, 0xb5, 0x1d, 0x89,
	0x38, 0xd7, 0x32, 0x63, 0x05, 0xc5, 0x9a, 0x98, 0x57, 0x9f, 0x54, 0x1d, 0x4a, 0x66, 0xf7, 0xa9,
	0x16, 0x90, 0xbd, 0x31, 0xb9, 0x90, 0x32, 0xab, 0x35, 0xc5, 0x95, 0x32, 0x4f, 0x94, 0xad, 0x24,
	0xe6, 0x81, 0xb2, 0x55, 0x3d, 0x9c, 0xbe, 0x74, 0x5d, 0x08, 0x68, 0xda, 0x8c, 0xc4, 0x1e, 0x1f,
	0xa1, 0xe0, 0xa9, 0x94, 0x44, 0xcb, 0xe6, 0xb2, 0xfe, 0x5b, 0xba, 0x8e, 0x8d, 0xf5, 0xce, 0x85,
	0xf7, 0x90, 0x05, 0x41, 0xb6, 0x96, 0x0d, 0xa6, 0x6e, 0xd9, 0x2c, 0xba, 0x1a, 0xe0, 0xec, 0x17,
	0xc5, 0x83, 0x35, 0x4c, 0xfb, 0xb3, 0x16, 0x90, 0x3b, 0x39, 0x4a, 0xa4, 0x5d, 0xa4, 0xfe, 0xe9,
	0x60, 0x3a, 0x7a, 0x9a, 0x77, 0x28, 0x03, 0x6e, 0x49, 0xee, 0xbb, 0x6f, 0x09, 0x5f, 0xa1, 0x88,
	0x05, 0xf0, 0x26, 0xcd, 0xa5, 0x27, 0xb8, 0xba, 0xfd, 0x47, 0x9d, 0x5a, 0xab, 0x83, 0x8f, 0x5c,
	0x5e, 0xf3, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x59, 0x96, 0x02, 0x75, 0xdb, 0x01, 0x00, 0x00,
}
