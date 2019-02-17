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

type OHLC struct {
	O float64
	H float64
	L float64
	C float64
}
type Output struct {
	Company string
	D       time.Time
	O       float64
	H       float64
	L       float64
	C       float64
}

const layout = "2006-01-02T15:04:05Z"

func writeCandlesToFile(interval int) string {
	tradeStart, _ := time.Parse("2006-01-02 15:04:05.999999", "2019-01-30 07:00:00.000000")
	file, _ := os.Open("trades.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	fileName := fmt.Sprintf("trades_%vmin.csv", interval)
	output, _ := os.Create(fileName)
	defer output.Close()
	writer := csv.NewWriter(output)
	defer writer.Flush()
	tickerMap := make(map[string]*OHLC)
	for curDate := tradeStart; ; {
		line, err := reader.Read()
		if err == io.EOF {
			for k, v := range tickerMap {
				result := []string{k, curDate.Format(time.RFC3339),
					strconv.FormatFloat(v.O, 'f', 2, 64),
					strconv.FormatFloat(v.H, 'f', 2, 64),
					strconv.FormatFloat(v.L, 'f', 2, 64),
					strconv.FormatFloat(v.C, 'f', 2, 64),
				}
				writer.Write(result)
				delete(tickerMap, k)
			}
			break
		}
		ticker := line[0]
		price, _ := strconv.ParseFloat(line[1], 64)
		dateTime, _ := time.Parse("2006-01-02 15:04:05.999999", line[3])
		//setting up the interval ending time
		nextDate := curDate.Add(time.Duration(interval) * time.Minute)
		//skipping all the unnecessary transactions
		if dateTime.Hour() >= 0 && dateTime.Hour() < 7 {
			continue
		}
		//if dateTime of transaction exceeds the interval ending time, we write the ohlc result to the file,
		//delete all the contents of our map and set curDate to reflect the relevant time interval (some intervals might
		// be skipped)
		if nextDate.Sub(dateTime) <= 0 {
			for k, v := range tickerMap {
				result := []string{k, curDate.Format(time.RFC3339),
					strconv.FormatFloat(v.O, 'f', 2, 64),
					strconv.FormatFloat(v.H, 'f', 2, 64),
					strconv.FormatFloat(v.L, 'f', 2, 64),
					strconv.FormatFloat(v.C, 'f', 2, 64),
				}
				writer.Write(result)
				delete(tickerMap, k)
			}
			//this is the part where we find out the relevant date
			curDate = nextDate.Add(time.Minute * time.Duration(interval) * time.Duration(int(dateTime.Sub(nextDate).Minutes())/interval))
		}
		//checking whether the ticker is present in the map
		ohlc, ok := tickerMap[ticker]
		if !ok {
			//if it isn't we create it and assign a new ohlc struct to it
			tickerMap[ticker] = &OHLC{price, price, price, price}
			continue
		}
		//if it is we update ohlc if necessary
		if price > ohlc.H {
			ohlc.H = price
		} else if price < ohlc.L {
			ohlc.L = price
		}
		//we also update the closing price each time (we don't know if this is going to be the last transaction)
		ohlc.C = price

	}
	return fileName
}

func fRomCSVtoInput(path string) []Output {
	var arr []Output
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
		t, _ := time.Parse(layout, records[1])
		O, _ := strconv.ParseFloat(records[2], 64)
		H, _ := strconv.ParseFloat(records[3], 64)
		L, _ := strconv.ParseFloat(records[4], 64)
		C, _ := strconv.ParseFloat(records[5], 64)
		//value, _ := strconv.Atoi(records[2])

		//str := "2014-11-12T11:45:26.371Z"

		arr = append(arr, Output{
			Company: records[0],
			D:       t,
			O:       O,
			H:       H,
			L:       L,
			C:       C,
		})

	}
	return arr

}

func less(lhs Output, rhs Output) bool {
	if lhs.D.Day()*24*60+lhs.D.Hour()*60+lhs.D.Minute() < rhs.D.Day()*24*60+rhs.D.Hour()*60+rhs.D.Minute() {
		return true
	} else if lhs.D.Day()*24*60+lhs.D.Hour()*60+lhs.D.Minute() == rhs.D.Day()*24*60+rhs.D.Hour()*60+rhs.D.Minute() && lhs.O < rhs.O {
		return true
	} else {
		return false
	}
}

func sortedTotal(csv_path string) {
	obj := fRomCSVtoInput(csv_path)
	//for i, _ := range obj {
	//	fmt.Println(obj[i])
	//}

	sort.Slice(obj[:], func(i, j int) bool {
		//return (obj[i].D.Minute() < obj[j].D.Minute() && obj[i].D.Hour() < obj[j].D.Hour() && obj[i].D.Day() < obj[j].D.Day() && obj[i].O < obj[j].O)

		return less(obj[i], obj[j])
	})

	output, _ := os.Create("sorted_" + csv_path)
	defer output.Close()
	writer := csv.NewWriter(output)
	defer writer.Flush()

	for i, _ := range obj {
		result := []string{obj[i].Company, obj[i].D.Format(time.RFC3339),
			strconv.FormatFloat(obj[i].O, 'f', 2, 64),
			strconv.FormatFloat(obj[i].H, 'f', 2, 64),
			strconv.FormatFloat(obj[i].L, 'f', 2, 64),
			strconv.FormatFloat(obj[i].C, 'f', 2, 64),
		}
		writer.Write(result)
	}

}

func main() {
	sortedTotal(writeCandlesToFile(5))
	sortedTotal(writeCandlesToFile(30))
	sortedTotal(writeCandlesToFile(240))
	//csv30 := writeCandlesToFile(30)
	//csv240 := writeCandlesToFile(240)

}
