package main

import (
	"fmt"
	"os"
)

func main()  {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("系统是：%s\n", goos)
	var path string = os.Getenv("PATH")
	fmt.Printf("PATH是：%s\n", path)
}
