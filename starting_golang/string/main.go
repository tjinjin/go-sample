package main

import "fmt"

func main() {
	s := `
	Goの
	Raw文字列リテラルによる
	複数行に渡る
	文字列
	`
	fmt.Printf("%v", s)
}
