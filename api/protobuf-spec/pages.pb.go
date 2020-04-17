// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pages.proto

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

type CreatePageRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Locale               string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePageRequest) Reset()         { *m = CreatePageRequest{} }
func (m *CreatePageRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePageRequest) ProtoMessage()    {}
func (*CreatePageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20be0c8f23cb364f, []int{0}
}

func (m *CreatePageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePageRequest.Unmarshal(m, b)
}
func (m *CreatePageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePageRequest.Marshal(b, m, deterministic)
}
func (m *CreatePageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePageRequest.Merge(m, src)
}
func (m *CreatePageRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePageRequest.Size(m)
}
func (m *CreatePageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePageRequest proto.InternalMessageInfo

func (m *CreatePageRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePageRequest) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type CreatePageResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePageResponse) Reset()         { *m = CreatePageResponse{} }
func (m *CreatePageResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePageResponse) ProtoMessage()    {}
func (*CreatePageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20be0c8f23cb364f, []int{1}
}

func (m *CreatePageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePageResponse.Unmarshal(m, b)
}
func (m *CreatePageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePageResponse.Marshal(b, m, deterministic)
}
func (m *CreatePageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePageResponse.Merge(m, src)
}
func (m *CreatePageResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePageResponse.Size(m)
}
func (m *CreatePageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePageResponse proto.InternalMessageInfo

func (m *CreatePageResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_success
}

type FetchPageRequest struct {
	PageId               string   `protobuf:"bytes,1,opt,name=Page_id,json=PageId,proto3" json:"Page_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchPageRequest) Reset()         { *m = FetchPageRequest{} }
func (m *FetchPageRequest) String() string { return proto.CompactTextString(m) }
func (*FetchPageRequest) ProtoMessage()    {}
func (*FetchPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20be0c8f23cb364f, []int{2}
}

func (m *FetchPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchPageRequest.Unmarshal(m, b)
}
func (m *FetchPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchPageRequest.Marshal(b, m, deterministic)
}
func (m *FetchPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPageRequest.Merge(m, src)
}
func (m *FetchPageRequest) XXX_Size() int {
	return xxx_messageInfo_FetchPageRequest.Size(m)
}
func (m *FetchPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPageRequest proto.InternalMessageInfo

func (m *FetchPageRequest) GetPageId() string {
	if m != nil {
		return m.PageId
	}
	return ""
}

type FetchPageResponse struct {
	Page                 *Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchPageResponse) Reset()         { *m = FetchPageResponse{} }
func (m *FetchPageResponse) String() string { return proto.CompactTextString(m) }
func (*FetchPageResponse) ProtoMessage()    {}
func (*FetchPageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20be0c8f23cb364f, []int{3}
}

func (m *FetchPageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchPageResponse.Unmarshal(m, b)
}
func (m *FetchPageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchPageResponse.Marshal(b, m, deterministic)
}
func (m *FetchPageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPageResponse.Merge(m, src)
}
func (m *FetchPageResponse) XXX_Size() int {
	return xxx_messageInfo_FetchPageResponse.Size(m)
}
func (m *FetchPageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPageResponse proto.InternalMessageInfo

func (m *FetchPageResponse) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

type Page struct {
	Pid                  uint64   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TimeLimit            uint64   `protobuf:"varint,3,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
	MemoryLimit          uint64   `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	InputFormat          string   `protobuf:"bytes,6,opt,name=input_format,json=inputFormat,proto3" json:"input_format,omitempty"`
	OutputFormat         string   `protobuf:"bytes,7,opt,name=output_format,json=outputFormat,proto3" json:"output_format,omitempty"`
	Example              string   `protobuf:"bytes,8,opt,name=example,proto3" json:"example,omitempty"`
	HintAndLimit         string   `protobuf:"bytes,9,opt,name=hint_and_limit,json=hintAndLimit,proto3" json:"hint_and_limit,omitempty"`
	PublicizerId         uint64   `protobuf:"varint,10,opt,name=publicizer_id,json=publicizerId,proto3" json:"publicizer_id,omitempty"`
	FileIoInputName      string   `protobuf:"bytes,11,opt,name=file_io_input_name,json=fileIoInputName,proto3" json:"file_io_input_name,omitempty"`
	FileIoOutputName     string   `protobuf:"bytes,12,opt,name=file_io_output_name,json=fileIoOutputName,proto3" json:"file_io_output_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_20be0c8f23cb364f, []int{4}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Page) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Page) GetTimeLimit() uint64 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *Page) GetMemoryLimit() uint64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *Page) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Page) GetInputFormat() string {
	if m != nil {
		return m.InputFormat
	}
	return ""
}

func (m *Page) GetOutputFormat() string {
	if m != nil {
		return m.OutputFormat
	}
	return ""
}

func (m *Page) GetExample() string {
	if m != nil {
		return m.Example
	}
	return ""
}

func (m *Page) GetHintAndLimit() string {
	if m != nil {
		return m.HintAndLimit
	}
	return ""
}

func (m *Page) GetPublicizerId() uint64 {
	if m != nil {
		return m.PublicizerId
	}
	return 0
}

func (m *Page) GetFileIoInputName() string {
	if m != nil {
		return m.FileIoInputName
	}
	return ""
}

func (m *Page) GetFileIoOutputName() string {
	if m != nil {
		return m.FileIoOutputName
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePageRequest)(nil), "CreatePageRequest")
	proto.RegisterType((*CreatePageResponse)(nil), "CreatePageResponse")
	proto.RegisterType((*FetchPageRequest)(nil), "FetchPageRequest")
	proto.RegisterType((*FetchPageResponse)(nil), "FetchPageResponse")
	proto.RegisterType((*Page)(nil), "Page")
}

func init() { proto.RegisterFile("pages.proto", fileDescriptor_20be0c8f23cb364f) }

var fileDescriptor_20be0c8f23cb364f = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x69, 0x77, 0x93, 0xb8, 0x27, 0x69, 0xdd, 0x3d, 0x15, 0x8d, 0x0b, 0x62, 0x4d, 0xbd,
	0x10, 0x8a, 0xb9, 0x58, 0xe9, 0x03, 0x54, 0xa1, 0xb0, 0x20, 0x2a, 0xeb, 0x9d, 0x37, 0x61, 0x9a,
	0x9c, 0xb6, 0x03, 0x99, 0xcc, 0x98, 0x4c, 0x40, 0x7d, 0x15, 0x5f, 0x56, 0xe6, 0xcc, 0xd4, 0x0d,
	0xd6, 0xab, 0xdd, 0xef, 0x3b, 0xbf, 0x9c, 0x3f, 0xc9, 0x07, 0xa9, 0x11, 0xb7, 0x34, 0x94, 0xa6,
	0xd7, 0x56, 0xaf, 0xb3, 0xc1, 0x0a, 0x3b, 0x06, 0x55, 0x5c, 0xc2, 0xea, 0x43, 0x4f, 0xc2, 0xd2,
	0x17, 0x71, 0x4b, 0x3b, 0xfa, 0x3e, 0xd2, 0x60, 0xf1, 0x09, 0x44, 0x56, 0xda, 0x96, 0xf2, 0x83,
	0xd3, 0x83, 0x37, 0x8b, 0x9d, 0x17, 0xf8, 0x14, 0xe2, 0x56, 0xd7, 0xa2, 0xa5, 0xfc, 0x90, 0xed,
	0xa0, 0x8a, 0x0b, 0xc0, 0x69, 0x8b, 0xc1, 0xe8, 0x6e, 0x20, 0x7c, 0x09, 0xb1, 0x1f, 0xc4, 0x4d,
	0x8e, 0x37, 0x49, 0xf9, 0x95, 0xe5, 0x2e, 0xd8, 0xc5, 0x39, 0x2c, 0xaf, 0xc8, 0xd6, 0x77, 0xd3,
	0xc1, 0xcf, 0x20, 0x71, 0xb2, 0x92, 0x4d, 0x18, 0x1d, 0x3b, 0xb9, 0x6d, 0x8a, 0x12, 0x56, 0x13,
	0x38, 0x8c, 0x78, 0x0e, 0x73, 0x77, 0x18, 0xa3, 0xe9, 0x26, 0x2a, 0xb9, 0xc8, 0x56, 0xf1, 0x7b,
	0x06, 0x73, 0x27, 0x71, 0x09, 0x33, 0x13, 0xba, 0xcd, 0x77, 0xee, 0xef, 0xfe, 0xb8, 0xc3, 0xe9,
	0x71, 0x2f, 0x00, 0xac, 0x54, 0x54, 0xb5, 0x52, 0x49, 0x9b, 0xcf, 0x18, 0x5f, 0x38, 0xe7, 0xa3,
	0x33, 0xf0, 0x15, 0x64, 0x8a, 0x94, 0xee, 0x7f, 0x06, 0x60, 0xce, 0x40, 0xea, 0x3d, 0x8f, 0x9c,
	0x42, 0xda, 0xd0, 0x50, 0xf7, 0xd2, 0x58, 0xa9, 0xbb, 0x3c, 0xe2, 0xee, 0x53, 0xcb, 0x35, 0x91,
	0x9d, 0x19, 0x6d, 0x75, 0xa3, 0x7b, 0x25, 0x6c, 0x1e, 0x7b, 0x84, 0xbd, 0x2b, 0xb6, 0xf0, 0x0c,
	0x8e, 0xf4, 0x68, 0x27, 0x4c, 0xc2, 0x4c, 0xe6, 0xcd, 0x00, 0xe5, 0x90, 0xd0, 0x0f, 0xa1, 0x4c,
	0x4b, 0xf9, 0x23, 0x2e, 0xdf, 0x4b, 0x7c, 0x0d, 0xc7, 0x77, 0xb2, 0xb3, 0x95, 0xe8, 0x9a, 0xb0,
	0xe8, 0xc2, 0x3f, 0xef, 0xdc, 0xcb, 0xae, 0xf1, 0x9b, 0x9e, 0xc1, 0x91, 0x19, 0xaf, 0x5b, 0x59,
	0xcb, 0x5f, 0xd4, 0xbb, 0x77, 0x0d, 0x7c, 0x4d, 0xb6, 0x37, 0xb7, 0x0d, 0x9e, 0x03, 0xde, 0xc8,
	0x96, 0x2a, 0xa9, 0x2b, 0xbf, 0x74, 0x27, 0x14, 0xe5, 0x29, 0xb7, 0x7b, 0xec, 0x2a, 0x5b, 0xbd,
	0x75, 0xfe, 0x27, 0xa1, 0x08, 0xdf, 0xc2, 0xc9, 0x3d, 0x1c, 0xd6, 0x67, 0x3a, 0x63, 0x7a, 0xe9,
	0xe9, 0xcf, 0x5c, 0x70, 0xf8, 0xa6, 0x87, 0xc8, 0x7d, 0x9c, 0x01, 0x2f, 0x00, 0xf6, 0xd1, 0x41,
	0x2c, 0x1f, 0x44, 0x71, 0x7d, 0x52, 0xfe, 0x27, 0x5b, 0x1b, 0x58, 0xfc, 0x4d, 0x03, 0xae, 0xca,
	0x7f, 0x63, 0xb4, 0xc6, 0xf2, 0x41, 0x58, 0xde, 0x27, 0xdf, 0x22, 0x4e, 0xfc, 0x75, 0xcc, 0x3f,
	0xef, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xae, 0x0a, 0x84, 0x15, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PagesClient is the client API for Pages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PagesClient interface {
	CreatePage(ctx context.Context, in *CreatePageRequest, opts ...grpc.CallOption) (*CreatePageResponse, error)
	FetchPage(ctx context.Context, in *FetchPageRequest, opts ...grpc.CallOption) (*FetchPageResponse, error)
}

type pagesClient struct {
	cc grpc.ClientConnInterface
}

func NewPagesClient(cc grpc.ClientConnInterface) PagesClient {
	return &pagesClient{cc}
}

func (c *pagesClient) CreatePage(ctx context.Context, in *CreatePageRequest, opts ...grpc.CallOption) (*CreatePageResponse, error) {
	out := new(CreatePageResponse)
	err := c.cc.Invoke(ctx, "/Pages/CreatePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagesClient) FetchPage(ctx context.Context, in *FetchPageRequest, opts ...grpc.CallOption) (*FetchPageResponse, error) {
	out := new(FetchPageResponse)
	err := c.cc.Invoke(ctx, "/Pages/FetchPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PagesServer is the server API for Pages service.
type PagesServer interface {
	CreatePage(context.Context, *CreatePageRequest) (*CreatePageResponse, error)
	FetchPage(context.Context, *FetchPageRequest) (*FetchPageResponse, error)
}

// UnimplementedPagesServer can be embedded to have forward compatible implementations.
type UnimplementedPagesServer struct {
}

func (*UnimplementedPagesServer) CreatePage(ctx context.Context, req *CreatePageRequest) (*CreatePageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (*UnimplementedPagesServer) FetchPage(ctx context.Context, req *FetchPageRequest) (*FetchPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPage not implemented")
}

func RegisterPagesServer(s *grpc.Server, srv PagesServer) {
	s.RegisterService(&_Pages_serviceDesc, srv)
}

func _Pages_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pages/CreatePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).CreatePage(ctx, req.(*CreatePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pages_FetchPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).FetchPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pages/FetchPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).FetchPage(ctx, req.(*FetchPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Pages",
	HandlerType: (*PagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePage",
			Handler:    _Pages_CreatePage_Handler,
		},
		{
			MethodName: "FetchPage",
			Handler:    _Pages_FetchPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pages.proto",
}
