package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("Result:", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i:", i)
	}
}

func main() {
	// fmt.Println("Welcome")

	// // defer ทำงานหลังจบโปรแกรม ถ้ามีหลายตัวจะทำงานแบบ First in Last out
	// defer fmt.Println("End")
	// defer add(20, 10)
	// defer add(15, 15)
	// defer add(12, 10)
	loop()
	deferloop()

}
