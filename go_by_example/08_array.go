package main

import "fmt"

func main() {
	// 初始化数组的三种方式
	a := [5]int{1, 2, 3, 4, 5}
	b := make([]int, 3)
	var c = [3]int{1, 2, 3}

	fmt.Printf("a is %v\n", a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v] is %v\n", i, a[i])
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%v] is %v\n", i, b[i])
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%v] is %v\n", i, c[i])
	}
}
