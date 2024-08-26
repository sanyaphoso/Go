package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

// ตัวแปร pointer ต้องมี *
func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i=", i)

	zerovalue(i)
	fmt.Println("i=", i)

	// ใช้ & เพื่อเข้าถึง pointer
	zeropointer(&i)
	fmt.Println("i=", i)
	fmt.Println("i address=", &i)

}
