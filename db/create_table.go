package db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTable(db *sql.DB) {
	createTableSql := `
	 CREATE TABLE IF NOT EXISTS transactions(
	 	id SERIAL PRIMARY KEY,
		amount NUMBER,
		phone VARCHAR(100),
		reason VARCHAR(100),
		reference VARCHAR(100),
		status VARCHAR(100),
		created_at DATE
	 )
	`
	_, err := db.Exec(createTableSql)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully.")
}
