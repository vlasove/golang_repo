package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Ready")

	var str string
	var hello string = "Hello\n\t"
	world := "WOrld!\n"

	fmt.Println(hello, world, str)

	var rusWorld = "Привет, Пидр"
	fmt.Println(rusWorld)

	fmt.Println(len(rusWorld), utf8.RuneCountInString(rusWorld))

	byteString := []byte(rusWorld)
	helloW := string(byteString)

	fmt.Println(byteString)
	fmt.Println(helloW)

}
