package database

import (
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/comOpenPositions/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	return db
}

// Init initializes the database connection using GORM
func Init() {
	dbConfig, err := config.LoadConfig("dbconfig.init")
	if err!=nil{
		log.Fatal("Error Loading config: %v\n",err)
	}
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	// Verify the connection to the database
	sqlDB, err := db.DB() // Get the underlying *sql.DB
	if err != nil {
		log.Fatalf("Failed to get underlying *sql.DB: %v\n", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Check the PostgreSQL version
	var version string
	err = db.Raw("SELECT version()").Scan(&version).Error
	if err != nil {
		log.Fatalf("Error querying the PostgreSQL version: %v\n", err)
	}

	fmt.Println("PostgreSQL version:", version)
}

// CCP_COD_SPN_CNTRCT_PSTN represents a commodity position in the database
type CCP_COD_SPN_CNTRCT_PSTN struct {
	Ccp_clm_mtch_accnt       string  // Field 1
	Ccp_xchng_cd             string  // Field 2
	Ccp_prdct_typ            string  // Field 3
	Ccp_indstk               string  // Field 4
	Ccp_undrlyng             string  // Field 5
	Ccp_expry_dt             string  // Field 6 (could be time.Time if you prefer)
	Ccp_exer_typ             string  // Field 7
	Ccp_strk_prc             int64   // Field 8
	Ccp_opt_typ              string  // Field 9
	Ccp_ibuy_qty             int64   // Field 10
	Ccp_ibuy_ord_val         float64 // Field 11
	Ccp_isell_qty            int64   // Field 12
	Ccp_isell_ord_val        float64 // Field 13
	Ccp_exbuy_qty            int64   // Field 14
	Ccp_exbuy_ord_val        float64 // Field 15
	Ccp_exsell_qty           int64   // Field 16
	Ccp_exsell_ord_val       float64 // Field 17
	Ccp_buy_exctd_qty        int64   // Field 18
	Ccp_sell_exctd_qty       int64   // Field 19
	Ccp_opnpstn_flw          string  // Field 20
	Ccp_opnpstn_qty          int64   // Field 21
	Ccp_opnpstn_val          float64 // Field 22
	Ccp_exrc_qty             int64   // Field 23
	Ccp_asgnd_qty            int64   // Field 24
	Ccp_opt_premium          float64 // Field 25
	Ccp_mtm_opn_val          float64 // Field 26
	Ccp_imtm_opn_val         float64 // Field 27
	Ccp_extrmloss_mrgn_extra float64 // Field 28
	Ccp_addnl_mrgn           float64 // Field 29
	Ccp_spcl_mrgn            float64 // Field 30
	Ccp_tndr_mrgn            float64 // Field 31
	Ccp_dlvry_mrgn           float64 // Field 32
	Ccp_extrm_min_loss_mrgn  float64 // Field 33
	Ccp_mtm_flg              string  // Field 34
	Ccp_extrm_loss_mrgn      float64 // Field 35
	Ccp_flat_val_mrgn        float64 // Field 36
	Ccp_trg_prc              float64 // Field 37
	Ccp_min_trg_prc          float64 // Field 38
	Ccp_devolmnt_mrgn        float64 // Field 39
	Ccp_mtmsq_ordcnt         int32   // Field 40
	Ccp_avg_prc              float64 // Field 41
}
