package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
/*
发送和接收两个操作都使用 <-   运算符。
在发送语句中， <-   运算符分割 channel 和要发送的值。
在接收语句中, <-   运算符写在 channel  对象之前。
*/
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
