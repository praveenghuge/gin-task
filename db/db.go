package db

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func db() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
