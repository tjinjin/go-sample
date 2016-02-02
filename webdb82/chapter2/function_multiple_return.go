package main

import "fmt"

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y) //4,3

	//	x = swap(x, y) //compile error 引数が合わない

	x, _ = swap(x, y) // 第二を無視
	fmt.Println(x)    // 3

	swap(x, y)
}
