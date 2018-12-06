// Code generated by protoc-gen-go. DO NOT EDIT.
// source: expert/crud-api-using-mongodb/proto/informatica.proto

package informaticapb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Informatica struct {
	Sequence             int32    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	HostName             string   `protobuf:"bytes,4,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Informatica) Reset()         { *m = Informatica{} }
func (m *Informatica) String() string { return proto.CompactTextString(m) }
func (*Informatica) ProtoMessage()    {}
func (*Informatica) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3c05494cc022a7, []int{0}
}

func (m *Informatica) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Informatica.Unmarshal(m, b)
}
func (m *Informatica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Informatica.Marshal(b, m, deterministic)
}
func (m *Informatica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Informatica.Merge(m, src)
}
func (m *Informatica) XXX_Size() int {
	return xxx_messageInfo_Informatica.Size(m)
}
func (m *Informatica) XXX_DiscardUnknown() {
	xxx_messageInfo_Informatica.DiscardUnknown(m)
}

var xxx_messageInfo_Informatica proto.InternalMessageInfo

func (m *Informatica) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Informatica) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Informatica) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Informatica) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func init() {
	proto.RegisterType((*Informatica)(nil), "informatica.Informatica")
}

func init() {
	proto.RegisterFile("expert/crud-api-using-mongodb/proto/informatica.proto", fileDescriptor_be3c05494cc022a7)
}

var fileDescriptor_be3c05494cc022a7 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xad, 0x28, 0x48,
	0x2d, 0x2a, 0xd1, 0x4f, 0x2e, 0x2a, 0x4d, 0xd1, 0x4d, 0x2c, 0xc8, 0xd4, 0x2d, 0x2d, 0xce, 0xcc,
	0x4b, 0xd7, 0xcd, 0xcd, 0xcf, 0x4b, 0xcf, 0x4f, 0x49, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0xcf, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0x4d, 0x2c, 0xc9, 0x4c, 0x4e, 0xd4, 0x03, 0x8b, 0x08, 0x71,
	0x23, 0x09, 0x29, 0x15, 0x70, 0x71, 0x7b, 0x22, 0xb8, 0x42, 0x52, 0x5c, 0x1c, 0xc5, 0xa9, 0x85,
	0xa5, 0xa9, 0x79, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x70, 0xbe, 0x90, 0x08,
	0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23,
	0x24, 0xc4, 0xc5, 0x02, 0x32, 0x4f, 0x82, 0x19, 0x2c, 0x08, 0x66, 0x0b, 0x49, 0x73, 0x71, 0x66,
	0xe4, 0x17, 0x97, 0xc4, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0xb0, 0x80, 0x25, 0x38, 0x40, 0x02, 0x7e,
	0x89, 0xb9, 0xa9, 0x46, 0x22, 0x5c, 0x42, 0x48, 0x36, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x3a, 0xf1, 0x47, 0xf1, 0x22, 0x39, 0xab, 0x20, 0x29, 0x89, 0x0d, 0xec, 0x58, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x55, 0x02, 0x63, 0xd8, 0xe5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InformaticaServiceClient is the client API for InformaticaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InformaticaServiceClient interface {
}

type informaticaServiceClient struct {
	cc *grpc.ClientConn
}

func NewInformaticaServiceClient(cc *grpc.ClientConn) InformaticaServiceClient {
	return &informaticaServiceClient{cc}
}

// InformaticaServiceServer is the server API for InformaticaService service.
type InformaticaServiceServer interface {
}

func RegisterInformaticaServiceServer(s *grpc.Server, srv InformaticaServiceServer) {
	s.RegisterService(&_InformaticaService_serviceDesc, srv)
}

var _InformaticaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "informatica.InformaticaService",
	HandlerType: (*InformaticaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "expert/crud-api-using-mongodb/proto/informatica.proto",
}