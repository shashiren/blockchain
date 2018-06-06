package main

import (
	"os"
	"fmt"
	"net"

)

func main()  {
	//链接网络服务器，IP，端口
	conn,err:=net.Dial("udp","127.0.0.1:8848")
	//关闭网络
	defer conn.Close()
	if err!=nil{
		fmt.Println("网络连接出错")
		os.Exit(1)
	}
	conn.Write([]byte("hello server"))//写入代表发送消息
	fmt.Println("发送消息","hello server")

	var  msg[30] byte //读取代表收取消息
	conn.Read(msg[0:])

	fmt.Println("收到消息",string(msg[:]))




}