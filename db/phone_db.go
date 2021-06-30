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
	return res, rows.Err()
}

// Normalize iterates over all the numbers from database and normalizes them.
func Normalize(numbers []PhoneNumber) error {
	log.Println("Normalizing all the phone numbers...")
	for _, phn := range numbers {
		// try to normalize current phone number
		nrm, err := normalize.Number(phn.Val)
		if err != nil {
			// if normalization fails, delete current phone number
			if err = DeleteNumber(phn, err.Error()); err != nil {
				return err
			}
			continue
		}

		// if phone number changes, check if duplicate exists
		old, err := FindNumber(nrm)
		if err != nil {
			return err
		} else if old != nil {
			// if duplicate found, delete curren phone number
			if err = DeleteNumber(phn, "duplicate phone number"); err != nil {
				return err
			}
		} else if nrm != phn.Val {
			// if no duplicates, update the value of current phone number
			if err = UpdateNumber(phn, nrm); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteNumber deletes given phone number from database
func DeleteNumber(phn PhoneNumber, msg string) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE uid=?", tableName)
	res, err := db.Exec(query, phn.Uid)
	defer func() {
		if err != nil {
			return
		}
		count, _ := res.RowsAffected()
		log.Printf(
			"Deleted %-20s: (uid=%-3d, rows affected=%d), Reason: %s",
			phn.Val, phn.Uid, count, msg,
		)
	}()
	return
}

// FindNumber looks up given number in database and returns matched row
func FindNumber(val string) (*PhoneNumber, error) {
	var phn PhoneNumber
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE number=?", tableName)
	err := db.QueryRow(stmt, val).Scan(&phn.Uid, &phn.Val)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &phn, err
}

// UpdateNumber updates new val at given uid, else returns error.
func UpdateNumber(phn PhoneNumber, new string) (err error) {
	query := fmt.Sprintf("UPDATE %s SET number=? WHERE uid=?", tableName)
	res, err := db.Exec(query, new, phn.Uid)
	defer func() {
		if err != nil {
			return
		}
		count, _ := res.RowsAffected()
		log.Printf(
			"Updated %s to %s: (uid=%-3d, rows affected=%d)",
			phn.Val, new, phn.Uid, count,
		)
	}()
	return
}

// PhoneNumber is Go representation for phone_number table in database.
type PhoneNumber struct {
	Uid int64
	Val string
}

// Format displays the phone number in "(223) 456-7890" format.
func (pn PhoneNumber) Format() string { return normalize.Format(pn.Val) }
