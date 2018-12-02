package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100) //随机生成1-100的数
	var input int
	for {
		fmt.Scanf("%d\n", &input)
		flag := false
		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("太大")
		case input < n:
			fmt.Println("太小")
		}
		if flag {
			break
		}
	}
}
