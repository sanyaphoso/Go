package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

// func plus(value1 int, value2 int) {
// 	result := value1 + value2
// 	fmt.Println("Result=", result)

// }

func plus(value1 int, value2 int) int {
	return value1 + value2

}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	// plus(5, 5)
	result := plus(5, 5)
	fmt.Println("Result=", result)

	result2 := plus3value(5, 5, 5)
	fmt.Println("Result2=", result2)
}
