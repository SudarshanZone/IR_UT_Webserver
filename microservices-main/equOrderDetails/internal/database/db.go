package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/equOrderDetails/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	dbConfig, err :=config.LoadConfig("dbconfig.init")
	
	if err!=nil{
		log.Fatal("Error Loading config: %v\n",err)
	}
	// Construct the PostgreSQL connection string
	

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	log.Println("Connecting to the PostgreSQL database...")

	//var err error
	// Open a connection to the database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		
	}

	// Ping the database to ensure connection is established
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database.")

	// Fetch and log the PostgreSQL version
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Error fetching PostgreSQL version: %v", err)
	}
	log.Printf("PostgreSQL version: %s", version)
}

func GetDB() *sql.DB {
	log.Println("Returning the database connection instance.")
	return db
}

type EquityOrderDetails struct {
	Ord_Clm_Mtch_Accnt sql.NullString
	Ord_Stck_Cd        sql.NullString
	Ord_Ordr_Dt        sql.NullString
	Ord_Ordr_Flw       sql.NullString
	Ord_Ordr_Qty       sql.NullInt32
	Ord_Lmt_Rt         sql.NullFloat64
	Ord_Ordr_Stts      sql.NullString
	// Additional fields can be uncommented as needed
}

// Function to log details about the order
func (o *EquityOrderDetails) LogOrderDetails() {
	log.Printf("Order Details: Account: %s, Stock Code: %s, Order Date: %s, Order Flow: %s, Order Quantity: %d, Limit Rate: %f, Order Status: %s",
		o.Ord_Clm_Mtch_Accnt.String, o.Ord_Stck_Cd.String, o.Ord_Ordr_Dt.String, o.Ord_Ordr_Flw.String, o.Ord_Ordr_Qty.Int32, o.Ord_Lmt_Rt.Float64, o.Ord_Ordr_Stts.String)
}
