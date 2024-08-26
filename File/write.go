package main

import "os"

func main() {
	data1 := []byte("Hello\n Sanya")
	err := os.WriteFile("data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	f, ferr := os.Create("employeeName")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()

	data2 := []byte("Sanya\nMako")
	os.WriteFile("employeeName.txt", data2, 0644)
}
