package main

import "fmt"

func one(m int) {
	m = 0
}

func two(n *int) {
	*n = 100
}

func main() {
	x, y := 1, 2
	a, b := &x, &y
	tmp := 0
	tmp = *a
	*a = *b
	*b = tmp
	fmt.Println("a is ", *a, "b is ", *b)

	c := 50
	one(c)
	fmt.Println("c is ", c)
	two(&c) // 传入指针地址会改变对象的值
	fmt.Println("c is ", c)

}
