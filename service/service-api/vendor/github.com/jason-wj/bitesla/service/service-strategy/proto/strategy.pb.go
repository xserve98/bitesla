// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/strategy.proto

package bitesla_srv_strategy

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Language int32

const (
	Language_GOLANG Language = 0
	Language_PYTHON Language = 1
)

var Language_name = map[int32]string{
	0: "GOLANG",
	1: "PYTHON",
}

var Language_value = map[string]int32{
	"GOLANG": 0,
	"PYTHON": 1,
}

func (x Language) String() string {
	return proto.EnumName(Language_name, int32(x))
}

func (Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{0}
}

type StrategyInfo struct {
	CurrentLoginUserID   int64    `protobuf:"varint,1,opt,name=currentLoginUserID,proto3" json:"currentLoginUserID,omitempty"`
	StrategyId           int64    `protobuf:"varint,2,opt,name=strategyId,proto3" json:"strategyId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Script               string   `protobuf:"bytes,5,opt,name=script,proto3" json:"script,omitempty"`
	Page                 int32    `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	CreateTime           int64    `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Language             int32    `protobuf:"varint,10,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StrategyInfo) Reset()         { *m = StrategyInfo{} }
func (m *StrategyInfo) String() string { return proto.CompactTextString(m) }
func (*StrategyInfo) ProtoMessage()    {}
func (*StrategyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{0}
}

func (m *StrategyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyInfo.Unmarshal(m, b)
}
func (m *StrategyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyInfo.Marshal(b, m, deterministic)
}
func (m *StrategyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyInfo.Merge(m, src)
}
func (m *StrategyInfo) XXX_Size() int {
	return xxx_messageInfo_StrategyInfo.Size(m)
}
func (m *StrategyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyInfo proto.InternalMessageInfo

func (m *StrategyInfo) GetCurrentLoginUserID() int64 {
	if m != nil {
		return m.CurrentLoginUserID
	}
	return 0
}

func (m *StrategyInfo) GetStrategyId() int64 {
	if m != nil {
		return m.StrategyId
	}
	return 0
}

func (m *StrategyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StrategyInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StrategyInfo) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *StrategyInfo) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *StrategyInfo) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *StrategyInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *StrategyInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *StrategyInfo) GetLanguage() int32 {
	if m != nil {
		return m.Language
	}
	return 0
}

type StrategyInfos struct {
	StrategyInfos        []*StrategyInfo `protobuf:"bytes,1,rep,name=strategyInfos,proto3" json:"strategyInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StrategyInfos) Reset()         { *m = StrategyInfos{} }
func (m *StrategyInfos) String() string { return proto.CompactTextString(m) }
func (*StrategyInfos) ProtoMessage()    {}
func (*StrategyInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9cfd4cb1c623bb, []int{1}
}

func (m *StrategyInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyInfos.Unmarshal(m, b)
}
func (m *StrategyInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyInfos.Marshal(b, m, deterministic)
}
func (m *StrategyInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyInfos.Merge(m, src)
}
func (m *StrategyInfos) XXX_Size() int {
	return xxx_messageInfo_StrategyInfos.Size(m)
}
func (m *StrategyInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyInfos.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyInfos proto.InternalMessageInfo

func (m *StrategyInfos) GetStrategyInfos() []*StrategyInfo {
	if m != nil {
		return m.StrategyInfos
	}
	return nil
}

func init() {
	proto.RegisterEnum("bitesla.srv.strategy.Language", Language_name, Language_value)
	proto.RegisterType((*StrategyInfo)(nil), "bitesla.srv.strategy.StrategyInfo")
	proto.RegisterType((*StrategyInfos)(nil), "bitesla.srv.strategy.StrategyInfos")
}

func init() { proto.RegisterFile("proto/strategy.proto", fileDescriptor_8a9cfd4cb1c623bb) }

var fileDescriptor_8a9cfd4cb1c623bb = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0x8d, 0x7f, 0xf2, 0xc6, 0x51, 0x5f, 0xec, 0x22, 0x65, 0xf1, 0x50, 0x42, 0x7a, 0x09,
	0x3d, 0xa4, 0x60, 0x3f, 0x41, 0x41, 0x50, 0x21, 0xa8, 0xa4, 0x96, 0x62, 0xa1, 0x87, 0xa8, 0xd3,
	0xb0, 0x10, 0x93, 0xb0, 0xbb, 0x29, 0xb4, 0x9f, 0xaa, 0x5f, 0xaf, 0xb7, 0x92, 0x8d, 0xd1, 0x14,
	0x84, 0x5a, 0xf0, 0x36, 0xcf, 0xf3, 0xcc, 0xfe, 0x66, 0x18, 0x58, 0xe8, 0x25, 0x3c, 0x96, 0xf1,
	0xad, 0x90, 0xdc, 0x97, 0x18, 0xbc, 0x3b, 0x4a, 0x92, 0xde, 0x8a, 0x49, 0x14, 0xa1, 0xef, 0x08,
	0xfe, 0xe6, 0x14, 0x99, 0xf5, 0x59, 0x85, 0xf6, 0xc3, 0x4e, 0x4c, 0xa2, 0xd7, 0x98, 0x38, 0x40,
	0xd6, 0x29, 0xe7, 0x18, 0x49, 0x37, 0x0e, 0x58, 0xf4, 0x28, 0x90, 0x4f, 0x86, 0x54, 0x33, 0x35,
	0xbb, 0xe6, 0x1d, 0x49, 0xc8, 0x15, 0x40, 0x01, 0x9b, 0x6c, 0x68, 0x55, 0xf5, 0x95, 0x1c, 0x42,
	0xa0, 0x1e, 0xf9, 0x5b, 0xa4, 0x35, 0x53, 0xb3, 0x9b, 0x9e, 0xaa, 0x89, 0x09, 0xad, 0x0d, 0x8a,
	0x35, 0x67, 0x89, 0x64, 0x71, 0x44, 0xeb, 0x2a, 0x2a, 0x5b, 0xe4, 0x12, 0xf4, 0x5c, 0xd0, 0x86,
	0x0a, 0x77, 0x2a, 0xa3, 0x25, 0x7e, 0x80, 0x54, 0x37, 0x35, 0xbb, 0xe1, 0xa9, 0x3a, 0xf3, 0x04,
	0xfb, 0x40, 0xfa, 0x2f, 0xf7, 0xb2, 0x3a, 0xdb, 0x6a, 0xcd, 0xd1, 0x97, 0xb8, 0x60, 0x5b, 0xa4,
	0x46, 0xbe, 0xd5, 0xc1, 0xc9, 0xf2, 0x34, 0xd9, 0x14, 0x79, 0x33, 0xcf, 0x0f, 0x0e, 0xe9, 0x83,
	0x11, 0xfa, 0x51, 0x90, 0x66, 0xb3, 0x40, 0x71, 0xf7, 0xda, 0x5a, 0x42, 0xa7, 0x7c, 0x31, 0x41,
	0xc6, 0xd0, 0x11, 0x65, 0x83, 0x6a, 0x66, 0xcd, 0x6e, 0x0d, 0x2c, 0xe7, 0xd8, 0xc5, 0x9d, 0xf2,
	0x5b, 0xef, 0xe7, 0xc3, 0x1b, 0x0b, 0x0c, 0x77, 0x37, 0x86, 0x00, 0xe8, 0xa3, 0x99, 0x7b, 0x3f,
	0x1d, 0x75, 0x2b, 0x59, 0x3d, 0x5f, 0x2e, 0xc6, 0xb3, 0x69, 0x57, 0x1b, 0x7c, 0x55, 0xc1, 0x28,
	0x18, 0x64, 0x09, 0x6d, 0x97, 0x09, 0xb9, 0xd7, 0x27, 0xcc, 0xec, 0x5f, 0xff, 0xde, 0x23, 0xac,
	0x0a, 0x79, 0x82, 0xd6, 0x3c, 0xfd, 0x1b, 0xf9, 0x84, 0x1e, 0xab, 0x42, 0x5e, 0xe0, 0x62, 0x84,
	0x7b, 0xf0, 0x10, 0xa5, 0xcf, 0xc2, 0x33, 0xe2, 0x9f, 0xe1, 0xff, 0x10, 0x43, 0x94, 0x78, 0xfe,
	0xd5, 0x57, 0xba, 0xfa, 0x4a, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xc9, 0x2e, 0xa3,
	0x62, 0x03, 0x00, 0x00,
}
