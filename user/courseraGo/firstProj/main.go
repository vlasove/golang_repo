package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func unique(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	var prevSeen string
	for in.Scan() {
		txt := in.Text()

		if txt == prevSeen {
			continue
		}
		if txt < prevSeen {
			return fmt.Errorf("file not sorted")
		}
		prevSeen = txt
		fmt.Fprintln(output, txt)

	}
	return nil

}

func main() {
	fmt.Println("Ready")
	err := unique(os.Stdin, os.Stdout)
	if err != nil {
		panic("Blyat")
	}

}
