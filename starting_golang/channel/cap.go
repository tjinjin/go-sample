package main

import "fmt"

func main() {
	// channelのバッファサイズは初期化時にしか変更できない
	ch := make(chan string)
	fmt.Println(cap(ch))

	ch2 := make(chan string, 3)
	fmt.Println(cap(ch2))
}
