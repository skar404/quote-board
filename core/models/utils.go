package models

import "log"

// PingDB ping data base and logs error
func PingDB() error {
	err := db.Ping()
	if err != nil {
		log.Println(err)
	}
	return err
}
