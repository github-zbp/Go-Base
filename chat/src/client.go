package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	//"os"
)

var msg string

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal()
	}

	defer func(){
		conn.Close()
		fmt.Println("服务端停止服务或客户端关闭连接或发送消息发生异常")
	}()

	// 接收消息
	go func(){
		for {
			resp, err := ioutil.ReadAll(conn)
			if err != nil{
				fmt.Println(err)
				break
			}
			fmt.Println(resp)
		}
	}()

	// 发送消息
	for {
		fmt.Scanln(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		conn.Write([]byte(msg))
fmt.Printf("send '%s' successfully ", msg)
	}
}