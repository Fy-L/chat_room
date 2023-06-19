package conns

import (
	pb_conn "chat_room/api/conn"
	"chat_room/internal/conns/config"
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/proto"
)

//存放ws连接
type Bucket struct {
	cLock   sync.RWMutex     //读写锁
	conns   map[string]*Conn //连接
	rooms   map[string]*Room //房间
	cnts    int              //连接数量
	maxCnts int              //允许最大连接数量

	routines       []chan *pb_conn.BroadcastRoomReq
	routinesNum    uint64 //当前协程的编号
	routinesAmount uint64 //协程数量
}

// 获取实例
func NewBucket(c *config.Bucket) *Bucket {
	b := new(Bucket)
	b.conns = make(map[string]*Conn)
	b.rooms = make(map[string]*Room)
	b.routines = make([]chan *pb_conn.BroadcastRoomReq, c.RoutineAmount)
	b.routinesAmount = c.RoutineAmount
	b.maxCnts = c.Conns
	for i := uint64(0); i < b.routinesAmount; i++ {
		c := make(chan *pb_conn.BroadcastRoomReq, c.Size)
		b.routines[i] = c
		go b.roomproc(c)
	}
	return b
}

// 存放conn到bucket中
func (b *Bucket) Put(roomId string, conn *Conn) {
	var (
		room *Room
		ok   bool
	)
	b.cLock.RLock()
	//超过最大连接
	if b.cnts >= b.maxCnts {
		// fmt.Println("close--", conn.UserId)
		b.cLock.RUnlock()
		b.Del(conn)
		return
	}
	b.cLock.RUnlock()

	b.cLock.Lock()
	//添加用户连接conn到bucket
	if dconn := b.conns[conn.Key]; dconn != nil {
		//删除旧连接
		dconn.Close()
	}

	b.conns[conn.Key] = conn
	b.cnts++

	//如果室新的聊天室，则添加到bucket中
	if room, ok = b.rooms[roomId]; !ok {
		room = NewRoom(roomId)
		b.rooms[roomId] = room
	}
	b.cLock.Unlock()

	//更新conn的room
	conn.Room = room
	//添加用户到room
	err := room.Put(conn)
	if err != nil {
		//加入房间错误
		// log.Fatalf("connect err %s \n", err)
		b.Del(conn)
	}
}

//删除连接
func (b *Bucket) Del(dconn *Conn) {
	// fmt.Printf("%+v\n", dconn)
	room := dconn.Room
	b.cLock.Lock()
	if conn, ok := b.conns[dconn.Key]; ok {
		if conn == dconn {
			delete(b.conns, dconn.Key)
		}
		b.cnts--
	}
	b.cLock.Unlock()
	if room != nil && room.Del(dconn) {
		//删除room
		b.cLock.Lock()
		delete(b.rooms, room.id)
		b.cLock.Unlock()
		//close room
		room.Close()
	}
}

// 获取room
func (b *Bucket) Room(rid string) (room *Room) {
	b.cLock.RLock()
	room = b.rooms[rid]
	b.cLock.RUnlock()
	return
}

//用于传输聊天室信息通道
func (b *Bucket) roomproc(c chan *pb_conn.BroadcastRoomReq) {
	for {
		arg := <-c
		// var msg Req
		// json.Unmarshal(arg, &msg)
		//封装pb_conn.Req
		var reply = new(pb_conn.Reply)
		var msg_b, _ = proto.Marshal(arg)
		reply.Type = pb_conn.PackageType_MESSAGE
		reply.Data = msg_b
		reply_b, _ := proto.Marshal(reply)
		if room := b.Room(arg.RoomID); room != nil {
			room.Push(reply_b)
		}
	}
}

//广播
func (b *Bucket) BroadcastRoom(req *pb_conn.BroadcastRoomReq) {
	num := atomic.AddUint64(&b.routinesNum, 1) % b.routinesAmount
	b.routines[num] <- req
}

//获取room在线人数
func (b *Bucket) OnlineMembers(rid string) uint32 {
	room := b.Room(rid)
	if room != nil {
		return room.OnlineMembers()
	}
	return 0
}
