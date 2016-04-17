package main

import "fmt"

func main() {
	var i uint8 = 3
	var j uint32 = uint32(i) // uint8 -> 32
	fmt.Println(j)

	var s string = "abc"
	var b []byte = []byte(s) // string -> []byte
	fmt.Println(b)

	//cannot convert "a" (type string) to type int
	a := int("a")
}
