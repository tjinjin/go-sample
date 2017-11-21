package main

import "fmt"

func inifinit_loop() {
	s := []string{"Apple", "Banana", "Cherry"}
	// 無限ループする
	for i := 0; i < len(s); i++ {
		fmt.Printf("[%d] => %s\n", i, v)
		s = append(s, "Mellon")
	}
}

func main() {
	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	// 無限ループにはならない
	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
		s = append(s, "Mellon")
	}
	fmt.Println(s)
}
