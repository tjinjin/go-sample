// ゴルーチンを使った場合
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://tjinjin.github.io",
		"http://yahoo.co.jp",
		"http://google.com",
	}
	for _, url := range urls {
		// 取得処理をゴルーチンで実行する。無名関数
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
		}(url)
	}
	// main()が終わらないように待ち合わせる
	//mainはゴルーチンの終了を待ってはくれない
	time.Sleep(time.Second)
}
