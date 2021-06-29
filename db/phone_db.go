package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hauntarl/phone-number/normalize"
	_ "modernc.org/sqlite"
)

const (
	dbName    = "phone-numbers.db"
	tableName = "phone_numbers"
)

var db *sql.DB

// Open creates a new sqlite database file, returns the pointer to database.
func Open() (err error) {
	defer func() {
		if err == nil {
			abs, _ := filepath.Abs(dbName)
			log.Printf("Database Created: (location=%v)", abs)
		}
	}()

	f, err := os.Create(dbName)
	if err != nil {
		return err
	}
	f.Close()

	db, err = sql.Open("sqlite", dbName)
	return err
}

// Close closes the database and prevents new queries from starting.
func Close() {
	db.Close()
	log.Println("Database Closed")
}

// CreateTable runs a query to create a table to store phone number data.
func CreateTable() (err error) {
	defer func() {
		if err == nil {
			log.Printf("Table Created: (name=%v)", tableName)
		}
	}()

	stmt := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		number VARCHAR(255)
	)`, tableName)
	_, err = db.Exec(stmt)
	return
}

// InsertNumber inserts given phone number into database and returns uid
func InsertNumber(value string) (uid int64, err error) {
	defer func() {
		if err == nil {
			log.Printf("Inserted '%v': (uid=%v)", value, uid)
		}
	}()

	stmt := fmt.Sprintf(`INSERT INTO %s(number) values(?)`, tableName)
	res, err := db.Exec(stmt, value)
	if err != nil {
		return
	}

	return res.LastInsertId()
}

func SelectAll() ([]PhoneNumber, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	var (
		res []PhoneNumber
		uid int64
		num string
	)
	for rows.Next() {
		err = rows.Scan(&uid, &num)
		if err != nil {
			return nil, err
		}
		res = append(res, PhoneNumber{uid, num})
	}
	return res, nil
}

// PhoneNumber is Go representation for phone_number table in database.
type PhoneNumber struct {
	uid int64
	num string
}

// Uid simply returns the generated uid in database for given number.
func (pn PhoneNumber) Uid() int64 { return pn.uid }

// Generic displays the phone number in "(223) 456-7890" format, works on the
// assumption that number is normalized.
func (pn PhoneNumber) Generic() string { return normalize.Format(pn.num) }

// String simply returns the number stored in database as is.
func (pn PhoneNumber) String() string { return pn.num }
