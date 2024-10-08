// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetHistoryKL/Qot_GetHistoryKL.proto

package Qot_GetHistoryKL

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
	RehabType            *int32               `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`
	KlType               *int32               `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`
	Security             *Qot_Common.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`
	BeginTime            *string              `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string              `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`
	MaxAckKLNum          *int32               `protobuf:"varint,6,opt,name=maxAckKLNum" json:"maxAckKLNum,omitempty"`
	NeedKLFieldsFlag     *int64               `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476, []int{0}
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

func (m *C2S) GetRehabType() int32 {
	if m != nil && m.RehabType != nil {
		return *m.RehabType
	}
	return 0
}

func (m *C2S) GetKlType() int32 {
	if m != nil && m.KlType != nil {
		return *m.KlType
	}
	return 0
}

func (m *C2S) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *C2S) GetMaxAckKLNum() int32 {
	if m != nil && m.MaxAckKLNum != nil {
		return *m.MaxAckKLNum
	}
	return 0
}

func (m *C2S) GetNeedKLFieldsFlag() int64 {
	if m != nil && m.NeedKLFieldsFlag != nil {
		return *m.NeedKLFieldsFlag
	}
	return 0
}

type S2C struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	KlList               []*Qot_Common.KLine  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`
	NextKLTime           *string              `protobuf:"bytes,3,opt,name=nextKLTime" json:"nextKLTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476, []int{1}
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

func (m *S2C) GetKlList() []*Qot_Common.KLine {
	if m != nil {
		return m.KlList
	}
	return nil
}

func (m *S2C) GetNextKLTime() string {
	if m != nil && m.NextKLTime != nil {
		return *m.NextKLTime
	}
	return ""
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
	return fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476, []int{2}
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
	return fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetHistoryKL.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetHistoryKL.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetHistoryKL.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetHistoryKL.Response")
}

func init() {
	proto.RegisterFile("Qot_GetHistoryKL/Qot_GetHistoryKL.proto", fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476)
}

var fileDescriptor_Qot_GetHistoryKL_51b858e71f3ae476 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0xab, 0xda, 0x40,
	0x14, 0xc5, 0x99, 0xc4, 0x18, 0x1d, 0x37, 0x76, 0xfa, 0x87, 0xc1, 0x16, 0x19, 0xb2, 0x31, 0x2d,
	0x54, 0x65, 0xe8, 0xaa, 0xbb, 0x12, 0xb0, 0x85, 0xa4, 0x85, 0x4e, 0xdc, 0x17, 0x4d, 0x2e, 0x36,
	0x68, 0x32, 0x76, 0x66, 0x04, 0xdd, 0x16, 0xfa, 0x89, 0xfb, 0x05, 0x4a, 0xfe, 0xd8, 0xe4, 0xe9,
	0x5b, 0xbc, 0x55, 0x72, 0x7e, 0xe7, 0xc0, 0x3d, 0x73, 0x2f, 0x9e, 0x7d, 0x97, 0xe6, 0xc7, 0x67,
	0x30, 0x5f, 0x32, 0x6d, 0xa4, 0xba, 0x84, 0xd1, 0xe2, 0x16, 0xcc, 0x8f, 0x4a, 0x1a, 0x49, 0xc6,
	0xb7, 0x7c, 0xf2, 0x3c, 0x90, 0x79, 0x2e, 0x8b, 0x45, 0xfd, 0xa9, 0x63, 0x93, 0xd7, 0x65, 0xac,
	0x31, 0xda, 0xdf, 0xda, 0xf4, 0xfe, 0x22, 0x6c, 0x07, 0x3c, 0x26, 0x6f, 0xf0, 0x50, 0xc1, 0xcf,
	0xcd, 0x76, 0x7d, 0x39, 0x02, 0x45, 0xcc, 0xf2, 0x1d, 0xd1, 0x02, 0xf2, 0x0a, 0xf7, 0xf7, 0x87,
	0xca, 0xb2, 0x2a, 0xab, 0x51, 0x64, 0x89, 0x07, 0x1a, 0x92, 0x93, 0xca, 0xcc, 0x85, 0xda, 0xcc,
	0xf2, 0x47, 0xfc, 0xc5, 0xbc, 0x33, 0x22, 0x6e, 0x3c, 0xf1, 0x3f, 0x55, 0xce, 0xd9, 0xc2, 0x2e,
	0x2b, 0xd6, 0x59, 0x0e, 0xb4, 0xc7, 0x2c, 0x7f, 0x28, 0x5a, 0x40, 0x28, 0x76, 0xa1, 0x48, 0x2b,
	0xcf, 0xa9, 0xbc, 0xab, 0x24, 0x0c, 0x8f, 0xf2, 0xcd, 0xf9, 0x53, 0xb2, 0x0f, 0xa3, 0x6f, 0xa7,
	0x9c, 0xf6, 0x19, 0xf2, 0x1d, 0xd1, 0x45, 0xe4, 0x1d, 0x1e, 0x17, 0x00, 0x69, 0x18, 0xad, 0x32,
	0x38, 0xa4, 0x7a, 0x75, 0xd8, 0xec, 0xa8, 0xcb, 0x90, 0x6f, 0x8b, 0x3b, 0xee, 0xfd, 0x46, 0xd8,
	0x8e, 0x79, 0xf0, 0xa0, 0x3f, 0x7a, 0x52, 0xff, 0xb7, 0xe5, 0x26, 0xa2, 0x4c, 0x1b, 0x6a, 0x31,
	0xdb, 0x1f, 0xf1, 0x67, 0xdd, 0x7c, 0x18, 0x65, 0x05, 0x88, 0x26, 0x40, 0xa6, 0x18, 0x17, 0x70,
	0x36, 0x61, 0x54, 0xbd, 0xc7, 0x66, 0xc8, 0x1f, 0x8a, 0x0e, 0xf1, 0x38, 0x76, 0x05, 0xfc, 0x3a,
	0x81, 0x36, 0x64, 0x86, 0xed, 0x84, 0xeb, 0xa6, 0xc2, 0xcb, 0xf9, 0xdd, 0xbd, 0x03, 0x1e, 0x8b,
	0x32, 0xe1, 0xfd, 0x41, 0x78, 0x20, 0x40, 0x1f, 0x65, 0xa1, 0x81, 0x4c, 0xb1, 0xab, 0xc0, 0xb4,
	0x17, 0xfb, 0xd8, 0x7b, 0xff, 0x61, 0xb9, 0x14, 0x57, 0x58, 0x5e, 0x4d, 0x81, 0xf9, 0xaa, 0x77,
	0xd4, 0xaa, 0x86, 0x37, 0xaa, 0xda, 0xb2, 0x52, 0x81, 0x4c, 0xeb, 0x56, 0x8e, 0xb8, 0xca, 0xb2,
	0x87, 0xe6, 0x09, 0xed, 0x31, 0xf4, 0x78, 0x8f, 0x98, 0x07, 0xa2, 0x4c, 0xfc, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x33, 0x4a, 0x9c, 0xa4, 0x02, 0x00, 0x00,
}
