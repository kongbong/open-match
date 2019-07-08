// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/matchfunction.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RunRequest struct {
	// The MatchProfile that describes the Match that this MatchFunction needs to
	// generate proposals for.
	Profile              *MatchProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b5069a21f149a55, []int{0}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetProfile() *MatchProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type RunResponse struct {
	// The proposal generated by this MatchFunction Run.
	// Note that OpenMatch will validate the proposals, a valid match should contain at least one ticket.
	Proposals            []*Match `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b5069a21f149a55, []int{1}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetProposals() []*Match {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func init() {
	proto.RegisterType((*RunRequest)(nil), "api.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "api.RunResponse")
}

func init() { proto.RegisterFile("api/matchfunction.proto", fileDescriptor_2b5069a21f149a55) }

var fileDescriptor_2b5069a21f149a55 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x1d, 0xd4, 0x8a, 0xad, 0x80, 0xb2, 0x12, 0x10, 0x45, 0x1c, 0x96, 0x70, 0x89, 0x02,
	0xf1, 0xa6, 0x26, 0x12, 0x22, 0x08, 0xa9, 0x05, 0x8a, 0x54, 0x54, 0x7e, 0x64, 0x24, 0x0e, 0x70,
	0xda, 0x38, 0x13, 0x7b, 0x21, 0xd9, 0x59, 0x76, 0x76, 0x5b, 0xce, 0x3c, 0x02, 0xdc, 0x78, 0x12,
	0xde, 0x83, 0x07, 0xe0, 0xc2, 0x8d, 0x97, 0x40, 0x59, 0x13, 0x19, 0x54, 0x4e, 0x96, 0xbf, 0x9f,
	0xf9, 0xc6, 0x9f, 0x87, 0x5d, 0x53, 0x56, 0xcb, 0x95, 0xf2, 0x65, 0xbd, 0x08, 0xa6, 0xf4, 0x1a,
	0x4d, 0x66, 0x1d, 0x7a, 0xe4, 0x1d, 0x65, 0x75, 0x8f, 0x47, 0x16, 0x88, 0x54, 0x05, 0xd4, 0x10,
	0xbd, 0xeb, 0x15, 0x62, 0xb5, 0x04, 0xb9, 0xa6, 0x94, 0x31, 0xe8, 0xd5, 0xda, 0xb5, 0x61, 0x6f,
	0xc7, 0x47, 0x39, 0xaa, 0xc0, 0x8c, 0xe8, 0x54, 0x55, 0x15, 0x38, 0x89, 0x36, 0x2a, 0xce, 0xaa,
	0xfb, 0xf7, 0x18, 0x2b, 0x82, 0x29, 0xe0, 0x43, 0x00, 0xf2, 0xfc, 0x16, 0xdb, 0xb6, 0x0e, 0x17,
	0x7a, 0x09, 0xdd, 0x44, 0x24, 0x83, 0x9d, 0xfc, 0x72, 0xa6, 0xac, 0xce, 0x9e, 0xad, 0xb7, 0x7b,
	0xd9, 0x10, 0xc5, 0x46, 0xd1, 0xbf, 0xcb, 0x76, 0xa2, 0x95, 0x2c, 0x1a, 0x02, 0x3e, 0x60, 0xe7,
	0xad, 0x43, 0x8b, 0xa4, 0x96, 0xd4, 0x4d, 0x44, 0x67, 0xb0, 0x93, 0xb3, 0xd6, 0x5d, 0xb4, 0x64,
	0xfe, 0x96, 0x5d, 0x88, 0xd8, 0x93, 0x3f, 0xdf, 0xcb, 0x9f, 0xb2, 0x4e, 0x11, 0x0c, 0xbf, 0x14,
	0xe5, 0xed, 0x3a, 0xbd, 0xdd, 0x16, 0x68, 0x42, 0xfa, 0xe2, 0xd3, 0xf7, 0x9f, 0x5f, 0xd2, 0x5e,
	0xff, 0x8a, 0x3c, 0xd9, 0xfb, 0xb7, 0xb4, 0xa9, 0x0b, 0x66, 0x9a, 0x0c, 0x1f, 0xfe, 0x4a, 0x3f,
	0x1f, 0xfc, 0x48, 0xf9, 0xb7, 0x84, 0x5d, 0x8c, 0x21, 0x62, 0x93, 0xd2, 0x3f, 0x62, 0xec, 0x85,
	0x05, 0x23, 0x22, 0xcc, 0xaf, 0xd6, 0xde, 0x5b, 0x9a, 0x4a, 0x89, 0x16, 0xcc, 0x28, 0x8e, 0xca,
	0xe6, 0x70, 0xd2, 0xbb, 0xd9, 0xbe, 0x8f, 0xe6, 0x9a, 0xca, 0x40, 0xb4, 0xdf, 0xf4, 0x5d, 0x39,
	0x0c, 0x96, 0xb2, 0x12, 0x57, 0xc3, 0xd7, 0x8c, 0x1f, 0x58, 0x55, 0xd6, 0x20, 0xf2, 0x6c, 0x2c,
	0x8e, 0x75, 0x09, 0xeb, 0x06, 0xf6, 0x37, 0x23, 0x2b, 0xed, 0xeb, 0x30, 0x5b, 0x2b, 0x65, 0x63,
	0x5d, 0xa0, 0xab, 0xd4, 0x0a, 0xe8, 0xaf, 0x30, 0x39, 0x5b, 0xe2, 0x4c, 0xae, 0x14, 0x79, 0x70,
	0xf2, 0xf8, 0xe8, 0xd1, 0xe1, 0xf3, 0x57, 0x87, 0x79, 0x67, 0x2f, 0x1b, 0x0f, 0xd3, 0x24, 0xcd,
	0x77, 0x95, 0xb5, 0x4b, 0x5d, 0xc6, 0x5f, 0x25, 0xdf, 0x11, 0x9a, 0xe9, 0x19, 0xa4, 0xb8, 0xcf,
	0x3a, 0x93, 0xf1, 0x84, 0x4f, 0xd8, 0xb0, 0x00, 0x1f, 0x9c, 0x81, 0xb9, 0x38, 0xad, 0xc1, 0x08,
	0x5f, 0x83, 0x70, 0x40, 0x18, 0x5c, 0x09, 0x62, 0x8e, 0x40, 0xc2, 0xa0, 0x17, 0xf0, 0x51, 0x93,
	0xcf, 0xf8, 0x16, 0x3b, 0xf7, 0x35, 0x4d, 0xb6, 0xdd, 0x03, 0xd6, 0x6d, 0xcb, 0x10, 0x8f, 0xb1,
	0x0c, 0x2b, 0x30, 0xcd, 0x69, 0xf0, 0x1b, 0xff, 0xaf, 0x46, 0x92, 0xf6, 0x20, 0xe7, 0x58, 0x92,
	0x7c, 0xb3, 0x65, 0xdf, 0x57, 0xd2, 0xce, 0x66, 0x5b, 0xf1, 0x8a, 0xee, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xed, 0x4f, 0xf9, 0xaa, 0xc5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MatchFunctionClient is the client API for MatchFunction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatchFunctionClient interface {
	// This is the function that is executed when by the Open Match backend to
	// generate Match proposals.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type matchFunctionClient struct {
	cc *grpc.ClientConn
}

func NewMatchFunctionClient(cc *grpc.ClientConn) MatchFunctionClient {
	return &matchFunctionClient{cc}
}

func (c *matchFunctionClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/api.MatchFunction/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchFunctionServer is the server API for MatchFunction service.
type MatchFunctionServer interface {
	// This is the function that is executed when by the Open Match backend to
	// generate Match proposals.
	Run(context.Context, *RunRequest) (*RunResponse, error)
}

func RegisterMatchFunctionServer(s *grpc.Server, srv MatchFunctionServer) {
	s.RegisterService(&_MatchFunction_serviceDesc, srv)
}

func _MatchFunction_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchFunctionServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MatchFunction/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchFunctionServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MatchFunction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MatchFunction",
	HandlerType: (*MatchFunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _MatchFunction_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/matchfunction.proto",
}