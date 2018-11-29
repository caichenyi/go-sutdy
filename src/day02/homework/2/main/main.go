package main

import "fmt"

func shuixianhuashu(num int) {
    a := num/100
    b := num/10%10
    c := num%10
    if (a * a * a + b * b * b + c * c * c == num) {
        fmt.Printf("水仙花数：%d\n", num)
    }
}

func main() {
    for i:= 100; i <= 999; i++ {
        go shuixianhuashu(i)
    }
}