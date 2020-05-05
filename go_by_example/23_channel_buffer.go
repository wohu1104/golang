package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan int, 5)
	go func() {
		for i := 0; i <= 3; i++ {
			msg <- i
		}
	}()

	time.Sleep(1 * time.Second)
	/*
		for j := range msg {
			fmt.Println("out is ", j)
			if j == 3 {
				break
			}
		}
	*/
	msgLength := len(msg)
	for i := 1; i <= msgLength; i++ {
		out, flag := <-msg
		if flag {
			fmt.Println("out is ", out)
		}
	}

}
