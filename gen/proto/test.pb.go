// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto

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

type ResponseRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseRequest) Reset()         { *m = ResponseRequest{} }
func (m *ResponseRequest) String() string { return proto.CompactTextString(m) }
func (*ResponseRequest) ProtoMessage()    {}
func (*ResponseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *ResponseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseRequest.Unmarshal(m, b)
}
func (m *ResponseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseRequest.Marshal(b, m, deterministic)
}
func (m *ResponseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseRequest.Merge(m, src)
}
func (m *ResponseRequest) XXX_Size() int {
	return xxx_messageInfo_ResponseRequest.Size(m)
}
func (m *ResponseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseRequest proto.InternalMessageInfo

func (m *ResponseRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UserRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseRequest)(nil), "main.ResponseRequest")
	proto.RegisterType((*UserRequest)(nil), "main.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "main.UserResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0x80, 0x30,
	0x14, 0xc6, 0x35, 0x35, 0xf1, 0x15, 0x54, 0x8f, 0x02, 0xf1, 0x64, 0xeb, 0xe2, 0x69, 0x81, 0x41,
	0xf7, 0x82, 0x08, 0x3a, 0x8e, 0xba, 0x74, 0x5b, 0xf5, 0xb0, 0x41, 0x73, 0xe6, 0xb6, 0xff, 0x3f,
	0xb6, 0x25, 0x48, 0x74, 0xda, 0x6f, 0x8f, 0xdf, 0x3e, 0xbe, 0x3d, 0x00, 0x47, 0xd6, 0xf1, 0x65,
	0x35, 0xce, 0x60, 0xa9, 0xa5, 0x9a, 0xd9, 0x15, 0x9c, 0x08, 0xb2, 0x8b, 0x99, 0x2d, 0x09, 0xfa,
	0xf6, 0x64, 0x1d, 0x9e, 0x42, 0xa1, 0xed, 0xd4, 0xe6, 0x7d, 0x3e, 0x34, 0x22, 0x20, 0xbb, 0x84,
	0xa3, 0x17, 0x4b, 0xeb, 0x26, 0x20, 0x94, 0xde, 0xab, 0x8f, 0x5f, 0x23, 0x32, 0x7b, 0x82, 0xe3,
	0xa4, 0xa4, 0xac, 0xe0, 0xcc, 0x52, 0xd3, 0xe6, 0x04, 0x0e, 0xc1, 0x72, 0xa2, 0xf6, 0xa0, 0xcf,
	0x87, 0x4a, 0x04, 0xc4, 0x73, 0xa8, 0x48, 0x4b, 0xf5, 0xd5, 0x16, 0x51, 0x4b, 0x97, 0xd1, 0x43,
	0xfd, 0x4c, 0xd6, 0xdd, 0x2d, 0x0a, 0x6f, 0xa1, 0x7c, 0x78, 0xff, 0x34, 0x78, 0xc1, 0x43, 0x5b,
	0xfe, 0xa7, 0x6a, 0xf7, 0xff, 0x98, 0x65, 0x38, 0x42, 0xfd, 0x48, 0x2e, 0x34, 0xc2, 0xb3, 0xe4,
	0xec, 0x3e, 0xd0, 0xe1, 0x7e, 0x94, 0x9e, 0xb2, 0xec, 0xbe, 0x79, 0xad, 0xf9, 0x75, 0xdc, 0xcd,
	0xdb, 0x61, 0x3c, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x9a, 0x37, 0xd6, 0x30, 0x01,
	0x00, 0x00,
}