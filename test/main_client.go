package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	//在线人数
	on = flag.Int64("on", 10000, "在线人数")
	//ws
	addr  = flag.String("addr", "127.0.0.1:8081", "ws地址")
	addr2 = flag.String("addr2", "127.0.0.1:8081", "ws2地址")
	//运行时间
	ti = flag.Int64("ti", 5, "运行时间")
	//没每秒发送消息
	per = flag.Int("per", 1001, "在线聊天人数")
	mrg *Mrg
)

func main() {
	flag.Parse()
	if int((*on)) < (*per) {
		fmt.Println("在线聊天人数，不能大于在线人数")
		return
	}
	mrg = newMrg(*on, *addr, *addr2)
	mrg.Run()
	// fmt.Println("真实在线人数", mrg.online_user)
	t := time.NewTicker(time.Duration(*ti) * time.Minute)
	for {
		select {
		case <-t.C:
			return
		}

	}
}

type Mrg struct {
	online_user int64
	addr        []string
}

func newMrg(on int64, addr string, addr2 string) *Mrg {
	return &Mrg{
		online_user: on,
		addr:        []string{addr, addr2},
	}
}

func (m *Mrg) Run() {

	room := []string{"room1", "room2", "room3"}
	//创建对应数量的client
	for i := int64(1); i <= m.online_user; i++ {
		url := fmt.Sprintf("ws://%s/room", m.addr[rand.Intn(2)])
		go Client(url, int(i), *per, room[rand.Intn(3)])
		// go Client(url, int(i), "room1")
	}

}
