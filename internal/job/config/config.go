package config

import (
	"flag"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "job-example.toml", "default config path.")
}
func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(confPath, &Conf)
	return err

}

type Config struct {
	Rpc  *Rpc
	Conn *Conn
	Room *Room
	Etcd *Etcd
	Srv  *Srv
	Nsq  *Nsq
}

type Rpc struct {
	Prefix        string
	ConnRpcPrefix string
}
type Conn struct {
	RoutineChan int //每个协程的大小
	RoutineSize int //每个connServer的协程数量
}
type Room struct {
	Batch         int           //批量发送数量
	Signal        time.Duration //触发信号时间
	Idle          time.Duration //空闲时间
	MaxBufferSize int           //包大小
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
type Nsq struct {
	Addr  string
	Topic string
}

func Default() *Config {
	return &Config{
		Rpc: &Rpc{
			Prefix:        "etcd://chat_room/",
			ConnRpcPrefix: "live_conn",
		},
		Conn: &Conn{
			RoutineChan: 1024,
			RoutineSize: 32,
		},
		Room: &Room{
			Batch:         20,
			Signal:        time.Second,
			Idle:          time.Minute * 10,
			MaxBufferSize: 4096,
		},
		Etcd: &Etcd{
			Addrs:       []string{"127.0.0.1:2379"},
			DialTimeOut: 5,
		},
		Srv: &Srv{
			SrvName:    "room_job",
			ListenAddr: "127.0.0.1:6002",
			TTL:        5,
		},
		Nsq: &Nsq{
			Addr:  "127.0.0.1:4150",
			Topic: "chat_room",
		},
	}
}
