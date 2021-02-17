package main

import "fmt"

func numberOne(firstNumber int) {
	fmt.Println(firstNumber)
}
func numberTwo(secondNumber int) {
	fmt.Println(secondNumber)
}
func finalNumber(first func(int), second func(int), firstNumber, secondNumber int) {
	second(secondNumber)
	first(firstNumber)
}
func main() {
	finalNumber(numberOne, numberTwo, 55, 23)
}
