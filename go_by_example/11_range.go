package main

import "fmt"

func main() {
	s := "abcd"

	for i, v := range s {
		fmt.Println("i is ", i, "v is ", v)
	}

	d := map[string]int{"a": 1, "b": 2}
	for k, v := range d {
		fmt.Println("k is ", k, "v is ", v)
	}

	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println("i is ", i, "v is ", v)
	}
}
