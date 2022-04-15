package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-go")

	if err != nil {
		log.Fatal(err)
	}
}

func CreateCon() *sql.DB {
	return db
}
