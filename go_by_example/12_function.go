package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	ret := add(3, 10)
	fmt.Println("ret is ", ret)
}
