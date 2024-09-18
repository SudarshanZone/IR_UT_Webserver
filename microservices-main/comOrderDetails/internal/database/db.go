package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/comOrderDetails/config"
	_ "github.com/lib/pq"
)



var db *sql.DB

// Init initializes the database connection
func Init() {
	

	dbConfig, err := config.LoadConfig("dbconfig.init")
	if err!=nil{
		log.Fatal("Error Loading config: %v\n",err)
	}
	
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	log.Println("Opening database connection...")
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	log.Println("Pinging database to verify connection...")
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	log.Println("Querying PostgreSQL version...")
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Error querying PostgreSQL version: %v\n", err)
	}

	log.Printf("PostgreSQL version: %s", version)
}

// GetDB returns the database connection instance
func GetDB() *sql.DB {
	if db == nil {
		log.Println("Warning: GetDB called but db is nil. Make sure Init is called first.")
	}
	return db
}

// OrdOrdrDtls represents a row in the COD_COD_ORDR_DTLS table
type ComOrdrDtls struct {
	COD_CLM_MTCH_ACCNT sql.NullString
	COD_UNDRLYNG       sql.NullString
	COD_PRDCT_TYP      sql.NullString
	COD_EXPRY_DT       sql.NullString
	COD_ORDR_VALID_DT  sql.NullString
	COD_LMT_RT         sql.NullFloat64
	COD_ORDR_FLW       sql.NullString
	COD_ORDR_TOT_QTY   sql.NullInt32
	COD_ORDR_STTS      sql.NullString
	CCP_OPNPSTN_QTY    sql.NullInt32
}
