package main

import "fmt"

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)

}

func sum(nums ...int) int {
	fmt.Println("nums is ", nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("total is ", total)
	return total
}
