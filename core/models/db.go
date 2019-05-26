package models

import (
	_ "github.com/lib/pq" // Init is Go Postgres driver for the database/sql package.

	"database/sql"
	"log"
)

var db *sql.DB

// InitDB init data base
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Printf("error open conenct to db")
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("error ping db")

		log.Panic(err)
	}
}
