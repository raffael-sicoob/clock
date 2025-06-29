package database

import (
	"log"
	"os"

	bolt "go.etcd.io/bbolt"
)




var db *bolt.DB
const (
	ClockBucket = "clock"
)




func InitDB() {
	basePath := os.Getenv("HOME") + "/.clock/data"

	os.MkdirAll(basePath, 0755)
	dbPath := basePath + "/clock.db"

	var err error

	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}


	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(ClockBucket))
		if err != nil {
			return err
		}
		return nil
	})

	
}


func CloseDB() {
	db.Close()
}




