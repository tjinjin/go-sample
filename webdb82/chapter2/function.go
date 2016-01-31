package main

import "fmt"

/*
func hello() {
	fmt.Println("hello world")
}

func main() {
	hello()// hello world
}
*/
/*
func sum(i, j int) {
	fmt.Println(i + j)
}

func main() {
	sum(1, 2) //3
}
*/
func sum(i, j int) int {
	return i + j
}

func main() {
	n := sum(1, 2)
	fmt.Println(n) //3
}
