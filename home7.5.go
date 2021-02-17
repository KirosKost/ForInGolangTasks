//package main
//
//import (
//  "fmt"
//  // "math"
//)
//func main() {
//  fmt.Println(" Зеркальные билеты")
//
//  sum := 0
//  begin := 100000
//  end := 999999
//
//  for i := begin; i <= end; i++ {
//
//    number1 := i / 100000
//    number2 := (i % 100000) / 10000
//    number3 := (i % 10000) / 1000
//    number4 := (i % 1000) / 100
//    number5 := (i % 100) / 10
//    number6 := i % 10
//
//    if number1 == number6 && number2 == number5 && number3 == number4  {
//      sum++
//
//    }
//
//}
//    fmt.Println(sum)
//
//}
//
//
//// package main
//
//// import (
//// 	"fmt"
//// 	"math"
//// )
//
//// func main() {
////   fmt.Println("Шахматная доска")
//
////   fmt.Println("Введите размер доски в клетках:")
////   var user int
////   fmt.Scan(&user)
//
//
////   for i := 0; i < user; i++ {
//
////    fmt.Println(" * *")
////    fmt.Println("* * ")
////   }
//
//// }
//
//// package main
//
//// import "fmt"
//
//// func main() {
////   fmt.Println("Вывод елочки")
//
//// 	fmt.Println("Введите желаемый размер ёлочки:")
//// 	var user int
//// 	fmt.Scan(&user)
//
//
//// 	for i := 1; i <= user; i++ {
//// 		for j := 1; j <= user-i; j++ {
//// 			fmt.Print(" ")
//// 		}
//// 		for j := 1; j <= 2 * i - 1; j++ {
//// 			fmt.Print("*")
//// 		}
//// 		fmt.Print("\n")
//// 	}
//// }
//
//
//
//
//// Reverse чисел в Go
//// package main
//
//// import (
//// 	"fmt"
//// )
//
//// func main() {
//// 	var n int
//// 	var reverse int = 0
//// 	fmt.Println("Enter a number")
//// 	fmt.Scanln(&n) // Take input from user
//// 	for n != 0 {
//// 		reverse = reverse * 10
//// 		reverse = reverse + n%10
//// 		n = (n / 10)
//// 	}
//// 	fmt.Println(reverse)
//// }