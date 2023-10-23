// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservices/chatServer/gen_files/chat.proto

package chat

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

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	ChatID               int32    `protobuf:"varint,2,opt,name=chatID,proto3" json:"chatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d796cb2faf1b7d7, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetChatID() int32 {
	if m != nil {
		return m.ChatID
	}
	return 0
}

type Status struct {
	IsSuccessful         bool     `protobuf:"varint,1,opt,name=isSuccessful,proto3" json:"isSuccessful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d796cb2faf1b7d7, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetIsSuccessful() bool {
	if m != nil {
		return m.IsSuccessful
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterType((*Status)(nil), "chat.Status")
}

func init() {
	proto.RegisterFile("microservices/chatServer/gen_files/chat.proto", fileDescriptor_6d796cb2faf1b7d7)
}

var fileDescriptor_6d796cb2faf1b7d7 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x3d, 0x4f, 0x80, 0x30,
	0x10, 0x86, 0xc5, 0x60, 0xc1, 0x0b, 0x2e, 0x1d, 0x0c, 0x71, 0x22, 0x9d, 0x18, 0x14, 0x12, 0x89,
	0x7f, 0x00, 0x5d, 0x1c, 0x5c, 0xca, 0xe6, 0x62, 0x6a, 0x73, 0x40, 0x13, 0xb4, 0xa6, 0x77, 0x10,
	0x7f, 0xbe, 0xa1, 0xb0, 0xb8, 0xbd, 0xef, 0x93, 0xfb, 0x78, 0xe0, 0xe1, 0xcb, 0xd9, 0xe0, 0x09,
	0xc3, 0xe6, 0x2c, 0x52, 0x6b, 0x67, 0xc3, 0x03, 0x86, 0x0d, 0x43, 0x3b, 0xe1, 0xf7, 0xc7, 0xe8,
	0x96, 0x13, 0x36, 0x3f, 0xc1, 0xb3, 0x97, 0xe9, 0x9e, 0xd5, 0x13, 0x64, 0x6f, 0x48, 0x64, 0x26,
	0x94, 0x12, 0x52, 0xc6, 0x5f, 0x2e, 0x93, 0x2a, 0xa9, 0xaf, 0x75, 0xcc, 0xf2, 0x16, 0xc4, 0x3e,
	0xf6, 0xfa, 0x52, 0x5e, 0x56, 0x49, 0x7d, 0xa5, 0xcf, 0xa6, 0xee, 0x41, 0x0c, 0x6c, 0x78, 0x25,
	0xa9, 0xa0, 0x70, 0x34, 0xac, 0xd6, 0x22, 0xd1, 0xb8, 0x2e, 0x71, 0x3b, 0xd7, 0xff, 0xd8, 0x63,
	0x07, 0x59, 0xef, 0xf9, 0x79, 0x36, 0x2c, 0x6b, 0xc8, 0x34, 0x5a, 0x87, 0x1b, 0xca, 0x9b, 0x26,
	0xda, 0x9c, 0xef, 0xef, 0x8a, 0xa3, 0x1e, 0x67, 0xd5, 0x45, 0x9f, 0xbf, 0x8b, 0x26, 0xfa, 0x7e,
	0x8a, 0x28, 0xdc, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x3d, 0x3f, 0x7e, 0xe1, 0x00, 0x00,
	0x00,
}