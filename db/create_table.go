package db

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	createUserTableSql := `
	 CREATE TABLE IF NOT EXISTS users(
	 	id SERIAL PRIMARY KEY,
		username VARCHAR(100),
		phone VARCHAR(100),
		role VARCHAR(100),
		status VARCHAR(100),
		created_at DATE
	 )
	`
	_, err := db.Exec(createUserTableSql)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	createTransactionTableSql := `
	 CREATE TABLE IF NOT EXISTS transactions(
	 	id SERIAL PRIMARY KEY,
		user_id INT, 
		amount NUMERIC,
		phone VARCHAR(100),
		reason VARCHAR(100),
		reference VARCHAR(100),
		status VARCHAR(100),
		created_at DATE, 
    CONSTRAINT fk_user
      FOREIGN KEY(user_id)
        REFERENCES users(id)
	 )
	`
	_, err = db.Exec(createTransactionTableSql)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	fmt.Println("Tables created successfully.")
	return nil
}
