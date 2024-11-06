package atdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PostgresConnect(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
