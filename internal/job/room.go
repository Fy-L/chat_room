package job

import (
	"chat_room/api/conn"
	"chat_room/internal/job/config"
	"chat_room/pkg/bytes"
	"chat_room/pkg/datapack"
	"chat_room/pkg/logger"
	"errors"
	"time"
)

var roomReady = new(conn.BroadcastRoomReq)

type Room struct {
	id    string
	job   *Job
	c     *config.Room
	proto chan *conn.BroadcastRoomReq
}

func NewRoom(j *Job, roomID string, c *config.Room) *Room {
	r := &Room{
		id:    roomID,
		job:   j,
		c:     c,
		proto: make(chan *conn.BroadcastRoomReq, c.Batch*2),
	}
	go r.pushproccess(c.Batch, c.Signal)
	return r
}

// push msg to the room, if chan full discard it.
// 消息推送，如果chan满了，则舍弃部分消息
func (r *Room) pushmsg(req *conn.BroadcastRoomReq) (err error) {
	select {
	case r.proto <- req:
	default:
		err = errors.New("room proto chan full")
	}
	return
}

// 监听chan，推送消息
// btach 合并最大值，超过则会发送消息
// signTime 触发发送消息时间
func (r *Room) pushproccess(batch int, signTime time.Duration) {
	var (
		n    int
		last time.Time
		p    *conn.BroadcastRoomReq
		buf  = bytes.NewWriterSize(r.c.MaxBufferSize)
	)
	logger.Sugar.Infof("start room:%s goroutine\n", r.id)
	//signTime时间后，执行func
	td := time.AfterFunc(signTime, func() {
		select {
		case r.proto <- roomReady:
		default:
		}
	})
	defer td.Stop()
	for {
		p = <-r.proto
		if p == nil {
			break //退出
		} else if p != roomReady {
			//如果不是ready，则代表有消息进入

			//打包
			packData, err := datapack.NewDataPack().Pack(p.Data)
			if err != nil {
				continue
			}
			//判断是否重要信息
			if p.MsgLv == conn.MsgLevel_IMPORTANT {
				//则优先发送
				r.job.broadcastRoom(r.id, packData, conn.MsgLevel_IMPORTANT)
				continue
			}

			//合并buffer
			buf.Write(packData)
			n++
			if n == 1 {
				last = time.Now()
				td.Reset(signTime)
				continue
			} else if n < batch {
				//如果n 小于 给定batch合并发送量，而且运行时间还未超过signTime则继续，无需发送消息
				if signTime > time.Since(last) {
					continue
				}
			}
		} else {
			if n == 0 {
				break
			}
		}

		//发送消息
		r.job.broadcastRoom(r.id, buf.Buffer(), conn.MsgLevel_NORMAL)
		//重置buffer
		buf = bytes.NewWriterSize(buf.Size())
		n = 0
		if r.c.Idle != 0 {
			td.Reset(time.Duration(r.c.Idle))
		} else {
			td.Reset(time.Minute)
		}
	}
	//删除room
	r.job.delRoom(r.id)
	logger.Sugar.Infof("room:%s goroutine exit", r.id)
}

//删除room
func (j *Job) delRoom(roomID string) {
	j.roomRWLock.Lock()
	delete(j.rooms, roomID)
	j.roomRWLock.Unlock()
}

// 获取room
func (j *Job) getRoom(roomID string) *Room {
	j.roomRWLock.RLock()
	room, ok := j.rooms[roomID]
	j.roomRWLock.RUnlock()
	if !ok {
		j.roomRWLock.Lock()
		if room, ok = j.rooms[roomID]; !ok {
			room = NewRoom(j, roomID, j.c.Room)
			j.rooms[roomID] = room
		}
		j.roomRWLock.Unlock()
	}
	return room
}
