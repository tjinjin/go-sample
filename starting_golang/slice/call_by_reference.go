package main

import "fmt"

// call by reference => 参照渡し
// call by name => 値渡し

func pow_array(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
}

func pow_slice(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
}

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	// call by value
	pow_array(a)
	// call by reference
	pow_slice(b)
	fmt.Println(a)
	fmt.Println(b)
}
