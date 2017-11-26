package main

import "fmt"

func main() {
	//keyがintでvalueが[]int
	m := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m)

	// 省略可能
	s := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(s)
}
