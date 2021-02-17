package main

import "fmt"

func main() {

	fmt.Println("Пользователь вводит месяц, программа должна вывести, на какое время года (зима, весна, лето, осень) этот месяц выпадает.")

	fmt.Print("Введите название месяца: ")
	var season string
	fmt.Scan(&season)

	switch season {
	case "декабрь", "январь", "февраль":
		fmt.Println(season, "является зимним месяцем")

	case "март", "апрель", "май":
		fmt.Println(season, "является весенним месяцем")

	case "июнь", "июль", "август":
		fmt.Println(season, "является летним месяцем")

	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println(season, "является осенним месяцем")

	default:
		fmt.Println("Такого месяца не существует")
	}
}

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Пользователь вводит будний день недели в сокращенной форме(пн, вт, ср, чт, пт)и получает развернутый список всех последующих рабочих дней, включая пятницу.  ")
//
//	fmt.Println("Введите день недели, что бы узнать какие рабочие дни остались: ")
//	var days string
//	fmt.Scan(&days)
//
//	switch days {
//
//	case "пн":
//		fmt.Println("Понедельник","Вторник","Среда","Четверг","Пятница")
//
//	case "вт":
//		fmt.Println("Вторик","Среда","Четверг","Пятница")
//
//	case "ср":
//		fmt.Println("Среда","Четверг","Пятница")
//
//	case "чт":
//		fmt.Println("Четверг","Пятница")
//
//	case "пт":
//		fmt.Println("Пятница, это последний рабочий день на этой неделе")
//
//	default:
//		fmt.Println("Такого дня нет, либо это выходной")
//
//	}
//
//}

//
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Каждый покупатель может купить только один лимонад и заплатить купюрой на 5, 10 или 20 долларов. Вы должны дать каждому покупателю сдачу с его купюры.")
//	var bills = []int{5, 10, 5, 20}
//	fmt.Println(lemonadeChange(bills))
//	bills = []int{5, 10, 5, 20, 20}
//	fmt.Println(lemonadeChange(bills))
//}
//
//func lemonadeChange(bills []int) bool {
//	five, ten := 0, 0
//
//	for _, i := range bills {
//		switch i {
//		case 5:
//			five++
//		case 10:
//			five--
//			ten++
//		case 20:
//			if ten >= 0 {
//				ten--
//				five--
//			} else {
//				five -= 3
//			}
//		}
//		if five < 0 || ten < 0 {
//			return false
//		}
//	}
//	return true
//}
