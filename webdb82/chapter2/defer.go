package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./error.go")
	if err != nil {

		// エラー処理
		fmt.Println("error")
	}
	// 正常処理
	// deferがあるので、main関数を抜ける際に必ず実行される
	defer file.Close()
}
