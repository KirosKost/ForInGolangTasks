package main
//
//import "fmt"
//
//func main() {
//
//	var userNumber int
//
//
//	fmt.Println("Введите число:")
//	fmt.Scan(&userNumber)
//
//	for i := 0; i < userNumber; i = i + 1{
//		fmt.Println("Вывод:", i)
//	}
//
//}

//----------------------------------------
//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Сумма двух чисел по единице")
//
//	var numberOne int
//	var numberTwo int
//
//	fmt.Println("Введите первое число:")
//	fmt.Scan(&numberOne)
//
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&numberTwo)
//
//	result := numberOne + numberTwo
//	for i := 1; i <= result; i++ {
//		fmt.Println(i)
//	}
//}

//---------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Расчёт суммы скидки")
//
//	var cost int
//	discount := 30
//
//
//	fmt.Println("Введите цену товара:")
//	fmt.Scan(&cost)
//
//
//	for i := 0; i < 2000; i = i + 1{
//		result := cost * discount / 100
//		// fmt.Println("Цена скидки:", result)
//		if result < 2000{
//			fmt.Println("Цена скидки:", result)
//			break
//		} else {
//			fmt.Println("Стоимость товара без скидки:", cost)
//
//		}
//		break
//	}
//}

//-----------------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//
//	fmt.Println("Предыдущий урок на if")
//
//	first := 0
//	second := 0
//	third := 0
//
//	for i := 0; i < 10000; i++ {
//		if first < 10 {
//			first++
//			continue
//		}
//
//		if second < 100 {
//			second++
//			continue
//		}
//
//		if third < 1000 {
//			third++
//			if third >= 1000 {
//				break
//			}
//
//		}
//
//	}
//	fmt.Println(first, second, third)
//
//}
//-------------------------

//package main
//
//import (
//"fmt"
//)
//
//
//func main() {
//	var capacity1, capacity2, capacity3 int
//	fmt.Print("Введите ёмкость первой корзины:")
//	fmt.Scan(&capacity1)
//	fmt.Print("Введите ёмкость второй корзины:")
//	fmt.Scan(&capacity2)
//	fmt.Print("Введите ёмкость третьей корзины:")
//	fmt.Scan(&capacity3)
//
//	var basket1, basket2, basket3 int
//	for {
//		if basket1 < capacity1 {
//			basket1++
//			continue
//		}
//
//		if basket2 < capacity2 {
//			basket2++
//			continue
//		}
//
//		if basket3 < capacity3 {
//			basket3++
//			continue
//		}
//
//		break
//	}
//
//	fmt.Println("Корзины заполнены")
//	fmt.Print(basket1, basket2, basket3)
}