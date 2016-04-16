/*
ゼロ除算や範囲外の配列にアクセスした場合に通常のエラーではなくpanicを
という方法でエラー発生する。
その場合recover()という関数でエラーを取得する
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			// runtime error
			log.Fatal(err)
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10])

}
