package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/go_api_gin?parseTime=true")
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	return db, nil
}
