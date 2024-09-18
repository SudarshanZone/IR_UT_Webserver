package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pb "github.com/krishnakashyap0704/microservices/comSquareOff/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

// CcpCodSpnCntrctPstn represents the contract position structure
type CcpCodSpnCntrctPstn struct {
	CcpAddnlMrgn        float64 `json:"ccp_addnl_mrgn"`
	CcpAsgndQty         int64   `json:"ccp_asgnd_qty"`
	CcpAvgPrc           float64 `json:"ccp_avg_prc"`
	CcpBuyExctdQty      int64   `json:"ccp_buy_exctd_qty"`
	CcpClmMtchAccnt     string  `json:"ccp_clm_mtch_accnt"`
	CcpXchngCd          string  `json:"ccp_xchng_cd"`
	CcpPrdctTyp         string  `json:"ccp_prdct_typ"`
	CcpIndstk           string  `json:"ccp_indstk"`
	CcpUndrlyng         string  `json:"ccp_undrlyng"`
	CcpExpryDt          string  `json:"ccp_expry_dt"`
	CcpExerTyp          string  `json:"ccp_exer_typ"`
	CcpStrkPrc          int64   `json:"ccp_strk_prc"`
	CcpOptTyp           string  `json:"ccp_opt_typ"`
	CcpIbuyQty          int64   `json:"ccp_ibuy_qty"`
	CcpIbuyOrdVal       float64 `json:"ccp_ibuy_ord_val"`
	CcpIsellQty         int64   `json:"ccp_isell_qty"`
	CcpIsellOrdVal      float64 `json:"ccp_isell_ord_val"`
	CcpExbuyQty         int64   `json:"ccp_exbuy_qty"`
	CcpExbuyOrdVal      float64 `json:"ccp_exbuy_ord_val"`
	CcpExsellQty        int64   `json:"ccp_exsell_qty"`
	CcpExsellOrdVal     float64 `json:"ccp_exsell_ord_val"`
	CcpSellExctdQty     int64   `json:"ccp_sell_exctd_qty"`
	CcpOpnpstnFlw       string  `json:"ccp_opnpstn_flw"`
	CcpOpnpstnQty       int64   `json:"ccp_opnpstn_qty"`
	CcpOpnpstnVal       float64 `json:"ccp_opnpstn_val"`
	CcpExrcQty          int64   `json:"ccp_exrc_qty"`
	CcpOptPremium       float64 `json:"ccp_opt_premium"`
	CcpMtmOpnVal        float64 `json:"ccp_mtm_opn_val"`
	CcpImtmOpnVal       float64 `json:"ccp_imtm_opn_val"`
	CcpExtrmlossMrgn    float64 `json:"ccp_extrmloss_mrgn"`
	CcpSpclMrgn         float64 `json:"ccp_spcl_mrgn"`
	CcpTndrMrgn         float64 `json:"ccp_tndr_mrgn"`
	CcpDlvryMrgn        float64 `json:"ccp_dlvry_mrgn"`
	CcpExtrmMinLossMrgn float64 `json:"ccp_extrm_min_loss_mrgn"`
	CcpMtmFlg           string  `json:"ccp_mtm_flg"`
	CcpExtLossMrgn      float64 `json:"ccp_ext_loss_mrgn"`
	CcpFlatValMrgn      float64 `json:"ccp_flat_val_mrgn"`
	CcpTrgPrc           float64 `json:"ccp_trg_prc"`
	CcpMinTrgPrc        float64 `json:"ccp_min_trg_prc"`
	CcpDevolmntMrgn     float64 `json:"ccp_devolmnt_mrgn"`
	CcpMtmsqOrdcnt      int32   `json:"ccp_mtmsq_ordcnt"`
}

func main() {
	r := gin.Default()

	// CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Update with your frontend URL
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// gRPC client connection
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSquareoffClient(conn)

	r.POST("/squareoff/order/:id/:type", func(ctx *gin.Context) {
		var ccp CcpCodSpnCntrctPstn

		// Bind incoming JSON body directly to struct
		if err := ctx.ShouldBindJSON(&ccp); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			log.Println("Error binding JSON:", err)
			return
		}
		//fmt.Print(ccp)
		// Create gRPC request
		req := &pb.SquareoffRequest{
			Ccp: &pb.CcpCodSpnCntrctPstn{
				CcpClmMtchAccnt:     ccp.CcpClmMtchAccnt,
				CcpXchngCd:          ccp.CcpXchngCd,
				CcpPrdctTyp:         ccp.CcpPrdctTyp,
				CcpIndstk:           ccp.CcpIndstk,
				CcpUndrlyng:         ccp.CcpUndrlyng,
				CcpExpryDt:          ccp.CcpExpryDt,
				CcpExerTyp:          ccp.CcpExerTyp,
				CcpStrkPrc:          ccp.CcpStrkPrc,
				CcpOptTyp:           ccp.CcpOptTyp,
				CcpIbuyQty:          ccp.CcpIbuyQty,
				CcpIbuyOrdVal:       ccp.CcpIbuyOrdVal,
				CcpIsellQty:         ccp.CcpIsellQty,
				CcpIsellOrdVal:      ccp.CcpIsellOrdVal,
				CcpExbuyQty:         ccp.CcpExbuyQty,
				CcpExbuyOrdVal:      ccp.CcpExbuyOrdVal,
				CcpExsellQty:        ccp.CcpExsellQty,
				CcpExsellOrdVal:     ccp.CcpExsellOrdVal,
				CcpBuyExctdQty:      ccp.CcpBuyExctdQty,
				CcpSellExctdQty:     ccp.CcpSellExctdQty,
				CcpOpnpstnFlw:       ccp.CcpOpnpstnFlw,
				CcpOpnpstnQty:       ccp.CcpOpnpstnQty,
				CcpOpnpstnVal:       ccp.CcpOpnpstnVal,
				CcpExrcQty:          ccp.CcpExrcQty,
				CcpAsgndQty:         ccp.CcpAsgndQty,
				CcpOptPremium:       ccp.CcpOptPremium,
				CcpMtmOpnVal:        ccp.CcpMtmOpnVal,
				CcpImtmOpnVal:       ccp.CcpImtmOpnVal,
				CcpExtrmlossMrgn:    ccp.CcpExtrmlossMrgn,
				CcpSpclMrgn:         ccp.CcpSpclMrgn,
				CcpTndrMrgn:         ccp.CcpTndrMrgn,
				CcpDlvryMrgn:        ccp.CcpDlvryMrgn,
				CcpExtrmMinLossMrgn: ccp.CcpExtrmMinLossMrgn,
				CcpMtmFlg:           ccp.CcpMtmFlg,
				CcpExtLossMrgn:      ccp.CcpExtLossMrgn,
				CcpFlatValMrgn:      ccp.CcpFlatValMrgn,
				CcpTrgPrc:           ccp.CcpTrgPrc,
				CcpMinTrgPrc:        ccp.CcpMinTrgPrc,
				CcpDevolmntMrgn:     ccp.CcpDevolmntMrgn,
				CcpMtmsqOrdcnt:      ccp.CcpMtmsqOrdcnt,
				CcpAvgPrc:           ccp.CcpAvgPrc,
			},
			CCP_USR_ID:    ctx.Param("id"),
			CCP_PRDCT_TYP: ctx.Param("type"),
		}

		fmt.Println(req.Ccp)

		// Call gRPC method
		res, err := client.Squareoff(context.Background(), req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(res)

		// Respond with gRPC result
		ctx.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	})

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
