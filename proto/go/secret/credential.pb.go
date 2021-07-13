// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credential.proto

package secret

import (
	fmt "fmt"
	permission "github.com/appootb/substratum/proto/go/permission"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Secret type.
type Type int32

const (
	Type_CLIENT Type = 0
	Type_SERVER Type = 1
)

var Type_name = map[int32]string{
	0: "CLIENT",
	1: "SERVER",
}

var Type_value = map[string]int32{
	"CLIENT": 0,
	"SERVER": 1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1720e53dfb4809d1, []int{0}
}

// Token algorithm
type Algorithm int32

const (
	Algorithm_None  Algorithm = 0
	Algorithm_HMAC  Algorithm = 1
	Algorithm_RSA   Algorithm = 2
	Algorithm_PSS   Algorithm = 3
	Algorithm_ECDSA Algorithm = 4
	Algorithm_EdDSA Algorithm = 5
)

var Algorithm_name = map[int32]string{
	0: "None",
	1: "HMAC",
	2: "RSA",
	3: "PSS",
	4: "ECDSA",
	5: "EdDSA",
}

var Algorithm_value = map[string]int32{
	"None":  0,
	"HMAC":  1,
	"RSA":   2,
	"PSS":   3,
	"ECDSA": 4,
	"EdDSA": 5,
}

func (x Algorithm) String() string {
	return proto.EnumName(Algorithm_name, int32(x))
}

func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1720e53dfb4809d1, []int{1}
}

// Token secret info.
type Info struct {
	Type                 Type                 `protobuf:"varint,1,opt,name=type,proto3,enum=appootb.secret.Type" json:"type,omitempty"`
	Algorithm            Algorithm            `protobuf:"varint,2,opt,name=algorithm,proto3,enum=appootb.secret.Algorithm" json:"algorithm,omitempty"`
	Issuer               string               `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Account              uint64               `protobuf:"varint,4,opt,name=account,proto3" json:"account,omitempty"`
	KeyId                int64                `protobuf:"varint,5,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Roles                []string             `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Subject              permission.Subject   `protobuf:"varint,11,opt,name=subject,proto3,enum=appootb.permission.method.Subject" json:"subject,omitempty"`
	IssuedAt             *timestamp.Timestamp `protobuf:"bytes,21,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,22,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_1720e53dfb4809d1, []int{0}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CLIENT
}

func (m *Info) GetAlgorithm() Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return Algorithm_None
}

func (m *Info) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Info) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *Info) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Info) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Info) GetSubject() permission.Subject {
	if m != nil {
		return m.Subject
	}
	return permission.Subject_NONE
}

func (m *Info) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *Info) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("appootb.secret.Type", Type_name, Type_value)
	proto.RegisterEnum("appootb.secret.Algorithm", Algorithm_name, Algorithm_value)
	proto.RegisterType((*Info)(nil), "appootb.secret.Info")
}

func init() { proto.RegisterFile("credential.proto", fileDescriptor_1720e53dfb4809d1) }

var fileDescriptor_1720e53dfb4809d1 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6a, 0xdb, 0x4e,
	0x14, 0xc5, 0x33, 0xfa, 0x72, 0x74, 0x03, 0xf9, 0x8b, 0xc1, 0x09, 0xfa, 0x7b, 0xd1, 0x88, 0xac,
	0x44, 0xa0, 0x12, 0xa4, 0x8b, 0x52, 0xe8, 0x46, 0x56, 0x4d, 0x63, 0x68, 0x83, 0x19, 0x99, 0x2e,
	0x4a, 0x20, 0xe8, 0x63, 0x22, 0xab, 0xb1, 0x34, 0x42, 0x33, 0x82, 0x7a, 0xd9, 0x57, 0xe9, 0xb2,
	0x8f, 0xd2, 0x27, 0x2a, 0x5d, 0x15, 0xcd, 0x58, 0x0e, 0xcd, 0xa6, 0xbb, 0x73, 0x86, 0xdf, 0x19,
	0x5d, 0x9d, 0xb9, 0xe0, 0xe4, 0x1d, 0x2d, 0x68, 0x23, 0xaa, 0x74, 0x1b, 0xb4, 0x1d, 0x13, 0x0c,
	0x9f, 0xa6, 0x6d, 0xcb, 0x98, 0xc8, 0x02, 0x4e, 0xf3, 0x8e, 0x8a, 0xd9, 0xc5, 0xde, 0x87, 0x2d,
	0xed, 0xea, 0x8a, 0xf3, 0x8a, 0x35, 0x61, 0x4d, 0xc5, 0x86, 0x15, 0x2a, 0x30, 0xbb, 0x28, 0x19,
	0x2b, 0xb7, 0x34, 0x94, 0x2e, 0xeb, 0x1f, 0x42, 0x51, 0xd5, 0x94, 0x8b, 0xb4, 0x6e, 0x15, 0x70,
	0xf9, 0x5b, 0x03, 0x63, 0xd9, 0x3c, 0x30, 0xec, 0x83, 0x21, 0x76, 0x2d, 0x75, 0x91, 0x87, 0xfc,
	0xd3, 0xeb, 0x69, 0xf0, 0xf7, 0x97, 0x82, 0xf5, 0xae, 0xa5, 0x44, 0x12, 0xf8, 0x35, 0xd8, 0xe9,
	0xb6, 0x64, 0x5d, 0x25, 0x36, 0xb5, 0xab, 0x49, 0xfc, 0xff, 0xe7, 0x78, 0x34, 0x02, 0xe4, 0x89,
	0xc5, 0xe7, 0x60, 0x55, 0x9c, 0xf7, 0xb4, 0x73, 0x75, 0x0f, 0xf9, 0x36, 0xd9, 0x3b, 0xec, 0xc2,
	0x24, 0xcd, 0x73, 0xd6, 0x37, 0xc2, 0x35, 0x3c, 0xe4, 0x1b, 0x64, 0xb4, 0xf8, 0x0c, 0xac, 0x47,
	0xba, 0xbb, 0xaf, 0x0a, 0xd7, 0xf4, 0x90, 0xaf, 0x13, 0xf3, 0x91, 0xee, 0x96, 0x05, 0x9e, 0x82,
	0xd9, 0xb1, 0x2d, 0xe5, 0xae, 0xe5, 0xe9, 0xbe, 0x4d, 0x94, 0xc1, 0x6f, 0x61, 0xc2, 0xfb, 0xec,
	0x0b, 0xcd, 0x85, 0x7b, 0x22, 0xa7, 0xba, 0x3c, 0x4c, 0xf5, 0x54, 0x4f, 0xb0, 0xaf, 0x27, 0x51,
	0x24, 0x19, 0x23, 0xc3, 0x5f, 0xc9, 0x71, 0x8a, 0xfb, 0x54, 0xb8, 0x67, 0x1e, 0xf2, 0x4f, 0xae,
	0x67, 0x81, 0x6a, 0x2f, 0x18, 0xdb, 0x0b, 0xd6, 0x63, 0x7b, 0xe4, 0x58, 0xc1, 0x91, 0xc0, 0x6f,
	0x00, 0xe8, 0xd7, 0xb6, 0xea, 0x54, 0xf2, 0xfc, 0x9f, 0x49, 0x7b, 0x4f, 0x47, 0xe2, 0xea, 0x05,
	0x18, 0x43, 0xaf, 0x18, 0xc0, 0x8a, 0x3f, 0x2c, 0x17, 0xb7, 0x6b, 0xe7, 0x68, 0xd0, 0xc9, 0x82,
	0x7c, 0x5a, 0x10, 0x07, 0x5d, 0xbd, 0x07, 0xfb, 0x50, 0x24, 0x3e, 0x06, 0xe3, 0x96, 0x35, 0xd4,
	0x39, 0x1a, 0xd4, 0xcd, 0xc7, 0x28, 0x76, 0x10, 0x9e, 0x80, 0x4e, 0x92, 0xc8, 0xd1, 0x06, 0xb1,
	0x4a, 0x12, 0x47, 0xc7, 0x36, 0x98, 0x8b, 0xf8, 0x5d, 0x12, 0x39, 0x86, 0x94, 0xc5, 0x20, 0xcd,
	0xf9, 0x37, 0x04, 0x38, 0x67, 0xf5, 0xb3, 0x57, 0x9a, 0xff, 0x17, 0x1f, 0x16, 0x6c, 0x35, 0x0c,
	0x7a, 0x83, 0x56, 0xe8, 0xf3, 0xcb, 0xb2, 0x12, 0x9b, 0x3e, 0x0b, 0x72, 0x56, 0x87, 0xe3, 0x7a,
	0xf1, 0x3e, 0xe3, 0xa2, 0x4b, 0x45, 0x5f, 0xab, 0x4d, 0x0a, 0x4b, 0x16, 0xaa, 0x3b, 0x7e, 0x21,
	0xf4, 0x5d, 0xd3, 0xe3, 0xd5, 0xfc, 0x87, 0x66, 0x25, 0xf2, 0xe4, 0xa7, 0x36, 0x8d, 0x54, 0xec,
	0x4e, 0x5e, 0x7a, 0xa7, 0x8e, 0x33, 0x4b, 0x26, 0x5f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3b,
	0x18, 0xd7, 0xfe, 0xd6, 0x02, 0x00, 0x00,
}
