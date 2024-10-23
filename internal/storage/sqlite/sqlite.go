package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // init sqlite3 driver
)

type Storage struct {
	db *sql.DB
}

// конструктор стораджа
func New(storagePath string) (*Storage, error) {

	// для дирекции ошибок
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS url (
			id INTEGER PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON url (alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
