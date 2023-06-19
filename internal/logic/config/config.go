package config

import (
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     *Config
)

type Config struct {
	Srv  *Srv
	Etcd *Etcd
	Nsq  *Nsq
}

type Srv struct {
	SrvName    string
	ListenAddr string
	TTL        int64
}
type Etcd struct {
	Addrs       []string
	DialTimeOut int
}

type Nsq struct {
	Addr  string
	Topic string
}

func init() {
	flag.StringVar(&confPath, "conf", "logic-example.toml", "default config path.")
}

func Init() error {
	Conf = Default()
	_, err := toml.DecodeFile(confPath, &Conf)
	return err
}

func Default() *Config {
	return &Config{
		Srv: &Srv{
			SrvName:    "room_logic",
			ListenAddr: "127.0.0.1:6001",
			TTL:        5,
		},
		Etcd: &Etcd{
			Addrs:       []string{"127.0.0.1:2379"},
			DialTimeOut: 5,
		},
		Nsq: &Nsq{
			Addr:  "127.0.0.1:4150",
			Topic: "chat_room",
		},
	}
}
