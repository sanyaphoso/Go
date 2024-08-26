package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {
	data, _ := json.Marshal(employee{
		101,
		"Sanya Phoso",
		"0628618027",
		"sanyazza.nat@gmail.com",
	})
	fmt.Println(string(data))
}
