package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		massage, _ := fmt.Scanf("%v must number only", promt)
		panic(massage)
	}

	return value
}

func getOperator() string {
	fmt.Println("Operator is (+, -, *, /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func minus(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

var result float64

func main() {
	value1 := getInput(" value1 = ")
	value2 := getInput(" value2 = ")

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = minus(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("Wrong Operator")
	}

	fmt.Println("Result:", result)
}
