// サーバを作る

package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!!")
}

func main() {
	// ルーティングの設定
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
