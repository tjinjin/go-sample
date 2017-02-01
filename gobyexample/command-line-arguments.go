package main

import (
	"fmt"
	"os"
)

func main() {
	// 引数を取得
	argsWithProg := os.Args
	// 純粋に引数のみ取得
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
