package main

import (
	"log"
	"path/filepath"

	"github.com/hauntarl/phone-number/db"
)

const (
	dbName    = "phone-numbers.db"
	tableName = "phone_numbers"
)

func main() {
	log.Printf("Creating Database: %v", dbName)
	dp, err := db.Open(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer dp.Close()
	abs, _ := filepath.Abs(dbName)
	log.Printf("Database Location: %v", abs)

	if err := db.CreateTable(dp, tableName); err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Table: %v", tableName)
}
