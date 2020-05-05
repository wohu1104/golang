package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i is %v\n", i)
	}

	j := 10
	for j > 0 {
		j--
		if j < 3 {
			fmt.Printf("j is %v\n", j)
		}
	}

	k := 0
	for {
		k++
		fmt.Printf("k is %v\n", k)
		if k >= 5 {
			break
		}
	}
}
