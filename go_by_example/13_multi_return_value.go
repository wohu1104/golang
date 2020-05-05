package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	a, b := swap(10, 20)
	fmt.Println("after swap value is ", a, b)
	_, second := swap(1, 2)
	fmt.Println("second is ", second)
}
