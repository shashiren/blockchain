package main

import (
	"fmt"
	"time"
)

func  worker(done chan bool)  {
	fmt.Print("worker...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main()  {
    //启动一个goroutine工作程序，给它一个同通知通道。
    // 如果从程序中删除<-done行，程序将在工作程序（worker）工作之前退出
	done := make(chan bool,1)
	go worker(done)
    //锁住，直到在通道上收到工作程序的通知。
	<- done

}