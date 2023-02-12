// Code generated by protoc-gen-go. DO NOT EDIT.
// source: excelize.proto

package excelize

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

type TestEntry struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Lp                   string   `protobuf:"bytes,2,opt,name=lp,proto3" json:"lp,omitempty"`
	Value                float32  `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestEntry) Reset()         { *m = TestEntry{} }
func (m *TestEntry) String() string { return proto.CompactTextString(m) }
func (*TestEntry) ProtoMessage()    {}
func (*TestEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c373e2fa0823fe5a, []int{0}
}

func (m *TestEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEntry.Unmarshal(m, b)
}
func (m *TestEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEntry.Marshal(b, m, deterministic)
}
func (m *TestEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEntry.Merge(m, src)
}
func (m *TestEntry) XXX_Size() int {
	return xxx_messageInfo_TestEntry.Size(m)
}
func (m *TestEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TestEntry proto.InternalMessageInfo

func (m *TestEntry) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestEntry) GetLp() string {
	if m != nil {
		return m.Lp
	}
	return ""
}

func (m *TestEntry) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Response struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c373e2fa0823fe5a, []int{1}
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

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Entries struct {
	Entries              []*TestEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Entries) Reset()         { *m = Entries{} }
func (m *Entries) String() string { return proto.CompactTextString(m) }
func (*Entries) ProtoMessage()    {}
func (*Entries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c373e2fa0823fe5a, []int{2}
}

func (m *Entries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entries.Unmarshal(m, b)
}
func (m *Entries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entries.Marshal(b, m, deterministic)
}
func (m *Entries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entries.Merge(m, src)
}
func (m *Entries) XXX_Size() int {
	return xxx_messageInfo_Entries.Size(m)
}
func (m *Entries) XXX_DiscardUnknown() {
	xxx_messageInfo_Entries.DiscardUnknown(m)
}

var xxx_messageInfo_Entries proto.InternalMessageInfo

func (m *Entries) GetEntries() []*TestEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*TestEntry)(nil), "excelize.TestEntry")
	proto.RegisterType((*Response)(nil), "excelize.Response")
	proto.RegisterType((*Entries)(nil), "excelize.Entries")
}

func init() { proto.RegisterFile("excelize.proto", fileDescriptor_c373e2fa0823fe5a) }

var fileDescriptor_c373e2fa0823fe5a = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbf, 0x4f, 0x87, 0x30,
	0x10, 0xc5, 0x6d, 0xbf, 0xf9, 0x0a, 0x9c, 0x09, 0xc6, 0xea, 0xd0, 0x30, 0x35, 0x1d, 0x4c, 0x17,
	0x19, 0x70, 0xd1, 0xd1, 0x81, 0xcd, 0xc1, 0x54, 0xff, 0x01, 0x85, 0x1b, 0x9a, 0x34, 0xb4, 0x69,
	0x91, 0xa8, 0x7f, 0xbd, 0xe1, 0x47, 0x61, 0x7b, 0x9f, 0x77, 0x97, 0x77, 0x2f, 0x07, 0x25, 0xfe,
	0x74, 0x68, 0xcd, 0x1f, 0xd6, 0x3e, 0xb8, 0xd1, 0xb1, 0x3c, 0xb1, 0x7c, 0x81, 0xe2, 0x03, 0xe3,
	0xd8, 0x0e, 0x63, 0xf8, 0x65, 0x25, 0x50, 0xd3, 0x73, 0x22, 0x88, 0x3a, 0x6b, 0x6a, 0xfa, 0x99,
	0xad, 0xe7, 0x54, 0x10, 0x55, 0x68, 0x6a, 0x3d, 0xbb, 0x83, 0xf3, 0xf4, 0x69, 0xbf, 0x91, 0x9f,
	0x04, 0x51, 0x54, 0xaf, 0x20, 0xef, 0x21, 0xd7, 0x18, 0xbd, 0x1b, 0x22, 0xb2, 0x0a, 0xf2, 0xb0,
	0xe9, 0x25, 0xa7, 0xd0, 0x3b, 0xcb, 0x27, 0xc8, 0xe6, 0x33, 0x06, 0x23, 0x7b, 0x80, 0x0c, 0x57,
	0xc9, 0x89, 0x38, 0xa9, 0xab, 0xe6, 0xb6, 0xde, 0x1b, 0xee, 0x75, 0x74, 0xda, 0x69, 0x5e, 0xe1,
	0xba, 0xdd, 0xc6, 0xef, 0x18, 0x26, 0xd3, 0x21, 0x7b, 0x86, 0xf2, 0x2d, 0xb8, 0xc9, 0xf4, 0x98,
	0x32, 0x6f, 0x8e, 0x88, 0xcd, 0xaa, 0xd8, 0x61, 0xa5, 0x86, 0xf2, 0xe2, 0xeb, 0x72, 0xf9, 0xc1,
	0xe3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf9, 0xa0, 0x3a, 0x15, 0x01, 0x00, 0x00,
}
