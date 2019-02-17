package main

import "fmt"

func main() {
	fmt.Println("Ready")
	var num0 int
	var num1 int = 10
	var num2 = 20
	fmt.Println(num0, num1, num2)

	num := 31
	//num := 30
	num = 30
	fmt.Println(num)
	num++
	fmt.Println(num)

	//stule
	camelCase := 10
	user_index := 20
	fmt.Println(camelCase, user_index)

	var weight, height = 10, 20
	weight, height = 20, 30

	weight, age := 15, 25
	fmt.Println(weight, height, age)
}
