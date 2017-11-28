package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// チャネルのクローズを検出するタイミングがない
	// forの無限ループだとifでcheckができる
	for i := range ch {
		fmt.Println(i)
	}
}
