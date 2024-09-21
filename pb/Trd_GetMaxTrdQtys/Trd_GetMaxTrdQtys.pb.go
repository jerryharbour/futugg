// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_GetMaxTrdQtys/Trd_GetMaxTrdQtys.proto

/*
Package Trd_GetMaxTrdQtys is a generated protocol buffer package.

It is generated from these files:
	Trd_GetMaxTrdQtys/Trd_GetMaxTrdQtys.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Trd_GetMaxTrdQtys

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
	Header    *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderType *int32                `protobuf:"varint,2,req,name=orderType" json:"orderType,omitempty"`
	Code      *string               `protobuf:"bytes,3,req,name=code" json:"code,omitempty"`
	Price     *float64              `protobuf:"fixed64,4,req,name=price" json:"price,omitempty"`
	OrderID   *uint64               `protobuf:"varint,5,opt,name=orderID" json:"orderID,omitempty"`
	// 为保证与下单的价格同步，也提供调整价格选项，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice        *bool    `protobuf:"varint,6,opt,name=adjustPrice" json:"adjustPrice,omitempty"`
	AdjustSideAndLimit *float64 `protobuf:"fixed64,7,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetOrderType() int32 {
	if m != nil && m.OrderType != nil {
		return *m.OrderType
	}
	return 0
}

func (m *C2S) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *C2S) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *C2S) GetOrderID() uint64 {
	if m != nil && m.OrderID != nil {
		return *m.OrderID
	}
	return 0
}

func (m *C2S) GetAdjustPrice() bool {
	if m != nil && m.AdjustPrice != nil {
		return *m.AdjustPrice
	}
	return false
}

func (m *C2S) GetAdjustSideAndLimit() float64 {
	if m != nil && m.AdjustSideAndLimit != nil {
		return *m.AdjustSideAndLimit
	}
	return 0
}

type S2C struct {
	Header           *Trd_Common.TrdHeader  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	MaxTrdQtys       *Trd_Common.MaxTrdQtys `protobuf:"bytes,2,opt,name=maxTrdQtys" json:"maxTrdQtys,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetMaxTrdQtys() *Trd_Common.MaxTrdQtys {
	if m != nil {
		return m.MaxTrdQtys
	}
	return nil
}

type Request struct {
	C2S              *C2S   `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType          *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg           *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode          *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C              *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	proto.RegisterType((*C2S)(nil), "Trd_GetMaxTrdQtys.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_GetMaxTrdQtys.S2C")
	proto.RegisterType((*Request)(nil), "Trd_GetMaxTrdQtys.Request")
	proto.RegisterType((*Response)(nil), "Trd_GetMaxTrdQtys.Response")
}

func init() { proto.RegisterFile("Trd_GetMaxTrdQtys/Trd_GetMaxTrdQtys.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4b, 0x4b, 0xeb, 0x40,
	0x14, 0x66, 0xf2, 0x68, 0xda, 0xd3, 0xd5, 0x9d, 0x7b, 0x6f, 0x19, 0x7a, 0x2f, 0x32, 0x64, 0x35,
	0x2e, 0xfa, 0x60, 0x14, 0x17, 0xee, 0x24, 0x82, 0x0a, 0x16, 0x74, 0x92, 0xbd, 0x94, 0xcc, 0xa0,
	0x11, 0xd3, 0xc4, 0x99, 0x29, 0xd8, 0x3f, 0xe0, 0x6f, 0xf5, 0x67, 0x48, 0x1e, 0xb5, 0x29, 0x75,
	0xe3, 0x2a, 0xf3, 0xbd, 0xc8, 0x39, 0xdf, 0x81, 0xe3, 0x44, 0xcb, 0x87, 0x2b, 0x65, 0x17, 0xcb,
	0xb7, 0x44, 0xcb, 0x7b, 0xbb, 0x31, 0xb3, 0x03, 0x66, 0x5a, 0xea, 0xc2, 0x16, 0xf8, 0xd7, 0x81,
	0x30, 0xfe, 0x1d, 0x15, 0x79, 0x5e, 0xac, 0x66, 0xcd, 0xa7, 0xf1, 0x8d, 0xff, 0x55, 0xbe, 0x56,
	0xd8, 0x3d, 0x1b, 0x31, 0xfc, 0x40, 0xe0, 0x46, 0x3c, 0xc6, 0x13, 0xe8, 0x3d, 0xa9, 0xa5, 0x54,
	0x9a, 0x20, 0xea, 0xb0, 0x21, 0xff, 0x3b, 0xed, 0x58, 0x13, 0x2d, 0xaf, 0x6b, 0x51, 0xb4, 0x26,
	0xfc, 0x1f, 0x06, 0x85, 0x96, 0x4a, 0x27, 0x9b, 0x52, 0x11, 0x87, 0x3a, 0xcc, 0x17, 0x3b, 0x02,
	0x63, 0xf0, 0xd2, 0x42, 0x2a, 0xe2, 0x52, 0x87, 0x0d, 0x44, 0xfd, 0xc6, 0x7f, 0xc0, 0x2f, 0x75,
	0x96, 0x2a, 0xe2, 0x51, 0x87, 0x21, 0xd1, 0x00, 0x4c, 0x20, 0xa8, 0x63, 0x37, 0x97, 0xc4, 0xa7,
	0x88, 0x79, 0x62, 0x0b, 0x31, 0x85, 0xe1, 0x52, 0x3e, 0xaf, 0x8d, 0xbd, 0xab, 0x53, 0x3d, 0x8a,
	0x58, 0x5f, 0x74, 0x29, 0x3c, 0x05, 0xdc, 0xc0, 0x38, 0x93, 0xea, 0x62, 0x25, 0x6f, 0xb3, 0x3c,
	0xb3, 0x24, 0xa0, 0x88, 0x21, 0xf1, 0x8d, 0x12, 0xbe, 0x80, 0x1b, 0xf3, 0xe8, 0xa7, 0x9b, 0x9e,
	0x01, 0xe4, 0x5f, 0x05, 0x13, 0x87, 0x22, 0x36, 0xe4, 0xa3, 0x6e, 0x64, 0x57, 0xbf, 0xe8, 0x38,
	0xc3, 0x13, 0x08, 0x84, 0x7a, 0x5d, 0x2b, 0x63, 0x31, 0x03, 0x37, 0xe5, 0xa6, 0xfd, 0x5d, 0x93,
	0xdd, 0xbf, 0x67, 0xc4, 0x63, 0x51, 0x59, 0xc2, 0x77, 0x04, 0x7d, 0xa1, 0x4c, 0x59, 0xac, 0x8c,
	0xc2, 0x47, 0x10, 0x68, 0x65, 0xeb, 0x86, 0xab, 0xa8, 0x7f, 0xee, 0x4d, 0x4e, 0xe7, 0x73, 0xb1,
	0x25, 0xf1, 0x08, 0x7a, 0x5a, 0xd9, 0x85, 0x79, 0xac, 0xa7, 0x1a, 0x88, 0x16, 0x55, 0x9d, 0x2a,
	0xad, 0xa3, 0xe6, 0x00, 0x88, 0xf9, 0x62, 0x0b, 0xab, 0x41, 0x0c, 0x4f, 0x89, 0xd7, 0x59, 0x62,
	0x7f, 0x90, 0x98, 0x47, 0xa2, 0xb2, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xa8, 0x4c, 0x21,
	0x87, 0x02, 0x00, 0x00,
}
