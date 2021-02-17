//package main
//
//import "fmt"
//
//var a int
//
//func numberOne(a int) bool {
//	if a%2 == 0 {
//		fmt.Printf("Число %d четное\n", a)
//	}
//	return true
//}
//
//func numberTwo(b int) bool {
//	if b%2 != 0 {
//		fmt.Printf("Число %d не четное\n", b)
//	}
//	return false
//}
//
//func main() {
//	sum1 := numberOne(10)
//	fmt.Println(sum1)
//
//	sum2 := numberTwo(11)
//	fmt.Println(sum2)
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//func checkNum(num int) (int, bool) {
//	if num%2 == 0 {
//		return 1, true
//	} else {
//		return 0, false
//	}
//}
//
//func main() {
//	fmt.Println(checkNum(8))
//}



package main

import (
"fmt"
"math/rand"
"time"
)

func randomPoint() (int, int) {
	return rand.Int(), rand.Int()
}

func transformPoint(x, y int) (int, int) {
	return 2*x + 10, -3*y - 5 // возможно переполнение int, из-за этого x1 может стать негативным
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i<4; i++{
		x, y := randomPoint()
		xT, yT := transformPoint(x, y)
		fmt.Println("x",i," = ",xT, ", y",i," = ", yT)
	}

}

//package main
//
//import (
//	"fmt"
//)
//
//func numberOne(a int) (c int) {
//	c = a * 15
//	return
//}
//func numberTwo(b int) (d int) {
//	d = b + 23
//	return
//}
//
//func main() {
//	fmt.Println(numberTwo(numberOne(13)))
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//var a = 3
//var b = 5
//var c = 7
//
//func numOne(x int) (res int) {
//	res = x + b + c
//	return
//}
//
//func numTwo(x int) (res int) {
//	res = x + a + c
//	return
//}
//func numThree(x int) (res int) {
//	res = x + b + c
//	return
//}
//
//func main() {
//	fmt.Println(numOne(numTwo(numThree(23))))
//}

//package main
//
//import "fmt"
//
//var a int
//
//func number(a int) bool {
//	if a%2 == 0 {
//		return true
//	} else {
//		return false
//	}
//
//}
//
//func main() {
//	fmt.Println(number(13))
//}
