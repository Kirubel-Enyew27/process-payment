package db

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return DB, fmt.Errorf("failed to load .env values: %v", err)
	}

	viper.AutomaticEnv()

	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")
	user := viper.GetString("POSTGRES_USERNAME")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbname := viper.GetString("POSTGRES_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return DB, fmt.Errorf("Failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return DB, fmt.Errorf("failed to ping DB: %v", err)
	}

	if err := CreateTables(DB); err != nil {
		return DB, err
	}

	return DB, nil
}
