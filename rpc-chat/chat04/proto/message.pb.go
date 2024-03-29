// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

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

type Msg1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Sex                  bool     `protobuf:"varint,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Project              string   `protobuf:"bytes,5,opt,name=Project,proto3" json:"Project,omitempty"`
	Job                  uint32   `protobuf:"varint,6,opt,name=Job,proto3" json:"Job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg1) Reset()         { *m = Msg1{} }
func (m *Msg1) String() string { return proto.CompactTextString(m) }
func (*Msg1) ProtoMessage()    {}
func (*Msg1) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Msg1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg1.Unmarshal(m, b)
}
func (m *Msg1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg1.Marshal(b, m, deterministic)
}
func (m *Msg1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg1.Merge(m, src)
}
func (m *Msg1) XXX_Size() int {
	return xxx_messageInfo_Msg1.Size(m)
}
func (m *Msg1) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg1.DiscardUnknown(m)
}

var xxx_messageInfo_Msg1 proto.InternalMessageInfo

func (m *Msg1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg1) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Msg1) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *Msg1) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Msg1) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Msg1) GetJob() uint32 {
	if m != nil {
		return m.Job
	}
	return 0
}

type Msg2 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Sex                  bool     `protobuf:"varint,3,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Job                  uint32   `protobuf:"varint,6,opt,name=Job,proto3" json:"Job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg2) Reset()         { *m = Msg2{} }
func (m *Msg2) String() string { return proto.CompactTextString(m) }
func (*Msg2) ProtoMessage()    {}
func (*Msg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Msg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg2.Unmarshal(m, b)
}
func (m *Msg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg2.Marshal(b, m, deterministic)
}
func (m *Msg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg2.Merge(m, src)
}
func (m *Msg2) XXX_Size() int {
	return xxx_messageInfo_Msg2.Size(m)
}
func (m *Msg2) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg2.DiscardUnknown(m)
}

var xxx_messageInfo_Msg2 proto.InternalMessageInfo

func (m *Msg2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg2) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Msg2) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *Msg2) GetJob() uint32 {
	if m != nil {
		return m.Job
	}
	return 0
}

func init() {
	proto.RegisterType((*Msg1)(nil), "Msg1")
	proto.RegisterType((*Msg2)(nil), "Msg2")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xaa, 0xe3, 0x62, 0xf1, 0x2d, 0x4e,
	0x37, 0x14, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x1d, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83,
	0x40, 0x4c, 0x90, 0x48, 0x70, 0x6a, 0x85, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x88, 0x09,
	0xd2, 0xe7, 0x92, 0x5a, 0x9c, 0x2c, 0xc1, 0x02, 0xd1, 0x07, 0x62, 0x0b, 0x49, 0x70, 0xb1, 0x07,
	0x14, 0xe5, 0x67, 0xa5, 0x26, 0x97, 0x48, 0xb0, 0x82, 0x85, 0x61, 0x5c, 0x90, 0x7e, 0xaf, 0xfc,
	0x24, 0x09, 0x36, 0x88, 0x89, 0x5e, 0xf9, 0x49, 0x4a, 0x01, 0x60, 0xfb, 0x8d, 0xc8, 0xb6, 0x1f,
	0xc3, 0x44, 0x27, 0xf6, 0x28, 0x56, 0xb0, 0xd7, 0x92, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x1d, 0xc0, 0x61, 0x4b, 0xf2, 0x00, 0x00, 0x00,
}
