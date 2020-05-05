package main

import (
	"fmt"
	"time"
)

func main() {
	a := 10
	switch a {
	case 10:
		fmt.Printf("a is equal 10\n")
	case 5:
		fmt.Printf("a is equal 5\n")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // case 语句中，可以使用逗号来分隔多个表达式
		fmt.Println("it's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。
	case t.Hour() < 12:
		fmt.Println(" it's morning")
	case t.Hour() > 12:
		fmt.Println("it's afternoon")
	}
}
