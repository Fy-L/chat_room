package conns

import (
	pb_conn "chat_room/api/conn"
	pb_logic "chat_room/api/logic"
	"context"
	"fmt"

	"chat_room/pkg/logger"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

//用户连接
type Conn struct {
	//所属room
	UserId   uint32
	Nickname string
	Key      string
	Room     *Room
	Next     *Conn
	Prev     *Conn
	//ws连接
	ws    *websocket.Conn
	wLock sync.Mutex
}

// 获取实例
func NewConnect() *Conn {
	return &Conn{}
}

func (c *Conn) HandleMessage(b []byte) {
	// return
	//处理消息
	var req = new(pb_conn.Req)
	err := proto.Unmarshal(b, req)
	if err != nil {
		log.Printf("unmarshal err:%s\n", err)

		return
	}
	// fmt.Println("-----------------------")
	// fmt.Printf("%+v\n", req)
	if req.Type != pb_conn.PackageType_SIGN_IN && c.UserId == 0 {
		//未登录
		var (
			reply pb_conn.Reply
			e     pb_conn.Err
		)
		reply.Type = pb_conn.PackageType_UNKNOWN
		e.Code = 100
		e.Msg = "请登录"
		reply.Data, _ = proto.Marshal(&e)
		b, _ := proto.Marshal(&reply)
		c.WriteMsg(b)
		return
	}

	switch req.Type {
	case pb_conn.PackageType_SIGN_IN:
		//登录
		c.signIn(req.Data)
	case pb_conn.PackageType_MESSAGE:
		//群推送
		c.roomPush(req.Data)
	case pb_conn.PackageType_HEARTBEAT:
		//心跳
		c.heartbeat(req.Type)
	case pb_conn.PackageType_MEMBERS:
		//获取在线人数
		c.onlineMembers()
	}
}

//写信息回ws
func (c *Conn) WriteMsg(b []byte) {
	c.wLock.Lock()
	defer c.wLock.Unlock()
	c.ws.WriteMessage(websocket.BinaryMessage, b)
}

//登录
func (c *Conn) signIn(b []byte) {
	//调用logic的rpc接口signin
	var (
		signIn = new(pb_conn.SignIn)
		resp   = new(pb_conn.Reply)
		errRep = new(pb_conn.Err)
	)
	err := proto.Unmarshal(b, signIn)
	if err != nil {
		logger.Sugar.Errorf("unmarshal conn.signIn err :%s\n", err)
		return
	}
	resp.Type = pb_conn.PackageType_UNKNOWN
	errRep.Code = 0
	errRep.Msg = "suc"
	reply, err := SrvMrg.logicRpcClient.Auth(context.Background(), &pb_logic.AuthReq{Token: signIn.Token})
	if err != nil {
		errRep.Code = 100
		errRep.Msg = err.Error()
		//断开链接
		// c.Close()
	} else {
		c.UserId = uint32(reply.Uid)
		c.Nickname = reply.Nickname
		c.Key = uuid.New().String()
		//分配bucket
		bucket := SrvMrg.Bucket(uint32(reply.Uid))
		bucket.Put(signIn.RoomID, c)
	}
	err_b, _ := proto.Marshal(errRep)
	resp.Data = err_b
	resp_b, _ := proto.Marshal(resp)
	c.WriteMsg(resp_b)
}

func (c *Conn) heartbeat(pt pb_conn.PackageType) {
	var resp = new(pb_conn.Reply)
	resp.Type = pt
	resp.Data = []byte("pong")
	resp_b, err := proto.Marshal(resp)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	c.WriteMsg(resp_b)
}

//房间推送
func (c *Conn) roomPush(b []byte) {
	//TODO::调用job的rpc接口roomPush
	var (
		req = new(pb_conn.BroadcastRoomReq)
		// logicReq = new(pb_logic.PushMsgReq)
		err error
	)
	err = proto.Unmarshal(b, req)
	if err != nil {
		logger.Sugar.Errorf("unmarshal conn.roomPush err :%s\n", err)
		return
	}
	// logicReq.RoomID = req.RoomID
	// logicReq.Data = req.Data
	_, err = SrvMrg.logicRpcClient.PushMsg(context.Background(), req)
	// fmt.Println(err)
	if err != nil {
		logger.Sugar.Errorf(" jobRpc.PushMsg err :%s\n", err)
		return
	}

}

// 获取在线人数
func (c *Conn) onlineMembers() {
	var (
		n    = uint32(0)
		resp = new(pb_conn.Reply)
	)
	//获取当前roomid
	roomID := c.Room.id
	for _, b := range SrvMrg.buckets {
		n += b.OnlineMembers(roomID)
	}
	resp.Type = pb_conn.PackageType_MEMBERS
	resp.Data = []byte(fmt.Sprintf("%d", n))
	resp_b, _ := proto.Marshal(resp)
	c.WriteMsg(resp_b)
}

func (c *Conn) Close() {
	c.ws.Close()
}
