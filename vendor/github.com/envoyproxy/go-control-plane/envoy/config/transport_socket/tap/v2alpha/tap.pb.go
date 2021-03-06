// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/transport_socket/tap/v2alpha/tap.proto

package envoy_config_transport_socket_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type Tap struct {
	CommonConfig         *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	TransportSocket      *core.TransportSocket          `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_07cb8c0b42756e40, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.transport_socket.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/tap/v2alpha/tap.proto", fileDescriptor_07cb8c0b42756e40)
}

var fileDescriptor_07cb8c0b42756e40 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x66, 0x2b, 0x95, 0x12, 0x15, 0x4b, 0x2e, 0x96, 0x22, 0x22, 0x3d, 0x29, 0xca, 0x2e, 0x24,
	0xa0, 0xf7, 0x14, 0xef, 0x45, 0x03, 0x1e, 0xcb, 0x34, 0x5d, 0xeb, 0x62, 0xb3, 0xb3, 0xec, 0xae,
	0x4b, 0xfb, 0x0a, 0x5e, 0xbc, 0xfa, 0x42, 0xbe, 0x90, 0x47, 0x0f, 0x22, 0xd9, 0x4d, 0xa0, 0x89,
	0x17, 0x6f, 0xc3, 0xcc, 0xf7, 0x33, 0xf3, 0x4d, 0x94, 0x72, 0xe9, 0x70, 0xcb, 0x0a, 0x94, 0x4f,
	0x62, 0xc5, 0xac, 0x06, 0x69, 0x14, 0x6a, 0x3b, 0x37, 0x58, 0xbc, 0x70, 0xcb, 0x2c, 0x28, 0xe6,
	0x12, 0x58, 0xab, 0x67, 0xa8, 0x6a, 0xaa, 0x34, 0x5a, 0x8c, 0x2f, 0x3d, 0x89, 0x06, 0x12, 0xed,
	0x92, 0x68, 0x05, 0xac, 0x49, 0xe3, 0xd3, 0xa0, 0x0f, 0x4a, 0x30, 0x97, 0xb0, 0x02, 0x35, 0x67,
	0x0b, 0x30, 0x3c, 0x08, 0x8d, 0xaf, 0x5b, 0xee, 0x05, 0x96, 0x25, 0xca, 0x96, 0x67, 0x68, 0xd5,
	0xe8, 0xb3, 0xd7, 0xa5, 0x02, 0x06, 0x52, 0xa2, 0x05, 0x2b, 0x50, 0x1a, 0x56, 0x8a, 0x95, 0x06,
	0xdb, 0xa8, 0x9d, 0x38, 0x58, 0x8b, 0x25, 0x58, 0xce, 0x9a, 0x22, 0x0c, 0x26, 0x9f, 0x24, 0xda,
	0xcb, 0x41, 0xc5, 0x3c, 0x3a, 0x0a, 0x82, 0xf3, 0xe0, 0x38, 0x22, 0xe7, 0xe4, 0xe2, 0x20, 0xb9,
	0xa1, 0xad, 0x7b, 0x6a, 0xcf, 0x9d, 0x2b, 0xe8, 0xd4, 0xb7, 0xee, 0x36, 0x96, 0x4b, 0x23, 0x50,
	0x4e, 0x3d, 0x30, 0x1b, 0x7c, 0x67, 0xfd, 0x37, 0xd2, 0x1b, 0x92, 0xfb, 0xc3, 0xc0, 0x09, 0xfd,
	0xf8, 0x31, 0x1a, 0x76, 0x33, 0x19, 0xf5, 0xbc, 0xd3, 0xa4, 0x76, 0x02, 0x25, 0xa8, 0x4b, 0x68,
	0x15, 0x07, 0xcd, 0x1b, 0xe8, 0x83, 0x47, 0xee, 0xa8, 0x1e, 0xdb, 0xce, 0xc8, 0x7d, 0x7d, 0xfc,
	0xbc, 0xf7, 0xaf, 0x9a, 0xfc, 0x79, 0xb3, 0x8f, 0xf9, 0xf3, 0x03, 0x13, 0xd6, 0x4f, 0xa3, 0x5b,
	0x81, 0xc1, 0x53, 0x69, 0xdc, 0x6c, 0xe9, 0xbf, 0x1f, 0x97, 0x0d, 0x72, 0x50, 0xb3, 0x2a, 0xbd,
	0x19, 0x59, 0xec, 0xfb, 0x18, 0xd3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x21, 0x73, 0xaf,
	0x2d, 0x02, 0x00, 0x00,
}
