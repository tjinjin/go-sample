package main

import "fmt"

func later() func(string) string {
	// ローカル変数だけど、clojure内の関数として扱われる
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome!"))
}
