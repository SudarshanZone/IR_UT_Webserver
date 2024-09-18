package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/krishnakashyap0704/microservices/equSquareOff/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

type Epb struct {
	EpbClmMtchAccnt      string  `json:"epbClmMtchAccnt"`
	EpbXchngCd           string  `json:"epbXchngCd"`
	EpbXchngSgmntCd      string  `json:"epbXchngSgmntCd"`
	EpbXchngSgmntSttlmnt int64   `json:"epbXchngSgmntSttlmnt"`
	EpbStckCd            string  `json:"epbStckCd"`
	EpbOrgnlPstnQty      int64   `json:"epbOrgnlPstnQty"`
	EpbRate              float32 `json:"epbRate"`
	EpbOrgnlAmtPayble    float32 `json:"epbOrgnlAmtPayble"`
	EpbOrgnlMrgnAmt      float32 `json:"epbOrgnlMrgnAmt"`
	EpbSellQty           int64   `json:"epbSellQty"`
	EpbCvrOrdQty         int64   `json:"epbCvrOrdQty"`
	EpbNetMrgnAmt        float32 `json:"epbNetMrgnAmt"`
	EpbNetAmtPayble      float32 `json:"epbNetAmtPayble"`
	EpbNetPstnQty        int64   `json:"epbNetPstnQty"`
	EpbCtdQty            int64   `json:"epbCtdQty"`
	EpbPstnStts          string  `json:"epbPstnStts"`
	EpbLpcCalcStts       string  `json:"epbLpcCalcStts"`
	EpbSqroffMode        string  `json:"epbSqroffMode"`
	EpbPstnTrdDt         string  `json:"epbPstnTrdDt"`
	EpbMtmPrcsFlg        string  `json:"epbMtmPrcsFlg"`
	EpbLastMdfcnDt       string  `json:"epbLastMdfcnDt"`
	EpbInsDate           string  `json:"epbInsDate"`
	EpbCloseDate         string  `json:"epbCloseDate"`
	EpbSysFailFlg        string  `json:"epbSysFailFlg"`
	EpbLastPymntDt       string  `json:"epbLastPymntDt"`
	EpbLpcCalcEndDt      string  `json:"epbLpcCalcEndDt"`
	EpbMtmCansq          string  `json:"epbMtmCansq"`
	EpbExpiryDt          string  `json:"epbExpiryDt"`
	EpbMinMrgn           float32 `json:"epbMinMrgn"`
	EpbMrgnDbcrPrcsFlg   string  `json:"epbMrgnDbcrPrcsFlg"`
	EpbDpId              string  `json:"epbDpId"`
	EpbDpClntId          string  `json:"epbDpClntId"`
	EpbPledgeStts        string  `json:"epbPledgeStts"`
	EpbBtstNetMrgnAmt    float32 `json:"epbBtstNetMrgnAmt"`
	EpbBtstMrgnBlckd     float32 `json:"epbBtstMrgnBlckd"`
	EpbBtstMrgnDbcrFlg   string  `json:"epbBtstMrgnDbcrFlg"`
	EpbBtstSgmntCd       string  `json:"epbBtstSgmntCd"`
	EpbBtstStlmnt        int64   `json:"epbBtstStlmnt"`
	EpbBtstCshBlckd      float32 `json:"epbBtstCshBlckd"`
	EpbBtstSamBlckd      float32 `json:"epbBtstSamBlckd"`
	EpbBtstCalcDt        string  `json:"epbBtstCalcDt"`
	EpbDbcrCalcDt        string  `json:"epbDbcrCalcDt"`
	EpbNsdlRefNo         string  `json:"epbNsdlRefNo"`
}

type OtpTrdPstn struct {
	OtpClmMtchAcct       string  `json:"otp_clm_mtch_acct"`
	OtpXchngCd           string  `json:"otp_xchng_cd,omitempty"`
	OtpXchngSgmntCd      string  `json:"otp_xchng_sgmnt_cd,omitempty"`
	OtpXchngSgmntSttlmnt int     `json:"otp_xchng_sgmnt_sttlmnt,omitempty"`
	OtpStckCd            string  `json:"otp_stck_cd,omitempty"`
	OtpFlw               string  `json:"otp_flw"`
	OtpQty               int64   `json:"otp_qty"`
	OtpCnvrtDlvryQty     int64   `json:"otp_cnvrt_dlvry_qty"`
	OtpCvrdQty           int64   `json:"otp_cvrd_qty"`
	OtpRt                float64 `json:"otp_rt"`
	OtpMrgnAmt           float64 `json:"otp_mrgn_amt"`
	OtpTrdVal            float64 `json:"otp_trd_val"`
	OtpRmrks             string  `json:"otp_rmrks,omitempty"`
	OtpXferMrgnStts      string  `json:"otp_xfer_mrgn_stts,omitempty"`
	OtpSellOpnPrccsd     string  `json:"otp_sell_opn_prccsd,omitempty"`
	OtpBuyOpnPrccsd      string  `json:"otp_buy_opn_prccsd,omitempty"`
	OtpMrgnSqroffMode    string  `json:"otp_mrgn_sqroff_mode,omitempty"`
	OtpEmTrdspltPrcsFlg  string  `json:"otp_em_trdsplt_prcs_flg,omitempty"`
	OtpMtmFlg            string  `json:"otp_mtm_flg,omitempty"`
	OtpMtmCansq          string  `json:"otp_mtm_cansq,omitempty"`
	OtpEosCan            string  `json:"otp_eos_can,omitempty"`
	OtpTrgrPrc           float64 `json:"otp_trgr_prc,omitempty"`
	Otp16TrgrPrc         float64 `json:"otp_16_trgr_prc,omitempty"`
	OtpMinMrgn           float64 `json:"otp_min_mrgn,omitempty"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to gRPC server")

	client := pb.NewSqureoffClient(conn)
	r := gin.Default()

	r.POST("/:id/:type", func(ctx *gin.Context) {
		var requests []map[string]interface{}
		if err := ctx.BindJSON(&requests); err != nil {
			log.Printf("Invalid request data: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}
		log.Printf("Received %d request(s)", len(requests))

		for _, request := range requests {
			if hasField(request, "epbClmMtchAccnt") {
				var epb Epb
				if err := mapToStruct(request, &epb); err != nil {
					log.Printf("Error mapping request to Epb struct: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				req := &pb.Epb_SquareoffRequest{
					Id:   ctx.Param("id"),
					Type: ctx.Param("type"),
					Epb: &pb.EpbEq{
						EpbClmMtchAccnt:      epb.EpbClmMtchAccnt,
						EpbXchngCd:           epb.EpbXchngCd,
						EpbXchngSgmntCd:      epb.EpbXchngSgmntCd,
						EpbXchngSgmntSttlmnt: epb.EpbXchngSgmntSttlmnt,
						EpbStckCd:            epb.EpbStckCd,
						EpbOrgnlPstnQty:      epb.EpbOrgnlPstnQty,
						EpbRate:              epb.EpbRate,
						EpbOrgnlAmtPayble:    epb.EpbOrgnlAmtPayble,
						EpbOrgnlMrgnAmt:      epb.EpbOrgnlMrgnAmt,
						EpbSellQty:           epb.EpbSellQty,
						EpbCvrOrdQty:         epb.EpbCvrOrdQty,
						EpbNetMrgnAmt:        epb.EpbNetMrgnAmt,
						EpbNetAmtPayble:      epb.EpbNetAmtPayble,
						EpbNetPstnQty:        epb.EpbNetPstnQty,
						EpbCtdQty:            epb.EpbCtdQty,
						EpbPstnStts:          epb.EpbPstnStts,
						EpbLpcCalcStts:       epb.EpbLpcCalcStts,
						EpbSqroffMode:        epb.EpbSqroffMode,
						EpbPstnTrdDt:         epb.EpbPstnTrdDt,
						EpbMtmPrcsFlg:        epb.EpbMtmPrcsFlg,
						EpbLastMdfcnDt:       epb.EpbLastMdfcnDt,
						EpbInsDate:           epb.EpbInsDate,
						EpbCloseDate:         epb.EpbCloseDate,
						EpbSysFailFlg:        epb.EpbSysFailFlg,
						EpbLastPymntDt:       epb.EpbLastPymntDt,
						EpbLpcCalcEndDt:      epb.EpbLpcCalcEndDt,
						EpbMtmCansq:          epb.EpbMtmCansq,
						EpbExpiryDt:          epb.EpbExpiryDt,
						EpbMinMrgn:           epb.EpbMinMrgn,
						EpbMrgnDbcrPrcsFlg:   epb.EpbMrgnDbcrPrcsFlg,
						EpbDpId:              epb.EpbDpId,
						EpbDpClntId:          epb.EpbDpClntId,
						EpbPledgeStts:        epb.EpbPledgeStts,
						EpbBtstNetMrgnAmt:    epb.EpbBtstNetMrgnAmt,
						EpbBtstMrgnBlckd:     epb.EpbBtstMrgnBlckd,
						EpbBtstMrgnDbcrFlg:   epb.EpbBtstMrgnDbcrFlg,
						EpbBtstSgmntCd:       epb.EpbBtstSgmntCd,
						EpbBtstStlmnt:        epb.EpbBtstStlmnt,
						EpbBtstCshBlckd:      epb.EpbBtstCshBlckd,
						EpbBtstSamBlckd:      epb.EpbBtstSamBlckd,
						EpbBtstCalcDt:        epb.EpbBtstCalcDt,
						EpbDbcrCalcDt:        epb.EpbDbcrCalcDt,
						EpbNsdlRefNo:         epb.EpbNsdlRefNo,
					},
				}
				fmt.Println("-------------=======================--------------------------")
				fmt.Println(req.Epb)
				grpcResp, err := client.Squareoff_Epb(context.Background(), req)
				if err != nil {
					log.Printf("Error calling gRPC Squareoff_Epb: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				log.Printf("gRPC response for Epb: %v", grpcResp)
				ctx.JSON(http.StatusOK, gin.H{"message": grpcResp})
			} else if hasField(request, "otp_clm_mtch_acct") {
				var otp OtpTrdPstn
				if err := mapToStruct(request, &otp); err != nil {
					log.Printf("Error mapping request to OtpTrdPstn struct: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				req := &pb.Opt_SquareoffRequest{
					Id:   ctx.Param("id"),
					Type: ctx.Param("type"),
					Otp: &pb.OtpEq{
						OtpClmMtchAcct:       otp.OtpClmMtchAcct,
						OtpXchngCd:           otp.OtpXchngCd,
						OtpXchngSgmntCd:      otp.OtpXchngSgmntCd,
						OtpXchngSgmntSttlmnt: int64(otp.OtpXchngSgmntSttlmnt),
						OtpStckCd:            otp.OtpStckCd,
						OtpFlw:               otp.OtpFlw,
						OtpQty:               otp.OtpQty,
						OtpCnvrtDlvryQty:     otp.OtpCnvrtDlvryQty,
						OtpCvrdQty:           otp.OtpCvrdQty,
						OtpRt:                float32(otp.OtpRt),
						OtpMrgnAmt:           float32(otp.OtpMrgnAmt),
						OtpTrdVal:            float32(otp.OtpTrdVal),
						OtpRmrks:             otp.OtpRmrks,
						OtpXferMrgnStts:      otp.OtpXferMrgnStts,
						OtpSellOpnPrccsd:     otp.OtpSellOpnPrccsd,
						OtpBuyOpnPrccsd:      otp.OtpBuyOpnPrccsd,
						OtpMrgnSqroffMode:    otp.OtpMrgnSqroffMode,
						OtpEmTrdspltPrcsFlg:  otp.OtpEmTrdspltPrcsFlg,
						OtpMtmFlg:            otp.OtpMtmFlg,
						OtpMtmCansq:          otp.OtpMtmCansq,
						OtpEosCan:            otp.OtpEosCan,
						OtpTrgrPrc:           float32(otp.OtpTrgrPrc),
						Otp_16TrgrPrc:        float32(otp.Otp16TrgrPrc),
						OtpMinMrgn:           float32(otp.OtpMinMrgn),
					},
				}
				grpcResp, err := client.Squareoff_Otp(context.Background(), req)
				if err != nil {
					log.Printf("Error calling gRPC Squareoff_Otp: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				log.Printf("gRPC response for Otp: %v", grpcResp)
				ctx.JSON(http.StatusOK, gin.H{"message": grpcResp})
			} else {
				log.Println("Unknown request type")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown request type"})
				return
			}
		}
	})

	log.Println("Starting server on :8000")
	r.Run(":8000")
}

func hasField(m map[string]interface{}, field string) bool {
	_, exists := m[field]
	return exists
}

func mapToStruct(m map[string]interface{}, result interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return json.Unmarshal(data, result)
}
