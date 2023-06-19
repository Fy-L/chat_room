package conns

import (
	"chat_room/api/conn"
	pb_logic "chat_room/api/logic"
	"chat_room/pkg/token"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	_ "net/http/pprof"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//启动ws服务器
func StartWSServer(addr string) {
	http.HandleFunc("/room", wsHandler)
	http.Handle("/msg", crosMiddleware(sendMsgHandler))
	http.Handle("/tk", crosMiddleware(tokenHandler))
	log.Println("websocket server start.listening ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("connect err %s \n", err)
		return
	}

	var conn = NewConnect()
	conn.ws = wsConn

	//5秒后判断如果没有登录，则关闭连接
	time.AfterFunc(5*time.Second, func() {
		if conn.UserId <= 0 {
			conn.Close()
		}
	})
	for {
		//设置10分钟i/o timeout时间
		// err := conn.ws.SetReadDeadline(time.Now().Add(10 * time.Minute))
		// if err != nil {
		// 	fmt.Println(conn.UserId)
		// 	errHandle(conn, err)
		// 	return
		// }
		_, data, err := conn.ws.ReadMessage()
		if err != nil {
			errHandle(conn, err)
			return
		}
		//处理data信息
		conn.HandleMessage(data)
	}
}

// @Summary 发送重要消息
// @Description 发送重要消息
// @Tags 发送重要消息
// @Param room_id formData string  true "群id"
// @Param msg formData string true "消息（字符串，前端自定义json字符串格式）"
// @Security Authorization
// @param Authorization header string true "token"
// @Success 200 {string} string "{"code":1, "msg": "suc"}"
// @Failure 200 {string} string "{"code":0, "msg": "errMsg"}"
// @Router /msg [post]
func sendMsgHandler(w http.ResponseWriter, r *http.Request) {

	var (
		h = map[string]interface{}{
			"code": 1,
			"msg":  "suc",
		}
		push = new(conn.BroadcastRoomReq)
	)
	//先判断是否post
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//先判断token
	tokenString := r.Header.Get("Authorization")
	//验证前端传过来的token格式，不为空，开头为Bearer
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		h["code"] = 0
		h["msg"] = "无效token"
		sendJson(w, h)
		return
	}
	token := strings.TrimPrefix(tokenString, "Bearer ")
	//调用gprc验证token
	_, err := SrvMrg.logicRpcClient.Auth(context.Background(), &pb_logic.AuthReq{Token: token})
	if err != nil {
		h["code"] = 0
		h["msg"] = err.Error()
		sendJson(w, h)
		return
	}

	//获取post内容
	roomId := r.PostFormValue("room_id")
	content := r.PostFormValue("msg")
	if roomId == "" || content == "" {
		h["code"] = 0
		h["msg"] = "参数错误"
		sendJson(w, h)
		return
	}
	//grpc调用
	push.RoomID = roomId
	push.MsgLv = conn.MsgLevel_IMPORTANT
	push.Data, _ = json.Marshal(content)
	_, err = SrvMrg.logicRpcClient.PushMsg(context.Background(), push)

	if err != nil {
		h["code"] = 0
		h["msg"] = err.Error()
	}
	sendJson(w, h)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	var (
		h = map[string]interface{}{
			"code": 1,
			"msg":  "suc",
		}
	)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	uid := r.PostFormValue("uid")
	//生成token
	tk, err := token.GenToken(uid)
	if err != nil {
		h["code"] = 0
		h["msg"] = err.Error()
		sendJson(w, h)
		return
	}
	h["data"] = tk
	sendJson(w, h)
}

func sendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(data)
	w.Write(b)
}

func errHandle(conn *Conn, err error) {
	//处理错误
	// log.Printf("read msg err %s \n", err)
	if conn.UserId > 0 {
		bucket := SrvMrg.Bucket(conn.UserId)
		bucket.Del(conn)
	}
	conn.Close()
}

func crosMiddleware(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                         // 指明哪些请求源被允许访问资源，值可以为 "*"，"null"，或者单个源地址。
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")                              //对于预请求来说，指明了哪些头信息可以用于实际的请求中。
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                                                       //对于预请求来说，哪些请求方式可以用于实际的请求。
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") //对于预请求来说，指明哪些头信息可以安全的暴露给 CORS API 规范的 API
		w.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                 //指明当请求中省略 creadentials 标识时响应是否暴露。对于预请求来说，它表明实际的请求中可以包含用户凭证。

		//放行所有OPTIONS方法
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
