package main

func main() {
	// intのチャネル
	var ch chan int
	// 受信専用チャネル
	var ch1 <-chan int
	// 送信専用チャネル
	var ch2 chan<- int

	/*
		ch1 = ch0 OK
		ch2 = ch0 OK
		ch0 = ch1 NG
		ch0 = ch2 NG
		ch1 = ch2 NG
		ch2 = ch1 NG
	*/

	// チャネルはキューのデータ構造.FIFO
	// バッファサイズ0
	ch01 := make(chan int)
	// バッファサイズ8
	ch08 := make(chan int, 8)

	ch10 := make(chan int, 10)
	//チャネルに5を送信
	ch10 <- 5
	// チャネルから整数値を受信
	i := <-ch10
}
