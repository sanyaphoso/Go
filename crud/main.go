package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// สร้างโครงสร้างข้อมูล
type Book struct {
	ID     int     `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// ประก่าศตัวแปลที่ใข้งาน struct
var books []Book

// สร้าง function อ่านข้อมูล มีการรับพารามิเตอร์สองตัว (w คือการเขียนข้อมูลลงไป หรือ response), (r การ request)
// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// **** อ่านข้อมูล ****
	// สร้าง Header
	w.Header().Set("Content-Type", "application/json")
	// return ด้วยการ encoder
	json.NewEncoder(w).Encode(books)

}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// สร้าง parameter มาอ่าน id
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create [POST]
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//สร้างตัวแปลรับข้อมูล
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	// เอาข้อมูลใส่ Book
	books = append(books, book)

	// อ่านข้อมูลออกมา
	json.NewEncoder(w).Encode(book)
}

// Update [PUT]
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// เช็ค ID
	for index, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			// เอา index ตัวสุดท้าย และตัวถัดๆไป
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// **** เรียกใช้ ****
	// สร้างตัวจัดการ route (Handle) ใช้ตัว Gorilla Mux มาช่วยจัดการ function
	r := mux.NewRouter()

	// สร้างข้อมูลจำลอง
	books = append(books, Book{ID: 1, Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: 2, Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
