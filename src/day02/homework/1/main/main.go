package main

import "fmt"

func sushu(num int) {
    count := 1
    for i := 2; i < num/2; i++ {
        if (num % i == 0) {
            count ++
        }
    }
    if (count == 1) {
        fmt.Printf("素数：%d\n", num)}
}

func main() {
    for i := 101; i <= 200; i++ {
	    go sushu(i)
    }
}