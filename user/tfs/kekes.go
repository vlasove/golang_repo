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

	// Отсекаем 'm' для преобразования string в int
	inter = strings.Trim(inter, "m")

	// Из string в int для работы с разными интервалами
	intInterval, _ := strconv.Atoi(inter)
	gap := intInterval

	// Имя файла для записи
	var nameWriter string

	// Имя файла в зависимости от интервала
	switch {
	case intInterval == 5:
		nameWriter = "candles_5min.csv"
	case intInterval == 30:
		nameWriter = "candles_30min.csv"
	case intInterval == 240:
		nameWriter = "candles_240min.csv"
	default:
		fmt.Println("error_naming_out_csv")
	}

	// Создаем файл записи
	file, _ := os.Create(nameWriter)

	// Writer для записи в csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Используем функцию readFile, получаем все элементы из csv, фильрованные по времени [07:00 - 00:00)
	data := readFile(nameOfFile)

	// Для вывода в csv
	var apple, sberbank, amazon Out

	// Проименование
	apple.Name = "AAPL"
	sberbank.Name = "SBER"
	amazon.Name = "AMZN"

	// Т.к. далее начинается проход по слайсу со 2 элемента, задаем время от первого элемента тут
	outTime := data[0].TimeStamp

	// Проверка, задано ли время для out или еще нет
	var isOutTimeThere = true

	switch {
	// Если заданное время <= 60м
	case intInterval <= 60:

		// Проходим от первого до последнего элемента слайса
		for num := range data {

			// Если время еще не задано - задать
			if isOutTimeThere == false {
				outTime = data[num].TimeStamp
				isOutTimeThere = true
			}

			// Пропускаем 1 элемент, т.к. далее будет проверка num-1 элемента (чтобы не было выхода за границы слайса)
			if num == 0 {
				continue
			}

			switch {
			case data[num].Name == "AAPL":

				// Первое вхождение AAPL, для Open переменной (на интервале), для начального значения LOW и времени
				if apple.isOpenThere == false {
					apple.open = data[num].Price
					apple.low = data[num].Price
					apple.TimeStamp = resetSeconds(outTime)
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

				// Первое вхождение SBER, для Open переменной (на интервале), для начального значения LOW и времени
				if sberbank.isOpenThere == false {
					sberbank.open = data[num].Price
					sberbank.low = data[num].Price
					sberbank.TimeStamp = resetSeconds(outTime)
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

				// Первое вхождение AMZN, для Open переменной (на интервале), для начального значения LOW и времени
				if amazon.isOpenThere == false {
					amazon.open = data[num].Price
					amazon.low = data[num].Price
					amazon.TimeStamp = resetSeconds(outTime)
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

				//////////////////////////////////////////////////////////
				// Работа с готовыми данными (здесь окончание интервала)//
				//////////////////////////////////////////////////////////

				// Запись (если не 0 значения) в файл
				if apple.open != 0 && apple.close != 0 {
					writer.Write(OutToString(apple))
					//fmt.Println(apple)
				}

				if sberbank.open != 0 && sberbank.close != 0 {
					writer.Write(OutToString(sberbank))
					//fmt.Println(sberbank)
				}

				if amazon.open != 0 && amazon.close != 0 {
					writer.Write(OutToString(amazon))
					//fmt.Println(amazon)
				}

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

				// Обнуление, для изменения времени в out.tTimeStamp
				isOutTimeThere = false

				//fmt.Println("NEW ITTERATION-----------------------------------------------------------------------------------------------------------")
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
	// Если заданное время > 60м
	case intInterval > 60:

		// Делим минуты на 60 для получения кол-ва часов
		intInterval = intInterval / 60

		// Для проверки времени по часам
		checkHour := data[0]

		for num := range data {

			// Перврое вхождение времени каждого интервала
			if isOutTimeThere == false {
				outTime = data[num].TimeStamp
				isOutTimeThere = true
			}

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
					apple.TimeStamp = resetSeconds(outTime)
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
					sberbank.TimeStamp = resetSeconds(outTime)
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
					amazon.TimeStamp = resetSeconds(outTime)
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

				// Запись (если не 0 значения) в файл
				if apple.open != 0 && apple.close != 0 {
					writer.Write(OutToString(apple))
					//fmt.Println(apple)
				}

				if sberbank.open != 0 && sberbank.close != 0 {
					writer.Write(OutToString(sberbank))
					//fmt.Println(sberbank)
				}

				if amazon.open != 0 && amazon.close != 0 {
					writer.Write(OutToString(amazon))
					//fmt.Println(amazon)
				}

				// Сброс данных
				apple = resetData(apple)
				sberbank = resetData(sberbank)
				amazon = resetData(amazon)

				// Обнуление, для изменения времени в out.tTimeStamp
				isOutTimeThere = false

				//fmt.Println("NEW ITTERATION-----------------------------------------------------------------------------------------------------------")
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

// Сброс секунд
func resetSeconds(mytime time.Time) string {

	for {

		// Обнуляем секунды
		if mytime.Second() == 0 {
			break
		} else {
			mytime = mytime.Add(time.Second * (-1))
		}

	}

	// Убираем миллисекунды с помощью формата
	mTime := mytime.Format(time.RFC3339)

	return mTime
}

// Делает из Out слайс строк (условие записи в csv файл)
func OutToString(out Out) []string {

	var outString []string

	outString = append(outString, out.Name)
	outString = append(outString, out.TimeStamp)

	floatStr := fmt.Sprintf("%g", out.open)
	outString = append(outString, floatStr)
	floatStr = fmt.Sprintf("%g", out.high)
	outString = append(outString, floatStr)
	floatStr = fmt.Sprintf("%g", out.low)
	outString = append(outString, floatStr)
	floatStr = fmt.Sprintf("%g", out.close)
	outString = append(outString, floatStr)

	return outString

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
	TimeStamp              string
	open, high, low, close float64
	isOpenThere            bool
}

func main() {

	CandleData("trades.csv", "5m")
	CandleData("trades.csv", "30m")
	CandleData("trades.csv", "240m")

}
