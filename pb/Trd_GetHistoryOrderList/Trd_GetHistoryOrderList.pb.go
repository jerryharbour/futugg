// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_GetHistoryOrderList/Trd_GetHistoryOrderList.proto

package Trd_GetHistoryOrderList

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"
import Trd_Common "github.com/jerryharbour/futugg/pb/Trd_Common"

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
	Header               *Trd_Common.TrdHeader           `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	FilterConditions     *Trd_Common.TrdFilterConditions `protobuf:"bytes,2,req,name=filterConditions" json:"filterConditions,omitempty"`
	FilterStatusList     []int32                         `protobuf:"varint,3,rep,name=filterStatusList" json:"filterStatusList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356, []int{0}
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

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetFilterConditions() *Trd_Common.TrdFilterConditions {
	if m != nil {
		return m.FilterConditions
	}
	return nil
}

func (m *C2S) GetFilterStatusList() []int32 {
	if m != nil {
		return m.FilterStatusList
	}
	return nil
}

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderList            []*Trd_Common.Order   `protobuf:"bytes,2,rep,name=orderList" json:"orderList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356, []int{1}
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

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrderList() []*Trd_Common.Order {
	if m != nil {
		return m.OrderList
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
	return fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356, []int{2}
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
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
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
	return fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Trd_GetHistoryOrderList.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_GetHistoryOrderList.S2C")
	proto.RegisterType((*Request)(nil), "Trd_GetHistoryOrderList.Request")
	proto.RegisterType((*Response)(nil), "Trd_GetHistoryOrderList.Response")
}

func init() {
	proto.RegisterFile("Trd_GetHistoryOrderList/Trd_GetHistoryOrderList.proto", fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356)
}

var fileDescriptor_Trd_GetHistoryOrderList_12491c5130306356 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xb6, 0x7f, 0xde, 0x4e, 0x2f, 0xaf, 0x2b, 0xea, 0x52, 0x45, 0x43, 0x4e, 0x41,
	0x68, 0x5a, 0x16, 0x3d, 0xe8, 0x75, 0x41, 0x0b, 0x2a, 0xc2, 0xa6, 0x77, 0x29, 0xcd, 0xa8, 0x01,
	0x9b, 0xad, 0xb3, 0xdb, 0x43, 0xbf, 0x83, 0x5f, 0xc5, 0xef, 0x28, 0x49, 0x13, 0x5a, 0x5a, 0x2a,
	0x78, 0x4a, 0x66, 0x9e, 0xdf, 0x33, 0xfb, 0xec, 0x2c, 0x5c, 0x8f, 0x29, 0x7d, 0xb9, 0x47, 0x37,
	0xca, 0xac, 0x33, 0xb4, 0x7c, 0xa6, 0x14, 0xe9, 0x31, 0xb3, 0x6e, 0xb0, 0xa7, 0x1f, 0xcf, 0xc9,
	0x38, 0xc3, 0x4f, 0xf6, 0xc8, 0xbd, 0x43, 0x65, 0x66, 0x33, 0x93, 0x0f, 0x56, 0x9f, 0x15, 0xdd,
	0x3b, 0x2d, 0xe8, 0x4a, 0x58, 0xff, 0xae, 0xc4, 0xf0, 0xdb, 0x03, 0xa6, 0x64, 0xc2, 0xfb, 0xd0,
	0x7a, 0xc7, 0x49, 0x8a, 0x24, 0xbc, 0xc0, 0x8f, 0xba, 0xf2, 0x28, 0xde, 0x40, 0xc7, 0x94, 0x8e,
	0x4a, 0x51, 0x57, 0x10, 0x7f, 0x80, 0xff, 0xaf, 0xd9, 0x87, 0x43, 0x52, 0x26, 0x4f, 0x33, 0x97,
	0x99, 0xdc, 0x0a, 0xbf, 0x34, 0x5e, 0x6c, 0x19, 0xef, 0xb6, 0x30, 0xbd, 0x63, 0xe4, 0x97, 0xf5,
	0xb0, 0xc4, 0x4d, 0xdc, 0xc2, 0x16, 0x37, 0x11, 0x2c, 0x60, 0x51, 0x53, 0xef, 0xf4, 0x43, 0x04,
	0x96, 0x48, 0xf5, 0xd7, 0xb8, 0x03, 0xe8, 0x98, 0x7a, 0x49, 0xc2, 0x0f, 0x58, 0xd4, 0x95, 0x07,
	0x9b, 0x8e, 0x72, 0x83, 0x7a, 0xcd, 0x84, 0x37, 0xd0, 0xd6, 0xf8, 0xb9, 0x40, 0xeb, 0x78, 0x0c,
	0x6c, 0x2a, 0x6d, 0x75, 0xce, 0x59, 0xbc, 0xef, 0x65, 0x94, 0x4c, 0x74, 0x01, 0x86, 0x5f, 0x1e,
	0xfc, 0xd3, 0x68, 0xe7, 0x26, 0xb7, 0xc8, 0xcf, 0xa1, 0x4d, 0xe8, 0xc6, 0xcb, 0x39, 0x96, 0x03,
	0x9a, 0xb7, 0x8d, 0xfe, 0xd5, 0x70, 0xa8, 0xeb, 0x26, 0x3f, 0x86, 0x16, 0xa1, 0x7b, 0xb2, 0x6f,
	0xc2, 0x0f, 0xbc, 0xa8, 0xa3, 0xab, 0x8a, 0x0b, 0x68, 0x23, 0x91, 0x32, 0x29, 0x0a, 0x16, 0x78,
	0x51, 0x53, 0xd7, 0x65, 0x11, 0xc7, 0xca, 0xa9, 0x68, 0x04, 0xde, 0xaf, 0x71, 0x12, 0xa9, 0x74,
	0x01, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x0f, 0x5d, 0x1d, 0x63, 0x02, 0x00, 0x00,
}
