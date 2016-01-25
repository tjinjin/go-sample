package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//	a := 0
	//	b := 1
	//	c := 0
	//
	//	return func() int {
	//		c = a + b
	//		a += b
	//		b += a
	//		return c
	//	}
	f, g := 0, 1
	return func() int {
		f, g = g, f+g
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
