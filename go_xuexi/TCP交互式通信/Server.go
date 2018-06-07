package main

import (
	"fmt"
	"net"
)

func CheckError1(err error)  {
	if err !=nil {
		fmt.Println("网络错误",err.Error())
	}
}

func Processinfo(conn net.Conn)  {
	buf :=make([]byte,1024)//开创缓冲区
	defer conn.Close()//关闭链接
	for  {
		numOfBytes,err:=conn.Read(buf)//读取数据
		if err !=nil {
			break
		}
		if numOfBytes!=0 {

			remoteAddr:=conn.RemoteAddr()
			fmt.Println("收到消息",string(buf),"来自",remoteAddr)
		}

	}

}

func main()  {
	//建立TCP服务器
	listen_socket,err:=net.Listen("tcp","127.0.0.1:8898")
	defer listen_socket.Close()//关闭网络
	CheckError1(err)
	fmt.Println("服务器正在等待")
	for {
		conn,err:=listen_socket.Accept()//新的客户端连接
		CheckError1(err)

		//处理每一个客户端

		go Processinfo(conn)

	}



}

