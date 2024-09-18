package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"

	pb "github.com/krishnakashyap0704/microservices/comOrderDetails/generated"
	"github.com/krishnakashyap0704/microservices/comOrderDetails/internal/database"
	logger "github.com/krishnakashyap0704/microservices/comOrderDetails/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func getString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "" // or return a default value
}

func getInt32(ni sql.NullInt32) int32 {
	if ni.Valid {
		return ni.Int32
	}
	return 0 // or return a default value
}

type server struct {
	p2 []database.ComOrdrDtls // p2 is now a slice
	pb.UnimplementedC_OrderServiceServer
	db *sql.DB
}

func (s *server) GetComOrder(ctx context.Context, req *pb.ComOrderRequest) (*pb.ComOrderResponse, error) {
	logger.InfoLogger.Printf("Received GetOrder request for COD_CLM_MTCH_ACCNT: %s", req.GetCodClmMtchAccnt())

	var query string
	var args []interface{}
	var rows *sql.Rows
	var err error

	logger.InfoLogger.Printf("Processing commodity order details for user: %s", req.GetCodClmMtchAccnt())

	// Prepare the SQL query
	query = `SELECT
    a.COD_CLM_MTCH_ACCNT,
    CONCAT(a.COD_PRDCT_TYP,'-', a.COD_UNDRLYNG, '-', a.COD_EXPRY_DT) AS COD_UNDRLYNG,
    a.COD_LMT_RT,
    a.COD_ORDR_VALID_DT,
    a.COD_ORDR_FLW,
    a.COD_ORDR_TOT_QTY,
    a.COD_ORDR_STTS
FROM
    COD_COD_ORDR_DTLS a
WHERE
    a.COD_CLM_MTCH_ACCNT = $1;`

	args = []interface{}{req.GetCodClmMtchAccnt()}
	logger.InfoLogger.Printf("Executing query: %s", query)
	logger.InfoLogger.Printf("With parameters: %v", args)

	// Execute the query
	rows, err = s.db.Query(query, args...)
	if err != nil {
		logger.ErrorLogger.Printf("Error querying Commodity order details: %v", err)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.ErrorLogger.Printf("Error closing rows: %v", err)
		}
	}()

	// Initialize the slice
	s.p2 = []database.ComOrdrDtls{}
	logger.InfoLogger.Println("Initialized empty slice for order details.")

	logger.InfoLogger.Println("Starting to iterate over rows.")
	// Iterate over the rows and scan each row into the struct
	for rows.Next() {
		var detail database.ComOrdrDtls
		var codOrdrFlw []byte // Create a variable to hold the scanned data
		if err := rows.Scan(
			&detail.COD_CLM_MTCH_ACCNT,
			&detail.COD_UNDRLYNG,
			&detail.COD_LMT_RT,
			&detail.COD_ORDR_VALID_DT,
			&codOrdrFlw, // Scan into the byte slice
			&detail.COD_ORDR_TOT_QTY,
			&detail.COD_ORDR_STTS,
		); err != nil {
			log.Printf("Error scanning Commodity order details: %v", err)
			return nil, err
		}

		// Convert byte slice to sql.NullString for COD_ORDR_FLW
		detail.COD_ORDR_FLW = sql.NullString{
			String: string(codOrdrFlw),
			Valid:  len(codOrdrFlw) > 0, // Set Valid if the slice is non-empty
		}

		logger.InfoLogger.Printf("Scanned detail: %+v", detail)
		// Append the scanned details to the slice
		s.p2 = append(s.p2, detail)
	}

	if len(s.p2) == 0 {
		log.Println("Commodity order details not found.")
		return nil, errors.New("commodity order details not found")
	}

	logger.InfoLogger.Printf("Successfully retrieved %d order details.", len(s.p2))

	// Convert s.p2 to the correct type
	var ordDtls []*pb.ComOrdrDtls
	for _, detail := range s.p2 {
		ordDtls = append(ordDtls, &pb.ComOrdrDtls{
			CodClmMtchAccnt: getString(detail.COD_CLM_MTCH_ACCNT),
			CodUndrlyng:     getString(detail.COD_UNDRLYNG),
			CodOrdrValidDt:  getString(detail.COD_ORDR_VALID_DT),
			CodOrdrStts:     getString(detail.COD_ORDR_STTS),
			CodLmtRt:        float32(detail.COD_LMT_RT.Float64),
			CodOrdrFlw:      getString(detail.COD_ORDR_FLW),
			CodOrdrTotQty:   getInt32(detail.COD_ORDR_TOT_QTY),
		})
	}

	logger.InfoLogger.Printf("Returning response with %d order details.", len(ordDtls))
	// Return the response with all the commodity order details data
	return &pb.ComOrderResponse{
		OrdDtls: ordDtls, // Use the converted slice here
	}, nil
}

func main() {

	// Initialize the database connection
	logger.InfoLogger.Println("Initializing database connection...")
	database.Init()
	db := database.GetDB()
	defer db.Close()

	// Check if the db is nil
	if db == nil {
		logger.ErrorLogger.Fatalf("Database connection is nil.")
	} else {
		logger.ErrorLogger.Println("Database connection is successfully initialized.")
	}

	// Test the database connection
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error querying database: %v", err)
	}
	logger.InfoLogger.Println("Database version:", version)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.InfoLogger.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	C_OrderService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterC_OrderServiceServer(s, C_OrderService)
	reflection.Register(s)

	logger.InfoLogger.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		logger.InfoLogger.Fatalf("Failed to serve: %v", err)
	}
}
