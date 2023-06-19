package job

import (
	pb_conn "chat_room/api/conn"
	"context"
)

// type Job struct {
// 	Srv *Server
// }

//job对外的rpc接口，用于对外rpc的消息推送。已经nsq替代
func (j *Job) PushMsg(ctx context.Context, req *pb_conn.BroadcastRoomReq) (*pb_conn.BroadcastRoomReply, error) {
	err := j.push(req)
	if err != nil {
		return nil, err
	}
	return &pb_conn.BroadcastRoomReply{}, nil
}
