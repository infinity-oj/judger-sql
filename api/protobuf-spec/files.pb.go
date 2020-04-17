// Code generated by protoc-gen-go. DO NOT EDIT.
// source: files.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FileMeta struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Length               string   `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	Sha256               string   `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMeta) Reset()         { *m = FileMeta{} }
func (m *FileMeta) String() string { return proto.CompactTextString(m) }
func (*FileMeta) ProtoMessage()    {}
func (*FileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{0}
}

func (m *FileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMeta.Unmarshal(m, b)
}
func (m *FileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMeta.Marshal(b, m, deterministic)
}
func (m *FileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMeta.Merge(m, src)
}
func (m *FileMeta) XXX_Size() int {
	return xxx_messageInfo_FileMeta.Size(m)
}
func (m *FileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FileMeta proto.InternalMessageInfo

func (m *FileMeta) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileMeta) GetLength() string {
	if m != nil {
		return m.Length
	}
	return ""
}

func (m *FileMeta) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

type File struct {
	Filename             string    `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Meta                 *FileMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{1}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *File) GetMeta() *FileMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Directory struct {
	DirectoryName        string   `protobuf:"bytes,1,opt,name=directory_name,json=directoryName,proto3" json:"directory_name,omitempty"`
	File                 []string `protobuf:"bytes,2,rep,name=File,proto3" json:"File,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Directory) Reset()         { *m = Directory{} }
func (m *Directory) String() string { return proto.CompactTextString(m) }
func (*Directory) ProtoMessage()    {}
func (*Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{2}
}

func (m *Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Directory.Unmarshal(m, b)
}
func (m *Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Directory.Marshal(b, m, deterministic)
}
func (m *Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Directory.Merge(m, src)
}
func (m *Directory) XXX_Size() int {
	return xxx_messageInfo_Directory.Size(m)
}
func (m *Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Directory proto.InternalMessageInfo

func (m *Directory) GetDirectoryName() string {
	if m != nil {
		return m.DirectoryName
	}
	return ""
}

func (m *Directory) GetFile() []string {
	if m != nil {
		return m.File
	}
	return nil
}

// CreateFileSpace
type CreateFileSpaceRequest struct {
	SpaceName            string   `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileSpaceRequest) Reset()         { *m = CreateFileSpaceRequest{} }
func (m *CreateFileSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFileSpaceRequest) ProtoMessage()    {}
func (*CreateFileSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{3}
}

func (m *CreateFileSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileSpaceRequest.Unmarshal(m, b)
}
func (m *CreateFileSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileSpaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateFileSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileSpaceRequest.Merge(m, src)
}
func (m *CreateFileSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFileSpaceRequest.Size(m)
}
func (m *CreateFileSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileSpaceRequest proto.InternalMessageInfo

func (m *CreateFileSpaceRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

type CreateFileSpaceResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileSpaceResponse) Reset()         { *m = CreateFileSpaceResponse{} }
func (m *CreateFileSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFileSpaceResponse) ProtoMessage()    {}
func (*CreateFileSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{4}
}

func (m *CreateFileSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileSpaceResponse.Unmarshal(m, b)
}
func (m *CreateFileSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileSpaceResponse.Marshal(b, m, deterministic)
}
func (m *CreateFileSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileSpaceResponse.Merge(m, src)
}
func (m *CreateFileSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFileSpaceResponse.Size(m)
}
func (m *CreateFileSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileSpaceResponse proto.InternalMessageInfo

func (m *CreateFileSpaceResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_success
}

// CreateDirectory
type CreateDirectoryRequest struct {
	FileSpace            string   `protobuf:"bytes,1,opt,name=file_space,json=fileSpace,proto3" json:"file_space,omitempty"`
	Directory            string   `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDirectoryRequest) Reset()         { *m = CreateDirectoryRequest{} }
func (m *CreateDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDirectoryRequest) ProtoMessage()    {}
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{5}
}

func (m *CreateDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDirectoryRequest.Unmarshal(m, b)
}
func (m *CreateDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDirectoryRequest.Merge(m, src)
}
func (m *CreateDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDirectoryRequest.Size(m)
}
func (m *CreateDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDirectoryRequest proto.InternalMessageInfo

func (m *CreateDirectoryRequest) GetFileSpace() string {
	if m != nil {
		return m.FileSpace
	}
	return ""
}

func (m *CreateDirectoryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type CreateDirectoryResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDirectoryResponse) Reset()         { *m = CreateDirectoryResponse{} }
func (m *CreateDirectoryResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDirectoryResponse) ProtoMessage()    {}
func (*CreateDirectoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{6}
}

func (m *CreateDirectoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDirectoryResponse.Unmarshal(m, b)
}
func (m *CreateDirectoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDirectoryResponse.Marshal(b, m, deterministic)
}
func (m *CreateDirectoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDirectoryResponse.Merge(m, src)
}
func (m *CreateDirectoryResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDirectoryResponse.Size(m)
}
func (m *CreateDirectoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDirectoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDirectoryResponse proto.InternalMessageInfo

func (m *CreateDirectoryResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_success
}

// CreateFile
type CreateFileRequest struct {
	FileSpace            string   `protobuf:"bytes,1,opt,name=file_space,json=fileSpace,proto3" json:"file_space,omitempty"`
	FilePath             string   `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileRequest) Reset()         { *m = CreateFileRequest{} }
func (m *CreateFileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFileRequest) ProtoMessage()    {}
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{7}
}

func (m *CreateFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileRequest.Unmarshal(m, b)
}
func (m *CreateFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileRequest.Marshal(b, m, deterministic)
}
func (m *CreateFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileRequest.Merge(m, src)
}
func (m *CreateFileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFileRequest.Size(m)
}
func (m *CreateFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileRequest proto.InternalMessageInfo

func (m *CreateFileRequest) GetFileSpace() string {
	if m != nil {
		return m.FileSpace
	}
	return ""
}

func (m *CreateFileRequest) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *CreateFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateFileResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileResponse) Reset()         { *m = CreateFileResponse{} }
func (m *CreateFileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFileResponse) ProtoMessage()    {}
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{8}
}

func (m *CreateFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileResponse.Unmarshal(m, b)
}
func (m *CreateFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileResponse.Marshal(b, m, deterministic)
}
func (m *CreateFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileResponse.Merge(m, src)
}
func (m *CreateFileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFileResponse.Size(m)
}
func (m *CreateFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileResponse proto.InternalMessageInfo

func (m *CreateFileResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_success
}

// CheckFileSpaceExists
type FetchFileSpaceRequest struct {
	SpaceName            string   `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchFileSpaceRequest) Reset()         { *m = FetchFileSpaceRequest{} }
func (m *FetchFileSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*FetchFileSpaceRequest) ProtoMessage()    {}
func (*FetchFileSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{9}
}

func (m *FetchFileSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchFileSpaceRequest.Unmarshal(m, b)
}
func (m *FetchFileSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchFileSpaceRequest.Marshal(b, m, deterministic)
}
func (m *FetchFileSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchFileSpaceRequest.Merge(m, src)
}
func (m *FetchFileSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_FetchFileSpaceRequest.Size(m)
}
func (m *FetchFileSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchFileSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchFileSpaceRequest proto.InternalMessageInfo

func (m *FetchFileSpaceRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

type FetchFileSpaceResponse struct {
	Exist                bool     `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	PrevFileSpace        string   `protobuf:"bytes,2,opt,name=prev_file_space,json=prevFileSpace,proto3" json:"prev_file_space,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchFileSpaceResponse) Reset()         { *m = FetchFileSpaceResponse{} }
func (m *FetchFileSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*FetchFileSpaceResponse) ProtoMessage()    {}
func (*FetchFileSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{10}
}

func (m *FetchFileSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchFileSpaceResponse.Unmarshal(m, b)
}
func (m *FetchFileSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchFileSpaceResponse.Marshal(b, m, deterministic)
}
func (m *FetchFileSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchFileSpaceResponse.Merge(m, src)
}
func (m *FetchFileSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_FetchFileSpaceResponse.Size(m)
}
func (m *FetchFileSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchFileSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchFileSpaceResponse proto.InternalMessageInfo

func (m *FetchFileSpaceResponse) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

func (m *FetchFileSpaceResponse) GetPrevFileSpace() string {
	if m != nil {
		return m.PrevFileSpace
	}
	return ""
}

type FetchFileRequest struct {
	SpaceName            string   `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	FilePath             string   `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchFileRequest) Reset()         { *m = FetchFileRequest{} }
func (m *FetchFileRequest) String() string { return proto.CompactTextString(m) }
func (*FetchFileRequest) ProtoMessage()    {}
func (*FetchFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{11}
}

func (m *FetchFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchFileRequest.Unmarshal(m, b)
}
func (m *FetchFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchFileRequest.Marshal(b, m, deterministic)
}
func (m *FetchFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchFileRequest.Merge(m, src)
}
func (m *FetchFileRequest) XXX_Size() int {
	return xxx_messageInfo_FetchFileRequest.Size(m)
}
func (m *FetchFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchFileRequest proto.InternalMessageInfo

func (m *FetchFileRequest) GetSpaceName() string {
	if m != nil {
		return m.SpaceName
	}
	return ""
}

func (m *FetchFileRequest) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

type FetchFileResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	File                 *File    `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchFileResponse) Reset()         { *m = FetchFileResponse{} }
func (m *FetchFileResponse) String() string { return proto.CompactTextString(m) }
func (*FetchFileResponse) ProtoMessage()    {}
func (*FetchFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac8f32ecfdd343c, []int{12}
}

func (m *FetchFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchFileResponse.Unmarshal(m, b)
}
func (m *FetchFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchFileResponse.Marshal(b, m, deterministic)
}
func (m *FetchFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchFileResponse.Merge(m, src)
}
func (m *FetchFileResponse) XXX_Size() int {
	return xxx_messageInfo_FetchFileResponse.Size(m)
}
func (m *FetchFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchFileResponse proto.InternalMessageInfo

func (m *FetchFileResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_success
}

func (m *FetchFileResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*FileMeta)(nil), "FileMeta")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*Directory)(nil), "Directory")
	proto.RegisterType((*CreateFileSpaceRequest)(nil), "CreateFileSpaceRequest")
	proto.RegisterType((*CreateFileSpaceResponse)(nil), "CreateFileSpaceResponse")
	proto.RegisterType((*CreateDirectoryRequest)(nil), "CreateDirectoryRequest")
	proto.RegisterType((*CreateDirectoryResponse)(nil), "CreateDirectoryResponse")
	proto.RegisterType((*CreateFileRequest)(nil), "CreateFileRequest")
	proto.RegisterType((*CreateFileResponse)(nil), "CreateFileResponse")
	proto.RegisterType((*FetchFileSpaceRequest)(nil), "FetchFileSpaceRequest")
	proto.RegisterType((*FetchFileSpaceResponse)(nil), "FetchFileSpaceResponse")
	proto.RegisterType((*FetchFileRequest)(nil), "FetchFileRequest")
	proto.RegisterType((*FetchFileResponse)(nil), "FetchFileResponse")
}

func init() { proto.RegisterFile("files.proto", fileDescriptor_cac8f32ecfdd343c) }

var fileDescriptor_cac8f32ecfdd343c = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xda, 0x24, 0x8d, 0xa7, 0x5f, 0x64, 0x80, 0xd4, 0x18, 0x2a, 0xaa, 0x95, 0x40, 0x3d,
	0xed, 0xc1, 0x28, 0x45, 0xe2, 0x08, 0x55, 0x6e, 0x14, 0xb4, 0x55, 0x7b, 0xe0, 0x12, 0x2d, 0xe9,
	0x84, 0x44, 0x4a, 0x13, 0xe3, 0xdd, 0x22, 0xf8, 0xf3, 0x08, 0xed, 0xd8, 0x5e, 0x5b, 0x71, 0x88,
	0xdc, 0x53, 0xf6, 0xbd, 0x64, 0xde, 0x9b, 0x99, 0x7d, 0x1b, 0xd8, 0x9f, 0xce, 0x17, 0x64, 0x64,
	0x92, 0xae, 0xec, 0x2a, 0x3a, 0x30, 0x56, 0xdb, 0x87, 0x1c, 0x89, 0x5b, 0xe8, 0x8d, 0xe6, 0x0b,
	0xfa, 0x4c, 0x56, 0x63, 0x04, 0x3d, 0xf7, 0xc3, 0xa5, 0xbe, 0xa7, 0xb0, 0x75, 0xd6, 0x3a, 0x0f,
	0x94, 0xc7, 0x38, 0x80, 0xee, 0x82, 0x96, 0x3f, 0xec, 0x2c, 0xdc, 0xe1, 0x6f, 0x72, 0xe4, 0x78,
	0x33, 0xd3, 0xf1, 0xf0, 0x22, 0xdc, 0xcd, 0xf8, 0x0c, 0x89, 0x1b, 0x68, 0x3b, 0xdd, 0xad, 0x9a,
	0xa7, 0xd0, 0xbe, 0x27, 0xab, 0x59, 0x71, 0x3f, 0x0e, 0x64, 0xd1, 0x88, 0x62, 0x1a, 0x11, 0xda,
	0x77, 0xda, 0x6a, 0x16, 0x3e, 0x50, 0x7c, 0x16, 0x23, 0x08, 0x2e, 0xe7, 0x29, 0x4d, 0xec, 0x2a,
	0xfd, 0x83, 0x6f, 0xe0, 0xe8, 0xae, 0x00, 0xe3, 0x8a, 0xc3, 0xa1, 0x67, 0xaf, 0x9c, 0x0d, 0x66,
	0xad, 0x84, 0x3b, 0x67, 0xbb, 0xe7, 0x81, 0xe2, 0xb3, 0x78, 0x0f, 0x83, 0x4f, 0x29, 0x69, 0x4b,
	0x0e, 0x5d, 0x27, 0x7a, 0x42, 0x8a, 0x7e, 0x3e, 0x90, 0xb1, 0x78, 0x0a, 0x60, 0x1c, 0xae, 0x0a,
	0x06, 0xcc, 0x38, 0x31, 0xf1, 0x01, 0x4e, 0x6a, 0x85, 0x26, 0x59, 0x2d, 0x0d, 0xe1, 0x6b, 0xe8,
	0x66, 0xab, 0xe5, 0xaa, 0xa3, 0x78, 0x4f, 0x5e, 0x33, 0x54, 0x39, 0x2d, 0x6e, 0x0a, 0x53, 0x3f,
	0x42, 0xc5, 0xd4, 0x6d, 0x65, 0xcc, 0x3e, 0x85, 0xe9, 0xb4, 0x70, 0xc0, 0x57, 0x10, 0xf8, 0x91,
	0xf2, 0xfd, 0x97, 0x44, 0xd9, 0x52, 0x45, 0xb6, 0x69, 0x4b, 0x13, 0xe8, 0x97, 0xe3, 0x34, 0xec,
	0xe6, 0x25, 0x30, 0x18, 0x27, 0xda, 0xa7, 0x81, 0xef, 0xf4, 0xab, 0xb6, 0xb3, 0x8d, 0x97, 0x36,
	0x04, 0xac, 0x9a, 0x34, 0xed, 0xed, 0x02, 0x9e, 0x8f, 0xc8, 0x4e, 0x66, 0x8f, 0xbd, 0xa2, 0x5b,
	0x18, 0xac, 0xd7, 0xe5, 0x96, 0xcf, 0xa0, 0x43, 0xbf, 0xe7, 0xc6, 0x72, 0x4d, 0x4f, 0x65, 0x00,
	0xdf, 0xc2, 0x71, 0x92, 0xd2, 0xaf, 0x71, 0x65, 0xe6, 0x6c, 0xaa, 0x43, 0x47, 0x7b, 0x15, 0x71,
	0x05, 0x4f, 0xbc, 0x6e, 0xb3, 0x56, 0xb6, 0xae, 0x4a, 0x7c, 0x81, 0x7e, 0x45, 0xaf, 0xe1, 0x56,
	0xf0, 0x05, 0xb4, 0xa7, 0x59, 0x9a, 0xdd, 0xa3, 0xe9, 0xf0, 0xa3, 0x51, 0x4c, 0xc5, 0x7f, 0x5b,
	0xd0, 0x71, 0xd0, 0xe0, 0x25, 0x1c, 0xaf, 0xa5, 0x14, 0x4f, 0xe4, 0xe6, 0xc0, 0x47, 0xa1, 0xfc,
	0x5f, 0xa0, 0xbd, 0x4a, 0xf9, 0xe4, 0x0a, 0x95, 0xf5, 0x04, 0x7b, 0x95, 0x7a, 0x06, 0x87, 0x00,
	0xa5, 0x01, 0xa2, 0xac, 0xe5, 0x2d, 0x7a, 0x2a, 0x37, 0xc4, 0x23, 0x86, 0xc0, 0x6f, 0x07, 0xfb,
	0x72, 0x7d, 0xf3, 0x11, 0xca, 0xda, 0xf2, 0x3e, 0xee, 0x7d, 0xeb, 0xf0, 0xbf, 0xda, 0xf7, 0x2e,
	0x7f, 0xbc, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x50, 0x08, 0x91, 0xf9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilesClient interface {
	CreateFileSpace(ctx context.Context, in *CreateFileSpaceRequest, opts ...grpc.CallOption) (*CreateFileSpaceResponse, error)
	CreateDirectory(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*CreateDirectoryResponse, error)
	CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error)
	FetchFile(ctx context.Context, in *FetchFileRequest, opts ...grpc.CallOption) (*FetchFileResponse, error)
}

type filesClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesClient(cc grpc.ClientConnInterface) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) CreateFileSpace(ctx context.Context, in *CreateFileSpaceRequest, opts ...grpc.CallOption) (*CreateFileSpaceResponse, error) {
	out := new(CreateFileSpaceResponse)
	err := c.cc.Invoke(ctx, "/Files/CreateFileSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) CreateDirectory(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*CreateDirectoryResponse, error) {
	out := new(CreateDirectoryResponse)
	err := c.cc.Invoke(ctx, "/Files/CreateDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error) {
	out := new(CreateFileResponse)
	err := c.cc.Invoke(ctx, "/Files/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) FetchFile(ctx context.Context, in *FetchFileRequest, opts ...grpc.CallOption) (*FetchFileResponse, error) {
	out := new(FetchFileResponse)
	err := c.cc.Invoke(ctx, "/Files/FetchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesServer is the server API for Files service.
type FilesServer interface {
	CreateFileSpace(context.Context, *CreateFileSpaceRequest) (*CreateFileSpaceResponse, error)
	CreateDirectory(context.Context, *CreateDirectoryRequest) (*CreateDirectoryResponse, error)
	CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error)
	FetchFile(context.Context, *FetchFileRequest) (*FetchFileResponse, error)
}

// UnimplementedFilesServer can be embedded to have forward compatible implementations.
type UnimplementedFilesServer struct {
}

func (*UnimplementedFilesServer) CreateFileSpace(ctx context.Context, req *CreateFileSpaceRequest) (*CreateFileSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileSpace not implemented")
}
func (*UnimplementedFilesServer) CreateDirectory(ctx context.Context, req *CreateDirectoryRequest) (*CreateDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDirectory not implemented")
}
func (*UnimplementedFilesServer) CreateFile(ctx context.Context, req *CreateFileRequest) (*CreateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (*UnimplementedFilesServer) FetchFile(ctx context.Context, req *FetchFileRequest) (*FetchFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFile not implemented")
}

func RegisterFilesServer(s *grpc.Server, srv FilesServer) {
	s.RegisterService(&_Files_serviceDesc, srv)
}

func _Files_CreateFileSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).CreateFileSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Files/CreateFileSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).CreateFileSpace(ctx, req.(*CreateFileSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_CreateDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).CreateDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Files/CreateDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).CreateDirectory(ctx, req.(*CreateDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Files/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).CreateFile(ctx, req.(*CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_FetchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).FetchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Files/FetchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).FetchFile(ctx, req.(*FetchFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Files_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Files",
	HandlerType: (*FilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileSpace",
			Handler:    _Files_CreateFileSpace_Handler,
		},
		{
			MethodName: "CreateDirectory",
			Handler:    _Files_CreateDirectory_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _Files_CreateFile_Handler,
		},
		{
			MethodName: "FetchFile",
			Handler:    _Files_FetchFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "files.proto",
}
