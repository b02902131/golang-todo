package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// initDB initializes the database and returns a database handle
func InitDB() *sql.DB {
	log.Println("Initializing database...")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set.")
	} else {
		log.Printf("Connecting to database at: %s", dbURL)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	} else {
		log.Printf("Connected to database: %s", dbURL)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		description TEXT,
		completed BOOLEAN
	);`

	log.Println("Creating todos table...")
	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatalf("Error creating todos table: %s", err)
	} else {
		log.Println("Successfully created todos table.")
	}

	return db
}
