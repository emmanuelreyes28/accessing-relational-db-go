package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// db handle (pointer to an sql.DB)
// In production, you’d avoid the global variable, such as by passing
// the variable to functions that need it or by wrapping it in a struct.
var db *sql.DB

func main() {
	// capture connection properties and format them into a DSN for a connection string
	// DSN data source name
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// get a database handle and check for an error
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err) //ends execution and prints the error. In production, handle errors in a more graceful way
	}

	// use ping to confirm that connecting to db works
	// check for ping error if connection fails
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// print if Ping connects successfully
	fmt.Println("Connected!")
}
