package models

import "log"

// ping data base
func PingDB() error {
	err := db.Ping()
	if err != nil {
		log.Println(err)
	}
	return err
}
