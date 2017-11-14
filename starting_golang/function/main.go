package main

import "fmt"

// 返り値に変数を割り当てる
func doSomething() (a int) {
	return
}

// doSomethingの別バージョン
func doSomething2() int {
	var a int
	return a
}

func doSomething3() (x, y int) {
	y = 5
	return
}

// 3のちゃんと書いた版。3の方がスッキリ
func doSomething4() (int, int) {
	var x, y int
	y = 5
	return x, y
}

func main() {
	fmt.Println(doSomething())
	fmt.Println(doSomething2())
	fmt.Println(doSomething3())
	fmt.Println(doSomething4())
}
