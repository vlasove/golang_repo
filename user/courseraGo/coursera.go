package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	text := "Hello World!"
	fmt.Println(len(text))
	fmt.Println(utf8.RuneCountInString(text))

}
