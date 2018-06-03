// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ths.proto

/*
Package thspbs is a generated protocol buffer package.

It is generated from these files:
	ths.proto

It has these top-level messages:
	Event
	BaseInfo
	FriendInfo
	GroupInfo
	MemberInfo
	Message
	Messages
	EmptyReq
	HelloReq
	HelloResp
*/
package thspbs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type MemberInfo_MemType int32

const (
	MemberInfo_UNKNOWN MemberInfo_MemType = 0
	MemberInfo_FRIEND  MemberInfo_MemType = 1
	MemberInfo_GROUP   MemberInfo_MemType = 2
	MemberInfo_PEER    MemberInfo_MemType = 3
)

var MemberInfo_MemType_name = map[int32]string{
	0: "UNKNOWN",
	1: "FRIEND",
	2: "GROUP",
	3: "PEER",
}
var MemberInfo_MemType_value = map[string]int32{
	"UNKNOWN": 0,
	"FRIEND":  1,
	"GROUP":   2,
	"PEER":    3,
}

func (x MemberInfo_MemType) String() string {
	return proto.EnumName(MemberInfo_MemType_name, int32(x))
}
func (MemberInfo_MemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Event struct {
	EventId int64             `protobuf:"varint,1,opt,name=EventId" json:"EventId,omitempty"`
	Name    string            `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Args    []string          `protobuf:"bytes,3,rep,name=Args" json:"Args,omitempty"`
	Margs   []string          `protobuf:"bytes,4,rep,name=Margs" json:"Margs,omitempty"`
	Nargs   map[string]string `protobuf:"bytes,5,rep,name=Nargs" json:"Nargs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ErrCode int32             `protobuf:"varint,6,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg  string            `protobuf:"bytes,7,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventId() int64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Event) GetMargs() []string {
	if m != nil {
		return m.Margs
	}
	return nil
}

func (m *Event) GetNargs() map[string]string {
	if m != nil {
		return m.Nargs
	}
	return nil
}

func (m *Event) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Event) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type BaseInfo struct {
	ToxId      string                 `protobuf:"bytes,1,opt,name=ToxId" json:"ToxId,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Stmsg      string                 `protobuf:"bytes,3,opt,name=Stmsg" json:"Stmsg,omitempty"`
	Status     uint32                 `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	Friends    map[uint32]*FriendInfo `protobuf:"bytes,5,rep,name=Friends" json:"Friends,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Groups     map[uint32]*GroupInfo  `protobuf:"bytes,6,rep,name=Groups" json:"Groups,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ConnStatus int32                  `protobuf:"varint,7,opt,name=ConnStatus" json:"ConnStatus,omitempty"`
	NextBatch  int64                  `protobuf:"varint,8,opt,name=NextBatch" json:"NextBatch,omitempty"`
}

func (m *BaseInfo) Reset()                    { *m = BaseInfo{} }
func (m *BaseInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseInfo) ProtoMessage()               {}
func (*BaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BaseInfo) GetToxId() string {
	if m != nil {
		return m.ToxId
	}
	return ""
}

func (m *BaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BaseInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *BaseInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BaseInfo) GetFriends() map[uint32]*FriendInfo {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *BaseInfo) GetGroups() map[uint32]*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *BaseInfo) GetConnStatus() int32 {
	if m != nil {
		return m.ConnStatus
	}
	return 0
}

func (m *BaseInfo) GetNextBatch() int64 {
	if m != nil {
		return m.NextBatch
	}
	return 0
}

type FriendInfo struct {
	Fnum       uint32 `protobuf:"varint,1,opt,name=Fnum" json:"Fnum,omitempty"`
	Status     uint32 `protobuf:"varint,2,opt,name=Status" json:"Status,omitempty"`
	Pubkey     string `protobuf:"bytes,3,opt,name=Pubkey" json:"Pubkey,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=Name" json:"Name,omitempty"`
	Stmsg      string `protobuf:"bytes,5,opt,name=Stmsg" json:"Stmsg,omitempty"`
	Avatar     string `protobuf:"bytes,6,opt,name=Avatar" json:"Avatar,omitempty"`
	Seen       uint64 `protobuf:"varint,7,opt,name=Seen" json:"Seen,omitempty"`
	ConnStatus int32  `protobuf:"varint,8,opt,name=ConnStatus" json:"ConnStatus,omitempty"`
}

func (m *FriendInfo) Reset()                    { *m = FriendInfo{} }
func (m *FriendInfo) String() string            { return proto.CompactTextString(m) }
func (*FriendInfo) ProtoMessage()               {}
func (*FriendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FriendInfo) GetFnum() uint32 {
	if m != nil {
		return m.Fnum
	}
	return 0
}

func (m *FriendInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FriendInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *FriendInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FriendInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *FriendInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *FriendInfo) GetSeen() uint64 {
	if m != nil {
		return m.Seen
	}
	return 0
}

func (m *FriendInfo) GetConnStatus() int32 {
	if m != nil {
		return m.ConnStatus
	}
	return 0
}

type GroupInfo struct {
	Gnum    uint32                 `protobuf:"varint,1,opt,name=Gnum" json:"Gnum,omitempty"`
	Mtype   uint32                 `protobuf:"varint,2,opt,name=Mtype" json:"Mtype,omitempty"`
	GroupId string                 `protobuf:"bytes,3,opt,name=GroupId" json:"GroupId,omitempty"`
	Title   string                 `protobuf:"bytes,4,opt,name=Title" json:"Title,omitempty"`
	Stmsg   string                 `protobuf:"bytes,5,opt,name=Stmsg" json:"Stmsg,omitempty"`
	Ours    bool                   `protobuf:"varint,6,opt,name=Ours" json:"Ours,omitempty"`
	Members map[uint32]*MemberInfo `protobuf:"bytes,7,rep,name=Members" json:"Members,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GroupInfo) Reset()                    { *m = GroupInfo{} }
func (m *GroupInfo) String() string            { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()               {}
func (*GroupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupInfo) GetGnum() uint32 {
	if m != nil {
		return m.Gnum
	}
	return 0
}

func (m *GroupInfo) GetMtype() uint32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *GroupInfo) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *GroupInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GroupInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *GroupInfo) GetOurs() bool {
	if m != nil {
		return m.Ours
	}
	return false
}

func (m *GroupInfo) GetMembers() map[uint32]*MemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

// = ContactInfo
// 可用于friend,group,peer
type MemberInfo struct {
	Pnum   uint32             `protobuf:"varint,1,opt,name=Pnum" json:"Pnum,omitempty"`
	Pubkey string             `protobuf:"bytes,2,opt,name=Pubkey" json:"Pubkey,omitempty"`
	Name   string             `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Mtype  MemberInfo_MemType `protobuf:"varint,4,opt,name=Mtype,enum=thspbs.MemberInfo_MemType" json:"Mtype,omitempty"`
}

func (m *MemberInfo) Reset()                    { *m = MemberInfo{} }
func (m *MemberInfo) String() string            { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()               {}
func (*MemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MemberInfo) GetPnum() uint32 {
	if m != nil {
		return m.Pnum
	}
	return 0
}

func (m *MemberInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *MemberInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberInfo) GetMtype() MemberInfo_MemType {
	if m != nil {
		return m.Mtype
	}
	return MemberInfo_UNKNOWN
}

type Message struct {
	MsgId   uint64      `protobuf:"varint,1,opt,name=MsgId" json:"MsgId,omitempty"`
	Content string      `protobuf:"bytes,2,opt,name=Content" json:"Content,omitempty"`
	Peer    *MemberInfo `protobuf:"bytes,3,opt,name=Peer" json:"Peer,omitempty"`
	Created uint64      `protobuf:"varint,4,opt,name=Created" json:"Created,omitempty"`
	Updated uint64      `protobuf:"varint,5,opt,name=Updated" json:"Updated,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Message) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetPeer() *MemberInfo {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *Message) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Message) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type Messages struct {
	Msgs []*Message `protobuf:"bytes,1,rep,name=Msgs" json:"Msgs,omitempty"`
}

func (m *Messages) Reset()                    { *m = Messages{} }
func (m *Messages) String() string            { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()               {}
func (*Messages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Messages) GetMsgs() []*Message {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type EmptyReq struct {
}

func (m *EmptyReq) Reset()                    { *m = EmptyReq{} }
func (m *EmptyReq) String() string            { return proto.CompactTextString(m) }
func (*EmptyReq) ProtoMessage()               {}
func (*EmptyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type HelloReq struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type HelloResp struct {
	Code   int32 `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Status int64 `protobuf:"varint,2,opt,name=Status" json:"Status,omitempty"`
}

func (m *HelloResp) Reset()                    { *m = HelloResp{} }
func (m *HelloResp) String() string            { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()               {}
func (*HelloResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *HelloResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HelloResp) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "thspbs.Event")
	proto.RegisterType((*BaseInfo)(nil), "thspbs.BaseInfo")
	proto.RegisterType((*FriendInfo)(nil), "thspbs.FriendInfo")
	proto.RegisterType((*GroupInfo)(nil), "thspbs.GroupInfo")
	proto.RegisterType((*MemberInfo)(nil), "thspbs.MemberInfo")
	proto.RegisterType((*Message)(nil), "thspbs.Message")
	proto.RegisterType((*Messages)(nil), "thspbs.Messages")
	proto.RegisterType((*EmptyReq)(nil), "thspbs.EmptyReq")
	proto.RegisterType((*HelloReq)(nil), "thspbs.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "thspbs.HelloResp")
	proto.RegisterEnum("thspbs.MemberInfo_MemType", MemberInfo_MemType_name, MemberInfo_MemType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Toxhs service

type ToxhsClient interface {
	PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error)
	GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error)
	RmtCall(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Ping(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyReq, error)
}

type toxhsClient struct {
	cc *grpc.ClientConn
}

func NewToxhsClient(cc *grpc.ClientConn) ToxhsClient {
	return &toxhsClient{cc}
}

func (c *toxhsClient) PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Toxhs_serviceDesc.Streams[0], c.cc, "/thspbs.Toxhs/PollCallback", opts...)
	if err != nil {
		return nil, err
	}
	x := &toxhsPollCallbackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Toxhs_PollCallbackClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type toxhsPollCallbackClient struct {
	grpc.ClientStream
}

func (x *toxhsPollCallbackClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toxhsClient) GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error) {
	out := new(BaseInfo)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/GetBaseInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toxhsClient) RmtCall(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/RmtCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toxhsClient) Ping(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyReq, error) {
	out := new(EmptyReq)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Toxhs service

type ToxhsServer interface {
	PollCallback(*EmptyReq, Toxhs_PollCallbackServer) error
	GetBaseInfo(context.Context, *EmptyReq) (*BaseInfo, error)
	RmtCall(context.Context, *Event) (*Event, error)
	Ping(context.Context, *EmptyReq) (*EmptyReq, error)
}

func RegisterToxhsServer(s *grpc.Server, srv ToxhsServer) {
	s.RegisterService(&_Toxhs_serviceDesc, srv)
}

func _Toxhs_PollCallback_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ToxhsServer).PollCallback(m, &toxhsPollCallbackServer{stream})
}

type Toxhs_PollCallbackServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type toxhsPollCallbackServer struct {
	grpc.ServerStream
}

func (x *toxhsPollCallbackServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Toxhs_GetBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).GetBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/GetBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).GetBaseInfo(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Toxhs_RmtCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).RmtCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/RmtCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).RmtCall(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Toxhs_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).Ping(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Toxhs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Toxhs",
	HandlerType: (*ToxhsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBaseInfo",
			Handler:    _Toxhs_GetBaseInfo_Handler,
		},
		{
			MethodName: "RmtCall",
			Handler:    _Toxhs_RmtCall_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Toxhs_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PollCallback",
			Handler:       _Toxhs_PollCallback_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ths.proto",
}

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHellox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *EmptyReq) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(context.Context, *HelloReq) (*HelloReq, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHellox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHellox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHellox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHellox(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHellox",
			Handler:    _Greeter_SayHellox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ths.proto",
}

func init() { proto.RegisterFile("ths.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xee, 0x8a, 0xa4, 0x44, 0x8e, 0xe2, 0x56, 0x5d, 0x04, 0x01, 0x21, 0xa4, 0x81, 0xc0, 0x02,
	0x8d, 0xd0, 0x03, 0xe3, 0xa8, 0x05, 0x6c, 0xf4, 0x96, 0xb8, 0xb2, 0x6b, 0xb4, 0xa2, 0x89, 0x75,
	0x8c, 0x9e, 0xa9, 0x68, 0x2b, 0xbb, 0xa1, 0x48, 0x86, 0xbb, 0x32, 0xac, 0x07, 0xe9, 0xa1, 0x0f,
	0xd1, 0x73, 0xef, 0x3d, 0xf4, 0xa5, 0x7a, 0x29, 0x76, 0x76, 0x57, 0xa2, 0x7e, 0x6a, 0xa0, 0xb7,
	0xf9, 0x66, 0x76, 0x66, 0x67, 0xbe, 0x6f, 0x7f, 0x20, 0x90, 0xb7, 0x22, 0xae, 0xea, 0x52, 0x96,
	0xb4, 0x2d, 0x6f, 0x45, 0x35, 0x15, 0xd1, 0x3f, 0x04, 0xbc, 0xf1, 0x3d, 0x2f, 0x24, 0x0d, 0xa1,
	0x83, 0xc6, 0xe5, 0x2c, 0x24, 0x03, 0x32, 0x74, 0x98, 0x85, 0x94, 0x82, 0x9b, 0x64, 0x0b, 0x1e,
	0xb6, 0x06, 0x64, 0x18, 0x30, 0xb4, 0x95, 0xef, 0x4d, 0x3d, 0x17, 0xa1, 0x33, 0x70, 0x94, 0x4f,
	0xd9, 0xf4, 0x29, 0x78, 0x93, 0x4c, 0x39, 0x5d, 0x74, 0x6a, 0x40, 0x63, 0xf0, 0x12, 0xf4, 0x7a,
	0x03, 0x67, 0xd8, 0x1d, 0x85, 0xb1, 0xde, 0x39, 0xc6, 0xea, 0x31, 0x86, 0xc6, 0x85, 0xac, 0x57,
	0x4c, 0x2f, 0xc3, 0x3e, 0xea, 0xfa, 0xac, 0x9c, 0xf1, 0xb0, 0x3d, 0x20, 0x43, 0x8f, 0x59, 0x48,
	0x9f, 0x41, 0x7b, 0x5c, 0xd7, 0x13, 0x31, 0x0f, 0x3b, 0xd8, 0x89, 0x41, 0xfd, 0x53, 0x80, 0x4d,
	0x19, 0xda, 0x03, 0xe7, 0x03, 0x5f, 0xe1, 0x0c, 0x01, 0x53, 0xa6, 0xea, 0xeb, 0x3e, 0xcb, 0x97,
	0x76, 0x00, 0x0d, 0xbe, 0x6b, 0x9d, 0x92, 0xe8, 0x0f, 0x07, 0xfc, 0xb7, 0x99, 0xe0, 0x97, 0xc5,
	0x2f, 0xa5, 0x5a, 0xf6, 0xae, 0x7c, 0x30, 0xe3, 0x07, 0x4c, 0x83, 0x83, 0xc3, 0x3f, 0x05, 0xef,
	0x5a, 0x2e, 0xc4, 0x3c, 0x74, 0xf4, 0x4a, 0x04, 0xaa, 0xbd, 0x6b, 0x99, 0xc9, 0xa5, 0x9a, 0x9f,
	0x0c, 0x8f, 0x98, 0x41, 0xf4, 0x04, 0x3a, 0xe7, 0xf5, 0x1d, 0x2f, 0x66, 0x96, 0x82, 0x2f, 0x2c,
	0x05, 0x76, 0xeb, 0xd8, 0xc4, 0x35, 0x0f, 0x76, 0x35, 0xfd, 0x16, 0xda, 0x17, 0x75, 0xb9, 0xac,
	0x44, 0xd8, 0xc6, 0xbc, 0xe7, 0x7b, 0x79, 0x3a, 0xac, 0xd3, 0xcc, 0x5a, 0xfa, 0x02, 0xe0, 0xac,
	0x2c, 0x0a, 0xd3, 0x4a, 0x07, 0x29, 0x6c, 0x78, 0xe8, 0x73, 0x08, 0x12, 0xfe, 0x20, 0xdf, 0x66,
	0xf2, 0xfd, 0x6d, 0xe8, 0xa3, 0xd2, 0x1b, 0x47, 0x3f, 0x81, 0x27, 0xcd, 0x66, 0x9a, 0x6c, 0x1e,
	0x69, 0x36, 0x87, 0x4d, 0x36, 0xbb, 0x23, 0x6a, 0x9b, 0xd2, 0x69, 0xaa, 0xad, 0x06, 0xc3, 0xfd,
	0x9f, 0xa0, 0xdb, 0x68, 0xf2, 0x40, 0xb9, 0x97, 0xdb, 0xe5, 0x3e, 0xb7, 0xe5, 0x30, 0x6b, 0xa7,
	0x5a, 0xf4, 0x37, 0x01, 0xd8, 0xec, 0xa3, 0xb4, 0x39, 0x2f, 0x96, 0x0b, 0x53, 0x0e, 0xed, 0x86,
	0x0a, 0xad, 0x2d, 0x15, 0x9e, 0x41, 0x3b, 0x5d, 0x4e, 0xd5, 0xe6, 0x5a, 0x34, 0x83, 0xd6, 0xfa,
	0xba, 0x87, 0xf4, 0xf5, 0x76, 0xf4, 0x7d, 0x73, 0x9f, 0xc9, 0xac, 0xc6, 0x73, 0x19, 0x30, 0x83,
	0x54, 0x85, 0x6b, 0xce, 0x0b, 0xa4, 0xda, 0x65, 0x68, 0xef, 0x88, 0xe0, 0xef, 0x8a, 0x10, 0xfd,
	0xde, 0x82, 0x60, 0x3d, 0xa1, 0xaa, 0x70, 0xd1, 0x98, 0x43, 0xd9, 0x78, 0x99, 0xe4, 0xaa, 0xe2,
	0x66, 0x0c, 0x0d, 0xd4, 0xe5, 0xd0, 0x69, 0x33, 0x33, 0x86, 0x85, 0x78, 0x7a, 0xef, 0x64, 0x6e,
	0x07, 0xd1, 0xe0, 0x3f, 0x26, 0xa1, 0xe0, 0x5e, 0x2d, 0x6b, 0x81, 0x73, 0xf8, 0x0c, 0x6d, 0x7a,
	0x0a, 0x9d, 0x09, 0x5f, 0x4c, 0x79, 0xad, 0xce, 0x8c, 0x3a, 0x6d, 0x2f, 0xf6, 0x94, 0x88, 0xcd,
	0x02, 0x73, 0x4c, 0x0d, 0x52, 0x47, 0xa6, 0x19, 0xf8, 0x1f, 0x47, 0x46, 0xa7, 0xed, 0x8a, 0xfc,
	0x27, 0x01, 0xd8, 0x44, 0x54, 0xb3, 0x69, 0x83, 0x9c, 0xd4, 0x88, 0x6c, 0xc4, 0x6c, 0x1d, 0x14,
	0xd3, 0x69, 0x88, 0x79, 0x6c, 0x89, 0x54, 0xc4, 0x7c, 0x3a, 0xea, 0xef, 0x6f, 0xae, 0xcc, 0x77,
	0xab, 0x8a, 0x1b, 0x92, 0xa3, 0x13, 0xa4, 0x42, 0x79, 0x68, 0x17, 0x3a, 0x37, 0xc9, 0x8f, 0xc9,
	0xd5, 0xcf, 0x49, 0xef, 0x13, 0x0a, 0xd0, 0x3e, 0x67, 0x97, 0xe3, 0xe4, 0xfb, 0x1e, 0xa1, 0x01,
	0x78, 0x17, 0xec, 0xea, 0x26, 0xed, 0xb5, 0xa8, 0x0f, 0x6e, 0x3a, 0x1e, 0xb3, 0x9e, 0x13, 0xfd,
	0x46, 0x54, 0xa6, 0x10, 0xd9, 0x1c, 0x99, 0x9f, 0x88, 0xb9, 0x79, 0x4d, 0x5c, 0xa6, 0x81, 0xd2,
	0xef, 0xac, 0x2c, 0x24, 0x2f, 0xa4, 0xe9, 0xdc, 0x42, 0xfa, 0x15, 0xb8, 0x29, 0xe7, 0x35, 0xb6,
	0x7e, 0x98, 0x22, 0x8c, 0x63, 0x85, 0x9a, 0x67, 0x92, 0xcf, 0x70, 0x20, 0x97, 0x59, 0xa8, 0x22,
	0x37, 0xd5, 0x0c, 0x23, 0x9e, 0x8e, 0x18, 0x18, 0xbd, 0x02, 0xdf, 0xb4, 0x25, 0xe8, 0x97, 0xe0,
	0x4e, 0xc4, 0x5c, 0x84, 0x04, 0x45, 0xfe, 0x6c, 0xb3, 0x0f, 0xc6, 0x19, 0x06, 0x23, 0x00, 0x7f,
	0xbc, 0xa8, 0xe4, 0x8a, 0xf1, 0x8f, 0xd1, 0x31, 0xf8, 0x3f, 0xf0, 0x3c, 0x2f, 0x19, 0xff, 0xb8,
	0xe6, 0x97, 0x34, 0xf8, 0xed, 0x81, 0xa3, 0x9e, 0x64, 0x3d, 0x8e, 0x32, 0xa3, 0x13, 0x08, 0x4c,
	0x86, 0xa8, 0x54, 0x0a, 0xbe, 0xe5, 0x04, 0xef, 0x80, 0x6b, 0x1f, 0xf2, 0xc6, 0x1d, 0x75, 0xec,
	0x1d, 0x1d, 0xfd, 0x45, 0xf0, 0x09, 0xbe, 0x15, 0xf4, 0x35, 0x3c, 0x49, 0xcb, 0x3c, 0x3f, 0xcb,
	0xf2, 0x7c, 0x9a, 0xbd, 0xff, 0x40, 0x7b, 0xeb, 0x5f, 0xc3, 0xb4, 0xd5, 0x3f, 0xda, 0xfa, 0x47,
	0x8e, 0x09, 0x7d, 0x0d, 0xdd, 0x0b, 0x2e, 0xd7, 0xaf, 0xf9, 0x7e, 0x46, 0x6f, 0xf7, 0xf9, 0xa4,
	0x2f, 0xa1, 0xc3, 0x16, 0x52, 0x6d, 0x42, 0xb7, 0xcb, 0xed, 0x54, 0xa7, 0x5f, 0x83, 0x9b, 0xde,
	0x15, 0xf3, 0xc7, 0x8a, 0x5a, 0xcf, 0xe8, 0x57, 0x75, 0x45, 0x39, 0x97, 0xbc, 0xa6, 0x31, 0xf8,
	0xd7, 0xd9, 0x0a, 0xb9, 0x78, 0x2c, 0x75, 0x4d, 0xef, 0x2b, 0x08, 0xec, 0xfa, 0x07, 0xba, 0x17,
	0xde, 0x4f, 0x98, 0xb6, 0xf1, 0x33, 0xff, 0xe6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x5a,
	0xb7, 0x36, 0xd9, 0x07, 0x00, 0x00,
}
