package main

import "fmt"

// どちらも同じ
//arr:= [4]string{"a","b","c","d"}
//arr:= [...]string{"a","b","c","d"}

func fn(arr [4]string) {
	fmt.Println(arr)
}

func main() {
	var arr1 [4]string
	//	var arr2 [5]string

	fn(arr1)
	//	fn(arr2) コンパイルエラー
}
