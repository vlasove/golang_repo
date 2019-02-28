package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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

func main() {
	flag.Parse()
	file, err := os.Open(*path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()

	scanner := bufio.NewScanner(file)
	in := make(chan string)
	//out := make(chan string)
	go func() {
		for scanner.Scan() {
			in <- scanner.Text()
		}
		close(in)
	}()

	fmt.Println(t.Truncate(240 * time.Minute))
	fmt.Println(t.Add(-180 * time.Minute).Truncate(240 * time.Minute).Add(180 * time.Minute))

}
