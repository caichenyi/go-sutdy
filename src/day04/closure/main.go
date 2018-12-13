package main

import "fmt"

func Adder() func(int) int {
	var x int
	fmt.Println(x)
	return func(d int) int {
		x += d
		fmt.Printf("x=%d, d=%d\n", x, d)
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}