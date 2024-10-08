// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_UpdateBroker/Qot_UpdateBroker.proto

package Qot_UpdateBroker

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

type S2C struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	BrokerAskList        []*Qot_Common.Broker `protobuf:"bytes,2,rep,name=brokerAskList" json:"brokerAskList,omitempty"`
	BrokerBidList        []*Qot_Common.Broker `protobuf:"bytes,3,rep,name=brokerBidList" json:"brokerBidList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_UpdateBroker_271461a08044f2e5, []int{0}
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

func (m *S2C) GetBrokerAskList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerAskList
	}
	return nil
}

func (m *S2C) GetBrokerBidList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerBidList
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
	return fileDescriptor_Qot_UpdateBroker_271461a08044f2e5, []int{1}
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
	proto.RegisterType((*S2C)(nil), "Qot_UpdateBroker.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateBroker.Response")
}

func init() {
	proto.RegisterFile("Qot_UpdateBroker/Qot_UpdateBroker.proto", fileDescriptor_Qot_UpdateBroker_271461a08044f2e5)
}

var fileDescriptor_Qot_UpdateBroker_271461a08044f2e5 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x0f, 0xcc, 0x2f, 0x89,
	0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x75, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x47, 0x17, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x97, 0x12, 0x76, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x87, 0x50, 0x10, 0x65, 0x52, 0xd2, 0x20, 0x65, 0x50, 0x09, 0x04, 0x13, 0x22, 0xa9,
	0xb4, 0x9a, 0x91, 0x8b, 0x39, 0xd8, 0xc8, 0x59, 0xc8, 0x80, 0x8b, 0xa3, 0x38, 0x35, 0xb9, 0xb4,
	0x28, 0xb3, 0xa4, 0x52, 0x82, 0x51, 0x81, 0x49, 0x83, 0xdb, 0x48, 0x44, 0x0f, 0x49, 0x71, 0x30,
	0x54, 0x2e, 0x08, 0xae, 0x4a, 0xc8, 0x82, 0x8b, 0x37, 0x09, 0x6c, 0xab, 0x63, 0x71, 0xb6, 0x4f,
	0x66, 0x71, 0x89, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x10, 0xb2, 0x36, 0x88, 0xb3, 0x82,
	0x50, 0x15, 0x22, 0x74, 0x3a, 0x65, 0xa6, 0x80, 0x75, 0x32, 0x13, 0xd2, 0x09, 0x55, 0xa8, 0xd4,
	0xca, 0xc8, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc7, 0xc5, 0x5e,
	0x94, 0x5a, 0x12, 0x52, 0x59, 0x90, 0x0a, 0x76, 0x31, 0xab, 0x15, 0x8b, 0xae, 0x89, 0x81, 0x41,
	0x10, 0x4c, 0x50, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5, 0xc4, 0xb7, 0x38, 0x5d, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x92, 0xe0, 0x62, 0x4f, 0x2d, 0x2a, 0x72, 0xce, 0x4f, 0x49,
	0x95, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x85, 0xd4, 0xb9, 0x98, 0x8b, 0x8d, 0x92,
	0x25, 0x58, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0x30, 0x82, 0x3d, 0xd8, 0xc8, 0x39, 0x08,
	0xa4, 0x02, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x96, 0x89, 0xd1, 0xa3, 0x01, 0x00, 0x00,
}
