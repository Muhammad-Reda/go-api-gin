package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDB() (*sql.DB, error) {
	// Load env
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	driverName := os.Getenv("DRIVER_NAME")
	host := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := dbUser + ":" + dbPassword + "@" + "tcp(" + host + ")/" + dbName + "?parseTime=true"

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	log.Println("Database connected")
	return db, nil
}
