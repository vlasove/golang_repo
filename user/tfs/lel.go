package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Для парсинга даты из строки
const dateMask = "2006-01-02 15:04:05"

func readFile(nameOfFile string) []Data {

	// Открытие файла
	csvFile, _ := os.Open(nameOfFile)
	reader := csv.NewReader(csvFile)

	// Экземпляр структуры Data
	var out []Data

	// Пробегаемся по всему файлу
	for {

		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
			fmt.Println("ERROR")
		}

		// Временная для цены
		var tPrice float64
		tPrice, _ = strconv.ParseFloat(line[1], 64)

		// Временная для кол-ва
		var tAmount int64
		tAmount, _ = strconv.ParseInt(line[2], 10, 64)

		// Временная для времени :)
		var tTimeStamp time.Time
		tTimeStamp, error = time.Parse(dateMask, line[3])
		if error != nil {
			fmt.Println(error)
		}

		// Фильтр по времени, запись в слайс out
		if tTimeStamp.Hour() >= 7 && tTimeStamp.Hour() <= 23 {

			// Присвоение очередной строки в слайс out структуры Data
			out = append(out, Data{
				Name:      line[0],
				Price:     tPrice,
				Amount:    tAmount,
				TimeStamp: tTimeStamp,
			})
		}
	}
	return out
}

func CandleData(nameOfFile string, inter string) {

	// Для добавления к outTime
	//dataInterval, _ := time.ParseDuration(inter)

	// Отсекаем 'm' для преобразования string в int
	inter = strings.Trim(inter, "m")

	// Из string в int для работы с разными интервалами
	intInterval, _ := strconv.Atoi(inter)
	gap := intInterval

	// Используем функцию readFile, получаем все элементы из csv, фильрованные по времени [07:00 - 00:00)
	data := readFile(nameOfFile)

	// Для вывода в csv
	var apple, sberbank, amazon Out

	// Проименование
	apple.Name = "AAPL"
	sberbank.Name = "SBER"
	amazon.Name = "AMZN"

	// Работа со временем в вывод csv
	//outTime := time.Date(0001, 01, 01, 07, 00, 00, 00, time.UTC)
	//outTime = outTime.AddDate(data[0].TimeStamp.Year()-1, 0, data[0].TimeStamp.YearDay()-1)

	switch {
	// Если заданное время <= 60м
	case intInterval <= 60:
		// Проходим от первого до последнего элемента слайса
		for num := range data {
			//for num := 0; num < 15000; num++ {

			// Добавление +7 часов, если время достигло 0
			//if outTime.Hour() == 0 {
			//	outTime = outTime.Add(time.Hour * 7)
			//}

			// Пропускаем 1 элемент, т.к. далее будет проверка num-1 элемента (чтобы не было выхода за границы слайса)
			if num == 0 {
				continue
			}

			switch {
			case data[num].Name == "AAPL":

				// Первое вхождение AAPL, для Open переменной (на интервале)
				if apple.isOpenThere == false {
					apple.open = data[num].Price
					apple.low = data[num].Price
					apple.isOpenThere = true
				}

				// Максимальная цена
				if apple.high < data[num].Price {
					apple.high = data[num].Price
				}

				// Минимальная цена
				if apple.low > data[num].Price {
					apple.low = data[num].Price
				}

			case data[num].Name == "SBER":

				// Первое вхождение SBER, для Open переменной (на интервале)
				if sberbank.isOpenThere == false {
					sberbank.open = data[num].Price
					sberbank.low = data[num].Price
					sberbank.isOpenThere = true
				}

				// Максимальная цена
				if sberbank.high < data[num].Price {
					sberbank.high = data[num].Price
				}

				// Минимальная цена
				if sberbank.low > data[num].Price {
					sberbank.low = data[num].Price
				}

			case data[num].Name == "AMZN":

				// Первое вхождение AMZN, для Open переменной (на интервале)
				if amazon.isOpenThere == false {
					amazon.open = data[num].Price
					amazon.low = data[num].Price
					amazon.isOpenThere = true
				}

				// Максимальная цена
				if amazon.high < data[num].Price {
					amazon.high = data[num].Price
				}

				// Минимальная цена
				if amazon.low > data[num].Price {
					amazon.low = data[num].Price
				}
			}

			// Деление слайса по заданным интервалам
			// Вторая часть условия (после или) для корректного перехода от 55 минут к 00 минутам следующего часа
			if data[num].TimeStamp.Minute() >= intInterval || (data[num].TimeStamp.Minute() == 0 && data[num-1].TimeStamp.Minute() != 0) || data[num].TimeStamp.Hour() > data[num-1].TimeStamp.Hour() {

				// Добавление времени
				//apple.TimeStamp = outTime
				//sberbank.TimeStamp = outTime
				//amazon.TimeStamp = outTime

				//////////////////////////////////////////////////////////
				// Работа с готовыми данными (здесь окончание интервала)//
				//////////////////////////////////////////////////////////
				fmt.Println(apple)
				fmt.Println(sberbank)
				fmt.Println(amazon)

				// Прибавка интервала
				intInterval += gap

				// Сброс интервала при переходе на след час
				if intInterval > 60 {
					intInterval = gap
				}

				// Сброс данных
				apple = resetData(apple)
				sberbank = resetData(sberbank)
				amazon = resetData(amazon)

				// Добавление интервала ко времени
				//outTime = outTime.Add(dataInterval)

				fmt.Println("NEW ITTERATION-----------------------------------------------------------------------------------------------------------")
				fmt.Println(intInterval)

			} else {
				switch {
				case data[num].Name == "AAPL":
					// Последнее вхождение AAPL, для Close переменной (на интервале)
					apple.close = data[num].Price

				case data[num].Name == "SBER":
					// Последнее вхождение SBER, для Close переменной (на интервале)
					sberbank.close = data[num].Price

				case data[num].Name == "AMZN":
					// Последнее вхождение AMZN, для Close переменной (на интервале)
					amazon.close = data[num].Price
				}
			}
		}
	// Если заданное время > 60м
	case intInterval > 60:

		// Делим минуты на 60 для получения кол-ва часов
		intInterval = intInterval / 60

		checkHour := data[0]

		for num := range data {
			//for num := 0; num < 20000; num++ {

			// Пропускаем 1 элемент, т.к. далее будет проверка num-1 элемента (чтобы не было выхода за границы слайса)
			if num == 0 {
				continue
			}

			switch {
			case data[num].Name == "AAPL":

				// Первое вхождение AAPL, для Open переменной (на интервале)
				if apple.isOpenThere == false {
					apple.open = data[num].Price
					apple.low = data[num].Price
					apple.isOpenThere = true
				}

				// Максимальная цена
				if apple.high < data[num].Price {
					apple.high = data[num].Price
				}

				// Минимальная цена
				if apple.low > data[num].Price {
					apple.low = data[num].Price
				}

			case data[num].Name == "SBER":

				// Первое вхождение SBER, для Open переменной (на интервале)
				if sberbank.isOpenThere == false {
					sberbank.open = data[num].Price
					sberbank.low = data[num].Price
					sberbank.isOpenThere = true
				}

				// Максимальная цена
				if sberbank.high < data[num].Price {
					sberbank.high = data[num].Price
				}

				// Минимальная цена
				if sberbank.low > data[num].Price {
					sberbank.low = data[num].Price
				}

			case data[num].Name == "AMZN":

				// Первое вхождение AMZN, для Open переменной (на интервале)
				if amazon.isOpenThere == false {
					amazon.open = data[num].Price
					amazon.low = data[num].Price
					amazon.isOpenThere = true
				}

				// Максимальная цена
				if amazon.high < data[num].Price {
					amazon.high = data[num].Price
				}

				// Минимальная цена
				if amazon.low > data[num].Price {
					amazon.low = data[num].Price
				}
			}

			// Деление слайса по заданным интервалам
			// Если есть разница между часами равна интервалу или изменился день
			if data[num].TimeStamp.Hour()-checkHour.TimeStamp.Hour() == intInterval || data[num].TimeStamp.Day() > data[num-1].TimeStamp.Day() {

				checkHour = data[num]

				//////////////////////////////////////////////////////////
				// Работа с готовыми данными (здесь окончание интервала)//
				//////////////////////////////////////////////////////////
				fmt.Println(apple)
				fmt.Println(sberbank)
				fmt.Println(amazon)

				// Сброс данных
				apple = resetData(apple)
				sberbank = resetData(sberbank)
				amazon = resetData(amazon)

				fmt.Println("NEW ITTERATION-----------------------------------------------------------------------------------------------------------")
				//fmt.Println(intInterval)

			} else {

				switch {
				case data[num].Name == "AAPL":
					// Последнее вхождение AAPL, для Close переменной (на интервале)
					apple.close = data[num].Price

				case data[num].Name == "SBER":
					// Последнее вхождение SBER, для Close переменной (на интервале)
					sberbank.close = data[num].Price

				case data[num].Name == "AMZN":
					// Последнее вхождение AMZN, для Close переменной (на интервале)
					amazon.close = data[num].Price
				}
			}
		}
	}
}

// Сброс данных apple, sberbank и amazon
func resetData(name Out) Out {
	name.isOpenThere = false
	name.open = 0
	name.high = 0
	name.low = 0
	name.close = 0
	return name
}

// Структура для работы с информацией из csv
type Data struct {
	Name      string
	Price     float64
	Amount    int64
	TimeStamp time.Time
}

// Структура для записи в csv
type Out struct {
	Name                   string
	TimeStamp              time.Time
	open, high, low, close float64
	isOpenThere            bool
}

func main() {

	CandleData("trades.csv", "240m")
}
