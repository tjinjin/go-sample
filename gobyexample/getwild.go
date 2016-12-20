//see more: http://qiita.com/takochuu/items/20f639194ac885fb3e1e
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	getWildWords = []string{"G", "e", "t", "W", "i", "l", "d"}
)

func main() {
	// string型のチャネル
	c := make(chan string)
	w := ""

	for {
		// gorutines
		go getwildAndTough(c)
		// チャネルに結果を送信
		w += <-c
		if len(w) == len(getWildWords) {
			fmt.Println(w)
			w = ""
		}
	}
}

func getwildAndTough(c chan string) {
	rand.Seed(time.Now().UnixNano())
	// Intn(n) で0からnまでの数字からランダムで数字を選ぶ
	c <- getWildWords[rand.Intn(len(getWildWords))]
}
