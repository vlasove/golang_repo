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
type OutputStruct struct {
	Company    string  //default:""
	TimeStamp  string  //default: ""
	openPrice  float64 //default:0
	highPrice  float64 //default:0
	lowPrice   float64 //default:0
	closePrice float64 //default:0
	checkMe    bool    //default: false
}

func DefaultOutput(Output OutputStruct) OutputStruct {
	Output.Company = ""
	Output.TimeStamp = ""
	Output.openPrice = 0.0
	Output.highPrice = -1.0
	Output.lowPrice = 1000000.0
	Output.closePrice = 0.0
	Output.checkMe = false

	return Output
}

func ConvertStrToString(Output OutputStruct) []string {
	var totalStr []string

	totalStr = append(totalStr, Output.Company)
	totalStr = append(totalStr, Output.TimeStamp)
	totalStr = append(totalStr, fmt.Sprintf("%v", Output.openPrice))
	totalStr = append(totalStr, fmt.Sprintf("%v", Output.highPrice))
	totalStr = append(totalStr, fmt.Sprintf("%v", Output.lowPrice))
	totalStr = append(totalStr, fmt.Sprintf("%v", Output.closePrice))

	return totalStr
}

func Filer(path string, lhs int, rhs int) []MegaStruct {
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
		layout := "2006-01-02 15:04:05"
		//str := "2014-11-12T11:45:26.371Z"

		t, _ := time.Parse(layout, records[3])

		if t.Hour() >= lhs && t.Hour() <= rhs {
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

//func MainParser(total1 []MegaStruct, interval int) {
//	fileName := "candles_" + strconv.Itoa(interval) + "min.csv"
//	MakeCsvFromScratch(fileName, total1, interval) // GENERATE Csv files
//	fmt.Println(fileName, " done! Check this in folder.\n")

//}

func MakeCsvFromScratch(total []MegaStruct, interval int) {
	fileName := "candles_" + strconv.Itoa(interval) + "min.csv"
	//MakeCsvFromScratch(fileName, total1, interval) // GENERATE Csv files
	//defer fmt.Println(fileName, " done! Check this in folder.\n")
	template, _ := os.Create(fileName)
	writer := csv.NewWriter(template)
	defer fmt.Println(fileName, " almost done! Check this folder.\n")

	//MakeCsvFromScratch(writer)
	defer writer.Flush()

	startTime := total[0].TimeStamp
	var aaplC, sberC, amznC OutputStruct
	//for aapl,sber,amzn default
	aaplC, sberC, amznC = DefaultOutput(aaplC), DefaultOutput(sberC), DefaultOutput(amznC)
	isTime := true
	temp := interval

	if interval <= 60 {

		for i := range total {
			if isTime == false {
				startTime = total[i].TimeStamp
				isTime = true
			} //TODO make nice code

			//aapl part
			if total[i].Company == "AAPL" {
				if aaplC.checkMe == false {
					aaplC.Company = "AAPL"
					aaplC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					aaplC.checkMe = true
					aaplC.openPrice = total[i].PriceOpen

				}
				if aaplC.lowPrice > total[i].PriceOpen {
					aaplC.lowPrice = total[i].PriceOpen
				}

				if aaplC.highPrice < total[i].PriceOpen {
					aaplC.highPrice = total[i].PriceOpen
				}

			}
			//sber part
			if total[i].Company == "SBER" {
				if sberC.checkMe == false {
					sberC.Company = "SBER"
					sberC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					sberC.checkMe = true
					sberC.openPrice = total[i].PriceOpen

				}
				if sberC.lowPrice > total[i].PriceOpen {
					sberC.lowPrice = total[i].PriceOpen
				}

				if sberC.highPrice < total[i].PriceOpen {
					sberC.highPrice = total[i].PriceOpen
				}

			}
			//amzn part
			if total[i].Company == "AMZN" {
				if amznC.checkMe == false {
					amznC.Company = "AMZN"
					amznC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					amznC.checkMe = true
					amznC.openPrice = total[i].PriceOpen

				}
				if amznC.lowPrice > total[i].PriceOpen {
					amznC.lowPrice = total[i].PriceOpen
				}

				if amznC.highPrice < total[i].PriceOpen {
					amznC.highPrice = total[i].PriceOpen
				}

			}

			//writer
			if i == 0 {
				break
			}
			if (total[i].TimeStamp.Minute() == 0 && total[i-1].TimeStamp.Minute() != 0) || total[i].TimeStamp.Minute() >= interval || total[i].TimeStamp.Hour() > total[i-1].TimeStamp.Hour() {
				if aaplC.closePrice*aaplC.openPrice > 0 {
					obj := ConvertStrToString(aaplC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}
				if sberC.closePrice*sberC.openPrice > 0 {
					obj := ConvertStrToString(sberC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}
				if amznC.closePrice*amznC.openPrice > 0 {
					obj := ConvertStrToString(amznC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}
				interval += temp
				if interval > 60 {
					interval = temp
				}

				aaplC, sberC, amznC = DefaultOutput(aaplC), DefaultOutput(sberC), DefaultOutput(amznC)
				isTime = false

			} else {
				if total[i].Company == "AAPL" {
					aaplC.closePrice = total[i].PriceOpen
				}
				if total[i].Company == "SBER" {
					sberC.closePrice = total[i].PriceOpen
				}
				if total[i].Company == "AMZN" {
					amznC.closePrice = total[i].PriceOpen
				}
			}
		}

	} else {
		interval = interval / 60
		check := total[0]

		for i := range total {
			if isTime == false {
				startTime = total[i].TimeStamp
				isTime = true
			} //TODO make nice code

			//aapl part
			if total[i].Company == "AAPL" {
				if aaplC.checkMe == false {
					aaplC.Company = "AAPL"
					aaplC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					aaplC.checkMe = true
					aaplC.openPrice = total[i].PriceOpen

				}
				if aaplC.lowPrice > total[i].PriceOpen {
					aaplC.lowPrice = total[i].PriceOpen
				}

				if aaplC.highPrice < total[i].PriceOpen {
					aaplC.highPrice = total[i].PriceOpen
				}

			}
			//sber part
			if total[i].Company == "SBER" {
				if sberC.checkMe == false {
					sberC.Company = "SBER"
					sberC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					sberC.checkMe = true
					sberC.openPrice = total[i].PriceOpen

				}
				if sberC.lowPrice > total[i].PriceOpen {
					sberC.lowPrice = total[i].PriceOpen
				}

				if sberC.highPrice < total[i].PriceOpen {
					sberC.highPrice = total[i].PriceOpen
				}

			}
			//amzn part
			if total[i].Company == "AMZN" {
				if amznC.checkMe == false {
					amznC.Company = "AMZN"
					amznC.TimeStamp = startTime.Add(-time.Second).Format(time.RFC3339)
					amznC.checkMe = true
					amznC.openPrice = total[i].PriceOpen

				}
				if amznC.lowPrice > total[i].PriceOpen {
					amznC.lowPrice = total[i].PriceOpen
				}

				if amznC.highPrice < total[i].PriceOpen {
					amznC.highPrice = total[i].PriceOpen
				}

			}

			//writer
			if i == 0 {
				continue
			}
			if total[i].TimeStamp.Hour()-check.TimeStamp.Hour() == interval || total[i].TimeStamp.Day() > total[i-1].TimeStamp.Day() {

				check = total[i]
				if aaplC.closePrice*aaplC.openPrice > 0 {
					obj := ConvertStrToString(aaplC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}
				if sberC.closePrice*sberC.openPrice > 0 {
					obj := ConvertStrToString(sberC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}
				if amznC.closePrice*amznC.openPrice > 0 {
					obj := ConvertStrToString(amznC)
					writer.Write(obj)
					//aaplC = DefaultOutput(aaplC)
				}

				aaplC, sberC, amznC = DefaultOutput(aaplC), DefaultOutput(sberC), DefaultOutput(amznC)
				isTime = false

			} else {
				if total[i].Company == "AAPL" {
					aaplC.closePrice = total[i].PriceOpen
				}
				if total[i].Company == "SBER" {
					sberC.closePrice = total[i].PriceOpen
				}
				if total[i].Company == "AMZN" {
					amznC.closePrice = total[i].PriceOpen
				}
			}
		}

	}
}

func main() {

	MakeCsvFromScratch(Filer("trades.csv", 7, 23), 5)
	MakeCsvFromScratch(Filer("trades.csv", 7, 23), 30)
	MakeCsvFromScratch(Filer("trades.csv", 7, 23), 240)

}
