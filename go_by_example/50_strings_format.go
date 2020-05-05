package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)   // %+v 的格式化输出内容将包括结构体
	fmt.Printf("%#v\n", p)   // %#v 根据 Go 语法输出值，即会产生该值的源码片段
	fmt.Printf("%T\n", p)    // %T  打印值的类型
	fmt.Printf("%t\n", true) // %t  打印 bool 值

	fmt.Printf("%d\n", 123) // %d 进行标准的十进制格式化。
	fmt.Printf("%b\n", 123) // %b 二进制形式
	fmt.Printf("%c\n", 123) // %c 输出给定整数的对应字符。
	fmt.Printf("%x\n", 123) // %x 16 进制编码

	fmt.Printf("%f\n", 123.45) //  %f 进行最基本的浮点数打印。
	fmt.Printf("%e\n", 123000000000.0)
	fmt.Printf("%E\n", 123000000000.0)

	fmt.Printf("%s\n", "\"strings\"") // %s 进行基本的字符串输出
	fmt.Printf("%q\n", "\"strings\"") // %q 原始字符串输出

	fmt.Printf("%x\n", "hex this") // %x 输出使用 base-16 编码的字符串， 每个字节使用 2 个字符表示
	fmt.Printf("%p\n", &p)         // %p 输出指针值

	fmt.Printf("|%6d|%6d|\n", 123, 45) // %6d 按照指定的宽度和精度进行打印
	fmt.Printf("|%6.2f|%6.2f|\n", 123.0, 45.0)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 要左对齐，使用 - 标志

	s := fmt.Sprintf("a %s", "string") //  Sprintf 则格式化并返回一个字符串而没有任何输出。
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // 可以使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout。
}
