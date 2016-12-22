package main

import "fmt"

func main() {
	fmt.Println("aaa")

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
