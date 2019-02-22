package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

type News struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Provider string    `json:"provider"`
	Time     time.Time `json:"published_at"`
	Tickers  []string  `json:"tickers"`
}

type Feed struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type CompanyNews struct {
	PayloadCompany []News    `json:"items"`
	PublishedAt    time.Time `json:"published_at"`
	Tickers        []string  `json:"tickers"`
}

func main() {
	//var name string
	//flag.StringVar(&name, "file", "", "a file name")
	//flag.Parse()

	news, err := readNews("news.json")
	if err != nil {
		log.Fatal("can't get slice of news:", err)
	}

	resultData := groupNews(news)

	err = writeResult(resultData, "out.json")
	if err != nil {
		log.Fatal("can't write a result:", err)
	}
}

func compareSlice(sl1 []string, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	attribute := make(map[string]bool)
	for i := 0; i < len(sl1); i++ {
		attribute[sl1[i]] = true
	}
	for i := 0; i < len(sl1); i++ {
		if !attribute[sl2[i]] {
			return false
		}
	}
	return true
}

func readNews(file string) ([]News, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("can't read file: %s", err)
	}

	var news []News
	err = json.Unmarshal(byteValue, &news)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal file: %s", err)
	}

	sort.SliceStable(news, func(i, j int) bool {
		return news[i].Time.Before(news[j].Time)
	})

	return news, nil
}

func writeResult(result []Feed, file string) error {
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return fmt.Errorf("can't marshal feed: %s", err)
	}

	jsonResult, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("can't create file: %s", err)
	}
	defer jsonResult.Close()

	writer := bufio.NewWriter(jsonResult)
	defer writer.Flush()

	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("can't write a result: %s", err)
	}
	return nil
}

func groupNews(news []News) []Feed {
	seenNews := make(map[int64]bool)
	getCompanyNews := func(indexs []int) []News {
		result := make([]News, 0, 0)
		for _, in := range indexs {
			result = append(result, news[in])
		}
		return result
	}
	resIndex := 0
	var result []Feed
	for i, in := range news {
		if seenNews[in.ID] {
			continue
		}
		tIndexOfCurrentNews := make([]int, 0, len(news))
		tIndexOfCurrentNews = append(tIndexOfCurrentNews, i)
		for next := i + 1; next < len(news); next++ {
			if in.Time.Day() != news[next].Time.Day() || news[next].Time.Month() != in.Time.Month() {
				break
			}
			if !seenNews[news[next].ID] && compareSlice(in.Tickers, news[next].Tickers) {
				tIndexOfCurrentNews = append(tIndexOfCurrentNews, next)
				seenNews[news[next].ID] = true
			}
		}
		var item Feed
		if len(tIndexOfCurrentNews) == 1 {
			item = Feed{
				Type:    "news",
				Payload: news[tIndexOfCurrentNews[0]],
			}
		} else {
			item = Feed{
				Type: "company_news",
				Payload: CompanyNews{
					PayloadCompany: getCompanyNews(tIndexOfCurrentNews),
					PublishedAt:    news[tIndexOfCurrentNews[0]].Time,
					Tickers:        news[tIndexOfCurrentNews[0]].Tickers,
				},
			}
		}
		result = append(result, item)
		if resIndex > 0 {
			changeSequenceByPrioritety(result, resIndex)
		}
		resIndex += 1
	}
	return result
}

func changeSequenceByPrioritety(result []Feed, index int) {
	if groupNews, ok := result[index].Payload.(CompanyNews); ok { // если одинаковое время у группы новостей и новости, то меняем их местами
		if singleNews, ok2 := result[index-1].Payload.(News); ok2 {
			if groupNews.PublishedAt.Equal(singleNews.Time) {
				result[index-1], result[index] = result[index], result[index-1]
			}
		}
	}
}
