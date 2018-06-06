package main

import (
	"fmt"
	"net"
)

func CheckError(err error)  {
	if err !=nil {
		fmt.Println("网络错误",err.Error())
	}
}
func main()  {
	conn,err :=net.Dial("tcp","127.0.0.1:8898")//建立网络连接
	defer conn.Close()//延迟关闭网络连接
	CheckError(err) //检查错误

	conn.Write([]byte("hello nimei"))
	fmt.Println("发送消息","hello nimei")

}
