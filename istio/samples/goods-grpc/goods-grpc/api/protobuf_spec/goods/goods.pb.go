// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf_spec/goods/goods.proto

/*
Package goods is a generated protocol buffer package.

It is generated from these files:
	api/protobuf_spec/goods/goods.proto

It has these top-level messages:
	GetGoodRequest
	GetGoodResponse
*/
package goods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import comments "goods-grpc/api/protobuf_spec/comments"

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

type GetGoodRequest struct {
	GoodID int64 `protobuf:"varint,1,opt,name=goodID" json:"goodID,omitempty"`
}

func (m *GetGoodRequest) Reset()                    { *m = GetGoodRequest{} }
func (m *GetGoodRequest) String() string            { return proto.CompactTextString(m) }
func (*GetGoodRequest) ProtoMessage()               {}
func (*GetGoodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetGoodRequest) GetGoodID() int64 {
	if m != nil {
		return m.GoodID
	}
	return 0
}

type GetGoodResponse struct {
	GoodID    int64                   `protobuf:"varint,1,opt,name=goodID" json:"goodID,omitempty"`
	GoodStock int64                   `protobuf:"varint,2,opt,name=goodStock" json:"goodStock,omitempty"`
	Comments  []*comments.CommentInfo `protobuf:"bytes,3,rep,name=comments" json:"comments,omitempty"`
}

func (m *GetGoodResponse) Reset()                    { *m = GetGoodResponse{} }
func (m *GetGoodResponse) String() string            { return proto.CompactTextString(m) }
func (*GetGoodResponse) ProtoMessage()               {}
func (*GetGoodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetGoodResponse) GetGoodID() int64 {
	if m != nil {
		return m.GoodID
	}
	return 0
}

func (m *GetGoodResponse) GetGoodStock() int64 {
	if m != nil {
		return m.GoodStock
	}
	return 0
}

func (m *GetGoodResponse) GetComments() []*comments.CommentInfo {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGoodRequest)(nil), "goods.GetGoodRequest")
	proto.RegisterType((*GetGoodResponse)(nil), "goods.GetGoodResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoodsService service

type GoodsServiceClient interface {
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
}

type goodsServiceClient struct {
	cc *grpc.ClientConn
}

func NewGoodsServiceClient(cc *grpc.ClientConn) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := grpc.Invoke(ctx, "/goods.GoodsService/GetGood", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoodsService service

type GoodsServiceServer interface {
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
}

func RegisterGoodsServiceServer(s *grpc.Server, srv GoodsServiceServer) {
	s.RegisterService(&_GoodsService_serviceDesc, srv)
}

func _GoodsService_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.GoodsService/GetGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoodsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goods.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGood",
			Handler:    _GoodsService_GetGood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf_spec/goods/goods.proto",
}

func init() { proto.RegisterFile("api/protobuf_spec/goods/goods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0x8b, 0x2f, 0x2e, 0x48, 0x4d, 0xd6, 0x4f, 0xcf,
	0xcf, 0x4f, 0x29, 0x86, 0x90, 0x7a, 0x60, 0x19, 0x21, 0x56, 0x30, 0x47, 0xca, 0x04, 0x4c, 0xe9,
	0xa6, 0x17, 0x15, 0x24, 0xeb, 0x63, 0x6a, 0x4b, 0xce, 0xcf, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x86,
	0x33, 0x20, 0x9a, 0x95, 0x34, 0xb8, 0xf8, 0xdc, 0x53, 0x4b, 0xdc, 0xf3, 0xf3, 0x53, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x40, 0x26, 0x79, 0xba, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x30, 0x07, 0x41, 0x79, 0x4a, 0x55, 0x5c, 0xfc, 0x70, 0x95, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0xb8, 0x94, 0x0a, 0xc9, 0x70, 0x71, 0x82, 0x58, 0xc1, 0x25, 0xf9, 0xc9, 0xd9, 0x12,
	0x4c, 0x60, 0x29, 0x84, 0x80, 0x90, 0x21, 0x17, 0x07, 0xcc, 0x11, 0x12, 0xcc, 0x0a, 0xcc, 0x1a,
	0xdc, 0x46, 0xa2, 0x7a, 0x70, 0x57, 0x39, 0x43, 0x18, 0x9e, 0x79, 0x69, 0xf9, 0x41, 0x70, 0x65,
	0x46, 0x5e, 0x5c, 0x3c, 0x20, 0x8b, 0x8b, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xac,
	0xb8, 0xd8, 0xa1, 0x6e, 0x11, 0x12, 0xd5, 0x83, 0x84, 0x05, 0xaa, 0x2f, 0xa4, 0xc4, 0xd0, 0x85,
	0x21, 0x4e, 0x56, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x8e, 0x01, 0xec, 0x5c, 0x01, 0x00, 0x00,
}