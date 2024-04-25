// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DataUpPayload.proto

package out

import (
	fmt "fmt"
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

type DataUpPayload struct {
	DevEui                  string               `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	JoinEui                 string               `protobuf:"bytes,2,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	Fcnt                    int32                `protobuf:"varint,3,opt,name=fcnt,proto3" json:"fcnt,omitempty"`
	Fport                   int32                `protobuf:"varint,4,opt,name=fport,proto3" json:"fport,omitempty"`
	Data                    []byte               `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	RxInfoGatewayId         string               `protobuf:"bytes,6,opt,name=rx_info_gateway_id,json=rxInfoGatewayId,proto3" json:"rx_info_gateway_id,omitempty"`
	RxInfoTime              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=rx_info_time,json=rxInfoTime,proto3" json:"rx_info_time,omitempty"`
	RxInfoRssi              int32                `protobuf:"varint,8,opt,name=rx_info_rssi,json=rxInfoRssi,proto3" json:"rx_info_rssi,omitempty"`
	RxInfoLoraSnr           float32              `protobuf:"fixed32,9,opt,name=rx_info_lora_snr,json=rxInfoLoraSnr,proto3" json:"rx_info_lora_snr,omitempty"`
	RxInfoLocationLatitude  float32              `protobuf:"fixed32,10,opt,name=rx_info_location_latitude,json=rxInfoLocationLatitude,proto3" json:"rx_info_location_latitude,omitempty"`
	RxInfoLocationLongitude float32              `protobuf:"fixed32,11,opt,name=rx_info_location_longitude,json=rxInfoLocationLongitude,proto3" json:"rx_info_location_longitude,omitempty"`
	Frequency               int32                `protobuf:"varint,12,opt,name=frequency,proto3" json:"frequency,omitempty"`
	GatewayCount            int32                `protobuf:"varint,13,opt,name=gateway_count,json=gatewayCount,proto3" json:"gateway_count,omitempty"`
	LoseCnt                 int32                `protobuf:"varint,14,opt,name=lose_cnt,json=loseCnt,proto3" json:"lose_cnt,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *DataUpPayload) Reset()         { *m = DataUpPayload{} }
func (m *DataUpPayload) String() string { return proto.CompactTextString(m) }
func (*DataUpPayload) ProtoMessage()    {}
func (*DataUpPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9266492bbaaf8738, []int{0}
}

func (m *DataUpPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataUpPayload.Unmarshal(m, b)
}
func (m *DataUpPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataUpPayload.Marshal(b, m, deterministic)
}
func (m *DataUpPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataUpPayload.Merge(m, src)
}
func (m *DataUpPayload) XXX_Size() int {
	return xxx_messageInfo_DataUpPayload.Size(m)
}
func (m *DataUpPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_DataUpPayload.DiscardUnknown(m)
}

var xxx_messageInfo_DataUpPayload proto.InternalMessageInfo

func (m *DataUpPayload) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *DataUpPayload) GetJoinEui() string {
	if m != nil {
		return m.JoinEui
	}
	return ""
}

func (m *DataUpPayload) GetFcnt() int32 {
	if m != nil {
		return m.Fcnt
	}
	return 0
}

func (m *DataUpPayload) GetFport() int32 {
	if m != nil {
		return m.Fport
	}
	return 0
}

func (m *DataUpPayload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataUpPayload) GetRxInfoGatewayId() string {
	if m != nil {
		return m.RxInfoGatewayId
	}
	return ""
}

func (m *DataUpPayload) GetRxInfoTime() *timestamp.Timestamp {
	if m != nil {
		return m.RxInfoTime
	}
	return nil
}

func (m *DataUpPayload) GetRxInfoRssi() int32 {
	if m != nil {
		return m.RxInfoRssi
	}
	return 0
}

func (m *DataUpPayload) GetRxInfoLoraSnr() float32 {
	if m != nil {
		return m.RxInfoLoraSnr
	}
	return 0
}

func (m *DataUpPayload) GetRxInfoLocationLatitude() float32 {
	if m != nil {
		return m.RxInfoLocationLatitude
	}
	return 0
}

func (m *DataUpPayload) GetRxInfoLocationLongitude() float32 {
	if m != nil {
		return m.RxInfoLocationLongitude
	}
	return 0
}

func (m *DataUpPayload) GetFrequency() int32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DataUpPayload) GetGatewayCount() int32 {
	if m != nil {
		return m.GatewayCount
	}
	return 0
}

func (m *DataUpPayload) GetLoseCnt() int32 {
	if m != nil {
		return m.LoseCnt
	}
	return 0
}

func init() {
	proto.RegisterType((*DataUpPayload)(nil), "out.DataUpPayload")
}

func init() { proto.RegisterFile("DataUpPayload.proto", fileDescriptor_9266492bbaaf8738) }

var fileDescriptor_9266492bbaaf8738 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x6b, 0xe3, 0x30,
	0x14, 0xc4, 0x71, 0xfe, 0x39, 0x51, 0xec, 0xdd, 0x45, 0xbb, 0x6c, 0x94, 0xb0, 0xb0, 0xa6, 0x3d,
	0xd4, 0x50, 0x70, 0xa0, 0x3d, 0x95, 0xf6, 0x96, 0x96, 0x12, 0xc8, 0xa1, 0xb8, 0xed, 0xd9, 0x28,
	0xb6, 0x6c, 0x54, 0x1c, 0x3d, 0x57, 0x96, 0xd3, 0xe4, 0x43, 0xf4, 0x3b, 0x17, 0x49, 0x31, 0xa1,
	0xed, 0x4d, 0x6f, 0x66, 0x7e, 0xb6, 0x46, 0x0f, 0xfd, 0xbe, 0xa5, 0x8a, 0x3e, 0x57, 0x0f, 0x74,
	0x5f, 0x02, 0xcd, 0xa2, 0x4a, 0x82, 0x02, 0xdc, 0x85, 0x46, 0xcd, 0xfe, 0x17, 0x00, 0x45, 0xc9,
	0xe6, 0x46, 0x5a, 0x37, 0xf9, 0x5c, 0xf1, 0x0d, 0xab, 0x15, 0xdd, 0x54, 0x36, 0x75, 0xf2, 0xde,
	0x43, 0xfe, 0x27, 0x1a, 0x4f, 0x90, 0x9b, 0xb1, 0x6d, 0xc2, 0x1a, 0x4e, 0x9c, 0xc0, 0x09, 0x47,
	0xf1, 0x20, 0x63, 0xdb, 0xbb, 0x86, 0xe3, 0x29, 0x1a, 0xbe, 0x00, 0x17, 0xc6, 0xe9, 0x18, 0xc7,
	0xd5, 0xb3, 0xb6, 0x30, 0xea, 0xe5, 0xa9, 0x50, 0xa4, 0x1b, 0x38, 0x61, 0x3f, 0x36, 0x67, 0xfc,
	0x07, 0xf5, 0xf3, 0x0a, 0xa4, 0x22, 0x3d, 0x23, 0xda, 0x41, 0x27, 0x33, 0xaa, 0x28, 0xe9, 0x07,
	0x4e, 0xe8, 0xc5, 0xe6, 0x8c, 0xcf, 0x11, 0x96, 0xbb, 0x84, 0x8b, 0x1c, 0x92, 0x82, 0x2a, 0xf6,
	0x46, 0xf7, 0x09, 0xcf, 0xc8, 0xc0, 0xfc, 0xe2, 0xa7, 0xdc, 0x2d, 0x45, 0x0e, 0xf7, 0x56, 0x5f,
	0x66, 0xf8, 0x06, 0x79, 0x6d, 0x58, 0x77, 0x21, 0x6e, 0xe0, 0x84, 0xe3, 0x8b, 0x59, 0x64, 0x8b,
	0x46, 0x6d, 0xd1, 0xe8, 0xa9, 0x2d, 0x1a, 0x23, 0xfb, 0x09, 0x2d, 0xe0, 0xe0, 0x48, 0xcb, 0xba,
	0xe6, 0x64, 0x68, 0xee, 0x76, 0x48, 0xc4, 0x75, 0xcd, 0xf1, 0x19, 0xfa, 0xd5, 0x26, 0x4a, 0x90,
	0x34, 0xa9, 0x85, 0x24, 0xa3, 0xc0, 0x09, 0x3b, 0xb1, 0x6f, 0x53, 0x2b, 0x90, 0xf4, 0x51, 0x48,
	0x7c, 0x85, 0xa6, 0xc7, 0x60, 0x4a, 0x15, 0x07, 0x91, 0x94, 0x54, 0x71, 0xd5, 0x64, 0x8c, 0x20,
	0x43, 0xfc, 0x6d, 0x09, 0x6b, 0xaf, 0x0e, 0x2e, 0xbe, 0x46, 0xb3, 0xef, 0x28, 0x88, 0xc2, 0xb2,
	0x63, 0xc3, 0x4e, 0xbe, 0xb0, 0xad, 0x8d, 0xff, 0xa1, 0x51, 0x2e, 0xd9, 0x6b, 0xc3, 0x44, 0xba,
	0x27, 0x9e, 0xb9, 0xff, 0x51, 0xc0, 0xa7, 0xc8, 0x6f, 0xdf, 0x30, 0x85, 0x46, 0x28, 0xe2, 0x9b,
	0x84, 0x77, 0x10, 0x17, 0x5a, 0xd3, 0x9b, 0x2c, 0xa1, 0x66, 0x89, 0x5e, 0xd9, 0x0f, 0xe3, 0xbb,
	0x7a, 0x5e, 0x08, 0xb5, 0x1e, 0x98, 0x07, 0xbc, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x99,
	0x46, 0xf0, 0x53, 0x02, 0x00, 0x00,
}