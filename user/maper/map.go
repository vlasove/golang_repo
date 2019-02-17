package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	temp := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		//fmt.Println(word)
		if _, err := temp[word]; err == false {
			temp[word] = 1
			continue
		} else {
			temp[word]++
		}

	}
	return temp
}

func main() {
	wc.Test(WordCount)

}
