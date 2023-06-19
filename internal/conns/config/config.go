package config

import (
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// wsBind   string
	// srvAddr  string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "conn-example.toml", "default config path.")
	// flag.StringVar(&wsBind, "ws", ":8081", "ws listen addr. default :8081")
	// flag.StringVar(&srvAddr, "srvaddr", "127.0.0.1:6000", "srv listen addr. default 127.0.0.1:6000")
}

func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(confPath, &Conf)
	return err
}

func Default() *Config {
	return &Config{
		Bucket: &Bucket{
			Size:          32,
			Conns:         1024,
			Room:          1024,
			RoutineAmount: 32,
			RoutineSize:   1024,
		},
		Websocket: &Websocket{
			Bind: ":8081",
		},
		Rpc: &Rpc{
			Prefix:         "etcd://chat_room/",
			LogicRpcPrefix: "room_logic",
			JobRpcPrefix:   "room_job",
		},
		Srv: &Srv{
			SrvName:    "room_conn",
			ListenAddr: "127.0.0.1:6000",
			TTL:        5,
		},
		Etcd: &Etcd{
			Addrs:       []string{"127.0.0.1:2379"},
			DialTimeOut: 5,
		},
	}
}

type Config struct {
	Bucket    *Bucket
	Websocket *Websocket

	Rpc  *Rpc
	Srv  *Srv
	Etcd *Etcd
}

type Websocket struct {
	Bind string
}

type Bucket struct {
	Size          int    //分多少个bucket
	Conns         int    //每个bucket连接的数量
	Room          int    //每个bucket的room的数量
	RoutineAmount uint64 //每个bucket协程数量
	RoutineSize   int    //每个协程的size
}

type Rpc struct {
	Prefix         string
	LogicRpcPrefix string
	JobRpcPrefix   string
}

type Etcd struct {
	Addrs       []string
	DialTimeOut int
}

type Srv struct {
	SrvName    string
	ListenAddr string
	TTL        int64
}
