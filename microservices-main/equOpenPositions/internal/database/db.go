package database

import (
	"fmt"
	"log"

	"github.com/krishnakashyap0704/microservices/equOpenPositions/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var db *gorm.DB
var gormDB *gorm.DB

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	return db
}

// Init initializes the database connection using GORM
func Init() {
	dbConfig, err :=config.LoadConfig("dbconfig.init")
	//var err error
	// Construct the PostgreSQL connection string
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

		if err!=nil{
			log.Fatal("Error loading config: %v\n,", err)
		}

	//var err error
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	// Verify the connection to the database
	sqlDB, err := db.DB() // Get the underlying *sql.DB
	if err != nil {
		log.Printf(" Failed to get underlying *sql.DB: %v", err)

	}
	if err := sqlDB.Ping(); err != nil {
		log.Printf(" Failed to ping database: %v", err)

	}

	fmt.Println("Successfully connected to PostgreSQL!")

	var version string
	err = db.Raw("SELECT version()").Scan(&version).Error
	if err != nil {
		log.Fatalf("Error querying the PostgreSQL version: %v\n", err)
	}

	fmt.Println("PostgreSQL version:", version)
}

// Equity_Positions represents an equity position in the database
type Epb_em_pstn_book struct {
	Epb_clm_mtch_accnt      string
	Epb_xchng_cd            string
	Epb_xchng_sgmnt_cd      string
	Epb_xchng_sgmnt_sttlmnt int32
	Epb_stck_cd             string
	Epb_orgnl_pstn_qty      int32
	Epb_rate                float64
	Epb_orgnl_amt_payble    float64
	Epb_orgnl_mrgn_amt      float64
	Epb_sell_qty            int32
	Epb_cvr_ord_qty         int32
	Epb_net_mrgn_amt        float64
	Epb_net_amt_payble      float64
	Epb_net_pstn_qty        int32
	Epb_ctd_qty             int32
	Epb_pstn_stts           string
	Epb_lpc_calc_stts       string
	Epb_sqroff_mode         string
	Epb_pstn_trd_dt         string
	Epb_mtm_prcs_flg        string
	Epb_last_mdfcn_dt       string
	Epb_ins_date            string
	Epb_close_date          string
	Epb_sys_fail_flg        string
	Epb_last_pymnt_dt       string
	Epb_lpc_calc_end_dt     string
	Epb_mtm_cansq           string
	Epb_expiry_dt           string
	Epb_min_mrgn            float64
	Epb_mrgn_dbcr_prcs_flg  string
	Epb_dp_id               string
	Epb_dp_clnt_id          string
	Epb_pledge_stts         string
	Epb_btst_net_mrgn_amt   float64
	Epb_btst_mrgn_blckd     float64
	Epb_btst_mrgn_dbcr_flg  string
	Epb_btst_sgmnt_cd       string
	Epb_btst_stlmnt         int32
	Epb_btst_csh_blckd      float64
	Epb_btst_sam_blckd      float64
	Epb_btst_calc_dt        string
	Epb_dbcr_calc_dt        string
	Epb_nsdl_ref_no         string
	Epb_mrgn_withheld_flg   string
}
