package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	fmt.Println(pos)
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*1),
		)
	}
	fmt.Println(pos(10))
}
