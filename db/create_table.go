package db

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	createTableSql := `
	 CREATE TABLE IF NOT EXISTS transactions(
	 	id SERIAL PRIMARY KEY,
		amount NUMERIC,
		phone VARCHAR(100),
		reason VARCHAR(100),
		reference VARCHAR(100),
		status VARCHAR(100),
		created_at DATE
	 )
	`
	_, err := db.Exec(createTableSql)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
	return nil
}
