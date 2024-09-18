package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/krishnakashyap0704/microservices/equSquareOff/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Trd_Trd_Dtls struct {
	Trd_trd_ref             string
	Trd_clm_mtch_accnt      string
	Trd_xchng_cd            string
	Trd_stck_cd             string
	Trd_xchng_sgmnt_cd      string
	Trd_xchng_sgmnt_sttlmnt int64
	Trd_ordr_rfrnc          string
	Trd_trd_dt              time.Time
	Trd_trnsctn_typ         string
	Trd_trd_flw             string
	Trd_exctd_qty           int64
	Trd_exctd_rt            float64
	Trd_trd_vl              int64
	Trd_brkrg_vl            int64
	Trd_net_vl              int64
	Trd_amt_blckd           int64
	Trd_xchng_rfrnc         int64
	Trd_upld_mtch_flg       string
	Trd_buy_oblg_flg        string
	Trd_prtfl_flg           *string
	Trd_usr_id              string
	Trd_stmp_duty           float64
	Trd_trnx_chrg           float64
	Trd_sebi_chrg_val       float64
	Trd_stt                 int64
	Trd_cntrct_nmbr         string
	Trd_brkrg_flg           string
	Trd_srvc_tax            int64
	Trd_brkrg_typ           string
	Trd_cgst_amt            int64
	Trd_sgst_amt            int64
	Trd_ugst_amt            int64
	Trd_igst_amt            int64
	Trd_atm_upfront_amt     int64
	Trd_fixed_brkg          int64
	Trd_variable_brkg       int64
	Trd_brkrg_mdl           string
	Trd_csh_wthld_amt       *int64
	Trd_ins_dt              time.Time
}

type EPBEmPstnBook struct {
	EpbClmMtchAccnt      string         `gorm:"type:char(10);column:epb_clm_mtch_accnt;not null"`
	EpbXchngCd           string         `gorm:"type:char(3);column:epb_xchng_cd;not null"`
	EpbXchngSgmntCd      string         `gorm:"type:char(2);column:epb_xchng_sgmnt_cd;not null"`
	EpbXchngSgmntSttlmnt int            `gorm:"column:epb_xchng_sgmnt_sttlmnt;not null"`
	EpbStckCd            string         `gorm:"type:char(6);column:epb_stck_cd;not null"`
	EpbOrgnlPstnQty      int            `gorm:"column:epb_orgnl_pstn_qty;not null"`
	EpbRate              float64        `gorm:"column:epb_rate;not null"`
	EpbOrgnlAmtPayble    float64        `gorm:"type:numeric(18,2);column:epb_orgnl_amt_payble;not null"`
	EpbOrgnlMrgnAmt      float64        `gorm:"type:numeric(18,2);column:epb_orgnl_mrgn_amt;not null"`
	EpbSellQty           int            `gorm:"column:epb_sell_qty;not null"`
	EpbCvrOrdQty         int            `gorm:"column:epb_cvr_ord_qty;not null"`
	EpbNetMrgnAmt        float64        `gorm:"type:numeric(18,2);column:epb_net_mrgn_amt;not null"`
	EpbNetAmtPayble      float64        `gorm:"type:numeric(18,2);column:epb_net_amt_payble;not null"`
	EpbNetPstnQty        int            `gorm:"column:epb_net_pstn_qty;not null"`
	EpbCtdQty            int            `gorm:"column:epb_ctd_qty;not null"`
	EpbPstnStts          string         `gorm:"type:char(1);column:epb_pstn_stts;not null"`
	EpbLpcCalcStts       string         `gorm:"type:char(1);column:epb_lpc_calc_stts;not null"`
	EpbSqroffMode        string         `gorm:"type:char(1);column:epb_sqroff_mode;not null"`
	EpbPstnTrdDt         sql.NullString `gorm:"column:epb_pstn_trd_dt;not null"`
	EpbMtmPrcsFlg        string         `gorm:"type:char(1);column:epb_mtm_prcs_flg;not null"`
	EpbLastMdfcnDt       sql.NullString `gorm:"column:epb_last_mdfcn_dt;not null"`
	EpbInsDate           sql.NullString `gorm:"column:epb_ins_date;not null"`
	EpbCloseDate         sql.NullString `gorm:"column:epb_close_date"`
	EpbSysFailFlg        string         `gorm:"type:char(1);column:epb_sys_fail_flg;not null"`
	EpbLastPymntDt       sql.NullString `gorm:"column:epb_last_pymnt_dt"`
	EpbLpcCalcEndDt      sql.NullString `gorm:"column:epb_lpc_calc_end_dt"`
	EpbMtmCansq          string         `gorm:"type:char(1);column:epb_mtm_cansq"`
	EpbExpiryDt          sql.NullString `gorm:"column:epb_expiry_dt"`
	EpbMinMrgn           float64        `gorm:"column:epb_min_mrgn"`
	EpbMrgnDbcrPrcsFlg   string         `gorm:"type:char(1);column:epb_mrgn_dbcr_prcs_flg"`
	EpbDpId              string         `gorm:"type:char(8);column:epb_dp_id"`
	EpbDpClntId          string         `gorm:"type:char(8);column:epb_dp_clnt_id"`
	EpbPledgeStts        string         `gorm:"type:char(1);column:epb_pledge_stts"`
	EpbBtstNetMrgnAmt    float64        `gorm:"type:numeric(18,2);column:epb_btst_net_mrgn_amt"`
	EpbBtstMrgnBlckd     float64        `gorm:"type:numeric(18,2);column:epb_btst_mrgn_blckd"`
	EpbBtstMrgnDbcrFlg   string         `gorm:"type:char(1);column:epb_btst_mrgn_dbcr_flg"`
	EpbBtstSgmntCd       string         `gorm:"type:char(2);column:epb_btst_sgmnt_cd"`
	EpbBtstStlmnt        int            `gorm:"column:epb_btst_stlmnt"`
	EpbBtstCshBlckd      float64        `gorm:"type:numeric(18,2);column:epb_btst_csh_blckd"`
	EpbBtstSamBlckd      float64        `gorm:"type:numeric(18,2);column:epb_btst_sam_blckd"`
	EpbBtstCalcDt        sql.NullString `gorm:"column:epb_btst_calc_dt"`
	EpbDbcrCalcDt        sql.NullString `gorm:"column:epb_dbcr_calc_dt"`
	EpbNsdlRefNo         string         `gorm:"type:varchar(16);column:epb_nsdl_ref_no"`
	EpbMrgnWithheldFlg   string         `gorm:"type:char(2);column:epb_mrgn_withheld_flg"`
}

type Otp_trd_pstns struct {
	Otp_clm_mtch_acct       string  `db:"otp_clm_mtch_acct"`
	Otp_xchng_cd            string  `db:"otp_xchng_cd"`
	Otp_xchng_sgmnt_cd      string  `db:"otp_xchng_sgmnt_cd"`
	Otp_xchng_sgmnt_sttlmnt int32   `db:"otp_xchng_sgmnt_sttlmnt"`
	Otp_stck_cd             string  `db:"otp_stck_cd"`
	Otp_flw                 string  `db:"otp_flw"`
	Otp_qty                 int64   `db:"otp_qty"`
	Otp_cnvrt_dlvry_qty     int64   `db:"otp_cnvrt_dlvry_qty"`
	Otp_cvrd_qty            int64   `db:"otp_cvrd_qty"`
	Otp_rt                  float64 `db:"otp_rt"`
	Otp_mrgn_amt            float64 `db:"otp_mrgn_amt"`
	Otp_trd_val             float64 `db:"otp_trd_val"`
	Otp_rmrks               string  `db:"otp_rmrks"`
	Otp_xfer_mrgn_stts      string  `db:"otp_xfer_mrgn_stts"`
	Otp_sell_opn_prccsd     string  `db:"otp_sell_opn_prccsd"`
	Otp_buy_opn_prccsd      string  `db:"otp_buy_opn_prccsd"`
	Otp_mrgn_sqroff_mode    string  `db:"otp_mrgn_sqroff_mode"`
	Otp_em_trdsplt_prcs_flg string  `db:"otp_em_trdsplt_prcs_flg"`
	Otp_mtm_flg             string  `db:"otp_mtm_flg"`
	Otp_mtm_cansq           string  `db:"otp_mtm_cansq"`
	Otp_eos_can             string  `db:"otp_eos_can"`
	Otp_trgr_prc            float64 `db:"otp_trgr_prc"`
	Otp_16_trgr_prc         float64 `db:"otp_16_trgr_prc"`
	Otp_min_mrgn            float64 `db:"otp_min_mrgn"`
}

type Ord_Ordr_Dtls struct {
	Ord_clm_mtch_accnt      string     `json:"ord_clm_mtch_accnt"`
	Ord_ordr_rfrnc          string     `json:"ord_ordr_rfrnc"`
	Ord_xchng_cd            string     `json:"ord_xchng_cd"`
	Ord_stck_cd             string     `json:"ord_stck_cd"`
	Ord_xchng_sgmnt_cd      string     `json:"ord_xchng_sgmnt_cd"`
	Ord_xchng_sgmnt_sttlmnt int        `json:"ord_xchng_sgmnt_sttlmnt"`
	Ord_ordr_dt             time.Time  `json:"ord_ordr_dt"`
	Ord_ordr_flw            string     `json:"ord_ordr_flw"`
	Ord_prdct_typ           string     `json:"ord_prdct_typ,omitempty"`
	Ord_ordr_qty            int        `json:"ord_ordr_qty"`
	Ord_lmt_mrkt_flg        string     `json:"ord_lmt_mrkt_flg"`
	Ord_lmt_rt              float64    `json:"ord_lmt_rt,omitempty"`
	Ord_dsclsd_qty          int        `json:"ord_dsclsd_qty,omitempty"`
	Ord_stp_lss_tgr         float64    `json:"ord_stp_lss_tgr,omitempty"`
	Ord_ordr_stts           string     `json:"ord_ordr_stts"`
	Ord_trd_dt              time.Time  `json:"ord_trd_dt"`
	Ord_sub_brkr_tag        *string    `json:"ord_sub_brkr_tag,omitempty"`
	Ord_mdfctn_cntr         int        `json:"ord_mdfctn_cntr"`
	Ord_ack_nmbr            *string    `json:"ord_ack_nmbr,omitempty"`
	Ord_xchng_ack_old       *float64   `json:"ord_xchng_ack_old,omitempty"`
	Ord_exctd_qty           int        `json:"ord_exctd_qty,omitempty"`
	Ord_amt_blckd           float64    `json:"ord_amt_blckd,omitempty"`
	Ord_brkrg_val           *float64   `json:"ord_brkrg_val,omitempty"`
	Ord_dp_id               string     `json:"ord_dp_id,omitempty"`
	Ord_dp_clnt_id          string     `json:"ord_dp_clnt_id,omitempty"`
	Ord_phy_qty             *int       `json:"ord_phy_qty,omitempty"`
	Ord_isin_nmbr           *string    `json:"ord_isin_nmbr,omitempty"`
	Ord_nd_flg              *string    `json:"ord_nd_flg,omitempty"`
	Ord_msc_char            *string    `json:"ord_msc_char,omitempty"`
	Ord_msc_varchar         *string    `json:"ord_msc_varchar,omitempty"`
	Ord_msc_int             *float64   `json:"ord_msc_int,omitempty"`
	Ord_plcd_stts           string     `json:"ord_plcd_stts,omitempty"`
	Ord_qty_blckd           int        `json:"ord_qty_blckd,omitempty"`
	Ord_mrgn_prcntg         *float64   `json:"ord_mrgn_prcntg,omitempty"`
	Ord_ipo_flg             *string    `json:"ord_ipo_flg,omitempty"`
	Ord_lss_amt_blckd       float64    `json:"ord_lss_amt_blckd,omitempty"`
	Ord_lss_qty             int64      `json:"ord_lss_qty,omitempty"`
	Ord_mtm_flg             *string    `json:"ord_mtm_flg,omitempty"`
	Ord_sq_flg              string     `json:"ord_sq_flg,omitempty"`
	Ord_schm_id             *string    `json:"ord_schm_id,omitempty"`
	Ord_pipe_id             string     `json:"ord_pipe_id,omitempty"`
	Ord_prtctn_rt           float64    `json:"ord_prtctn_rt,omitempty"`
	Ord_sl_trg_flg          string     `json:"ord_sl_trg_flg,omitempty"`
	Ord_xchng_usr_id        *int       `json:"ord_xchng_usr_id,omitempty"`
	Ord_btst_sttlmnt_nmbr   *int       `json:"ord_btst_sttlmnt_nmbr,omitempty"`
	Ord_btst_sgmnt_cd       *string    `json:"ord_btst_sgmnt_cd,omitempty"`
	Ord_channel             string     `json:"ord_channel,omitempty"`
	Ord_bp_id               string     `json:"ord_bp_id,omitempty"`
	Ord_sltp_ordr_rfrnc     *string    `json:"ord_sltp_ordr_rfrnc,omitempty"`
	Ord_ctcl_id             string     `json:"ord_ctcl_id,omitempty"`
	Ord_usr_id              string     `json:"ord_usr_id,omitempty"`
	Ord_cnt_id              *int       `json:"ord_cnt_id,omitempty"`
	Ord_em_settlmnt_nmbr    int        `json:"ord_em_settlmnt_nmbr,omitempty"`
	Ord_mrgn_sqroff_mode    string     `json:"ord_mrgn_sqroff_mode,omitempty"`
	Ord_cncl_qty            *int       `json:"ord_cncl_qty,omitempty"`
	Ord_ordr_typ            string     `json:"ord_ordr_typ,omitempty"`
	Ord_valid_dt            *time.Time `json:"ord_valid_dt,omitempty"`
	Ord_cal_flg             string     `json:"ord_cal_flg,omitempty"`
	Ord_xchng_ack           *string    `json:"ord_xchng_ack,omitempty"`
	Ord_em_rollovr_flg      string     `json:"ord_em_rollovr_flg,omitempty"`
	Ord_trd_val             *float64   `json:"ord_trd_val,omitempty"`
	Ord_trd_cntrct_nmbr     *string    `json:"ord_trd_cntrct_nmbr,omitempty"`
	Ord_avg_exctd_rt        *float64   `json:"ord_avg_exctd_rt,omitempty"`
	Ord_prc_imp_flg         *string    `json:"ord_prc_imp_flg,omitempty"`
	Ord_mbc_flg             *string    `json:"ord_mbc_flg,omitempty"`
	Ord_trl_amt             *float64   `json:"ord_trl_amt,omitempty"`
	Ord_lmt_offst           *float64   `json:"ord_lmt_offst,omitempty"`
	Ord_source_flg          *string    `json:"ord_source_flg,omitempty"`
	Ord_pan_no              string     `json:"ord_pan_no,omitempty"`
	Ord_atm_payout_stts     *string    `json:"ord_atm_payout_stts,omitempty"`
	Ord_esp_cd              *string    `json:"ord_esp_cd,omitempty"`
	Ord_remarks             *string    `json:"ord_remarks,omitempty"`
	Ord_wthld_amt_stts      *string    `json:"ord_wthld_amt_stts,omitempty"`
	Ord_pstn_xchng_cd       *string    `json:"ord_pstn_xchng_cd,omitempty"`
	Ord_interop_ord_flg     string     `json:"ord_interop_ord_flg,omitempty"`
	Ord_settlement_period   int        `json:"ord_settlement_period,omitempty"`
	Ord_algo_id             *string    `json:"ord_algo_id,omitempty"`
	Ord_bundle_name         *string    `json:"ord_bundle_name,omitempty"`
	Ord_prt_flg             *string    `json:"ord_prt_flg,omitempty"`
	Ord_src_tag             *string    `json:"ord_src_tag,omitempty"`
	Ord_rls_amt             *float64   `json:"ord_rls_amt,omitempty"`
	Ord_rls_date            *time.Time `json:"ord_rls_date,omitempty"`
	Ord_mtf_unplg_sqroff    string     `json:"ord_mtf_unplg_sqroff,omitempty"`
	Ord_n_ordr_qty          *int       `json:"ord_n_ordr_qty,omitempty"`
	Ord_ack_date            *time.Time `json:"ord_ack_date,omitempty"`
	Ord_last_activity_ref   *string    `json:"ord_last_activity_ref,omitempty"`
	Ord_clm_clnt_cd         string     `json:"ord_clm_clnt_cd,omitempty"`
}

type Clm_clnt_mstr struct {
	Clm_mtch_accnt string
	Clm_bp_id      string
	Clm_clnt_cd    string
}

type Esp_em_systm_prmtr struct {
	Esp_dp_id      string
	Esp_dp_clnt_id string
}

type ICD_INFO_CLNT_DTLS struct {
	Icd_Pan_No string
}

func init() {
	DatabaseConnection()
}

func DatabaseConnection() {

	dbConfig, err := config.LoadConfig("dbconfig.ini")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	fmt.Println("Connecting to the database...")
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	} else {
		fmt.Println("Database connection successful.")
	}
}
