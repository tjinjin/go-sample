package main

import "fmt"

func sendChannel() {
	ch := make(chan int, 1)
	// channelのclose
	close(ch)
	ch <- 1 // error
}

func receiveChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	// errorにならない
	i := <-ch
	fmt.Println(i)
	j := <-ch
	fmt.Println(j)
	k := <-ch
	fmt.Println(k)
	// 型の初期値が引ける
	l := <-ch
	fmt.Println(l)
}

func isClosed() {
	ch := make(chan int)
	close(ch)
	// iが値, okがチャネルのバッファ内が空でかつクローズされている状態の場合にfalse
	_, ok := <-ch
	fmt.Println(ok)
}

func main() {
	//sendChannel()
	receiveChannel()
	isClosed()
}
