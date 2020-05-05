package main

import "fmt"

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return recursion(n-1) + recursion(n-2)

}

func main() {
	ret := recursion(16)
	fmt.Println("ret is ", ret)
}
