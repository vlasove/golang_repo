package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type MegaStruct struct {
	Company   string
	PriceOpen float64
	Value     int
	TimeStamp time.Time
}

func Filer(path string) []MegaStruct {
	var arr []MegaStruct
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvFile)
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		price, _ := strconv.ParseFloat(records[1], 64)
		value, _ := strconv.Atoi(records[2])
		layout := "2006-01-02 15:04:05.000000"
		//str := "2014-11-12T11:45:26.371Z"
		t, _ := time.Parse(layout, records[3])
		arr = append(arr, MegaStruct{
			Company:   records[0],
			PriceOpen: price,
			Value:     value,
			TimeStamp: t,
		})
	}
	return arr

}
func main() {

	newObj := Filer("trades.csv")
	for i := 0; i < len(newObj); i++ {
		fmt.Println(newObj[i].Company, newObj[i].PriceOpen, newObj[i].Value, newObj[i].TimeStamp)
		fmt.Println(newObj[i].TimeStamp.Minute())
		//TODO добавить функцию прохода окном с заданной длиной
		//TODO функция сортировки полученных данных
		//TODO функция записи в файл полученных данных
	}

}
