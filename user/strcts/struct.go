package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 4})

	v := Vertex{19, 20}
	v.X = 4
	fmt.Println(v.X, v.Y)

	p := &v
	p.X = 19
	fmt.Println(v.X, v.Y)

}
