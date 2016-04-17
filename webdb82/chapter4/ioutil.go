// ファイルの読み書きに特化

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// ファイルにメッセージを書き込む
	hello := []byte("hello world\n")
	// 0666はファイルのパーミッション
	err := ioutil.WriteFile("./file.txt", hello, 0666)
	if err != nil { // エラー処理
		log.Fatal(err)
	}

	// ファイルの中身を全て読みだす
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil { // エラー処理
		log.Fatal(err)
	}
	fmt.Print(string(message)) // 文字列にして表示
}
