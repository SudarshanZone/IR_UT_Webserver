package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/krishnakashyap0704/microservices/fnoSquareOff/generated"
	"github.com/krishnakashyap0704/microservices/fnoSquareOff/internal/database"
	logger "github.com/krishnakashyap0704/microservices/fnoSquareOff/utils"
)

type server struct {
	pb.UnimplementedSquareOffServiceServer
	db *sql.DB
}

var OrderSequence int64 = 0
var ORDR_FLW string

func (s *server) SquareOffOrder(ctx context.Context, req *pb.SquareOffRequest) (*pb.SquareOffResponse, error) {

	logger.InfoLogger.Println("Starting SquareOffOrder transaction")

	tx, err := s.db.Begin()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to start transaction: %v", err)
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			logger.ErrorLogger.Printf("Rolling back transaction due to error: %v", err)
			tx.Rollback()
		}
	}()

	//Pipe ID and Formated Date
	today := time.Now().Format("2006-01-02 15:04:05")
	pipeID := 99
	logger.InfoLogger.Printf("PipeID: %d", pipeID)

	//Other Variables For FXB/FOD
	LMT_MRKT_SL_FLG := "M"
	DSCLSD_QTY := 0
	LMT_RT := 0
	STP_LSS_TGR := 0
	MDFCTN_CNTR := 1
	ORDR_TYPE := "I"
	SPRD_ORD_IND := "*"
	RQST_TYP := "N"
	QUOTE := 0
	PLCD_STTS := "Q"
	EX_ORDR_TYP := "O"
	SPL_FLAG := "C"
	MRKT_TYP := "N"
	SPN_FLG := "S"
	INIT_SLTP_RT := 0
	INIT_LMT_RT := 0
	LTP_RT := 0
	TRAIL_AMT := 0
	LMT_OFFSET := 0
	PRCIMPV_FLG := "N"
	CLNT_CTGRY := 1
	CTCL_ID := "1111111111111"
	ORDR_STTS := "Q"
	EXEC_QTY := 0
	CNCL_QTY := 0
	CUST_TYP := "RI"
	SQROFF_TM := 0
	EXPRD_QTY := 0
	CLICK_FLG := "N"
	PRO_CLI_IND := "C"

	// Log each variable
	logger.InfoLogger.Printf("LMT_MRKT_SL_FLG: %s", LMT_MRKT_SL_FLG)
	logger.InfoLogger.Printf("DSCLSD_QTY: %d", DSCLSD_QTY)
	logger.InfoLogger.Printf("LMT_RT: %d", LMT_RT)
	logger.InfoLogger.Printf("STP_LSS_TGR: %d", STP_LSS_TGR)
	logger.InfoLogger.Printf("MDFCTN_CNTR: %d", MDFCTN_CNTR)
	logger.InfoLogger.Printf("ORDR_TYPE: %s", ORDR_TYPE)
	logger.InfoLogger.Printf("SPRD_ORD_IND: %s", SPRD_ORD_IND)
	logger.InfoLogger.Printf("RQST_TYP: %s", RQST_TYP)
	logger.InfoLogger.Printf("QUOTE: %d", QUOTE)
	logger.InfoLogger.Printf("PLCD_STTS: %s", PLCD_STTS)
	logger.InfoLogger.Printf("EX_ORDR_TYP: %s", EX_ORDR_TYP)
	logger.InfoLogger.Printf("SPL_FLAG: %s", SPL_FLAG)
	logger.InfoLogger.Printf("MRKT_TYP: %s", MRKT_TYP)
	logger.InfoLogger.Printf("SPN_FLG: %s", SPN_FLG)
	logger.InfoLogger.Printf("INIT_SLTP_RT: %d", INIT_SLTP_RT)
	logger.InfoLogger.Printf("INIT_LMT_RT: %d", INIT_LMT_RT)
	logger.InfoLogger.Printf("LTP_RT: %d", LTP_RT)
	logger.InfoLogger.Printf("TRAIL_AMT: %d", TRAIL_AMT)
	logger.InfoLogger.Printf("LMT_OFFSET: %d", LMT_OFFSET)
	logger.InfoLogger.Printf("PRCIMPV_FLG: %s", PRCIMPV_FLG)
	logger.InfoLogger.Printf("CLNT_CTGRY: %d", CLNT_CTGRY)
	logger.InfoLogger.Printf("CTCL_ID: %s", CTCL_ID)
	logger.InfoLogger.Printf("ORDR_STTS: %s", ORDR_STTS)
	logger.InfoLogger.Printf("EXEC_QTY: %d", EXEC_QTY)
	logger.InfoLogger.Printf("CNCL_QTY: %d", CNCL_QTY)
	logger.InfoLogger.Printf("CUST_TYP: %s", CUST_TYP)
	logger.InfoLogger.Printf("SQROFF_TM: %d", SQROFF_TM)
	logger.InfoLogger.Printf("EXPRD_QTY: %d", EXPRD_QTY)
	logger.InfoLogger.Printf("CLICK_FLG: %s", CLICK_FLG)
	logger.InfoLogger.Printf("PRO_CLI_IND: %s", PRO_CLI_IND)

	//Switch case for FNO Commidity and Equity
	switch req.FFO_PRDCT_TYP {
	case "F":
		logger.InfoLogger.Printf("Future Processing Started")
		if len(req.FcpDetails) > 0 {
			for _, fno := range req.FcpDetails {
				logger.DebugLogger.Printf("Processing FNO: %+v", fno)

				if err := validateData(fno); err != nil {
					logger.ErrorLogger.Printf("Validation failed: %v", err)
					return &pb.SquareOffResponse{
						Status:  "false",
						Message: fmt.Sprintf("validation failed: %v", err),
					}, nil
				}

				token, err := GetTokenNo(s.db, fno)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to get token number: %v", err)
					return nil, fmt.Errorf("failed to get token number: %v", err)
				}
				Token_No := token
				logger.InfoLogger.Printf("Token number retrieved: %d", Token_No)

				//Order Reference
				OrderReference, err := GenerateOrderReference(s.db, pipeID)
				if err != nil {
					logger.ErrorLogger.Printf("failed to generate order reference: %v", err)
					return nil, fmt.Errorf("failed to generate order reference: %v", err)
				}
				logger.InfoLogger.Printf("Generated Order Reference: %s", OrderReference)

				//Order Sequence
				OrderSequence += 1
				CurrOrderSeq := OrderSequence
				logger.DebugLogger.Printf("Current Order Sequence: %d", CurrOrderSeq)

				// Order Flow
				if fno.FCP_OPNPSTN_FLW == "B" {
					ORDR_FLW = "S"
				} else if fno.FCP_OPNPSTN_FLW == "S" {
					ORDR_FLW = "B"
				}
				logger.DebugLogger.Printf("Order Flow set to: %s", ORDR_FLW)

				logger.InfoLogger.Printf("Insertion Processing on FXB_FO_XCHNG_BOOK")
				_, err = tx.Exec(
					`INSERT INTO FXB_FO_XCHNG_BOOK (
					FXB_XCHNG_CD, FXB_ORDR_RFRNC, FXB_PIPE_ID, FXB_MOD_TRD_DT, FXB_ORDR_SQNC,
					FXB_LMT_MRKT_SL_FLG, FXB_DSCLSD_QTY, FXB_ORDR_TOT_QTY, FXB_LMT_RT, FXB_STP_LSS_TGR,
					FXB_MDFCTN_CNTR, FXB_ORDR_VALID_DT, FXB_ORDR_TYPE, FXB_SPRD_ORD_IND, FXB_RQST_TYP,
					FXB_QUOTE, FXB_RQST_TM, FXB_FRWD_TM, FXB_PLCD_STTS, FXB_EX_ORDR_TYP, FXB_SPL_FLAG, 
					FXB_MRKT_TYP, FXB_SPN_FLG, FXB_INIT_SLTP_RT, FXB_INIT_LMT_RT, FXB_LTP_RT, 
					FXB_TRAIL_AMT, FXB_LMT_OFFSET, FXB_PRCIMPV_FLG, FXB_INSRT_DT
					) VALUES (
					 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
					 $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
					)
				`,
					fno.FCP_XCHNG_CD, OrderReference, pipeID, today, CurrOrderSeq,
					LMT_MRKT_SL_FLG, DSCLSD_QTY, fno.FCP_OPNPSTN_QTY, LMT_RT, STP_LSS_TGR,
					MDFCTN_CNTR, today, ORDR_TYPE, SPRD_ORD_IND, RQST_TYP,
					QUOTE, today, today, PLCD_STTS, EX_ORDR_TYP, SPL_FLAG,
					MRKT_TYP, SPN_FLG, INIT_SLTP_RT, INIT_LMT_RT, LTP_RT,
					TRAIL_AMT, LMT_OFFSET, PRCIMPV_FLG, today,
				)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to insert order into FXB_FO_XCHNG_BOOK: %v", err)
					tx.Rollback()
					return nil, fmt.Errorf("failed to insert order into FXB_FO_XCHNG_BOOK: %v", err)
				}
				logger.InfoLogger.Println("Inserted order into FXB_FO_XCHNG_BOOK")

				logger.InfoLogger.Printf("Insertion Processing on FOD_FO_ORDR_DTLS")
				{
					_, err := tx.Exec(
						`INSERT INTO FOD_FO_ORDR_DTLS (
						FOD_CLM_MTCH_ACCNT, FOD_CLNT_CTGRY, FOD_ORDR_RFRNC, FOD_PIPE_ID, FOD_XCHNG_CD, FOD_PRDCT_TYP,
						FOD_INDSTK, FOD_UNDRLYNG, FOD_EXPRY_DT, FOD_EXER_TYP, FOD_OPT_TYP, FOD_STRK_PRC,
						FOD_ORDR_FLW, FOD_LMT_MRKT_SL_FLG, FOD_DSCLSD_QTY, FOD_ORDR_TOT_QTY, FOD_LMT_RT, FOD_STP_LSS_TGR,
						FOD_ORDR_TYPE, FOD_ORDR_VALID_DT, FOD_TRD_DT, FOD_ORDR_STTS, FOD_EXEC_QTY, FOD_CNCL_QTY,
						FOD_EXPRD_QTY, FOD_MDFCTN_CNTR, FOD_CTCL_ID, FOD_SPL_FLAG, FOD_PRO_CLI_IND, FOD_1CLICK_FLG,
						FOD_SQROFF_TM,  FOD_UCC_CD, FOD_TOKEN_NO,  FOD_CUST_TYP, FOD_EXGUCC_CD
						) VALUES (
						$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
						$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
						$21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
						$31, $32, $33, $34, $35
					)`,
						fno.FCP_CLM_MTCH_ACCNT, CLNT_CTGRY, OrderReference, pipeID, fno.FCP_XCHNG_CD, fno.FCP_PRDCT_TYP,
						fno.FCP_INDSTK, fno.FCP_UNDRLYNG, fno.FCP_EXPRY_DT, fno.FCP_EXER_TYP, fno.FCP_OPT_TYP, fno.FCP_STRK_PRC,
						ORDR_FLW, LMT_MRKT_SL_FLG, DSCLSD_QTY, fno.FCP_OPNPSTN_QTY, LMT_RT, STP_LSS_TGR,
						ORDR_TYPE, today, today, ORDR_STTS, EXEC_QTY, CNCL_QTY,
						EXPRD_QTY, MDFCTN_CNTR, CTCL_ID, SPL_FLAG, PRO_CLI_IND, CLICK_FLG,
						SQROFF_TM, fno.FCP_UCC_CD, Token_No, CUST_TYP, fno.FCP_UCC_CD,
					)

					if err != nil {
						logger.ErrorLogger.Printf("Failed to insert order into FOD_FO_ORDR_DTLS: %v", err)
						tx.Rollback()
						return nil, fmt.Errorf("failed to insert order into FOD_FO_ORDR_DTLS: %v", err)
					}
					logger.InfoLogger.Printf("Inserted order into FOD_FO_ORDR_DTLS")
				}

				logger.InfoLogger.Printf("Insertion Processing on FCP_FO_SPN_CNTRCT_PSTN")
				{
					_, err := tx.Exec(
						`INSERT INTO FCP_FO_SPN_CNTRCT_PSTN (
							FCP_CLM_MTCH_ACCNT, FCP_XCHNG_CD, FCP_PRDCT_TYP, FCP_INDSTK, FCP_UNDRLYNG,
							FCP_EXPRY_DT, FCP_EXER_TYP, FCP_STRK_PRC, FCP_OPT_TYP, FCP_IBUY_QTY,
							FCP_IBUY_ORD_VAL, FCP_ISELL_QTY, FCP_ISELL_ORD_VAL, FCP_EXBUY_QTY, FCP_EXBUY_ORD_VAL,
							FCP_EXSELL_QTY, FCP_EXSELL_ORD_VAL, FCP_BUY_EXCTD_QTY, FCP_SELL_EXCTD_QTY, FCP_OPNPSTN_FLW,
							FCP_OPNPSTN_QTY, FCP_OPNPSTN_VAL, FCP_OPT_PREMIUM, FCP_MTM_OPN_VAL, 
							FCP_TRG_PRC, FCP_MIN_TRG_PRC, FCP_AVG_PRC, FCP_UCC_CD
						) VALUES (
							$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
							$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
							$21, $22, $23, $24, $25, $26, $27, $28
						)`,
						fno.FCP_CLM_MTCH_ACCNT, fno.FCP_XCHNG_CD, fno.FCP_PRDCT_TYP, fno.FCP_INDSTK, fno.FCP_UNDRLYNG,
						fno.FCP_EXPRY_DT, fno.FCP_EXER_TYP, fno.FCP_STRK_PRC, fno.FCP_OPT_TYP, fno.FCP_IBUY_QTY,
						fno.FCP_IBUY_ORD_VAL, -1*fno.FCP_OPNPSTN_QTY, -1*fno.FCP_OPNPSTN_VAL, fno.FCP_EXBUY_QTY, fno.FCP_EXBUY_ORD_VAL,
						fno.FCP_EXSELL_QTY, fno.FCP_EXSELL_ORD_VAL, fno.FCP_BUY_EXCTD_QTY, fno.FCP_SELL_EXCTD_QTY, fno.FCP_OPNPSTN_FLW,
						fno.FCP_OPNPSTN_QTY, fno.FCP_OPNPSTN_VAL, fno.FCP_OPT_PREMIUM, fno.FCP_MTM_OPN_VAL,
						fno.FCP_TRG_PRC, fno.FCP_MIN_TRG_PRC, fno.FCP_AVG_PRC, fno.FCP_UCC_CD,
					)

					if err != nil {
						logger.ErrorLogger.Printf("Failed to insert order into FCP_FO_SPN_CNTRCT_PSTN: %v", err)
						tx.Rollback()
						return nil, fmt.Errorf("failed to insert order into FCP_FO_SPN_CNTRCT_PSTN: %v", err)
					}
					logger.InfoLogger.Printf("Inserted order into FCP_FO_SPN_CNTRCT_PSTN")
				}
			}
		}
		logger.InfoLogger.Printf("Future Processing Completed")

	case "O":
		logger.InfoLogger.Printf("Options Processing Started")
		if len(req.FcpDetails) > 0 {
			for _, fno := range req.FcpDetails {
				logger.DebugLogger.Printf("Processing FNO: %+v", fno)
				if err := validateData(fno); err != nil {
					logger.ErrorLogger.Printf("Validation failed: %v", err)
					return &pb.SquareOffResponse{
						Status:  "false",
						Message: fmt.Sprintf("validation failed: %v", err),
					}, nil
				}

				token, err := GetTokenNo(s.db, fno)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to get token number: %v", err)
					return nil, fmt.Errorf("failed to get token number: %v", err)
				}
				Token_No := token
				logger.InfoLogger.Printf("Token number retrieved: %d", Token_No)

				//Order Reference
				OrderReference, err := GenerateOrderReference(s.db, pipeID)
				if err != nil {
					logger.ErrorLogger.Printf("failed to generate order reference: %v", err)
					return nil, fmt.Errorf("failed to generate order reference: %v", err)
				}
				logger.InfoLogger.Printf("Generated Order Reference: %s", OrderReference)

				//Order Sequence
				OrderSequence += 1
				CurrOrderSeq := OrderSequence
				logger.DebugLogger.Printf("Current Order Sequence: %d", CurrOrderSeq)

				logger.InfoLogger.Printf("Insertion Processing on FXB_FO_XCHNG_BOOK")
				_, err = tx.Exec(
					`INSERT INTO FXB_FO_XCHNG_BOOK (
					FXB_XCHNG_CD, FXB_ORDR_RFRNC, FXB_PIPE_ID, FXB_MOD_TRD_DT, FXB_ORDR_SQNC,
					FXB_LMT_MRKT_SL_FLG, FXB_DSCLSD_QTY, FXB_ORDR_TOT_QTY, FXB_LMT_RT, FXB_STP_LSS_TGR,
					FXB_MDFCTN_CNTR, FXB_ORDR_VALID_DT, FXB_ORDR_TYPE, FXB_SPRD_ORD_IND, FXB_RQST_TYP,
					FXB_QUOTE, FXB_RQST_TM, FXB_FRWD_TM, FXB_PLCD_STTS, FXB_EX_ORDR_TYP, FXB_SPL_FLAG, 
					FXB_MRKT_TYP, FXB_SPN_FLG, FXB_INIT_SLTP_RT, FXB_INIT_LMT_RT, FXB_LTP_RT, 
					FXB_TRAIL_AMT, FXB_LMT_OFFSET, FXB_PRCIMPV_FLG, FXB_INSRT_DT
					) VALUES (
					 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
					 $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
					)
				`,
					fno.FCP_XCHNG_CD, OrderReference, pipeID, today, CurrOrderSeq,
					LMT_MRKT_SL_FLG, DSCLSD_QTY, fno.FCP_OPNPSTN_QTY, LMT_RT, STP_LSS_TGR,
					MDFCTN_CNTR, today, ORDR_TYPE, SPRD_ORD_IND, RQST_TYP,
					QUOTE, today, today, PLCD_STTS, EX_ORDR_TYP, SPL_FLAG,
					MRKT_TYP, SPN_FLG, INIT_SLTP_RT, INIT_LMT_RT, LTP_RT,
					TRAIL_AMT, LMT_OFFSET, PRCIMPV_FLG, today,
				)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to insert order into FXB_FO_XCHNG_BOOK: %v", err)
					tx.Rollback()
					return nil, fmt.Errorf("failed to insert order into FXB_FO_XCHNG_BOOK: %v", err)
				}
				logger.InfoLogger.Println("Inserted order into FXB_FO_XCHNG_BOOK")

				logger.InfoLogger.Printf("Insertion Processing on FOD_FO_ORDR_DTLS")
				{
					_, err := tx.Exec(
						`INSERT INTO FOD_FO_ORDR_DTLS (
						FOD_CLM_MTCH_ACCNT, FOD_CLNT_CTGRY, FOD_ORDR_RFRNC, FOD_PIPE_ID, FOD_XCHNG_CD, FOD_PRDCT_TYP,
						FOD_INDSTK, FOD_UNDRLYNG, FOD_EXPRY_DT, FOD_EXER_TYP, FOD_OPT_TYP, FOD_STRK_PRC,
						FOD_ORDR_FLW, FOD_LMT_MRKT_SL_FLG, FOD_DSCLSD_QTY, FOD_ORDR_TOT_QTY, FOD_LMT_RT, FOD_STP_LSS_TGR,
						FOD_ORDR_TYPE, FOD_ORDR_VALID_DT, FOD_TRD_DT, FOD_ORDR_STTS, FOD_EXEC_QTY, FOD_CNCL_QTY,
						FOD_EXPRD_QTY, FOD_MDFCTN_CNTR, FOD_CTCL_ID, FOD_SPL_FLAG, FOD_PRO_CLI_IND, FOD_1CLICK_FLG,
						FOD_SQROFF_TM,  FOD_UCC_CD, FOD_TOKEN_NO,  FOD_CUST_TYP, FOD_EXGUCC_CD
						) VALUES (
						$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
						$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
						$21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
						$31, $32, $33, $34, $35
					)`,
						fno.FCP_CLM_MTCH_ACCNT, CLNT_CTGRY, OrderReference, pipeID, fno.FCP_XCHNG_CD, fno.FCP_PRDCT_TYP,
						fno.FCP_INDSTK, fno.FCP_UNDRLYNG, fno.FCP_EXPRY_DT, fno.FCP_EXER_TYP, fno.FCP_OPT_TYP, fno.FCP_STRK_PRC,
						ORDR_FLW, LMT_MRKT_SL_FLG, DSCLSD_QTY, fno.FCP_OPNPSTN_QTY, LMT_RT, STP_LSS_TGR,
						ORDR_TYPE, today, today, ORDR_STTS, EXEC_QTY, CNCL_QTY,
						EXPRD_QTY, MDFCTN_CNTR, CTCL_ID, SPL_FLAG, PRO_CLI_IND, CLICK_FLG,
						SQROFF_TM, fno.FCP_UCC_CD, Token_No, CUST_TYP, fno.FCP_UCC_CD,
					)

					if err != nil {
						logger.ErrorLogger.Printf("Failed to insert order into FOD_FO_ORDR_DTLS: %v", err)
						tx.Rollback()
						return nil, fmt.Errorf("failed to insert order into FOD_FO_ORDR_DTLS: %v", err)
					}
					logger.InfoLogger.Println("Inserted order into FOD_FO_ORDR_DTLS")
				}

				logger.InfoLogger.Printf("Insertion Processing on FCP_FO_SPN_CNTRCT_PSTN")
				{
					_, err := tx.Exec(
						`INSERT INTO FCP_FO_SPN_CNTRCT_PSTN (
							FCP_CLM_MTCH_ACCNT, FCP_XCHNG_CD, FCP_PRDCT_TYP, FCP_INDSTK, FCP_UNDRLYNG,
							FCP_EXPRY_DT, FCP_EXER_TYP, FCP_STRK_PRC, FCP_OPT_TYP, FCP_IBUY_QTY,
							FCP_IBUY_ORD_VAL, FCP_ISELL_QTY, FCP_ISELL_ORD_VAL, FCP_EXBUY_QTY, FCP_EXBUY_ORD_VAL,
							FCP_EXSELL_QTY, FCP_EXSELL_ORD_VAL, FCP_BUY_EXCTD_QTY, FCP_SELL_EXCTD_QTY, FCP_OPNPSTN_FLW,
							FCP_OPNPSTN_QTY, FCP_OPNPSTN_VAL, FCP_OPT_PREMIUM, FCP_MTM_OPN_VAL, 
							FCP_TRG_PRC, FCP_MIN_TRG_PRC, FCP_AVG_PRC, FCP_UCC_CD
						) VALUES (
							$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
							$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
							$21, $22, $23, $24, $25, $26, $27, $28
						)`,
						fno.FCP_CLM_MTCH_ACCNT, fno.FCP_XCHNG_CD, fno.FCP_PRDCT_TYP, fno.FCP_INDSTK, fno.FCP_UNDRLYNG,
						fno.FCP_EXPRY_DT, fno.FCP_EXER_TYP, fno.FCP_STRK_PRC, fno.FCP_OPT_TYP, fno.FCP_IBUY_QTY,
						fno.FCP_IBUY_ORD_VAL, -1*fno.FCP_OPNPSTN_QTY, -1*fno.FCP_OPNPSTN_VAL, fno.FCP_EXBUY_QTY, fno.FCP_EXBUY_ORD_VAL,
						fno.FCP_EXSELL_QTY, fno.FCP_EXSELL_ORD_VAL, fno.FCP_BUY_EXCTD_QTY, fno.FCP_SELL_EXCTD_QTY, fno.FCP_OPNPSTN_FLW,
						fno.FCP_OPNPSTN_QTY, fno.FCP_OPNPSTN_VAL, fno.FCP_OPT_PREMIUM, fno.FCP_MTM_OPN_VAL,
						fno.FCP_TRG_PRC, fno.FCP_MIN_TRG_PRC, fno.FCP_AVG_PRC, fno.FCP_UCC_CD,
					)

					if err != nil {
						logger.ErrorLogger.Printf("Failed to insert order into FCP_FO_SPN_CNTRCT_PSTN: %v", err)
						tx.Rollback()
						return nil, fmt.Errorf("failed to insert order into FCP_FO_SPN_CNTRCT_PSTN: %v", err)
					}
					logger.InfoLogger.Printf("Inserted order into FCP_FO_SPN_CNTRCT_PSTN")
				}
			}
		}
		logger.InfoLogger.Printf("Options Processing Completed")

	default:
		return nil, fmt.Errorf("unsupported product type: %v", req.FFO_PRDCT_TYP)

	}

	if err := tx.Commit(); err != nil {
		logger.ErrorLogger.Printf("Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	logger.InfoLogger.Printf("Transaction committed successfully")
	return &pb.SquareOffResponse{
		Status:  "true",
		Message: "order squared off and inserted successfully",
	}, nil

}

var orderCounter int = 0

// Validation for incoming data
func validateData(fno *pb.FnoData) error {

	if _, ok := interface{}(fno.FCP_CLM_MTCH_ACCNT).(string); !ok {
		return fmt.Errorf("FCP_CLM_MTCH_ACCNT must be a string")
	}

	if _, ok := interface{}(fno.FCP_XCHNG_CD).(string); !ok {
		return fmt.Errorf("FCP_XCHNG_CD must be a string")
	}

	if _, ok := interface{}(fno.FCP_PRDCT_TYP).(string); !ok {
		return fmt.Errorf("FCP_PRDCT_TYP must be a string")
	}

	if _, ok := interface{}(fno.FCP_INDSTK).(string); !ok {
		return fmt.Errorf("FCP_INDSTK must be a string")
	}

	if _, ok := interface{}(fno.FCP_UNDRLYNG).(string); !ok {
		return fmt.Errorf("FCP_UNDRLYNG must be a string")
	}

	if _, ok := interface{}(fno.FCP_EXPRY_DT).(string); !ok {
		return fmt.Errorf("FCP_EXPRY_DT must be a string")
	}

	if _, ok := interface{}(fno.FCP_EXER_TYP).(string); !ok {
		return fmt.Errorf("FCP_EXER_TYP must be a string")
	}

	if _, ok := interface{}(fno.FCP_OPT_TYP).(string); !ok {
		return fmt.Errorf("FCP_OPT_TYP must be a string")
	}

	if _, ok := interface{}(fno.FCP_UCC_CD).(string); !ok {
		return fmt.Errorf("FCP_UCC_CD must be a string")
	}

	if _, ok := interface{}(fno.FCP_OPNPSTN_FLW).(string); !ok {
		return fmt.Errorf("FCP_UCC_CD must be a string")
	}

	if _, ok := interface{}(fno.FCP_STRK_PRC).(int64); !ok {
		return fmt.Errorf("FCP_STRK_PRC must be a int64")
	}

	if _, ok := interface{}(fno.FCP_IBUY_QTY).(int64); !ok {
		return fmt.Errorf("FCP_IBUY_QTY must be a int64")
	}

	return nil
}

func GetTokenNo(db *sql.DB, req *pb.FnoData) (int32, error) {
	var tokenNo int32

	// Check if FCP_EXPRY_DT is empty
	if req.FCP_EXPRY_DT == "" {
		return 0, fmt.Errorf("expiry date is empty")
	}

	// Parse the timestamp to ensure correct format
	expiryDate, err := time.Parse(time.RFC3339, req.FCP_EXPRY_DT)
	if err != nil {
		return 0, fmt.Errorf("invalid expiry date format: %v", err)
	}

	// Format expiry date for SQL query
	formattedExpiryDate := expiryDate.Format("2006-01-02 15:04:05")

	query := `
        SELECT FTQ_TOKEN_NO 
        FROM FTQ_FO_TRD_QT 
        WHERE FTQ_STRK_PRC = $1 
        AND FTQ_EXER_TYP = $2 
        AND FTQ_EXPRY_DT = $3 
        AND FTQ_PRDCT_TYP = $4 
        AND FTQ_UNDRLYNG = $5
    `
	err = db.QueryRow(query, req.FCP_STRK_PRC, req.FCP_EXER_TYP, formattedExpiryDate, req.FCP_PRDCT_TYP, req.FCP_UNDRLYNG).Scan(&tokenNo)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no token found for the given criteria")
		}
		return 0, fmt.Errorf("query error: %v", err)
	}

	return tokenNo, nil
}

// GenerateOrderReference generates a new order reference based on the current date, pipe ID, and database counter.
func GenerateOrderReference(db *sql.DB, pipeID int) (string, error) {

	// Step 1: Get the current date in YYYYMMDD format
	currentDate := time.Now().Format("20060102")

	// Step 2: Add the Pipe ID (ensure it's a 2-digit number)
	pipeIDStr := fmt.Sprintf("%02d", pipeID)

	// Step 3: Query the database to get the max order reference for today
	query := `
		SELECT MAX(fxb_ordr_rfrnc)
		FROM fxb_fo_xchng_book
		WHERE fxb_ordr_rfrnc LIKE $1;
	`
	pattern := currentDate + "%"
	var maxOrderRef sql.NullString
	err := db.QueryRow(query, pattern).Scan(&maxOrderRef)
	if err != nil {
		return "", fmt.Errorf("failed to query max order reference: %v", err)
	}

	// Step 4: If there is a non-null result, extract the counter part and increment it
	if maxOrderRef.Valid {
		lastCounter := maxOrderRef.String[len(maxOrderRef.String)-8:] // Extract the last 8 digits
		fmt.Sscanf(lastCounter, "%d", &orderCounter)
		orderCounter += 1
	} else {
		// If no previous order reference, start with 1
		orderCounter += 1
	}

	// Step 5: Generate the last 8 digits by incrementing the counter
	counterStr := fmt.Sprintf("%08d", orderCounter)

	// Combine all parts to create the order reference number
	orderReference := currentDate + pipeIDStr + counterStr

	return orderReference, nil
}

// Main Connection
func main() {

	//Initialize the logger
	logger.InitLogger()

	// Initialize the database connection
	database.Init()
	db := database.GetDB()
	defer db.Close()

	// Ensure the database connection is closed on exit
	defer func() {
		if err := db.Close(); err != nil {
			logger.ErrorLogger.Printf("Failed to close the database connection: %v", err)
		} else {
			logger.InfoLogger.Println("Database connection closed successfully")
		}
	}()

	// Check if the db is nil
	if db == nil {
		logger.ErrorLogger.Printf("Database connection is nil")
	} else {
		logger.InfoLogger.Println("Database connection is successfully initialized")
	}

	// Test the database connection
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error querying database: %v", err)
		log.Printf("Error querying database: %v", err)
	}
	logger.InfoLogger.Println("Database version:", version)
	log.Println("Database version:", version)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.ErrorLogger.Fatalf("failed to listen: %v", err)
		log.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	squareOffService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterSquareOffServiceServer(s, squareOffService)
	reflection.Register(s)

	logger.InfoLogger.Println("gRPC server is running...")
	log.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		logger.ErrorLogger.Fatalf("Failed to serve gRPC server: %v", err)
		log.Printf("failed to serve: %v", err)
	}
}
