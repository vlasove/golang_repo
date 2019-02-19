package main

import "fmt"

type IPAddr [4]byte

func main() {
	a := []byte{1, 2, 3, 4, 5}
	temp := []rune{}
	for _, v := range a {
		fmt.Println(rune(v))
		temp = append(temp, rune(v), []rune(".")[0])

	}
	str := fmt.Sprintf("%s", a)
	fmt.Println(str)

}
