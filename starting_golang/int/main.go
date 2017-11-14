package main

import (
	"fmt"
	"math"
)

func main() {
	// ラップアラウンドが怒らないように定数を利用する
	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	wrapArround()

	fmt.Println("uint8 overflow test")
	overflowUint8Max()
}

func wrapArround() {
	n := 256
	// 256のままのはずが…
	fmt.Println(byte(n)) // 0
}

func overflowUint8Max() {
	var a, b, c, d uint8
	a = 255
	b = 0
	c = 1
	d = math.MaxUint8

	fmt.Println(a, "+", b, "->", a+b)
	fmt.Println(a, "+", c, "->", a+c)
	fmt.Println(d, "+", b, "->", d+b)
	fmt.Println(d, "+", c, "->", d+c)
	if doSomething(d, c) {
		fmt.Println(d, "+", c, "->", d+c)
	}
}

func doSomething(a, b uint8) bool {
	// x - a < b
	// x < a + b
	if (math.MaxUint8 - a) < b {
		// オーバフローした場合はfalse
		return false
	} else {
		sum := a + b
		fmt.Println(sum)
		return true
	}
}
