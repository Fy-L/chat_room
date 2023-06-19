package main

import (
	"chat_room/internal/job"
	"chat_room/internal/job/config"
	"chat_room/pkg/discovery"
	"chat_room/pkg/logger"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建并监听 gops agent，gops 命令会通过连接 agent 来读取进程信息
	// 若需要远程访问，可配置 agent.Options{Addr: "0.0.0.0:6060"}，否则默认仅允许本地访问
	// if err := agent.Listen(agent.Options{}); err != nil {
	// 	log.Fatalf("agent.Listen err: %v", err)
	// }
	flag.Parse()
	err := config.Init()
	if err != nil {
		panic(err)
	}
	//开启服务
	// j := job.NewJob(config.Conf)
	_, err = job.NewJob(config.Conf)
	if err != nil {
		log.Printf("create job err : %s\n", err)
		return
	}
	log.Printf("job-srv [%s] start!", config.Conf.Srv.ListenAddr)

	/*
		//注册到etcd中
		dis, err := register()
		if err != nil {
			fmt.Printf("服务注册失败：%s", err)
			return
		}

		//开启rpc服务
		addr := config.Conf.Srv.ListenAddr
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			log.Printf("监听网络失败, err : %s\n", err)
			return
		}
		defer listen.Close()
		log.Printf("Listenning :%s\n", addr)

		//实例化grpc
		s := grpc.NewServer()
		defer s.GracefulStop()
		//在grpc上注册
		api.RegisterJobServer(s, j)
		//启动服务
		logger.Logger.Info("rpc启动成功")
		log.Println("rpc启动成功")
		go func() {
			if err = s.Serve(listen); err != nil {
				fmt.Println(err)
				panic("rpc启动失败")
			}
		}()
	*/

	//关闭信号处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// dis.Stop()
			logger.Sugar.Infof("job-srv [%s] exit!", config.Conf.Srv.ListenAddr)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

//服务注册
func register() (*discovery.Register, error) {
	dis := discovery.NewRegister(config.Conf.Etcd.Addrs, config.Conf.Etcd.DialTimeOut, logger.Logger)
	_, err := dis.Register(discovery.Server{Name: config.Conf.Srv.SrvName, Addr: config.Conf.Srv.ListenAddr}, config.Conf.Srv.TTL)
	if err != nil {
		return nil, err
	}
	info, _ := dis.GetServerInfo()
	fmt.Printf("注册成功:[%s]\n", discovery.BuildRegPath(info))
	return dis, nil
}
