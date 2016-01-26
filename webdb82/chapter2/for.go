package main

import "fmt"

/*
func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
*/

/*
whileの変わりの書き方
func main() {
	n := 0
	for n < 10 {
		fmt.Printf("n = %d\n", n)
		n++
	}
}
*/

/*
infinite loop
for {

	doSomething()
}

*/

//break, continue
func main() {
	n := 0
	for {
		n++
		if n > 10 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
