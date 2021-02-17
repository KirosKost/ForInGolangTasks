package main

import(
	"fmt"
	"os"
)

func main() {
	text := "Hello"
	file, err := os.Create("helloGo.txt")
	if err != nil{
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}