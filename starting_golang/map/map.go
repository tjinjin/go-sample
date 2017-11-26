package main

import "fmt"

func main() {
	// int型のkeyとstring型の値を保持する
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	fmt.Println(m)
}
