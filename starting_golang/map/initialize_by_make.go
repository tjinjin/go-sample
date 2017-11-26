package main

import "fmt"

func main() {
	// 初期スペース100で作成（メモリ的な話？
	// sliceのcapとは違うので巨大なマップでなければ不要そう
	m := make(map[int]string, 100)
	fmt.Println(m)
}
