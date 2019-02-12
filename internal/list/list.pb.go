// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package list

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

type File struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_9170fc6eca7a456d, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Path struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_9170fc6eca7a456d, []int{1}
}
func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (dst *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(dst, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type FileList struct {
	Files                []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileList) Reset()         { *m = FileList{} }
func (m *FileList) String() string { return proto.CompactTextString(m) }
func (*FileList) ProtoMessage()    {}
func (*FileList) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_9170fc6eca7a456d, []int{2}
}
func (m *FileList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileList.Unmarshal(m, b)
}
func (m *FileList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileList.Marshal(b, m, deterministic)
}
func (dst *FileList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileList.Merge(dst, src)
}
func (m *FileList) XXX_Size() int {
	return xxx_messageInfo_FileList.Size(m)
}
func (m *FileList) XXX_DiscardUnknown() {
	xxx_messageInfo_FileList.DiscardUnknown(m)
}

var xxx_messageInfo_FileList proto.InternalMessageInfo

func (m *FileList) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "list.File")
	proto.RegisterType((*Path)(nil), "list.Path")
	proto.RegisterType((*FileList)(nil), "list.FileList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilesClient interface {
	List(ctx context.Context, in *Path, opts ...grpc.CallOption) (*FileList, error)
}

type filesClient struct {
	cc *grpc.ClientConn
}

func NewFilesClient(cc *grpc.ClientConn) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) List(ctx context.Context, in *Path, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/list.Files/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesServer is the server API for Files service.
type FilesServer interface {
	List(context.Context, *Path) (*FileList, error)
}

func RegisterFilesServer(s *grpc.Server, srv FilesServer) {
	s.RegisterService(&_Files_serviceDesc, srv)
}

func _Files_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Path)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/list.Files/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).List(ctx, req.(*Path))
	}
	return interceptor(ctx, in, info, handler)
}

var _Files_serviceDesc = grpc.ServiceDesc{
	ServiceName: "list.Files",
	HandlerType: (*FilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Files_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "list.proto",
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_list_9170fc6eca7a456d) }

var fileDescriptor_list_9170fc6eca7a456d = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x2c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb8, 0x58, 0xdc, 0x32,
	0x73, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x90, 0x5c, 0x40, 0x62, 0x49, 0x06, 0x56, 0x39, 0x1d, 0x2e, 0x0e, 0x90, 0x3e,
	0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x05, 0x2e, 0xd6, 0xb4, 0xcc, 0x9c, 0xd4, 0x62, 0x09, 0x46, 0x05,
	0x66, 0x0d, 0x6e, 0x23, 0x2e, 0x3d, 0xb0, 0x2d, 0x20, 0xe9, 0x20, 0x88, 0x84, 0x91, 0x2e, 0x17,
	0x2b, 0x88, 0x5b, 0x2c, 0xa4, 0xc2, 0xc5, 0x02, 0xd6, 0x02, 0x55, 0x03, 0x32, 0x5e, 0x8a, 0x0f,
	0xa1, 0x1e, 0x24, 0xa7, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0xe5, 0xfb, 0x9c, 0xaf, 0x00, 0x00, 0x00,
}
