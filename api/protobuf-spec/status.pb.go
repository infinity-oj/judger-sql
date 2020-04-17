// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

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

type Status int32

const (
	Status_success Status = 0
	Status_error   Status = 1
)

var Status_name = map[int32]string{
	0: "success",
	1: "error",
}

var Status_value = map[string]int32{
	"success": 0,
	"error":   1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}

func init() {
	proto.RegisterEnum("Status", Status_name, Status_value)
}

func init() {
	proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b)
}

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 74 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x52, 0xe0, 0x62, 0x0b, 0x06, 0xf3, 0x85,
	0xb8, 0xb9, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x05, 0x18, 0x84, 0x38, 0xb9, 0x58,
	0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0x04, 0x18, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0x90, 0x7d, 0x86, 0x38, 0x00, 0x00, 0x00,
}
