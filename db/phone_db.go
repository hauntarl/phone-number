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
		if err != nil {
			return
		}

		abs, _ := filepath.Abs(dbName)
		log.Printf("Database Created: (location=%v)", abs)
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

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		number VARCHAR(255)
	)`, tableName)
	_, err = db.Exec(query)
	return
}

// InsertNumber inserts given phone number into database and returns uid.
func InsertNumber(value string) (uid int64, err error) {
	query := fmt.Sprintf(`INSERT INTO %s(number) values(?)`, tableName)
	res, err := db.Exec(query, value)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			return
		}

		count, _ := res.RowsAffected()
		log.Printf(
			"Inserted %-20v: (uid=%-3d, rows affected=%v)",
			value, uid, count,
		)
	}()
	return res.LastInsertId()
}

// SelectOne queries a single row from database based on given uid.
func SelectOne(uid int64) (phn PhoneNumber, err error) {
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE uid=?", tableName)
	err = db.QueryRow(stmt, uid).Scan(&phn.Uid, &phn.Val)
	return
}

// SelectAll fetches all the rows from phone_numbers table.
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
	Uid int64
	Val string
}

// Format displays the phone number in "(223) 456-7890" format, works on the
// assumption that number is normalized.
func (pn PhoneNumber) Format() string { return normalize.Format(pn.Val) }
