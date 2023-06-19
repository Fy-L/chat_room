package conns

import (
	"errors"
	"sync"
)

type Room struct {
	id     string
	next   *Conn //用户连接链表
	rLock  sync.RWMutex
	online uint32
	drop   bool //是否丢弃
}

func NewRoom(id string) *Room {
	r := new(Room)
	r.id = id
	r.next = nil
	r.online = 0
	r.drop = false
	return r
}

// put conn into room
//存放conn
func (r *Room) Put(conn *Conn) (err error) {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	if r.drop {
		err = errors.New("房间已丢弃")
	} else {
		//新的连接从头部插入
		if r.next != nil {
			//如果room已经存在连接，则改变当前room第一个连接的prev为新的conn
			r.next.Prev = conn
		}
		//新的连接的next 改为room的next
		conn.Next = r.next
		//新的连接prev为nil，方便下个连接的加入
		conn.Prev = nil
		//插入当前room的头部
		r.next = conn
		r.online++
	}

	return
}

//移除conn
func (r *Room) Del(conn *Conn) bool {
	r.rLock.Lock()
	if conn.Next != nil {
		//如果不是尾部
		conn.Next.Prev = conn.Prev
	}
	if conn.Prev != nil {
		// 如果不是头部
		conn.Prev.Next = conn.Next
	} else {
		r.next = conn.Next
	}
	conn.Next = nil
	conn.Prev = nil
	r.online--
	r.drop = r.online == 0
	r.rLock.Unlock()
	return r.drop
}

//推送消息
func (r *Room) Push(b []byte) {
	//1.获取当前直播间所有连接
	//2.循环推送消息
	r.rLock.RLock()

	for user := r.next; user != nil; user = user.Next {
		user.WriteMsg(b)
	}
	r.rLock.RUnlock()
}

//获取当前在线人数
func (r *Room) OnlineMembers() uint32 {
	r.rLock.RLock()
	defer r.rLock.RUnlock()
	return r.online
}
func (r *Room) Close() {
	r.rLock.RLock()
	for conn := r.next; conn != nil; conn = conn.Next {
		conn.Close()
	}
	r.rLock.RUnlock()
}
