package main

import "fmt"

type ErrNegativSqrt struct {
	what string
}

func (e ErrNegativSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %s", e.what)
}

func Sqrt(x float64) (float64, error) {
	if error != nil {
		fmt.Println(error)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
