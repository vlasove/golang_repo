package main

import (
	"golang.org/x/tour/pic"
)

func Converter(i, j int) uint8 {
	return uint8((i + j) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	var temp [][]uint8
	for i := 0; i < dy; i++ {
		inTemp := make([]uint8, dy, dy)
		for j := 0; j < dy; j++ {
			inTemp[j] = Converter(i, j)
		}
		temp = append(temp, inTemp)

	}

	return temp
}

func main() {
	//pic.Show(Pic(8, 8))

	pic.Show(Pic)
}
