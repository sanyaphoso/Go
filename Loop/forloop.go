package main

import "fmt"

const count = 15

var input string

func main() {
	// j := 0
	// for j < 10 {
	// 	fmt.Println("j:", j)
	// 	j = j + 1
	// }

	// for i := 0; i < count; i++ {
	// 	fmt.Println("i:", i)
	// }

	// in go don,t have while loop
	// While Loop

	for {
		fmt.Scanf("%s", &input)
		fmt.Println("Input:", input)
		if input == "exit" {
			break
		}
	}

}
