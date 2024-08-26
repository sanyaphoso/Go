// demoGoWithMysql.go

// demoConnectDB.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// creat table
func creatingTable(db, *sql.DB) {
	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Insert into table
func Insert(db, *sql.DB) {
	var username string
	var password string

	fmt.Scan(&username)
	fmt.Scan(&password)

	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO user (username, password, created_at) VALUE (?, ?, ?)`, username, password, createdAt)

	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}

// Delete from id
func delete(db, *sql.DB) {
	var deleteId int
	fmt.Scan(&deleteId)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, deleteId)
	if err != nil {
		log.Fatal(err)
	}
}

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)

	var inputId int
	fmt.Scan(&inputId)
	query := "SELECT id, coursename, price, instructor FROM coursedb.onlinecourse WHERE id = ?"
	if err := db.QueryRow(query, inputId).Scan(&id, &coursename, &price, &instructor); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, coursename, price, instructor)
}

func main() {
	db, err := sql.Open("mysql", "root")
	if err != nil {
		fmt.Println("Failed to connect")
	} else {
		fmt.Println("Connect successfully")
	}
	// query(db)
	// creatingTable(db)
	// Insert(db)
	delete(db)
}
