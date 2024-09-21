// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetReference/Qot_GetReference.proto

/*
Package Qot_GetReference is a generated protocol buffer package.

It is generated from these files:
	Qot_GetReference/Qot_GetReference.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_GetReference

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

type ReferenceType int32

const (
	ReferenceType_ReferenceType_Unknow  ReferenceType = 0
	ReferenceType_ReferenceType_Warrant ReferenceType = 1
)

var ReferenceType_name = map[int32]string{
	0: "ReferenceType_Unknow",
	1: "ReferenceType_Warrant",
}
var ReferenceType_value = map[string]int32{
	"ReferenceType_Unknow":  0,
	"ReferenceType_Warrant": 1,
}

func (x ReferenceType) Enum() *ReferenceType {
	p := new(ReferenceType)
	*p = x
	return p
}
func (x ReferenceType) String() string {
	return proto.EnumName(ReferenceType_name, int32(x))
}
func (x *ReferenceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReferenceType_value, data, "ReferenceType")
	if err != nil {
		return err
	}
	*x = ReferenceType(value)
	return nil
}
func (ReferenceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type C2S struct {
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	ReferenceType    *int32               `protobuf:"varint,2,req,name=referenceType" json:"referenceType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetReferenceType() int32 {
	if m != nil && m.ReferenceType != nil {
		return *m.ReferenceType
	}
	return 0
}

type S2C struct {
	StaticInfoList   []*Qot_Common.SecurityStaticInfo `protobuf:"bytes,2,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetStaticInfoList() []*Qot_Common.SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
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
	proto.RegisterType((*C2S)(nil), "Qot_GetReference.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetReference.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetReference.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetReference.Response")
	proto.RegisterEnum("Qot_GetReference.ReferenceType", ReferenceType_name, ReferenceType_value)
}

func init() { proto.RegisterFile("Qot_GetReference/Qot_GetReference.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0x3a, 0x41,
	0x14, 0xc5, 0xff, 0xbb, 0xab, 0x7f, 0xed, 0x8a, 0x21, 0x93, 0xc6, 0x64, 0x20, 0xcb, 0x12, 0xb8,
	0x04, 0xa9, 0x0c, 0x3d, 0xf5, 0xba, 0x51, 0x04, 0xf9, 0xd0, 0x9d, 0xa2, 0xa7, 0x10, 0xd9, 0xae,
	0x21, 0xe1, 0x8c, 0xcd, 0x8c, 0x84, 0x1f, 0xa0, 0xef, 0x1d, 0x9b, 0xab, 0xe9, 0xe6, 0xd3, 0xcc,
	0x3d, 0xe7, 0x77, 0xb9, 0x87, 0x03, 0xdd, 0x07, 0xed, 0x46, 0xb7, 0xe4, 0x90, 0x26, 0x64, 0x48,
	0xa5, 0xd4, 0x2f, 0x0a, 0xbd, 0xb9, 0xd1, 0x4e, 0xb3, 0x46, 0x51, 0x6f, 0x1f, 0x25, 0x7a, 0x36,
	0xd3, 0xaa, 0xbf, 0x7a, 0x56, 0x58, 0xfb, 0x34, 0xc3, 0x72, 0xe3, 0xf7, 0xbb, 0x32, 0xa3, 0x17,
	0x08, 0x12, 0x21, 0xd9, 0x00, 0xaa, 0x96, 0xd2, 0x85, 0x99, 0xba, 0x25, 0xf7, 0x42, 0x3f, 0xae,
	0x89, 0x66, 0x6f, 0x8b, 0x95, 0xb9, 0x87, 0x1b, 0x8a, 0x9d, 0x41, 0xdd, 0xac, 0xef, 0x3e, 0x2e,
	0xe7, 0xc4, 0xfd, 0xd0, 0x8f, 0xcb, 0xb8, 0x2b, 0x46, 0x43, 0x08, 0xa4, 0x48, 0xd8, 0x0d, 0x1c,
	0x5a, 0x37, 0x76, 0xd3, 0xf4, 0x4e, 0x4d, 0xf4, 0xfd, 0xd4, 0x3a, 0xee, 0x87, 0x41, 0x5c, 0x13,
	0x9d, 0x7d, 0x47, 0xe4, 0x86, 0xc4, 0xc2, 0x56, 0x24, 0xa0, 0x82, 0xf4, 0xb1, 0x20, 0xeb, 0x58,
	0x17, 0x82, 0x54, 0xd8, 0x3c, 0x6c, 0xab, 0xf7, 0xa7, 0xa2, 0x44, 0x48, 0xcc, 0x88, 0xe8, 0xcb,
	0x83, 0x2a, 0x92, 0x9d, 0x6b, 0x65, 0x89, 0x75, 0xa0, 0x62, 0xc8, 0xfd, 0xe4, 0xcd, 0x36, 0xcb,
	0x57, 0xa5, 0x8b, 0xcb, 0xc1, 0x00, 0xd7, 0x22, 0x3b, 0x86, 0xff, 0x86, 0xdc, 0xd0, 0xbe, 0x71,
	0x3f, 0xf4, 0xe2, 0x03, 0xcc, 0x27, 0xc6, 0xa1, 0x42, 0xc6, 0x24, 0xfa, 0x95, 0x78, 0x10, 0x7a,
	0x71, 0x19, 0xd7, 0x63, 0x96, 0xc3, 0x8a, 0x94, 0x97, 0x42, 0x6f, 0x7f, 0x0e, 0x29, 0x12, 0xcc,
	0x88, 0xf3, 0x6b, 0xa8, 0xe3, 0x76, 0x37, 0x8c, 0x43, 0x73, 0x47, 0x18, 0x3d, 0xa9, 0x77, 0xa5,
	0x3f, 0x1b, 0xff, 0xd8, 0x09, 0xb4, 0x76, 0x9d, 0xe7, 0xb1, 0x31, 0x63, 0xe5, 0x1a, 0xde, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x5e, 0x8b, 0xe9, 0x1d, 0x02, 0x00, 0x00,
}
