// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocs/service.proto

package packets

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Ack struct {
	Namespace            string                      `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Server               string                      `protobuf:"bytes,2,opt,name=Server,proto3" json:"Server,omitempty"`
	Checksum             string                      `protobuf:"bytes,3,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	ProtectedRoutes      map[string]*ProtectedRoutes `protobuf:"bytes,4,rep,name=ProtectedRoutes,proto3" json:"ProtectedRoutes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{0}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Ack) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Ack) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Ack) GetProtectedRoutes() map[string]*ProtectedRoutes {
	if m != nil {
		return m.ProtectedRoutes
	}
	return nil
}

type AckRequest struct {
	From                 string   `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckRequest) Reset()         { *m = AckRequest{} }
func (m *AckRequest) String() string { return proto.CompactTextString(m) }
func (*AckRequest) ProtoMessage()    {}
func (*AckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{1}
}

func (m *AckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckRequest.Unmarshal(m, b)
}
func (m *AckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckRequest.Marshal(b, m, deterministic)
}
func (m *AckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckRequest.Merge(m, src)
}
func (m *AckRequest) XXX_Size() int {
	return xxx_messageInfo_AckRequest.Size(m)
}
func (m *AckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AckRequest proto.InternalMessageInfo

func (m *AckRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type ProtectedRoute struct {
	IsStrict             bool     `protobuf:"varint,1,opt,name=IsStrict,proto3" json:"IsStrict,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=Method,proto3" json:"Method,omitempty"`
	Pattern              string   `protobuf:"bytes,2,opt,name=Pattern,proto3" json:"Pattern,omitempty"`
	Metas                *any.Any `protobuf:"bytes,4,opt,name=Metas,proto3" json:"Metas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtectedRoute) Reset()         { *m = ProtectedRoute{} }
func (m *ProtectedRoute) String() string { return proto.CompactTextString(m) }
func (*ProtectedRoute) ProtoMessage()    {}
func (*ProtectedRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{2}
}

func (m *ProtectedRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtectedRoute.Unmarshal(m, b)
}
func (m *ProtectedRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtectedRoute.Marshal(b, m, deterministic)
}
func (m *ProtectedRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtectedRoute.Merge(m, src)
}
func (m *ProtectedRoute) XXX_Size() int {
	return xxx_messageInfo_ProtectedRoute.Size(m)
}
func (m *ProtectedRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtectedRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ProtectedRoute proto.InternalMessageInfo

func (m *ProtectedRoute) GetIsStrict() bool {
	if m != nil {
		return m.IsStrict
	}
	return false
}

func (m *ProtectedRoute) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ProtectedRoute) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ProtectedRoute) GetMetas() *any.Any {
	if m != nil {
		return m.Metas
	}
	return nil
}

type ProtectedRoutes struct {
	Routes               []*ProtectedRoute `protobuf:"bytes,1,rep,name=Routes,proto3" json:"Routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProtectedRoutes) Reset()         { *m = ProtectedRoutes{} }
func (m *ProtectedRoutes) String() string { return proto.CompactTextString(m) }
func (*ProtectedRoutes) ProtoMessage()    {}
func (*ProtectedRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{3}
}

func (m *ProtectedRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtectedRoutes.Unmarshal(m, b)
}
func (m *ProtectedRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtectedRoutes.Marshal(b, m, deterministic)
}
func (m *ProtectedRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtectedRoutes.Merge(m, src)
}
func (m *ProtectedRoutes) XXX_Size() int {
	return xxx_messageInfo_ProtectedRoutes.Size(m)
}
func (m *ProtectedRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtectedRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_ProtectedRoutes proto.InternalMessageInfo

func (m *ProtectedRoutes) GetRoutes() []*ProtectedRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Request struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{4}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Server               string   `protobuf:"bytes,1,opt,name=Server,proto3" json:"Server,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7842d62b85457464, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Ack)(nil), "packets.Ack")
	proto.RegisterMapType((map[string]*ProtectedRoutes)(nil), "packets.Ack.ProtectedRoutesEntry")
	proto.RegisterType((*AckRequest)(nil), "packets.AckRequest")
	proto.RegisterType((*ProtectedRoute)(nil), "packets.ProtectedRoute")
	proto.RegisterType((*ProtectedRoutes)(nil), "packets.ProtectedRoutes")
	proto.RegisterType((*Request)(nil), "packets.Request")
	proto.RegisterType((*Response)(nil), "packets.Response")
}

func init() { proto.RegisterFile("protocs/service.proto", fileDescriptor_7842d62b85457464) }

var fileDescriptor_7842d62b85457464 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x8f, 0xd3, 0x3c,
	0x10, 0xdd, 0xf4, 0x77, 0xa7, 0xab, 0xef, 0x5b, 0x86, 0x05, 0x42, 0xc5, 0xa1, 0xe4, 0xb4, 0x42,
	0x22, 0x85, 0x22, 0x21, 0x04, 0xa7, 0xb6, 0x80, 0x58, 0xa1, 0x5d, 0x55, 0xee, 0x0d, 0x71, 0xf1,
	0xba, 0xc3, 0xa6, 0x4a, 0x1a, 0x07, 0xdb, 0xa9, 0x94, 0x2b, 0x47, 0xfe, 0x6a, 0x14, 0xc7, 0x2d,
	0x69, 0x59, 0x24, 0x6e, 0xf3, 0xc6, 0xcf, 0xe3, 0xf7, 0x9e, 0x07, 0x1e, 0x64, 0x4a, 0x1a, 0x29,
	0xf4, 0x58, 0x93, 0xda, 0xae, 0x05, 0x85, 0x16, 0x63, 0x37, 0xe3, 0x22, 0x26, 0xa3, 0x87, 0x8f,
	0x6f, 0xa5, 0xbc, 0x4d, 0x68, 0x6c, 0xdb, 0x37, 0xf9, 0xb7, 0x31, 0x4f, 0x8b, 0x8a, 0x13, 0xfc,
	0x68, 0x40, 0x73, 0x2a, 0x62, 0x7c, 0x02, 0xfd, 0x6b, 0xbe, 0x21, 0x9d, 0x71, 0x41, 0xbe, 0x37,
	0xf2, 0x2e, 0xfa, 0xec, 0x77, 0x03, 0x1f, 0x42, 0x67, 0x49, 0x6a, 0x4b, 0xca, 0x6f, 0xd8, 0x23,
	0x87, 0x70, 0x08, 0xbd, 0x79, 0x44, 0x22, 0xd6, 0xf9, 0xc6, 0x6f, 0xda, 0x93, 0x3d, 0xc6, 0xcf,
	0xf0, 0xff, 0x42, 0x49, 0x43, 0xc2, 0xd0, 0x8a, 0xc9, 0xdc, 0x90, 0xf6, 0x5b, 0xa3, 0xe6, 0xc5,
	0x60, 0xf2, 0x34, 0x74, 0xba, 0xc2, 0xa9, 0x88, 0xc3, 0x23, 0xce, 0x87, 0xd4, 0xa8, 0x82, 0x1d,
	0xdf, 0x1c, 0x7e, 0x85, 0xf3, 0xbb, 0x88, 0x78, 0x06, 0xcd, 0x98, 0x0a, 0x27, 0xb8, 0x2c, 0x31,
	0x84, 0xf6, 0x96, 0x27, 0x39, 0x59, 0xa5, 0x83, 0x89, 0xbf, 0x7f, 0xec, 0xe8, 0x3e, 0xab, 0x68,
	0x6f, 0x1b, 0x6f, 0xbc, 0x60, 0x04, 0x30, 0x15, 0x31, 0xa3, 0xef, 0x39, 0x69, 0x83, 0x08, 0xad,
	0x8f, 0x4a, 0x6e, 0x9c, 0x55, 0x5b, 0x07, 0x3f, 0x3d, 0xf8, 0xef, 0x70, 0x40, 0xe9, 0xfd, 0x52,
	0x2f, 0x8d, 0x5a, 0x0b, 0x63, 0xdf, 0xef, 0xb1, 0x3d, 0x2e, 0xf3, 0xba, 0x22, 0x13, 0xc9, 0x95,
	0x4b, 0xc5, 0x21, 0xf4, 0xa1, 0xbb, 0xe0, 0xc6, 0x90, 0x4a, 0xdd, 0xf4, 0x1d, 0xc4, 0x67, 0xd0,
	0xbe, 0x22, 0xc3, 0xcb, 0x8c, 0x4a, 0xd9, 0xe7, 0x61, 0xf5, 0x65, 0xe1, 0xee, 0xcb, 0xc2, 0x69,
	0x5a, 0xb0, 0x8a, 0x12, 0xcc, 0xfe, 0x48, 0x16, 0xc7, 0xd0, 0x71, 0x19, 0x7b, 0x36, 0xe3, 0x47,
	0x7f, 0xb1, 0xcd, 0x1c, 0x2d, 0xb8, 0x84, 0x6e, 0xcd, 0xef, 0x82, 0x9b, 0xc8, 0x85, 0x68, 0xeb,
	0x9a, 0x81, 0xc6, 0x81, 0x01, 0x84, 0xd6, 0x4c, 0xae, 0x0a, 0x6b, 0xeb, 0x94, 0xd9, 0x3a, 0xb8,
	0x86, 0x1e, 0x23, 0x9d, 0xc9, 0x54, 0xd7, 0x17, 0xc5, 0x3b, 0x58, 0x94, 0xb2, 0x6f, 0xb8, 0xc9,
	0xb5, 0x9d, 0xd7, 0x66, 0x0e, 0xdd, 0x35, 0x6f, 0x92, 0x42, 0x77, 0x59, 0xed, 0x31, 0xbe, 0x80,
	0xfe, 0x27, 0x9e, 0xae, 0x74, 0xc4, 0x63, 0xc2, 0xfb, 0xf5, 0xbd, 0x71, 0xe2, 0x87, 0xa7, 0xf5,
	0x66, 0x70, 0x82, 0x2f, 0xa1, 0xf7, 0x7e, 0xad, 0x33, 0x6e, 0x44, 0x84, 0x67, 0xfb, 0xb3, 0x1d,
	0xfb, 0x5e, 0xad, 0x53, 0x29, 0x0e, 0x4e, 0x26, 0xaf, 0xa1, 0x3d, 0x4f, 0x68, 0xcb, 0xf1, 0x39,
	0xb4, 0xe6, 0x3c, 0x49, 0xfe, 0xf1, 0xde, 0x6c, 0xf0, 0xa5, 0x1f, 0xbe, 0x73, 0xfd, 0x9b, 0x8e,
	0xfd, 0xa8, 0x57, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x69, 0x42, 0xc8, 0x8b, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handshake(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*Ack, error)
	Dispatch(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handshake(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/packets.Service/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Dispatch(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/packets.Service/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handshake(context.Context, *AckRequest) (*Ack, error)
	Dispatch(context.Context, *Request) (*Response, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handshake(ctx context.Context, req *AckRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (*UnimplementedServiceServer) Dispatch(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.Service/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handshake(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.Service/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Dispatch(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packets.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Service_Handshake_Handler,
		},
		{
			MethodName: "Dispatch",
			Handler:    _Service_Dispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocs/service.proto",
}

// ClevaClient is the client API for Cleva service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClevaClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type clevaClient struct {
	cc *grpc.ClientConn
}

func NewClevaClient(cc *grpc.ClientConn) ClevaClient {
	return &clevaClient{cc}
}

func (c *clevaClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/packets.Cleva/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClevaServer is the server API for Cleva service.
type ClevaServer interface {
	Call(context.Context, *Request) (*Response, error)
}

// UnimplementedClevaServer can be embedded to have forward compatible implementations.
type UnimplementedClevaServer struct {
}

func (*UnimplementedClevaServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterClevaServer(s *grpc.Server, srv ClevaServer) {
	s.RegisterService(&_Cleva_serviceDesc, srv)
}

func _Cleva_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClevaServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.Cleva/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClevaServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cleva_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packets.Cleva",
	HandlerType: (*ClevaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Cleva_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocs/service.proto",
}
