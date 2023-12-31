// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.2
// source: logic/logic.proto

package logic

import (
	conn "chat_room/api/conn"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RoomID   string `protobuf:"bytes,3,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{1}
}

func (x *AuthReply) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AuthReply) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AuthReply) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

var File_logic_logic_proto protoreflect.FileDescriptor

var file_logic_logic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x09,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x32,
	0x6e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x12, 0x28, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x15, 0x5a, 0x13, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logic_logic_proto_rawDescOnce sync.Once
	file_logic_logic_proto_rawDescData = file_logic_logic_proto_rawDesc
)

func file_logic_logic_proto_rawDescGZIP() []byte {
	file_logic_logic_proto_rawDescOnce.Do(func() {
		file_logic_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_logic_logic_proto_rawDescData)
	})
	return file_logic_logic_proto_rawDescData
}

var file_logic_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logic_logic_proto_goTypes = []interface{}{
	(*AuthReq)(nil),                 // 0: logic.AuthReq
	(*AuthReply)(nil),               // 1: logic.AuthReply
	(*conn.BroadcastRoomReq)(nil),   // 2: conn.BroadcastRoomReq
	(*conn.BroadcastRoomReply)(nil), // 3: conn.BroadcastRoomReply
}
var file_logic_logic_proto_depIdxs = []int32{
	0, // 0: logic.Logic.Auth:input_type -> logic.AuthReq
	2, // 1: logic.Logic.PushMsg:input_type -> conn.BroadcastRoomReq
	1, // 2: logic.Logic.Auth:output_type -> logic.AuthReply
	3, // 3: logic.Logic.PushMsg:output_type -> conn.BroadcastRoomReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logic_logic_proto_init() }
func file_logic_logic_proto_init() {
	if File_logic_logic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logic_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logic_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logic_logic_proto_goTypes,
		DependencyIndexes: file_logic_logic_proto_depIdxs,
		MessageInfos:      file_logic_logic_proto_msgTypes,
	}.Build()
	File_logic_logic_proto = out.File
	file_logic_logic_proto_rawDesc = nil
	file_logic_logic_proto_goTypes = nil
	file_logic_logic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicClient interface {
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error)
	PushMsg(ctx context.Context, in *conn.BroadcastRoomReq, opts ...grpc.CallOption) (*conn.BroadcastRoomReply, error)
}

type logicClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicClient(cc grpc.ClientConnInterface) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/logic.Logic/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) PushMsg(ctx context.Context, in *conn.BroadcastRoomReq, opts ...grpc.CallOption) (*conn.BroadcastRoomReply, error) {
	out := new(conn.BroadcastRoomReply)
	err := c.cc.Invoke(ctx, "/logic.Logic/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
type LogicServer interface {
	Auth(context.Context, *AuthReq) (*AuthReply, error)
	PushMsg(context.Context, *conn.BroadcastRoomReq) (*conn.BroadcastRoomReply, error)
}

// UnimplementedLogicServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (*UnimplementedLogicServer) Auth(context.Context, *AuthReq) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedLogicServer) PushMsg(context.Context, *conn.BroadcastRoomReq) (*conn.BroadcastRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}

func RegisterLogicServer(s *grpc.Server, srv LogicServer) {
	s.RegisterService(&_Logic_serviceDesc, srv)
}

func _Logic_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(conn.BroadcastRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).PushMsg(ctx, req.(*conn.BroadcastRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logic.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Logic_Auth_Handler,
		},
		{
			MethodName: "PushMsg",
			Handler:    _Logic_PushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic/logic.proto",
}
