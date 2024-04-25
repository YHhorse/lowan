// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geo/geo.proto

package geo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "go.lowaniot.com/lowan/ns-api/go/common"
	gw "go.lowaniot.com/lowan/ns-api/go/gw"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ResolveResult struct {
	// Resolved location.
	Location             *common.Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ResolveResult) Reset()         { *m = ResolveResult{} }
func (m *ResolveResult) String() string { return proto.CompactTextString(m) }
func (*ResolveResult) ProtoMessage()    {}
func (*ResolveResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{0}
}

func (m *ResolveResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveResult.Unmarshal(m, b)
}
func (m *ResolveResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveResult.Marshal(b, m, deterministic)
}
func (m *ResolveResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResult.Merge(m, src)
}
func (m *ResolveResult) XXX_Size() int {
	return xxx_messageInfo_ResolveResult.Size(m)
}
func (m *ResolveResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResult.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResult proto.InternalMessageInfo

func (m *ResolveResult) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type FrameRXInfo struct {
	// Uplink Gateway meta-data.
	RxInfo               []*gw.UplinkRXInfo `protobuf:"bytes,1,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FrameRXInfo) Reset()         { *m = FrameRXInfo{} }
func (m *FrameRXInfo) String() string { return proto.CompactTextString(m) }
func (*FrameRXInfo) ProtoMessage()    {}
func (*FrameRXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{1}
}

func (m *FrameRXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameRXInfo.Unmarshal(m, b)
}
func (m *FrameRXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameRXInfo.Marshal(b, m, deterministic)
}
func (m *FrameRXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameRXInfo.Merge(m, src)
}
func (m *FrameRXInfo) XXX_Size() int {
	return xxx_messageInfo_FrameRXInfo.Size(m)
}
func (m *FrameRXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameRXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FrameRXInfo proto.InternalMessageInfo

func (m *FrameRXInfo) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type ResolveTDOARequest struct {
	// Device ID.
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Frame meta-data.
	FrameRxInfo *FrameRXInfo `protobuf:"bytes,2,opt,name=frame_rx_info,json=frameRXInfo,proto3" json:"frame_rx_info,omitempty"`
	// Device reference altitude.
	DeviceReferenceAltitude float64  `protobuf:"fixed64,3,opt,name=device_reference_altitude,json=deviceReferenceAltitude,proto3" json:"device_reference_altitude,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ResolveTDOARequest) Reset()         { *m = ResolveTDOARequest{} }
func (m *ResolveTDOARequest) String() string { return proto.CompactTextString(m) }
func (*ResolveTDOARequest) ProtoMessage()    {}
func (*ResolveTDOARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{2}
}

func (m *ResolveTDOARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveTDOARequest.Unmarshal(m, b)
}
func (m *ResolveTDOARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveTDOARequest.Marshal(b, m, deterministic)
}
func (m *ResolveTDOARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveTDOARequest.Merge(m, src)
}
func (m *ResolveTDOARequest) XXX_Size() int {
	return xxx_messageInfo_ResolveTDOARequest.Size(m)
}
func (m *ResolveTDOARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveTDOARequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveTDOARequest proto.InternalMessageInfo

func (m *ResolveTDOARequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *ResolveTDOARequest) GetFrameRxInfo() *FrameRXInfo {
	if m != nil {
		return m.FrameRxInfo
	}
	return nil
}

func (m *ResolveTDOARequest) GetDeviceReferenceAltitude() float64 {
	if m != nil {
		return m.DeviceReferenceAltitude
	}
	return 0
}

type ResolveMultiFrameTDOARequest struct {
	// Device ID.
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Frames meta-data.
	FrameRxInfoSet []*FrameRXInfo `protobuf:"bytes,2,rep,name=frame_rx_info_set,json=frameRXInfoSet,proto3" json:"frame_rx_info_set,omitempty"`
	// Device reference altitude.
	DeviceReferenceAltitude float64  `protobuf:"fixed64,3,opt,name=device_reference_altitude,json=deviceReferenceAltitude,proto3" json:"device_reference_altitude,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ResolveMultiFrameTDOARequest) Reset()         { *m = ResolveMultiFrameTDOARequest{} }
func (m *ResolveMultiFrameTDOARequest) String() string { return proto.CompactTextString(m) }
func (*ResolveMultiFrameTDOARequest) ProtoMessage()    {}
func (*ResolveMultiFrameTDOARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{3}
}

func (m *ResolveMultiFrameTDOARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveMultiFrameTDOARequest.Unmarshal(m, b)
}
func (m *ResolveMultiFrameTDOARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveMultiFrameTDOARequest.Marshal(b, m, deterministic)
}
func (m *ResolveMultiFrameTDOARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveMultiFrameTDOARequest.Merge(m, src)
}
func (m *ResolveMultiFrameTDOARequest) XXX_Size() int {
	return xxx_messageInfo_ResolveMultiFrameTDOARequest.Size(m)
}
func (m *ResolveMultiFrameTDOARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveMultiFrameTDOARequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveMultiFrameTDOARequest proto.InternalMessageInfo

func (m *ResolveMultiFrameTDOARequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *ResolveMultiFrameTDOARequest) GetFrameRxInfoSet() []*FrameRXInfo {
	if m != nil {
		return m.FrameRxInfoSet
	}
	return nil
}

func (m *ResolveMultiFrameTDOARequest) GetDeviceReferenceAltitude() float64 {
	if m != nil {
		return m.DeviceReferenceAltitude
	}
	return 0
}

type ResolveTDOAResponse struct {
	// Resolve result.
	Result               *ResolveResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResolveTDOAResponse) Reset()         { *m = ResolveTDOAResponse{} }
func (m *ResolveTDOAResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveTDOAResponse) ProtoMessage()    {}
func (*ResolveTDOAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{4}
}

func (m *ResolveTDOAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveTDOAResponse.Unmarshal(m, b)
}
func (m *ResolveTDOAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveTDOAResponse.Marshal(b, m, deterministic)
}
func (m *ResolveTDOAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveTDOAResponse.Merge(m, src)
}
func (m *ResolveTDOAResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveTDOAResponse.Size(m)
}
func (m *ResolveTDOAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveTDOAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveTDOAResponse proto.InternalMessageInfo

func (m *ResolveTDOAResponse) GetResult() *ResolveResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ResolveMultiFrameTDOAResponse struct {
	// Resolve result.
	Result               *ResolveResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResolveMultiFrameTDOAResponse) Reset()         { *m = ResolveMultiFrameTDOAResponse{} }
func (m *ResolveMultiFrameTDOAResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveMultiFrameTDOAResponse) ProtoMessage()    {}
func (*ResolveMultiFrameTDOAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{5}
}

func (m *ResolveMultiFrameTDOAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveMultiFrameTDOAResponse.Unmarshal(m, b)
}
func (m *ResolveMultiFrameTDOAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveMultiFrameTDOAResponse.Marshal(b, m, deterministic)
}
func (m *ResolveMultiFrameTDOAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveMultiFrameTDOAResponse.Merge(m, src)
}
func (m *ResolveMultiFrameTDOAResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveMultiFrameTDOAResponse.Size(m)
}
func (m *ResolveMultiFrameTDOAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveMultiFrameTDOAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveMultiFrameTDOAResponse proto.InternalMessageInfo

func (m *ResolveMultiFrameTDOAResponse) GetResult() *ResolveResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveResult)(nil), "geo.ResolveResult")
	proto.RegisterType((*FrameRXInfo)(nil), "geo.FrameRXInfo")
	proto.RegisterType((*ResolveTDOARequest)(nil), "geo.ResolveTDOARequest")
	proto.RegisterType((*ResolveMultiFrameTDOARequest)(nil), "geo.ResolveMultiFrameTDOARequest")
	proto.RegisterType((*ResolveTDOAResponse)(nil), "geo.ResolveTDOAResponse")
	proto.RegisterType((*ResolveMultiFrameTDOAResponse)(nil), "geo.ResolveMultiFrameTDOAResponse")
}

func init() { proto.RegisterFile("geo/geo.proto", fileDescriptor_d1967117b7b016d8) }

var fileDescriptor_d1967117b7b016d8 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x6f, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0xdd, 0x2b, 0xf4, 0x64, 0x62, 0x45, 0xf7, 0x90, 0xc6, 0xa2, 0x50, 0x23, 0xc2, 0x29,
	0xda, 0xc0, 0xe9, 0x0b, 0x51, 0x7c, 0xd1, 0xc3, 0x3f, 0x1c, 0x2a, 0xc2, 0x9e, 0x07, 0xe2, 0x9b,
	0x18, 0xd3, 0xc9, 0xb2, 0x98, 0xee, 0xd4, 0xcd, 0x26, 0xb9, 0x6f, 0xe3, 0x57, 0xf0, 0x23, 0xf8,
	0xd1, 0x24, 0x9b, 0x8d, 0xa4, 0x7a, 0x8a, 0xe2, 0x9b, 0x2e, 0x9d, 0x79, 0x66, 0x9f, 0xdf, 0xce,
	0x64, 0x60, 0x22, 0x91, 0x62, 0x89, 0xb4, 0xd8, 0x18, 0xb2, 0xc4, 0x47, 0x12, 0x69, 0x16, 0xc8,
	0x26, 0x96, 0x4d, 0x17, 0x99, 0xed, 0x65, 0xb4, 0x5e, 0x93, 0x8e, 0xbb, 0xa3, 0x0b, 0x46, 0x4f,
	0x60, 0x22, 0xb0, 0xa4, 0xa2, 0x46, 0x81, 0x65, 0x55, 0x58, 0x7e, 0x17, 0xce, 0x17, 0x94, 0xa5,
	0x56, 0x91, 0x0e, 0xd9, 0x9c, 0xed, 0x07, 0x07, 0x97, 0x16, 0xbe, 0xe2, 0x95, 0x8f, 0x8b, 0x1f,
	0x8a, 0xe8, 0x21, 0x04, 0xcf, 0x4d, 0xba, 0x46, 0xf1, 0xee, 0x48, 0xe7, 0xc4, 0x6f, 0xc3, 0xae,
	0x39, 0x4d, 0x94, 0xce, 0x29, 0x64, 0xf3, 0x91, 0xab, 0x95, 0xcd, 0xe2, 0x64, 0x53, 0x28, 0xfd,
	0xa9, 0x93, 0x88, 0xb1, 0x39, 0x6d, 0xcf, 0xe8, 0x0b, 0x03, 0xee, 0x9d, 0xdf, 0x3e, 0x7d, 0xb3,
	0x14, 0xf8, 0xb9, 0xc2, 0xd2, 0xf2, 0x29, 0xec, 0xae, 0xb0, 0x4e, 0xb0, 0x52, 0xce, 0xfd, 0x82,
	0x18, 0xaf, 0xb0, 0x7e, 0x76, 0x72, 0xc4, 0x1f, 0xc0, 0x24, 0x6f, 0x9d, 0x92, 0xde, 0x60, 0xc7,
	0xc3, 0xb5, 0x4f, 0x1e, 0x30, 0x88, 0x20, 0x1f, 0x00, 0x3d, 0x82, 0xab, 0x2b, 0xac, 0x55, 0x86,
	0x89, 0xc1, 0x1c, 0x0d, 0xea, 0x0c, 0x93, 0xb4, 0xb0, 0xca, 0x56, 0x2b, 0x0c, 0x47, 0x73, 0xb6,
	0xcf, 0xc4, 0xb4, 0x13, 0x88, 0x3e, 0xbf, 0xf4, 0xe9, 0xe8, 0x2b, 0x83, 0x6b, 0x9e, 0xf0, 0x75,
	0x55, 0x58, 0xe5, 0x4c, 0xfe, 0x8a, 0xf5, 0x31, 0x5c, 0xde, 0x62, 0x4d, 0x4a, 0xb4, 0xe1, 0x4e,
	0xdf, 0x90, 0x9f, 0x78, 0x2f, 0x0e, 0x78, 0x8f, 0xd1, 0xfe, 0x17, 0xf2, 0x12, 0xf6, 0xb6, 0x7a,
	0x5a, 0x6e, 0x48, 0x97, 0xc8, 0xef, 0xc0, 0xd8, 0xb8, 0xe9, 0xfa, 0x89, 0x72, 0x07, 0xb1, 0x35,
	0x77, 0xe1, 0x15, 0xd1, 0x4b, 0xb8, 0xfe, 0x9b, 0x47, 0xff, 0xfb, 0x65, 0x07, 0xdf, 0x18, 0x84,
	0x2f, 0x90, 0xfa, 0xcf, 0xe5, 0x18, 0x4d, 0x8d, 0xa6, 0xfd, 0x55, 0x19, 0xf2, 0x43, 0x08, 0x06,
	0xb0, 0x7c, 0x3a, 0xbc, 0x67, 0xd0, 0xe6, 0x59, 0xf8, 0x6b, 0xa2, 0x43, 0x89, 0xce, 0xf1, 0x0f,
	0x70, 0xe5, 0x4c, 0x5a, 0x7e, 0x63, 0x58, 0x74, 0xe6, 0xf8, 0x66, 0xd1, 0x9f, 0x24, 0xbd, 0xc3,
	0xe1, 0xad, 0xf7, 0x37, 0x25, 0x2d, 0x0a, 0x6a, 0x52, 0xad, 0xc8, 0xb6, 0xab, 0x10, 0xbb, 0x3f,
	0xb1, 0x2e, 0xef, 0xa5, 0x1b, 0x15, 0x4b, 0xb7, 0x74, 0x1f, 0xc7, 0x6e, 0x9d, 0xee, 0x7f, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x94, 0xc0, 0x9a, 0x86, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeolocationServerServiceClient is the client API for GeolocationServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeolocationServerServiceClient interface {
	// ResolveTDOA resolves the location based on TDOA.
	ResolveTDOA(ctx context.Context, in *ResolveTDOARequest, opts ...grpc.CallOption) (*ResolveTDOAResponse, error)
	// ResolveMultiFrameTDOA resolves the location using TDOA, based on
	// multiple frames.
	ResolveMultiFrameTDOA(ctx context.Context, in *ResolveMultiFrameTDOARequest, opts ...grpc.CallOption) (*ResolveMultiFrameTDOAResponse, error)
}

type geolocationServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeolocationServerServiceClient(cc *grpc.ClientConn) GeolocationServerServiceClient {
	return &geolocationServerServiceClient{cc}
}

func (c *geolocationServerServiceClient) ResolveTDOA(ctx context.Context, in *ResolveTDOARequest, opts ...grpc.CallOption) (*ResolveTDOAResponse, error) {
	out := new(ResolveTDOAResponse)
	err := c.cc.Invoke(ctx, "/geo.GeolocationServerService/ResolveTDOA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geolocationServerServiceClient) ResolveMultiFrameTDOA(ctx context.Context, in *ResolveMultiFrameTDOARequest, opts ...grpc.CallOption) (*ResolveMultiFrameTDOAResponse, error) {
	out := new(ResolveMultiFrameTDOAResponse)
	err := c.cc.Invoke(ctx, "/geo.GeolocationServerService/ResolveMultiFrameTDOA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeolocationServerServiceServer is the server API for GeolocationServerService service.
type GeolocationServerServiceServer interface {
	// ResolveTDOA resolves the location based on TDOA.
	ResolveTDOA(context.Context, *ResolveTDOARequest) (*ResolveTDOAResponse, error)
	// ResolveMultiFrameTDOA resolves the location using TDOA, based on
	// multiple frames.
	ResolveMultiFrameTDOA(context.Context, *ResolveMultiFrameTDOARequest) (*ResolveMultiFrameTDOAResponse, error)
}

// UnimplementedGeolocationServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeolocationServerServiceServer struct {
}

func (*UnimplementedGeolocationServerServiceServer) ResolveTDOA(ctx context.Context, req *ResolveTDOARequest) (*ResolveTDOAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveTDOA not implemented")
}
func (*UnimplementedGeolocationServerServiceServer) ResolveMultiFrameTDOA(ctx context.Context, req *ResolveMultiFrameTDOARequest) (*ResolveMultiFrameTDOAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveMultiFrameTDOA not implemented")
}

func RegisterGeolocationServerServiceServer(s *grpc.Server, srv GeolocationServerServiceServer) {
	s.RegisterService(&_GeolocationServerService_serviceDesc, srv)
}

func _GeolocationServerService_ResolveTDOA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveTDOARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeolocationServerServiceServer).ResolveTDOA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.GeolocationServerService/ResolveTDOA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeolocationServerServiceServer).ResolveTDOA(ctx, req.(*ResolveTDOARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeolocationServerService_ResolveMultiFrameTDOA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveMultiFrameTDOARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeolocationServerServiceServer).ResolveMultiFrameTDOA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.GeolocationServerService/ResolveMultiFrameTDOA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeolocationServerServiceServer).ResolveMultiFrameTDOA(ctx, req.(*ResolveMultiFrameTDOARequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeolocationServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geo.GeolocationServerService",
	HandlerType: (*GeolocationServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveTDOA",
			Handler:    _GeolocationServerService_ResolveTDOA_Handler,
		},
		{
			MethodName: "ResolveMultiFrameTDOA",
			Handler:    _GeolocationServerService_ResolveMultiFrameTDOA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geo/geo.proto",
}
