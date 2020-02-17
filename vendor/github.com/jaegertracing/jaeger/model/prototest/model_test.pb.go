// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model_test.proto

package prototest

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

type SpanRefType int32

const (
	SpanRefType_CHILD_OF     SpanRefType = 0
	SpanRefType_FOLLOWS_FROM SpanRefType = 1
)

var SpanRefType_name = map[int32]string{
	0: "CHILD_OF",
	1: "FOLLOWS_FROM",
}

var SpanRefType_value = map[string]int32{
	"CHILD_OF":     0,
	"FOLLOWS_FROM": 1,
}

func (x SpanRefType) String() string {
	return proto.EnumName(SpanRefType_name, int32(x))
}

func (SpanRefType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d019f60590a05da, []int{0}
}

type SpanRef struct {
	TraceId              []byte      `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId               []byte      `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	RefType              SpanRefType `protobuf:"varint,3,opt,name=ref_type,json=refType,proto3,enum=prototest.SpanRefType" json:"ref_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SpanRef) Reset()         { *m = SpanRef{} }
func (m *SpanRef) String() string { return proto.CompactTextString(m) }
func (*SpanRef) ProtoMessage()    {}
func (*SpanRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d019f60590a05da, []int{0}
}

func (m *SpanRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanRef.Unmarshal(m, b)
}
func (m *SpanRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanRef.Marshal(b, m, deterministic)
}
func (m *SpanRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanRef.Merge(m, src)
}
func (m *SpanRef) XXX_Size() int {
	return xxx_messageInfo_SpanRef.Size(m)
}
func (m *SpanRef) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanRef.DiscardUnknown(m)
}

var xxx_messageInfo_SpanRef proto.InternalMessageInfo

func (m *SpanRef) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *SpanRef) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *SpanRef) GetRefType() SpanRefType {
	if m != nil {
		return m.RefType
	}
	return SpanRefType_CHILD_OF
}

func init() {
	proto.RegisterEnum("prototest.SpanRefType", SpanRefType_name, SpanRefType_value)
	proto.RegisterType((*SpanRef)(nil), "prototest.SpanRef")
}

func init() { proto.RegisterFile("model_test.proto", fileDescriptor_7d019f60590a05da) }

var fileDescriptor_7d019f60590a05da = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcd, 0x4f, 0x49,
	0xcd, 0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53,
	0x20, 0x01, 0xa5, 0x02, 0x2e, 0xf6, 0xe0, 0x82, 0xc4, 0xbc, 0xa0, 0xd4, 0x34, 0x21, 0x49, 0x2e,
	0x8e, 0x92, 0xa2, 0xc4, 0xe4, 0xd4, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x76, 0x30, 0xdf, 0x33, 0x45, 0x48, 0x9c, 0x8b, 0xbd, 0xb8, 0x20, 0x31, 0x0f, 0x24, 0xc3, 0x04,
	0x96, 0x61, 0x03, 0x71, 0x3d, 0x53, 0x84, 0x0c, 0xb9, 0x38, 0x8a, 0x52, 0xd3, 0xe2, 0x4b, 0x2a,
	0x0b, 0x52, 0x25, 0x98, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0xc4, 0xf4, 0xe0, 0x86, 0xeb, 0x41, 0x4d,
	0x0e, 0xa9, 0x2c, 0x48, 0x0d, 0x62, 0x2f, 0x82, 0x30, 0xb4, 0x74, 0xb9, 0xb8, 0x91, 0xc4, 0x85,
	0x78, 0xb8, 0x38, 0x9c, 0x3d, 0x3c, 0x7d, 0x5c, 0xe2, 0xfd, 0xdd, 0x04, 0x18, 0x84, 0x04, 0xb8,
	0x78, 0xdc, 0xfc, 0x7d, 0x7c, 0xfc, 0xc3, 0x83, 0xe3, 0xdd, 0x82, 0xfc, 0x7d, 0x05, 0x18, 0x93,
	0xd8, 0xc0, 0xc6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x29, 0x73, 0x43, 0x66, 0xc6, 0x00,
	0x00, 0x00,
}
