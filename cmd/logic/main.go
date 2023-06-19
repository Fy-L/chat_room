package main

import (
	api "chat_room/api/logic"
	"chat_room/internal/logic"
	"chat_room/internal/logic/config"
	"chat_room/pkg/discovery"
	"chat_room/pkg/logger"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	err := config.Init()
	if err != nil {
		panic(err)
	}

	//注册到etcd
	dis, err := register()
	if err != nil {
		fmt.Printf("服务注册失败：%s", err)
		return
	}
	//开启rpc
	addr := config.Conf.Srv.ListenAddr
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	l, err := logic.NewLogicSrv(config.Conf.Nsq.Addr, config.Conf.Nsq.Topic)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("logic-srv [%s] start!", addr)

	//实例化grpc
	s := grpc.NewServer()
	defer s.GracefulStop()
	//在grpc上注册微服务
	api.RegisterLogicServer(s, l)
	logger.Logger.Info("rpc启动成功")
	log.Println("rpc启动成功")
	//启动服务
	go func() {
		if err = s.Serve(listen); err != nil {
			panic("rpc启动失败")
		}
	}()

	//关闭信号处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			dis.Stop()
			logger.Sugar.Infof("logic-srv [%s] exit!", config.Conf.Srv.ListenAddr)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

//服务注册
func register() (*discovery.Register, error) {
	etcdAddrs := config.Conf.Etcd.Addrs
	dis := discovery.NewRegister(etcdAddrs, config.Conf.Etcd.DialTimeOut, logger.Logger)
	_, err := dis.Register(discovery.Server{Name: config.Conf.Srv.SrvName, Addr: config.Conf.Srv.ListenAddr}, config.Conf.Srv.TTL)
	if err != nil {
		return nil, err
	}
	info, _ := dis.GetServerInfo()
	fmt.Printf("注册成功:[%s]\n", discovery.BuildRegPath(info))
	return dis, nil
}
