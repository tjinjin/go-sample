package main

import "fmt"

func main() {
	b := byte(255)
	fmt.Println(b) // 255
	b = b + 1
	fmt.Println(b) // 0

	n := -1
	c := byte(n)
	fmt.Println(c) // 255
}
