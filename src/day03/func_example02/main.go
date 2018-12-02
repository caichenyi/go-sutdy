package main

func add(a int, arg ...int) int {
	sum := a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func main() {
	print(add(1, 2, 3, 4, 5, 6))
}
