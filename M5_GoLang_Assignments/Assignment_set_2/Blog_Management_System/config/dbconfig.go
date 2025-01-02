package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		log.Fatal(err)
	}

	// create blogs table
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		author TEXT
		);`)

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}
