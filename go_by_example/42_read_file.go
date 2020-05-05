package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("/tmp/data") // 将文件内容读取到内存中。
	check(err)
	fmt.Println(string(data))

	// 通常会希望对文件的读取方式和内容进行更多控制。
	// 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值。
	f, err := os.Open("/tmp/data")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // 从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	check(err)

	fmt.Printf("%d bytes:%s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0) // 也可以 Seek 到一个文件中已知的位置，并从这个位置开始 Read。
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2) // io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件。
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	/*
		bufio 包实现了带缓冲的读取， 这不仅对于很多小的读取操作能够提升性能， 也提供了很多附加的读取函数。

		bufio 包实现了带缓冲的 Reader， 这对于提高读取大量小文件的效率， 以及它提供的其他读取方法可能都有用。
	*/
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // 任务结束后要关闭这个文件 （通常这个操作应该在 Open 操作后立即使用 defer 来完成）
}
