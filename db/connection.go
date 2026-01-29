package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env values: %v", err)
	}

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
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	CreateTable(DB)
}
