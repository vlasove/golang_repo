package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - 8
	return
}

var python, c, java bool

var i, j int = 1, 2

func main() {
	var t int
	fmt.Println(split(15))
	fmt.Println(python, java, c, t)

	var kek, lol = false, "NONONONONO"
	fmt.Println(i, j, kek, lol)
}
