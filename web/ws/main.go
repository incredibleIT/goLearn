package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

//
//import (
//	json "encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"github.com/gorilla/websocket"
//	"net/http"
//)
//
// 升级ws
// 注册客户端
// 作握手
// 添加用户
// 为每个客户端开启一个协程监听消息
// 一旦有消息就广播

// 抽取每一个会话的结构体
type Client struct {
	// 包含ws连接器
	ws *websocket.Conn
	// 客户端IP地址
	Ip string `json:"ip"`
}

type Msg struct {
	// 标识信息类型的字段
	// HandShake：握手信息，弹出登录框
	// Login：正在登录
	// Chat：正常普通用户会话, 由前端传
	// Logout：退出的信息
	Type string `json:"type"`
	// 消息内容
	Content string `json:"content"`
	// 发送消息的用户名
	User string `json:"user"`
	// 聊天室用户列表
	UserList []string `json:"user_list"`
}

var onlineClients = make(map[string]Client) //保存所有用户  {"Addr":{"Addr":Addr,"Chan":Chan}}
var broadcast = make(chan *Msg)             // 广播管道
// 定义全局用户列表
var userList []string

// 从用户列表切片删除一个用户
func delUser(slice []string, user string) []string {
	for i := 0; i < len(slice); i++ {
		if user == slice[i] {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

//
////监听broadcast通道中的数据,一旦有数据，循环写入每一个Client的Chan中,进而是每一个客户端收到该广播数据
//func MessageManager() {
//	for {
//		msg := <-broadcast
//		switch msg.Type {
//		case "Login":
//
//			// 登录记录
//			userList = append(userList, msg.User)
//			msg.UserList = userList
//			msg.User = msg.Content
//		case "Logout":
//			userList = delUser(userList, msg.User)
//			msg.UserList = userList
//			// 发送Logout请求时，msg的Content需要的IP
//			delete(onlineClients, msg.Content)
//		}
//
//		for _, client := range onlineClients {
//			data, _ := json.Marshal(*msg)
//			client.ws.WriteMessage(websocket.TextMessage, data)
//		}
//	}
//
//}
//
////监听该客户端的socket管道，一旦有数据,写入broadcast管道，进而发给每一个客户端
//func readMsgFromClient(client Client) {
//	// 如果写外面丢失用户数据  User(string)
//	var msg Msg
//	for {
//		_, message, _ := client.ws.ReadMessage()
//		json.Unmarshal(message, &msg)
//		fmt.Println(msg)
//		broadcast <- &msg
//	}
//
//	func() {
//		msg.Type = "Logout"
//		msg.Content = client.Ip
//		broadcast <- &msg
//	}()
//}
//
////为每一个客户端开启协程处理函数
//func HandleConnect(w http.ResponseWriter, r *http.Request) {
//	var upgrader = &websocket.Upgrader{ReadBufferSize: 1024,
//		WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
//			return true
//		}}
//	// 获取升级信息，拿到ws连接
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		return
//	}
//
//	client := Client{Ip: r.RemoteAddr, ws: ws}
//	onlineClients[client.Ip] = client
//	go readMsgFromClient(client)
//	// 响应这次握手请求
//	data := Msg{Type: "HandShake"}
//	jsonData, _ := json.Marshal(data)
//	client.ws.WriteMessage(websocket.TextMessage, jsonData)
//
//}
//
////主协程
//func main() {
//	// 创建路由
//	router := mux.NewRouter()
//	fmt.Println("服务端已开启,等待客户端连接中...")
//	router.HandleFunc("/ws", HandleConnect)
//	//负责监听广播通道中的数据
//	go MessageManager()
//	// 开启服务
//	if err := http.ListenAndServe("127.0.0.1:8060", router); err != nil {
//		fmt.Println("err :", err)
//	}
//}

/*
关键：每一个客户端维持一个msg
*/

func messageManager() {
	for {
		msg := <-broadcast
		switch msg.Type {
		case "Login": // 做登录
			fmt.Println("user", msg.User, "content", msg.Content)
			userList = append(userList, msg.Content)
			msg.UserList = userList
			msg.User = msg.Content
		case "Logout":
			delete(onlineClients, msg.Content)
			userList = delUser(userList, msg.User)
			msg.UserList = userList
		}

		for _, client := range onlineClients {
			msgJson, _ := json.Marshal(*msg)
			client.ws.WriteMessage(websocket.TextMessage, msgJson)
		}
	}
}

func readMsgFromClient(client Client) {
	var msg Msg
	for {
		_, message, _ := client.ws.ReadMessage()
		json.Unmarshal(message, &msg)
		broadcast <- &msg
	}

	func() {
		msg.Type = "Logout"
		msg.Content = client.Ip
		broadcast <- &msg
	}()
}

func HandleConnect(w http.ResponseWriter, r *http.Request) {

	// 升级
	var upgrader = &websocket.Upgrader{ReadBufferSize: 1024,
		WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
			return true
		}}
	// 获取升级信息，拿到ws连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := Client{ws: ws, Ip: r.RemoteAddr}
	onlineClients[client.Ip] = client
	go readMsgFromClient(client)

	// 握手
	msg := Msg{Type: "HandShake"}
	msgData, _ := json.Marshal(msg)
	ws.WriteMessage(websocket.TextMessage, msgData)

}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/ws", HandleConnect)

	go messageManager()

	http.ListenAndServe(":8060", mux)
}
