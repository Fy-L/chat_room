package conns

import (
	"chat_room/api/conn"
	"context"
	"errors"
)

type ConnSrv struct{}

//广播到某个房间
func (c *ConnSrv) BroadcastRoom(ctx context.Context, req *conn.BroadcastRoomReq) (*conn.BroadcastRoomReply, error) {
	if req.RoomID == "" || req.Data == nil {
		return nil, errors.New("缺少roomID或data参数")
	}
	// fmt.Printf("%+v\n", req)
	go func() {
		for _, bucket := range SrvMrg.Buckets() {
			bucket.BroadcastRoom(req)
		}
	}()
	return &conn.BroadcastRoomReply{}, nil
}
