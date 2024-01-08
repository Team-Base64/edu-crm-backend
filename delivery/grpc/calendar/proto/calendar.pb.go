// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery/grpc/calendar/proto/calendar.proto

package calendar

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

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{0}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

type EventData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate            string   `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	ClassID              int32    `protobuf:"varint,6,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventData) Reset()         { *m = EventData{} }
func (m *EventData) String() string { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()    {}
func (*EventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{1}
}

func (m *EventData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventData.Unmarshal(m, b)
}
func (m *EventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventData.Marshal(b, m, deterministic)
}
func (m *EventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventData.Merge(m, src)
}
func (m *EventData) XXX_Size() int {
	return xxx_messageInfo_EventData.Size(m)
}
func (m *EventData) XXX_DiscardUnknown() {
	xxx_messageInfo_EventData.DiscardUnknown(m)
}

var xxx_messageInfo_EventData proto.InternalMessageInfo

func (m *EventData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventData) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *EventData) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *EventData) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type GetEventsRequestCalendar struct {
	TeacherID            int32    `protobuf:"varint,1,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsRequestCalendar) Reset()         { *m = GetEventsRequestCalendar{} }
func (m *GetEventsRequestCalendar) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequestCalendar) ProtoMessage()    {}
func (*GetEventsRequestCalendar) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{2}
}

func (m *GetEventsRequestCalendar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequestCalendar.Unmarshal(m, b)
}
func (m *GetEventsRequestCalendar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequestCalendar.Marshal(b, m, deterministic)
}
func (m *GetEventsRequestCalendar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequestCalendar.Merge(m, src)
}
func (m *GetEventsRequestCalendar) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequestCalendar.Size(m)
}
func (m *GetEventsRequestCalendar) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequestCalendar.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequestCalendar proto.InternalMessageInfo

func (m *GetEventsRequestCalendar) GetTeacherID() int32 {
	if m != nil {
		return m.TeacherID
	}
	return 0
}

type GetEventsResponse struct {
	Events               []*EventData `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetEventsResponse) Reset()         { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()    {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{3}
}

func (m *GetEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsResponse.Unmarshal(m, b)
}
func (m *GetEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsResponse.Merge(m, src)
}
func (m *GetEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsResponse.Size(m)
}
func (m *GetEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsResponse proto.InternalMessageInfo

func (m *GetEventsResponse) GetEvents() []*EventData {
	if m != nil {
		return m.Events
	}
	return nil
}

type CreateEventRequest struct {
	CalendarID           string     `protobuf:"bytes,1,opt,name=calendarID,proto3" json:"calendarID,omitempty"`
	Event                *EventData `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{4}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetCalendarID() string {
	if m != nil {
		return m.CalendarID
	}
	return ""
}

func (m *CreateEventRequest) GetEvent() *EventData {
	if m != nil {
		return m.Event
	}
	return nil
}

type DeleteEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CalendarID           string   `protobuf:"bytes,2,opt,name=calendarID,proto3" json:"calendarID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb2b1a167884712, []int{5}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteEventRequest) GetCalendarID() string {
	if m != nil {
		return m.CalendarID
	}
	return ""
}

func init() {
	proto.RegisterType((*Nothing)(nil), "calendar.Nothing")
	proto.RegisterType((*EventData)(nil), "calendar.EventData")
	proto.RegisterType((*GetEventsRequestCalendar)(nil), "calendar.GetEventsRequestCalendar")
	proto.RegisterType((*GetEventsResponse)(nil), "calendar.GetEventsResponse")
	proto.RegisterType((*CreateEventRequest)(nil), "calendar.CreateEventRequest")
	proto.RegisterType((*DeleteEventRequest)(nil), "calendar.DeleteEventRequest")
}

func init() {
	proto.RegisterFile("delivery/grpc/calendar/proto/calendar.proto", fileDescriptor_cdb2b1a167884712)
}

var fileDescriptor_cdb2b1a167884712 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xfa, 0xa5, 0xfd, 0x72, 0x23, 0x42, 0xaf, 0x2e, 0x86, 0x5a, 0x24, 0x64, 0x55,
	0x29, 0xb4, 0x50, 0x11, 0x5c, 0x15, 0xb1, 0x11, 0xe9, 0xc6, 0x45, 0xd0, 0x8d, 0x1b, 0x19, 0x93,
	0x4b, 0x1b, 0x08, 0x49, 0x9c, 0x19, 0x0b, 0x3e, 0x8e, 0x0b, 0xdf, 0x53, 0x3a, 0xcd, 0x3f, 0xdb,
	0x74, 0x79, 0xcf, 0xb9, 0xa7, 0x67, 0xe6, 0xd7, 0x0c, 0x8c, 0x23, 0x4a, 0xe2, 0x0d, 0x89, 0xaf,
	0xe9, 0x4a, 0xe4, 0xe1, 0x34, 0xe4, 0x09, 0xa5, 0x11, 0x17, 0xd3, 0x5c, 0x64, 0x2a, 0xab, 0xc6,
	0x89, 0x1e, 0xf1, 0x7f, 0x39, 0x7b, 0x36, 0xf4, 0x9e, 0x32, 0xb5, 0x8e, 0xd3, 0x95, 0xf7, 0x63,
	0x80, 0xfd, 0xb0, 0xa1, 0x54, 0xf9, 0x5c, 0x71, 0x3c, 0x05, 0x33, 0x8e, 0x98, 0xe1, 0x1a, 0x23,
	0x3b, 0x30, 0xe3, 0x08, 0xcf, 0xc1, 0x52, 0xb1, 0x4a, 0x88, 0x99, 0x5a, 0xda, 0x0d, 0xe8, 0x82,
	0x13, 0x91, 0x0c, 0x45, 0x9c, 0xab, 0x38, 0x4b, 0x59, 0x47, 0x7b, 0x4d, 0x09, 0x87, 0x60, 0x4b,
	0xc5, 0xc5, 0xf6, 0x47, 0x89, 0xfd, 0xd3, 0x7e, 0x2d, 0x20, 0x83, 0x1e, 0xa5, 0x91, 0xf6, 0x2c,
	0xed, 0x95, 0xe3, 0xd6, 0x09, 0x13, 0x2e, 0xe5, 0xd2, 0x67, 0x5d, 0xd7, 0x18, 0x59, 0x41, 0x39,
	0x7a, 0xb7, 0xc0, 0x1e, 0x49, 0xe9, 0x93, 0xca, 0x80, 0x3e, 0x3e, 0x49, 0xaa, 0x45, 0x71, 0x9d,
	0x6d, 0x9b, 0x22, 0x1e, 0xae, 0x49, 0x2c, 0x7d, 0x7d, 0x78, 0x2b, 0xa8, 0x05, 0xef, 0x0e, 0xfa,
	0x8d, 0xa4, 0xcc, 0xb3, 0x54, 0x12, 0x8e, 0xa1, 0x4b, 0x5a, 0x61, 0x86, 0xdb, 0x19, 0x39, 0xb3,
	0xb3, 0x49, 0x05, 0xab, 0xa2, 0x11, 0x14, 0x2b, 0xde, 0x1b, 0xe0, 0x42, 0x10, 0x57, 0xa4, 0xad,
	0xa2, 0x1d, 0x2f, 0x01, 0xca, 0x4c, 0x51, 0x6b, 0x07, 0x0d, 0x05, 0xaf, 0xc0, 0xd2, 0x79, 0xcd,
	0xee, 0x48, 0xc3, 0x6e, 0xc3, 0xf3, 0x01, 0x7d, 0x4a, 0x68, 0xaf, 0x60, 0xff, 0xcf, 0xf8, 0x5b,
	0x68, 0xee, 0x17, 0xce, 0xbe, 0x4d, 0xc0, 0x92, 0xc9, 0x22, 0x4b, 0x95, 0xc8, 0x92, 0x84, 0x04,
	0x3e, 0x37, 0xee, 0x5f, 0x21, 0xf3, 0xea, 0xd3, 0x1c, 0xc3, 0x3a, 0xb8, 0x68, 0xdd, 0x29, 0x00,
	0xce, 0xc1, 0x69, 0x30, 0xc1, 0x61, 0xbd, 0x7b, 0x88, 0x6a, 0xd0, 0xaf, 0xdd, 0xe2, 0xbb, 0xc3,
	0x1b, 0x70, 0x5e, 0xf2, 0xa8, 0xca, 0xb7, 0xd1, 0x69, 0x8b, 0xcd, 0xc1, 0x69, 0x90, 0x6a, 0xd6,
	0x1e, 0x02, 0x6c, 0xc9, 0xdf, 0x9f, 0xbc, 0xc2, 0xa4, 0x7a, 0x17, 0xef, 0x5d, 0xfd, 0x30, 0xae,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x01, 0xc5, 0x52, 0x1d, 0x47, 0x03, 0x00, 0x00,
}
