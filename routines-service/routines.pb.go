// Code generated by protoc-gen-go.
// source: routines.proto
// DO NOT EDIT!

/*
Package routines is a generated protocol buffer package.

It is generated from these files:
	routines.proto

It has these top-level messages:
	ScrapTrelloResetRequest
	ScrapTrelloResetReply
	GetRoutinesRequest
	GetRoutinesReply
	Days
	Items
	Item
	Bool
*/
package routines

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/zaquestion/routines/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ScrapTrelloResetRequest struct {
}

func (m *ScrapTrelloResetRequest) Reset()                    { *m = ScrapTrelloResetRequest{} }
func (m *ScrapTrelloResetRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapTrelloResetRequest) ProtoMessage()               {}
func (*ScrapTrelloResetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ScrapTrelloResetReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *ScrapTrelloResetReply) Reset()                    { *m = ScrapTrelloResetReply{} }
func (m *ScrapTrelloResetReply) String() string            { return proto.CompactTextString(m) }
func (*ScrapTrelloResetReply) ProtoMessage()               {}
func (*ScrapTrelloResetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetRoutinesRequest struct {
	DateStart string `protobuf:"bytes,1,opt,name=date_start,json=dateStart" json:"date_start,omitempty"`
	DateEnd   string `protobuf:"bytes,2,opt,name=date_end,json=dateEnd" json:"date_end,omitempty"`
}

func (m *GetRoutinesRequest) Reset()                    { *m = GetRoutinesRequest{} }
func (m *GetRoutinesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRoutinesRequest) ProtoMessage()               {}
func (*GetRoutinesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetRoutinesReply struct {
	// CL NAME  data
	Routines map[string]*Days `protobuf:"bytes,1,rep,name=routines" json:"routines,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Err      string           `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetRoutinesReply) Reset()                    { *m = GetRoutinesReply{} }
func (m *GetRoutinesReply) String() string            { return proto.CompactTextString(m) }
func (*GetRoutinesReply) ProtoMessage()               {}
func (*GetRoutinesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRoutinesReply) GetRoutines() map[string]*Days {
	if m != nil {
		return m.Routines
	}
	return nil
}

type Days struct {
	Day   string `protobuf:"bytes,1,opt,name=day" json:"day,omitempty"`
	Items *Items `protobuf:"bytes,2,opt,name=items" json:"items,omitempty"`
}

func (m *Days) Reset()                    { *m = Days{} }
func (m *Days) String() string            { return proto.CompactTextString(m) }
func (*Days) ProtoMessage()               {}
func (*Days) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Days) GetItems() *Items {
	if m != nil {
		return m.Items
	}
	return nil
}

type Items struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Item *Item  `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
}

func (m *Items) Reset()                    { *m = Items{} }
func (m *Items) String() string            { return proto.CompactTextString(m) }
func (*Items) ProtoMessage()               {}
func (*Items) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Items) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type Item struct {
	Checked     *Bool  `protobuf:"bytes,1,opt,name=checked" json:"checked,omitempty"`
	LastUpdated string `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Item) GetChecked() *Bool {
	if m != nil {
		return m.Checked
	}
	return nil
}

type Bool struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*ScrapTrelloResetRequest)(nil), "routines.ScrapTrelloResetRequest")
	proto.RegisterType((*ScrapTrelloResetReply)(nil), "routines.ScrapTrelloResetReply")
	proto.RegisterType((*GetRoutinesRequest)(nil), "routines.GetRoutinesRequest")
	proto.RegisterType((*GetRoutinesReply)(nil), "routines.GetRoutinesReply")
	proto.RegisterType((*Days)(nil), "routines.Days")
	proto.RegisterType((*Items)(nil), "routines.Items")
	proto.RegisterType((*Item)(nil), "routines.Item")
	proto.RegisterType((*Bool)(nil), "routines.Bool")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RoutinesService service

type RoutinesServiceClient interface {
	ScrapTrelloReset(ctx context.Context, in *ScrapTrelloResetRequest, opts ...grpc.CallOption) (*ScrapTrelloResetReply, error)
	GetRoutines(ctx context.Context, in *GetRoutinesRequest, opts ...grpc.CallOption) (*GetRoutinesReply, error)
}

type routinesServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoutinesServiceClient(cc *grpc.ClientConn) RoutinesServiceClient {
	return &routinesServiceClient{cc}
}

func (c *routinesServiceClient) ScrapTrelloReset(ctx context.Context, in *ScrapTrelloResetRequest, opts ...grpc.CallOption) (*ScrapTrelloResetReply, error) {
	out := new(ScrapTrelloResetReply)
	err := grpc.Invoke(ctx, "/routines.RoutinesService/ScrapTrelloReset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routinesServiceClient) GetRoutines(ctx context.Context, in *GetRoutinesRequest, opts ...grpc.CallOption) (*GetRoutinesReply, error) {
	out := new(GetRoutinesReply)
	err := grpc.Invoke(ctx, "/routines.RoutinesService/GetRoutines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoutinesService service

type RoutinesServiceServer interface {
	ScrapTrelloReset(context.Context, *ScrapTrelloResetRequest) (*ScrapTrelloResetReply, error)
	GetRoutines(context.Context, *GetRoutinesRequest) (*GetRoutinesReply, error)
}

func RegisterRoutinesServiceServer(s *grpc.Server, srv RoutinesServiceServer) {
	s.RegisterService(&_RoutinesService_serviceDesc, srv)
}

func _RoutinesService_ScrapTrelloReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapTrelloResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutinesServiceServer).ScrapTrelloReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routines.RoutinesService/ScrapTrelloReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutinesServiceServer).ScrapTrelloReset(ctx, req.(*ScrapTrelloResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutinesService_GetRoutines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoutinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutinesServiceServer).GetRoutines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routines.RoutinesService/GetRoutines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutinesServiceServer).GetRoutines(ctx, req.(*GetRoutinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutinesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routines.RoutinesService",
	HandlerType: (*RoutinesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScrapTrelloReset",
			Handler:    _RoutinesService_ScrapTrelloReset_Handler,
		},
		{
			MethodName: "GetRoutines",
			Handler:    _RoutinesService_GetRoutines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("routines.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x53, 0x87, 0x26, 0x63, 0xda, 0x84, 0x11, 0x15, 0x69, 0x14, 0x44, 0xbb, 0x02, 0x29,
	0x5c, 0x62, 0x29, 0x5c, 0x10, 0x97, 0x4a, 0xa8, 0x15, 0x42, 0x48, 0x1c, 0x36, 0x70, 0xe1, 0x52,
	0x2d, 0xf1, 0xa8, 0x58, 0x71, 0xbc, 0x66, 0xbd, 0xa9, 0xe4, 0x2b, 0xaf, 0xc0, 0xc3, 0xf0, 0x20,
	0x5c, 0x39, 0xf2, 0x20, 0x68, 0x36, 0x5e, 0x87, 0x34, 0x4a, 0x6f, 0x3b, 0xdf, 0xf7, 0xcd, 0x37,
	0x3f, 0x1e, 0xc3, 0xb1, 0xd1, 0x2b, 0x9b, 0xe6, 0x54, 0x4e, 0x0a, 0xa3, 0xad, 0xc6, 0x8e, 0x8f,
	0x87, 0xa3, 0x1b, 0xad, 0x6f, 0x32, 0x8a, 0x55, 0x91, 0xc6, 0x2a, 0xcf, 0xb5, 0x55, 0x36, 0xd5,
	0x79, 0xad, 0x13, 0xa7, 0xf0, 0x64, 0x36, 0x37, 0xaa, 0xf8, 0x64, 0x28, 0xcb, 0xb4, 0xa4, 0x92,
	0xac, 0xa4, 0xef, 0x2b, 0x2a, 0xad, 0x78, 0x09, 0x27, 0xbb, 0x54, 0x91, 0x55, 0xd8, 0x87, 0x03,
	0x32, 0x66, 0x10, 0x9c, 0x05, 0xe3, 0xae, 0xe4, 0xa7, 0xf8, 0x08, 0xf8, 0x8e, 0xac, 0xac, 0x4b,
	0xd6, 0x06, 0xf8, 0x14, 0x20, 0x51, 0x96, 0xae, 0x4b, 0xab, 0x8c, 0xad, 0xe5, 0x5d, 0x46, 0x66,
	0x0c, 0xe0, 0x29, 0x74, 0x1c, 0x4d, 0x79, 0x32, 0x68, 0x39, 0xf2, 0x90, 0xe3, 0xab, 0x3c, 0x11,
	0xbf, 0x02, 0xe8, 0x6f, 0x19, 0x72, 0xd9, 0x4b, 0x68, 0x86, 0x1a, 0x04, 0x67, 0x07, 0xe3, 0x68,
	0x3a, 0x9e, 0x34, 0x53, 0xdf, 0x55, 0x4f, 0x7c, 0x74, 0x95, 0x5b, 0x53, 0xc9, 0x26, 0xd3, 0x37,
	0xdf, 0x6a, 0x9a, 0x1f, 0x7e, 0x80, 0xa3, 0x2d, 0x31, 0x4b, 0x16, 0x54, 0xf9, 0xf9, 0x16, 0x54,
	0xe1, 0x73, 0x68, 0xdf, 0xaa, 0x6c, 0x45, 0x2e, 0x2d, 0x9a, 0x1e, 0x6f, 0xea, 0x5e, 0xaa, 0xaa,
	0x94, 0x6b, 0xf2, 0x4d, 0xeb, 0x75, 0x20, 0x2e, 0x20, 0x64, 0x88, 0x3d, 0x12, 0xd5, 0x78, 0x24,
	0xaa, 0xc2, 0x17, 0xd0, 0x4e, 0x2d, 0x2d, 0xcb, 0xda, 0xa3, 0xb7, 0xf1, 0x78, 0xcf, 0xb0, 0x5c,
	0xb3, 0xe2, 0x02, 0xda, 0x2e, 0x46, 0x84, 0x30, 0x57, 0x4b, 0xaa, 0x2d, 0xdc, 0x1b, 0x05, 0x84,
	0xac, 0xda, 0x6d, 0x83, 0x53, 0xa4, 0xe3, 0xc4, 0x0c, 0x42, 0x8e, 0x70, 0x0c, 0x87, 0xf3, 0x6f,
	0x34, 0x5f, 0x50, 0xe2, 0x2c, 0xb6, 0xe4, 0x6f, 0xb5, 0xce, 0xa4, 0xa7, 0xf1, 0x1c, 0x1e, 0x66,
	0xaa, 0xb4, 0xd7, 0xab, 0x82, 0xf7, 0xef, 0x3f, 0x46, 0xc4, 0xd8, 0xe7, 0x35, 0x24, 0x46, 0x10,
	0x72, 0x0e, 0x3e, 0xf6, 0x8b, 0x60, 0xcb, 0x4e, 0x3d, 0xf8, 0xf4, 0x4f, 0x00, 0x3d, 0xbf, 0xc2,
	0x19, 0x99, 0xdb, 0x74, 0x4e, 0xb8, 0x84, 0xfe, 0xdd, 0xeb, 0xc1, 0xf3, 0x4d, 0x07, 0x7b, 0x8e,
	0x6e, 0xf8, 0xec, 0x3e, 0x49, 0x91, 0x55, 0xe2, 0xe4, 0xc7, 0xef, 0xbf, 0x3f, 0x5b, 0x3d, 0x3c,
	0x8a, 0x4b, 0xe6, 0x63, 0xeb, 0x04, 0xf8, 0x05, 0xa2, 0xff, 0x4e, 0x00, 0x47, 0x7b, 0x2e, 0x63,
	0x5d, 0x64, 0xb8, 0xff, 0x6e, 0xc4, 0x23, 0xe7, 0x1f, 0x61, 0x37, 0xf6, 0x9a, 0xaf, 0x0f, 0xdc,
	0xaf, 0xf2, 0xea, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x93, 0x97, 0x71, 0x64, 0x03, 0x00,
	0x00,
}