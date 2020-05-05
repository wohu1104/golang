package main

import (
	"fmt"
	"reflect"
)

const s = "hello word"

func main() {
	const PI float32 = 3.0 // 3.14 constant 3.14 truncated to integer
	/*
			如果 浮点数字面值常量 能被精确的表示为int, 那么可以当作数字类型使用(untyped).
		    如果 浮点数字面值常量 不能精确转换为 int , 应该会有编译错误.
	*/
	fmt.Println("PI is ", PI)
	fmt.Println("s is ", s)
	fmt.Println(int64(PI))
	t1 := reflect.ValueOf(PI)
	fmt.Printf("t1 type is %v\n", t1.Type())
}
