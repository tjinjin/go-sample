package main

import "fmt"

//「引数も戻り値もない関数」を戻り値として返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	fmt.Printf("%#v\n", func(x, y int) int { return x + y })       // メモリのアドレス
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3)) // 関数結果

	f := returnFunc()
	f()

	//returnFuncの戻り値である関数をそのまま直接呼び出す
	returnFunc()()
}
