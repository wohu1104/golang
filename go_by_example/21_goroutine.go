package main

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println("s ", i)
	}
}

func main() {
	f("hello")

	go f("going") // Go 协程

	go func(msg string) { // 匿名函数协程
		fmt.Println(msg)
	}("world")

	var input string
	fmt.Scanln(&input) // Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Println("done")

}
