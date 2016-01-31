package main

import "fmt"

/*
func main() {
	n := 10
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
/*/

// fallthrough
func main() {
	n := 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		//		fallthrough
	case 1:
		n = n - 1
		fmt.Println(n)
	}
	fmt.Println(n)
}

//func main() {
//	n := 10
//	switch {
//	case n%15 == 0:
//		fmt.Println("FizzBuzz")
//	case n%5 == 0:
//		fmt.Println("Buzz")
//	case n%3 == 0:
//		fmt.Println("Fizz")
//	default:
//		fmt.Println(n)
//	}
//}
