package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {


	var b bytes.Buffer
		stantion1 := "Будапештская"
		b.WriteString(fmt.Sprintf(stantion1))
		stantion2 := "ул. Славы 4"
	b.WriteString(fmt.Sprint(stantion2))
		stantion3 := "Вечный огонь"
	b.WriteString(fmt.Sprint(stantion3))
		stantion4 := "Аэропорт"
	b.WriteString(fmt.Sprint(stantion4))
		passagir := 0
	b.WriteString(fmt.Sprintf("%d", passagir))
		ticketAmount := 20
	b.WriteString(fmt.Sprintf("%D", ticketAmount))
		var totalSum int

		b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion1, passagir))
		var passagirBus int
		b.WriteString(fmt.Sprintln("Сколько пассажиров зашло на остановке?"))
		_, _ = fmt.Scan(&passagir)
		b.WriteString(fmt.Sprintf("Сколько пассажиров зашло на остановке?: %d \n", passagir))
		passagirBus += passagir
		totalSum += (passagir * ticketAmount)
		b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion2, passagirBus))
		b.WriteString(fmt.Sprintln("----------Едем-------------"))

		b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion2, passagirBus))
		b.WriteString(fmt.Sprintln("Сколько пассажиров вышло на остановке?"))
		_, _ = fmt.Scan(&passagir)
		passagirBus -= passagir
	b.WriteString(fmt.Sprintln("Сколько пассажиров зашло на остановке?"))
	_, _ = fmt.Scan(&passagir)
		passagirBus += passagir
		totalSum += passagir * ticketAmount
		b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion3, passagirBus))
		b.WriteString(fmt.Sprintln("----------Едем-------------"))

		b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion3, passagirBus))
		b.WriteString(fmt.Sprintln("Сколько пассажиров вышло на остановке?"))
	_, _ = fmt.Scan(&passagir)
		passagirBus -= passagir
		b.WriteString(fmt.Sprintln("Сколько пассажиров зашло на остановке?"))
	_, _ = fmt.Scan(&passagir)
		passagirBus += passagir
		totalSum += passagir * ticketAmount
		b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion4, passagirBus))
		b.WriteString(fmt.Sprintln("----------Едем-------------"))

		b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion4, passagirBus))
		b.WriteString(fmt.Sprintf("На остановке вышло %d\n", passagirBus))
		passagirBus -= passagirBus
		driverSalary := totalSum / 4
		taxes := totalSum / 5
		carRepair := totalSum / 5
		fuelConsumption := totalSum / 5
		income := totalSum - (driverSalary + taxes + carRepair + fuelConsumption)

		b.WriteString(fmt.Sprintf("Всего заработанно %d руб.\n", totalSum))
		b.WriteString(fmt.Sprintf("Зарплата водителя: %d руб.\n", driverSalary))
		b.WriteString(fmt.Sprintf("Расходы на топливо: %d руб.\n", fuelConsumption))
		b.WriteString(fmt.Sprintf("Налоги %d руб.\n", taxes))
		b.WriteString(fmt.Sprintf("Расходы на ремонт машины: %d руб.\n", carRepair))
		b.WriteString(fmt.Sprintf("Итого доход: %d руб.\n", income))

		fileName := "minibus.txt"
		if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}
		file, err := os.Open(fileName)
		if err != nil{
			panic(err)
		}
		defer file.Close()
		resultBytes, err := ioutil.ReadAll(file)
		if err != nil{
			panic(err)
		}

		fmt.Println(string(resultBytes))
	}