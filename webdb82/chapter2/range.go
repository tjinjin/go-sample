package main

import "fmt"

func main() {
	var arr [4]string

	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	arr[3] = "d"

	for i, s := range arr {
		// i = 添字、s = 値
		fmt.Println(i, s)
	}

	z := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(z[2:4])
	fmt.Println(z[0:len(z)])
	fmt.Println(z[:3])
	fmt.Println(z[3:])
	fmt.Println(z[:])
	fmt.Println(sum(1, 2, 3, 4))
}

func sum(nums ...int) (result int) {
	// numsは[]int型
	for _, n := range nums {
		result += n
	}
	return
}
