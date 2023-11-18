package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// database name datakom_db. poop:poop
func ConnectDB() {
	db, err := sql.Open("mysql", "poop:poop@/datakom_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db
}
