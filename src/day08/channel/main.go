package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan *student
	stuChan = make(chan *student, 10) //初始化intChan管道
	stu := student{name: "stu01111"}
	stuChan <- &stu

	var stu01 *student
	stu01 = <-stuChan

	fmt.Println(stu01)
}