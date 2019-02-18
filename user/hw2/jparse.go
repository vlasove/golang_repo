package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Bird struct {
	Species     string
	Description string
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("bird.json")
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened bird.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on

	//birdJson := `[{"species" : "pigeon" , "decription" : "asdkASDlikes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal(byteValue, &birds)
	for _, v := range birds {
		fmt.Printf("Species: %s, Description: %s\n", v.Species, v.Description)
	}

}
