package job

import (
	"chat_room/api/conn"
	"chat_room/internal/job/config"
	"chat_room/pkg/logger"
	"context"
	"fmt"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	client      conn.ConnClient
	addr        string
	roomChan    []chan *conn.BroadcastRoomReq
	roomChanNum uint64
	routineSize uint64

	ctx    context.Context
	cancel context.CancelFunc
	c      *config.Config
}

//初始化服务
func NewServer(c *config.Config, addr string) *Server {
	s := &Server{
		client:      newRpcClient(addr),
		roomChan:    make([]chan *conn.BroadcastRoomReq, c.Conn.RoutineSize),
		addr:        addr,
		routineSize: uint64(c.Conn.RoutineSize),
		c:           c,
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	//开启对应数量的协程
	for i := 0; i < c.Conn.RoutineSize; i++ {
		s.roomChan[i] = make(chan *conn.BroadcastRoomReq, c.Conn.RoutineChan)
		go s.process(s.roomChan[i])
	}
	return s
}

// 广播
func (s *Server) BroadcastRoom(arg *conn.BroadcastRoomReq) {
	idx := atomic.AddUint64(&s.roomChanNum, 1) % s.routineSize
	s.roomChan[idx] <- arg
}

func (s *Server) process(roomChan chan *conn.BroadcastRoomReq) {
	for {
		select {
		case roomArg := <-roomChan:
			_, err := s.client.BroadcastRoom(context.Background(), roomArg)
			if err != nil {
				logger.Sugar.Errorf("s.client.BroadcastRoom(%s, reply) error(%v)", roomArg, err)
			}
		case <-s.ctx.Done():
			return
		}
	}
}

// conn的rpc.client
func newRpcClient(addr string) conn.ConnClient {
	connect, err := grpc.DialContext(context.TODO(), addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		panic(err)
	}
	return conn.NewConnClient(connect)
}
