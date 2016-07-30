package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i //&iでiのメモリアドレスを指定
	fmt.Println(*p)
	fmt.Println(&i)
	*p = 21
	fmt.Println(i)
	fmt.Println(p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

/*
42
0x20818a220
21
0x20818a220
73
*/
