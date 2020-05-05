package main

import "fmt"

// Go 支持在结构体类型中定义方法 。
type rect struct {
	width  int
	height int
}

func (r *rect) area() int { // area 方法有一个接收器类型 rect
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
