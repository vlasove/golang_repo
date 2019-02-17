package main

import "fmt"

func printSlice(s []uint64) {
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))
}
func maximize(s []uint64) {
	var i uint64
	for i = 0; i < 1<<32-1; i++ {
		s = append(s, i)
		if cap(s)%8 == 0 {
			printSlice(s)
		}
	}

}

func main() {
	var ex []uint64
	maximize(ex)

}
