package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"net"
)

var clients map[string]chan string		// 记录连接进来的客户端连接
var msg chan string				// 传输客户端消息的channel
var entering chan string		// 传输客户端连接进来时服务端产生的通知消息
var leaving chan string			// 传输客户端连接断开时服务端产生的通知消息（channel中放客户端的地址即可）

func main(){
	msg = make(chan string)
	entering = make(chan string)
	leaving = make(chan string)
	clients = map[string]chan string{}

	// 创建广播服务
	go broadcast()

	// 创建网络服务（阻塞）
	AcceptConn()
}

// 创建套接字监听端口，接收连接
func AcceptConn() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil{
		log.Fatal(err)		// 创建服务失败
	}

	defer func() {
		listener.Close()
		fmt.Println("服务端停止服务")
	}()

fmt.Println("创建服务成功！")

	for {
		// 接收连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("客户端连接服务器失败：%s\n", err)
			continue
		}

		addr := conn.RemoteAddr().String()
		clients[addr] = make(chan string)
fmt.Println("客户端连接成功：" + addr)
		entering <- addr

		// 并发的处理客户端连接之后发送过来的所有请求
		go handleConn(conn)
	}
}

// 处理单个客户端的所有请求
func handleConn(conn net.Conn){
	defer conn.Close()

	// 为每个连接开一个goroutine用于将这个消息从服务端发送给这个客户端
	go sendMsg(conn)

	// 不断接收客户端消息
	recvMsg(conn)
}

// 发送消息给客户端
func sendMsg(conn net.Conn){
	addr := conn.RemoteAddr().String()
	msgChan := clients[addr]	// 获取对应客户端的消息通道，该消息通道可以接收msg/entering/leaving这3种消息
	for {
		// 接收所有消息（msg/entering/leaving这3种消息）
		cont := <-msgChan
		_,err := fmt.Fprint(conn, "%s", cont)
		if err != nil {
			fmt.Printf("消息：%s 无法发送给客户端 %s ，原因：%s\n", cont, addr, err)
		}
	}
}

// 接收客户端消息并发送给其他客户端
func recvMsg(conn net.Conn){
	addr := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)

	// 服务端接收客户端发送过来的消息
	for input.Scan(){
		cont := input.Bytes()

		// 将消息广播给其他客户端
		message := fmt.Sprintf("%s : %s\n", addr, cont)
		msg <- message
	}

	fmt.Println(addr + "客户端断开连接")
	delete(clients, addr)		// 将客户端信息从map中剔除
	leaving <- addr
}

// 广播消息,负责接收msg/leaving/entering这3中消息并发送给所有客户端
func broadcast() {
	for{
		select{
		case c := <-msg:
			go _broadcast(c, "msg", "")
		case addr := <-entering:
			c := fmt.Sprintf("客户端 %s 加入群聊~\n", addr)
			go _broadcast(c, "entering", addr)
		case addr := <-leaving:
			c := fmt.Sprintf("客户端 %s 退出群聊\n", addr)
			go _broadcast(c, "leaving", addr)
		}
	}
}

// 这里的message参数可能是msg或leaving或entering中的任意一种消息
func _broadcast(message, typ, sender  string){
	for addr, client_chan := range clients{
		// 如果message是客户端退出的消息，则这个消息不发送给退出的那个客户端
		if typ == "leaving" && sender == addr{
			continue
		}
		client_chan <- message
	}
}