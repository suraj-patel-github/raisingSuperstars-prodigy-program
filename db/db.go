package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=rspp dbname=rspp sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error connecting to the Database")
		log.Fatal(err)
	}
}
