package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(time.Second * 30)
	fmt.Println("TIme is out")
	os.Exit(0)
}
