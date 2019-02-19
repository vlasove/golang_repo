package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

//1 func --- from json to byteValue  TODO ----- done
//2 func --- from byteValue to []News TODO ----- done
//3 func --- from unsorted []News to sortedByTime []News TODO ----- done
//4 func --- from []News get newsByIndexes func TODO  ---- done
//5 func ---- groupNews TODO
//6 func --- sortByPriority TODO
//7 func ---- from sorted slice []News to .json file TODO
//8 check ---- check lang style TODO

func openJson(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, fmt.Errorf("can't open file: %s \nmessage : %s", path, err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("can't read (.ReadAll) from json : %s \nmessage : %s", path, err)
	}
	return byteValue, nil

}

func byteUnmarshal(valByte []byte) ([]News, error) {
	var news []News
	err := json.Unmarshal(valByte, &news)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal file : %s", err)
	}

	return news, err
}

func newsSort(news []News) {
	sort.SliceStable(news, func(i, j int) bool {
		return news[i].PublushTime.Before(news[j].PublushTime)
	})

}

func getNewsByIndexes(indexes []int, news []News) []News {
	newsByIndexes := make([]News, 0)
	for _, val := range indexes {
		newsByIndexes = append(newsByIndexes, news[val])
	}

	return newsByIndexes
}

func groupNews(news []News) {}

func main() {

	byteValue, err := openJson("news.json")
	if err != nil {
		log.Fatal("can't execute openJson func : ", err)
	}
	news, err := byteUnmarshal(byteValue)
	if err != nil {
		log.Fatal("can't unmarshal json after openJson func execution : ", err)
	}

	newsSort(news)
	//done by this point

	for i, v := range news {
		fmt.Println("==============================================NEWS==========================================================")
		fmt.Printf("Published at %v\nProvider is %v\nTickers are %+v\n", v.PublushTime, v.Provider, v.Ticker)
		fmt.Printf("News #%v\n", i)

	}

}
