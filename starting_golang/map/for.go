package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	for k, v := range m {
		// mapから取り出す順は不定
		fmt.Printf("%d => %s\n", k, v)
	}

	//len
	fmt.Println(len(m))

	// delete
	delete(m, 2)
	fmt.Println(m)
}
