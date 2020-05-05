package main

import "fmt"

func main() {
	a := 10
	if a > 4 {
		fmt.Printf("a is bigger than 4\n")
	} else {
		fmt.Printf("a is smaller than 4\n")
	}

	if b := 5; b > 10 {
		fmt.Printf("b is bigger than 10\n")
	} else if b > 3 {
		fmt.Printf("b is bigger than 3\n")
	} else {
		fmt.Printf("b is smaller than 3\n")
	}
}
