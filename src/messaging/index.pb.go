// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging/index.proto

package messaging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "messaging/channels/event"
import _ "messaging/channels/handshake"
import _ "messaging/channels/user"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("messaging/index.proto", fileDescriptor_index_7b3aa5ef6630466a) }

var fileDescriptor_index_7b3aa5ef6630466a = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0x97, 0xbf, 0xb0, 0x6c, 0x5e, 0xb2, 0x75, 0x83, 0xea, 0x41, 0x87, 0x47, 0xe9,
	0x40, 0x41, 0x10, 0x04, 0x61, 0x32, 0xf0, 0xe0, 0xc5, 0xcd, 0x5d, 0x3c, 0x19, 0xdb, 0x97, 0xb6,
	0xa8, 0x49, 0xcd, 0x9b, 0xa9, 0x5f, 0xd0, 0xef, 0x25, 0x6d, 0x9a, 0xa4, 0x89, 0xc5, 0x4b, 0x68,
	0x9f, 0x3c, 0xbf, 0xe7, 0x6d, 0x9f, 0xb4, 0x24, 0x7a, 0x03, 0x44, 0x96, 0x97, 0x3c, 0x5f, 0x94,
	0x3c, 0x83, 0xaf, 0xa4, 0x92, 0x42, 0x09, 0x3a, 0xb0, 0x72, 0x7c, 0xea, 0x1c, 0x69, 0xc1, 0x38,
	0x87, 0x57, 0x5c, 0x14, 0x8c, 0x67, 0x58, 0xb0, 0x17, 0x70, 0x57, 0x1a, 0x8c, 0xe7, 0x3d, 0xee,
	0x1d, 0x82, 0x6c, 0x96, 0xd6, 0x73, 0xd2, 0xe3, 0x81, 0x0f, 0xe0, 0x4a, 0xaf, 0xda, 0x75, 0xf6,
	0xfd, 0x9f, 0x8c, 0x36, 0xc0, 0x64, 0x5a, 0xac, 0x78, 0x5e, 0x72, 0xa0, 0x17, 0x64, 0x70, 0x6b,
	0xa6, 0x51, 0x9a, 0xb8, 0xc9, 0x6b, 0x78, 0xdf, 0x01, 0xaa, 0x78, 0xec, 0x69, 0x58, 0x09, 0x8e,
	0x30, 0xff, 0x47, 0xaf, 0xc9, 0x68, 0x8b, 0x20, 0x37, 0x85, 0xf8, 0x4c, 0x19, 0x02, 0x8d, 0x92,
	0xe6, 0x59, 0xcc, 0xbd, 0xa1, 0xa7, 0xa1, 0x6c, 0x03, 0x2e, 0x09, 0xb9, 0x91, 0xc0, 0x14, 0xd4,
	0x31, 0x74, 0xac, 0x7d, 0x5a, 0x31, 0xf0, 0xc4, 0x17, 0xbb, 0xe8, 0xb6, 0xca, 0x02, 0x54, 0x2b,
	0x01, 0x6a, 0x44, 0x8b, 0xde, 0x13, 0xea, 0xd0, 0x3b, 0x91, 0x32, 0x55, 0x0a, 0x4e, 0x0f, 0xba,
	0x6e, 0xa3, 0x9a, 0xa8, 0xc3, 0xfe, 0x4d, 0x2f, 0x12, 0x41, 0xae, 0x21, 0x2f, 0x51, 0x81, 0x7c,
	0x60, 0x32, 0x07, 0x65, 0x22, 0x7d, 0x35, 0x88, 0x0c, 0x37, 0x6d, 0xe4, 0x92, 0xec, 0xaf, 0xea,
	0x43, 0xb3, 0xed, 0x4e, 0x13, 0x7d, 0x88, 0x61, 0xbd, 0xb3, 0x5f, 0xba, 0xcd, 0xb8, 0x22, 0x43,
	0x5d, 0x5c, 0x93, 0x44, 0x27, 0xad, 0xd3, 0x6f, 0x38, 0x0a, 0xd4, 0x2e, 0xad, 0x5f, 0xd8, 0xa7,
	0xfd, 0x92, 0xa3, 0x40, 0xb5, 0xf4, 0x13, 0x99, 0x75, 0x2b, 0x69, 0x32, 0xda, 0x5e, 0x8e, 0x5a,
	0xc6, 0xec, 0xd5, 0x3e, 0xbf, 0x9c, 0xe3, 0x3f, 0x1c, 0x66, 0xc2, 0x72, 0xf8, 0xe8, 0x7e, 0xa6,
	0xe7, 0xbd, 0xe6, 0xdb, 0x3e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x74, 0x88, 0xe2, 0xc2, 0x77,
	0x03, 0x00, 0x00,
}
