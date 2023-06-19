package job

import (
	"chat_room/api/conn"
	"chat_room/internal/job/config"
	"chat_room/pkg/logger"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Job struct {
	c *config.Config
	// consumer //消费者，
	connServers map[string]*Server //记录ip 与server关系
	rooms       map[string]*Room
	cli         *clientv3.Client
	roomRWLock  sync.RWMutex
}

// 获取Job实例
func NewJob(c *config.Config) (*Job, error) {
	j := &Job{
		c:           c,
		connServers: make(map[string]*Server),
		rooms:       make(map[string]*Room),
	}
	go j.connsproc()
	err := j.consumer()
	if err != nil {
		return nil, err
	}
	return j, nil
}

// nsq消费者
func (j *Job) consumer() error {
	return initConsumer(j)

}

func (j *Job) push(req *conn.BroadcastRoomReq) error {
	err := j.getRoom(req.RoomID).pushmsg(req)
	return err
}

//发送msg到每个connServer
func (j *Job) broadcastRoom(roomID string, b []byte, msglv conn.MsgLevel) {
	arg := &conn.BroadcastRoomReq{
		RoomID: roomID,
		MsgLv:  msglv,
		Data:   b,
	}
	for _, s := range j.connServers {
		s.BroadcastRoom(arg)
	}
}

func (j *Job) connsproc() {
	//连接etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   j.c.Etcd.Addrs,
		DialTimeout: time.Duration(j.c.Etcd.DialTimeOut) * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	j.cli = cli

	if err := j.syncConns(); err != nil {
		logger.Sugar.Errorf("sync faild %s", err)
	}

	//监听 key = "/live_conn/" 的变化
	//获取key对应的值放到connServers中
	key := fmt.Sprintf("/%s/", j.c.Rpc.ConnRpcPrefix)
	ticker := time.NewTicker(time.Minute)
	watch := j.cli.Watch(context.Background(), key, clientv3.WithPrefix())
	for {
		select {
		case res, ok := <-watch:
			if ok {
				j.updateConns(res.Events)
			}
		case <-ticker.C:
			if err := j.syncConns(); err != nil {
				logger.Sugar.Errorf("sync faild %s", err)
			}
		}
	}

}

//更新
func (j *Job) updateConns(events []*clientv3.Event) {
	for _, ev := range events {

		addr := strings.Split(string(ev.Kv.Key), "/")[2]
		switch ev.Type {
		case mvccpb.PUT:
			if _, ok := j.connServers[addr]; ok {
				continue
			}
			j.connServers[addr] = NewServer(j.c, addr)
		case mvccpb.DELETE:
			if _, ok := j.connServers[addr]; !ok {
				continue
			}
			//删除
			delete(j.connServers, addr)
		}
	}
}

//获取所有connSrv
func (j *Job) syncConns() error {
	key := fmt.Sprintf("/%s/", j.c.Rpc.ConnRpcPrefix)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(j.c.Etcd.DialTimeOut)*time.Second)
	defer cancel()
	res, err := j.cli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, v := range res.Kvs {
		tmp := strings.Split(string(v.Key), "/")
		if _, ok := j.connServers[tmp[2]]; ok {
			continue
		}
		j.connServers[tmp[2]] = NewServer(j.c, tmp[2])
	}
	return nil
}
