package example1

import "fmt"

func main() {
	sum := 5
	for i:=0; i <= 5; i++{
		fmt.Println(i, "+", sum - i, "=", sum)
		fmt.Printf("%d+%d=%d\n", i, sum - i, sum)
	}
}
