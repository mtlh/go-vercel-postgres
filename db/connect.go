package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// Construct connection string from environment variables
	connStr := GetConnectionString()
	// Open database connection
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func GetConnectionString() string {
	return "postgresql://" + os.Getenv("PGUSER") + ":" + os.Getenv("PGPASSWORD") + "@" + os.Getenv("PGHOST") + "/" + os.Getenv("PGDATABASE") + "?sslmode=require"
}
