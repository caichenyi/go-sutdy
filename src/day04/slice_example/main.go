package main

import (
	"fmt"
	"sort"
)

func testSlice() {
	s1 := []int{1, 4, 7, 3, 2}
	sort.Ints(s1)
	fmt.Println(s1)
	//sort.Ints(s1)
	fmt.Println(s1)
	fmt.Println(sort.SearchInts(s1, 3))

}

func main() {
	testSlice()
}