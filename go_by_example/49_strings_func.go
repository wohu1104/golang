package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	/*
		这是一些 strings 中有用的函数例子。 由于它们都是包的函数，而不是字符串对象自身的方法，
		这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。
	*/
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "es"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("hasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "t"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("hello", "l", "a", -1))
	p("Replace:   ", s.Replace("hello", "l", "a", 1))
	p("Split:     ", s.Split("a-b-c-d", "-"))
	p("ToLower    ", s.ToLower("HELLO"))
	p("ToUpper    ", s.ToUpper("hello"))
	p()

	// 虽然不是 strings 的函数，但仍然值得一提的是， 获取字符串长度（以字节为单位）以及通过索引获取一个字节的机制。
	p("Len:       ", len("hello"))
	p("Char:      ", "hello"[0])
	// 注意，上面的 len 以及索引工作在字节级别上。 Go 使用 UTF-8 编码字符串，因此通常按原样使用。

}
