// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println("пользователь вводит 2 числа (int16), а программа выводит в какой минимальный тип данных можно сохранить результат умножения этих чисел.")

// 	fmt.Println("Введите первое число")
// 	var numberOne int16
// 	fmt.Scan(&numberOne)

// 	fmt.Println("Введите второе число")
// 	var numberTwo int16
// 	fmt.Scan(&numberTwo)
	
// 	result := int64(numberOne) * int64(numberTwo)
// 	if result > 0{
// 	if result <= math.MaxUint8 {
// 		fmt.Println(numberOne, numberTwo, "результат uint8")
	
// 		}else if result <= math.MaxUint16 {
// 			fmt.Println(numberOne, numberTwo, "результат uint16")
	
// 		}else{
// 			fmt.Println(numberOne, numberTwo, "результат uint32")
// 		}
	
// 	}else{
// 		if result >= math.MinInt8 {
// 			fmt.Println(numberOne, numberTwo, "результат int8")
	
// 		 }else if result >= math.MinInt16{
// 			fmt.Print(numberOne, numberTwo, "результат int16")
	
// 		}else{
// 			fmt.Println(numberOne, numberTwo, "результат int32")
// 	}
// }
// }



package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println("Подсчитывать сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32")

// 	var a uint8 
// 	var b uint16


// 	var counter1 uint32
// 	var counter2 uint32

// 	for i := 1; i <= math.MaxUint32; i++ {	
// 		if a+1 == 0 {
// 			counter1++
// 			a = 0
// 		}else{
// 			a++
// 		}
// 		if b+1 == 0 {
// 			counter2++
// 			b = 0
// 		}else{
// 			b++
// 		}
// 	}
// 	fmt.Println(counter1)
// 	fmt.Println(counter2)



// }