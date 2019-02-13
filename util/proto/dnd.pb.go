// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dnd.proto

package dnd

import (
	fmt "fmt"
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

type Files struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Files) Reset()         { *m = Files{} }
func (m *Files) String() string { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()    {}
func (*Files) Descriptor() ([]byte, []int) {
	return fileDescriptor_80917f3bcd409d31, []int{0}
}

func (m *Files) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Files.Unmarshal(m, b)
}
func (m *Files) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Files.Marshal(b, m, deterministic)
}
func (m *Files) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Files.Merge(m, src)
}
func (m *Files) XXX_Size() int {
	return xxx_messageInfo_Files.Size(m)
}
func (m *Files) XXX_DiscardUnknown() {
	xxx_messageInfo_Files.DiscardUnknown(m)
}

var xxx_messageInfo_Files proto.InternalMessageInfo

func (m *Files) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Files)(nil), "dnd.Files")
}

func init() { proto.RegisterFile("dnd.proto", fileDescriptor_80917f3bcd409d31) }

var fileDescriptor_80917f3bcd409d31 = []byte{
	// 71 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xc9, 0x4b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xc9, 0x4b, 0x51, 0x92, 0xe6, 0x62, 0x75, 0xcb,
	0xcc, 0x49, 0x2d, 0x16, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae,
	0x11, 0xf6, 0x4b, 0x35, 0x00, 0x00, 0x00,
}