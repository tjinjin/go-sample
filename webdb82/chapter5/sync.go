// 同期処理
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
	for _, url := range urls {
		// 前のレスポンスが返ってくるまで、新しいリクエストはしない
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Println(url, res.Status)
	}
}
