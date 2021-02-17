//package main
//
//import(
//	"fmt"
//	"math"
//)
//// Функция Тейлора для sin
//func main() {
//
//	var epsilon float64
//	epsilon = 0.0001
//
//	fmt.Println("Введите х для которого необходимо расчитать sin: ")
//	var x float64
//	fmt.Scan(&x)
//
//	result := 0.0
//	prevResult := 0.0
//	fact := 1
//	k := 0
//
//	for {
//		if k > 0 {
//			fact *= 2 * k *(2 * k + 1)
//		}
//		result += math.Pow(-1, float64(k))* math.Pow(x, float64(2*k+1))/ float64(fact)
//		if math.Abs(result - prevResult) < epsilon {
//			fmt.Println(result)
//			break
//		}
//		k++
//		prevResult = result
//	}
//
//}
////
////package main
////
////import(
////	"fmt"
////	"math"
////)
////
////func main() {
////
////	fmt.Println("Введите сумму вклада: ")
////	var sum float64
////	fmt.Scan(&sum)
////
////	fmt.Println("Введите процент: ")
////	var percent float64
////	fmt.Scan(&percent)
////
////	fmt.Println("Введите количество лет вклада: ")
////	var depositTime int
////	fmt.Scan(&depositTime)
////
////	var tail float64 = 0
////
////	for i := 0; i < depositTime * 12; i++ {
////		totalSum := sum + sum * percent / 100
////		sum = math.Trunc(totalSum * 100) / 100
////		tail += totalSum - sum
////	}
////
////	fmt.Printf("Сумма %v p. Срезано %v p.", sum, tail)
////
////}
//
//
////package main
////
////import (
////  "fmt"
////  "math"
////)
////
////func main() {
////
////  const max = 500
////
////  fmt.Println("Введите значение х: ")
////  var x float64
////  fmt.Scan(&x)
////
////  fmt.Println("Введите значение n: ")
////  var n float64
////  fmt.Scan(&n)
////
////  var epsilon float64 = 0.0000000001
////  var next float64 = 1.0
////  var result float64 = 1.0
////  var k float64 = 1
////  next = result
////  fact := 1.0
////  for  i := 1; i < 500 && math.Abs(next) > epsilon; i++{
////    k = float64(i)
////    fact *= k
////    result += math.Pow(x, k)/fact
////  }
////  fmt.Println(result)
////  fmt.Println(math.Exp(x))
////
////
////}
////
////
////
////
////
////
////package main
////
////import (
////	"fmt"
////	"math"
////)
////
////func main() {
////
////	fmt.Println("Enter x and n")
////	var x float64
////	var n float64
////	fmt.Scan(&x)
////	fmt.Scan(&n)
////
////	fact := 1.0
////	sum := 0.0
////	var i float64 = 1
////	for ; i < n; i++{
////		fact *= i
////		sum += math.Pow(x, i)/fact
////	}
////	fmt.Println(math.Exp(sum))
////	fmt.Println(math.Exp(x))
////
////
////
////}