package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB connect and validates DB connection
func ConnectDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
