package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	//func (v Vertex) Abs() float64 {
	fmt.Println(v.X)
	fmt.Println(v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	//	fmt.Println(Vertex{3, 4}.Abs())
}
