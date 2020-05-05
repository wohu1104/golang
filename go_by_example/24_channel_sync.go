package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done..")
	done <- true
}

func main() {
	done := make(chan bool, 1) //创建一个 bool 值的通道
	go worker(done)            // 启动一个 go 协程
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
	// 如果把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
}
