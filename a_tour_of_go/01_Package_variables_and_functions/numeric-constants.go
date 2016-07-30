package main

import "fmt"

/*
ビット演算。
数値を2進数に直して左や右に動かす
https://teratail.com/questions/18602

*/
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//	fmt.Println(Big)
	//	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
