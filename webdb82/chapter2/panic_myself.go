/*
panicを自分で起こす場合。
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		if i > len(a) {
			panic(errors.New("index out of range"))
		}
		fmt.Println(a[i])
	}
}
