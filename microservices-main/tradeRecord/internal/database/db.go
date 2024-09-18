package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/tradeRecord/config"
	_ "github.com/lib/pq"
)

func init() {
	dbConfig, err := config.LoadConfig("dbconfig.ini")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	fmt.Println("Connecting to the database...")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}