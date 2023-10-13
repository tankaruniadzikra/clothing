package handler

import "database/sql"

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func teardownTestDB(db *sql.DB) {
}
