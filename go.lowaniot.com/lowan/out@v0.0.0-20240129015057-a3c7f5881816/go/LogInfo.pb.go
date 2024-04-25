// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LogInfo.proto

package out

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	gw "go.lowaniot.com/lowan/ns-api/go/gw"
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

type LogInfo struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Direction            string               `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	Reference            string               `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Fport                int32                `protobuf:"varint,5,opt,name=fport,proto3" json:"fport,omitempty"`
	Data                 []byte               `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	AppData              []byte               `protobuf:"bytes,7,opt,name=app_data,json=appData,proto3" json:"app_data,omitempty"`
	LoseCnt              int32                `protobuf:"varint,8,opt,name=lose_cnt,json=loseCnt,proto3" json:"lose_cnt,omitempty"`
	Confirmed            bool                 `protobuf:"varint,9,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	Ack                  bool                 `protobuf:"varint,10,opt,name=ack,proto3" json:"ack,omitempty"`
	Fcnt                 int32                `protobuf:"varint,11,opt,name=fcnt,proto3" json:"fcnt,omitempty"`
	Tmst                 int32                `protobuf:"varint,12,opt,name=tmst,proto3" json:"tmst,omitempty"`
	Slot                 int32                `protobuf:"varint,13,opt,name=slot,proto3" json:"slot,omitempty"`
	GwDelay              int32                `protobuf:"varint,14,opt,name=gw_delay,json=gwDelay,proto3" json:"gw_delay,omitempty"`
	MacVersion           string               `protobuf:"bytes,15,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	DevAddr              []byte               `protobuf:"bytes,16,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
	DevEui               []byte               `protobuf:"bytes,17,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	JoinEui              []byte               `protobuf:"bytes,18,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	Adr                  bool                 `protobuf:"varint,19,opt,name=adr,proto3" json:"adr,omitempty"`
	TxPowerIndex         int32                `protobuf:"varint,20,opt,name=tx_power_index,json=txPowerIndex,proto3" json:"tx_power_index,omitempty"`
	NbTrans              int32                `protobuf:"varint,21,opt,name=nb_trans,json=nbTrans,proto3" json:"nb_trans,omitempty"`
	FcntUp               int32                `protobuf:"varint,22,opt,name=fcnt_up,json=fcntUp,proto3" json:"fcnt_up,omitempty"`
	NFcntDown            int32                `protobuf:"varint,23,opt,name=n_fcnt_down,json=nFcntDown,proto3" json:"n_fcnt_down,omitempty"`
	AFcntDown            int32                `protobuf:"varint,24,opt,name=a_fcnt_down,json=aFcntDown,proto3" json:"a_fcnt_down,omitempty"`
	ConfFcnt             int32                `protobuf:"varint,25,opt,name=conf_fcnt,json=confFcnt,proto3" json:"conf_fcnt,omitempty"`
	GatewayCount         int32                `protobuf:"varint,26,opt,name=gateway_count,json=gatewayCount,proto3" json:"gateway_count,omitempty"`
	GatewayId            []byte               `protobuf:"bytes,27,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	RxInfoSet            []*gw.UplinkRXInfo   `protobuf:"bytes,28,rep,name=rx_info_set,json=rxInfoSet,proto3" json:"rx_info_set,omitempty"`
	TxInfo               *UpDownTXInfo        `protobuf:"bytes,29,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	MacCommands          []*MACCommand        `protobuf:"bytes,30,rep,name=mac_commands,json=macCommands,proto3" json:"mac_commands,omitempty"`
	Error                string               `protobuf:"bytes,31,opt,name=error,proto3" json:"error,omitempty"`
	ErrorType            int32                `protobuf:"varint,32,opt,name=error_type,json=errorType,proto3" json:"error_type,omitempty"`
	ErrorFrame           []byte               `protobuf:"bytes,33,opt,name=error_frame,json=errorFrame,proto3" json:"error_frame,omitempty"`
	IsFinalProcess       bool                 `protobuf:"varint,34,opt,name=is_final_process,json=isFinalProcess,proto3" json:"is_final_process,omitempty"`
	RfRegion             string               `protobuf:"bytes,35,opt,name=rf_region,json=rfRegion,proto3" json:"rf_region,omitempty"`
	DeviceMode           string               `protobuf:"bytes,36,opt,name=device_mode,json=deviceMode,proto3" json:"device_mode,omitempty"`
	Dr                   int32                `protobuf:"varint,37,opt,name=dr,proto3" json:"dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60a416064a69f13b, []int{0}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *LogInfo) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *LogInfo) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *LogInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LogInfo) GetFport() int32 {
	if m != nil {
		return m.Fport
	}
	return 0
}

func (m *LogInfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LogInfo) GetAppData() []byte {
	if m != nil {
		return m.AppData
	}
	return nil
}

func (m *LogInfo) GetLoseCnt() int32 {
	if m != nil {
		return m.LoseCnt
	}
	return 0
}

func (m *LogInfo) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *LogInfo) GetAck() bool {
	if m != nil {
		return m.Ack
	}
	return false
}

func (m *LogInfo) GetFcnt() int32 {
	if m != nil {
		return m.Fcnt
	}
	return 0
}

func (m *LogInfo) GetTmst() int32 {
	if m != nil {
		return m.Tmst
	}
	return 0
}

func (m *LogInfo) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *LogInfo) GetGwDelay() int32 {
	if m != nil {
		return m.GwDelay
	}
	return 0
}

func (m *LogInfo) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *LogInfo) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *LogInfo) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *LogInfo) GetJoinEui() []byte {
	if m != nil {
		return m.JoinEui
	}
	return nil
}

func (m *LogInfo) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *LogInfo) GetTxPowerIndex() int32 {
	if m != nil {
		return m.TxPowerIndex
	}
	return 0
}

func (m *LogInfo) GetNbTrans() int32 {
	if m != nil {
		return m.NbTrans
	}
	return 0
}

func (m *LogInfo) GetFcntUp() int32 {
	if m != nil {
		return m.FcntUp
	}
	return 0
}

func (m *LogInfo) GetNFcntDown() int32 {
	if m != nil {
		return m.NFcntDown
	}
	return 0
}

func (m *LogInfo) GetAFcntDown() int32 {
	if m != nil {
		return m.AFcntDown
	}
	return 0
}

func (m *LogInfo) GetConfFcnt() int32 {
	if m != nil {
		return m.ConfFcnt
	}
	return 0
}

func (m *LogInfo) GetGatewayCount() int32 {
	if m != nil {
		return m.GatewayCount
	}
	return 0
}

func (m *LogInfo) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *LogInfo) GetRxInfoSet() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfoSet
	}
	return nil
}

func (m *LogInfo) GetTxInfo() *UpDownTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *LogInfo) GetMacCommands() []*MACCommand {
	if m != nil {
		return m.MacCommands
	}
	return nil
}

func (m *LogInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LogInfo) GetErrorType() int32 {
	if m != nil {
		return m.ErrorType
	}
	return 0
}

func (m *LogInfo) GetErrorFrame() []byte {
	if m != nil {
		return m.ErrorFrame
	}
	return nil
}

func (m *LogInfo) GetIsFinalProcess() bool {
	if m != nil {
		return m.IsFinalProcess
	}
	return false
}

func (m *LogInfo) GetRfRegion() string {
	if m != nil {
		return m.RfRegion
	}
	return ""
}

func (m *LogInfo) GetDeviceMode() string {
	if m != nil {
		return m.DeviceMode
	}
	return ""
}

func (m *LogInfo) GetDr() int32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func init() {
	proto.RegisterType((*LogInfo)(nil), "out.LogInfo")
}

func init() { proto.RegisterFile("LogInfo.proto", fileDescriptor_60a416064a69f13b) }

var fileDescriptor_60a416064a69f13b = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x93, 0xdb, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0xe1, 0x38, 0xf1, 0x81, 0x76, 0xbc, 0x5e, 0x76, 0xdb, 0x9d, 0xf5, 0x1e, 0xe2, 0xee,
	0x6e, 0x01, 0xa3, 0x17, 0x4a, 0x91, 0x3e, 0x41, 0xe0, 0x34, 0x40, 0x80, 0x06, 0x08, 0x54, 0xa7,
	0xe8, 0x1d, 0x41, 0x8b, 0x94, 0xc0, 0xc6, 0x22, 0x05, 0x8a, 0xb2, 0xec, 0x87, 0xec, 0x3b, 0x15,
	0x33, 0x74, 0xe2, 0x5c, 0x69, 0xe6, 0xfb, 0x49, 0xea, 0x9f, 0xe1, 0x90, 0x9d, 0xff, 0xe9, 0x8a,
	0x3b, 0x9b, 0xbb, 0xa4, 0xf2, 0x2e, 0x38, 0xde, 0x75, 0x4d, 0x98, 0x5d, 0x14, 0xce, 0x15, 0x1b,
	0x7d, 0x49, 0x68, 0xdd, 0xe4, 0x97, 0xc1, 0x94, 0xba, 0x0e, 0xb2, 0xac, 0xe2, 0xaa, 0xd9, 0xa8,
	0x68, 0x2f, 0x8b, 0xf6, 0x90, 0xf0, 0xc7, 0xea, 0xc6, 0xb5, 0x76, 0xf5, 0xcf, 0xf1, 0x98, 0xd9,
	0xf4, 0xfe, 0x7a, 0xb9, 0x74, 0x65, 0x29, 0xad, 0x8a, 0xe4, 0xeb, 0x7f, 0x03, 0xd6, 0x3f, 0xfc,
	0x8a, 0x73, 0x76, 0xda, 0x34, 0x46, 0x41, 0x67, 0xde, 0x59, 0x0c, 0x53, 0x8a, 0xf9, 0x27, 0x36,
	0x54, 0xc6, 0xeb, 0x2c, 0x18, 0x67, 0xe1, 0x84, 0x84, 0x23, 0x40, 0xd5, 0xeb, 0x5c, 0x7b, 0x6d,
	0x33, 0x0d, 0xdd, 0xa8, 0xbe, 0x00, 0x9e, 0xb0, 0x53, 0x74, 0x08, 0xa7, 0xf3, 0xce, 0x62, 0x74,
	0x35, 0x4b, 0xa2, 0xfd, 0xe4, 0xd9, 0x7e, 0xb2, 0x7a, 0xb6, 0x9f, 0xd2, 0x3a, 0xfe, 0x8e, 0x9d,
	0xe5, 0x95, 0xf3, 0x01, 0xce, 0xe6, 0x9d, 0xc5, 0x59, 0x1a, 0x13, 0x74, 0xa5, 0x64, 0x90, 0xd0,
	0x9b, 0x77, 0x16, 0xe3, 0x94, 0x62, 0xfe, 0x81, 0x0d, 0x64, 0x55, 0x09, 0xe2, 0x7d, 0xe2, 0x7d,
	0x59, 0x55, 0x37, 0x07, 0x69, 0xe3, 0x6a, 0x2d, 0x32, 0x1b, 0x60, 0x40, 0xe7, 0xf4, 0x31, 0x5f,
	0xda, 0x80, 0x6e, 0x33, 0x67, 0x73, 0xe3, 0x4b, 0xad, 0x60, 0x38, 0xef, 0x2c, 0x06, 0xe9, 0x11,
	0xf0, 0x29, 0xeb, 0xca, 0xec, 0x09, 0x18, 0x71, 0x0c, 0xf1, 0xcf, 0x39, 0x1e, 0x33, 0xa2, 0x63,
	0x28, 0x46, 0x16, 0xca, 0x3a, 0xc0, 0x38, 0x32, 0x8c, 0x91, 0xd5, 0x1b, 0x17, 0xe0, 0x3c, 0x32,
	0x8c, 0xd1, 0x46, 0xd1, 0x0a, 0xa5, 0x37, 0x72, 0x0f, 0x93, 0x68, 0xa3, 0x68, 0x6f, 0x30, 0xe5,
	0x17, 0x6c, 0x54, 0xca, 0x4c, 0x6c, 0xb5, 0xaf, 0xb1, 0xa9, 0x6f, 0xa8, 0x6d, 0xac, 0x94, 0xd9,
	0xdf, 0x91, 0xe0, 0x5e, 0xa5, 0xb7, 0x42, 0x2a, 0xe5, 0x61, 0x1a, 0xab, 0x53, 0x7a, 0x7b, 0xad,
	0x94, 0xe7, 0xef, 0x19, 0x86, 0x42, 0x37, 0x06, 0xde, 0x92, 0xd2, 0x53, 0x7a, 0xfb, 0x47, 0x63,
	0x70, 0xcf, 0xbf, 0xce, 0x58, 0x52, 0x78, 0xdc, 0x83, 0x39, 0x4a, 0x58, 0x98, 0xf2, 0xf0, 0xc3,
	0xa1, 0x30, 0xe5, 0xf9, 0x77, 0x36, 0x09, 0x3b, 0x51, 0xb9, 0x56, 0x7b, 0x61, 0xac, 0xd2, 0x3b,
	0x78, 0x47, 0x16, 0xc7, 0x61, 0xf7, 0x80, 0xf0, 0x0e, 0x19, 0x1e, 0x69, 0xd7, 0x22, 0x78, 0x69,
	0x6b, 0xf8, 0x31, 0x96, 0x60, 0xd7, 0x2b, 0x4c, 0xd1, 0x06, 0x76, 0x43, 0x34, 0x15, 0xfc, 0x44,
	0x4a, 0x0f, 0xd3, 0xc7, 0x8a, 0x7f, 0x61, 0x23, 0x2b, 0x48, 0x52, 0xae, 0xb5, 0xf0, 0x9e, 0xc4,
	0xa1, 0xbd, 0xcd, 0x6c, 0xc0, 0x61, 0x44, 0x5d, 0xbe, 0xd2, 0x21, 0xea, 0xf2, 0x45, 0xff, 0x18,
	0xaf, 0x88, 0x96, 0xc0, 0x07, 0x52, 0x07, 0x08, 0x70, 0x01, 0xff, 0xc6, 0xce, 0x0b, 0x19, 0x74,
	0x2b, 0xf7, 0x22, 0x73, 0x8d, 0x0d, 0x30, 0x8b, 0xae, 0x0f, 0x70, 0x89, 0x8c, 0x7f, 0x66, 0xec,
	0x79, 0x91, 0x51, 0xf0, 0x91, 0x5a, 0x31, 0x3c, 0x90, 0x3b, 0xc5, 0x7f, 0x63, 0x23, 0xbf, 0x13,
	0xc6, 0xe6, 0x4e, 0xd4, 0x3a, 0xc0, 0xa7, 0x79, 0x77, 0x31, 0xba, 0x9a, 0x26, 0x45, 0x9b, 0x3c,
	0x56, 0x1b, 0x63, 0x9f, 0x52, 0x7a, 0x2e, 0xe9, 0xd0, 0xef, 0xf0, 0xfb, 0x97, 0x0e, 0xfc, 0x57,
	0xd6, 0x0f, 0x71, 0x07, 0x7c, 0xa6, 0x41, 0x7e, 0x9b, 0xb8, 0x26, 0x24, 0xaf, 0x5f, 0x57, 0xda,
	0x0b, 0xb4, 0x9c, 0x5f, 0xb1, 0x31, 0x5e, 0x6d, 0x16, 0x9f, 0x58, 0x0d, 0x5f, 0xe8, 0xf8, 0x37,
	0xb4, 0xe1, 0xf8, 0xf4, 0x52, 0xbc, 0xff, 0x43, 0x5c, 0xe3, 0xd4, 0x6b, 0xef, 0x9d, 0x87, 0x0b,
	0x1a, 0x84, 0x98, 0x60, 0x19, 0x14, 0x88, 0xb0, 0xaf, 0x34, 0xcc, 0x63, 0x9f, 0x88, 0xac, 0xf6,
	0x95, 0xc6, 0x19, 0x8a, 0x72, 0xee, 0x65, 0xa9, 0xe1, 0x67, 0x2a, 0x33, 0xee, 0xb8, 0x45, 0xc2,
	0x17, 0x6c, 0x6a, 0x6a, 0x91, 0x1b, 0x2b, 0x37, 0xa2, 0xf2, 0x2e, 0xd3, 0x75, 0x0d, 0x5f, 0x69,
	0x02, 0x26, 0xa6, 0xbe, 0x45, 0xfc, 0x10, 0x29, 0xb6, 0xdc, 0xe7, 0xc2, 0xeb, 0x02, 0x87, 0xf1,
	0x1b, 0x79, 0x18, 0xf8, 0x3c, 0xa5, 0x1c, 0xff, 0xa3, 0xf4, 0xd6, 0x64, 0x5a, 0x94, 0x4e, 0x69,
	0xf8, 0x1e, 0x67, 0x35, 0xa2, 0x7b, 0xa7, 0x34, 0x9f, 0xb0, 0x13, 0xe5, 0xe1, 0x17, 0xf2, 0x77,
	0xa2, 0xfc, 0xba, 0x47, 0xaf, 0xfb, 0xf7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x51, 0x43, 0x3e,
	0xbf, 0xc0, 0x04, 0x00, 0x00,
}