package main

import (
	"errors"
	"fmt"
	"log"
)

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("divided by zero")
	}
	return i / j, nil
}

/*
名前付き戻り値
func div(i, j int) (result int, err error) {
	if j ==0 {
		err = errors.New("divided by zero")
		return //return 0, errと同様
	}
	result = i /j
	return //return result, nilと同様
}
*/

func main() {
	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
