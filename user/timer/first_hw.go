package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type MegaStruct struct {
	Company   string
	PriceOpen float64
	Value     int
	TimeStamp time.Time
}

func MinMaxSlice(v []float64) (float64, float64) {
	sort.Float64s(v)
	return v[0], v[len(v)-1]
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

		//parsing
		price, _ := strconv.ParseFloat(records[1], 64)
		value, _ := strconv.Atoi(records[2])
		layout := "2006-01-02 15:04:05.000000"
		//str := "2014-11-12T11:45:26.371Z"

		t, _ := time.Parse(layout, records[3])
		if t.Hour() >= 7 {
			arr = append(arr, MegaStruct{
				Company:   records[0],
				PriceOpen: price,
				Value:     value,
				TimeStamp: t,
			})

		}

	}
	return arr

}

func ValueUnique(arr []MegaStruct) []string {
	var unique []string
	for i := 0; i < len(arr); i++ {
		if len(unique) == 0 {
			unique = append(unique, arr[i].Company)
		} else {
			for _, b := range unique {
				if b != arr[i].Company {
					unique = append(unique, arr[i].Company)
				}
			}

		}

	}
	return unique

}

func sepFunc(arr []MegaStruct) ([]MegaStruct, []MegaStruct, []MegaStruct) {
	uniqueName := []string{"SBER", "AMZN", "AAPL"}
	var sberList, aaplList, amznList []MegaStruct
	for i := 0; i < len(arr); i++ {
		if arr[i].Company == uniqueName[0] {
			sberList = append(sberList, MegaStruct{
				Company:   arr[i].Company,
				PriceOpen: arr[i].PriceOpen,
				Value:     arr[i].Value,
				TimeStamp: arr[i].TimeStamp,
			})
		}
		if arr[i].Company == uniqueName[1] {
			amznList = append(amznList, MegaStruct{
				Company:   arr[i].Company,
				PriceOpen: arr[i].PriceOpen,
				Value:     arr[i].Value,
				TimeStamp: arr[i].TimeStamp,
			})

		}
		if arr[i].Company == uniqueName[2] {
			aaplList = append(aaplList, MegaStruct{
				Company:   arr[i].Company,
				PriceOpen: arr[i].PriceOpen,
				Value:     arr[i].Value,
				TimeStamp: arr[i].TimeStamp,
			})

		}

	}
	return sberList, amznList, aaplList

}

type NewStruct struct {
	CompanyName   string
	TimeSeparator time.Time
	openPrice     float64
	maxPrice      float64
	minPrice      float64
	closedPrice   float64
}

func TickerSorter(arr []MegaStruct, interval int) []NewStruct {
	fmt.Println(arr[0].TimeStamp)
	v := arr[0].TimeStamp
	v_end := arr[len(arr)-1].TimeStamp
	currentDay := 30
	v_end_h := v.Hour()
	v_end2_h := v_end.Hour()
	for i, _ := range arr {
		if v.Day() == currentDay {
			if arr[i].TimeStamp.Hour() > v_end_h && arr[i].TimeStamp.Day() == currentDay {
				v_end_h = arr[i].TimeStamp.Hour()
			}

		}
	}
	fmt.Println(v_end_h, v_end2_h)

	start_h := 7
	start_m := interval
	start_s := 0
	start_ns := 0

	//var temp NewStruct
	var new_arr []NewStruct
	tempMax := -1.0
	tempMin := 10000000.0
	openIndex := 0

	for i, _ := range arr {

		now := time.Date(v.Year(), v.Month(), currentDay, start_h, start_m, start_s, start_ns, time.UTC)
		if arr[i].TimeStamp.Day() == currentDay {
			if (arr[i].TimeStamp.Hour()*60 + arr[i].TimeStamp.Minute()) < (now.Hour()*60 + now.Minute()) {
				//fmt.Println(now)
				if arr[i].PriceOpen > tempMax {
					tempMax = arr[i].PriceOpen
				}
				if arr[i].PriceOpen < tempMin {
					tempMin = arr[i].PriceOpen
				}
			} else {
				if arr[i].PriceOpen > tempMax {
					tempMax = arr[i].PriceOpen
				}
				if arr[i].PriceOpen < tempMin {
					tempMin = arr[i].PriceOpen
				}
				timer := time.Date(v.Year(), v.Month(), currentDay, start_h, start_m-interval, start_s, start_ns, time.UTC)
				new_arr = append(new_arr, NewStruct{
					CompanyName:   arr[i].Company,
					TimeSeparator: timer,
					openPrice:     arr[openIndex].PriceOpen,
					maxPrice:      tempMax,
					minPrice:      tempMin,
					closedPrice:   arr[i].PriceOpen,
				})
				tempMax = -1.0
				tempMin = 10000000.0
				openIndex = i + 1
				start_m += interval
			}

		} else if currentDay < 30 {
			currentDay += 1
			start_h = 7
			start_m = interval
			now = time.Date(v.Year(), v.Month(), currentDay, start_h, start_m, start_s, start_ns, time.UTC)
		}
	}

	return new_arr

}

func main() {

	//var temp []float64

	newObj := Filer("trades.csv")
	sber, _, _ := sepFunc(newObj)
	obj := TickerSorter(sber, 5)
	for _, o := range obj {
		fmt.Println(o)
		//if i == 15 {
		//	break
		//}

	}

	//	fmt.Println(len(sber), len(amzn), len(aapl))
	//pop := TickerSorter(sber, 5)
	//for i := 0; i < len(pop); i++ {
	//	fmt.Println(pop[i])
	//}
	//	fmt.Printf("%T and %v\n", sber[0].TimeStamp, sber[0].TimeStamp)
	//uniqueName := []string{"SBER", "AMZN", "AAPL"}

	//for i := 0; i < len(newObj); i++ {

	//	if newObj[i].Company == "SBER" && newObj[i].TimeStamp.Hour() <= 7 && newObj[i].TimeStamp.Minute() <= 4 && newObj[i].TimeStamp.Day() == 30 {
	//		fmt.Println(newObj[i].Company, newObj[i].PriceOpen, newObj[i].Value, newObj[i].TimeStamp)
	//		fmt.Println(newObj[i].TimeStamp.Day())
	//		temp = append(temp, newObj[i].PriceOpen)

	//	}

	//TODO добавить функцию прохода окном с заданной длиной
	//TODO функция сортировки полученных данных
	//TODO функция записи в файл полученных данных
	//}
	//openVal, closedVal := temp[0], temp[len(temp)-1]
	//	minVal, maxVal := MinMaxSlice(temp)
	//fmt.Println("SBER", "TIMER", openVal, maxVal, minVal, closedVal)
	//fmt.Println(unique)

}
