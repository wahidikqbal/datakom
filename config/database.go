package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// new database name datakom_db. root:root
func ConnectDB() {
	db, err := sql.Open("mysql", "root:root@/datakom_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db
}
