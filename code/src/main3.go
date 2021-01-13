package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 该方法是对 io.Copy 的简单封装，这样的话这个方法既能从服务端接收消息并输出到客户端，又能把客户端的消息发送给服务端
// r是消息发送端，w是消息接收端，mustCopy的意思是从一个可读对象r读取数据，再发送到可写对象w中
// 该方法 内部会阻塞，会无限循环（io.Copy内部实现了死循环）接收Reader的消息
func mustCopy(w io.Writer, r io.Reader){
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil{
		log.Fatal()
	}

	defer conn.Close()

	// 开一个子协程用于接收服务端的消息
	go mustCopy(os.Stdout, conn)	// 将服务端发过来的消息发送到客户端的标准输出（屏幕上）

	// 主协程则向服务端发送消息
	mustCopy(conn, os.Stdin)    // 将客户端标准输入拷贝（发送）到服务端
}

