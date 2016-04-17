//ゴルーチンの終了を待ち合わせる
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://tjinjin.github.io",
		"http://yahoo.co.jp",
		"http://google.com",
	}
	for _, url := range urls {
		// waitGroupに追加
		wait.Add(1)
		// 取得処理をゴルーチンで実行する
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			// waitGroupから削除
			wait.Done()
		}(url)
	}
	// 待ち合わせ waitGroupが0になったら終了する
	wait.Wait()
}
