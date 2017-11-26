package main

import "fmt"

func printer(v string, ok bool) {
	if !ok {
		fmt.Println("Not exist index. value:", v)
		return
	}
	fmt.Println("string value:", v, ".")
}

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println("1:", m[1])
	// stringの初期値の空文字が入ってる
	fmt.Println("9:", m[9])

	s := map[int]int{1: 0}
	fmt.Println("1:", s[1])
	//区別がつかない
	fmt.Println("7:", s[7])

	// okにはキーが存在するかがboolで表現される
	r1, ok := m[1]
	printer(r1, ok)
	r2, ok := m[9]
	printer(r2, ok)
	//_, ok := m2[3] // ok == true

	// 頻出
	if _, ok := m[1]; ok {
		fmt.Println("1 aru!")
	}
}
