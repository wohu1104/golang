package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "hello"
	}()

	select {
	case res := <-c1:
		fmt.Println("res is", res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "world"
	}()

	select {
	case res := <-c2:
		fmt.Println("res is ", res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
