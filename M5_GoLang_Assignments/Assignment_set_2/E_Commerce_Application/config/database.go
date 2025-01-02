package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// DB will hold the database connection pool.
var DB *sql.DB

// InitializeDatabase initializes the database connection.
func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	// Create the products table if not exists
	createTableSQL := `CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        description TEXT,
        price REAL,
        stock INTEGER,
        category_id INTEGER
    );`
	if _, err := DB.Exec(createTableSQL); err != nil {
		log.Fatal("Error creating table: ", err)
	}
}
