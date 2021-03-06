// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flowlog/flowlog.proto

package flowlog

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

type ACLAction int32

const (
	ACLAction_ActionUnknown ACLAction = 0
	ACLAction_ActionDrop    ACLAction = 1
	ACLAction_ActionAccept  ACLAction = 2
)

var ACLAction_name = map[int32]string{
	0: "ActionUnknown",
	1: "ActionDrop",
	2: "ActionAccept",
}

var ACLAction_value = map[string]int32{
	"ActionUnknown": 0,
	"ActionDrop":    1,
	"ActionAccept":  2,
}

func (x ACLAction) String() string {
	return proto.EnumName(ACLAction_name, int32(x))
}

func (ACLAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{0}
}

type IpFlow struct {
	Src                  string   `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	SrcPort              int32    `protobuf:"varint,2,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	Dest                 string   `protobuf:"bytes,3,opt,name=dest,proto3" json:"dest,omitempty"`
	DestPort             int32    `protobuf:"varint,4,opt,name=destPort,proto3" json:"destPort,omitempty"`
	Protocol             int32    `protobuf:"varint,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpFlow) Reset()         { *m = IpFlow{} }
func (m *IpFlow) String() string { return proto.CompactTextString(m) }
func (*IpFlow) ProtoMessage()    {}
func (*IpFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{0}
}

func (m *IpFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpFlow.Unmarshal(m, b)
}
func (m *IpFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpFlow.Marshal(b, m, deterministic)
}
func (m *IpFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpFlow.Merge(m, src)
}
func (m *IpFlow) XXX_Size() int {
	return xxx_messageInfo_IpFlow.Size(m)
}
func (m *IpFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_IpFlow.DiscardUnknown(m)
}

var xxx_messageInfo_IpFlow proto.InternalMessageInfo

func (m *IpFlow) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *IpFlow) GetSrcPort() int32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *IpFlow) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

func (m *IpFlow) GetDestPort() int32 {
	if m != nil {
		return m.DestPort
	}
	return 0
}

func (m *IpFlow) GetProtocol() int32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

type ScopeInfo struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Intf                 string   `protobuf:"bytes,2,opt,name=intf,proto3" json:"intf,omitempty"`
	LocalIntf            string   `protobuf:"bytes,3,opt,name=localIntf,proto3" json:"localIntf,omitempty"`
	NetInstUUID          string   `protobuf:"bytes,4,opt,name=netInstUUID,proto3" json:"netInstUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScopeInfo) Reset()         { *m = ScopeInfo{} }
func (m *ScopeInfo) String() string { return proto.CompactTextString(m) }
func (*ScopeInfo) ProtoMessage()    {}
func (*ScopeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{1}
}

func (m *ScopeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScopeInfo.Unmarshal(m, b)
}
func (m *ScopeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScopeInfo.Marshal(b, m, deterministic)
}
func (m *ScopeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopeInfo.Merge(m, src)
}
func (m *ScopeInfo) XXX_Size() int {
	return xxx_messageInfo_ScopeInfo.Size(m)
}
func (m *ScopeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ScopeInfo proto.InternalMessageInfo

func (m *ScopeInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ScopeInfo) GetIntf() string {
	if m != nil {
		return m.Intf
	}
	return ""
}

func (m *ScopeInfo) GetLocalIntf() string {
	if m != nil {
		return m.LocalIntf
	}
	return ""
}

func (m *ScopeInfo) GetNetInstUUID() string {
	if m != nil {
		return m.NetInstUUID
	}
	return ""
}

type FlowRecord struct {
	Flow    *IpFlow `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Inbound bool    `protobuf:"varint,2,opt,name=inbound,proto3" json:"inbound,omitempty"`
	AclId   int32   `protobuf:"varint,3,opt,name=aclId,proto3" json:"aclId,omitempty"`
	AclName string  `protobuf:"bytes,4,opt,name=aclName,proto3" json:"aclName,omitempty"`
	// deprecated = 5;
	StartTime            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=endTime,proto3" json:"endTime,omitempty"`
	TxBytes              int64                `protobuf:"varint,8,opt,name=txBytes,proto3" json:"txBytes,omitempty"`
	TxPkts               int64                `protobuf:"varint,9,opt,name=txPkts,proto3" json:"txPkts,omitempty"`
	RxBytes              int64                `protobuf:"varint,10,opt,name=rxBytes,proto3" json:"rxBytes,omitempty"`
	RxPkts               int64                `protobuf:"varint,11,opt,name=rxPkts,proto3" json:"rxPkts,omitempty"`
	Action               ACLAction            `protobuf:"varint,12,opt,name=action,proto3,enum=ACLAction" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlowRecord) Reset()         { *m = FlowRecord{} }
func (m *FlowRecord) String() string { return proto.CompactTextString(m) }
func (*FlowRecord) ProtoMessage()    {}
func (*FlowRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{2}
}

func (m *FlowRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRecord.Unmarshal(m, b)
}
func (m *FlowRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRecord.Marshal(b, m, deterministic)
}
func (m *FlowRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRecord.Merge(m, src)
}
func (m *FlowRecord) XXX_Size() int {
	return xxx_messageInfo_FlowRecord.Size(m)
}
func (m *FlowRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRecord.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRecord proto.InternalMessageInfo

func (m *FlowRecord) GetFlow() *IpFlow {
	if m != nil {
		return m.Flow
	}
	return nil
}

func (m *FlowRecord) GetInbound() bool {
	if m != nil {
		return m.Inbound
	}
	return false
}

func (m *FlowRecord) GetAclId() int32 {
	if m != nil {
		return m.AclId
	}
	return 0
}

func (m *FlowRecord) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *FlowRecord) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *FlowRecord) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *FlowRecord) GetTxBytes() int64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *FlowRecord) GetTxPkts() int64 {
	if m != nil {
		return m.TxPkts
	}
	return 0
}

func (m *FlowRecord) GetRxBytes() int64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *FlowRecord) GetRxPkts() int64 {
	if m != nil {
		return m.RxPkts
	}
	return 0
}

func (m *FlowRecord) GetAction() ACLAction {
	if m != nil {
		return m.Action
	}
	return ACLAction_ActionUnknown
}

type DnsRequest struct {
	HostName             string               `protobuf:"bytes,1,opt,name=hostName,proto3" json:"hostName,omitempty"`
	Addrs                []string             `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	RequestTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=requestTime,proto3" json:"requestTime,omitempty"`
	AclNum               int32                `protobuf:"varint,4,opt,name=aclNum,proto3" json:"aclNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DnsRequest) Reset()         { *m = DnsRequest{} }
func (m *DnsRequest) String() string { return proto.CompactTextString(m) }
func (*DnsRequest) ProtoMessage()    {}
func (*DnsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{3}
}

func (m *DnsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsRequest.Unmarshal(m, b)
}
func (m *DnsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsRequest.Marshal(b, m, deterministic)
}
func (m *DnsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsRequest.Merge(m, src)
}
func (m *DnsRequest) XXX_Size() int {
	return xxx_messageInfo_DnsRequest.Size(m)
}
func (m *DnsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DnsRequest proto.InternalMessageInfo

func (m *DnsRequest) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *DnsRequest) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *DnsRequest) GetRequestTime() *timestamp.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *DnsRequest) GetAclNum() int32 {
	if m != nil {
		return m.AclNum
	}
	return 0
}

// This is the request payload for POST /api/v1/edgeDevice/flowlog
// FlowMessage carries device logs to the controller.
// The message is assumed to be protected by a TLS session bound to the
// device certificate.
type FlowMessage struct {
	DevId                string        `protobuf:"bytes,1,opt,name=devId,proto3" json:"devId,omitempty"`
	Scope                *ScopeInfo    `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Flows                []*FlowRecord `protobuf:"bytes,3,rep,name=flows,proto3" json:"flows,omitempty"`
	DnsReqs              []*DnsRequest `protobuf:"bytes,4,rep,name=dnsReqs,proto3" json:"dnsReqs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FlowMessage) Reset()         { *m = FlowMessage{} }
func (m *FlowMessage) String() string { return proto.CompactTextString(m) }
func (*FlowMessage) ProtoMessage()    {}
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_469cd0148fb841a1, []int{4}
}

func (m *FlowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowMessage.Unmarshal(m, b)
}
func (m *FlowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowMessage.Marshal(b, m, deterministic)
}
func (m *FlowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowMessage.Merge(m, src)
}
func (m *FlowMessage) XXX_Size() int {
	return xxx_messageInfo_FlowMessage.Size(m)
}
func (m *FlowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FlowMessage proto.InternalMessageInfo

func (m *FlowMessage) GetDevId() string {
	if m != nil {
		return m.DevId
	}
	return ""
}

func (m *FlowMessage) GetScope() *ScopeInfo {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *FlowMessage) GetFlows() []*FlowRecord {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *FlowMessage) GetDnsReqs() []*DnsRequest {
	if m != nil {
		return m.DnsReqs
	}
	return nil
}

func init() {
	proto.RegisterEnum("ACLAction", ACLAction_name, ACLAction_value)
	proto.RegisterType((*IpFlow)(nil), "IpFlow")
	proto.RegisterType((*ScopeInfo)(nil), "ScopeInfo")
	proto.RegisterType((*FlowRecord)(nil), "FlowRecord")
	proto.RegisterType((*DnsRequest)(nil), "DnsRequest")
	proto.RegisterType((*FlowMessage)(nil), "FlowMessage")
}

func init() {
	proto.RegisterFile("flowlog/flowlog.proto", fileDescriptor_469cd0148fb841a1)
}

var fileDescriptor_469cd0148fb841a1 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0x7e, 0x8e, 0xf3, 0xa3, 0x1e, 0xf7, 0x55, 0x79, 0xab, 0x07, 0xb2, 0x0a, 0x52, 0x4d, 0xa4,
	0x4a, 0x11, 0x12, 0xb6, 0x14, 0x38, 0x70, 0xe0, 0x40, 0x4b, 0x85, 0xb0, 0x04, 0xa8, 0x5a, 0xda,
	0x0b, 0x37, 0x67, 0x77, 0xe3, 0x5a, 0xb5, 0xbd, 0x66, 0x77, 0xdd, 0x16, 0x4e, 0xfc, 0x07, 0x5c,
	0xf8, 0x2f, 0xf9, 0x27, 0xd0, 0xce, 0xda, 0x49, 0x6f, 0x3d, 0xed, 0x7c, 0x33, 0xdf, 0x97, 0x99,
	0x7c, 0x33, 0x32, 0x3c, 0xda, 0x54, 0xf2, 0xb6, 0x92, 0x45, 0xda, 0xbf, 0x49, 0xab, 0xa4, 0x91,
	0x87, 0x47, 0x85, 0x94, 0x45, 0x25, 0x52, 0x44, 0xeb, 0x6e, 0x93, 0x9a, 0xb2, 0x16, 0xda, 0xe4,
	0x75, 0xeb, 0x08, 0x8b, 0x9f, 0x1e, 0x4c, 0xb3, 0xf6, 0x7d, 0x25, 0x6f, 0xc9, 0x1c, 0x7c, 0xad,
	0x58, 0xe4, 0xc5, 0xde, 0x32, 0xa0, 0x36, 0x24, 0x11, 0xcc, 0xb4, 0x62, 0xe7, 0x52, 0x99, 0x68,
	0x14, 0x7b, 0xcb, 0x09, 0x1d, 0x20, 0x21, 0x30, 0xe6, 0x42, 0x9b, 0xc8, 0x47, 0x32, 0xc6, 0xe4,
	0x10, 0xf6, 0xec, 0x8b, 0xf4, 0x31, 0xd2, 0xb7, 0xd8, 0xd6, 0xb0, 0x1f, 0x93, 0x55, 0x34, 0x71,
	0xb5, 0x01, 0x2f, 0x34, 0x04, 0x5f, 0x98, 0x6c, 0x45, 0xd6, 0x6c, 0xa4, 0xfd, 0xe1, 0xae, 0x2b,
	0x79, 0x3f, 0x05, 0xc6, 0x36, 0x57, 0x36, 0x66, 0x83, 0x33, 0x04, 0x14, 0x63, 0xf2, 0x14, 0x82,
	0x4a, 0xb2, 0xbc, 0xca, 0x6c, 0xc1, 0x4d, 0xb1, 0x4b, 0x90, 0x18, 0xc2, 0x46, 0x98, 0xac, 0xd1,
	0xe6, 0xf2, 0x32, 0x3b, 0xc3, 0x69, 0x02, 0x7a, 0x3f, 0xb5, 0xf8, 0x33, 0x02, 0xb0, 0xff, 0x9a,
	0x0a, 0x26, 0x15, 0x27, 0x4f, 0x60, 0x6c, 0x8d, 0xc3, 0xb6, 0xe1, 0x6a, 0x96, 0x38, 0x4b, 0x28,
	0x26, 0xad, 0x0d, 0x65, 0xb3, 0x96, 0x5d, 0xc3, 0x71, 0x84, 0x3d, 0x3a, 0x40, 0xf2, 0x3f, 0x4c,
	0x72, 0x56, 0x65, 0x1c, 0x27, 0x98, 0x50, 0x07, 0x2c, 0x3f, 0x67, 0xd5, 0xe7, 0xbc, 0x16, 0x7d,
	0xe7, 0x01, 0x92, 0xd7, 0x10, 0x68, 0x93, 0x2b, 0x73, 0x51, 0xd6, 0x22, 0x9a, 0x62, 0xaf, 0xc3,
	0xc4, 0xad, 0x28, 0x19, 0x56, 0x94, 0x5c, 0x0c, 0x2b, 0xa2, 0x3b, 0x32, 0x79, 0x05, 0x33, 0xd1,
	0x70, 0xd4, 0xcd, 0x1e, 0xd4, 0x0d, 0x54, 0x3b, 0x89, 0xb9, 0x3b, 0xfd, 0x6e, 0x84, 0x8e, 0xf6,
	0x62, 0x6f, 0xe9, 0xd3, 0x01, 0x92, 0xc7, 0x30, 0x35, 0x77, 0xe7, 0xd7, 0x46, 0x47, 0x01, 0x16,
	0x7a, 0x64, 0x15, 0xaa, 0x57, 0x80, 0x53, 0xa8, 0x9d, 0x42, 0x39, 0x45, 0xe8, 0x14, 0x0e, 0x91,
	0x05, 0x4c, 0x73, 0x66, 0x4a, 0xd9, 0x44, 0xfb, 0xb1, 0xb7, 0x3c, 0x58, 0x41, 0x72, 0xf2, 0xee,
	0xe3, 0x09, 0x66, 0x68, 0x5f, 0x59, 0xfc, 0xf6, 0x00, 0xce, 0x1a, 0x4d, 0xc5, 0xb7, 0xae, 0xbf,
	0x94, 0x2b, 0xa9, 0x0d, 0x3a, 0xe4, 0x16, 0xbd, 0xc5, 0x68, 0x29, 0xe7, 0x4a, 0x47, 0xa3, 0xd8,
	0x5f, 0x06, 0xd4, 0x01, 0xf2, 0x06, 0x42, 0xe5, 0xc4, 0x68, 0x81, 0xff, 0xa0, 0x05, 0xf7, 0xe9,
	0x76, 0x74, 0xbb, 0x81, 0xae, 0xee, 0xef, 0xb2, 0x47, 0x8b, 0x5f, 0x1e, 0x84, 0x76, 0xcf, 0x9f,
	0x84, 0xd6, 0x79, 0x81, 0xbd, 0xb9, 0xb8, 0xc9, 0x86, 0xeb, 0x73, 0x80, 0xc4, 0x30, 0xd1, 0xf6,
	0x3e, 0x71, 0xf9, 0xe1, 0x0a, 0x92, 0xed, 0xb5, 0x52, 0x57, 0x20, 0xcf, 0x60, 0x62, 0x0f, 0x45,
	0x47, 0x7e, 0xec, 0x2f, 0xc3, 0x55, 0x98, 0xec, 0x2e, 0x8b, 0xba, 0x0a, 0x39, 0x86, 0x19, 0x47,
	0x03, 0x74, 0x34, 0xee, 0x49, 0x3b, 0x43, 0xe8, 0x50, 0x7b, 0xfe, 0x16, 0x82, 0xad, 0x7b, 0xe4,
	0x3f, 0xf8, 0xd7, 0x45, 0x97, 0xcd, 0x75, 0x23, 0x6f, 0x9b, 0xf9, 0x3f, 0xe4, 0x00, 0xc0, 0xa5,
	0xce, 0x94, 0x6c, 0xe7, 0x1e, 0x99, 0xc3, 0xbe, 0xc3, 0x27, 0x8c, 0x89, 0xd6, 0xcc, 0x47, 0xa7,
	0x1f, 0xe0, 0x88, 0xc9, 0x3a, 0xf9, 0x21, 0xb8, 0xe0, 0x79, 0xc2, 0x2a, 0xd9, 0xf1, 0xa4, 0xd3,
	0x42, 0xdd, 0x94, 0xac, 0x37, 0xea, 0xeb, 0x71, 0x51, 0x9a, 0xab, 0x6e, 0x9d, 0x30, 0x59, 0xa7,
	0xd5, 0xe6, 0x85, 0xe0, 0x85, 0x48, 0xc5, 0x8d, 0x48, 0xf3, 0xb6, 0x4c, 0x0b, 0x39, 0x7c, 0x41,
	0xd6, 0x53, 0x64, 0xbf, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x47, 0x38, 0xc8, 0x5b, 0x04,
	0x00, 0x00,
}
