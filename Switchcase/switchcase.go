package main

import (
	"fmt"
)

var input string
var code string

func main() {
	fmt.Scanf("%s", &input)
	switch input {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("FFFF00")
	// เมื่อไม่เข้าเงื่อนไขใดใด
	default:
		fmt.Println("No Color")
	}
}
