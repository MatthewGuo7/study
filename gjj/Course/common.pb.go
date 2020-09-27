// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

package Course

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TimeStamp struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeStamp) Reset()         { *m = TimeStamp{} }
func (m *TimeStamp) String() string { return proto.CompactTextString(m) }
func (*TimeStamp) ProtoMessage()    {}
func (*TimeStamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}
func (m *TimeStamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeStamp.Unmarshal(m, b)
}
func (m *TimeStamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeStamp.Marshal(b, m, deterministic)
}
func (m *TimeStamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeStamp.Merge(m, src)
}
func (m *TimeStamp) XXX_Size() int {
	return xxx_messageInfo_TimeStamp.Size(m)
}
func (m *TimeStamp) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeStamp.DiscardUnknown(m)
}

var xxx_messageInfo_TimeStamp proto.InternalMessageInfo

func (m *TimeStamp) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PageInfo struct {
	PageNo               int32    `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TotalCount           int32    `protobuf:"varint,3,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	TotalPage            int32    `protobuf:"varint,4,opt,name=totalPage,proto3" json:"totalPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}
func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPageNo() int32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func (m *PageInfo) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PageInfo) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *PageInfo) GetTotalPage() int32 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func init() {
	proto.RegisterType((*TimeStamp)(nil), "Course.TimeStamp")
	proto.RegisterType((*PageInfo)(nil), "Course.PageInfo")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x73, 0xce, 0x2f, 0x2d, 0x2a, 0x4e,
	0x95, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x26, 0x95, 0xa6, 0xe9, 0x97,
	0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x14, 0x2a, 0xb9, 0x72, 0x71, 0x86, 0x64,
	0xe6, 0xa6, 0x06, 0x83, 0x84, 0x84, 0x2c, 0xb8, 0x38, 0xe1, 0xf2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x52, 0x7a, 0x10, 0x13, 0xf4, 0x60, 0x26, 0xe8, 0x85, 0xc0, 0x54, 0x04, 0x21, 0x14,
	0x2b, 0xd5, 0x70, 0x71, 0x04, 0x24, 0xa6, 0xa7, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x89, 0x71, 0xb1,
	0x15, 0x24, 0xa6, 0xa7, 0xfa, 0xe5, 0x83, 0x8d, 0x60, 0x0d, 0x82, 0xf2, 0x84, 0xa4, 0xb8, 0x38,
	0x40, 0xac, 0xe0, 0xcc, 0xaa, 0x54, 0x09, 0x26, 0xb0, 0x0c, 0x9c, 0x2f, 0x24, 0xc7, 0xc5, 0x55,
	0x92, 0x5f, 0x92, 0x98, 0xe3, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0x0c, 0x96, 0x45, 0x12, 0x11,
	0x92, 0xe1, 0xe2, 0x04, 0xf3, 0x40, 0x96, 0x48, 0xb0, 0x80, 0xa5, 0x11, 0x02, 0x49, 0x6c, 0x60,
	0xc7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb9, 0x44, 0x63, 0x04, 0x01, 0x00, 0x00,
}