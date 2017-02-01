package main

import "fmt"

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)

	//section with switch/case gives unexpected output
	for i := 1; i <= 30; i++ {
		switch {
		case i%FIZZ == 0 && i%BUZZ == 0:
			fmt.Printf("%d fizz\t", i%3)
			fmt.Printf("%d buzz\t", i%5)
		case i%FIZZ == 0:
			fmt.Printf("%d fizz\t", i%3)
		case i%BUZZ == 0:
			fmt.Printf("%d buzz\t", i%5)
		}
		fmt.Printf("\t%d\n", i)
	}

	fmt.Printf("now towards the if/else\n")

	//
	for i := 1; i <= 30; i++ {
		switch {
		case i%FIZZ == 0 && i%BUZZ == 0:
			fmt.Printf("%d fizz\t", i%3)
			fmt.Printf("%d buzz\t", i%5)
		case i%FIZZ == 0:
			fmt.Printf("%d fizz\t", i%3)
		case i%BUZZ == 0:
			fmt.Printf("%d buzz\t", i%5)
		}
		fmt.Printf("\t%d\n", i)
	}

	fmt.Printf("now towards the if/else\n")

	//section with if/else works as expected
	for i := 1; i <= 30; i++ {
		if i%FIZZ == 0 {
			fmt.Printf("%d fizz\t", i%3)
		}
		if i%BUZZ == 0 {
			fmt.Printf("%d buzz\t", i%5)
		}
		fmt.Printf("\t%d\n", i)
	}
}
