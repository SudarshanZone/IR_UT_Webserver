package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"

	pb "github.com/krishnakashyap0704/microservices/equOrderDetails/generated"
	"github.com/krishnakashyap0704/microservices/equOrderDetails/internal/database"
	logger "github.com/krishnakashyap0704/microservices/equOrderDetails/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedEquityOrderServiceServer
	db *sql.DB
}

func (s *server) GetEquityOrder(ctx context.Context, req *pb.EquityOrderRequest) (*pb.EquityOrderResponse, error) {
	log.Printf("Received GetOrder request for account: %s", req.GetOrdClmMtchAccnt())
	logger.InfoLogger.Println("Starting equity open position service")

	var query string
	var args []interface{}
	var rows *sql.Rows
	var err error

	log.Println("Preparing SQL query for fetching equity order details")
	query = `SELECT Ord_Clm_Mtch_Accnt, Ord_Stck_Cd, Ord_Ordr_Dt, Ord_Ordr_Flw, Ord_Ordr_Qty, Ord_Lmt_Rt, Ord_Ordr_Stts FROM ORD_ORDR_DTLS WHERE ord_clm_mtch_accnt = $1`
	args = []interface{}{req.GetOrdClmMtchAccnt()}

	// Execute the query
	log.Println("Executing SQL query")
	rows, err = s.db.Query(query, args...)
	if err != nil {
		//log.Printf("Error querying equity order details for account %s: %v", req.GetOrdClmMtchAccnt(), err)
		logger.ErrorLogger.Printf("Error querying equity order details for account %s: %v", req.GetOrdClmMtchAccnt(), err)
		return nil, err
	}
	defer func() {
		//log.Println("Closing rows after processing")
		logger.ErrorLogger.Printf("Closing rows after processing")
		rows.Close()
	}()

	var ordDtls []*pb.EquityOrderDetails

	// Iterate over the rows and scan each row into the struct
	//log.Println("Iterating over result rows")
	logger.ErrorLogger.Printf("Iterating over result rows")

	for rows.Next() {
		var p2 database.EquityOrderDetails
		if err := rows.Scan(
			&p2.Ord_Clm_Mtch_Accnt,
			&p2.Ord_Stck_Cd,
			&p2.Ord_Ordr_Dt,
			&p2.Ord_Ordr_Flw,
			&p2.Ord_Ordr_Qty,
			&p2.Ord_Lmt_Rt,
			&p2.Ord_Ordr_Stts,
		); err != nil {
			//log.Printf("Error scanning equity order details for account %s: %v", req.GetOrdClmMtchAccnt(), err)
			logger.ErrorLogger.Printf("Error scanning equity order details for account %s: %v", req.GetOrdClmMtchAccnt(), err)
			return nil, err
		}

		// Append the details to the slice
		ordDtls = append(ordDtls, &pb.EquityOrderDetails{
			OrdClmMtchAccnt: p2.Ord_Clm_Mtch_Accnt.String,
			OrdStckCd:       p2.Ord_Stck_Cd.String,
			OrdOrdrDt:       p2.Ord_Ordr_Dt.String,
			OrdOrdrFlw:      p2.Ord_Ordr_Flw.String,
			OrdOrdrQty:      p2.Ord_Ordr_Qty.Int32,
			OrdLmtRt:        p2.Ord_Lmt_Rt.Float64,
			OrdOrdrStts:     p2.Ord_Ordr_Stts.String,
		})
	}

	//log.Printf("Fetched %d order details for account: %s", len(ordDtls), req.GetOrdClmMtchAccnt())
	logger.ErrorLogger.Printf("Fetched %d order details for account: %s", len(ordDtls), req.GetOrdClmMtchAccnt())

	// Check if no rows were found
	if len(ordDtls) == 0 {
		//log.Printf("No equity order details found for account: %s", req.GetOrdClmMtchAccnt())
		logger.ErrorLogger.Printf("No equity order details found for account: %s", req.GetOrdClmMtchAccnt())
		return nil, errors.New("equity order details not found")
	}

	// Return the response with all the equity order details data
	//log.Printf("Returning order details response for account: %s", req.GetOrdClmMtchAccnt())
	logger.ErrorLogger.Printf("Returning order details response for account: %s", req.GetOrdClmMtchAccnt())

	return &pb.EquityOrderResponse{
		OrdDtls: ordDtls,
	}, nil
}

func main() {

	// Initialize the database connection
	log.Println("Initializing database connection")
	database.Init()
	db := database.GetDB()
	defer func() {
		log.Println("Closing database connection")
		db.Close()
	}()

	// Check if the db is nil
	if db == nil {
		//log.Fatalf("Database connection is nil")
		logger.InfoLogger.Printf("Database connection is nil")
	} else {
		//log.Println("Database connection is successfully initialized")
		logger.InfoLogger.Printf("Database connection is successfully initialized")
	}

	// Test the database connection

	logger.InfoLogger.Printf("Testing database connection by querying version")
	//log.Println("Testing database connection by querying version")
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {

		logger.InfoLogger.Printf("Error querying database: %v", err)
		//log.Fatalf("Error querying database: %v", err)
	}
	//log.Println("Database version:", version)
	logger.InfoLogger.Printf("Database version:", version)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.ErrorLogger.Printf("Failed to listen: %v", err)
		//log.Fatalf("Failed to listen: %v", err)
	}

	//log.Println("Starting gRPC server on port 8089")
	logger.InfoLogger.Printf("Starting gRPC server on port 8089")

	s := grpc.NewServer()
	OrderService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterEquityOrderServiceServer(s, OrderService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		//log.Fatalf("Failed to serve: %v", err)
		logger.ErrorLogger.Printf("Failed to serve: %v", err)
	}
}
