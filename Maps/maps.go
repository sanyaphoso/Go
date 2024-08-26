package main

import (
	"fmt"
)

var product = make(map[string]float64)

func main() {
	fmt.Println("Product: ", product)

	// append in maps
	product["Macbook"] = 40000
	product["IPhone"] = 30000
	product["IPad"] = 20000
	fmt.Println("Product: ", product)

	// delete
	delete(product, "IPad")
	fmt.Println("Product: ", product)

	//update
	product["Macbook"] = 50000
	fmt.Println("Product: ", product)

	// เข้าถึงข้อมูล
	value1 := product["IPhone"]
	fmt.Println("Value1: ", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName", courseName)
}
