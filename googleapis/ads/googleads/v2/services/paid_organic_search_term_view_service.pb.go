// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/paid_organic_search_term_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [PaidOrganicSearchTermViewService.GetPaidOrganicSearchTermView][google.ads.googleads.v2.services.PaidOrganicSearchTermViewService.GetPaidOrganicSearchTermView].
type GetPaidOrganicSearchTermViewRequest struct {
	// The resource name of the paid organic search term view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaidOrganicSearchTermViewRequest) Reset()         { *m = GetPaidOrganicSearchTermViewRequest{} }
func (m *GetPaidOrganicSearchTermViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetPaidOrganicSearchTermViewRequest) ProtoMessage()    {}
func (*GetPaidOrganicSearchTermViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056f29523bf8aaf, []int{0}
}

func (m *GetPaidOrganicSearchTermViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Unmarshal(m, b)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Marshal(b, m, deterministic)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Merge(m, src)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.Size(m)
}
func (m *GetPaidOrganicSearchTermViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaidOrganicSearchTermViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaidOrganicSearchTermViewRequest proto.InternalMessageInfo

func (m *GetPaidOrganicSearchTermViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPaidOrganicSearchTermViewRequest)(nil), "google.ads.googleads.v2.services.GetPaidOrganicSearchTermViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/paid_organic_search_term_view_service.proto", fileDescriptor_0056f29523bf8aaf)
}

var fileDescriptor_0056f29523bf8aaf = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x8a, 0xd5, 0x40,
	0x14, 0xc6, 0x49, 0x04, 0xc1, 0xa0, 0x4d, 0x1a, 0x97, 0xb8, 0x45, 0xd8, 0x5d, 0x41, 0xb6, 0x98,
	0x81, 0xd8, 0x8d, 0xab, 0x90, 0xc5, 0xe5, 0x8a, 0x88, 0x5e, 0x76, 0x25, 0x85, 0x04, 0xc2, 0x98,
	0x1c, 0xe2, 0xc0, 0xcd, 0x4c, 0x9c, 0x33, 0x37, 0x5b, 0x88, 0x8d, 0xb5, 0x9d, 0x6f, 0x60, 0xa9,
	0x6f, 0xb2, 0xad, 0xaf, 0x60, 0xa3, 0x4f, 0x21, 0xc9, 0x64, 0x72, 0x15, 0x8c, 0xb1, 0xfb, 0x98,
	0xf3, 0xf1, 0xfb, 0xce, 0x9f, 0x09, 0x9e, 0xd5, 0x4a, 0xd5, 0x1b, 0xa0, 0xbc, 0x42, 0x6a, 0x65,
	0xaf, 0xba, 0x84, 0x22, 0xe8, 0x4e, 0x94, 0x80, 0xb4, 0xe5, 0xa2, 0x2a, 0x94, 0xae, 0xb9, 0x14,
	0x65, 0x81, 0xc0, 0x75, 0xf9, 0xa6, 0x30, 0xa0, 0x9b, 0xa2, 0x13, 0x70, 0x59, 0x8c, 0x36, 0xd2,
	0x6a, 0x65, 0x54, 0x18, 0x5b, 0x04, 0xe1, 0x15, 0x92, 0x89, 0x46, 0xba, 0x84, 0x38, 0x5a, 0x74,
	0x36, 0x97, 0xa7, 0x01, 0xd5, 0x56, 0x2f, 0x06, 0xda, 0xa0, 0x68, 0xdf, 0x61, 0x5a, 0x41, 0xb9,
	0x94, 0xca, 0x70, 0x23, 0x94, 0xc4, 0xb1, 0x7a, 0xfb, 0xb7, 0x6a, 0xb9, 0x11, 0x20, 0x8d, 0x2d,
	0x1c, 0x3c, 0x0d, 0x0e, 0x57, 0x60, 0xd6, 0x5c, 0x54, 0x2f, 0x2c, 0xff, 0x62, 0xc0, 0xbf, 0x04,
	0xdd, 0x64, 0x02, 0x2e, 0xcf, 0xe1, 0xed, 0x16, 0xd0, 0x84, 0x87, 0xc1, 0x2d, 0xd7, 0x4e, 0x21,
	0x79, 0x03, 0x7b, 0x5e, 0xec, 0xdd, 0xbb, 0x71, 0x7e, 0xd3, 0x3d, 0x3e, 0xe7, 0x0d, 0x24, 0x5f,
	0xfd, 0x20, 0x9e, 0x25, 0x5d, 0xd8, 0x79, 0xc3, 0x1f, 0x5e, 0xb0, 0xff, 0xaf, 0xc4, 0xf0, 0x8c,
	0x2c, 0xad, 0x8c, 0xfc, 0x47, 0xc7, 0xd1, 0xc9, 0x2c, 0x66, 0xda, 0x2b, 0x99, 0x85, 0x1c, 0x3c,
	0xfe, 0xf0, 0xed, 0xfb, 0x27, 0xff, 0x51, 0x78, 0xd2, 0x1f, 0xe2, 0xdd, 0x1f, 0xa3, 0x3f, 0x2c,
	0xb7, 0x68, 0x54, 0x03, 0x1a, 0xe9, 0xf1, 0x70, 0x99, 0xbf, 0x12, 0x90, 0x1e, 0xbf, 0x8f, 0xee,
	0x5c, 0xa5, 0x7b, 0xbb, 0xe8, 0x51, 0xb5, 0x02, 0x49, 0xa9, 0x9a, 0xd3, 0x8f, 0x7e, 0x70, 0x54,
	0xaa, 0x66, 0x71, 0xda, 0xd3, 0xbb, 0x4b, 0x3b, 0x5d, 0xf7, 0x97, 0x5c, 0x7b, 0xaf, 0x9e, 0x8c,
	0xa8, 0x5a, 0x6d, 0xb8, 0xac, 0x89, 0xd2, 0x35, 0xad, 0x41, 0x0e, 0x77, 0xa6, 0xbb, 0xf0, 0xf9,
	0x8f, 0xfd, 0xc0, 0x89, 0xcf, 0xfe, 0xb5, 0x55, 0x9a, 0x7e, 0xf1, 0xe3, 0x95, 0x05, 0xa6, 0x15,
	0x12, 0x2b, 0x7b, 0x95, 0x25, 0x64, 0x0c, 0xc6, 0x2b, 0x67, 0xc9, 0xd3, 0x0a, 0xf3, 0xc9, 0x92,
	0x67, 0x49, 0xee, 0x2c, 0x3f, 0xfd, 0x23, 0xfb, 0xce, 0x58, 0x5a, 0x21, 0x63, 0x93, 0x89, 0xb1,
	0x2c, 0x61, 0xcc, 0xd9, 0x5e, 0x5f, 0x1f, 0xfa, 0xbc, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x56, 0x1b, 0xa9, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaidOrganicSearchTermViewServiceClient is the client API for PaidOrganicSearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaidOrganicSearchTermViewServiceClient interface {
	// Returns the requested paid organic search term view in full detail.
	GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error)
}

type paidOrganicSearchTermViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaidOrganicSearchTermViewServiceClient(cc grpc.ClientConnInterface) PaidOrganicSearchTermViewServiceClient {
	return &paidOrganicSearchTermViewServiceClient{cc}
}

func (c *paidOrganicSearchTermViewServiceClient) GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error) {
	out := new(resources.PaidOrganicSearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaidOrganicSearchTermViewServiceServer is the server API for PaidOrganicSearchTermViewService service.
type PaidOrganicSearchTermViewServiceServer interface {
	// Returns the requested paid organic search term view in full detail.
	GetPaidOrganicSearchTermView(context.Context, *GetPaidOrganicSearchTermViewRequest) (*resources.PaidOrganicSearchTermView, error)
}

// UnimplementedPaidOrganicSearchTermViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaidOrganicSearchTermViewServiceServer struct {
}

func (*UnimplementedPaidOrganicSearchTermViewServiceServer) GetPaidOrganicSearchTermView(ctx context.Context, req *GetPaidOrganicSearchTermViewRequest) (*resources.PaidOrganicSearchTermView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaidOrganicSearchTermView not implemented")
}

func RegisterPaidOrganicSearchTermViewServiceServer(s *grpc.Server, srv PaidOrganicSearchTermViewServiceServer) {
	s.RegisterService(&_PaidOrganicSearchTermViewService_serviceDesc, srv)
}

func _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaidOrganicSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, req.(*GetPaidOrganicSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaidOrganicSearchTermViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.PaidOrganicSearchTermViewService",
	HandlerType: (*PaidOrganicSearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaidOrganicSearchTermView",
			Handler:    _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/paid_organic_search_term_view_service.proto",
}
