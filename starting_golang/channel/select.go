package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// ch1から受信した整数を2倍して送信
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	// ch2から受信した整数を1減算してch3へ送信
	go func() {
		for {
			i := <-ch2
			ch3 <- (i * 2)
		}
	}()

	n := 1
	//ラベル付き文
LOOP:
	for {
		// selectを使うことで複数のチャネルに対する受信・送信でもゴルーチンを停止させずにコントロールできる
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
