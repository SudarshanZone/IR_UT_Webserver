package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/krishnakashyap0704/microservices/comSquareOff/generated"
	"github.com/krishnakashyap0704/microservices/comSquareOff/internal/database"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedSquareoffServer
}

func (s *server) Squareoff(ctx context.Context, req *pb.SquareoffRequest) (*pb.SquareoffResponse, error) {
	fmt.Println(req)
	if req == nil {
		return nil, errors.New("request is nil")
	}
	ccpPosition := req.GetCcp()

	if ccpPosition == nil {
		return nil, errors.New("ccpCodSpnCntrctPstn is nil")
	}

	// Parse expiry date
	const layout1 = "2006-01-02T15:04:05Z"
	ccpExpryDtStr := ccpPosition.CcpExpryDt // The string value
	ccpExpryDtTime, err := parseDate(ccpExpryDtStr, layout1)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expiry date: %w", err)
	}

	// Format the expiry date to a string if needed
	const dateFormat = "2006-01-02" // Adjust format as required
	ccpExpryDtStrFormatted := ccpExpryDtTime.Format(dateFormat)

	// Example of parsing another date format
	// layout2 := "2006-01-02 15:04:05" // Define layout if needed
	// expiryDateStr := "2024-08-30 15:04:05"
	// // expiryDate, err := time.Parse(layout2, expiryDateStr)
	// if err != nil {
	// 	return nil, fmt.Errorf("error parsing time: %w", err)
	// }

	var openpstnflw string
	if req.Ccp.CcpOpnpstnFlw == "B" {
		openpstnflw = "S"
	} else if req.Ccp.CcpOpnpstnFlw == "S" {
		openpstnflw = "B"
	}

	// fmt.Println(req)
	pipeID := 11
	orderReferenceNo := GenreateOrderReference(pipeID)

	codOrder := database.CodCodOrdrDtls{
		CodClmMtchAccnt:    req.CCP_USR_ID,
		CodClntCtgry:       0,
		CodOrdrRfrnc:       orderReferenceNo,
		CodPipeId:          string(pipeID),
		CodXchngCd:         ccpPosition.CcpXchngCd,
		CodPrdctTyp:        req.CCP_PRDCT_TYP,
		CodIndstk:          ccpPosition.CcpIndstk,
		CodUndrlyng:        ccpPosition.CcpUndrlyng,
		CodExpryDt:         ccpExpryDtStrFormatted, // Use formatted date string
		CodExerTyp:         ccpPosition.CcpExerTyp,
		CodOptTyp:          ccpPosition.CcpOptTyp,
		CodStrkPrc:         0,
		CodOrdrFlw:         openpstnflw,
		CodLmtMrktSlFlg:    "M",
		CodDsclsdQty:       0,
		CodOrdrTotQty:      float64(ccpPosition.CcpOpnpstnQty), // not confirmed
		CodLmtRt:           0,
		CodStpLssTgr:       0,
		CodOrdrType:        "T",
		CodOrdrValidDt:     time.Now(), // Use formatted date string
		CodTrdDt:           time.Now(), // Use formatted date string
		CodOrdrStts:        "R",        // bcz in the 1st stage
		CodSprdOrdrRef:     "",         // checked
		CodMdfctnCntr:      1,
		CodSettlor:         "",
		CodAckNmbr:         "ACK00000",
		CodSplFlag:         "I",
		CodOrdAckTm:        time.Now(),
		CodLstRqstAckTm:    time.Now(),
		CodProCliInd:       "C",
		CodExecQtyDay:      0,
		CodRemarks:         "test remark",
		CodChannel:         "WEB",
		CodBpId:            "BP00",
		CodCtclId:          "111111111111",
		CodUsrId:           "USR000", // not confirmed
		CodMrktTyp:         "N",
		CodCseId:           0,
		CodSpnFlg:          "L",
		CodSltpOrdrRfrnc:   "0000",
		CodAmtBlckd:        0,
		CodLssAmtBlckd:     0,
		CodFcFlag:          "M",
		CodDiffAmtBlckd:    0,
		CodDiffLssAmtBlckd: 0,
		CodTrdVal:          0,
		CodTrdBrkg:         0,
		CodCntrctntNmbr:    "CON0000",
		CodSourceFlg:       "O",
		CodEosFlg:          "",
		CodPrcimpvFlg:      "N",
		CodTrailAmt:        0,
		CodLmtOffset:       0,
		CodSrollDiffAmt:    0,
		CodSrollLssAmt:     0,
		CodSrollDiffAmtOld: 0,
		CodSrollLssAmtOld:  0,
		CodPanNo:           "PAN1234567",
		CodSetlmntFlg:      "O",
		CodLstActRef:       "00",
		CodQtyUnit:         "G",
		CodPrcUnit:         "RS/1KGS",
		CodPrcMltplr:       0,
		CodGenMltplr:       0,
		CodSsnTyp:          "W", // s1,s2
		CodIsFlg:           "N",
		CodExecQty:         float64(ccpPosition.CcpExrcQty), // not confirmed
		CodCnclQty:         0,
		CodExprdQty:        0,
		CodEspId:           "ESP0000",
		CodMinLotQty:       0,
		CodPrtctnRt:        0,
		CodSqroffTm:        3600,
		CodAvgExctdRt:      0,
	}

	cxbOrder := database.CxbCodXchngBook{
		CxbXchngCd:      "MCO",
		CxbOrdrRfrnc:    orderReferenceNo,
		CxbPipeId:       string(pipeID),
		CxbModTrdDt:     time.Now().Format(dateFormat), // Use formatted date string
		CxbOrdrSqnc:     0,
		CxbLmtMrktSlFlg: "M",
		CxbDsclsdQty:    0,
		CxbOrdrTotQty:   ccpPosition.CcpOpnpstnQty,
		CxbLmtRt:        0,
		CxbStpLssTgr:    0,
		CxbMdfctnCntr:   1,
		CxbOrdrValidDt:  ccpExpryDtStrFormatted, // Use formatted date string
		CxbOrdrType:     "M",
		CxbSprdOrdInd:   "*",
		CxbRqstTyp:      "N", // not confirmed
		CxbQuote:        0,
		CxbQtTm:         time.Now().Format(dateFormat), // Use formatted date string
		CxbRqstTm:       time.Now().Format(dateFormat), // Use formatted date string
		CxbFrwdTm:       time.Now().Format(dateFormat), // Use formatted date string
		CxbPlcdStts:     "0",                           // not confirmed
		CxbRmsPrcsdFlg:  "0",
		CxbPrcimpvFlg:   "N",
		CxbClntOrdId:    "0",
		CxbExgRespSeq:   0,
		CxbSssnId:       0,
		CxbAppmsgId:     "Appmsg01",
		// Populate additional fields as needed
	}

	// Insert orders into the database
	if err := database.DB.Debug().Create(&cxbOrder).Error; err != nil {
		return nil, fmt.Errorf("failed to insert into CXB table: %w", err)
	}
	log.Println("CXB order inserted successfully")

	if err := database.DB.Debug().Create(&codOrder).Error; err != nil {
		return nil, fmt.Errorf("failed to insert into COD table: %w", err)
	}
	log.Println("COD order inserted successfully")

	//added
	fmt.Printf("COD Order: %+v\n", codOrder)
	fmt.Printf("CXB Order: %+v\n", cxbOrder)

	return &pb.SquareoffResponse{Success: true}, nil
}

func parseDate(dateStr string, layout string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil // Return zero value for time.Time if the date string is empty
	}
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

var orderCounter = 0
var orderCounterMu sync.Mutex

func GenreateOrderReference(pipeID int) string {
	orderCounterMu.Lock()
	defer orderCounterMu.Unlock()
	currentDate := time.Now().Format("20060102")
	pipeIDStr := fmt.Sprintf("%02d", pipeID)
	orderCounter += 1
	counterStr := fmt.Sprintf("%08d", orderCounter)
	OrderReference := currentDate + pipeIDStr + counterStr
	return OrderReference
}

func main() {
	// Parse flags
	flag.Parse()

	fmt.Println("gRPC server running ....")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSquareoffServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
