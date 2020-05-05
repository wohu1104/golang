package main

// Go 提供了一个 flag 包，支持基本的命令行标志解析。
import (
	"flag"
	"fmt"
)

func main() {
	/*
		基本的标记声明仅支持字符串、整数和布尔值选项。
		这里我们声明一个默认值为 "foo" 的字符串标志 word 并带有一个简短的描述。
		这里的 flag.String 函数返回一个字符串指针（不是一个字符串值）
	*/
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。 注意在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析
	flag.Parse()

	// 这里我们将仅输出解析的选项以及后面的位置参数。 注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用， 从而得到选项真正的值。
	fmt.Println("word", *wordPtr)
	fmt.Println("numb", *numbPtr)
	fmt.Println("fork", *boolPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
	/*
		注意，flag 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）。

		wohu@wohu:~/gocode/src$ go run 36_cmd_line_flag.go   -word=opt a1 a2 a3 -numb=7
		word opt
		numb 42
		fork false
		svar bar
		tail [a1 a2 a3 -numb=7]
	*/

	/*
		使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。
		wohu@wohu:~/gocode/src$ go run 36_cmd_line_flag.go   -h
		Usage of /tmp/go-build190236654/b001/exe/36_cmd_line_flag:
		-fork
				a bool
		-numb int
				an int (default 42)
		-svar string
				a string var (default "bar")
		-word string
				a string (default "foo")
		exit status 2
		wohu@wohu:~/gocode/src$

	*/

}
