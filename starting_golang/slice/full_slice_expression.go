package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := a[2:4]
	fmt.Printf("len=%d, cap=%d, s1=%d \n", len(s1), cap(s1), s1)
	// 完全スライス式[low:high:max]を指定。maxで容量をコントロールする
	s2 := a[2:4:4]
	fmt.Printf("len=%d, cap=%d, s1=%d \n", len(s2), cap(s2), s2)
	//cap(s3) == (max-low)
	s3 := a[2:4:6]
	fmt.Printf("len=%d, cap=%d, s1=%d \n", len(s3), cap(s3), s3)
}
