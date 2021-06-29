package phonedb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

const (
	dbName    = "phone-numbers.db"
	tableName = "phone_numbers"
)

var db *sql.DB

// Open creates a new sqlite database file, returns the pointer to database.
func Open() error {
	log.Printf("Creating Database: %v", dbName)
	f, err := os.Create(dbName)
	if err != nil {
		return err
	}
	f.Close()

	db, err = sql.Open("sqlite", dbName)
	return err
}

// Close closes the database and prevents new queries from starting.
func Close() { db.Close() }

// CreateTable runs a query to create a table to store phone number data.
func CreateTable() error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		number VARCHAR(255)
	)`, tableName)
	_, err := db.Exec(query)
	return err
}

// InsertNumber inserts given phone number into database and returns uid
func InsertNumber(value string) (uid int64, err error) {
	query := fmt.Sprintf(`INSERT INTO %s(number) values(?)`, tableName)
	res, err := db.Exec(query, value)
	if err != nil {
		return
	}

	return res.LastInsertId()
}
