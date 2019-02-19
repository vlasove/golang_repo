package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Printf("%v and %v\n", v.X, v.Y)
	fmt.Println(v.Abs())
}
