// timeout処理のよくあるパターン
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getStatus(urls ...string) chan string {
	statusChan := make(chan string)
	for _, url := range urls {
		go func() {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			statusChan <- res.Status
		}()
	}
	return statusChan
}

func main() {
	statusChan := getStatus(
		"http://tjinjin.github.io",
		"http://yahoo.co.jp",
		"http://google.com",
	)

	// 1 秒後に値が読み出せる channel
	timeout := time.After(time.Second)
	//ラベル
	// 1秒たったときにtimeが読み出せたらループを抜ける
LOOP:
	for {
		select {
		case status := <-statusChan:
			fmt.Println(status) // 受信したデータを表示
		case <-timeout:
			break LOOP // この for-select を抜ける
		}
	}
}
