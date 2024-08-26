package main

import (
	"fmt"
	// "reflect"
)

// func main() {
// 	var x float64 = 3.4
// 	fmt.Println("Type:", reflect.TypeOf(x))
// 	fmt.Println("Value:", reflect.ValueOf(x))
// }

// var numberInt int = 1000
var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello"

func main() {
	// short format using only in main
	numberfloat := 25.5
	
	fmt.Println(numberInt, numberInt2)
	fmt.Println(numberInt + numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(numberfloat + float64(numberInt))
	fmt.Println(msg + " World")
	fmt.Println("value:", numberInt)
}
