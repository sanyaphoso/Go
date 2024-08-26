// package main

// import (
// 	"fmt"
// 	"time"
// )

// func process1(c chan string, data string) {
// 	c <- data
// }

// func main() {
// 	ch := make(chan string)
// 	go process1(ch, "Data1")
// 	fmt.Println(<-ch)
// 	time.Sleep(5 * time.Second)
// }

package main

import (
	"fmt"
	"time"
)

// ฟังก์ชันที่จะรันใน Goroutine
func sendData(ch chan string, data string, delay time.Duration) {
	time.Sleep(delay) // หน่วงเวลาเพื่อจำลองการทำงานที่ใช้เวลา
	ch <- data        // ส่งข้อมูลเข้าไปในช่องสัญญาณ
}

func main() {
	ch := make(chan string) // สร้างช่องสัญญาณ

	// เรียกใช้ Goroutines หลายตัว
	go sendData(ch, "Hello from Goroutine 1", 2*time.Second)
	go sendData(ch, "Hello from Goroutine 2", 1*time.Second)
	go sendData(ch, "Hello from Goroutine 3", 3*time.Second)

	// รับข้อมูลจากช่องสัญญาณและพิมพ์ออกมา
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
