package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var path = flag.String("path", "conf.json", "File to open as .json")

type config struct {
	Enabled bool
	Path    string
}

func main() {
	flag.Parse()
	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("can't open file.json: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("can't decode file : ", err)
	}

	fmt.Println(conf)

}
