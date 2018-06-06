package main

import (
	"os"
	"fmt"
	"net"

)

func checkError(err error)  {
	if err!=nil{//指针不为空
		fmt.Println("Error",err.Error())
		os.Exit(1)
	}
}

func recUDPMsg(conn*net.UDPConn)  {//接收消息
	var buf[30]  byte //缓冲数组
	n,raddr,err := conn.ReadFromUDP(buf[0:])//从udp接收数据
	if err!=nil{
		fmt.Println("Error",err.Error())//打印错误信息
		return
	}
	fmt.Println("消息是",string(buf[0:n]))//n代表读取了多少字节
	_,err=conn.WriteToUDP([]byte("hao client"),raddr)//写入UDP，根据地址发送消息


}

func main()  {
	//创造服务器
    udp_addr,err:=net.ResolveUDPAddr("udp",":8848")
    checkError(err)

    conn,err:=net.ListenUDP("udp",udp_addr)//开始监听
    defer conn.Close()//关闭链接
    checkError(err)
    recUDPMsg(conn)//收消息
}

