// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouderrorreporting/v1beta1/error_group_service.proto

package clouderrorreporting

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A request to return an individual group.
type GetGroupRequest struct {
	// Required. The group resource name. Written as
	// <code>projects/<var>projectID</var>/groups/<var>group_name</var></code>.
	// Call
	// <a href="/error-reporting/reference/rest/v1beta1/projects.groupStats/list">
	// <code>groupStats.list</code></a> to return a list of groups belonging to
	// this project.
	//
	// Example: <code>projects/my-project-123/groups/my-group</code>
	GroupName            string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupRequest) Reset()         { *m = GetGroupRequest{} }
func (m *GetGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupRequest) ProtoMessage()    {}
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cada5d349d61e784, []int{0}
}

func (m *GetGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupRequest.Unmarshal(m, b)
}
func (m *GetGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupRequest.Merge(m, src)
}
func (m *GetGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupRequest.Size(m)
}
func (m *GetGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupRequest proto.InternalMessageInfo

func (m *GetGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

// A request to replace the existing data for the given group.
type UpdateGroupRequest struct {
	// Required. The group which replaces the resource on the server.
	Group                *ErrorGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cada5d349d61e784, []int{1}
}

func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(m, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetGroup() *ErrorGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGroupRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.GetGroupRequest")
	proto.RegisterType((*UpdateGroupRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.UpdateGroupRequest")
}

func init() {
	proto.RegisterFile("google/devtools/clouderrorreporting/v1beta1/error_group_service.proto", fileDescriptor_cada5d349d61e784)
}

var fileDescriptor_cada5d349d61e784 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0x96, 0x53, 0x81, 0xe8, 0x75, 0x40, 0x78, 0xe0, 0xc3, 0x42, 0x02, 0x85, 0x05, 0x5a, 0xf5,
	0x4e, 0x09, 0x42, 0xa0, 0xf2, 0x69, 0xa3, 0x2a, 0x1b, 0xaa, 0x02, 0x64, 0x40, 0x11, 0xd1, 0xc5,
	0x79, 0xeb, 0x18, 0xd9, 0x7e, 0x8f, 0xbb, 0x73, 0x32, 0x20, 0x16, 0xfe, 0x02, 0xff, 0x82, 0x7f,
	0xc1, 0x48, 0x47, 0xd8, 0x3a, 0x75, 0xe0, 0x47, 0xa0, 0xb2, 0x20, 0xdf, 0x39, 0x8d, 0x9b, 0x06,
	0x81, 0xb3, 0xde, 0xf3, 0xde, 0xf3, 0x75, 0xef, 0x91, 0xdd, 0x08, 0x31, 0x4a, 0x80, 0x8d, 0x60,
	0xa2, 0x11, 0x13, 0xc5, 0xc2, 0x04, 0xf3, 0x11, 0x48, 0x89, 0x52, 0x82, 0x40, 0xa9, 0xe3, 0x2c,
	0x62, 0x93, 0xd6, 0x10, 0x34, 0x6f, 0x31, 0x73, 0x3c, 0x88, 0x24, 0xe6, 0x62, 0xa0, 0x40, 0x4e,
	0xe2, 0x10, 0xa8, 0x90, 0xa8, 0xd1, 0xdd, 0xb2, 0x34, 0x74, 0x46, 0x43, 0x97, 0xd0, 0xd0, 0x92,
	0xc6, 0xbb, 0x5e, 0x6a, 0x72, 0x11, 0x33, 0x9e, 0x65, 0xa8, 0xb9, 0x8e, 0x31, 0x53, 0x96, 0xca,
	0xbb, 0x52, 0x41, 0xc3, 0x24, 0x86, 0x4c, 0x97, 0xc0, 0x8d, 0x0a, 0xb0, 0x1f, 0x43, 0x32, 0x1a,
	0x0c, 0x61, 0xcc, 0x27, 0x31, 0xca, 0x72, 0xe0, 0x5a, 0x65, 0x40, 0x82, 0xc2, 0x5c, 0xce, 0xfc,
	0x79, 0x0f, 0xea, 0xc4, 0x0c, 0x31, 0x4d, 0x31, 0xb3, 0x37, 0x9b, 0x11, 0xb9, 0xd8, 0x01, 0xdd,
	0x29, 0x32, 0x77, 0xe1, 0x7d, 0x0e, 0x4a, 0xbb, 0xaf, 0x08, 0xb1, 0x1d, 0x64, 0x3c, 0x85, 0xab,
	0xce, 0x4d, 0xe7, 0xf6, 0x7a, 0x70, 0xef, 0xc8, 0x6f, 0x1c, 0xfb, 0x8c, 0x6c, 0x2f, 0x0b, 0x6e,
	0xc5, 0xb9, 0x88, 0x15, 0x0d, 0x31, 0x65, 0xbb, 0x05, 0x6a, 0x19, 0xd7, 0x0d, 0xd1, 0x0b, 0x9e,
	0x42, 0x73, 0x4c, 0xdc, 0xd7, 0x62, 0xc4, 0x35, 0x9c, 0xd2, 0xea, 0x92, 0x73, 0x66, 0xc4, 0xc8,
	0x6c, 0xb4, 0xef, 0xd3, 0x1a, 0x45, 0xd3, 0xb9, 0x50, 0xb0, 0x76, 0xe4, 0x37, 0xba, 0x96, 0xaa,
	0xfd, 0x7b, 0x8d, 0x5c, 0x9a, 0x43, 0x2f, 0xed, 0x43, 0xba, 0x5f, 0x1d, 0x72, 0x61, 0x96, 0xd4,
	0x7d, 0x54, 0x4b, 0x67, 0xa1, 0x20, 0x6f, 0x55, 0x97, 0xcd, 0x27, 0x87, 0x7e, 0xa5, 0xda, 0x4f,
	0x3f, 0x7e, 0x7e, 0x6e, 0x6c, 0xb9, 0x77, 0x4e, 0x1e, 0xe6, 0xc3, 0x1c, 0x7b, 0x2c, 0x24, 0xbe,
	0x83, 0x50, 0x2b, 0xb6, 0xc9, 0xcc, 0xa9, 0x62, 0x9b, 0x1f, 0xdd, 0x6f, 0x0e, 0xd9, 0xa8, 0x94,
	0xe8, 0x3e, 0xad, 0x65, 0xe4, 0x6c, 0xfd, 0xab, 0x27, 0x79, 0x76, 0xe8, 0xdb, 0xb6, 0x4d, 0x88,
	0xb6, 0xb7, 0x18, 0x82, 0xfe, 0x35, 0xc4, 0x8e, 0xbd, 0xe7, 0xf5, 0x0e, 0xfc, 0xe6, 0xbf, 0x97,
	0xe9, 0xbb, 0x4f, 0xc7, 0x5a, 0x0b, 0xb5, 0xc3, 0xd8, 0x74, 0x3a, 0x5d, 0xdc, 0x34, 0x9e, 0xeb,
	0xb1, 0x5d, 0xf5, 0x6d, 0x91, 0x70, 0xbd, 0x8f, 0x32, 0x0d, 0x8e, 0x1d, 0x52, 0x6c, 0x78, 0x9d,
	0x60, 0xc1, 0xe5, 0x33, 0xeb, 0xb2, 0x57, 0x7c, 0x8e, 0x3d, 0xe7, 0xcd, 0xdb, 0x92, 0x26, 0xc2,
	0x84, 0x67, 0x11, 0x45, 0x19, 0xb1, 0x08, 0x32, 0xf3, 0x75, 0xd8, 0xdc, 0xcc, 0x7f, 0xfd, 0xbb,
	0x87, 0x4b, 0xb0, 0x5f, 0x8e, 0xf3, 0xa5, 0x71, 0xab, 0x63, 0x35, 0x9e, 0x17, 0xb8, 0x2d, 0xb9,
	0x7b, 0x62, 0xb1, 0xd7, 0x0a, 0x8a, 0xcb, 0x07, 0xb3, 0xa9, 0xbe, 0x99, 0xea, 0x9f, 0x9e, 0xea,
	0xf7, 0xac, 0xc4, 0xf0, 0xbc, 0x71, 0x76, 0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x2f,
	0x47, 0xcc, 0xf7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ErrorGroupServiceClient is the client API for ErrorGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErrorGroupServiceClient interface {
	// Get the specified group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
}

type errorGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorGroupServiceClient(cc grpc.ClientConnInterface) ErrorGroupServiceClient {
	return &errorGroupServiceClient{cc}
}

func (c *errorGroupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorGroupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorGroupServiceServer is the server API for ErrorGroupService service.
type ErrorGroupServiceServer interface {
	// Get the specified group.
	GetGroup(context.Context, *GetGroupRequest) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*ErrorGroup, error)
}

// UnimplementedErrorGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedErrorGroupServiceServer struct {
}

func (*UnimplementedErrorGroupServiceServer) GetGroup(ctx context.Context, req *GetGroupRequest) (*ErrorGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (*UnimplementedErrorGroupServiceServer) UpdateGroup(ctx context.Context, req *UpdateGroupRequest) (*ErrorGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}

func RegisterErrorGroupServiceServer(s *grpc.Server, srv ErrorGroupServiceServer) {
	s.RegisterService(&_ErrorGroupService_serviceDesc, srv)
}

func _ErrorGroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ErrorGroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ErrorGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ErrorGroupService",
	HandlerType: (*ErrorGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _ErrorGroupService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _ErrorGroupService_UpdateGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/error_group_service.proto",
}
