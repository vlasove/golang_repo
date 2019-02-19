package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)

}

type MyFloat float64

func (f MyFloat) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("%v and %T\n", i, i)
}

func main() {
	var i I
	i = &T{"Hello\a"}
	describe(i)
	i.M()

	i = MyFloat(20.5)
	describe(i)
	i.M()
}
