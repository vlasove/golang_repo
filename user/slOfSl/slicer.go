package main

import (
	"fmt"
	"strings"
	"time"
)

type Board [][]string

func itemNext(i, j int, symbol string, plato Board) {
	fmt.Println("=================Next Step!=============")
	plato[i][j] = symbol

	for _, v := range plato {
		fmt.Printf("%s \n", strings.Join(v, " | "))

	}
	time.Sleep(3 * time.Second)

}

func main() {
	fmt.Println("Ready")
	plato := Board{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("Start Game")
	itemNext(0, 0, "X", plato)
	itemNext(2, 2, "Y", plato)
	itemNext(1, 2, "X", plato)
	itemNext(1, 0, "Y", plato)
	itemNext(0, 2, "X", plato)
	fmt.Println("Finsih Game")

}
