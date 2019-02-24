package main

import (
	"flag"
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

var path = flag.String("path", "conf.yaml", "Path to .yaml file")

func main() {
	flag.Parse()
	config, err := yaml.ReadFile(*path)
	if err != nil {
		fmt.Println("can't open .yaml : ", err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.Get("enabled"))

}
