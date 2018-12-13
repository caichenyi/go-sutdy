package main

func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}

func main() {
	result := factor(5)
	print(result)
}
