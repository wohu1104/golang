package main

import "fmt"

func main() {
	a := []int{}
	b := make([]int, 5, 10)

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	c := append(b, 1)
	c[0] = 0
	c[1] = 1
	c[2] = 2
	fmt.Println("c is ", c)
	d := append(b, 3, 4)

	f := make([]int, 5, 10)
	copy(f, d)
	fmt.Println("d is ", d)
	fmt.Println("f is ", f)

}
