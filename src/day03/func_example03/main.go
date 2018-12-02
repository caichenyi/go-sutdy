package main

import (
	"fmt"
)

func concat(a string, arg ...string) string {
	result := a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return result
}

func main() {
	fmt.Print(concat("hello", " ", "world", " ", "!"))
}
