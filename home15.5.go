//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	fmt.Println("----Ввод данных----")
//	var arr [100]int
//	var n int
//	fmt.Println("Введите колличество элементов в массиве: ")
//	fmt.Scan(&n)
//	for i := 0; i < n; i++ {
//		fmt.Printf("массив[%v]= ", i)
//		fmt.Scan(&arr[i])
//	}
//	sumOdd := 0
//	sumPar := 0
//
//	fmt.Println("----Вывод данных----")
//	for i := 0; i < n; i++ {
//		if arr[i]%2 == 0 {
//			sumPar++
//		} else {
//			sumOdd++
//		}
//	}
//
//	fmt.Printf("Нечетных чисел в массиве: %v\n", sumOdd)
//	fmt.Printf("Четных чисел в массиве: %v\n", sumPar)
//
//}

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Массив %v Реверсия: %v\n", arr, reverseArr(arr))
}

func reverseArr(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseArr(input[1:]), input[0])
}
