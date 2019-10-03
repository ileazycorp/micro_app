// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro_booking/protos/booking.proto

package micro_booking_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_795a5e335392c5fb, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

type HelloResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_795a5e335392c5fb, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HelloRequest)(nil), "micro_booking.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "micro_booking.HelloResponse")
}

func init() { proto.RegisterFile("micro_booking/protos/booking.proto", fileDescriptor_795a5e335392c5fb) }

var fileDescriptor_795a5e335392c5fb = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xcd, 0x4c, 0x2e,
	0xca, 0x8f, 0x4f, 0xca, 0xcf, 0xcf, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f,
	0xd6, 0x87, 0x72, 0xf5, 0xc0, 0x5c, 0x21, 0x5e, 0x14, 0x35, 0x4a, 0x7c, 0x5c, 0x3c, 0x1e, 0xa9,
	0x39, 0x39, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xfc, 0x5c, 0xbc, 0x50, 0x7e,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x51, 0x08, 0x17, 0x9f, 0x13, 0x44, 0x6d, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x90, 0x13, 0x17, 0x2b, 0x58, 0x89, 0x90, 0xb4, 0x1e, 0x8a, 0x59, 0x7a,
	0xc8, 0x06, 0x49, 0xc9, 0x60, 0x97, 0x84, 0x98, 0xea, 0x24, 0x14, 0x25, 0x80, 0x22, 0x1d, 0x5f,
	0x90, 0x94, 0xc4, 0x06, 0x76, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x75, 0xc6, 0x9d, 0x24,
	0xc6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookingServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type bookingServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookingServiceClient(cc *grpc.ClientConn) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/micro_booking.BookingService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
type BookingServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedBookingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (*UnimplementedBookingServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterBookingServiceServer(s *grpc.Server, srv BookingServiceServer) {
	s.RegisterService(&_BookingService_serviceDesc, srv)
}

func _BookingService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/micro_booking.BookingService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "micro_booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _BookingService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro_booking/protos/booking.proto",
}