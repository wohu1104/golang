package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/data1", d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件。打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
	f, err := os.Create("/tmp/data2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) //  Write 字节切片
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString 写入字符串
	n3, err := f.WriteString("world\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // 调用 Sync 将缓冲区的数据写入硬盘

	w := bufio.NewWriter(f) // bufio 还提供了的带缓冲的 Writer。
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// bufio 还提供了的带缓冲的 Writer。
	w.Flush()
}
