package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		msg <- "ping" // 使用 channel <- 语法 发送 一个新的值到通道中
	}()

	out := <-msg // 使用 <-channel 语法从通道中 接收 一个值。
	fmt.Println("out is ", out)
}
