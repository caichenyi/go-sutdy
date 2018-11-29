package main

import (
	"../calc"
	"fmt"
)

func main() {
	sum := calc.Add(100, 200)
	sub := calc.Sub(100, 200)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}
