package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Id           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"Id":101, "EmployeeName":"Sanya Phoso", "Phone":"0628618027", "Email":"0628618027"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.EmployeeName)
}
