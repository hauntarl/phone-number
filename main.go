package main

import (
	"log"

	"github.com/hauntarl/phone-number/db"
)

const database = "phone-number.db"

func main() {
	db, err := db.Open(database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
