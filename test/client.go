package main

import (
	"chat_room/api/conn"
	"chat_room/pkg/datapack"
	"chat_room/pkg/token"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// func client(url string) {
// 	//新建客户端

// 	log.Println(url)
// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer conn.Close()
// 	go Client(conn)
// 	// time.Sleep(5 * time.Second)

// 	//Stop here to prevent the program from exiting
// 	tick := time.NewTicker(10 * time.Second)
// 	for {
// 		<-tick.C
// 		return
// 	}
// }

func Client(url string, uid int, per int, roomId string) {

	var (
		req     = new(conn.Req)
		sign_in = new(conn.SignIn)
		// roomPush = new(conn.Req)
	)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println("启动链接失败:", err)
		return
	}
	defer c.Close()
	//根据uid生成token
	tokenS, err := token.GenToken(strconv.Itoa(uid))
	if err != nil {
		log.Println("生成token失败:", err)
		return
	}
	sign_in.Token = tokenS
	sign_in.RoomID = roomId
	sign_b, _ := proto.Marshal(sign_in)
	req.Type = 1
	req.Data = sign_b

	send_d, _ := proto.Marshal(req)

	c.WriteMessage(websocket.BinaryMessage, send_d)
	// if uid < per {
	go sendMsg(c, uid, roomId)
	// }
	messageHandle(c)
	// roomPush.Type = conn.PackageType_MESSAGE
	// sendMsg := fmt.Sprintf("hello im %d", uid)
	// msg := &conn.BroadcastRoomReq{
	// 	RoomID: roomId,
	// 	Data:   []byte(sendMsg),
	// }
	// msg_b, _ := proto.Marshal(msg)
	// roomPush.Data = msg_b
	// pushMsg_b, _ := proto.Marshal(roomPush)

	// time.Sleep(10 * time.Second)
	// c.WriteMessage(websocket.BinaryMessage, pushMsg_b)
	// tick := time.NewTicker(30 * time.Minute)
	// t1 := time.NewTicker(5 * time.Minute)
	// for {

	// c.WriteMessage(websocket.BinaryMessage, pushMsg_b)
	// t1.Reset(time.Duration(rand.Int63n(30)+10) * time.Second)
	// }

}

func sendMsg(c *websocket.Conn, uid int, roomId string) {
	var roomPush = new(conn.Req)
	roomPush.Type = conn.PackageType_MESSAGE
	sendMsg := fmt.Sprintf("hello im %d", uid)
	msg := &conn.BroadcastRoomReq{
		RoomID: roomId,
		MsgLv:  0,
		Data:   []byte(sendMsg),
	}
	msg_b, _ := proto.Marshal(msg)
	roomPush.Data = msg_b
	pushMsg_b, _ := proto.Marshal(roomPush)
	for {
		time.Sleep(1 * time.Second)
		c.WriteMessage(websocket.BinaryMessage, pushMsg_b)
	}
}

func messageHandle(c *websocket.Conn) {
	var (
		req      = new(conn.Req)
		roomPush = new(conn.BroadcastRoomReq)
		errR     = new(conn.Err)
	)
	for {

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read err %s:\n", err)
			return
		}
		proto.Unmarshal(message, req)
		switch req.Type {
		case conn.PackageType_UNKNOWN:
			//登录
			proto.Unmarshal(req.Data, errR)
			// log.Printf("sign in reply %+v\n", errR)
		case conn.PackageType_MESSAGE:
			//群推送
			proto.Unmarshal(req.Data, roomPush)
			log.Printf("recv msg %+v\n", roomPush)
			//解包
			d := roomPush.Data
			dp := datapack.NewDataPack()
			for len(d) > 0 {
				left, res, err := dp.Unpack(d)
				if err != nil {
					break
				}
				log.Printf("room push %s\n", res)
				d = left
			}
		case conn.PackageType_HEARTBEAT:
			//心跳

		}
	}

}
