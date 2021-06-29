package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Open creates a new sqlite database file,
// returns the pointer to database.
func Open(name string) (*sql.DB, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	f.Close()

	return sql.Open("sqlite", name)
}

func CreateTable(dp *sql.DB, name string) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL,
		value VARCHAR(255)
	)`, name)
	_, err := dp.Exec(query)
	return err
}
