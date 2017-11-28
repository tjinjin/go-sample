package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	// ゴルーチンを3つ起動
	// 3つのうちのどれかがchからデータを取り出し処理する
	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		// channelに値を入れる
		ch <- i
		i++
	}
	close(ch)

	// ゴルーチンの完了を3秒待つ
	// これは簡略化のための実装
	time.Sleep(3 * time.Second)
}
