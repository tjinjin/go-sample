package main

import "fmt"

func Print(value interface{}) {
	// sにはCastした結果、okには真偽値
	s, ok := value.(string) // Tyep Assertion
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}

func main() {
	Print("abc")
	Print(10)
}
