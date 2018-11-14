// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatbothub.proto

package chatbothub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type EventRequest struct {
	EventType            string   `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientType           string   `protobuf:"bytes,4,opt,name=clientType,proto3" json:"clientType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{0}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EventRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *EventRequest) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

type EventReply struct {
	EventType            string   `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientType           string   `protobuf:"bytes,4,opt,name=clientType,proto3" json:"clientType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventReply) Reset()         { *m = EventReply{} }
func (m *EventReply) String() string { return proto.CompactTextString(m) }
func (*EventReply) ProtoMessage()    {}
func (*EventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{1}
}

func (m *EventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventReply.Unmarshal(m, b)
}
func (m *EventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventReply.Marshal(b, m, deterministic)
}
func (m *EventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReply.Merge(m, src)
}
func (m *EventReply) XXX_Size() int {
	return xxx_messageInfo_EventReply.Size(m)
}
func (m *EventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReply.DiscardUnknown(m)
}

var xxx_messageInfo_EventReply proto.InternalMessageInfo

func (m *EventReply) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventReply) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EventReply) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *EventReply) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

type BotsRequest struct {
	Logins               []string `protobuf:"bytes,1,rep,name=logins,proto3" json:"logins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotsRequest) Reset()         { *m = BotsRequest{} }
func (m *BotsRequest) String() string { return proto.CompactTextString(m) }
func (*BotsRequest) ProtoMessage()    {}
func (*BotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{2}
}

func (m *BotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotsRequest.Unmarshal(m, b)
}
func (m *BotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotsRequest.Marshal(b, m, deterministic)
}
func (m *BotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsRequest.Merge(m, src)
}
func (m *BotsRequest) XXX_Size() int {
	return xxx_messageInfo_BotsRequest.Size(m)
}
func (m *BotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotsRequest proto.InternalMessageInfo

func (m *BotsRequest) GetLogins() []string {
	if m != nil {
		return m.Logins
	}
	return nil
}

type BotsReply struct {
	BotsInfo             []*BotsInfo `protobuf:"bytes,1,rep,name=botsInfo,proto3" json:"botsInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BotsReply) Reset()         { *m = BotsReply{} }
func (m *BotsReply) String() string { return proto.CompactTextString(m) }
func (*BotsReply) ProtoMessage()    {}
func (*BotsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{3}
}

func (m *BotsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotsReply.Unmarshal(m, b)
}
func (m *BotsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotsReply.Marshal(b, m, deterministic)
}
func (m *BotsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsReply.Merge(m, src)
}
func (m *BotsReply) XXX_Size() int {
	return xxx_messageInfo_BotsReply.Size(m)
}
func (m *BotsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotsReply proto.InternalMessageInfo

func (m *BotsReply) GetBotsInfo() []*BotsInfo {
	if m != nil {
		return m.BotsInfo
	}
	return nil
}

type BotsInfo struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientType           string   `protobuf:"bytes,2,opt,name=clientType,proto3" json:"clientType,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	StartAt              int64    `protobuf:"varint,4,opt,name=startAt,proto3" json:"startAt,omitempty"`
	LastPing             int64    `protobuf:"varint,5,opt,name=lastPing,proto3" json:"lastPing,omitempty"`
	Login                string   `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	LoginInfo            string   `protobuf:"bytes,7,opt,name=loginInfo,proto3" json:"loginInfo,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	FilterInfo           string   `protobuf:"bytes,9,opt,name=filterInfo,proto3" json:"filterInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotsInfo) Reset()         { *m = BotsInfo{} }
func (m *BotsInfo) String() string { return proto.CompactTextString(m) }
func (*BotsInfo) ProtoMessage()    {}
func (*BotsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{4}
}

func (m *BotsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotsInfo.Unmarshal(m, b)
}
func (m *BotsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotsInfo.Marshal(b, m, deterministic)
}
func (m *BotsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsInfo.Merge(m, src)
}
func (m *BotsInfo) XXX_Size() int {
	return xxx_messageInfo_BotsInfo.Size(m)
}
func (m *BotsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BotsInfo proto.InternalMessageInfo

func (m *BotsInfo) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *BotsInfo) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *BotsInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BotsInfo) GetStartAt() int64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *BotsInfo) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

func (m *BotsInfo) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *BotsInfo) GetLoginInfo() string {
	if m != nil {
		return m.LoginInfo
	}
	return ""
}

func (m *BotsInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BotsInfo) GetFilterInfo() string {
	if m != nil {
		return m.FilterInfo
	}
	return ""
}

type BotLoginRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientType           string   `protobuf:"bytes,2,opt,name=clientType,proto3" json:"clientType,omitempty"`
	Login                string   `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	NotifyUrl            string   `protobuf:"bytes,5,opt,name=notifyUrl,proto3" json:"notifyUrl,omitempty"`
	LoginInfo            string   `protobuf:"bytes,6,opt,name=loginInfo,proto3" json:"loginInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotLoginRequest) Reset()         { *m = BotLoginRequest{} }
func (m *BotLoginRequest) String() string { return proto.CompactTextString(m) }
func (*BotLoginRequest) ProtoMessage()    {}
func (*BotLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{5}
}

func (m *BotLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotLoginRequest.Unmarshal(m, b)
}
func (m *BotLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotLoginRequest.Marshal(b, m, deterministic)
}
func (m *BotLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotLoginRequest.Merge(m, src)
}
func (m *BotLoginRequest) XXX_Size() int {
	return xxx_messageInfo_BotLoginRequest.Size(m)
}
func (m *BotLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotLoginRequest proto.InternalMessageInfo

func (m *BotLoginRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *BotLoginRequest) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *BotLoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *BotLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BotLoginRequest) GetNotifyUrl() string {
	if m != nil {
		return m.NotifyUrl
	}
	return ""
}

func (m *BotLoginRequest) GetLoginInfo() string {
	if m != nil {
		return m.LoginInfo
	}
	return ""
}

type BotLoginReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotLoginReply) Reset()         { *m = BotLoginReply{} }
func (m *BotLoginReply) String() string { return proto.CompactTextString(m) }
func (*BotLoginReply) ProtoMessage()    {}
func (*BotLoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{6}
}

func (m *BotLoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotLoginReply.Unmarshal(m, b)
}
func (m *BotLoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotLoginReply.Marshal(b, m, deterministic)
}
func (m *BotLoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotLoginReply.Merge(m, src)
}
func (m *BotLoginReply) XXX_Size() int {
	return xxx_messageInfo_BotLoginReply.Size(m)
}
func (m *BotLoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotLoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotLoginReply proto.InternalMessageInfo

func (m *BotLoginReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type BotActionRequest struct {
	ActionRequestId      string   `protobuf:"bytes,1,opt,name=actionRequestId,proto3" json:"actionRequestId,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	ActionType           string   `protobuf:"bytes,3,opt,name=actionType,proto3" json:"actionType,omitempty"`
	ActionBody           string   `protobuf:"bytes,4,opt,name=actionBody,proto3" json:"actionBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotActionRequest) Reset()         { *m = BotActionRequest{} }
func (m *BotActionRequest) String() string { return proto.CompactTextString(m) }
func (*BotActionRequest) ProtoMessage()    {}
func (*BotActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{7}
}

func (m *BotActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotActionRequest.Unmarshal(m, b)
}
func (m *BotActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotActionRequest.Marshal(b, m, deterministic)
}
func (m *BotActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotActionRequest.Merge(m, src)
}
func (m *BotActionRequest) XXX_Size() int {
	return xxx_messageInfo_BotActionRequest.Size(m)
}
func (m *BotActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotActionRequest proto.InternalMessageInfo

func (m *BotActionRequest) GetActionRequestId() string {
	if m != nil {
		return m.ActionRequestId
	}
	return ""
}

func (m *BotActionRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *BotActionRequest) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *BotActionRequest) GetActionBody() string {
	if m != nil {
		return m.ActionBody
	}
	return ""
}

type BotActionReply struct {
	ActionRequestId      string   `protobuf:"bytes,1,opt,name=actionRequestId,proto3" json:"actionRequestId,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotActionReply) Reset()         { *m = BotActionReply{} }
func (m *BotActionReply) String() string { return proto.CompactTextString(m) }
func (*BotActionReply) ProtoMessage()    {}
func (*BotActionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1f640cec0d9d68, []int{8}
}

func (m *BotActionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotActionReply.Unmarshal(m, b)
}
func (m *BotActionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotActionReply.Marshal(b, m, deterministic)
}
func (m *BotActionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotActionReply.Merge(m, src)
}
func (m *BotActionReply) XXX_Size() int {
	return xxx_messageInfo_BotActionReply.Size(m)
}
func (m *BotActionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotActionReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotActionReply proto.InternalMessageInfo

func (m *BotActionReply) GetActionRequestId() string {
	if m != nil {
		return m.ActionRequestId
	}
	return ""
}

func (m *BotActionReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BotActionReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BotActionReply) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRequest)(nil), "chatbothub.EventRequest")
	proto.RegisterType((*EventReply)(nil), "chatbothub.EventReply")
	proto.RegisterType((*BotsRequest)(nil), "chatbothub.BotsRequest")
	proto.RegisterType((*BotsReply)(nil), "chatbothub.BotsReply")
	proto.RegisterType((*BotsInfo)(nil), "chatbothub.BotsInfo")
	proto.RegisterType((*BotLoginRequest)(nil), "chatbothub.BotLoginRequest")
	proto.RegisterType((*BotLoginReply)(nil), "chatbothub.BotLoginReply")
	proto.RegisterType((*BotActionRequest)(nil), "chatbothub.BotActionRequest")
	proto.RegisterType((*BotActionReply)(nil), "chatbothub.BotActionReply")
}

func init() { proto.RegisterFile("chatbothub.proto", fileDescriptor_0b1f640cec0d9d68) }

var fileDescriptor_0b1f640cec0d9d68 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x7e, 0xec, 0x09, 0x90, 0x68, 0x55, 0x8a, 0x71, 0xab, 0x2a, 0x58, 0x42, 0xca,
	0x29, 0x8a, 0xca, 0x11, 0x71, 0xa8, 0xa1, 0x2a, 0x95, 0x38, 0x44, 0x56, 0x79, 0x00, 0x3b, 0xd9,
	0x24, 0x16, 0xae, 0x37, 0x64, 0xc7, 0x44, 0xa9, 0x78, 0x0a, 0x5e, 0x80, 0x07, 0xe1, 0xd5, 0x38,
	0xa0, 0x1d, 0xff, 0x6d, 0x4c, 0x41, 0x48, 0x48, 0xdc, 0x66, 0xbe, 0x19, 0xcf, 0x7e, 0xdf, 0xb7,
	0x9e, 0x85, 0xe1, 0x7c, 0x1d, 0x62, 0x24, 0x70, 0x9d, 0x45, 0x93, 0xcd, 0x56, 0xa0, 0x60, 0x50,
	0x23, 0xde, 0x17, 0x78, 0x78, 0xf9, 0x99, 0xa7, 0x18, 0xf0, 0x4f, 0x19, 0x97, 0xc8, 0x4e, 0xc1,
	0xe6, 0x2a, 0xbf, 0xd9, 0x6f, 0xb8, 0x63, 0x8c, 0x8c, 0xb1, 0x1d, 0xd4, 0x00, 0x63, 0xd0, 0x8e,
	0xc4, 0x62, 0xef, 0xb4, 0xa8, 0x40, 0x31, 0x73, 0xc1, 0x9a, 0x27, 0x31, 0x4f, 0xf1, 0x7a, 0xe1,
	0x98, 0x84, 0x57, 0x39, 0x3b, 0x03, 0xc8, 0x63, 0x1a, 0xd7, 0xa6, 0xaa, 0x86, 0x78, 0x77, 0x00,
	0xc5, 0xe9, 0x9b, 0x64, 0xff, 0x9f, 0xcf, 0x7e, 0x01, 0x7d, 0x5f, 0xa0, 0x2c, 0x85, 0x1f, 0x43,
	0x37, 0x11, 0xab, 0x38, 0x95, 0x8e, 0x31, 0x32, 0xc7, 0x76, 0x50, 0x64, 0xde, 0x6b, 0xb0, 0xf3,
	0x36, 0xc5, 0x70, 0x0a, 0x56, 0x24, 0x50, 0x5e, 0xa7, 0x4b, 0x41, 0x6d, 0xfd, 0xf3, 0xa3, 0x89,
	0x66, 0xaf, 0x5f, 0xd4, 0x82, 0xaa, 0xcb, 0xfb, 0x61, 0x80, 0x55, 0xc2, 0x07, 0x74, 0x8d, 0x3f,
	0xd2, 0x6d, 0x35, 0xe9, 0x2a, 0xf9, 0x69, 0x78, 0xcb, 0x0b, 0x99, 0x14, 0x33, 0x07, 0x7a, 0x12,
	0xc3, 0x2d, 0x5e, 0x20, 0xe9, 0x33, 0x83, 0x32, 0x55, 0x27, 0x25, 0xa1, 0xc4, 0x59, 0x9c, 0xae,
	0x9c, 0x0e, 0x95, 0xaa, 0x9c, 0x1d, 0x41, 0x87, 0xb4, 0x39, 0x5d, 0x1a, 0x95, 0x27, 0xca, 0x7c,
	0x0a, 0x48, 0x5b, 0x2f, 0x37, 0xbf, 0x02, 0x94, 0x3b, 0x12, 0x43, 0xcc, 0xa4, 0x63, 0x8d, 0x8c,
	0x71, 0x27, 0x28, 0x32, 0xc5, 0x7a, 0x19, 0x27, 0xc8, 0xb7, 0xf4, 0x99, 0x9d, 0xb3, 0xae, 0x11,
	0xef, 0xbb, 0x01, 0x03, 0x5f, 0xe0, 0x7b, 0x35, 0xa8, 0x74, 0xfa, 0x5f, 0x5c, 0xa8, 0xb8, 0x9b,
	0x3a, 0x77, 0x17, 0xac, 0x4d, 0x28, 0xe5, 0x4e, 0x6c, 0x17, 0xc5, 0x45, 0x57, 0xb9, 0xd2, 0x95,
	0x0a, 0x8c, 0x97, 0xfb, 0x0f, 0xdb, 0x84, 0xac, 0xb0, 0x83, 0x1a, 0x38, 0x54, 0xdd, 0x6d, 0xa8,
	0xf6, 0x9e, 0xc3, 0xa3, 0x9a, 0xbc, 0xba, 0xff, 0x21, 0x98, 0xb7, 0x72, 0x55, 0xb0, 0x56, 0xa1,
	0xf7, 0xd5, 0x80, 0xa1, 0x2f, 0xf0, 0x62, 0x8e, 0xb1, 0xa8, 0x14, 0x8e, 0x61, 0x10, 0xea, 0x40,
	0x25, 0xb4, 0x09, 0xd7, 0x7a, 0x5a, 0xba, 0x9e, 0x33, 0x80, 0xbc, 0x91, 0x5c, 0xc8, 0xa5, 0x6a,
	0x48, 0x5d, 0xf7, 0xd5, 0x42, 0xb4, 0xf5, 0xba, 0x42, 0xbc, 0x3b, 0x78, 0xac, 0x71, 0x52, 0xc4,
	0xff, 0x9e, 0x91, 0xfa, 0xa7, 0xb2, 0xf9, 0x9c, 0x4b, 0x49, 0x9c, 0xac, 0xa0, 0x4c, 0x4b, 0xf1,
	0x66, 0x25, 0xbe, 0x5a, 0xc9, 0x76, 0xbd, 0x92, 0xe7, 0xdf, 0x5a, 0x00, 0x6f, 0xd6, 0x21, 0xfa,
	0x02, 0xdf, 0x65, 0x11, 0xbb, 0x84, 0x3e, 0x6d, 0xf8, 0x4d, 0x96, 0xa6, 0x3c, 0x61, 0x8e, 0xbe,
	0x2e, 0xfa, 0xc3, 0xe3, 0x1e, 0xdf, 0x53, 0xd9, 0x24, 0x7b, 0xef, 0xc1, 0xd8, 0x98, 0x1a, 0xec,
	0x15, 0xf4, 0xae, 0xb8, 0x9a, 0x29, 0xd9, 0xd3, 0xe6, 0xc6, 0x95, 0x13, 0x9e, 0xfc, 0x5a, 0xa0,
	0x01, 0xec, 0x2d, 0xad, 0x20, 0x5d, 0x23, 0x3b, 0x69, 0x34, 0xe9, 0x7f, 0xa6, 0xfb, 0xec, 0xfe,
	0x62, 0x3e, 0xe5, 0x8a, 0x1e, 0x82, 0xdc, 0x54, 0x76, 0xda, 0xe8, 0x3c, 0xb8, 0x7f, 0xd7, 0xfd,
	0x4d, 0x95, 0x06, 0xf9, 0x53, 0x38, 0x49, 0x39, 0x4e, 0xd6, 0xe1, 0xee, 0xe3, 0x2e, 0xc6, 0xf5,
	0x2e, 0x4e, 0x17, 0x5a, 0xbf, 0x3f, 0xa8, 0xdd, 0x9b, 0xa9, 0xe7, 0x7a, 0x66, 0x44, 0x5d, 0x7a,
	0xb7, 0x5f, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x8b, 0x82, 0xe3, 0xcb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatBotHubClient is the client API for ChatBotHub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatBotHubClient interface {
	// bots only use eventtunnel to communicate
	EventTunnel(ctx context.Context, opts ...grpc.CallOption) (ChatBotHub_EventTunnelClient, error)
	// below are for internal web api
	GetBots(ctx context.Context, in *BotsRequest, opts ...grpc.CallOption) (*BotsReply, error)
	BotLogin(ctx context.Context, in *BotLoginRequest, opts ...grpc.CallOption) (*BotLoginReply, error)
	BotAction(ctx context.Context, in *BotActionRequest, opts ...grpc.CallOption) (*BotActionReply, error)
}

type chatBotHubClient struct {
	cc *grpc.ClientConn
}

func NewChatBotHubClient(cc *grpc.ClientConn) ChatBotHubClient {
	return &chatBotHubClient{cc}
}

func (c *chatBotHubClient) EventTunnel(ctx context.Context, opts ...grpc.CallOption) (ChatBotHub_EventTunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatBotHub_serviceDesc.Streams[0], "/chatbothub.ChatBotHub/EventTunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBotHubEventTunnelClient{stream}
	return x, nil
}

type ChatBotHub_EventTunnelClient interface {
	Send(*EventRequest) error
	Recv() (*EventReply, error)
	grpc.ClientStream
}

type chatBotHubEventTunnelClient struct {
	grpc.ClientStream
}

func (x *chatBotHubEventTunnelClient) Send(m *EventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatBotHubEventTunnelClient) Recv() (*EventReply, error) {
	m := new(EventReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatBotHubClient) GetBots(ctx context.Context, in *BotsRequest, opts ...grpc.CallOption) (*BotsReply, error) {
	out := new(BotsReply)
	err := c.cc.Invoke(ctx, "/chatbothub.ChatBotHub/GetBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatBotHubClient) BotLogin(ctx context.Context, in *BotLoginRequest, opts ...grpc.CallOption) (*BotLoginReply, error) {
	out := new(BotLoginReply)
	err := c.cc.Invoke(ctx, "/chatbothub.ChatBotHub/BotLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatBotHubClient) BotAction(ctx context.Context, in *BotActionRequest, opts ...grpc.CallOption) (*BotActionReply, error) {
	out := new(BotActionReply)
	err := c.cc.Invoke(ctx, "/chatbothub.ChatBotHub/BotAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatBotHubServer is the server API for ChatBotHub service.
type ChatBotHubServer interface {
	// bots only use eventtunnel to communicate
	EventTunnel(ChatBotHub_EventTunnelServer) error
	// below are for internal web api
	GetBots(context.Context, *BotsRequest) (*BotsReply, error)
	BotLogin(context.Context, *BotLoginRequest) (*BotLoginReply, error)
	BotAction(context.Context, *BotActionRequest) (*BotActionReply, error)
}

func RegisterChatBotHubServer(s *grpc.Server, srv ChatBotHubServer) {
	s.RegisterService(&_ChatBotHub_serviceDesc, srv)
}

func _ChatBotHub_EventTunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatBotHubServer).EventTunnel(&chatBotHubEventTunnelServer{stream})
}

type ChatBotHub_EventTunnelServer interface {
	Send(*EventReply) error
	Recv() (*EventRequest, error)
	grpc.ServerStream
}

type chatBotHubEventTunnelServer struct {
	grpc.ServerStream
}

func (x *chatBotHubEventTunnelServer) Send(m *EventReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatBotHubEventTunnelServer) Recv() (*EventRequest, error) {
	m := new(EventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatBotHub_GetBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotHubServer).GetBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbothub.ChatBotHub/GetBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotHubServer).GetBots(ctx, req.(*BotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatBotHub_BotLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotHubServer).BotLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbothub.ChatBotHub/BotLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotHubServer).BotLogin(ctx, req.(*BotLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatBotHub_BotAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotHubServer).BotAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbothub.ChatBotHub/BotAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotHubServer).BotAction(ctx, req.(*BotActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatBotHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatbothub.ChatBotHub",
	HandlerType: (*ChatBotHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBots",
			Handler:    _ChatBotHub_GetBots_Handler,
		},
		{
			MethodName: "BotLogin",
			Handler:    _ChatBotHub_BotLogin_Handler,
		},
		{
			MethodName: "BotAction",
			Handler:    _ChatBotHub_BotAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventTunnel",
			Handler:       _ChatBotHub_EventTunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chatbothub.proto",
}
