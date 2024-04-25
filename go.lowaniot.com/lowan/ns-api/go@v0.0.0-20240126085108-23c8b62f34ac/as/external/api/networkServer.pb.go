// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/networkServer.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NetworkServer struct {
	// Network-server ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Network-server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server server.
	// Format: hostname:ip (e.g. localhost:8000).
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	// CA certificate (optional).
	CaCert string `protobuf:"bytes,4,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// TLS (client) certificate for connecting to the network-server (optional).
	TlsCert string `protobuf:"bytes,5,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// TLS (client) key for connecting to the network-server (optional).
	TlsKey string `protobuf:"bytes,6,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	// Routing-profile ca certificate (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileCaCert string `protobuf:"bytes,7,opt,name=routing_profile_ca_cert,json=routingProfileCACert,proto3" json:"routing_profile_ca_cert,omitempty"`
	// Routing-profile TLS certificate (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileTlsCert string `protobuf:"bytes,8,opt,name=routing_profile_tls_cert,json=routingProfileTLSCert,proto3" json:"routing_profile_tls_cert,omitempty"`
	// Routing-profile TLS key (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileTlsKey string `protobuf:"bytes,9,opt,name=routing_profile_tls_key,json=routingProfileTLSKey,proto3" json:"routing_profile_tls_key,omitempty"`
	// Enable gateway discovery for this network-server.
	GatewayDiscoveryEnabled bool `protobuf:"varint,10,opt,name=gateway_discovery_enabled,json=gatewayDiscoveryEnabled,proto3" json:"gateway_discovery_enabled,omitempty"`
	// The number of times per day the gateway discovery 'ping' must be
	// broadcasted per gateway.
	GatewayDiscoveryInterval uint32 `protobuf:"varint,11,opt,name=gateway_discovery_interval,json=gatewayDiscoveryInterval,proto3" json:"gateway_discovery_interval,omitempty"`
	// The frequency (Hz) of the gateway discovery 'ping'.
	GatewayDiscoveryTxFrequency uint32 `protobuf:"varint,12,opt,name=gateway_discovery_tx_frequency,json=gatewayDiscoveryTXFrequency,proto3" json:"gateway_discovery_tx_frequency,omitempty"`
	// The data-rate of the gateway discovery 'ping'.
	GatewayDiscoveryDr   uint32   `protobuf:"varint,13,opt,name=gateway_discovery_dr,json=gatewayDiscoveryDR,proto3" json:"gateway_discovery_dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkServer) Reset()         { *m = NetworkServer{} }
func (m *NetworkServer) String() string { return proto.CompactTextString(m) }
func (*NetworkServer) ProtoMessage()    {}
func (*NetworkServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{0}
}

func (m *NetworkServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServer.Unmarshal(m, b)
}
func (m *NetworkServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServer.Marshal(b, m, deterministic)
}
func (m *NetworkServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServer.Merge(m, src)
}
func (m *NetworkServer) XXX_Size() int {
	return xxx_messageInfo_NetworkServer.Size(m)
}
func (m *NetworkServer) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServer.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServer proto.InternalMessageInfo

func (m *NetworkServer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServer) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *NetworkServer) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *NetworkServer) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *NetworkServer) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileCaCert() string {
	if m != nil {
		return m.RoutingProfileCaCert
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileTlsCert() string {
	if m != nil {
		return m.RoutingProfileTlsCert
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileTlsKey() string {
	if m != nil {
		return m.RoutingProfileTlsKey
	}
	return ""
}

func (m *NetworkServer) GetGatewayDiscoveryEnabled() bool {
	if m != nil {
		return m.GatewayDiscoveryEnabled
	}
	return false
}

func (m *NetworkServer) GetGatewayDiscoveryInterval() uint32 {
	if m != nil {
		return m.GatewayDiscoveryInterval
	}
	return 0
}

func (m *NetworkServer) GetGatewayDiscoveryTxFrequency() uint32 {
	if m != nil {
		return m.GatewayDiscoveryTxFrequency
	}
	return 0
}

func (m *NetworkServer) GetGatewayDiscoveryDr() uint32 {
	if m != nil {
		return m.GatewayDiscoveryDr
	}
	return 0
}

type NetworkServerListItem struct {
	// Network-server ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Network-server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server server.
	// Format: hostname:ip (e.g. localhost:8000).
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetworkServerListItem) Reset()         { *m = NetworkServerListItem{} }
func (m *NetworkServerListItem) String() string { return proto.CompactTextString(m) }
func (*NetworkServerListItem) ProtoMessage()    {}
func (*NetworkServerListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{1}
}

func (m *NetworkServerListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServerListItem.Unmarshal(m, b)
}
func (m *NetworkServerListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServerListItem.Marshal(b, m, deterministic)
}
func (m *NetworkServerListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServerListItem.Merge(m, src)
}
func (m *NetworkServerListItem) XXX_Size() int {
	return xxx_messageInfo_NetworkServerListItem.Size(m)
}
func (m *NetworkServerListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServerListItem.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServerListItem proto.InternalMessageInfo

func (m *NetworkServerListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkServerListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServerListItem) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *NetworkServerListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NetworkServerListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateNetworkServerRequest struct {
	// Network-server object to create.
	NetworkServer        *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateNetworkServerRequest) Reset()         { *m = CreateNetworkServerRequest{} }
func (m *CreateNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkServerRequest) ProtoMessage()    {}
func (*CreateNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{2}
}

func (m *CreateNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNetworkServerRequest.Unmarshal(m, b)
}
func (m *CreateNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNetworkServerRequest.Marshal(b, m, deterministic)
}
func (m *CreateNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNetworkServerRequest.Merge(m, src)
}
func (m *CreateNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNetworkServerRequest.Size(m)
}
func (m *CreateNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNetworkServerRequest proto.InternalMessageInfo

func (m *CreateNetworkServerRequest) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

type CreateNetworkServerResponse struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNetworkServerResponse) Reset()         { *m = CreateNetworkServerResponse{} }
func (m *CreateNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkServerResponse) ProtoMessage()    {}
func (*CreateNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{3}
}

func (m *CreateNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNetworkServerResponse.Unmarshal(m, b)
}
func (m *CreateNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNetworkServerResponse.Marshal(b, m, deterministic)
}
func (m *CreateNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNetworkServerResponse.Merge(m, src)
}
func (m *CreateNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNetworkServerResponse.Size(m)
}
func (m *CreateNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNetworkServerResponse proto.InternalMessageInfo

func (m *CreateNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerRequest struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkServerRequest) Reset()         { *m = GetNetworkServerRequest{} }
func (m *GetNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkServerRequest) ProtoMessage()    {}
func (*GetNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{4}
}

func (m *GetNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkServerRequest.Unmarshal(m, b)
}
func (m *GetNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkServerRequest.Marshal(b, m, deterministic)
}
func (m *GetNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkServerRequest.Merge(m, src)
}
func (m *GetNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkServerRequest.Size(m)
}
func (m *GetNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkServerRequest proto.InternalMessageInfo

func (m *GetNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerResponse struct {
	// Network-server object.
	NetworkServer *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The LoRa Server version.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The LoRa Server region configured.
	Region               string   `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkServerResponse) Reset()         { *m = GetNetworkServerResponse{} }
func (m *GetNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*GetNetworkServerResponse) ProtoMessage()    {}
func (*GetNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{5}
}

func (m *GetNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkServerResponse.Unmarshal(m, b)
}
func (m *GetNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkServerResponse.Marshal(b, m, deterministic)
}
func (m *GetNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkServerResponse.Merge(m, src)
}
func (m *GetNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_GetNetworkServerResponse.Size(m)
}
func (m *GetNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkServerResponse proto.InternalMessageInfo

func (m *GetNetworkServerResponse) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

func (m *GetNetworkServerResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetNetworkServerResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *GetNetworkServerResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetNetworkServerResponse) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type UpdateNetworkServerRequest struct {
	// Network-server object to update.
	NetworkServer        *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateNetworkServerRequest) Reset()         { *m = UpdateNetworkServerRequest{} }
func (m *UpdateNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNetworkServerRequest) ProtoMessage()    {}
func (*UpdateNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{6}
}

func (m *UpdateNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNetworkServerRequest.Unmarshal(m, b)
}
func (m *UpdateNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNetworkServerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNetworkServerRequest.Merge(m, src)
}
func (m *UpdateNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNetworkServerRequest.Size(m)
}
func (m *UpdateNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNetworkServerRequest proto.InternalMessageInfo

func (m *UpdateNetworkServerRequest) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

type DeleteNetworkServerRequest struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNetworkServerRequest) Reset()         { *m = DeleteNetworkServerRequest{} }
func (m *DeleteNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkServerRequest) ProtoMessage()    {}
func (*DeleteNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{7}
}

func (m *DeleteNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNetworkServerRequest.Unmarshal(m, b)
}
func (m *DeleteNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNetworkServerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNetworkServerRequest.Merge(m, src)
}
func (m *DeleteNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNetworkServerRequest.Size(m)
}
func (m *DeleteNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNetworkServerRequest proto.InternalMessageInfo

func (m *DeleteNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListNetworkServerRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId       int64    `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNetworkServerRequest) Reset()         { *m = ListNetworkServerRequest{} }
func (m *ListNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*ListNetworkServerRequest) ProtoMessage()    {}
func (*ListNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{8}
}

func (m *ListNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworkServerRequest.Unmarshal(m, b)
}
func (m *ListNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworkServerRequest.Marshal(b, m, deterministic)
}
func (m *ListNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworkServerRequest.Merge(m, src)
}
func (m *ListNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_ListNetworkServerRequest.Size(m)
}
func (m *ListNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworkServerRequest proto.InternalMessageInfo

func (m *ListNetworkServerRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

type ListNetworkServerResponse struct {
	// Total number of network-servers.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Network-servers within the result-set.
	Result               []*NetworkServerListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListNetworkServerResponse) Reset()         { *m = ListNetworkServerResponse{} }
func (m *ListNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*ListNetworkServerResponse) ProtoMessage()    {}
func (*ListNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c681e00a45db98, []int{9}
}

func (m *ListNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworkServerResponse.Unmarshal(m, b)
}
func (m *ListNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworkServerResponse.Marshal(b, m, deterministic)
}
func (m *ListNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworkServerResponse.Merge(m, src)
}
func (m *ListNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_ListNetworkServerResponse.Size(m)
}
func (m *ListNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworkServerResponse proto.InternalMessageInfo

func (m *ListNetworkServerResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNetworkServerResponse) GetResult() []*NetworkServerListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServer)(nil), "api.NetworkServer")
	proto.RegisterType((*NetworkServerListItem)(nil), "api.NetworkServerListItem")
	proto.RegisterType((*CreateNetworkServerRequest)(nil), "api.CreateNetworkServerRequest")
	proto.RegisterType((*CreateNetworkServerResponse)(nil), "api.CreateNetworkServerResponse")
	proto.RegisterType((*GetNetworkServerRequest)(nil), "api.GetNetworkServerRequest")
	proto.RegisterType((*GetNetworkServerResponse)(nil), "api.GetNetworkServerResponse")
	proto.RegisterType((*UpdateNetworkServerRequest)(nil), "api.UpdateNetworkServerRequest")
	proto.RegisterType((*DeleteNetworkServerRequest)(nil), "api.DeleteNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerRequest)(nil), "api.ListNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerResponse)(nil), "api.ListNetworkServerResponse")
}

func init() {
	proto.RegisterFile("as/external/api/networkServer.proto", fileDescriptor_12c681e00a45db98)
}

var fileDescriptor_12c681e00a45db98 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0x7e, 0x9a, 0xb6, 0xa7, 0xb4, 0x48, 0xa3, 0x6c, 0xe3, 0xb8, 0x65, 0x1b, 0xcc, 0x05,
	0x61, 0x45, 0x6d, 0xe8, 0x0a, 0x21, 0x56, 0xdc, 0x94, 0x74, 0x59, 0x55, 0x54, 0x08, 0xb9, 0x45,
	0x20, 0x6e, 0xac, 0xa9, 0x7d, 0x12, 0x8d, 0xd6, 0xf1, 0x78, 0x3d, 0x93, 0x74, 0x03, 0xda, 0x1b,
	0x5e, 0x81, 0xf7, 0x40, 0xbc, 0x06, 0xd7, 0xbc, 0x02, 0xaf, 0x81, 0x84, 0xe6, 0xc7, 0xa8, 0x71,
	0x6c, 0x10, 0x0b, 0x37, 0x91, 0x67, 0xce, 0xf7, 0x9d, 0x6f, 0xce, 0x39, 0xdf, 0xd8, 0x81, 0x77,
	0xa8, 0x08, 0xf0, 0xa5, 0xc4, 0x22, 0xa3, 0x69, 0x40, 0x73, 0x16, 0x64, 0x28, 0xef, 0x78, 0xf1,
	0xfc, 0x1a, 0x8b, 0x25, 0x16, 0x7e, 0x5e, 0x70, 0xc9, 0x49, 0x87, 0xe6, 0xcc, 0x3d, 0x9e, 0x71,
	0x3e, 0x4b, 0x51, 0x83, 0x68, 0x96, 0x71, 0x49, 0x25, 0xe3, 0x99, 0x30, 0x10, 0xf7, 0xc4, 0x46,
	0xf5, 0xea, 0x76, 0x31, 0x0d, 0x24, 0x9b, 0xa3, 0x90, 0x74, 0x9e, 0x5b, 0xc0, 0x51, 0x15, 0x80,
	0xf3, 0x5c, 0xae, 0x4c, 0xd0, 0xfb, 0xa5, 0x0b, 0xfb, 0x5f, 0xde, 0x17, 0x26, 0x07, 0xd0, 0x66,
	0x89, 0xd3, 0x1a, 0xb5, 0xc6, 0x9d, 0xb0, 0xcd, 0x12, 0x42, 0xa0, 0x9b, 0xd1, 0x39, 0x3a, 0xed,
	0x51, 0x6b, 0xbc, 0x1b, 0xea, 0x67, 0x72, 0x08, 0x3d, 0xa1, 0xd1, 0x4e, 0x47, 0xef, 0xda, 0x15,
	0x19, 0xc0, 0x76, 0x4c, 0xa3, 0x18, 0x0b, 0xe9, 0x74, 0x4d, 0x20, 0xa6, 0x13, 0x2c, 0x24, 0x19,
	0xc2, 0x8e, 0x4c, 0x85, 0x89, 0x6c, 0xe9, 0xc8, 0xb6, 0x4c, 0x85, 0x0e, 0x0d, 0x40, 0x3d, 0x46,
	0xcf, 0x71, 0xe5, 0xf4, 0x0c, 0x47, 0xa6, 0xe2, 0x0b, 0x5c, 0x91, 0x8f, 0x60, 0x50, 0xf0, 0x85,
	0x64, 0xd9, 0x2c, 0xca, 0x0b, 0x3e, 0x65, 0x29, 0x46, 0x65, 0xf2, 0x6d, 0x0d, 0xec, 0xdb, 0xf0,
	0x57, 0x26, 0x3a, 0x39, 0xd7, 0xf9, 0x3e, 0x06, 0xa7, 0x4a, 0xfb, 0x4b, 0x7a, 0x47, 0xf3, 0x1e,
	0xac, 0xf3, 0x6e, 0xae, 0xae, 0x35, 0xb1, 0x46, 0xaf, 0x3c, 0xd8, 0x6e, 0x9d, 0xde, 0xcd, 0xd5,
	0xb5, 0x3a, 0xe6, 0x13, 0x18, 0xce, 0xa8, 0xc4, 0x3b, 0xba, 0x8a, 0x12, 0x26, 0x62, 0xbe, 0xc4,
	0x62, 0x15, 0x61, 0x46, 0x6f, 0x53, 0x4c, 0x1c, 0x18, 0xb5, 0xc6, 0x3b, 0xe1, 0xc0, 0x02, 0x2e,
	0xca, 0xf8, 0x53, 0x13, 0x26, 0x9f, 0x82, 0xbb, 0xc9, 0x65, 0x99, 0xc4, 0x62, 0x49, 0x53, 0x67,
	0x6f, 0xd4, 0x1a, 0xef, 0x87, 0x4e, 0x95, 0x7c, 0x69, 0xe3, 0x64, 0x02, 0x0f, 0x37, 0xd9, 0xf2,
	0x65, 0x34, 0x2d, 0xf0, 0xc5, 0x02, 0xb3, 0x78, 0xe5, 0xbc, 0xa1, 0x33, 0x1c, 0x55, 0x33, 0xdc,
	0x7c, 0xfb, 0x79, 0x09, 0x21, 0x1f, 0x40, 0x7f, 0x33, 0x49, 0x52, 0x38, 0xfb, 0x9a, 0x4a, 0xaa,
	0xd4, 0x8b, 0xd0, 0xfb, 0xb5, 0x05, 0x0f, 0xd6, 0x2c, 0x73, 0xc5, 0x84, 0xbc, 0x94, 0x38, 0xff,
	0x4f, 0xd6, 0xf9, 0x04, 0x20, 0x2e, 0x90, 0x4a, 0x4c, 0x22, 0x6a, 0xdc, 0xb3, 0x77, 0xe6, 0xfa,
	0xc6, 0xba, 0x7e, 0x69, 0x5d, 0xff, 0xa6, 0xf4, 0x76, 0xb8, 0x6b, 0xd1, 0xe7, 0x52, 0x51, 0x17,
	0x79, 0x52, 0x52, 0xb7, 0xfe, 0x99, 0x6a, 0xd1, 0xe7, 0xd2, 0xfb, 0x06, 0xdc, 0x89, 0xce, 0xb3,
	0x56, 0x50, 0xa8, 0x9a, 0x23, 0x54, 0xe2, 0x03, 0x7b, 0x29, 0x23, 0x7b, 0xe6, 0x96, 0x4e, 0x4e,
	0x7c, 0x9a, 0x33, 0x7f, 0x9d, 0xb2, 0xbf, 0x76, 0x7d, 0xbd, 0x53, 0x38, 0xaa, 0x4d, 0x2c, 0x72,
	0x9e, 0x09, 0xac, 0x76, 0xca, 0x7b, 0x0f, 0x06, 0xcf, 0x50, 0xd6, 0x1e, 0xa2, 0x0a, 0xfd, 0xa3,
	0x05, 0xce, 0x26, 0xd6, 0xe6, 0x7d, 0xfd, 0x13, 0x57, 0x06, 0xd0, 0x7e, 0xfd, 0x01, 0x74, 0xfe,
	0xc5, 0x00, 0x88, 0x03, 0xdb, 0x4b, 0x2c, 0x04, 0xe3, 0x99, 0x7d, 0x63, 0x94, 0x4b, 0x65, 0x94,
	0x02, 0x67, 0x2a, 0x60, 0x5e, 0x18, 0x76, 0xa5, 0x46, 0xf6, 0xb5, 0xa6, 0xff, 0xdf, 0x23, 0x7b,
	0x1f, 0xdc, 0x0b, 0x4c, 0xb1, 0x21, 0x71, 0x75, 0x0c, 0x2f, 0xc0, 0x51, 0xbe, 0xaf, 0xc5, 0xf6,
	0x61, 0x2b, 0x65, 0x73, 0x26, 0x2d, 0xdc, 0x2c, 0x54, 0x41, 0x7c, 0x3a, 0x15, 0x68, 0x9a, 0xdb,
	0x09, 0xed, 0x8a, 0xbc, 0x0b, 0x6f, 0xf2, 0x62, 0x46, 0x33, 0xf6, 0xbd, 0x7e, 0xaf, 0x47, 0x2c,
	0xd1, 0x2d, 0xec, 0x84, 0x07, 0xf7, 0xb7, 0x2f, 0x2f, 0xbc, 0x1c, 0x86, 0x35, 0x92, 0x76, 0xf2,
	0x27, 0xb0, 0x27, 0xb9, 0xa4, 0x69, 0x14, 0xf3, 0x45, 0x56, 0x2a, 0x83, 0xde, 0x9a, 0xa8, 0x1d,
	0x72, 0xa6, 0xfa, 0x29, 0x16, 0xa9, 0x92, 0xef, 0xe8, 0x01, 0x6d, 0x74, 0xa4, 0xbc, 0xc8, 0xa1,
	0x45, 0x9e, 0xfd, 0xdc, 0x85, 0xfe, 0x1a, 0x42, 0xfd, 0xb2, 0x18, 0x49, 0x0a, 0x3d, 0x63, 0x6f,
	0x72, 0xa2, 0xd3, 0x34, 0x5f, 0x22, 0x77, 0xd4, 0x0c, 0x30, 0x47, 0xf7, 0x4e, 0x7e, 0xfc, 0xed,
	0xf7, 0x9f, 0xda, 0x43, 0xaf, 0x7f, 0xff, 0x33, 0x78, 0x6a, 0xc6, 0x27, 0x9e, 0xb4, 0x1e, 0x11,
	0x84, 0xce, 0x33, 0x94, 0xe4, 0x58, 0x67, 0x6a, 0xb8, 0x27, 0xee, 0x5b, 0x0d, 0x51, 0x2b, 0xf2,
	0xb6, 0x16, 0x39, 0x22, 0xc3, 0x3a, 0x91, 0xe0, 0x07, 0x96, 0xbc, 0x22, 0x4b, 0xe8, 0x19, 0x67,
	0xd9, 0xa2, 0x9a, 0x6d, 0xe6, 0x1e, 0x6e, 0xb8, 0xfb, 0xa9, 0xfa, 0xa8, 0x7a, 0x8f, 0xb5, 0xca,
	0xa9, 0x3b, 0xae, 0x57, 0x59, 0xb7, 0xa6, 0xcf, 0x92, 0x57, 0xaa, 0xbc, 0x04, 0x7a, 0xc6, 0x78,
	0x56, 0xb7, 0xd9, 0x85, 0x8d, 0xba, 0xb6, 0xba, 0x47, 0x7f, 0x53, 0x5d, 0x0c, 0x5d, 0x35, 0x5f,
	0x62, 0xfa, 0xd4, 0xe4, 0x5d, 0xf7, 0x61, 0x53, 0xd8, 0xf6, 0xf1, 0x58, 0x2b, 0x1d, 0x92, 0xda,
	0x61, 0x7d, 0xf6, 0xe1, 0x77, 0xc1, 0x8c, 0xfb, 0x29, 0xbf, 0xa3, 0x19, 0xe3, 0xd2, 0x8f, 0xf9,
	0x3c, 0xd0, 0x8b, 0x20, 0x13, 0xa7, 0x0a, 0x3f, 0xe3, 0x41, 0xe5, 0x6f, 0xcf, 0x6d, 0x4f, 0x97,
	0xf2, 0xf8, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x8f, 0x45, 0xbb, 0x10, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServerServiceClient is the client API for NetworkServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServerServiceClient interface {
	// Create creates the given network-server.
	Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the network-server matching the given id.
	Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available network-servers.
	List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error)
}

type networkServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerServiceClient(cc *grpc.ClientConn) NetworkServerServiceClient {
	return &networkServerServiceClient{cc}
}

func (c *networkServerServiceClient) Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error) {
	out := new(CreateNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error) {
	out := new(GetNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error) {
	out := new(ListNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServerServiceServer is the server API for NetworkServerService service.
type NetworkServerServiceServer interface {
	// Create creates the given network-server.
	Create(context.Context, *CreateNetworkServerRequest) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(context.Context, *GetNetworkServerRequest) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(context.Context, *UpdateNetworkServerRequest) (*empty.Empty, error)
	// Delete deletes the network-server matching the given id.
	Delete(context.Context, *DeleteNetworkServerRequest) (*empty.Empty, error)
	// List lists the available network-servers.
	List(context.Context, *ListNetworkServerRequest) (*ListNetworkServerResponse, error)
}

// UnimplementedNetworkServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServerServiceServer struct {
}

func (*UnimplementedNetworkServerServiceServer) Create(ctx context.Context, req *CreateNetworkServerRequest) (*CreateNetworkServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNetworkServerServiceServer) Get(ctx context.Context, req *GetNetworkServerRequest) (*GetNetworkServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedNetworkServerServiceServer) Update(ctx context.Context, req *UpdateNetworkServerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedNetworkServerServiceServer) Delete(ctx context.Context, req *DeleteNetworkServerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedNetworkServerServiceServer) List(ctx context.Context, req *ListNetworkServerRequest) (*ListNetworkServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterNetworkServerServiceServer(s *grpc.Server, srv NetworkServerServiceServer) {
	s.RegisterService(&_NetworkServerService_serviceDesc, srv)
}

func _NetworkServerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Create(ctx, req.(*CreateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Get(ctx, req.(*GetNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Update(ctx, req.(*UpdateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Delete(ctx, req.(*DeleteNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).List(ctx, req.(*ListNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkServerService",
	HandlerType: (*NetworkServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkServerService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetworkServerService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkServerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkServerService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkServerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/networkServer.proto",
}
