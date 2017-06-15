// Code generated by protoc-gen-go. DO NOT EDIT.
// source: point.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	point.proto

It has these top-level messages:
	Point
	ThingId
	GroupId
	StringPoint
	FloatPoint
	IntegerPoint
	DurationPoint
	DateTimePoint
	BoolPoint
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf3 "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	User         string                           `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Tags         map[string]string                `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Fields       map[string]*google_protobuf3.Any `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DateCreation *google_protobuf1.Timestamp      `protobuf:"bytes,4,opt,name=dateCreation" json:"dateCreation,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto1.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Point) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Point) GetFields() map[string]*google_protobuf3.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Point) GetDateCreation() *google_protobuf1.Timestamp {
	if m != nil {
		return m.DateCreation
	}
	return nil
}

type ThingId struct {
	User    string                      `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	ThingId string                      `protobuf:"bytes,2,opt,name=thingId" json:"thingId,omitempty"`
	Start   *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	End     *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
}

func (m *ThingId) Reset()                    { *m = ThingId{} }
func (m *ThingId) String() string            { return proto1.CompactTextString(m) }
func (*ThingId) ProtoMessage()               {}
func (*ThingId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ThingId) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ThingId) GetThingId() string {
	if m != nil {
		return m.ThingId
	}
	return ""
}

func (m *ThingId) GetStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ThingId) GetEnd() *google_protobuf1.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type GroupId struct {
	User    string                      `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	GroupId string                      `protobuf:"bytes,2,opt,name=groupId" json:"groupId,omitempty"`
	Start   *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	End     *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
}

func (m *GroupId) Reset()                    { *m = GroupId{} }
func (m *GroupId) String() string            { return proto1.CompactTextString(m) }
func (*GroupId) ProtoMessage()               {}
func (*GroupId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GroupId) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GroupId) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *GroupId) GetStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *GroupId) GetEnd() *google_protobuf1.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type StringPoint struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringPoint) Reset()                    { *m = StringPoint{} }
func (m *StringPoint) String() string            { return proto1.CompactTextString(m) }
func (*StringPoint) ProtoMessage()               {}
func (*StringPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StringPoint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type FloatPoint struct {
	Value float32 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
}

func (m *FloatPoint) Reset()                    { *m = FloatPoint{} }
func (m *FloatPoint) String() string            { return proto1.CompactTextString(m) }
func (*FloatPoint) ProtoMessage()               {}
func (*FloatPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FloatPoint) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type IntegerPoint struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *IntegerPoint) Reset()                    { *m = IntegerPoint{} }
func (m *IntegerPoint) String() string            { return proto1.CompactTextString(m) }
func (*IntegerPoint) ProtoMessage()               {}
func (*IntegerPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IntegerPoint) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type DurationPoint struct {
	Value *google_protobuf2.Duration `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DurationPoint) Reset()                    { *m = DurationPoint{} }
func (m *DurationPoint) String() string            { return proto1.CompactTextString(m) }
func (*DurationPoint) ProtoMessage()               {}
func (*DurationPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DurationPoint) GetValue() *google_protobuf2.Duration {
	if m != nil {
		return m.Value
	}
	return nil
}

type DateTimePoint struct {
	Value *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DateTimePoint) Reset()                    { *m = DateTimePoint{} }
func (m *DateTimePoint) String() string            { return proto1.CompactTextString(m) }
func (*DateTimePoint) ProtoMessage()               {}
func (*DateTimePoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DateTimePoint) GetValue() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Value
	}
	return nil
}

type BoolPoint struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BoolPoint) Reset()                    { *m = BoolPoint{} }
func (m *BoolPoint) String() string            { return proto1.CompactTextString(m) }
func (*BoolPoint) ProtoMessage()               {}
func (*BoolPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BoolPoint) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto1.RegisterType((*Point)(nil), "proto.Point")
	proto1.RegisterType((*ThingId)(nil), "proto.ThingId")
	proto1.RegisterType((*GroupId)(nil), "proto.GroupId")
	proto1.RegisterType((*StringPoint)(nil), "proto.StringPoint")
	proto1.RegisterType((*FloatPoint)(nil), "proto.FloatPoint")
	proto1.RegisterType((*IntegerPoint)(nil), "proto.IntegerPoint")
	proto1.RegisterType((*DurationPoint)(nil), "proto.DurationPoint")
	proto1.RegisterType((*DateTimePoint)(nil), "proto.DateTimePoint")
	proto1.RegisterType((*BoolPoint)(nil), "proto.BoolPoint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PointSvc service

type PointSvcClient interface {
	CreatePoint(ctx context.Context, in *Point, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetPointsByThing(ctx context.Context, in *ThingId, opts ...grpc.CallOption) (PointSvc_GetPointsByThingClient, error)
	GetPointsByGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (PointSvc_GetPointsByGroupClient, error)
}

type pointSvcClient struct {
	cc *grpc.ClientConn
}

func NewPointSvcClient(cc *grpc.ClientConn) PointSvcClient {
	return &pointSvcClient{cc}
}

func (c *pointSvcClient) CreatePoint(ctx context.Context, in *Point, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/proto.PointSvc/CreatePoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointSvcClient) GetPointsByThing(ctx context.Context, in *ThingId, opts ...grpc.CallOption) (PointSvc_GetPointsByThingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PointSvc_serviceDesc.Streams[0], c.cc, "/proto.PointSvc/GetPointsByThing", opts...)
	if err != nil {
		return nil, err
	}
	x := &pointSvcGetPointsByThingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PointSvc_GetPointsByThingClient interface {
	Recv() (*Point, error)
	grpc.ClientStream
}

type pointSvcGetPointsByThingClient struct {
	grpc.ClientStream
}

func (x *pointSvcGetPointsByThingClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pointSvcClient) GetPointsByGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (PointSvc_GetPointsByGroupClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PointSvc_serviceDesc.Streams[1], c.cc, "/proto.PointSvc/GetPointsByGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &pointSvcGetPointsByGroupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PointSvc_GetPointsByGroupClient interface {
	Recv() (*Point, error)
	grpc.ClientStream
}

type pointSvcGetPointsByGroupClient struct {
	grpc.ClientStream
}

func (x *pointSvcGetPointsByGroupClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PointSvc service

type PointSvcServer interface {
	CreatePoint(context.Context, *Point) (*google_protobuf.Empty, error)
	GetPointsByThing(*ThingId, PointSvc_GetPointsByThingServer) error
	GetPointsByGroup(*GroupId, PointSvc_GetPointsByGroupServer) error
}

func RegisterPointSvcServer(s *grpc.Server, srv PointSvcServer) {
	s.RegisterService(&_PointSvc_serviceDesc, srv)
}

func _PointSvc_CreatePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointSvcServer).CreatePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PointSvc/CreatePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointSvcServer).CreatePoint(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointSvc_GetPointsByThing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ThingId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PointSvcServer).GetPointsByThing(m, &pointSvcGetPointsByThingServer{stream})
}

type PointSvc_GetPointsByThingServer interface {
	Send(*Point) error
	grpc.ServerStream
}

type pointSvcGetPointsByThingServer struct {
	grpc.ServerStream
}

func (x *pointSvcGetPointsByThingServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

func _PointSvc_GetPointsByGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GroupId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PointSvcServer).GetPointsByGroup(m, &pointSvcGetPointsByGroupServer{stream})
}

type PointSvc_GetPointsByGroupServer interface {
	Send(*Point) error
	grpc.ServerStream
}

type pointSvcGetPointsByGroupServer struct {
	grpc.ServerStream
}

func (x *pointSvcGetPointsByGroupServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

var _PointSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PointSvc",
	HandlerType: (*PointSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoint",
			Handler:    _PointSvc_CreatePoint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPointsByThing",
			Handler:       _PointSvc_GetPointsByThing_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPointsByGroup",
			Handler:       _PointSvc_GetPointsByGroup_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "point.proto",
}

func init() { proto1.RegisterFile("point.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0x6b, 0x3b, 0x6e, 0x9b, 0x71, 0x8a, 0xaa, 0x55, 0x55, 0xb9, 0x46, 0x82, 0x60, 0x78,
	0x88, 0x2a, 0xe4, 0x44, 0x06, 0x04, 0xe2, 0x01, 0xd1, 0xd2, 0x1f, 0xf5, 0x09, 0x94, 0xe6, 0x02,
	0x5b, 0xbc, 0x35, 0x16, 0xce, 0xae, 0xb5, 0x5e, 0x47, 0xf2, 0x49, 0x38, 0x05, 0xa7, 0xe0, 0x62,
	0xc8, 0xfb, 0x13, 0x39, 0xb1, 0xf9, 0x79, 0xe9, 0x53, 0x3c, 0x9e, 0x6f, 0xc6, 0xdf, 0x8e, 0xc7,
	0x01, 0xaf, 0x60, 0x19, 0x15, 0x51, 0xc1, 0x99, 0x60, 0xc8, 0x95, 0x3f, 0xc1, 0xe3, 0x94, 0xb1,
	0x34, 0x27, 0x53, 0x19, 0xdd, 0x55, 0xf7, 0x53, 0xb2, 0x2c, 0x44, 0xad, 0x98, 0xe0, 0xe9, 0x76,
	0x52, 0x64, 0x4b, 0x52, 0x0a, 0xbc, 0x2c, 0x34, 0xf0, 0x64, 0x1b, 0x48, 0x2a, 0x8e, 0x45, 0xc6,
	0xa8, 0xce, 0x9f, 0x6c, 0xe7, 0x31, 0xd5, 0xbd, 0xc3, 0x5f, 0x36, 0xb8, 0x5f, 0x1a, 0x1f, 0x84,
	0x60, 0x50, 0x95, 0x84, 0xfb, 0xd6, 0xd8, 0x9a, 0x0c, 0xe7, 0xf2, 0x1a, 0x9d, 0xc2, 0x40, 0xe0,
	0xb4, 0xf4, 0xed, 0xb1, 0x33, 0xf1, 0xe2, 0x63, 0x55, 0x13, 0x49, 0x3e, 0x5a, 0xe0, 0xb4, 0xbc,
	0xa4, 0x82, 0xd7, 0x73, 0xc9, 0xa0, 0x19, 0xec, 0xde, 0x67, 0x24, 0x4f, 0x4a, 0xdf, 0x91, 0xb4,
	0xbf, 0x41, 0x5f, 0xc9, 0x94, 0xe2, 0x35, 0x87, 0x3e, 0xc0, 0x28, 0xc1, 0x82, 0x7c, 0xe2, 0x44,
	0xca, 0xfa, 0x83, 0xb1, 0x35, 0xf1, 0xe2, 0x20, 0x52, 0xb6, 0x91, 0xb1, 0x8d, 0x16, 0xe6, 0xb8,
	0xf3, 0x0d, 0x3e, 0x78, 0x0b, 0xc3, 0xb5, 0x04, 0x3a, 0x04, 0xe7, 0x3b, 0xa9, 0xb5, 0x7d, 0x73,
	0x89, 0x8e, 0xc0, 0x5d, 0xe1, 0xbc, 0x22, 0xbe, 0x2d, 0xef, 0xa9, 0xe0, 0xbd, 0xfd, 0xce, 0x0a,
	0x3e, 0x83, 0xd7, 0xf2, 0xe9, 0x29, 0x3d, 0x6d, 0x97, 0x7a, 0xf1, 0x51, 0x47, 0xe9, 0x8c, 0xd6,
	0xad, 0x86, 0xe1, 0x0f, 0x0b, 0xf6, 0x16, 0xdf, 0x32, 0x9a, 0xde, 0x24, 0xbd, 0x73, 0xf4, 0x61,
	0x4f, 0xa8, 0xb4, 0x96, 0x31, 0x21, 0x9a, 0x81, 0x5b, 0x0a, 0xcc, 0x85, 0xef, 0xfc, 0xf3, 0xf0,
	0x0a, 0x44, 0x2f, 0xc1, 0x21, 0x34, 0xf9, 0x8f, 0x61, 0x35, 0x98, 0x34, 0xbb, 0xe6, 0xac, 0x2a,
	0xfe, 0x6c, 0x96, 0xaa, 0xb4, 0x31, 0xd3, 0xe1, 0x83, 0x9b, 0x3d, 0x07, 0xef, 0x56, 0xf0, 0x8c,
	0xa6, 0x6a, 0xfd, 0xd6, 0x6f, 0xcb, 0x6a, 0xbd, 0xad, 0x30, 0x04, 0xb8, 0xca, 0x19, 0x16, 0x3d,
	0x8c, 0x6d, 0x98, 0x17, 0x30, 0xba, 0xa1, 0x82, 0xa4, 0x84, 0xf7, 0x50, 0xae, 0xa1, 0x3e, 0xc2,
	0xc1, 0x85, 0xfe, 0x2a, 0x14, 0x36, 0x6d, 0x63, 0x5e, 0x7c, 0xd2, 0xf1, 0x35, 0xb8, 0xe9, 0x70,
	0x06, 0x07, 0x17, 0x58, 0x90, 0xe6, 0x18, 0xaa, 0xc3, 0x6c, 0xb3, 0xc3, 0x5f, 0x27, 0xa4, 0x5a,
	0x3c, 0x83, 0xe1, 0x39, 0x63, 0x79, 0x8f, 0xe7, 0xbe, 0x46, 0xe2, 0x9f, 0x16, 0xec, 0xcb, 0xfc,
	0xed, 0xea, 0x2b, 0x7a, 0x03, 0x9e, 0xdc, 0x76, 0xfd, 0xc0, 0x51, 0xfb, 0x93, 0x0a, 0x8e, 0x3b,
	0xcf, 0xbb, 0x6c, 0xfe, 0x34, 0xc2, 0x1d, 0xf4, 0x1a, 0x0e, 0xaf, 0x89, 0x9a, 0x59, 0x79, 0x5e,
	0xcb, 0xc5, 0x44, 0x8f, 0x74, 0xad, 0x5e, 0xd3, 0x60, 0xa3, 0x57, 0xb8, 0x33, 0xb3, 0xb6, 0xaa,
	0xe4, 0xd2, 0xac, 0xab, 0xf4, 0x0a, 0x75, 0xab, 0xee, 0x76, 0xe5, 0x8d, 0x57, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x06, 0xe4, 0xda, 0xcd, 0xd6, 0x04, 0x00, 0x00,
}