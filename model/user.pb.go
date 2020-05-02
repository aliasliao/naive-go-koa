// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/user.proto

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

type User_Gender int32

const (
	User_MALE   User_Gender = 0
	User_FEMALE User_Gender = 1
	User_OTHER  User_Gender = 2
)

var User_Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
	2: "OTHER",
}

var User_Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
	"OTHER":  2,
}

func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}

func (User_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a44479e244fa8f0, []int{0, 0}
}

type User struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32       `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Hobbies              []string    `protobuf:"bytes,4,rep,name=hobbies,proto3" json:"hobbies,omitempty"`
	Gender               User_Gender `protobuf:"varint,5,opt,name=gender,proto3,enum=model.User_Gender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a44479e244fa8f0, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetHobbies() []string {
	if m != nil {
		return m.Hobbies
	}
	return nil
}

func (m *User) GetGender() User_Gender {
	if m != nil {
		return m.Gender
	}
	return User_MALE
}

func init() {
	proto.RegisterEnum("model.User_Gender", User_Gender_name, User_Gender_value)
	proto.RegisterType((*User)(nil), "model.User")
}

func init() { proto.RegisterFile("model/user.proto", fileDescriptor_3a44479e244fa8f0) }

var fileDescriptor_3a44479e244fa8f0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xcd, 0xef, 0xb5, 0xb3, 0xb8, 0x84, 0x59, 0x65, 0x19, 0xba, 0x8a, 0x2e, 0x22, 0xe8,
	0x13, 0x28, 0x54, 0x5d, 0x28, 0x42, 0xd0, 0x8d, 0xbb, 0x96, 0x0c, 0xb5, 0x60, 0x1b, 0x49, 0xf5,
	0xb1, 0x7c, 0x47, 0x69, 0xea, 0xdd, 0x7d, 0x73, 0x66, 0xe0, 0x9b, 0x03, 0x66, 0xce, 0x89, 0x3e,
	0xaf, 0x7e, 0x56, 0x2a, 0xe1, 0xab, 0xe4, 0xef, 0x8c, 0xaa, 0x26, 0xed, 0x2f, 0x03, 0xf9, 0xb6,
	0x52, 0xc1, 0x23, 0xf0, 0x29, 0x59, 0xe6, 0x98, 0x17, 0x91, 0x4f, 0x09, 0x11, 0xe4, 0xd2, 0xcf,
	0x64, 0xb9, 0x63, 0xbe, 0x89, 0x95, 0xd1, 0x80, 0xe8, 0x47, 0xb2, 0xc2, 0x31, 0xaf, 0xe2, 0x86,
	0x68, 0xe1, 0xf0, 0x91, 0x87, 0x61, 0xa2, 0xd5, 0x4a, 0x27, 0x7c, 0x13, 0x4f, 0x23, 0x5e, 0x82,
	0x1e, 0x69, 0x49, 0x54, 0xac, 0x72, 0xcc, 0x1f, 0xaf, 0x31, 0x54, 0x61, 0xd8, 0x64, 0xe1, 0xa1,
	0x6e, 0xe2, 0xff, 0x45, 0x7b, 0x01, 0x7a, 0x4f, 0xf0, 0x1c, 0xe4, 0xf3, 0xed, 0x53, 0x67, 0xce,
	0x10, 0x40, 0xdf, 0x77, 0x95, 0x19, 0x36, 0xa0, 0x5e, 0x5e, 0x1f, 0xbb, 0x68, 0xf8, 0xdd, 0xe1,
	0x7d, 0x7f, 0x7c, 0xd0, 0xb5, 0xc6, 0xcd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x64, 0xfa,
	0x66, 0xda, 0x00, 0x00, 0x00,
}
