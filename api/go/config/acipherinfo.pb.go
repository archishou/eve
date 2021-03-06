// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/acipherinfo.proto

package config

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

// Security Key Exchange Method
type KeyExchangeScheme int32

const (
	KeyExchangeScheme_KEA_NONE KeyExchangeScheme = 0
	KeyExchangeScheme_KEA_ECDH KeyExchangeScheme = 1
)

var KeyExchangeScheme_name = map[int32]string{
	0: "KEA_NONE",
	1: "KEA_ECDH",
}

var KeyExchangeScheme_value = map[string]int32{
	"KEA_NONE": 0,
	"KEA_ECDH": 1,
}

func (x KeyExchangeScheme) String() string {
	return proto.EnumName(KeyExchangeScheme_name, int32(x))
}

func (KeyExchangeScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{0}
}

// Encryption Scheme for Cipher Payload
type EncryptionScheme int32

const (
	EncryptionScheme_SA_NONE        EncryptionScheme = 0
	EncryptionScheme_SA_AES_256_CFB EncryptionScheme = 1
)

var EncryptionScheme_name = map[int32]string{
	0: "SA_NONE",
	1: "SA_AES_256_CFB",
}

var EncryptionScheme_value = map[string]int32{
	"SA_NONE":        0,
	"SA_AES_256_CFB": 1,
}

func (x EncryptionScheme) String() string {
	return proto.EnumName(EncryptionScheme_name, int32(x))
}

func (EncryptionScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{1}
}

type CipherHashAlgorithm int32

const (
	CipherHashAlgorithm_HASH_NONE           CipherHashAlgorithm = 0
	CipherHashAlgorithm_HASH_SHA256_16bytes CipherHashAlgorithm = 1
	CipherHashAlgorithm_HASH_SHA256_32bytes CipherHashAlgorithm = 2
)

var CipherHashAlgorithm_name = map[int32]string{
	0: "HASH_NONE",
	1: "HASH_SHA256_16bytes",
	2: "HASH_SHA256_32bytes",
}

var CipherHashAlgorithm_value = map[string]int32{
	"HASH_NONE":           0,
	"HASH_SHA256_16bytes": 1,
	"HASH_SHA256_32bytes": 2,
}

func (x CipherHashAlgorithm) String() string {
	return proto.EnumName(CipherHashAlgorithm_name, int32(x))
}

func (CipherHashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{2}
}

// Cipher information to decrypt Sensitive Data
type CipherContext struct {
	// cipher context id, key to this structure
	ContextId string `protobuf:"bytes,1,opt,name=contextId,proto3" json:"contextId,omitempty"`
	// algorithm used to compute hash for certificates
	HashScheme CipherHashAlgorithm `protobuf:"varint,2,opt,name=hashScheme,proto3,enum=CipherHashAlgorithm" json:"hashScheme,omitempty"`
	// for key exchange scheme, like ECDH etc.
	KeyExchangeScheme KeyExchangeScheme `protobuf:"varint,3,opt,name=keyExchangeScheme,proto3,enum=KeyExchangeScheme" json:"keyExchangeScheme,omitempty"`
	// for encrypting sensitive data, like AES256 etc.
	EncryptionScheme EncryptionScheme `protobuf:"varint,4,opt,name=encryptionScheme,proto3,enum=EncryptionScheme" json:"encryptionScheme,omitempty"`
	// device public certificate hash
	DeviceCertHash []byte `protobuf:"bytes,5,opt,name=deviceCertHash,proto3" json:"deviceCertHash,omitempty"`
	// controller certificate hash
	ControllerCertHash   []byte   `protobuf:"bytes,6,opt,name=controllerCertHash,proto3" json:"controllerCertHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherContext) Reset()         { *m = CipherContext{} }
func (m *CipherContext) String() string { return proto.CompactTextString(m) }
func (*CipherContext) ProtoMessage()    {}
func (*CipherContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{0}
}

func (m *CipherContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherContext.Unmarshal(m, b)
}
func (m *CipherContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherContext.Marshal(b, m, deterministic)
}
func (m *CipherContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherContext.Merge(m, src)
}
func (m *CipherContext) XXX_Size() int {
	return xxx_messageInfo_CipherContext.Size(m)
}
func (m *CipherContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherContext.DiscardUnknown(m)
}

var xxx_messageInfo_CipherContext proto.InternalMessageInfo

func (m *CipherContext) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *CipherContext) GetHashScheme() CipherHashAlgorithm {
	if m != nil {
		return m.HashScheme
	}
	return CipherHashAlgorithm_HASH_NONE
}

func (m *CipherContext) GetKeyExchangeScheme() KeyExchangeScheme {
	if m != nil {
		return m.KeyExchangeScheme
	}
	return KeyExchangeScheme_KEA_NONE
}

func (m *CipherContext) GetEncryptionScheme() EncryptionScheme {
	if m != nil {
		return m.EncryptionScheme
	}
	return EncryptionScheme_SA_NONE
}

func (m *CipherContext) GetDeviceCertHash() []byte {
	if m != nil {
		return m.DeviceCertHash
	}
	return nil
}

func (m *CipherContext) GetControllerCertHash() []byte {
	if m != nil {
		return m.ControllerCertHash
	}
	return nil
}

// Encrypted sensitive data information
type CipherBlock struct {
	// cipher context id
	CipherContextId string `protobuf:"bytes,1,opt,name=cipherContextId,proto3" json:"cipherContextId,omitempty"`
	// Initial Value for Symmetric Key derivation
	InitialValue []byte `protobuf:"bytes,2,opt,name=initialValue,proto3" json:"initialValue,omitempty"`
	// encrypted sensitive data
	CipherData []byte `protobuf:"bytes,3,opt,name=cipherData,proto3" json:"cipherData,omitempty"`
	// sha256 of the plaintext sensitive data
	ClearTextSha256      []byte   `protobuf:"bytes,4,opt,name=clearTextSha256,proto3" json:"clearTextSha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherBlock) Reset()         { *m = CipherBlock{} }
func (m *CipherBlock) String() string { return proto.CompactTextString(m) }
func (*CipherBlock) ProtoMessage()    {}
func (*CipherBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{1}
}

func (m *CipherBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherBlock.Unmarshal(m, b)
}
func (m *CipherBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherBlock.Marshal(b, m, deterministic)
}
func (m *CipherBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherBlock.Merge(m, src)
}
func (m *CipherBlock) XXX_Size() int {
	return xxx_messageInfo_CipherBlock.Size(m)
}
func (m *CipherBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CipherBlock proto.InternalMessageInfo

func (m *CipherBlock) GetCipherContextId() string {
	if m != nil {
		return m.CipherContextId
	}
	return ""
}

func (m *CipherBlock) GetInitialValue() []byte {
	if m != nil {
		return m.InitialValue
	}
	return nil
}

func (m *CipherBlock) GetCipherData() []byte {
	if m != nil {
		return m.CipherData
	}
	return nil
}

func (m *CipherBlock) GetClearTextSha256() []byte {
	if m != nil {
		return m.ClearTextSha256
	}
	return nil
}

// This message will be filled with the
// details to be encrypted and shared across
// wire for data in transit, by the controller
// for encryption
type EncryptionBlock struct {
	DsAPIKey             string   `protobuf:"bytes,1,opt,name=dsAPIKey,proto3" json:"dsAPIKey,omitempty"`
	DsPassword           string   `protobuf:"bytes,2,opt,name=dsPassword,proto3" json:"dsPassword,omitempty"`
	WifiUserName         string   `protobuf:"bytes,3,opt,name=wifiUserName,proto3" json:"wifiUserName,omitempty"`
	WifiPassword         string   `protobuf:"bytes,4,opt,name=wifiPassword,proto3" json:"wifiPassword,omitempty"`
	ProtectedUserData    string   `protobuf:"bytes,5,opt,name=protectedUserData,proto3" json:"protectedUserData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptionBlock) Reset()         { *m = EncryptionBlock{} }
func (m *EncryptionBlock) String() string { return proto.CompactTextString(m) }
func (*EncryptionBlock) ProtoMessage()    {}
func (*EncryptionBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e17e005087ff211f, []int{2}
}

func (m *EncryptionBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionBlock.Unmarshal(m, b)
}
func (m *EncryptionBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionBlock.Marshal(b, m, deterministic)
}
func (m *EncryptionBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionBlock.Merge(m, src)
}
func (m *EncryptionBlock) XXX_Size() int {
	return xxx_messageInfo_EncryptionBlock.Size(m)
}
func (m *EncryptionBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionBlock.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionBlock proto.InternalMessageInfo

func (m *EncryptionBlock) GetDsAPIKey() string {
	if m != nil {
		return m.DsAPIKey
	}
	return ""
}

func (m *EncryptionBlock) GetDsPassword() string {
	if m != nil {
		return m.DsPassword
	}
	return ""
}

func (m *EncryptionBlock) GetWifiUserName() string {
	if m != nil {
		return m.WifiUserName
	}
	return ""
}

func (m *EncryptionBlock) GetWifiPassword() string {
	if m != nil {
		return m.WifiPassword
	}
	return ""
}

func (m *EncryptionBlock) GetProtectedUserData() string {
	if m != nil {
		return m.ProtectedUserData
	}
	return ""
}

func init() {
	proto.RegisterEnum("KeyExchangeScheme", KeyExchangeScheme_name, KeyExchangeScheme_value)
	proto.RegisterEnum("EncryptionScheme", EncryptionScheme_name, EncryptionScheme_value)
	proto.RegisterEnum("CipherHashAlgorithm", CipherHashAlgorithm_name, CipherHashAlgorithm_value)
	proto.RegisterType((*CipherContext)(nil), "CipherContext")
	proto.RegisterType((*CipherBlock)(nil), "CipherBlock")
	proto.RegisterType((*EncryptionBlock)(nil), "EncryptionBlock")
}

func init() {
	proto.RegisterFile("config/acipherinfo.proto", fileDescriptor_e17e005087ff211f)
}

var fileDescriptor_e17e005087ff211f = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xa5, 0x2d, 0xcd, 0xd4, 0x6d, 0x9d, 0x2d, 0x12, 0x11, 0x42, 0x50, 0x45, 0x08, 0x45,
	0x15, 0xd8, 0x22, 0xa1, 0xb9, 0x21, 0xe1, 0xa4, 0x86, 0x54, 0x91, 0x42, 0x65, 0x43, 0x0f, 0x5c,
	0xa2, 0xcd, 0x7a, 0x62, 0xaf, 0xea, 0x78, 0x23, 0x7b, 0xd3, 0x26, 0xbc, 0x0f, 0x6f, 0xc1, 0x8d,
	0x17, 0x43, 0x5e, 0x27, 0x4e, 0xea, 0xf4, 0xe6, 0xf9, 0x7e, 0x66, 0x67, 0x3f, 0xcf, 0x42, 0x8d,
	0x89, 0x78, 0xcc, 0x03, 0x8b, 0x32, 0x3e, 0x0d, 0x31, 0xe1, 0xf1, 0x58, 0x98, 0xd3, 0x44, 0x48,
	0x51, 0xff, 0xbb, 0x03, 0x47, 0x5d, 0x05, 0x76, 0x45, 0x2c, 0x71, 0x2e, 0xc9, 0x2b, 0xa8, 0xb0,
	0xfc, 0xf3, 0xca, 0xaf, 0x69, 0x67, 0x5a, 0xa3, 0xe2, 0xae, 0x01, 0xf2, 0x09, 0x20, 0xa4, 0x69,
	0xe8, 0xb1, 0x10, 0x27, 0x58, 0xdb, 0x39, 0xd3, 0x1a, 0xc7, 0xcd, 0xe7, 0x66, 0xde, 0xb6, 0x47,
	0xd3, 0xd0, 0x8e, 0x02, 0x91, 0x70, 0x19, 0x4e, 0xdc, 0x0d, 0x1d, 0xf9, 0x02, 0xd5, 0x5b, 0x5c,
	0x38, 0x73, 0x16, 0xd2, 0x38, 0xc0, 0xa5, 0xf9, 0xa9, 0x32, 0x13, 0xb3, 0x5f, 0x66, 0xdc, 0x6d,
	0x31, 0xf9, 0x0c, 0x06, 0xc6, 0x2c, 0x59, 0x4c, 0x25, 0x17, 0xf1, 0xb2, 0xc1, 0xae, 0x6a, 0x50,
	0x35, 0x9d, 0x12, 0xe1, 0x6e, 0x49, 0xc9, 0x3b, 0x38, 0xf6, 0xf1, 0x8e, 0x33, 0xec, 0x62, 0x22,
	0xb3, 0x39, 0x6b, 0x7b, 0x67, 0x5a, 0x43, 0x77, 0x4b, 0x28, 0x31, 0x81, 0x64, 0x77, 0x4d, 0x44,
	0x14, 0x61, 0x52, 0x68, 0xf7, 0x95, 0xf6, 0x11, 0xa6, 0xfe, 0x47, 0x83, 0xc3, 0x3c, 0xbe, 0x4e,
	0x24, 0xd8, 0x2d, 0x69, 0xc0, 0x09, 0xdb, 0x4c, 0xb3, 0x88, 0xb0, 0x0c, 0x93, 0x3a, 0xe8, 0x3c,
	0xe6, 0x92, 0xd3, 0xe8, 0x86, 0x46, 0xb3, 0x3c, 0x4a, 0xdd, 0x7d, 0x80, 0x91, 0xd7, 0x00, 0xb9,
	0xed, 0x92, 0x4a, 0xaa, 0xf2, 0xd2, 0xdd, 0x0d, 0x44, 0x9d, 0x16, 0x21, 0x4d, 0x7e, 0xe0, 0x5c,
	0x7a, 0x21, 0x6d, 0x5e, 0xb4, 0x55, 0x26, 0xba, 0x5b, 0x86, 0xeb, 0xff, 0x34, 0x38, 0x59, 0xc7,
	0x94, 0xcf, 0xfa, 0x12, 0x0e, 0xfc, 0xd4, 0xbe, 0xbe, 0xea, 0xe3, 0x62, 0x39, 0x64, 0x51, 0x67,
	0x27, 0xfb, 0xe9, 0x35, 0x4d, 0xd3, 0x7b, 0x91, 0xf8, 0x6a, 0xb6, 0x8a, 0xbb, 0x81, 0x64, 0xd3,
	0xdf, 0xf3, 0x31, 0xff, 0x99, 0x62, 0x32, 0xa0, 0xcb, 0x7f, 0x59, 0x71, 0x1f, 0x60, 0x2b, 0x4d,
	0xd1, 0x65, 0x77, 0xad, 0x29, 0xfa, 0xbc, 0x87, 0x6a, 0xb6, 0x87, 0xc8, 0x24, 0xfa, 0x99, 0x51,
	0x5d, 0x74, 0x4f, 0x09, 0xb7, 0x89, 0x73, 0x0b, 0xaa, 0x5b, 0xcb, 0x42, 0x74, 0x38, 0xe8, 0x3b,
	0xf6, 0x70, 0xf0, 0x7d, 0xe0, 0x18, 0x4f, 0x56, 0x95, 0xd3, 0xbd, 0xec, 0x19, 0xda, 0x79, 0x0b,
	0x8c, 0xf2, 0x72, 0x90, 0x43, 0x78, 0xe6, 0x15, 0x72, 0x02, 0xc7, 0x9e, 0x3d, 0xb4, 0x1d, 0x6f,
	0xd8, 0xbc, 0x68, 0x0f, 0xbb, 0x5f, 0x3b, 0x86, 0x76, 0x7e, 0x03, 0xa7, 0x8f, 0xec, 0x33, 0x39,
	0x82, 0x4a, 0xcf, 0xf6, 0x7a, 0x2b, 0xe7, 0x0b, 0x38, 0x55, 0xa5, 0xd7, 0xb3, 0x33, 0xeb, 0xc7,
	0xf6, 0x68, 0x21, 0x31, 0x35, 0xb4, 0x32, 0xd1, 0x6a, 0xe6, 0xc4, 0x4e, 0xe7, 0x1b, 0xbc, 0x61,
	0x62, 0x62, 0xfe, 0x46, 0x1f, 0x7d, 0x6a, 0xb2, 0x48, 0xcc, 0x7c, 0x73, 0x96, 0x62, 0x92, 0x2d,
	0x60, 0xfe, 0x1a, 0x7f, 0xbd, 0x0d, 0xb8, 0x0c, 0x67, 0x23, 0x93, 0x89, 0x89, 0x15, 0x8d, 0x3f,
	0xa0, 0x1f, 0xa0, 0x85, 0x77, 0x68, 0xd1, 0x29, 0xb7, 0x02, 0x61, 0xe5, 0xaf, 0x78, 0xb4, 0xaf,
	0xc4, 0xad, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x7b, 0x35, 0x40, 0xd6, 0x03, 0x00, 0x00,
}
