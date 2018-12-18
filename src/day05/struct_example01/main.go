package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float32
}

func main() {
	var stu Student
	stu.Name = "Jim"
	stu.Age = 18
	stu.score = 88.50

	fmt.Println(stu)
	fmt.Printf("Name: %s  Addr: %p\n", stu.Name, &stu.Name)
	fmt.Printf("Age: %d  Addr: %p\n", stu.Age, &stu.Age)
	fmt.Printf("Name: %f  Addr: %p\n", stu.score, &stu.score)
}
