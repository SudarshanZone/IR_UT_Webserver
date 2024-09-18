package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/fnoSquareOff/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Init initializes the database connection
func Init() {
	// Load the database configuration
	dbConfig, err := config.LoadConfig("dbconfig.ini")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL version:", version)
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
