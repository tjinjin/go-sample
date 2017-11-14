package main

import "fmt"

func main() {
	a1 := [3]int{}
	fmt.Println(a1)
	a2 := [...]int{1, 2, 3}
	fmt.Println(a2)

	// copy
	a3 := [3]int{1, 2, 3}
	a4 := [3]int{4, 5, 6}
	a3 = a4
	a3[0] = 0
	a3[2] = 0
	fmt.Printf("%v", a3)
	fmt.Printf("%v", a4)
}
