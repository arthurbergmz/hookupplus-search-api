// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging/channels/event/event.proto

package event

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShowcaseRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowcaseRequest) Reset()         { *m = ShowcaseRequest{} }
func (m *ShowcaseRequest) String() string { return proto.CompactTextString(m) }
func (*ShowcaseRequest) ProtoMessage()    {}
func (*ShowcaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{0}
}
func (m *ShowcaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowcaseRequest.Unmarshal(m, b)
}
func (m *ShowcaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowcaseRequest.Marshal(b, m, deterministic)
}
func (dst *ShowcaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowcaseRequest.Merge(dst, src)
}
func (m *ShowcaseRequest) XXX_Size() int {
	return xxx_messageInfo_ShowcaseRequest.Size(m)
}
func (m *ShowcaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowcaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowcaseRequest proto.InternalMessageInfo

func (m *ShowcaseRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ShowcaseResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowcaseResponse) Reset()         { *m = ShowcaseResponse{} }
func (m *ShowcaseResponse) String() string { return proto.CompactTextString(m) }
func (*ShowcaseResponse) ProtoMessage()    {}
func (*ShowcaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{1}
}
func (m *ShowcaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowcaseResponse.Unmarshal(m, b)
}
func (m *ShowcaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowcaseResponse.Marshal(b, m, deterministic)
}
func (dst *ShowcaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowcaseResponse.Merge(dst, src)
}
func (m *ShowcaseResponse) XXX_Size() int {
	return xxx_messageInfo_ShowcaseResponse.Size(m)
}
func (m *ShowcaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowcaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowcaseResponse proto.InternalMessageInfo

func (m *ShowcaseResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateResponse struct {
	Response             bool     `protobuf:"varint,1,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{3}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetResponse() bool {
	if m != nil {
		return m.Response
	}
	return false
}

type UpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{4}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateResponse struct {
	Response             bool     `protobuf:"varint,1,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{5}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetResponse() bool {
	if m != nil {
		return m.Response
	}
	return false
}

type RegisterUserTargetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserTargetRequest) Reset()         { *m = RegisterUserTargetRequest{} }
func (m *RegisterUserTargetRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUserTargetRequest) ProtoMessage()    {}
func (*RegisterUserTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{6}
}
func (m *RegisterUserTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserTargetRequest.Unmarshal(m, b)
}
func (m *RegisterUserTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserTargetRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterUserTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserTargetRequest.Merge(dst, src)
}
func (m *RegisterUserTargetRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUserTargetRequest.Size(m)
}
func (m *RegisterUserTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserTargetRequest proto.InternalMessageInfo

func (m *RegisterUserTargetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterUserTargetRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *RegisterUserTargetRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type RegisterUserTargetResponse struct {
	Response             bool     `protobuf:"varint,1,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserTargetResponse) Reset()         { *m = RegisterUserTargetResponse{} }
func (m *RegisterUserTargetResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterUserTargetResponse) ProtoMessage()    {}
func (*RegisterUserTargetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_fb27057cebc432e7, []int{7}
}
func (m *RegisterUserTargetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserTargetResponse.Unmarshal(m, b)
}
func (m *RegisterUserTargetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserTargetResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterUserTargetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserTargetResponse.Merge(dst, src)
}
func (m *RegisterUserTargetResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterUserTargetResponse.Size(m)
}
func (m *RegisterUserTargetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserTargetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserTargetResponse proto.InternalMessageInfo

func (m *RegisterUserTargetResponse) GetResponse() bool {
	if m != nil {
		return m.Response
	}
	return false
}

func init() {
	proto.RegisterType((*ShowcaseRequest)(nil), "event.ShowcaseRequest")
	proto.RegisterType((*ShowcaseResponse)(nil), "event.ShowcaseResponse")
	proto.RegisterType((*CreateRequest)(nil), "event.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "event.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "event.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "event.UpdateResponse")
	proto.RegisterType((*RegisterUserTargetRequest)(nil), "event.RegisterUserTargetRequest")
	proto.RegisterType((*RegisterUserTargetResponse)(nil), "event.RegisterUserTargetResponse")
}

func init() {
	proto.RegisterFile("messaging/channels/event/event.proto", fileDescriptor_event_fb27057cebc432e7)
}

var fileDescriptor_event_fb27057cebc432e7 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4b, 0xc7, 0x30,
	0x10, 0x85, 0xf9, 0x55, 0xad, 0xf5, 0xc0, 0x2a, 0x19, 0xa4, 0x76, 0x51, 0x82, 0x83, 0x82, 0xb4,
	0x83, 0x8b, 0xb3, 0x4e, 0xae, 0xd5, 0x22, 0xb8, 0xc5, 0xf6, 0x48, 0x03, 0x9a, 0xd4, 0x5c, 0xaa,
	0xff, 0xbe, 0x34, 0x69, 0x2b, 0x82, 0x55, 0x97, 0x70, 0x2f, 0x7c, 0xf9, 0x08, 0xef, 0xe0, 0xec,
	0x15, 0x89, 0x84, 0x54, 0x5a, 0x96, 0x4d, 0x27, 0xb4, 0xc6, 0x17, 0x2a, 0xf1, 0x1d, 0xb5, 0x0b,
	0x67, 0xd1, 0x5b, 0xe3, 0x0c, 0xdb, 0xf1, 0x81, 0x5f, 0xc0, 0xc1, 0x7d, 0x67, 0x3e, 0x1a, 0x41,
	0x58, 0xe1, 0xdb, 0x80, 0xe4, 0xd8, 0x11, 0xc4, 0x03, 0xa1, 0xbd, 0x6b, 0xb3, 0xcd, 0xe9, 0xe6,
	0x7c, 0xaf, 0x9a, 0x12, 0x2f, 0xe0, 0xf0, 0x0b, 0xa5, 0xde, 0x68, 0x42, 0x96, 0x43, 0x62, 0xa7,
	0x79, 0xa2, 0x97, 0xcc, 0x4f, 0x60, 0xff, 0xd6, 0xa2, 0x70, 0x8b, 0x38, 0x85, 0x48, 0xcd, 0xd2,
	0x48, 0xb5, 0xfc, 0x12, 0xd2, 0x19, 0x58, 0xd1, 0x25, 0xdf, 0x75, 0x75, 0xdf, 0xfe, 0xae, 0x9b,
	0x81, 0x7f, 0xe8, 0x1e, 0xe1, 0xb8, 0x42, 0xa9, 0xc8, 0xa1, 0xad, 0x09, 0xed, 0x83, 0xb0, 0x12,
	0xdd, 0x8a, 0x9a, 0x31, 0xd8, 0x1e, 0x4b, 0xc8, 0x22, 0x7f, 0xe3, 0xe7, 0xb1, 0x26, 0xe7, 0x1f,
	0x65, 0x5b, 0xa1, 0xa6, 0x90, 0xf8, 0x35, 0xe4, 0x3f, 0x89, 0xff, 0xfe, 0xd2, 0xcd, 0xee, 0x53,
	0x58, 0xca, 0x73, 0xec, 0x57, 0x74, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xf1, 0xf7, 0xde,
	0xca, 0x01, 0x00, 0x00,
}
