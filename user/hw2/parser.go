package main

import (
	"bufio"
	"encoding/json"
	"flag"
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
	Provider    string    `json:"provider"`
	PublushTime time.Time `json:"published_at"`
	Ticker      []string  `json:"tickers"`
}

type Feed struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type CompanyNews struct {
	PayloadCompany []News    `json:"items"`
	PublishTime    time.Time `json:"published_at"`
	Ticker         []string  `json:"tickers"`
}

//1 func --- from json to byteValue  TODO ----- done
//2 func --- from byteValue to []News TODO ----- done
//3 func --- from unsorted []News to sortedByTime []News TODO ----- done
//4 func --- from []News get newsByIndexes func TODO  ---- done
//5 func ---- groupNews TODO --- done
//6 func --- sortByPriority TODO ---- done
//7 func ---- from sorted slice []News to .json file TODO --- done
//8 check ---- check lang style TODO --- done

func openJson(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("can't open file: %s \nmessage : %s", path, err)
	}
	defer jsonFile.Close()
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
	var newsByIndexes []News
	for _, val := range indexes {
		newsByIndexes = append(newsByIndexes, news[val])
	}

	return newsByIndexes
}

func sortByPriority(feed []Feed) {
	for i := 1; i < len(feed); i++ {
		if groupedNews, existGroup := feed[i].Payload.(CompanyNews); existGroup {
			if singleNews, existSngle := feed[i-1].Payload.(News); existSngle {
				if groupedNews.PublishTime == singleNews.PublushTime {
					feed[i-1], feed[i] = feed[i], feed[i-1]
				}
			}
		}

	}

}

func testEq(a, b []string) bool {
	set := make(map[string]bool)

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		set[a[i]] = true
	}
	for i := 0; i < len(a); i++ {
		if !set[b[i]] {
			return false
		}
	}

	return true
}

func groupNews(news []News) []Feed {
	var feed []Feed
	alreadySeen := make(map[int64]bool)
	//feedIndex := 0
Loop:
	for i, val := range news {

		newsIndexer := make([]int, 0, len(news))
		newsIndexer = append(newsIndexer, i)

		if alreadySeen[val.Id] {
			continue Loop
		}

	LoopNextNews:
		for nextNews := i + 1; nextNews < len(news); nextNews++ {
			if val.PublushTime.Day() != news[nextNews].PublushTime.Day() {
				break LoopNextNews
			}

			if testEq(val.Ticker, news[nextNews].Ticker) && !alreadySeen[news[nextNews].Id] {
				alreadySeen[news[nextNews].Id] = true
				newsIndexer = append(newsIndexer, nextNews)

			}

		}

		var feedObj Feed
		if len(newsIndexer) == 1 {
			feedObj = Feed{
				Type:    "news",
				Payload: news[newsIndexer[0]],
			}

		} else {
			tempPayload := getNewsByIndexes(newsIndexer, news)
			feedObj = Feed{
				Type: "company_news",
				Payload: CompanyNews{
					PayloadCompany: tempPayload,
					PublishTime:    news[newsIndexer[0]].PublushTime,
					Ticker:         news[newsIndexer[0]].Ticker,
				},
			}

		}

		feed = append(feed, feedObj)

	}

	return feed

}

func writeToJson(feed []Feed, fileName string) error {
	sample, err := json.MarshalIndent(feed, "", "    ")
	if err != nil {
		return fmt.Errorf("can't marshal some feed: %s", err)
	}

	jsonFinal, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("can't create a file: %s", err)
	}
	defer jsonFinal.Close()

	writer := bufio.NewWriter(jsonFinal)
	defer writer.Flush()

	_, err = writer.Write(sample)
	if err != nil {
		return fmt.Errorf("can't write a result: %s", err)
	}
	return nil

}

func main() {
	path := flag.String("file", "newser.json", "filePath")
	flag.Parse()

	byteValue, err := openJson(*path)
	if err != nil {
		log.Fatal("can't execute openJson func : ", err)
	}
	news, err := byteUnmarshal(byteValue)
	if err != nil {
		log.Fatal("can't unmarshal json after openJson func execution : ", err)
	}

	newsSort(news)
	totalNews := groupNews(news)
	sortByPriority(totalNews)

	err = writeToJson(totalNews, "out.json")
	if err != nil {
		log.Fatal("can't write a result:", err)
	}

}
