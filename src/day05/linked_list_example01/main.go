package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float32
	Next *Student
}

func main() {
	var head Student
	head.Name = "Jim"
	head.Age = 18
	head.Score = 88.5

	var stu1 Student
	stu1.Name = "Ann"
	stu1.Age = 16
	stu1.Score = 99.5
	//stu1.Next = nil   默认为nil

	head.Next = &stu1

	var p *Student = &head

	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}