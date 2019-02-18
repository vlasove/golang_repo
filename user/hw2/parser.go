package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type News struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	PublushTime time.Time `json:"published_at"`
	Provider    string    `json:"provider"`
	Ticker      []string  `json:"tickers"`
}

//1 func --- from json to byteValue  TODO
//2 func --- from byteValue to []News TODO
//3 func --- form non grouped []News to grouped []News
//4 func --- sort grouped []News by time and Grouping
//5 func --- from grouped and sorted []News to json

func openJson(path string) []byte {
	jsonFile, errOpen := os.Open(path)
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if errOpen != nil {
		fmt.Println(errOpen)
	}
	fmt.Printf("Successfully Opened %v \n", path)
	byteValue, errRead := ioutil.ReadAll(jsonFile)
	if errRead != nil {
		fmt.Println(errRead)
	}
	return byteValue

}

func byteUnmarsh(valByte []byte) []News {
	var news []News
	json.Unmarshal(valByte, &news)
	return news
}

func main() {

	// Open our jsonFile
	///jsonFile, err := os.Open("news.json")
	///defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	///if err != nil {
	///		fmt.Println(err)
	///	}
	///fmt.Println("Successfully Opened news.json")
	///byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on

	byteValue := openJson("news.json")
	//birdJson := `[{"species" : "pigeon" , "decription" : "asdkASDlikes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	///var news []News
	///json.Unmarshal(byteValue, &news)
	news := byteUnmarsh(byteValue)
	for i, v := range news {
		fmt.Println("==============================================NEWS==========================================================")
		fmt.Printf("Id : %v  \n Title : %v \n", v.Id, v.Title)
		fmt.Printf("Body: \n %v\n", v.Body)
		fmt.Printf("Published at %v\nProvider is %v\nTickers are %+v\n", v.PublushTime, v.Provider, v.Ticker)
		fmt.Printf("News #%v\n", i)

		if i > 10 {
			break
		}
	}
}
