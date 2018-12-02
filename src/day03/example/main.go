package main

import "fmt"

// 指针

func main() {
	var a int = 10
	fmt.Println(a)  //值
	fmt.Println(&a) //内存地址

	var b *int = &a
	fmt.Println(b)  //内存地址
	fmt.Println(*b) //取到内存地址的值

}
