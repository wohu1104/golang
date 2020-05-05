package main

import "fmt"

type Student struct {
	name string
	age  int
	sex  string
}

func main() {
	var student Student

	student.name = "wohu"
	student.age = 18
	student.sex = "boy"
	fmt.Println("student is ", student)

	student.age = 26
	fmt.Println("student is ", student)
	// 初始化结构体时 Student{name: "ali", age: 18, sex: "boy"} 必须带有前面的 Student
	s := Student{name: "ali", age: 18, sex: "boy"}
	fmt.Println("s name is  ", s.name)

	sp := &s
	fmt.Println("sp is  ", sp)
}
