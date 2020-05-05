package main

import "fmt"

func main() {

	d1 := make(map[string]string)

	fmt.Println("d1 is ", d1)

	d1["a"] = "1"
	d1["b"] = "2"

	d2 := map[string]int{"aaa": 1, "bbb": 2}

	aaa := d2["aaa"]
	fmt.Println("aaa is ", aaa)
	fmt.Println("d1 is ", d1, "d2 is ", d2)

	delete(d2, "bbb")
	fmt.Println("d1 is ", d1, "d2 is ", d2)

}
