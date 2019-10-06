// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/health/health.proto

package health

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
	3: "SERVICE_UNKNOWN",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":         0,
	"SERVING":         1,
	"NOT_SERVING":     2,
	"SERVICE_UNKNOWN": 3,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7bb2e9a145ccb86d, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb2e9a145ccb86d, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb2e9a145ccb86d, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("health.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "health.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "health.HealthCheckResponse")
}

func init() { proto.RegisterFile("proto/health/health.proto", fileDescriptor_7bb2e9a145ccb86d) }

var fileDescriptor_7bb2e9a145ccb86d = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x4e, 0xc2, 0x40,
	0x14, 0xb5, 0x10, 0x4a, 0xfc, 0x88, 0x90, 0xc1, 0x18, 0xac, 0x2e, 0x48, 0x57, 0x9a, 0xd8, 0x0e,
	0xe0, 0x09, 0x80, 0x10, 0x35, 0x9a, 0x92, 0x14, 0x94, 0xc4, 0x0d, 0x29, 0x65, 0x32, 0x6d, 0xac,
	0x9d, 0xda, 0x99, 0x42, 0xdc, 0x7a, 0x04, 0xbd, 0x81, 0x3b, 0xcf, 0xe3, 0x15, 0x3c, 0x88, 0x61,
	0xda, 0x2e, 0x88, 0xc4, 0x85, 0xab, 0x9f, 0xf7, 0xff, 0x7b, 0x6f, 0xde, 0xe4, 0x7f, 0x38, 0x8a,
	0x62, 0x26, 0x18, 0xf6, 0x88, 0x13, 0x08, 0x2f, 0x2b, 0xa6, 0xec, 0x21, 0x35, 0x45, 0xda, 0x09,
	0x65, 0x8c, 0x06, 0x04, 0x3b, 0x91, 0x8f, 0x9d, 0x30, 0x64, 0xc2, 0x11, 0x3e, 0x0b, 0x79, 0xca,
	0xd2, 0xce, 0x65, 0x71, 0x0d, 0x4a, 0x42, 0x83, 0xaf, 0x1c, 0x4a, 0x49, 0x8c, 0x59, 0x24, 0x19,
	0xbf, 0xd9, 0xba, 0x09, 0xe8, 0x4a, 0xba, 0x0e, 0x3c, 0xe2, 0x3e, 0xda, 0xe4, 0x39, 0x21, 0x5c,
	0xa0, 0x26, 0x94, 0x39, 0x89, 0x97, 0xbe, 0x4b, 0x9a, 0x4a, 0x4b, 0x39, 0xdd, 0xb5, 0x73, 0xa8,
	0x7f, 0x2a, 0xd0, 0xd8, 0x10, 0xf0, 0x88, 0x85, 0x9c, 0xa0, 0x1e, 0xa8, 0x5c, 0x38, 0x22, 0xe1,
	0x52, 0xb0, 0xdf, 0x3d, 0x33, 0xb3, 0xe8, 0x5b, 0xc8, 0xe6, 0x78, 0x6d, 0x16, 0xd2, 0xb1, 0x14,
	0xd8, 0x99, 0x50, 0x1f, 0x41, 0x75, 0x63, 0x80, 0x2a, 0x50, 0xbe, 0xb3, 0x6e, 0xac, 0xd1, 0xd4,
	0xaa, 0xef, 0xac, 0xc1, 0x78, 0x68, 0xdf, 0x5f, 0x5b, 0x97, 0x75, 0x05, 0xd5, 0xa0, 0x62, 0x8d,
	0x26, 0xb3, 0xbc, 0x51, 0x40, 0x0d, 0xa8, 0x49, 0x30, 0x18, 0xce, 0x72, 0x49, 0xb1, 0xfb, 0xa1,
	0x80, 0x9a, 0x3e, 0x8f, 0x26, 0x50, 0x92, 0x11, 0x90, 0xb6, 0x35, 0x97, 0xfc, 0xb5, 0x76, 0xfc,
	0x47, 0x66, 0x1d, 0xbd, 0x7e, 0x7d, 0xbf, 0x17, 0xf6, 0x10, 0xe0, 0x65, 0x27, 0x5b, 0x0b, 0xea,
	0x43, 0x69, 0xea, 0x08, 0xd7, 0xfb, 0xb7, 0x6b, 0x5b, 0xe9, 0xdf, 0xbe, 0xf5, 0x0e, 0xd1, 0x01,
	0x54, 0xd3, 0x69, 0x4b, 0x8e, 0x79, 0xb7, 0xd8, 0x31, 0xdb, 0x0f, 0x1d, 0xea, 0x0b, 0x2f, 0x99,
	0x9b, 0x2e, 0x7b, 0xc2, 0x2f, 0x2c, 0x89, 0x63, 0x12, 0xad, 0x0f, 0x23, 0x08, 0x98, 0xb1, 0x62,
	0x71, 0xb0, 0x30, 0xfc, 0x45, 0x80, 0x29, 0x09, 0x31, 0xcd, 0xef, 0x65, 0xae, 0xca, 0xad, 0x5e,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xfa, 0x76, 0x68, 0x46, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (Health_WatchClient, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (Health_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Health_serviceDesc.Streams[0], "/health.Health/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpc.ClientStream
}

type healthWatchClient struct {
	grpc.ClientStream
}

func (x *healthWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Watch(*HealthCheckRequest, Health_WatchServer) error
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedHealthServer) Watch(req *HealthCheckRequest, srv Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).Watch(m, &healthWatchServer{stream})
}

type Health_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpc.ServerStream
}

type healthWatchServer struct {
	grpc.ServerStream
}

func (x *healthWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Health_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/health/health.proto",
}
