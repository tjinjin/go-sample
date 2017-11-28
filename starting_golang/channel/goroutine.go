package main

import "fmt"

// 受信用
func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
func main() {
	ch := make(chan int)

	// goroutineの生成
	// goステートメントに関数を渡すだけ
	go receiver(ch)

	i := 0
	for i < 10000 {
		ch <- i
		i++
	}
}
