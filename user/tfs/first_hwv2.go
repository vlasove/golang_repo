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
	Company string //default:""
	TimeStamp string //default: ""
	openPrice float64 //default:0
	highPrice float64 //default:0
	lowPrice float64 //default:0
	closePrice float64 //default:0
	checkMe bool //default: false
}

func DefaultOutput(Output OutputStruct) OutputStruct{
	Output.Company = ""
	Output.TimeStamp = ""
	Output.openPrice = 0
	Output.highPrice = -1.0
	Output.lowPrice = 100000000.0
	Output.closePrice = 0
	Output.checkMe = false

	return Output
}

func ConvertStrToString(Output OutputStruct) []string {
	var totalStr []string

	totalStr = append(totalStr, Output.Company)
	totalStr = append(totalStr, Output.TimeStamp)
	totalStr = append(totalStr, fmt.Sprintf("%f",Output.openPrice))
	totalStr = append(totalStr, fmt.Sprintf("%f",Output.highPrice))
	totalStr = append(totalStr, fmt.Sprintf("%f",Output.lowPrice))
	totalStr = append(totalStr, fmt.Sprintf("%f",Output.closePrice))

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
		layout := "2012-12-12 12:12:12"
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

func MainParser(total []MegaStruct, interval int)  {
	fileName := "candles_" + strconv.Itoa(interval) + "min.csv"
	MakeCsvFromScratch(fileName, total, interval) // GENERATE Csv files
	fmt.Println(fileName," done! Check this in folder.\n")
	//template, _ := os.Create(fileName)
	//writer := csv.NewWriter(template)

	//MakeCsvFromScratch(writer)
	//defer writer.Flush()
	//startTime :=total[0].TimeStamp
	//var aaplC,sberC,amznC OutputStruct
	//isTime :=false	

}

func MakeCsvFromScratch(fileName string, total []MegaStruct, interval int){
	template, _ := os.Create(fileName)
	writer := csv.NewWriter(template)

	//MakeCsvFromScratch(writer)
	defer writer.Flush()
	startTime :=total[0].TimeStamp
	var aaplC,sberC,amznC OutputStruct
	isTime :=true	

	if interval < 60{

		for i,_ := range total {
			if isTime == false{
				startTime = total[i].TimeStamp
				isTime = true
			}//TODO make nice code
		}
		
	}else{

	}
}

func main() {
	

}
