// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pilot.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HandleSubtaskRequest struct {
	SubtaskId     *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	SubtaskAction *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=subtask_action,json=subtaskAction" json:"subtask_action,omitempty"`
	Directive     *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=directive" json:"directive,omitempty"`
}

func (m *HandleSubtaskRequest) Reset()                    { *m = HandleSubtaskRequest{} }
func (m *HandleSubtaskRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleSubtaskRequest) ProtoMessage()               {}
func (*HandleSubtaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *HandleSubtaskRequest) GetSubtaskId() *google_protobuf2.StringValue {
	if m != nil {
		return m.SubtaskId
	}
	return nil
}

func (m *HandleSubtaskRequest) GetSubtaskAction() *google_protobuf2.StringValue {
	if m != nil {
		return m.SubtaskAction
	}
	return nil
}

func (m *HandleSubtaskRequest) GetDirective() *google_protobuf2.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type HandleSubtaskResponse struct {
	SubtaskId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
}

func (m *HandleSubtaskResponse) Reset()                    { *m = HandleSubtaskResponse{} }
func (m *HandleSubtaskResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleSubtaskResponse) ProtoMessage()               {}
func (*HandleSubtaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *HandleSubtaskResponse) GetSubtaskId() *google_protobuf2.StringValue {
	if m != nil {
		return m.SubtaskId
	}
	return nil
}

type GetSubtaskStatusRequest struct {
	SubtaskId []string `protobuf:"bytes,1,rep,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
}

func (m *GetSubtaskStatusRequest) Reset()                    { *m = GetSubtaskStatusRequest{} }
func (m *GetSubtaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubtaskStatusRequest) ProtoMessage()               {}
func (*GetSubtaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GetSubtaskStatusRequest) GetSubtaskId() []string {
	if m != nil {
		return m.SubtaskId
	}
	return nil
}

type SubtaskStatus struct {
	SubtaskId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=subtask_id,json=subtaskId" json:"subtask_id,omitempty"`
	Status    *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *SubtaskStatus) Reset()                    { *m = SubtaskStatus{} }
func (m *SubtaskStatus) String() string            { return proto.CompactTextString(m) }
func (*SubtaskStatus) ProtoMessage()               {}
func (*SubtaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *SubtaskStatus) GetSubtaskId() *google_protobuf2.StringValue {
	if m != nil {
		return m.SubtaskId
	}
	return nil
}

func (m *SubtaskStatus) GetStatus() *google_protobuf2.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetSubtaskStatusResponse struct {
	TotalCount       uint32           `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	SubtaskStatusSet []*SubtaskStatus `protobuf:"bytes,2,rep,name=subtask_status_set,json=subtaskStatusSet" json:"subtask_status_set,omitempty"`
}

func (m *GetSubtaskStatusResponse) Reset()                    { *m = GetSubtaskStatusResponse{} }
func (m *GetSubtaskStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSubtaskStatusResponse) ProtoMessage()               {}
func (*GetSubtaskStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *GetSubtaskStatusResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *GetSubtaskStatusResponse) GetSubtaskStatusSet() []*SubtaskStatus {
	if m != nil {
		return m.SubtaskStatusSet
	}
	return nil
}

func init() {
	proto.RegisterType((*HandleSubtaskRequest)(nil), "openpitrix.HandleSubtaskRequest")
	proto.RegisterType((*HandleSubtaskResponse)(nil), "openpitrix.HandleSubtaskResponse")
	proto.RegisterType((*GetSubtaskStatusRequest)(nil), "openpitrix.GetSubtaskStatusRequest")
	proto.RegisterType((*SubtaskStatus)(nil), "openpitrix.SubtaskStatus")
	proto.RegisterType((*GetSubtaskStatusResponse)(nil), "openpitrix.GetSubtaskStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PilotService service

type PilotServiceClient interface {
	HandleSubtask(ctx context.Context, in *HandleSubtaskRequest, opts ...grpc.CallOption) (*HandleSubtaskResponse, error)
	GetSubtaskStatus(ctx context.Context, in *GetSubtaskStatusRequest, opts ...grpc.CallOption) (*GetSubtaskStatusResponse, error)
}

type pilotServiceClient struct {
	cc *grpc.ClientConn
}

func NewPilotServiceClient(cc *grpc.ClientConn) PilotServiceClient {
	return &pilotServiceClient{cc}
}

func (c *pilotServiceClient) HandleSubtask(ctx context.Context, in *HandleSubtaskRequest, opts ...grpc.CallOption) (*HandleSubtaskResponse, error) {
	out := new(HandleSubtaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotService/HandleSubtask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetSubtaskStatus(ctx context.Context, in *GetSubtaskStatusRequest, opts ...grpc.CallOption) (*GetSubtaskStatusResponse, error) {
	out := new(GetSubtaskStatusResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotService/GetSubtaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PilotService service

type PilotServiceServer interface {
	HandleSubtask(context.Context, *HandleSubtaskRequest) (*HandleSubtaskResponse, error)
	GetSubtaskStatus(context.Context, *GetSubtaskStatusRequest) (*GetSubtaskStatusResponse, error)
}

func RegisterPilotServiceServer(s *grpc.Server, srv PilotServiceServer) {
	s.RegisterService(&_PilotService_serviceDesc, srv)
}

func _PilotService_HandleSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleSubtaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).HandleSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotService/HandleSubtask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).HandleSubtask(ctx, req.(*HandleSubtaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetSubtaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubtaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotService/GetSubtaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, req.(*GetSubtaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PilotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.PilotService",
	HandlerType: (*PilotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSubtask",
			Handler:    _PilotService_HandleSubtask_Handler,
		},
		{
			MethodName: "GetSubtaskStatus",
			Handler:    _PilotService_GetSubtaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pilot.proto",
}

// Client API for PilotServiceForFrontgate service

type PilotServiceForFrontgateClient interface {
	Channel(ctx context.Context, opts ...grpc.CallOption) (PilotServiceForFrontgate_ChannelClient, error)
}

type pilotServiceForFrontgateClient struct {
	cc *grpc.ClientConn
}

func NewPilotServiceForFrontgateClient(cc *grpc.ClientConn) PilotServiceForFrontgateClient {
	return &pilotServiceForFrontgateClient{cc}
}

func (c *pilotServiceForFrontgateClient) Channel(ctx context.Context, opts ...grpc.CallOption) (PilotServiceForFrontgate_ChannelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PilotServiceForFrontgate_serviceDesc.Streams[0], c.cc, "/openpitrix.PilotServiceForFrontgate/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilotServiceForFrontgateChannelClient{stream}
	return x, nil
}

type PilotServiceForFrontgate_ChannelClient interface {
	Send(*google_protobuf2.BytesValue) error
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ClientStream
}

type pilotServiceForFrontgateChannelClient struct {
	grpc.ClientStream
}

func (x *pilotServiceForFrontgateChannelClient) Send(m *google_protobuf2.BytesValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pilotServiceForFrontgateChannelClient) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PilotServiceForFrontgate service

type PilotServiceForFrontgateServer interface {
	Channel(PilotServiceForFrontgate_ChannelServer) error
}

func RegisterPilotServiceForFrontgateServer(s *grpc.Server, srv PilotServiceForFrontgateServer) {
	s.RegisterService(&_PilotServiceForFrontgate_serviceDesc, srv)
}

func _PilotServiceForFrontgate_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PilotServiceForFrontgateServer).Channel(&pilotServiceForFrontgateChannelServer{stream})
}

type PilotServiceForFrontgate_ChannelServer interface {
	Send(*google_protobuf2.BytesValue) error
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ServerStream
}

type pilotServiceForFrontgateChannelServer struct {
	grpc.ServerStream
}

func (x *pilotServiceForFrontgateChannelServer) Send(m *google_protobuf2.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pilotServiceForFrontgateChannelServer) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PilotServiceForFrontgate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.PilotServiceForFrontgate",
	HandlerType: (*PilotServiceForFrontgateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _PilotServiceForFrontgate_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pilot.proto",
}

func init() { proto.RegisterFile("pilot.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0x54, 0x94, 0x49, 0x83, 0xaa, 0x55, 0x11, 0x26, 0x14, 0x08, 0x86, 0x43, 0x90,
	0xa8, 0x0d, 0x81, 0x03, 0x82, 0x53, 0x1b, 0xa9, 0x81, 0x1b, 0x8a, 0x2b, 0x0e, 0x5c, 0xa2, 0x8d,
	0x33, 0x98, 0x15, 0xd6, 0xee, 0xb2, 0x3b, 0x4e, 0xe1, 0x5a, 0x71, 0xe0, 0x0c, 0x9f, 0x06, 0x9f,
	0xc0, 0x9d, 0x5f, 0x40, 0x59, 0xdb, 0xe0, 0x24, 0x10, 0x45, 0xea, 0x29, 0xca, 0xcc, 0x7b, 0x6f,
	0xe7, 0xbd, 0x19, 0x43, 0x5b, 0x8b, 0x4c, 0x51, 0xa8, 0x8d, 0x22, 0xc5, 0x40, 0x69, 0x94, 0x5a,
	0x90, 0x11, 0x1f, 0xbb, 0xb7, 0x52, 0xa5, 0xd2, 0x0c, 0x23, 0xd7, 0x99, 0xe6, 0x6f, 0xa3, 0x33,
	0xc3, 0xb5, 0x46, 0x63, 0x0b, 0x6c, 0xf7, 0xa0, 0xec, 0x73, 0x2d, 0x22, 0x2e, 0xa5, 0x22, 0x4e,
	0x42, 0xc9, 0xaa, 0xfb, 0xc0, 0xfd, 0x24, 0x87, 0x29, 0xca, 0x43, 0x7b, 0xc6, 0xd3, 0x14, 0x4d,
	0xa4, 0xb4, 0x43, 0xac, 0xa3, 0x83, 0x1f, 0x1e, 0xec, 0xbf, 0xe0, 0x72, 0x96, 0x61, 0x9c, 0x4f,
	0x89, 0xdb, 0xf7, 0x63, 0xfc, 0x90, 0xa3, 0x25, 0xf6, 0x1c, 0xc0, 0x16, 0x95, 0x89, 0x98, 0xf9,
	0x5e, 0xcf, 0xeb, 0xb7, 0x07, 0x07, 0x61, 0xf1, 0x72, 0x58, 0x4d, 0x16, 0xc6, 0x64, 0x84, 0x4c,
	0x5f, 0xf3, 0x2c, 0xc7, 0x71, 0xab, 0xc4, 0xbf, 0x9c, 0xb1, 0x21, 0x5c, 0xa9, 0xc8, 0x3c, 0x59,
	0x3c, 0xe7, 0x37, 0xb6, 0x10, 0xe8, 0x94, 0x9c, 0x23, 0x47, 0x61, 0xcf, 0xa0, 0x35, 0x13, 0x06,
	0x13, 0x12, 0x73, 0xf4, 0x9b, 0xdb, 0x0c, 0xf0, 0x07, 0x1e, 0x9c, 0xc2, 0xd5, 0x15, 0x57, 0x56,
	0x2b, 0x69, 0xf1, 0x42, 0xb6, 0x82, 0xa7, 0x70, 0x6d, 0x84, 0x54, 0x4a, 0xc6, 0xc4, 0x29, 0xb7,
	0x55, 0x5c, 0x37, 0x57, 0x74, 0x9b, 0xfd, 0x56, 0x9d, 0x79, 0xee, 0x41, 0x67, 0x89, 0x77, 0xb1,
	0x7c, 0x9f, 0xc0, 0x8e, 0x75, 0x32, 0x5b, 0xe5, 0x5a, 0x62, 0x83, 0xcf, 0x1e, 0xf8, 0xeb, 0xf3,
	0x97, 0xc1, 0xdc, 0x86, 0x36, 0x29, 0xe2, 0xd9, 0x24, 0x51, 0xb9, 0x24, 0x37, 0x50, 0x67, 0x0c,
	0xae, 0x34, 0x5c, 0x54, 0xd8, 0x08, 0x58, 0x35, 0x70, 0xa1, 0x37, 0xb1, 0x48, 0x7e, 0xa3, 0xd7,
	0xec, 0xb7, 0x07, 0xd7, 0xc3, 0xbf, 0xe7, 0x1b, 0x2e, 0xeb, 0xef, 0xd9, 0xfa, 0xdf, 0x18, 0x69,
	0xf0, 0xcb, 0x83, 0xdd, 0x57, 0x8b, 0xd3, 0x8f, 0xd1, 0xcc, 0x45, 0x82, 0xec, 0x14, 0x3a, 0x4b,
	0xcb, 0x62, 0xbd, 0xba, 0xdc, 0xbf, 0xae, 0xb3, 0x7b, 0x67, 0x03, 0xa2, 0x34, 0xf4, 0xc5, 0x83,
	0xbd, 0x55, 0xb7, 0xec, 0x6e, 0x9d, 0xf7, 0x9f, 0x5d, 0x76, 0xef, 0x6d, 0x06, 0x15, 0xfa, 0xc1,
	0xfd, 0xaf, 0x47, 0xfb, 0x8c, 0xa5, 0x48, 0xbd, 0xd2, 0x61, 0xaf, 0x88, 0xe5, 0xfc, 0xfb, 0xcf,
	0x6f, 0x8d, 0x5d, 0x06, 0xd1, 0xfc, 0x51, 0xe4, 0xbe, 0x70, 0x3b, 0x48, 0xc0, 0xaf, 0x1b, 0x3e,
	0x51, 0xe6, 0xc4, 0x28, 0x49, 0x29, 0x27, 0x64, 0x23, 0xb8, 0x3c, 0x7c, 0xc7, 0xa5, 0xc4, 0x8c,
	0xdd, 0x58, 0xdb, 0xe2, 0xf1, 0x27, 0x42, 0xeb, 0x96, 0xd8, 0xdd, 0xd4, 0xec, 0x7b, 0x0f, 0xbd,
	0xe3, 0x4b, 0x6f, 0x1a, 0x7a, 0x3a, 0xdd, 0x71, 0xed, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe2, 0xaa, 0xad, 0x86, 0x5d, 0x04, 0x00, 0x00,
}
