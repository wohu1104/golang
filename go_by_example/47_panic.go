package main

import "os"

func main() {
	panic("a issue")
	/*
		panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。
		如果创建新文件时遇到意外错误该如何处理？这里有一个很好的 panic 示例。
	*/
	_, err := os.Create("/tmp/data")
	if err != nil {
		panic(err)
	}

	/*
		注意，与某些使用 exception 处理错误的语言不同， 在 Go 中，通常会尽可能的使用返回值来标示错误。
	*/
}
