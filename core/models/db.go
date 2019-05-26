package models

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
)

var db *sql.DB

// InitDB init data base
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Printf("error level 1")
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("error level 1")

		log.Panic(err)
	}
}
