// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.2
// source: conn/conn.proto

package conn

import (
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

type PackageType int32

const (
	PackageType_UNKNOWN   PackageType = 0 // 未知
	PackageType_SIGN_IN   PackageType = 1 // 登录
	PackageType_HEARTBEAT PackageType = 2 // 心跳
	PackageType_MESSAGE   PackageType = 3 // 消息投递
	PackageType_MEMBERS   PackageType = 4 //获取在线人数
)

// Enum value maps for PackageType.
var (
	PackageType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SIGN_IN",
		2: "HEARTBEAT",
		3: "MESSAGE",
		4: "MEMBERS",
	}
	PackageType_value = map[string]int32{
		"UNKNOWN":   0,
		"SIGN_IN":   1,
		"HEARTBEAT": 2,
		"MESSAGE":   3,
		"MEMBERS":   4,
	}
)

func (x PackageType) Enum() *PackageType {
	p := new(PackageType)
	*p = x
	return p
}

func (x PackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_conn_proto_enumTypes[0].Descriptor()
}

func (PackageType) Type() protoreflect.EnumType {
	return &file_conn_conn_proto_enumTypes[0]
}

func (x PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageType.Descriptor instead.
func (PackageType) EnumDescriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{0}
}

type MsgLevel int32

const (
	MsgLevel_NORMAL    MsgLevel = 0 //普通消息
	MsgLevel_IMPORTANT MsgLevel = 1 //重要消息
)

// Enum value maps for MsgLevel.
var (
	MsgLevel_name = map[int32]string{
		0: "NORMAL",
		1: "IMPORTANT",
	}
	MsgLevel_value = map[string]int32{
		"NORMAL":    0,
		"IMPORTANT": 1,
	}
)

func (x MsgLevel) Enum() *MsgLevel {
	p := new(MsgLevel)
	*p = x
	return p
}

func (x MsgLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_conn_proto_enumTypes[1].Descriptor()
}

func (MsgLevel) Type() protoreflect.EnumType {
	return &file_conn_conn_proto_enumTypes[1]
}

func (x MsgLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgLevel.Descriptor instead.
func (MsgLevel) EnumDescriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{1}
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=conn.PackageType" json:"type,omitempty"` //类型
	Data []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetType() PackageType {
	if x != nil {
		return x.Type
	}
	return PackageType_UNKNOWN
}

func (x *Req) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=conn.PackageType" json:"type,omitempty"`
	Data []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetType() PackageType {
	if x != nil {
		return x.Type
	}
	return PackageType_UNKNOWN
}

func (x *Reply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Err struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` //错误码 0 代表ok 100需要登录
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    //错误信息
}

func (x *Err) Reset() {
	*x = Err{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Err) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Err) ProtoMessage() {}

func (x *Err) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Err.ProtoReflect.Descriptor instead.
func (*Err) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{2}
}

func (x *Err) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Err) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SignIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoomID string `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *SignIn) Reset() {
	*x = SignIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignIn) ProtoMessage() {}

func (x *SignIn) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignIn.ProtoReflect.Descriptor instead.
func (*SignIn) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{3}
}

func (x *SignIn) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignIn) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type BroadcastRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string   `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	MsgLv  MsgLevel `protobuf:"varint,2,opt,name=msgLv,proto3,enum=conn.MsgLevel" json:"msgLv,omitempty"`
	Data   []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BroadcastRoomReq) Reset() {
	*x = BroadcastRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRoomReq) ProtoMessage() {}

func (x *BroadcastRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRoomReq.ProtoReflect.Descriptor instead.
func (*BroadcastRoomReq) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastRoomReq) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *BroadcastRoomReq) GetMsgLv() MsgLevel {
	if x != nil {
		return x.MsgLv
	}
	return MsgLevel_NORMAL
}

func (x *BroadcastRoomReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BroadcastRoomReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastRoomReply) Reset() {
	*x = BroadcastRoomReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_conn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRoomReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRoomReply) ProtoMessage() {}

func (x *BroadcastRoomReply) ProtoReflect() protoreflect.Message {
	mi := &file_conn_conn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRoomReply.ProtoReflect.Descriptor instead.
func (*BroadcastRoomReply) Descriptor() ([]byte, []int) {
	return file_conn_conn_proto_rawDescGZIP(), []int{5}
}

var File_conn_conn_proto protoreflect.FileDescriptor

var file_conn_conn_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x6e, 0x22, 0x40, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x25,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a,
	0x03, 0x45, 0x72, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x22, 0x64, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x24,
	0x0a, 0x05, 0x6d, 0x73, 0x67, 0x4c, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6d,
	0x73, 0x67, 0x4c, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x50,
	0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49,
	0x47, 0x4e, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54,
	0x42, 0x45, 0x41, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x10, 0x04,
	0x2a, 0x25, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x50, 0x4f,
	0x52, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x32, 0x49, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x41, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x14, 0x5a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conn_conn_proto_rawDescOnce sync.Once
	file_conn_conn_proto_rawDescData = file_conn_conn_proto_rawDesc
)

func file_conn_conn_proto_rawDescGZIP() []byte {
	file_conn_conn_proto_rawDescOnce.Do(func() {
		file_conn_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_conn_conn_proto_rawDescData)
	})
	return file_conn_conn_proto_rawDescData
}

var file_conn_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conn_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_conn_conn_proto_goTypes = []interface{}{
	(PackageType)(0),           // 0: conn.PackageType
	(MsgLevel)(0),              // 1: conn.MsgLevel
	(*Req)(nil),                // 2: conn.Req
	(*Reply)(nil),              // 3: conn.Reply
	(*Err)(nil),                // 4: conn.Err
	(*SignIn)(nil),             // 5: conn.SignIn
	(*BroadcastRoomReq)(nil),   // 6: conn.BroadcastRoomReq
	(*BroadcastRoomReply)(nil), // 7: conn.BroadcastRoomReply
}
var file_conn_conn_proto_depIdxs = []int32{
	0, // 0: conn.Req.type:type_name -> conn.PackageType
	0, // 1: conn.Reply.type:type_name -> conn.PackageType
	1, // 2: conn.BroadcastRoomReq.msgLv:type_name -> conn.MsgLevel
	6, // 3: conn.Conn.BroadcastRoom:input_type -> conn.BroadcastRoomReq
	7, // 4: conn.Conn.BroadcastRoom:output_type -> conn.BroadcastRoomReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conn_conn_proto_init() }
func file_conn_conn_proto_init() {
	if File_conn_conn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conn_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_conn_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_conn_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Err); i {
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
		file_conn_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignIn); i {
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
		file_conn_conn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRoomReq); i {
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
		file_conn_conn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRoomReply); i {
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
			RawDescriptor: file_conn_conn_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conn_conn_proto_goTypes,
		DependencyIndexes: file_conn_conn_proto_depIdxs,
		EnumInfos:         file_conn_conn_proto_enumTypes,
		MessageInfos:      file_conn_conn_proto_msgTypes,
	}.Build()
	File_conn_conn_proto = out.File
	file_conn_conn_proto_rawDesc = nil
	file_conn_conn_proto_goTypes = nil
	file_conn_conn_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnClient interface {
	// 广播
	//rpc Broadcast(BroadcastReq) returns (BroadcastReply);
	// 某个房间广播
	BroadcastRoom(ctx context.Context, in *BroadcastRoomReq, opts ...grpc.CallOption) (*BroadcastRoomReply, error)
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

func (c *connClient) BroadcastRoom(ctx context.Context, in *BroadcastRoomReq, opts ...grpc.CallOption) (*BroadcastRoomReply, error) {
	out := new(BroadcastRoomReply)
	err := c.cc.Invoke(ctx, "/conn.Conn/BroadcastRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnServer is the server API for Conn service.
type ConnServer interface {
	// 广播
	//rpc Broadcast(BroadcastReq) returns (BroadcastReply);
	// 某个房间广播
	BroadcastRoom(context.Context, *BroadcastRoomReq) (*BroadcastRoomReply, error)
}

// UnimplementedConnServer can be embedded to have forward compatible implementations.
type UnimplementedConnServer struct {
}

func (*UnimplementedConnServer) BroadcastRoom(context.Context, *BroadcastRoomReq) (*BroadcastRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastRoom not implemented")
}

func RegisterConnServer(s *grpc.Server, srv ConnServer) {
	s.RegisterService(&_Conn_serviceDesc, srv)
}

func _Conn_BroadcastRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).BroadcastRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conn.Conn/BroadcastRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).BroadcastRoom(ctx, req.(*BroadcastRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Conn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conn.Conn",
	HandlerType: (*ConnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastRoom",
			Handler:    _Conn_BroadcastRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conn/conn.proto",
}