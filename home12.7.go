//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//	"time"
//)
//
//func main() {
//	ad := ""
//	scanner := bufio.NewReader(os.Stdin)
//	for i := 1; ; i++{
//		fmt.Println("Введите сообщение: ")
//	    inputRead, _ := scanner.ReadString('\n')
//		if strings.Contains(inputRead, "выход"){
//			break
//		}
//		dt := time.Now()
//		ad += fmt.Sprint(i," ",dt.Format("01-02-2006 15:04:05")," ",inputRead)
//	}
//	f, err := os.Create("data.txt")
//	if err != nil {
//		fmt.Println("Упс :( Не удалось зодать файл!", err)
//		return
//	}
//	defer f.Close()
//
//	f.WriteString(ad)
//}


//
//package main
//
//import (
//	"bytes"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main()  {
//	file, err := os.Create("home12.7.txt")
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	s := "a10 10 20b 20 30c30 30 dd"
//	space := " "
//	var b bytes.Buffer
//
//	for strings.Contains(s, space) {
//		spaceIndex := strings.Index(s, space)
//		words := s[:spaceIndex]
//		i, err := strconv.Atoi(words)
//		if err == nil {
//			fmt.Println(i)
//			b.WriteString(fmt.Sprint(i,"\n"))
//		}
//		s = s[spaceIndex+1:]
//		s = strings.Trim(s, space)
//		}
//	_, err = file.Write(b.Bytes())
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//
//}



//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	file, err := os.Create("file.txt")
//	if err := os.Chmod("file.txt", 0444); err != nil{
//		fmt.Println(err)
//	}
//	writer := bufio.NewWriter(file)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//
//	writer.WriteString("Hello")
//	writer.Flush()
//}


//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//	var b bytes.Buffer
//	var lap int = 4
//	b.WriteString(fmt.Sprintf("lap: %d \n", lap))
//	var engine int = 254
//	b.WriteString(fmt.Sprintf("engine: %d \n", engine))
//	var wheels int = 93
//	b.WriteString(fmt.Sprintf("wheels: %d \n", wheels))
//	var steeringWheel int = 49
//	b.WriteString(fmt.Sprintf("steeringWheel: %d \n", steeringWheel))
//	var wind int = 21
//	b.WriteString(fmt.Sprintf("wind: -%d \n", wind))
//	var rain int = 17
//	b.WriteString(fmt.Sprintf("rain: -%d \n", rain))
//	var speed int = engine + wheels + steeringWheel - wind - rain
//	b.WriteString(fmt.Sprintln("Итоговоя скорость гонщика: ", speed,"км/ч"))
//
//	fileName := "racer.txt"
//	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println(string(resultBytes))
//}





//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//
//	var b bytes.Buffer
//
//	var product int = 6400
//	b.WriteString(fmt.Sprintf("product: %d \n", product))
//	var delivery int = 350
//	b.WriteString(fmt.Sprintf("delivery: %d \n", delivery))
//	var discount int = 700
//	b.WriteString(fmt.Sprintf("ldiscount: %d \n", discount))
//	var result int = product + delivery - discount
//	b.WriteString(fmt.Sprintf("result: %d \n", result))
//
//		fileName := "discount.txt"
//		if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//			log.Fatal(err)
//		}
//		file, err := os.Open(fileName)
//		if err != nil{
//			panic(err)
//		}
//		defer file.Close()
//		resultBytes, err := ioutil.ReadAll(file)
//		if err != nil{
//			panic(err)
//		}
//		fmt.Println(string(resultBytes))
//
//
//}



//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//
//	var b bytes.Buffer
//
//	var change int = 480
//	b.WriteString(fmt.Sprintf("длительность смены: %d мин. \n", change))
//	var clientOrder int = 2
//	b.WriteString(fmt.Sprintf("клиент делает заказ: %d мин. \n", clientOrder))
//	var cashier int = 4
//	b.WriteString(fmt.Sprintf("кассир собирает заказ: %d мин. \n", cashier))
//	var result int = change / (clientOrder + cashier)
//	b.WriteString(fmt.Sprintf("За смену длиной %d минут кассир успеет обслужить %d, клиентов",change, result))
//
//	fileName := "cashier.txt"
//	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//
//	fmt.Println(string(resultBytes))
//}



//package main
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//	var b bytes.Buffer
//	var receipt int = 4000000
//	b.WriteString(fmt.Sprintf("Сумма, указанная в квитанции: %d руб. \n", receipt))
//	var entrances int = 10
//	b.WriteString(fmt.Sprintf("Подъездов в доме: %d \n", entrances))
//	var apartments int = 40
//	b.WriteString(fmt.Sprintf("Квартир в каждом подъезде: %d \n", apartments))
//	var result int = receipt / (entrances * apartments)
//	b.WriteString(fmt.Sprintf("Каждая квартира должна заплатить по: %d руб. \n", result))
//
//	fileName := "receipt.txt"
//	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//
//	fmt.Println(string(resultBytes))
//}



//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//
//	var b bytes.Buffer
//	var planet string
//	fmt.Println("Введите название планеты")
//	_, _ = fmt.Scan(&planet)
//	b.WriteString(fmt.Sprintf("Введите название планеты: %v \n", planet))
//
//	fmt.Println("Введите название звездной системы")
//	var star string
//	_, _ = fmt.Scan(&star)
//	b.WriteString(fmt.Sprintf("Введите название планеты: %v \n", star))
//
//	fmt.Println("Введите имя рейнджера")
//	var ranger string
//	_, _ = fmt.Scan(&ranger)
//	b.WriteString(fmt.Sprintf("Введите название планеты: %v \n", ranger))
//
//	fmt.Println("Введите количество вознаграждения")
//	var money int
//	_, _ = fmt.Scan(&money)
//	b.WriteString(fmt.Sprintf("Введите название планеты: %d  \n", money))
//	b.WriteString(fmt.Sprintln("========================="))
//	b.WriteString(fmt.Sprintf("Как вам %v известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты %v системы %s .", ranger, planet, star))
//	b.WriteString(fmt.Sprintln("\n"))
//	b.WriteString(fmt.Sprintf("Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты %v впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!", planet))
//	b.WriteString(fmt.Sprintln("\n"))
//	b.WriteString(fmt.Sprintf("%v вы должны прибыть на планету %v системы %v и помочь выполнить план призыва. Мы готовы выплатить вам премию в %d кредитов за эту маленькую услугу.",ranger, planet, star, money))
//
//	fileName := "planet.txt"
//	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//
//	fmt.Println(string(resultBytes))
//}





//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//
//
//	var b bytes.Buffer
//	stantion1 := "Будапештская"
//	b.WriteString(fmt.Sprintln(stantion1))
//	stantion2 := "ул. Славы 4"
//	b.WriteString(fmt.Sprintln(stantion2))
//	stantion3 := "Вечный огонь"
//	b.WriteString(fmt.Sprintln(stantion3))
//	stantion4 := "Аэропорт"
//	b.WriteString(fmt.Sprintln(stantion4))
//	b.WriteString(fmt.Sprint("\n"))
//	passagir := 0
//
//	ticketAmount := 20
//
//
//	var totalSum int
//	b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion1, passagir))
//
//	var passagirBus int
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	_, _ = fmt.Scan(&passagir)
//	b.WriteString(fmt.Sprintf("Сколько пассажиров зашло на остановке?: %d \n", passagir))
//
//	passagirBus += passagir
//	totalSum += (passagir * ticketAmount)
//
//	b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion2, passagirBus))
//	b.WriteString(fmt.Sprintln("----------Едем-------------"))
//
//	b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion2, passagirBus))
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	_, _ = fmt.Scan(&passagir)
//	b.WriteString(fmt.Sprintf("Сколько пассажиров вышло на остановке? %d \n", passagir))
//
//	passagirBus -= passagir
//
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	_, _ = fmt.Scan(&passagir)
//	b.WriteString(fmt.Sprintf("Сколько пассажиров зашло на остановке? %d \n", passagir))
//
//	passagirBus += passagir
//	totalSum += passagir * ticketAmount
//
//	b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion3, passagirBus))
//	b.WriteString(fmt.Sprintln("----------Едем-------------"))
//
//	b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion3, passagirBus))
//	fmt.Println("Сколько пассажиров вышло на остановке?")
//	_, _ = fmt.Scan(&passagir)
//	b.WriteString(fmt.Sprintf("Сколько пассажиров вышло на остановке? %d \n", passagir))
//
//	passagirBus -= passagir
//	fmt.Println("Сколько пассажиров зашло на остановке?")
//	_, _ = fmt.Scan(&passagir)
//	b.WriteString(fmt.Sprintf("Сколько пассажиров зашло на остановке? %d \n", passagir))
//
//	passagirBus += passagir
//	totalSum += passagir * ticketAmount
//
//	b.WriteString(fmt.Sprintf("Отправляемся с остановки %s. В салоне пассажиров: %d\n", stantion4, passagirBus))
//	b.WriteString(fmt.Sprintln("----------Едем-------------"))
//
//	b.WriteString(fmt.Sprintf("Прибываем на остановку %s. В салоне пассажиров: %d\n", stantion4, passagirBus))
//	b.WriteString(fmt.Sprintf("На остановке вышло %d\n", passagirBus))
//
//	passagirBus -= passagirBus
//	driverSalary := totalSum / 4
//	taxes := totalSum / 5
//	carRepair := totalSum / 5
//	fuelConsumption := totalSum / 5
//	income := totalSum - (driverSalary + taxes + carRepair + fuelConsumption)
//
//	b.WriteString(fmt.Sprint("\n"))
//	b.WriteString(fmt.Sprintf("Всего заработанно %d руб.\n", totalSum))
//	b.WriteString(fmt.Sprintf("Зарплата водителя: %d руб.\n", driverSalary))
//	b.WriteString(fmt.Sprintf("Расходы на топливо: %d руб.\n", fuelConsumption))
//	b.WriteString(fmt.Sprintf("Налоги %d руб.\n", taxes))
//	b.WriteString(fmt.Sprintf("Расходы на ремонт машины: %d руб.\n", carRepair))
//	b.WriteString(fmt.Sprintf("Итого доход: %d руб.\n", income))
//
//	fileName := "minibus.txt"
//	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//
//	fmt.Println(string(resultBytes))
//	}




//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//func main() {
//
//	var byteB bytes.Buffer
//	a := 42
//	byteB.WriteString(fmt.Sprintf("%d \n", a))
//	b := 153
//	byteB.WriteString(fmt.Sprintf("%d \n", b))
//
//	a, b = b, a
//
//	byteB.WriteString(fmt.Sprintf("%d \n", a))
//	byteB.WriteString(fmt.Sprintf("%d \n", b))
//
//	fileName := "change.txt"
//	if err := ioutil.WriteFile(fileName, byteB.Bytes(), 0666); err != nil {
//		log.Fatal(err)
//	}
//	file, err := os.Open(fileName)
//	if err != nil{
//		panic(err)
//	}
//	defer file.Close()
//	resultBytes, err := ioutil.ReadAll(file)
//	if err != nil{
//		panic(err)
//	}
//
//	fmt.Println(string(resultBytes))
//
//}





package main

import(
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var b bytes.Buffer


	fmt.Println("Ведите размер саженцов при высадке в см:")
	var height int
	_, _ = fmt.Scan(&height)
	b.WriteString(fmt.Sprintf("размер саженцов при высадке в см: %d \n", height))

	fmt.Println("Введите скорость роста бамбука в см:")
	var growth int
	_, _ = fmt.Scan(&growth)
	b.WriteString(fmt.Sprintf("скорость роста бамбука в см: %d \n", growth))

	fmt.Println("Введите скорость поедания бамбука гусеницами в см:")
	var eat int
	_, _ = fmt.Scan(&eat)
	b.WriteString(fmt.Sprintf("скорость поедания бамбука гусеницами в см: %d \n", eat))

	fmt.Println("Введите колличество дней:")
	var day int
	_, _ = fmt.Scan(&day)
	b.WriteString(fmt.Sprintf("колличество дней: %d \n", day))

	fmt.Println("Введите колличество ночей:")
	var night int
	_, _ = fmt.Scan(&night)
	b.WriteString(fmt.Sprintf("колличество ночей: %d \n", night))

	fmt.Println("Введите целевую высоту бамбука в см:")
	var target int
	_, _ = fmt.Scan(&target)
	b.WriteString(fmt.Sprintf("целевая высота бамбука в см: %d \n", target))

	target = height + day * growth - eat * night


	b.WriteString(fmt.Sprintln(target))

		fileName := "growth.txt"
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