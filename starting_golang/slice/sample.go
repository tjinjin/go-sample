package main

import "fmt"

func main() {
	// slice
	s := make([]int, 10)
	fmt.Println(s)
	// array
	var a [10]int
	fmt.Println(a)

	// 要素数
	fmt.Println(len(s))

	// 容量
	fmt.Println(cap(s))

	// 第3引数を渡すと要素数と容量をずらすことができる
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
