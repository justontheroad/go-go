package main

import "fmt"

func main() {
	var ii *int
	// *ii = 1 // runtime error: invalid memory address or nil pointer dereference
	ii = new(int)
	*ii = 10
	fmt.Println(ii, *ii)

	// slice 分配
	var sl []int = make([]int, 5)
	ps := make([]int, 10, 10)
	fmt.Println(sl, ps, len(sl), cap(sl))
}
