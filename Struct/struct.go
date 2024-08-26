package main

import "fmt"

type employee struct {
	employeeId   string
	employeeName string
	phone        string
}

func main() {
	// employee1 := employee{
	// 	employeeId:   "101",
	// 	employeeName: "Sanya",
	// 	phone:        "0628618207",
	// }
	// fmt.Println("Employee1:", employee1)

	// Struct with Array
	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeId:   "101",
	// 	employeeName: "Sanya",
	// 	phone:        "0628618207",
	// }
	// employeeList[1] = employee{
	// 	employeeId:   "102",
	// 	employeeName: "Mako",
	// 	phone:        "0628612315",
	// }
	// employeeList[2] = employee{
	// 	employeeId:   "103",
	// 	employeeName: "Imt",
	// 	phone:        "0628611564",
	// }

	// fmt.Println("Employee:", employeeList)

	// Struct with Slice
	employeeList := []employee{}
	employee1 := employee{
		employeeId:   "101",
		employeeName: "Sanya",
		phone:        "0628618207",
	}
	employee2 := employee{
		employeeId:   "102",
		employeeName: "Mako",
		phone:        "0628612315",
	}
	employee3 := employee{
		employeeId:   "103",
		employeeName: "Imt",
		phone:        "0628611564",
	}

	employeeList = append(employeeList, employee1, employee2, employee3)

	fmt.Println("Employee:", employeeList)
}
