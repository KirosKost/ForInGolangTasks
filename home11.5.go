package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	s := "a10 10 20b 20 30c30 30 dd"
	space := " "
	for strings.Contains(s, space) {
		spaceIndex := strings.Index(s, space)
		words := s[:spaceIndex]
		i, err := strconv.Atoi(words)
		if err == nil {
			fmt.Println(i)
		}
		s = s[spaceIndex+1:]
		s = strings.Trim(s, space)

	}
	i, err := strconv.Atoi(s)
	if err == nil{
		fmt.Println(i)
	}
}
//
//
//
//
////package main
////
////import (
////	"fmt"
////)
////
////func main()  {
////	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
////
////	count := 0
////	for i := 0; i < len(s); i++{
////		if s[i] < 91 && s[i] > 64{
////			count++
////		}
////	}
////	fmt.Println(count)
////
////}
//
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main()  {
//	s := "Go это Язык программирования с Открытым исходным кодом, который позволяет Легко создавать простое, надежное и эффективное Программное обеспечение"
//	v := strings.Fields(s)
//
//	count := 0
//	for i := 0; i < len(v); i++{
//		if strings.Compare(v[i],strings.Title(v[i])) == 0{
//			count++
//		}
//	}
//	fmt.Println(count)
//
//}
//
