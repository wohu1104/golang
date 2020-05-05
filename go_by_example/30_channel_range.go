package main

import "fmt"

func main() {

	ch := make(chan int, 4)
	go func() {
		for i := 0; i <= 3; i++ {
			ch <- i
			fmt.Println("send ", i)
		}
		close(ch) // 关闭函数非常重要,若不执行close(ch), 下面的循环将永不结束，从而造成死锁
	}()

	for c := range ch {
		fmt.Println("c is ", c)
	}

}
