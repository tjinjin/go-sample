// 複数のゴルーチン間でデータのやりとりをしたい場合チャネルで
// メッセージパッシングに寄ってデータのやりとりができる
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://tjinjin.github.io",
		"http://yahoo.co.jp",
		"http://google.com",
	}
	statusChan := make(chan string)
	for _, url := range urls {
		// 取得処理をゴルーチンで実行する
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// <- でチャネルの送信を示す
			statusChan <- res.Status
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		// チャネルを読みだす
		fmt.Println(<-statusChan)
	}
}
