// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryGetVendorInfoRequest struct {
	Vid int32 `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
}

func (m *QueryGetVendorInfoRequest) Reset()         { *m = QueryGetVendorInfoRequest{} }
func (m *QueryGetVendorInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetVendorInfoRequest) ProtoMessage()    {}
func (*QueryGetVendorInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{0}
}
func (m *QueryGetVendorInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVendorInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVendorInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVendorInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVendorInfoRequest.Merge(m, src)
}
func (m *QueryGetVendorInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVendorInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVendorInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVendorInfoRequest proto.InternalMessageInfo

func (m *QueryGetVendorInfoRequest) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

type QueryGetVendorInfoResponse struct {
	VendorInfo VendorInfo `protobuf:"bytes,1,opt,name=vendor_info,json=vendorInfo,proto3" json:"vendor_info"`
}

func (m *QueryGetVendorInfoResponse) Reset()         { *m = QueryGetVendorInfoResponse{} }
func (m *QueryGetVendorInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetVendorInfoResponse) ProtoMessage()    {}
func (*QueryGetVendorInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{1}
}
func (m *QueryGetVendorInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVendorInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVendorInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVendorInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVendorInfoResponse.Merge(m, src)
}
func (m *QueryGetVendorInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVendorInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVendorInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVendorInfoResponse proto.InternalMessageInfo

func (m *QueryGetVendorInfoResponse) GetVendorInfo() VendorInfo {
	if m != nil {
		return m.VendorInfo
	}
	return VendorInfo{}
}

type QueryAllVendorInfoRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVendorInfoRequest) Reset()         { *m = QueryAllVendorInfoRequest{} }
func (m *QueryAllVendorInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllVendorInfoRequest) ProtoMessage()    {}
func (*QueryAllVendorInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{2}
}
func (m *QueryAllVendorInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVendorInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVendorInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVendorInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVendorInfoRequest.Merge(m, src)
}
func (m *QueryAllVendorInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVendorInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVendorInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVendorInfoRequest proto.InternalMessageInfo

func (m *QueryAllVendorInfoRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllVendorInfoResponse struct {
	VendorInfo []VendorInfo        `protobuf:"bytes,1,rep,name=vendor_info,json=vendorInfo,proto3" json:"vendor_info"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVendorInfoResponse) Reset()         { *m = QueryAllVendorInfoResponse{} }
func (m *QueryAllVendorInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllVendorInfoResponse) ProtoMessage()    {}
func (*QueryAllVendorInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{3}
}
func (m *QueryAllVendorInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVendorInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVendorInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVendorInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVendorInfoResponse.Merge(m, src)
}
func (m *QueryAllVendorInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVendorInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVendorInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVendorInfoResponse proto.InternalMessageInfo

func (m *QueryAllVendorInfoResponse) GetVendorInfo() []VendorInfo {
	if m != nil {
		return m.VendorInfo
	}
	return nil
}

func (m *QueryAllVendorInfoResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetVendorInfoRequest)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryGetVendorInfoRequest")
	proto.RegisterType((*QueryGetVendorInfoResponse)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryGetVendorInfoResponse")
	proto.RegisterType((*QueryAllVendorInfoRequest)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryAllVendorInfoRequest")
	proto.RegisterType((*QueryAllVendorInfoResponse)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryAllVendorInfoResponse")
}

func init() { proto.RegisterFile("vendorinfo/query.proto", fileDescriptor_e4af47a04389bb67) }

var fileDescriptor_e4af47a04389bb67 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xc0, 0xb3, 0x8d, 0xf5, 0x30, 0x45, 0x90, 0x41, 0xd4, 0x86, 0xba, 0x91, 0x05, 0xff, 0x20,
	0x64, 0x86, 0x56, 0xfc, 0x00, 0xe9, 0xc1, 0x22, 0x5e, 0xda, 0x1c, 0x3c, 0x78, 0x91, 0xd9, 0xdd,
	0xd7, 0xe9, 0xc0, 0x64, 0xde, 0x76, 0x67, 0xb2, 0x58, 0xc5, 0x8b, 0x07, 0xcf, 0x82, 0xdf, 0xc6,
	0x4f, 0xd0, 0x9b, 0x05, 0x11, 0x7a, 0x12, 0x49, 0xfc, 0x20, 0xb2, 0x3b, 0x5b, 0x77, 0xb5, 0x09,
	0x42, 0x63, 0x6f, 0x8f, 0xc9, 0xfb, 0xf3, 0xfb, 0xbd, 0x3c, 0x96, 0xdc, 0x2c, 0xc0, 0xa4, 0x98,
	0x2b, 0xb3, 0x8f, 0xfc, 0x70, 0x02, 0xf9, 0x11, 0xcb, 0x72, 0x74, 0x48, 0x9f, 0xbc, 0x51, 0x32,
	0x06, 0x10, 0x5a, 0x2b, 0x61, 0x12, 0x60, 0xa9, 0xb2, 0x2e, 0x57, 0xf1, 0xc4, 0x41, 0x9a, 0xe0,
	0x38, 0xf3, 0xaf, 0x1a, 0x52, 0x09, 0x39, 0x6b, 0x5a, 0xf4, 0x36, 0x24, 0xa2, 0xd4, 0xc0, 0x45,
	0xa6, 0xb8, 0x30, 0x06, 0x9d, 0x70, 0x0a, 0x8d, 0xf5, 0x4d, 0x7b, 0x8f, 0x12, 0xb4, 0x63, 0xb4,
	0x3c, 0x16, 0x16, 0xfc, 0x34, 0x5e, 0x6c, 0xc6, 0xe0, 0xc4, 0x26, 0xcf, 0x84, 0x54, 0xa6, 0x4a,
	0xae, 0x73, 0x37, 0x5a, 0x60, 0x3e, 0x7c, 0x55, 0xc6, 0xf5, 0xaf, 0x37, 0x24, 0x4a, 0xac, 0x42,
	0x5e, 0x46, 0xfe, 0x35, 0x1a, 0x90, 0xf5, 0xbd, 0xb2, 0xeb, 0x0e, 0xb8, 0x17, 0x55, 0xc9, 0x33,
	0xb3, 0x8f, 0x23, 0x38, 0x9c, 0x80, 0x75, 0xf4, 0x3a, 0xe9, 0x16, 0x2a, 0xbd, 0x1d, 0xdc, 0x0d,
	0x1e, 0xae, 0x8e, 0xca, 0x30, 0xfa, 0x10, 0x90, 0xde, 0xbc, 0x7c, 0x9b, 0xa1, 0xb1, 0x40, 0x0f,
	0xc8, 0x5a, 0x6b, 0x70, 0x55, 0xb8, 0xb6, 0x35, 0x64, 0x17, 0x5a, 0x0c, 0x6b, 0xfa, 0x6f, 0x5f,
	0x39, 0xfe, 0xde, 0xef, 0x8c, 0x48, 0xf1, 0xfb, 0x25, 0x4a, 0x6a, 0xee, 0xa1, 0xd6, 0xe7, 0xb9,
	0x9f, 0x12, 0xd2, 0x2c, 0xa7, 0xa6, 0xb8, 0xcf, 0xfc, 0x26, 0x59, 0xb9, 0x49, 0xe6, 0xff, 0xb7,
	0x7a, 0x93, 0x6c, 0x57, 0x48, 0xa8, 0x6b, 0x47, 0xad, 0xca, 0xe8, 0xcb, 0x99, 0xed, 0x5f, 0x53,
	0x16, 0xd9, 0x76, 0x2f, 0xc9, 0x96, 0xee, 0xfc, 0x21, 0xb4, 0x52, 0x09, 0x3d, 0xf8, 0xa7, 0x90,
	0xc7, 0x6c, 0x1b, 0x6d, 0x7d, 0xee, 0x92, 0xd5, 0xca, 0x88, 0x9e, 0x06, 0x84, 0x34, 0x33, 0xe9,
	0xee, 0x05, 0xb1, 0x17, 0x1e, 0x4f, 0x6f, 0xef, 0x3f, 0x76, 0xf4, 0x26, 0xd1, 0xbd, 0xf7, 0x5f,
	0x7f, 0x7e, 0x5a, 0xe9, 0xd3, 0x3b, 0x3c, 0x4d, 0x34, 0x3f, 0x77, 0xed, 0x96, 0xbf, 0x2d, 0x54,
	0xfa, 0x8e, 0x7e, 0x0b, 0xc8, 0xb5, 0xa6, 0x7a, 0xa8, 0xf5, 0x72, 0x76, 0xf3, 0x4e, 0x6c, 0x39,
	0xbb, 0xb9, 0xe7, 0x14, 0xf5, 0x2b, 0xbb, 0x75, 0x7a, 0x6b, 0x81, 0xdd, 0x36, 0x1c, 0x4f, 0xc3,
	0xe0, 0x64, 0x1a, 0x06, 0x3f, 0xa6, 0x61, 0xf0, 0x71, 0x16, 0x76, 0x4e, 0x66, 0x61, 0xe7, 0x74,
	0x16, 0x76, 0x5e, 0x3e, 0x97, 0xca, 0x1d, 0x4c, 0x62, 0x96, 0xe0, 0x98, 0x7b, 0xae, 0xc1, 0x19,
	0x18, 0x6f, 0x81, 0x0d, 0x1a, 0xb2, 0x81, 0x47, 0xe3, 0xaf, 0xdb, 0x83, 0xdc, 0x51, 0x06, 0x36,
	0xbe, 0x5a, 0x7d, 0x19, 0x1e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x22, 0xa9, 0x2d, 0x37, 0xe8,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a vendorInfo by index.
	VendorInfo(ctx context.Context, in *QueryGetVendorInfoRequest, opts ...grpc.CallOption) (*QueryGetVendorInfoResponse, error)
	// Queries a list of vendorInfo items.
	VendorInfoAll(ctx context.Context, in *QueryAllVendorInfoRequest, opts ...grpc.CallOption) (*QueryAllVendorInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) VendorInfo(ctx context.Context, in *QueryGetVendorInfoRequest, opts ...grpc.CallOption) (*QueryGetVendorInfoResponse, error) {
	out := new(QueryGetVendorInfoResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VendorInfoAll(ctx context.Context, in *QueryAllVendorInfoRequest, opts ...grpc.CallOption) (*QueryAllVendorInfoResponse, error) {
	out := new(QueryAllVendorInfoResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a vendorInfo by index.
	VendorInfo(context.Context, *QueryGetVendorInfoRequest) (*QueryGetVendorInfoResponse, error)
	// Queries a list of vendorInfo items.
	VendorInfoAll(context.Context, *QueryAllVendorInfoRequest) (*QueryAllVendorInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) VendorInfo(ctx context.Context, req *QueryGetVendorInfoRequest) (*QueryGetVendorInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorInfo not implemented")
}
func (*UnimplementedQueryServer) VendorInfoAll(ctx context.Context, req *QueryAllVendorInfoRequest) (*QueryAllVendorInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorInfoAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_VendorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetVendorInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VendorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VendorInfo(ctx, req.(*QueryGetVendorInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VendorInfoAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllVendorInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VendorInfoAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VendorInfoAll(ctx, req.(*QueryAllVendorInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zigbeealliance.distributedcomplianceledger.vendorinfo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VendorInfo",
			Handler:    _Query_VendorInfo_Handler,
		},
		{
			MethodName: "VendorInfoAll",
			Handler:    _Query_VendorInfoAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendorinfo/query.proto",
}

func (m *QueryGetVendorInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVendorInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVendorInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vid != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetVendorInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVendorInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVendorInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VendorInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllVendorInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVendorInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVendorInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllVendorInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVendorInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVendorInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.VendorInfo) > 0 {
		for iNdEx := len(m.VendorInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VendorInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetVendorInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovQuery(uint64(m.Vid))
	}
	return n
}

func (m *QueryGetVendorInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VendorInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllVendorInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllVendorInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VendorInfo) > 0 {
		for _, e := range m.VendorInfo {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetVendorInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetVendorInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVendorInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetVendorInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetVendorInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVendorInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VendorInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllVendorInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllVendorInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVendorInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllVendorInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllVendorInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVendorInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorInfo = append(m.VendorInfo, VendorInfo{})
			if err := m.VendorInfo[len(m.VendorInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
