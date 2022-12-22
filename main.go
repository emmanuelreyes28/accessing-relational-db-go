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

// struct is used to hold row data returned from the query
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// capture connection properties and format them into a DSN for a connection string
	// DSN data source name
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
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

	albums, err := albumByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

// albumByArtist queries for albums that have the specified artist name.
func albumByArtist(name string) ([]Album, error) {
	// an album slice to hold data from the returned rows
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	defer rows.Close() // defer closing rows so that any resources it holds will be released when the function exits.

	// loop through rows, using scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
