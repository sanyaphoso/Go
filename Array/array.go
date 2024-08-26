package main

import "fmt"

// create array ขนาด 4 ช่อง มี Type เป็น String
var productName [4]string
var price [4]float64

func main() {
	productName[0] = "MacBook"
	productName[1] = "IPad"
	productName[2] = "IPhone"
	productName[3] = "AirPods"

	// create array short style
	price2 := [4]float64{40000, 30000, 20000, 2000}

	fmt.Println(productName)
	fmt.Println(price)
	fmt.Println(price2)
}
