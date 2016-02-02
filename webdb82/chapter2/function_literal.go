package main

import "fmt"

// 無名関数
func main() {
	sum(2, 5)
	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)
}

// 関数を変数に
var sum func(i, j int) = func(i, j int) {
	fmt.Println(i + j)
}
