package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var path = flag.String("file", "trades.csv", "Path to .csv file")

func firstStage(path *string) <-chan string {
	file, err := os.Open(*path)
	defer file.Close()
	if err != nil {
		fmt.Println("can't open file : ", err)
	}

	out := make(chan string)
	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			out <- scanner.Text()
		}
		close(out)
	}()

	return out

}



func fetchTrades(in <-chan string) []<-chan string, <-chan error{
	
}

func main() {
	flag.Parse()
	file, err := os.Open(*path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	in := make(chan string)
	//out := make(chan string)
	go func() {
		for scanner.Scan() {
			in <- scanner.Text()
		}
		close(in)
	}()

	fetchTrades(in chan<- string) []

}
