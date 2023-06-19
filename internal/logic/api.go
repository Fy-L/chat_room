package logic

import (
	pb_conn "chat_room/api/conn"
	"chat_room/api/logic"
	"chat_room/pkg/logger"
	"chat_room/pkg/token"
	"context"
	"errors"

	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
	"google.golang.org/protobuf/proto"
)

type LogicSrv struct {
	producer *nsq.Producer
	topic    string
}

func NewLogicSrv(nsqAddr string, topic string) (*LogicSrv, error) {
	l := &LogicSrv{}
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqAddr, config)
	if err != nil {
		log.Printf("create producer failed. err :%+v\n", err)
		return nil, err
	}
	l.producer = p
	l.topic = topic
	return l, nil
}

var _ logic.LogicServer = (*LogicSrv)(nil)

func (s *LogicSrv) Auth(ctx context.Context, req *logic.AuthReq) (*logic.AuthReply, error) {
	var reply = new(logic.AuthReply)
	var err error
	// fmt.Println("rpc token=", req.Token)
	//解析token
	claims, err := token.ParseToken(req.Token)
	if err != nil {
		return reply, err
	}
	now := time.Now().Unix()
	if claims.Iat > now {
		err = errors.New("无效token")
		return reply, err
	}
	if claims.Exp < now {
		err = errors.New("登录信息过期")
		return reply, err
	}

	uid, _ := strconv.Atoi(claims.Uid)
	// fmt.Print(uid)
	reply.Nickname = claims.Uid
	reply.Uid = int32(uid)
	return reply, err
}

func (s *LogicSrv) PushMsg(ctx context.Context, req *pb_conn.BroadcastRoomReq) (*pb_conn.BroadcastRoomReply, error) {
	var reply = new(pb_conn.BroadcastRoomReply)
	var err error
	//将消息推送到 nsq中
	data_b, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	s.publish(data_b)
	return reply, err
}

func (s *LogicSrv) publish(data []byte) {
	err := s.producer.Publish(s.topic, data)
	if err != nil {
		//记录错误
		logger.Sugar.Errorf("logic.publish err : %s\n", err)
	}
}
