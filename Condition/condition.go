package main

import (
	"fmt"
)

var score int

func main() {
	// myMoney := 100
	// if myMoney > 100 {
	// 	fmt.Println("You can")
	// } else {
	// 	fmt.Println("You can't")
	// }

	fmt.Println("Grade Calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("Score=", score)

	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
