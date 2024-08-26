// package main

// import "fmt"

// func main() {
// 	var courseName []string
// 	courseName = []string{"Java", "Python"}
// 	fmt.Println(courseName)
// 	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
// 	fmt.Println(courseName)

// 	courseWeb := courseName[4:7]
// 	fmt.Println(courseWeb)
// }

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before deletion:", slice)

	// ลบข้อมูลที่ตำแหน่ง index 2 (ค่า 3)
	index := 2
	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println("After deletion:", slice)
}

