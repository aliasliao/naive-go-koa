// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/city.proto

package model

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

type City struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Users                []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *City) Reset()         { *m = City{} }
func (m *City) String() string { return proto.CompactTextString(m) }
func (*City) ProtoMessage()    {}
func (*City) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec86b249bacb29f3, []int{0}
}

func (m *City) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_City.Unmarshal(m, b)
}
func (m *City) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_City.Marshal(b, m, deterministic)
}
func (m *City) XXX_Merge(src proto.Message) {
	xxx_messageInfo_City.Merge(m, src)
}
func (m *City) XXX_Size() int {
	return xxx_messageInfo_City.Size(m)
}
func (m *City) XXX_DiscardUnknown() {
	xxx_messageInfo_City.DiscardUnknown(m)
}

var xxx_messageInfo_City proto.InternalMessageInfo

func (m *City) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *City) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *City) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*City)(nil), "model.City")
}

func init() { proto.RegisterFile("model/city.proto", fileDescriptor_ec86b249bacb29f3) }

var fileDescriptor_ec86b249bacb29f3 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x4f, 0xce, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x8b,
	0x48, 0x41, 0x25, 0x4a, 0x8b, 0x53, 0x8b, 0x20, 0x12, 0x4a, 0xbe, 0x5c, 0x2c, 0xce, 0x99, 0x25,
	0x95, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x22, 0x17, 0x2b, 0x48, 0x67, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x11,
	0xb7, 0x1e, 0xd8, 0x34, 0xbd, 0xd0, 0xe2, 0xd4, 0xa2, 0x20, 0x88, 0x8c, 0x13, 0x7b, 0x14, 0xc4,
	0xa6, 0x24, 0x36, 0xb0, 0xf1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x77, 0x01, 0x4e,
	0x8b, 0x00, 0x00, 0x00,
}
