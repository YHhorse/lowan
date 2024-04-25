// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}

var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
}

var Region_value = map[string]int32{
	"EU868": 0,
	"US915": 2,
	"CN779": 3,
	"EU433": 4,
	"AU915": 5,
	"CN470": 6,
	"AS923": 7,
	"KR920": 8,
	"IN865": 9,
	"RU864": 10,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver.
	LocationSource_GEO_RESOLVER LocationSource = 3
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":      0,
	"GPS":          1,
	"CONFIG":       2,
	"GEO_RESOLVER": 3,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0xad, 0xf2, 0xe1, 0xd8, 0xb7, 0x5d, 0x11, 0x1a, 0x6c, 0x61, 0x1b, 0x2c, 0xf4, 0x61, 0x98,
	0xc0, 0x92, 0x2e, 0x49, 0xf3, 0xf1, 0xd8, 0x65, 0x6e, 0x28, 0xf6, 0xec, 0x21, 0xe3, 0x0d, 0xf6,
	0x12, 0x54, 0x57, 0x18, 0x63, 0xd5, 0x0a, 0xb1, 0xdd, 0xe1, 0xff, 0xb4, 0x1f, 0x39, 0x64, 0xb9,
	0x1d, 0x7d, 0xf2, 0x39, 0xf7, 0x9c, 0x73, 0x7d, 0xaf, 0xb8, 0xf0, 0x3a, 0x96, 0x0f, 0x0f, 0x32,
	0x9f, 0xea, 0xcf, 0xe4, 0x70, 0x94, 0xa5, 0x24, 0x86, 0x66, 0x17, 0x5b, 0x38, 0x75, 0x79, 0xed,
	0xe4, 0x8f, 0x5c, 0xc8, 0x03, 0x27, 0xef, 0xc1, 0xca, 0x78, 0xb6, 0x17, 0xec, 0x8e, 0x8b, 0x21,
	0x1a, 0x21, 0xdb, 0xa2, 0x66, 0xc6, 0x33, 0x4f, 0x71, 0xf2, 0x16, 0x06, 0x8c, 0x17, 0xfb, 0x8c,
	0xd7, 0xc3, 0xce, 0x08, 0xd9, 0x67, 0xd4, 0x60, 0xbc, 0x70, 0x79, 0x7d, 0xf1, 0x17, 0x81, 0xe9,
	0xc9, 0x98, 0x95, 0xa9, 0xcc, 0xc9, 0x3b, 0x30, 0x05, 0x2b, 0xd3, 0xb2, 0xba, 0xe7, 0x4d, 0x07,
	0x44, 0x9f, 0x39, 0xf9, 0x00, 0x96, 0x90, 0x79, 0xa2, 0xc5, 0x4e, 0x23, 0xfe, 0x2f, 0xa8, 0x24,
	0x13, 0x6d, 0xb2, 0xab, 0x93, 0x4f, 0x9c, 0x4c, 0xc0, 0x28, 0x64, 0x75, 0x8c, 0xf9, 0xb0, 0x37,
	0x42, 0xf6, 0xf9, 0xec, 0xcd, 0xa4, 0x5d, 0xe7, 0xe9, 0xbf, 0x61, 0xa3, 0xd2, 0xd6, 0xd5, 0xf4,
	0x8a, 0xe3, 0xea, 0xc8, 0xe2, 0x7a, 0xd8, 0x1f, 0x21, 0xfb, 0x15, 0x7d, 0xe6, 0xe3, 0x8f, 0x00,
	0xdf, 0xe5, 0x7d, 0x25, 0xf4, 0xbc, 0x26, 0xf4, 0xbc, 0x80, 0x5e, 0xe3, 0x13, 0x32, 0x80, 0xee,
	0x4d, 0xe8, 0x62, 0x34, 0x7e, 0x04, 0x83, 0xf2, 0x44, 0x89, 0x16, 0xf4, 0x9d, 0x68, 0xbd, 0x5c,
	0xe3, 0x13, 0x05, 0xa3, 0x70, 0xf3, 0xe5, 0x0a, 0x77, 0x14, 0xdc, 0xfa, 0xab, 0xd5, 0x06, 0x77,
	0xb5, 0x61, 0x31, 0x9f, 0xe3, 0x9e, 0x82, 0xd7, 0x91, 0x32, 0xf4, 0xb5, 0x61, 0xb1, 0xba, 0xc4,
	0x46, 0x53, 0x0d, 0x37, 0xb3, 0x39, 0x1e, 0x28, 0xe8, 0xd2, 0xcd, 0xec, 0x12, 0x9b, 0x0a, 0xde,
	0xfa, 0xeb, 0xe5, 0x15, 0xb6, 0x14, 0xa4, 0xd1, 0x7a, 0xb9, 0xc0, 0x30, 0xfe, 0x06, 0xe7, 0x2f,
	0xd7, 0x21, 0xa7, 0x30, 0x88, 0x7c, 0xd7, 0x0f, 0x7e, 0xf9, 0x7a, 0xbe, 0xdd, 0x8f, 0x10, 0x23,
	0x02, 0x60, 0x6c, 0x03, 0xff, 0xe6, 0x76, 0x87, 0x3b, 0x04, 0xc3, 0xd9, 0xce, 0x09, 0xf6, 0xd4,
	0x09, 0x03, 0xef, 0xa7, 0x43, 0x71, 0xf7, 0xab, 0xfd, 0xfb, 0x53, 0x22, 0x27, 0x42, 0xfe, 0x61,
	0x79, 0x2a, 0x4b, 0xf5, 0x4e, 0xd3, 0x86, 0x4c, 0xf3, 0xe2, 0x33, 0x3b, 0xa4, 0xd3, 0x44, 0xb6,
	0xa7, 0x70, 0x67, 0x34, 0xb7, 0x30, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x04, 0xd4, 0xaa, 0xb2,
	0x22, 0x02, 0x00, 0x00,
}
