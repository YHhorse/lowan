// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ErrorNotification.proto

package out

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

type ErrorNotification struct {
	DevEui               string   `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	JoinEui              string   `protobuf:"bytes,2,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Reference            string   `protobuf:"bytes,4,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorNotification) Reset()         { *m = ErrorNotification{} }
func (m *ErrorNotification) String() string { return proto.CompactTextString(m) }
func (*ErrorNotification) ProtoMessage()    {}
func (*ErrorNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_4294fdef53b7c5f6, []int{0}
}

func (m *ErrorNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorNotification.Unmarshal(m, b)
}
func (m *ErrorNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorNotification.Marshal(b, m, deterministic)
}
func (m *ErrorNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorNotification.Merge(m, src)
}
func (m *ErrorNotification) XXX_Size() int {
	return xxx_messageInfo_ErrorNotification.Size(m)
}
func (m *ErrorNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorNotification.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorNotification proto.InternalMessageInfo

func (m *ErrorNotification) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *ErrorNotification) GetJoinEui() string {
	if m != nil {
		return m.JoinEui
	}
	return ""
}

func (m *ErrorNotification) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorNotification) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorNotification)(nil), "out.ErrorNotification")
}

func init() { proto.RegisterFile("ErrorNotification.proto", fileDescriptor_4294fdef53b7c5f6) }

var fileDescriptor_4294fdef53b7c5f6 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x77, 0x2d, 0x2a, 0xca,
	0x2f, 0xf2, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2f, 0x2d, 0x51, 0xaa, 0xe6, 0x12, 0xc4, 0x90, 0x17, 0x12,
	0xe7, 0x62, 0x4f, 0x49, 0x2d, 0x8b, 0x4f, 0x2d, 0xcd, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x62, 0x4b, 0x49, 0x2d, 0x73, 0x2d, 0xcd, 0x14, 0x92, 0xe4, 0xe2, 0xc8, 0xca, 0xcf, 0xcc, 0x03,
	0xcb, 0x30, 0x81, 0x65, 0xd8, 0x41, 0x7c, 0x90, 0x94, 0x08, 0x17, 0x6b, 0x2a, 0xc8, 0x20, 0x09,
	0x66, 0xb0, 0x38, 0x84, 0x23, 0x24, 0xc3, 0xc5, 0x59, 0x94, 0x9a, 0x96, 0x5a, 0x94, 0x9a, 0x97,
	0x9c, 0x2a, 0xc1, 0x02, 0x96, 0x41, 0x08, 0x24, 0xb1, 0x81, 0x1d, 0x62, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x6c, 0x11, 0x03, 0xa3, 0x00, 0x00, 0x00,
}