// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Notify/Notify.proto

package Notify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NotifyType int32

const (
	NotifyType_NotifyType_None     NotifyType = 0
	NotifyType_NotifyType_GtwEvent NotifyType = 1
)

var NotifyType_name = map[int32]string{
	0: "NotifyType_None",
	1: "NotifyType_GtwEvent",
}
var NotifyType_value = map[string]int32{
	"NotifyType_None":     0,
	"NotifyType_GtwEvent": 1,
}

func (x NotifyType) Enum() *NotifyType {
	p := new(NotifyType)
	*p = x
	return p
}
func (x NotifyType) String() string {
	return proto.EnumName(NotifyType_name, int32(x))
}
func (x *NotifyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NotifyType_value, data, "NotifyType")
	if err != nil {
		return err
	}
	*x = NotifyType(value)
	return nil
}
func (NotifyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Notify_a623816a5806fe93, []int{0}
}

type GtwEventType int32

const (
	GtwEventType_GtwEventType_None                GtwEventType = 0
	GtwEventType_GtwEventType_LocalCfgLoadFailed  GtwEventType = 1
	GtwEventType_GtwEventType_APISvrRunFailed     GtwEventType = 2
	GtwEventType_GtwEventType_ForceUpdate         GtwEventType = 3
	GtwEventType_GtwEventType_LoginFailed         GtwEventType = 4
	GtwEventType_GtwEventType_UnAgreeDisclaimer   GtwEventType = 5
	GtwEventType_GtwEventType_NetCfgMissing       GtwEventType = 6
	GtwEventType_GtwEventType_KickedOut           GtwEventType = 7
	GtwEventType_GtwEventType_LoginPwdChanged     GtwEventType = 8
	GtwEventType_GtwEventType_BanLogin            GtwEventType = 9
	GtwEventType_GtwEventType_NeedPicVerifyCode   GtwEventType = 10
	GtwEventType_GtwEventType_NeedPhoneVerifyCode GtwEventType = 11
	GtwEventType_GtwEventType_AppDataNotExist     GtwEventType = 12
	GtwEventType_GtwEventType_NessaryDataMissing  GtwEventType = 13
	GtwEventType_GtwEventType_TradePwdChanged     GtwEventType = 14
	GtwEventType_GtwEventType_EnableDeviceLock    GtwEventType = 15
)

var GtwEventType_name = map[int32]string{
	0:  "GtwEventType_None",
	1:  "GtwEventType_LocalCfgLoadFailed",
	2:  "GtwEventType_APISvrRunFailed",
	3:  "GtwEventType_ForceUpdate",
	4:  "GtwEventType_LoginFailed",
	5:  "GtwEventType_UnAgreeDisclaimer",
	6:  "GtwEventType_NetCfgMissing",
	7:  "GtwEventType_KickedOut",
	8:  "GtwEventType_LoginPwdChanged",
	9:  "GtwEventType_BanLogin",
	10: "GtwEventType_NeedPicVerifyCode",
	11: "GtwEventType_NeedPhoneVerifyCode",
	12: "GtwEventType_AppDataNotExist",
	13: "GtwEventType_NessaryDataMissing",
	14: "GtwEventType_TradePwdChanged",
	15: "GtwEventType_EnableDeviceLock",
}
var GtwEventType_value = map[string]int32{
	"GtwEventType_None":                0,
	"GtwEventType_LocalCfgLoadFailed":  1,
	"GtwEventType_APISvrRunFailed":     2,
	"GtwEventType_ForceUpdate":         3,
	"GtwEventType_LoginFailed":         4,
	"GtwEventType_UnAgreeDisclaimer":   5,
	"GtwEventType_NetCfgMissing":       6,
	"GtwEventType_KickedOut":           7,
	"GtwEventType_LoginPwdChanged":     8,
	"GtwEventType_BanLogin":            9,
	"GtwEventType_NeedPicVerifyCode":   10,
	"GtwEventType_NeedPhoneVerifyCode": 11,
	"GtwEventType_AppDataNotExist":     12,
	"GtwEventType_NessaryDataMissing":  13,
	"GtwEventType_TradePwdChanged":     14,
	"GtwEventType_EnableDeviceLock":    15,
}

func (x GtwEventType) Enum() *GtwEventType {
	p := new(GtwEventType)
	*p = x
	return p
}
func (x GtwEventType) String() string {
	return proto.EnumName(GtwEventType_name, int32(x))
}
func (x *GtwEventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GtwEventType_value, data, "GtwEventType")
	if err != nil {
		return err
	}
	*x = GtwEventType(value)
	return nil
}
func (GtwEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Notify_a623816a5806fe93, []int{1}
}

type GtwEvent struct {
	EventType            *int32   `protobuf:"varint,1,req,name=eventType" json:"eventType,omitempty"`
	Desc                 *string  `protobuf:"bytes,2,req,name=desc" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GtwEvent) Reset()         { *m = GtwEvent{} }
func (m *GtwEvent) String() string { return proto.CompactTextString(m) }
func (*GtwEvent) ProtoMessage()    {}
func (*GtwEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_Notify_a623816a5806fe93, []int{0}
}
func (m *GtwEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GtwEvent.Unmarshal(m, b)
}
func (m *GtwEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GtwEvent.Marshal(b, m, deterministic)
}
func (dst *GtwEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GtwEvent.Merge(dst, src)
}
func (m *GtwEvent) XXX_Size() int {
	return xxx_messageInfo_GtwEvent.Size(m)
}
func (m *GtwEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GtwEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GtwEvent proto.InternalMessageInfo

func (m *GtwEvent) GetEventType() int32 {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return 0
}

func (m *GtwEvent) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type S2C struct {
	Type                 *int32    `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Event                *GtwEvent `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Notify_a623816a5806fe93, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *S2C) GetEvent() *GtwEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Notify_a623816a5806fe93, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*GtwEvent)(nil), "Notify.GtwEvent")
	proto.RegisterType((*S2C)(nil), "Notify.S2C")
	proto.RegisterType((*Response)(nil), "Notify.Response")
	proto.RegisterEnum("Notify.NotifyType", NotifyType_name, NotifyType_value)
	proto.RegisterEnum("Notify.GtwEventType", GtwEventType_name, GtwEventType_value)
}

func init() { proto.RegisterFile("Notify/Notify.proto", fileDescriptor_Notify_a623816a5806fe93) }

var fileDescriptor_Notify_a623816a5806fe93 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x5f, 0x6f, 0x12, 0x41,
	0x10, 0x97, 0x7f, 0x05, 0x86, 0x6a, 0xd7, 0x25, 0xad, 0x27, 0xa1, 0x78, 0xa2, 0x31, 0xa4, 0x89,
	0x6d, 0x43, 0x7c, 0x6a, 0x7c, 0xc1, 0x83, 0x1a, 0x23, 0x45, 0x72, 0xb4, 0xbe, 0x9a, 0xf5, 0x6e,
	0xb8, 0x6e, 0x0a, 0xbb, 0x97, 0xdd, 0x2d, 0x95, 0xf8, 0x51, 0xfc, 0xb2, 0x86, 0x3b, 0x36, 0x70,
	0xb9, 0xa7, 0x9d, 0xdf, 0x9f, 0x99, 0xf9, 0xed, 0x2e, 0x34, 0x27, 0xd2, 0xf0, 0xf9, 0xfa, 0x22,
	0x3d, 0xce, 0x63, 0x25, 0x8d, 0xa4, 0x07, 0x29, 0x6a, 0x35, 0x3d, 0xb9, 0x5c, 0x4a, 0x71, 0x91,
	0x1e, 0xa9, 0xd8, 0xfd, 0x0c, 0xb5, 0xaf, 0xe6, 0x69, 0xb4, 0x42, 0x61, 0x68, 0x1b, 0xea, 0xb8,
	0x29, 0x6e, 0xd7, 0x31, 0x3a, 0x05, 0xb7, 0xd8, 0xab, 0xf8, 0x3b, 0x82, 0x52, 0x28, 0x87, 0xa8,
	0x03, 0xa7, 0xe8, 0x16, 0x7b, 0x75, 0x3f, 0xa9, 0xbb, 0x03, 0x28, 0xcd, 0xfa, 0xde, 0x46, 0x32,
	0xbb, 0x9e, 0xa4, 0xa6, 0x1f, 0xa0, 0x92, 0xf4, 0x3a, 0x45, 0xb7, 0xd0, 0x6b, 0xf4, 0xc9, 0xf9,
	0x36, 0x93, 0xdd, 0xe6, 0xa7, 0x72, 0xf7, 0x2f, 0xd4, 0x7c, 0xd4, 0xb1, 0x14, 0x1a, 0x69, 0x07,
	0xaa, 0x0a, 0xf7, 0xd6, 0x5f, 0x95, 0x3f, 0x7e, 0xba, 0xbc, 0xf4, 0x2d, 0x49, 0x4f, 0xe0, 0x40,
	0xa1, 0xb9, 0xd1, 0x51, 0x32, 0xb4, 0xee, 0x6f, 0x11, 0x75, 0xa0, 0x8a, 0x4a, 0x79, 0x32, 0x44,
	0xa7, 0xe4, 0x16, 0x7a, 0x15, 0xdf, 0x42, 0x7a, 0x0a, 0x25, 0xdd, 0x0f, 0x9c, 0x72, 0x92, 0xa1,
	0x61, 0x33, 0xcc, 0xfa, 0x9e, 0xbf, 0xe1, 0xcf, 0xae, 0x00, 0x52, 0x2a, 0x19, 0xdf, 0x84, 0xa3,
	0x1d, 0xfa, 0x35, 0x91, 0x02, 0xc9, 0x33, 0xfa, 0xca, 0x3e, 0x6a, 0x42, 0xda, 0xf4, 0xa4, 0x70,
	0xf6, 0xaf, 0x0c, 0x87, 0x16, 0x26, 0xed, 0xc7, 0xf0, 0x72, 0x1f, 0xdb, 0x01, 0xef, 0xe0, 0x4d,
	0x86, 0x1e, 0xcb, 0x80, 0x2d, 0xbc, 0x79, 0x34, 0x96, 0x2c, 0xbc, 0x66, 0x7c, 0x81, 0x21, 0x29,
	0x50, 0x17, 0xda, 0x19, 0xd3, 0x60, 0xfa, 0x6d, 0xb6, 0x52, 0xfe, 0xa3, 0xd8, 0x3a, 0x8a, 0xb4,
	0x0d, 0x4e, 0xc6, 0x71, 0x2d, 0x55, 0x80, 0x77, 0x71, 0xc8, 0x0c, 0x92, 0x52, 0x4e, 0x1d, 0xcb,
	0x88, 0xdb, 0xde, 0x32, 0xed, 0x42, 0x27, 0xa3, 0xde, 0x89, 0x41, 0xa4, 0x10, 0x87, 0x5c, 0x07,
	0x0b, 0xc6, 0x97, 0xa8, 0x48, 0x85, 0x76, 0xa0, 0x95, 0x4d, 0x8f, 0xc6, 0x9b, 0x47, 0x37, 0x5c,
	0x6b, 0x2e, 0x22, 0x72, 0x40, 0x5b, 0x70, 0x92, 0xd1, 0xbf, 0xf3, 0xe0, 0x01, 0xc3, 0x1f, 0x8f,
	0x86, 0x54, 0x73, 0xe9, 0x93, 0xed, 0xd3, 0xa7, 0xd0, 0xbb, 0x67, 0x22, 0xc2, 0x90, 0xd4, 0xe8,
	0x6b, 0x38, 0xce, 0x38, 0xbe, 0x30, 0x91, 0x98, 0x48, 0x3d, 0x17, 0x6e, 0x82, 0x18, 0x4e, 0x79,
	0xf0, 0x13, 0x15, 0x9f, 0xaf, 0x37, 0x9f, 0x48, 0x80, 0xbe, 0x07, 0x37, 0xef, 0xb9, 0x97, 0x02,
	0xf7, 0x5c, 0x8d, 0xfc, 0x23, 0xc6, 0xf1, 0x90, 0x19, 0x36, 0x91, 0x66, 0xf4, 0x87, 0x6b, 0x43,
	0x0e, 0x73, 0x7f, 0x31, 0x41, 0xad, 0x99, 0x5a, 0x6f, 0x5c, 0xf6, 0xa6, 0xcf, 0x73, 0x63, 0x6e,
	0x15, 0x0b, 0x71, 0xef, 0x36, 0x2f, 0xe8, 0x5b, 0x38, 0xcd, 0x38, 0x46, 0x82, 0xfd, 0x5e, 0xe0,
	0x10, 0x57, 0x3c, 0xc0, 0xb1, 0x0c, 0x1e, 0xc8, 0xd1, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0xa1, 0xeb, 0x3c, 0x8a, 0x03, 0x00, 0x00,
}
