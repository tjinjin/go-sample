package main

import "fmt"

func main() {
	var s []string
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c", "d")
	s1 := []string{"a", "b"}
	s2 := []string{"c", "d"}
	s1 = append(s1, s2...) //...がないとだめ
	fmt.Println(s)
	fmt.Println(s1)
}

// var s string sliceの宣言
/*
s := []string{"a", "b", "c", "d"}
fmt.Println(s[0])
*/
