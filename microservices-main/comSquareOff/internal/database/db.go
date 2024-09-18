package database

import (
	"fmt"
	"log"
	"time"

	"github.com/krishnakashyap0704/microservices/comSquareOff/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Define structures for database tables
type CodCodOrdrDtls struct {
	CodClmMtchAccnt    string    `gorm:"column:cod_clm_mtch_accnt;type:char(10);not null"`
	CodClntCtgry       int       `gorm:"column:cod_clnt_ctgry;type:number(3);not null"`
	CodOrdrRfrnc       string    `gorm:"column:cod_ordr_rfrnc;type:char(18);not null"`
	CodPipeId          string    `gorm:"column:cod_pipe_id;type:char(2);not null"`
	CodXchngCd         string    `gorm:"column:cod_xchng_cd;type:char(3);not null"`
	CodPrdctTyp        string    `gorm:"column:cod_prdct_typ;type:char(1);not null"`
	CodIndstk          string    `gorm:"column:cod_indstk;type:char(1);not null"`
	CodUndrlyng        string    `gorm:"column:cod_undrlyng;type:char(6);not null"`
	CodExpryDt         string    `gorm:"column:cod_expry_dt;type:date;not null"`
	CodExerTyp         string    `gorm:"column:cod_exer_typ;type:char(1);not null"`
	CodOptTyp          string    `gorm:"column:cod_opt_typ;type:char(1);not null"`
	CodStrkPrc         float64   `gorm:"column:cod_strk_prc;type:number(10);not null"`
	CodOrdrFlw         string    `gorm:"column:cod_ordr_flw;type:char(1);not null"`
	CodLmtMrktSlFlg    string    `gorm:"column:cod_lmt_mrkt_sl_flg;type:char(1);not null"`
	CodDsclsdQty       float64   `gorm:"column:cod_dsclsd_qty;type:number(8)"`
	CodOrdrTotQty      float64   `gorm:"column:cod_ordr_tot_qty;type:number(8);not null"`
	CodLmtRt           float64   `gorm:"column:cod_lmt_rt;type:number(10)"`
	CodStpLssTgr       float64   `gorm:"column:cod_stp_lss_tgr;type:number(10)"`
	CodOrdrType        string    `gorm:"column:cod_ordr_type;type:char(1);not null"`
	CodOrdrValidDt     time.Time `gorm:"column:cod_ordr_valid_dt;type:date"`
	CodTrdDt           time.Time `gorm:"column:cod_trd_dt;type:date;not null"`
	CodOrdrStts        string    `gorm:"column:cod_ordr_stts;type:char(1);not null"`
	CodSprdOrdrRef     string    `gorm:"column:cod_sprd_ordr_ref;type:char(18)"`
	CodMdfctnCntr      int       `gorm:"column:cod_mdfctn_cntr;type:number(3);not null"`
	CodSettlor         string    `gorm:"column:cod_settlor;type:char(12)"`
	CodAckNmbr         string    `gorm:"column:cod_ack_nmbr;type:char(20)"`
	CodSplFlag         string    `gorm:"column:cod_spl_flag;type:char(1)"`
	CodOrdAckTm        time.Time `gorm:"column:cod_ord_ack_tm;type:date"`
	CodLstRqstAckTm    time.Time `gorm:"column:cod_lst_rqst_ack_tm;type:date"`
	CodProCliInd       string    `gorm:"column:cod_pro_cli_ind;type:char(1)"`
	CodExecQtyDay      float64   `gorm:"column:cod_exec_qty_day;type:number(8)"`
	CodRemarks         string    `gorm:"column:cod_remarks;type:varchar(256)"`
	CodChannel         string    `gorm:"column:cod_channel;type:varchar(3)"`
	CodBpId            string    `gorm:"column:cod_bp_id;type:varchar(8)"`
	CodCtclId          string    `gorm:"column:cod_ctcl_id;type:varchar(16)"`
	CodUsrId           string    `gorm:"column:cod_usr_id;type:varchar(15)"`
	CodMrktTyp         string    `gorm:"column:cod_mrkt_typ;type:char(1)"`
	CodCseId           int       `gorm:"column:cod_cse_id;type:number(6)"`
	CodSpnFlg          string    `gorm:"column:cod_spn_flg;type:char(1)"`
	CodSltpOrdrRfrnc   string    `gorm:"column:cod_sltp_ordr_rfrnc;type:char(18)"`
	CodAmtBlckd        float64   `gorm:"column:cod_amt_blckd;type:number(24,2)"`
	CodLssAmtBlckd     float64   `gorm:"column:cod_lss_amt_blckd;type:number(24,2)"`
	CodFcFlag          string    `gorm:"column:cod_fc_flag;type:char(1)"`
	CodDiffAmtBlckd    float64   `gorm:"column:cod_diff_amt_blckd;type:number(24,2)"`
	CodDiffLssAmtBlckd float64   `gorm:"column:cod_diff_lss_amt_blckd;type:number(24,2)"`
	CodTrdVal          float64   `gorm:"column:cod_trd_val;type:number(24,2)"`
	CodTrdBrkg         float64   `gorm:"column:cod_trd_brkg;type:number(12)"`
	CodCntrctntNmbr    string    `gorm:"column:cod_cntrctnt_nmbr;type:varchar(25)"`
	CodSourceFlg       string    `gorm:"column:cod_source_flg;type:char(1)"`
	CodEosFlg          string    `gorm:"column:cod_eos_flg;type:char(1)"`
	CodPrcimpvFlg      string    `gorm:"column:cod_prcimpv_flg;type:char(1)"`
	CodTrailAmt        float64   `gorm:"column:cod_trail_amt;type:number(12)"`
	CodLmtOffset       float64   `gorm:"column:cod_lmt_offset;type:number(12)"`
	CodSrollDiffAmt    float64   `gorm:"column:cod_sroll_diff_amt;type:number(24,2)"`
	CodSrollLssAmt     float64   `gorm:"column:cod_sroll_lss_amt;type:number(24,2)"`
	CodSrollDiffAmtOld float64   `gorm:"column:cod_sroll_diff_amt_old;type:number(24,2)"`
	CodSrollLssAmtOld  float64   `gorm:"column:cod_sroll_lss_amt_old;type:number(24,2)"`
	CodPanNo           string    `gorm:"column:cod_pan_no;type:varchar(10)"`
	CodSetlmntFlg      string    `gorm:"column:cod_setlmnt_flg;type:char(1)"`
	CodLstActRef       string    `gorm:"column:cod_lst_act_ref;type:char(21)"`
	CodQtyUnit         string    `gorm:"column:cod_qty_unit;type:varchar(8)"`
	CodPrcUnit         string    `gorm:"column:cod_prc_unit;type:varchar(20)"`
	CodPrcMltplr       float64   `gorm:"column:cod_prc_mltplr;type:number(20,4)"`
	CodGenMltplr       float64   `gorm:"column:cod_gen_mltplr;type:number(20,4)"`
	CodSsnTyp          string    `gorm:"column:cod_ssn_typ;type:char(2)"`
	CodIsFlg           string    `gorm:"column:cod_is_flg;type:char(1)"`
	CodExecQty         float64   `gorm:"column:cod_exec_qty;type:number(8)"`
	CodCnclQty         float64   `gorm:"column:cod_cncl_qty;type:number(8)"`
	CodExprdQty        float64   `gorm:"column:cod_exprd_qty;type:number(8)"`
	CodEspId           string    `gorm:"column:cod_esp_id;type:char(50)"`
	CodMinLotQty       float64   `gorm:"column:cod_min_lot_qty;type:number(10)"`
	CodPrtctnRt        float64   `gorm:"column:cod_prtctn_rt;type:number(12,2)"`
	CodSqroffTm        float64   `gorm:"column:cod_sqroff_tm;type:number(8)"`
	CodAvgExctdRt      float64   `gorm:"column:cod_avg_exctd_rt;type:number(17,7)"`
}

type CxbCodXchngBook struct {
	CxbXchngCd       string    `gorm:"type:varchar(10);not null"`          // Exchange Code
	CxbOrdrRfrnc     string    `gorm:"type:varchar(255);not null"`         // Order Reference Number
	CxbPipeId        string    `gorm:"type:varchar(50)"`                   // Pipeline ID
	CxbModTrdDt      string    `gorm:"column:cxb_mod_trd_dt;type:date;"`   // Modified Trade Date
	CxbOrdrSqnc      int64     `gorm:"type:bigint"`                        // Order Sequence
	CxbLmtMrktSlFlg  string    `gorm:"type:varchar(10)"`                   // Limit Market Sale Flag
	CxbDsclsdQty     int64     `gorm:"type:bigint"`                        // Disclosed Quantity
	CxbOrdrTotQty    int64     `gorm:"type:bigint"`                        // Order Total Quantity
	CxbLmtRt         float64   `gorm:"type:double precision"`              // Limit Rate
	CxbStpLssTgr     float64   `gorm:"type:double precision"`              // Stop Loss Trigger
	CxbMdfctnCntr    int       `gorm:"type:integer"`                       // Modification Counter
	CxbOrdrValidDt   string    `gorm:"column:cxb_ordr_valid_dt;type:date"` // Order Valid Date
	CxbOrdrType      string    `gorm:"type:varchar(20)"`                   // Order Type
	CxbSprdOrdInd    string    `gorm:"type:varchar(10)"`                   // Spread Order Indicator
	CxbRqstTyp       string    `gorm:"type:varchar(20)"`                   // Request Type
	CxbQuote         float64   `gorm:"type:double precision"`              // Quote
	CxbQtTm          string    `gorm:"column:cxb_qt_tm;type:date"`         // Quote Time
	CxbRqstTm        string    `gorm:"column:cxb_rqst_tm;type:date"`       // Request Time
	CxbFrwdTm        string    `gorm:"column:cxb_frwd_tm;type:date"`       // Forward Time
	CxbPlcdStts      string    `gorm:"type:varchar(20)"`                   // Placed Status
	CxbRmsPrcsdFlg   string    `gorm:"type:varchar(10)"`                   // RMS Processed Flag
	CxbOrsMsgTyp     float64   `gorm:"type:double precision"`              // ORS Message Type
	CxbAckTm         time.Time `gorm:"type:timestamp"`                     // Acknowledge Time
	CxbXchngRmrks    string    `gorm:"type:text"`                          // Exchange Remarks
	CxbExOrdrTyp     string    `gorm:"type:varchar(20)"`                   // Exchange Order Type
	CxbXchngCncldQty int64     `gorm:"type:bigint"`                        // Exchange Cancelled Quantity
	CxbSplFlag       string    `gorm:"type:varchar(10)"`                   // Special Flag
	CxbJiffy         float64   `gorm:"type:double precision"`              // Jiffy
	CxbMrktTyp       string    `gorm:"type:varchar(20)"`                   // Market Type
	CxbStreamNo      int64     `gorm:"type:bigint"`                        // Stream Number
	CxbSpnFlg        string    `gorm:"type:varchar(10)"`                   // SPN Flag
	CxbIp            string    `gorm:"type:varchar(50)"`                   // IP Address
	CxbInitSltpRt    float64   `gorm:"type:double precision"`              // Initial Stop Loss Trigger Rate
	CxbInitLmtRt     float64   `gorm:"type:double precision"`              // Initial Limit Rate
	CxbLtpRt         float64   `gorm:"type:double precision"`              // Last Traded Price Rate
	CxbIncrmntPrc    float64   `gorm:"type:double precision"`              // Incremental Price
	CxbTrailAmt      float64   `gorm:"type:double precision"`              // Trail Amount
	CxbLmtOffset     float64   `gorm:"type:double precision"`              // Limit Offset
	CxbTrlUpdCondVal float64   `gorm:"type:double precision"`              // Trail Update Condition Value
	CxbPrcimpvFlg    string    `gorm:"type:varchar(10)"`                   // Price Improvement Flag
	CxbClntOrdId     string    `gorm:"type:varchar(255)"`                  // Client Order ID
	CxbExgRespSeq    int64     `gorm:"type:bigint"`                        // Exchange Response Sequence
	CxbSssnId        int64     `gorm:"type:bigint"`                        // Session ID
	CxbAppmsgId      string    `gorm:"type:varchar(255)"`                  // App Message ID
}

func init() {
	DatabaseConnection()
}

func DatabaseConnection() {
	// Load the database configuration
	dbConfig, err := config.LoadConfig("dbconfig.ini")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to the database: %v", err)
	}

	log.Println("Database connection successful")
}
