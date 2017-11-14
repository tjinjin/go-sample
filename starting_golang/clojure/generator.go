package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	ints := integers()

	// 同じクロージャで同じ変数を参照
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3

	// 別のクロージャでは変数が初期化される
	otherInts := integers()
	fmt.Println(otherInts()) // 1
}
