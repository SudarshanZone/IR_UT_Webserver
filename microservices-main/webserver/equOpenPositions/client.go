package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/krishnakashyap0704/microservices/equOpenPositions/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		log.Println("Setting CORS headers")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			log.Println("CORS Preflight request received, aborting with status 200")
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Initialize gRPC connections
	log.Println("Attempting to connect to gRPC servers")

	// gRPC connection for EPBService
	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to EPB gRPC server: %v", err)
	}
	// defer func() {
	// 	log.Println("Closing EPB gRPC connection")
	// 	epbConn.Close()
	// }()
	// epbClient := pb.NewEPBServiceClient(epbConn)

	// // gRPC connection for OTPService
	// otpConn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials())) // Assuming OTPService runs on a different port
	// if err != nil {
	// 	log.Fatalf("Could not connect to OTP gRPC server: %v", err)
	// }
	// defer func() {
	// 	log.Println("Closing OTP gRPC connection")
	// 	otpConn.Close()
	// }()
	epbClient := pb.NewEPBServiceClient(conn)
	otpClient := pb.NewOTPServiceClient(conn)

	// Define the GET route for EPB positions
	router.GET("/EPBService/:epb_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("epb_clm_mtch_accnt")
		log.Printf("Received request for EPB account: %s", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Calling gRPC GetEquityPosition service for EPB")
		res, err := epbClient.GetEquityPosition(ctxWithTimeout, &pb.EquityPositionRequest{EpbClmMtchAccnt: accnt})
		if err != nil {
			log.Printf("Error fetching EPB position from gRPC service: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		if res.Equity != nil {
			log.Println("Processing EPB positions received from gRPC service")
			positions := []gin.H{}

			for _, p := range res.Equity {
				if p != nil {
					log.Printf("Processing EPB position for stock: %s", p.EpbStckCd)
					positions = append(positions, gin.H{
						"epb_clm_mtch_accnt":      checkNullString(p.EpbClmMtchAccnt),
						"epb_xchng_cd":            checkNullString(p.EpbXchngCd),
						"epb_xchng_sgmnt_cd":      checkNullString(p.EpbXchngSgmntCd),
						"epb_xchng_sgmnt_sttlmnt": p.EpbXchngSgmntSttlmnt,
						"epb_stck_cd":             checkNullString(p.EpbStckCd),
						"epb_orgnl_pstn_qty":      p.EpbOrgnlPstnQty,
						"epb_rate":                p.EpbRate,
						"epb_orgnl_amt_payble":    p.EpbOrgnlAmtPayble,
						"epb_orgnl_mrgn_amt":      p.EpbOrgnlMrgnAmt,
						"epb_sell_qty":            p.EpbSellQty,
						"epb_cvr_ord_qty":         p.EpbCvrOrdQty,
						"epb_net_mrgn_amt":        p.EpbNetMrgnAmt,
						"epb_net_amt_payble":      p.EpbNetAmtPayble,
						"epb_net_pstn_qty":        p.EpbNetPstnQty,
						"epb_ctd_qty":             p.EpbCtdQty,
						"epb_pstn_stts":           checkNullString(p.EpbPstnStts),
						"epb_lpc_calc_stts":       checkNullString(p.EpbLpcCalcStts),
						"epb_sqroff_mode":         checkNullString(p.EpbSqroffMode),
						"epb_pstn_trd_dt":         checkNullString(p.EpbPstnTrdDt),
						"epb_mtm_prcs_flg":        checkNullString(p.EpbMtmPrcsFlg),
						"epb_last_mdfcn_dt":       checkNullString(p.EpbLastMdfcnDt),
						"epb_ins_date":            checkNullString(p.EpbInsDate),
						"epb_close_date":          checkNullString(p.EpbCloseDate),
						"epb_sys_fail_flg":        checkNullString(p.EpbSysFailFlg),
						"epb_last_pymnt_dt":       checkNullString(p.EpbLastPymntDt),
						"epb_lpc_calc_end_dt":     checkNullString(p.EpbLpcCalcEndDt),
						"epb_mtm_cansq":           checkNullString(p.EpbMtmCansq),
						"epb_expiry_dt":           checkNullString(p.EpbExpiryDt),
						"epb_min_mrgn":            p.EpbMinMrgn,
						"epb_mrgn_dbcr_prcs_flg":  checkNullString(p.EpbMrgnDbcrPrcsFlg),
						"epb_dp_id":               checkNullString(p.EpbDpId),
						"epb_dp_clnt_id":          checkNullString(p.EpbDpClntId),
						"epb_pledge_stts":         checkNullString(p.EpbPledgeStts),
						"epb_btst_net_mrgn_amt":   p.EpbBtstNetMrgnAmt,
						"epb_btst_mrgn_blckd":     p.EpbBtstMrgnBlckd,
						"epb_btst_mrgn_dbcr_flg":  checkNullString(p.EpbBtstMrgnDbcrFlg),
						"epb_btst_sgmnt_cd":       checkNullString(p.EpbBtstSgmntCd),
						"epb_btst_stlmnt":         p.EpbBtstStlmnt,
						"epb_btst_csh_blckd":      p.EpbBtstCshBlckd,
						"epb_btst_sam_blckd":      p.EpbBtstSamBlckd,
						"epb_btst_calc_dt":        checkNullString(p.EpbBtstCalcDt),
						"epb_dbcr_calc_dt":        checkNullString(p.EpbDbcrCalcDt),
						"epb_nsdl_ref_no":         checkNullString(p.EpbNsdlRefNo),
						"epb_mrgn_withheld_flg":   checkNullString(p.EpbMrgnWithheldFlg),
					})
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"positions": positions,
			})
		} else {
			log.Println("No EPB positions found")
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "No positions found",
			})
		}
	})

	// Define the GET route for OTP positions
	router.GET("/OTPService/:otp_clm_mtch_acct", func(ctx *gin.Context) {
		accnt := ctx.Param("otp_clm_mtch_acct")
		log.Printf("Received request for OTP account: %s", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Calling gRPC GetOtpPosition service for OTP")
		res, err := otpClient.GetOtpPosition(ctxWithTimeout, &pb.OtpPositionRequest{OtpClmMtchAcct: accnt})
		if err != nil {
			log.Printf("Error fetching OTP position from gRPC service: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		if res.OtpPosition != nil {
			log.Println("Processing OTP positions received from gRPC service")
			positions := []gin.H{}

			for _, p := range res.OtpPosition {
				if p != nil {
					log.Printf("Processing OTP position for stock: %s", p.OtpStckCd)
					positions = append(positions, gin.H{
						"otp_clm_mtch_acct":       checkNullString(p.OtpClmMtchAcct),
						"otp_xchng_cd":            checkNullString(p.OtpXchngCd),
						"otp_xchng_sgmnt_cd":      checkNullString(p.OtpXchngSgmntCd),
						"otp_xchng_sgmnt_sttlmnt": p.OtpXchngSgmntSttlmnt,
						"otp_stck_cd":             checkNullString(p.OtpStckCd),
						"otp_flw":                 checkNullString(p.OtpFlw),
						"otp_qty":                 p.OtpQty,
						"otp_cnvrt_dlvry_qty":     p.OtpCnvrtDlvryQty,
						"otp_cvrd_qty":            p.OtpCvrdQty,
						"otp_rt":                  checkNullString(p.OtpRt),
						"otp_mrgn_amt":            p.OtpMrgnAmt,
						"otp_trd_val":             p.OtpTrdVal,
						"otp_rmrks":               checkNullString(p.OtpRmrks),
						"otp_xfer_mrgn_stts":      checkNullString(p.OtpXferMrgnStts),
						"otp_sell_opn_prccsd":     checkNullString(p.OtpSellOpnPrccsd),
						"otp_buy_opn_prccsd":      checkNullString(p.OtpBuyOpnPrccsd),
						"otp_mrgn_sqroff_mode":    checkNullString(p.OtpMrgnSqroffMode),
						"otp_em_trdsplt_prcs_flg": checkNullString(p.OtpEmTrdspltPrcsFlg),
						"otp_mtm_flg":             checkNullString(p.OtpMtmFlg),
						"otp_mtm_cansq":           checkNullString(p.OtpMtmCansq),
						"otp_eos_can":             checkNullString(p.OtpEosCan),
						"otp_trgr_prc":            p.OtpTrgrPrc,
						"otp_16_trgr_prc":         p.Otp_16TrgrPrc,
						"otp_min_mrgn":            p.OtpMinMrgn,
						"otp_t5_trdsplt_prcs_flg": checkNullString(p.OtpT5TrdspltPrcsFlg),
					})
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"positions": positions,
			})
		} else {
			log.Println("No OTP positions found")
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "No positions found",
			})
		}
	})

	// Start the server
	log.Println("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Helper function to handle null strings
func checkNullString(s string) string {
	if s == "" {
		return "NULL"
	}
	return s
}
