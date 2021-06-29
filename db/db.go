package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// Open creates a new sqlite database file,
// returns the pointer to database.
func Open(name string) (*sql.DB, error) {
	log.Printf("Creating Database: %v", name)
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	f.Close()

	abs, err := filepath.Abs(name)
	if err != nil {
		abs = "./" + name
	}
	log.Printf("Database Location: %v", abs)

	return sql.Open("sqlite", name)
}
