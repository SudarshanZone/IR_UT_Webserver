package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	pbEquOpen "github.com/SudarshanZone/webserver/generated"
	pbEquOrds "github.com/SudarshanZone/webserver/generated"
	pbFnoOpen "github.com/SudarshanZone/webserver/generated"
	pbFnoOrd "github.com/SudarshanZone/webserver/generated"
	pbFnoSquoff "github.com/SudarshanZone/webserver/generated"
	pbcommOpen "github.com/SudarshanZone/webserver/generated"   
	pbcommOrds "github.com/SudarshanZone/webserver/generated"   
	pbequSqoff "github.com/SudarshanZone/webserver/generated"
	pbTrdRec "github.com/SudarshanZone/webserver/generated"
	pbLogin "github.com/SudarshanZone/webserver/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log *logrus.Logger

func InitLogger() *logrus.Logger {
	log = logrus.New()

	logFilename := filepath.Join("logs", "ULOG."+time.Now().Format("2006-01-02")+".log")

	if err := os.MkdirAll(filepath.Dir(logFilename), 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		os.Exit(1)
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilename,
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	log.SetOutput(getLogOutput(lumberjackLogger))
	log.SetFormatter(&CustomFormatter{})
	log.SetLevel(getLogLevel())
	log.AddHook(NewContextHook())

	return log
}

func getLogOutput(lumberjackLogger *lumberjack.Logger) io.Writer {
	if os.Getenv("ENV") == "production" {
		return lumberjackLogger
	}
	return io.MultiWriter(lumberjackLogger, os.Stdout)
}

func getLogLevel() logrus.Level {
	return logrus.DebugLevel
	// switch os.Getenv("ENV") {
	// case "production":
	// 	return logrus.WarnLevel
	// case "staging":
	// 	return logrus.InfoLevel
	// default:
	// 	return logrus.DebugLevel
	// }
}

type ContextHook struct{}

func NewContextHook() *ContextHook {
	return &ContextHook{}
}

func (hook *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	if _, file, line, ok := runtime.Caller(6); ok {
		shortFile := filepath.Base(file)
		entry.Data["file"] = shortFile
		entry.Data["line"] = line
	}
	return nil
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	serviceName, _ := entry.Data["service"].(string)
	message := entry.Message

	formatted := timestamp + " " + serviceName + ": " + message + "\n"
	return []byte(formatted), nil
}

func (f *CustomFormatter) Format2(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006/01/02 - 15:04:05")
	statusCode, _ := entry.Data["status_code"].(int)
	duration, _ := entry.Data["duration"].(float64)
	clientIP, _ := entry.Data["client_ip"].(string)
	method, _ := entry.Data["method"].(string)
	path, _ := entry.Data["path"].(string)

	formatted := fmt.Sprintf("[GIN] %s | %d | %.6fs | %s | %s | %s\n",
		timestamp, statusCode, duration, clientIP, method, path)

	return []byte(formatted), nil
}



//Equity Square off Models
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	//Logger
	InitLogger()

	//Gin Instailization
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Initialize gRPC connection for Login Newclient is used instead of grpc.Dial()
	connLogin, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connLogin.Close()
	grpcLoginClient := pbLogin.NewUserLoginServiceClient(connLogin)

	// Initialize gRPC connection for FnoOpenPosService
	connFnoOpen, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server of fnoOpenPosService: %v", err)
	}
	defer connFnoOpen.Close()
	pbFnoOpenPosClient := pbFnoOpen.NewFnoPositionServiceClient(connFnoOpen)

	//Intialize grpc connection for FnoOrderDetails
	connFnoOrd, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not Connect to grpc server fnoOpenOrdService: %v", err)
	}
	defer connFnoOrd.Close()
	pbFnoOrdClient := pbFnoOrd.NewOrderDetailsServiceClient(connFnoOrd)


	//Intialize the grpc Connection for Equity Orders Service
	connEquOrds ,err := grpc.NewClient("localhost:50055",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could Not COnnect to grpc Server Equity Orders Service : %v",err)
	}
	defer connEquOrds.Close()
	pbEquOrdsClient := pbEquOrds.NewEquityOrderServiceClient(connEquOrds)
	

	//Intialize the grpc connection for Equity Open Positions
	connEquOpen,err := grpc.NewClient("localhost:50054",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could Not COnnect to grpc Server Equity Open Positions Service : %v",err)
	}
	defer connEquOpen.Close()
	pbEquOpenPosClient := pbEquOpen.NewPositionServiceClient(connEquOpen)



	connCommOpen ,err := grpc.NewClient("localhost:50057",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could not Connect to grpc server of Commodities Open position Service")
	}
	defer connCommOpen.Close()
	pbcommOpenClient := pbcommOpen.NewCCPServiceClient(connCommOpen)

	//Intialize the grpc connection for Commodities Orders Service
	connCommOrds,err := grpc.NewClient("localhost:50058",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could not Connect to grpc server of Commodities Orders Service")
	}
	defer connCommOrds.Close()
	pbcommOrdsClient := pbcommOrds.NewC_OrderServiceClient(connCommOrds)


	//Intilize the grpc connection for FNO Square off Service
	connFnoSqoff,err := grpc.NewClient("localhost:50059",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could not Connect to grpc server of FNO Square off Service")
	}
	defer connFnoSqoff.Close()
	pbFnoSquoffClient := pbFnoSquoff.NewSquareOffServiceClient(connFnoSqoff)


	//Intialize the grpc connection for Eqity Square off Service
	connEqsSqoff,err := grpc.NewClient("localhost:50056",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil{
		log.Fatalf("Could Not COnnect to grpc Server Equity Square off Service : %v",err)
	}
	defer connEqsSqoff.Close()
	pbEquSquareOffClient := pbequSqoff.NewSqureoffClient(connEqsSqoff)

	//Intialize grpc connection for Trade Records Service
	connTradeRec,err := grpc.NewClient("localhost:50053",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Could Not COnnect to grpc Server TradeRecords Service : %v",err)
	}
	defer connTradeRec.Close()
	pbTrdRecClient := pbTrdRec.NewDataServiceClient(connTradeRec)
	//NewDataServiceClient




	//Setup for FNO Square off Service
	
	// Struct to handle JSON input from frontend
	type SquareOffRequest struct {
		FcpDetails []*pbFnoSquoff.FnoData `json:"FcpDetails"`
	}

	// Handler function to process square off orders
	handleSquareOff := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			var reqBody SquareOffRequest

			// Extract user ID and product type from query parameters
			FFO_USR_ID_Str := c.Query("FFO_USR_ID")
			fmt.Println("Required ID: ", FFO_USR_ID_Str)
			if FFO_USR_ID_Str == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "FFO_USR_ID is required"})
				return
			}

			// Extract user ID and product type from query parameters
			FFO_USR_ID, err := strconv.ParseInt(c.Query("FFO_USR_ID"), 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FFO_USR_ID", "details": err.Error()})
				return
			}
			FFO_PRDCT_TYP := c.Query("FFO_PRDCT_TYP")

			// Bind JSON input to struct
			if err := c.ShouldBindJSON(&reqBody); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input", "details": err.Error()})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			// Create the gRPC request
			req := &pbFnoSquoff.SquareOffRequest{
				FFO_USR_ID:    FFO_USR_ID,
				FFO_PRDCT_TYP: FFO_PRDCT_TYP,
				FcpDetails:    reqBody.FcpDetails,
			}

			// Call the SquareOffOrder RPC method
			resp, err := pbFnoSquoffClient.SquareOffOrder(ctx, req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "gRPC square off order failed", "details": err.Error()})
				return
			}

			// Return the response to the client
			c.JSON(http.StatusOK, gin.H{"status": resp.Status, "message": resp.Message})
		}
	}
	
	//Router for FNO Square off Service
	r.POST("/squareoff/order", handleSquareOff())


	//Route for Trade records Service
	r.GET("/download/:type/:sub_type", func(c *gin.Context) {
		typ := c.Param("type")
		sub_typ := c.Param("sub_type")

		log.Infof("Received download request: type=%s, sub_type=%s", typ, sub_typ)

		
		
		stream, err := pbTrdRecClient.GetCSV(context.Background(), &pbTrdRec.CSVDataRequest{Type: typ, SubType: sub_typ})
		if err != nil {
			log.Infof("Failed to get CSV data from gRPC server: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Disposition", "attachment; filename=data.csv")
		c.Header("Content-Type", "text/csv")

		writer := c.Writer
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				log.Infof("Completed receiving data from gRPC server")
				break
			}
			if err != nil {
				log.Infof("Error receiving data from gRPC server: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			_, err = writer.Write([]byte(msg.Data))
			if err != nil {
				log.Infof("Error writing data to response: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			log.Info("Sent data  to client")
		}

	})


	//Route for Login Service
	r.POST("/login", func(c *gin.Context) {
		var formData struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&formData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Log 
		log.Printf("Attempting Login for user : %s",formData.Username)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := grpcLoginClient.Login(ctx, &pbLogin.LoginRequest{
			Username: formData.Username,
			Password: formData.Password,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": "gRPC login failed", "details": err.Error()})
			return
		}

		//
		log.Printf("Login successful for user: %s, UserType: %s", formData.Username, resp.UserType)
		c.JSON(200, gin.H{"success": resp.Success, "message": resp.Message, "userType": resp.UserType})
	})

	//Route for Equity Square off Service	
	r.POST("/:id/:type", func(ctx *gin.Context) {
		var requests []map[string]interface{}
		if err := ctx.BindJSON(&requests); err != nil {
			log.Infof("Invalid request data: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}
		log.Infof("Received %d request(s)", len(requests))

		for _, request := range requests {
			if hasField(request, "epbClmMtchAccnt") {
				var epb Epb
				if err := mapToStruct(request, &epb); err != nil {
					log.Infof("Error mapping request to Epb struct: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				req := &pbequSqoff.Epb_SquareoffRequest{
					Id:   ctx.Param("id"),
					Type: ctx.Param("type"),
					Epb: &pbequSqoff.EpbEq{
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
				grpcResp, err := pbEquSquareOffClient.Squareoff_Epb(context.Background(), req)
				if err != nil {
					log.Infof("Error calling gRPC Squareoff_Epb: %v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				log.Infof("gRPC response for Epb: %v", grpcResp)
				ctx.JSON(http.StatusOK, gin.H{"message": grpcResp})
			}else {
				log.Println("Unknown request type")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown request type"})
				return
			}
		}
	})

	//Route for FnoOpenPosition Service
	r.GET("/getFNOPosition/:UCCID", func(c *gin.Context) {
		UCCid := c.Param("UCCID")

		log.Printf("Received request to fetch Positions details for UCCID: %s", UCCid)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req := &pbFnoOpen.FnoPositionRequest{FCP_CLM_MTCH_ACCNT: UCCid}

		resp, err := pbFnoOpenPosClient.GetFNOPosition(ctx, req)
		if err != nil {
			log.Printf("Error fetching FNO positions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FNO positions"})
			return
		}

		log.Printf("Successfully fetched order details for UCCID: %s", UCCid)
		c.JSON(http.StatusOK, gin.H{"positions": resp.FcpDetails})
	})

	//Route for FnoOrderDetails Service
	r.GET("/getOrderDetails/:OrderID", func(c *gin.Context) {
		OrderID := c.Param("OrderID")
		log.Printf("Error fetching FNO Orders: %v", err)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Create request object for gRPC
		req := &pbFnoOrd.OrderDetailsRequest{FOD_CLM_MTCH_ACCNT: OrderID}

		resp, err := pbFnoOrdClient.GetOrderDetails(ctx, req)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"orderDetails": resp.OrdDetails})
	})

	//Route for Commodities Open Position Service
	r.GET("/C_OrderService/:cod_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("cod_clm_mtch_accnt")
		log.Printf("Received request for account: %s", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Calling GetOrder on gRPC client...")
		res, err := pbcommOpenClient.GetCCPData(ctxWithTimeout, &pbcommOpen.CCPRequest{CcpClmMtchAccnt: accnt})
		if err != nil {
			log.Printf("Error fetching order details: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Error fetching order details: %v", err),
			})
			return
		}

		response := gin.H{}
		if res.Commo != nil {
			log.Printf("Order details found for account: %s", accnt)
			response["Commodity_order_details"] = res.Commo
			ctx.JSON(http.StatusOK, response)
		} else {
			log.Printf("No order details found for account: %s", accnt)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Commodity order details not found",
			})
		}
	})

	//Route for Commodities Orders Service
	r.GET("/C_OrderService2/:cod_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("cod_clm_mtch_accnt")
		log.Printf("Received request for account: %s", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Calling GetOrder on gRPC client...")
		res, err := pbcommOrdsClient.GetComOrder(ctxWithTimeout, &pbcommOrds.ComOrderRequest{CodClmMtchAccnt: accnt})
		if err != nil {
			log.Printf("Error fetching order details: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Error fetching order details: %v", err),
			})
			return
		}

		response := gin.H{}
		if res.OrdDtls != nil {
			log.Printf("Order details found for account: %s", accnt)
			response["Commodity_order_details"] = res.OrdDtls
			ctx.JSON(http.StatusOK, response)
		} else {
			log.Printf("No order details found for account: %s", accnt)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Commodity order details not found",
			})
		}
	})

	//Route for Equity Open Position Service
	r.GET("/PositionService/:epb_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("epb_clm_mtch_accnt")
		fmt.Println("Account:", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := pbEquOpenPosClient.GetPosition(ctxWithTimeout, &pbEquOpen.PositionRequest{EpbClmMtchAccnt: accnt})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		if res.Equity != nil {
			positions := []gin.H{}

			for _, p := range res.Equity {
				if p != nil { // Check if p is not nil
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
						//"epb_btst_stlmnt":         checkNullString(p.EpbBtstStlmnt),
						"epb_btst_stlmnt":       p.EpbBtstStlmnt,
						"epb_btst_csh_blckd":    p.EpbBtstCshBlckd,
						"epb_btst_sam_blckd":    p.EpbBtstSamBlckd,
						"epb_btst_calc_dt":      checkNullString(p.EpbBtstCalcDt),
						"epb_dbcr_calc_dt":      checkNullString(p.EpbDbcrCalcDt),
						"epb_nsdl_ref_no":       checkNullString(p.EpbNsdlRefNo),
						"epb_mrgn_withheld_flg": checkNullString(p.EpbMrgnWithheldFlg),
					})
				}
			}
			ctx.JSON(http.StatusOK, gin.H{"Equity_MTF_positions ": positions})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Equity positions not found",
			})
		}
	})
	
	//Route for Equity Orders Service
	r.GET("/OrderService/:ord_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("ord_clm_mtch_accnt")
		log.Infof("Received request for account: %s", accnt)

		// Set a timeout for the gRPC call
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Call the gRPC method pbEquOrdsClient
		log.Infof("Calling gRPC method GetOrder for account: %s", accnt)
		//res, err := pbEquOrdsClient.GetOrder(ctxWithTimeout, &pbEquOrds.OrderRequest{OrdClmMtchAccnt: accnt})
		res ,err := pbEquOrdsClient.GetEquityOrder(ctxWithTimeout,&pbEquOrds.EquityOrderRequest{OrdClmMtchAccnt: accnt})
		if err != nil {
			log.Infof("Error fetching order details for account %s: %v", accnt, err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Error fetching order details: %v", err),
			})
			return
		}

		// Process the gRPC response
		if res.OrdDtls != nil {
			log.Infof("Successfully retrieved order details for account: %s", accnt)
			ctx.JSON(http.StatusOK, gin.H{
				"equity_order_details": res.OrdDtls,
			})
		} else {
			log.Infof("No order details found for account: %s", accnt)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Equity order details not found",
			})
		}
	})
	
	
	log.Println("Web server is running on :8080")
	if err = r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func checkNullString(str string) string {
	if str == "" {
		return "NULL"
	}
	return str
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