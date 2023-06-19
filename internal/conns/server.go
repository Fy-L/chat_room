package conns

import (
	"chat_room/api/logic"
	"chat_room/internal/conns/config"
	"chat_room/pkg/discovery"
	"chat_room/pkg/logger"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type Server struct {
	buckets        []*Bucket
	bucketIdx      uint32
	conf           *config.Config
	logicRpcClient logic.LogicClient
	// jobRpcClient   job.JobClient
}

var SrvMrg *Server

//启动服务
func NewServer(c *config.Config) *Server {
	s := &Server{
		buckets:        make([]*Bucket, c.Bucket.Size),
		bucketIdx:      uint32(c.Bucket.Size),
		conf:           c,
		logicRpcClient: newLogicRpcClient(c),
		// jobRpcClient:   newJobRpcClient(c),
	}
	for i := 0; i < c.Bucket.Size; i++ {
		s.buckets[i] = NewBucket(c.Bucket)
	}
	SrvMrg = s
	go s.onlineproc()
	return s
}

//logicRpcClient
func newLogicRpcClient(c *config.Config) logic.LogicClient {
	//etcd解析器
	r := discovery.NewResolver(c.Etcd.Addrs, c.Etcd.DialTimeOut, logger.Logger)
	resolver.Register(r)
	connect, err := grpc.DialContext(context.TODO(),
		fmt.Sprintf("%s%s", c.Rpc.Prefix, c.Rpc.LogicRpcPrefix),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		log.Println("logic rpc client 链接失败")
		panic(err)
	}
	return logic.NewLogicClient(connect)
}

//jobRpcClient
// func newJobRpcClient(c *config.Config) job.JobClient {
// 	//etcd解析器
// 	r := discovery.NewResolver(c.Etcd.Addrs, c.Etcd.DialTimeOut, logger.Logger)
// 	resolver.Register(r)
// 	connect, err := grpc.DialContext(context.Background(),
// 		fmt.Sprintf("%s%s", c.Rpc.Prefix, c.Rpc.JobRpcPrefix),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
// 	if err != nil {
// 		log.Println("job rpc client 链接失败")
// 		panic(err)
// 	}
// 	return job.NewJobClient(connect)
// }

//获取bucket
func (s *Server) Bucket(key uint32) *Bucket {
	hash := key % s.bucketIdx
	return s.buckets[hash]
}

func (s *Server) Buckets() []*Bucket {
	return s.buckets
}

func (s *Server) onlineproc() {
	for {
		time.Sleep(time.Second * 10)
		cts := 0
		for _, b := range s.buckets {
			if b.cnts > 0 {
				// fmt.Printf("Buckets %d cnts %d \n", i, b.cnts)
				cts += int(b.cnts)
				// fmt.Printf("room %+v\n", b.rooms)
			}

		}
		fmt.Printf("当前连接数量 %d\n", cts)
	}
}
