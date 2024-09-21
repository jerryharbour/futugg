// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetTicker/Qot_GetTicker.proto

package Qot_GetTicker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"
import Qot_Common "github.com/jerryharbour/futugg/pb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	MaxRetNum            *int32               `protobuf:"varint,2,req,name=maxRetNum" json:"maxRetNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetMaxRetNum() int32 {
	if m != nil && m.MaxRetNum != nil {
		return *m.MaxRetNum
	}
	return 0
}

type S2C struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	TickerList           []*Qot_Common.Ticker `protobuf:"bytes,2,rep,name=tickerList" json:"tickerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072, []int{1}
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

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetTickerList() []*Qot_Common.Ticker {
	if m != nil {
		return m.TickerList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
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
	return fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetTicker.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetTicker.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetTicker.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetTicker.Response")
}

func init() {
	proto.RegisterFile("Qot_GetTicker/Qot_GetTicker.proto", fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072)
}

var fileDescriptor_Qot_GetTicker_c83dc8a0bfb64072 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0xbb, 0xda, 0xed, 0x0d, 0x2f, 0x51, 0x24, 0x4c, 0x91, 0x58, 0x3c, 0xf4, 0xe2,
	0x5a, 0x82, 0x27, 0xaf, 0x39, 0x78, 0x51, 0xc1, 0xd7, 0x79, 0x16, 0xa9, 0x0f, 0x29, 0xa3, 0x4b,
	0x4d, 0x52, 0x70, 0x57, 0xff, 0x72, 0xe9, 0xda, 0xd9, 0x0e, 0x3c, 0x79, 0x6a, 0xbf, 0xef, 0xfb,
	0x7d, 0xc9, 0x7b, 0x81, 0xab, 0x67, 0xed, 0x5e, 0xef, 0xc9, 0xad, 0xca, 0x62, 0x4d, 0x26, 0x3d,
	0x50, 0xcb, 0xda, 0x68, 0xa7, 0xd9, 0xf1, 0x81, 0xb9, 0x38, 0x51, 0xba, 0xaa, 0xf4, 0x26, 0xed,
	0x3e, 0x1d, 0xb3, 0x38, 0x6f, 0x99, 0x3e, 0x18, 0x7e, 0xbb, 0x30, 0x7e, 0x81, 0x40, 0xc9, 0x9c,
	0x65, 0x30, 0xb5, 0x54, 0x34, 0xa6, 0x74, 0x5b, 0xee, 0x09, 0x3f, 0x99, 0xcb, 0xd3, 0xe5, 0x88,
	0xcd, 0xfb, 0x0c, 0x7f, 0x29, 0x76, 0x01, 0xb3, 0xea, 0xed, 0x0b, 0xc9, 0x3d, 0x35, 0x15, 0xf7,
	0x85, 0x9f, 0x84, 0x38, 0x18, 0xf1, 0x1a, 0x82, 0x5c, 0xaa, 0x7f, 0x1c, 0x2b, 0x01, 0xdc, 0x6e,
	0x97, 0x87, 0xd2, 0x3a, 0xee, 0x8b, 0x20, 0x99, 0x4b, 0x36, 0xee, 0x74, 0x9b, 0xe2, 0x88, 0x8a,
	0x53, 0x88, 0x90, 0x3e, 0x1b, 0xb2, 0x8e, 0x5d, 0x43, 0x50, 0x48, 0xdb, 0xdf, 0xd5, 0xf5, 0x86,
	0x27, 0x53, 0x32, 0xc7, 0x36, 0x8e, 0xbf, 0x3d, 0x98, 0x22, 0xd9, 0x5a, 0x6f, 0x2c, 0xb1, 0x4b,
	0x88, 0x0c, 0xb9, 0xd5, 0xb6, 0xa6, 0x5d, 0x2d, 0xbc, 0x9b, 0xdc, 0xdc, 0x66, 0x19, 0xee, 0x4d,
	0x76, 0x06, 0x47, 0x86, 0xdc, 0xa3, 0xfd, 0xe0, 0xbe, 0xf0, 0x92, 0x19, 0xf6, 0x8a, 0x71, 0x88,
	0xc8, 0x18, 0xa5, 0xdf, 0x89, 0x07, 0xc2, 0x4b, 0x42, 0xdc, 0xcb, 0x76, 0x08, 0x2b, 0x0b, 0x3e,
	0x11, 0xde, 0x1f, 0x43, 0xe4, 0x52, 0x61, 0x1b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xfe,
	0x8f, 0xe4, 0xde, 0x01, 0x00, 0x00,
}
