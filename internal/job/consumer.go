package job

import (
	"chat_room/api/conn"
	"chat_room/pkg/logger"
	"fmt"
	"strings"

	"github.com/nsqio/go-nsq"
	"google.golang.org/protobuf/proto"
)

// 消费者结构体，用于实现nsq.Handler 接口
type consumerHandle struct {
	j *Job
}

var _ (nsq.Handler) = (*consumerHandle)(nil)

// 处理消息
func (cus *consumerHandle) HandleMessage(message *nsq.Message) error {
	body := message.Body
	var req = new(conn.BroadcastRoomReq)
	// var push = new(conn.BroadcastRoomReq)

	err := proto.Unmarshal(body, req)
	if err != nil {
		logger.Sugar.Errorf("proto 解析失败: %s ", err)
	}
	// push.RoomID = req.RoomID
	// push.Data = req.Data
	err = cus.j.getRoom(req.RoomID).pushmsg(req)
	if err != nil {
		logger.Sugar.Errorf("room push err: %s", err)
	}
	return nil
}

// 初始化
func initConsumer(j *Job) error {
	config := nsq.NewConfig()
	ch := fmt.Sprintf("%s_%s", j.c.Srv.SrvName, strings.Split(j.c.Srv.ListenAddr, ":")[1])
	c, err := nsq.NewConsumer(j.c.Nsq.Topic, ch, config)
	if err != nil {
		logger.Sugar.Errorf("启动job consumer 失败：%s ", err)
		return err
	}
	// defer c.Stop()
	cus := &consumerHandle{j}
	c.AddHandler(cus)

	//连接nsq
	if err := c.ConnectToNSQD(j.c.Nsq.Addr); err != nil { // 直接连NSQD
		logger.Sugar.Error("连接nsq失败: %s ", err)
		return err
	}
	return nil
}
