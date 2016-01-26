package main

import "fmt"

func main() {
	a, b := 10, 20
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}
}

/*
三項演算子はない
波カッコの省略はできない
*/
