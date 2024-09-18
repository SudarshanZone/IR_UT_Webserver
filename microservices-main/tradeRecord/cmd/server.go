package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/krishnakashyap0704/microservices/tradeRecord/generated"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port     = ":5001"
	dbSource = "postgresql://postgres:shoeb@localhost:5432/postgres?sslmode=disable"
)

type server struct {
	pb.UnimplementedDataServiceServer
	db *sql.DB
}

func (s *server) GetCSV(req *pb.CSVDataRequest, stream pb.DataService_GetCSVServer) error {
	log.Printf("Received request: Type=%s, SubType=%s", req.Type, req.SubType)

	var query string
	switch {
	case req.Type == "fno" && req.SubType == "ordr_dtls":
		query = "SELECT * FROM fod_fo_ordr_dtls"
	case req.Type == "fno" && req.SubType == "trd_dtls":
		query = "SELECT * FROM ftq_fo_trd_dtls"
	case req.Type == "commodity" && req.SubType == "ordr_dtls":
		query = "SELECT * FROM cod_cod_ordr_dtls"
	case req.Type == "commodity" && req.SubType == "trd_dtls":
		query = "SELECT * FROM ctd_cod_trd_dtls"
	case req.Type == "equity" && req.SubType == "ordr_dtls":
		query = "SELECT * FROM ord_ordr_dtls"
	case req.Type == "equity" && req.SubType == "trd_dtls":
		query = "SELECT * FROM clm_clnt_mstr where "
	default:
		return fmt.Errorf("invalid type or subtype: Type=%s, SubType=%s", req.Type, req.SubType)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Printf("Failed to get columns: %v", err)
		return err
	}

	log.Printf("Sending columns: %v", columns)
	if err := stream.Send(&pb.CSVDataResponse{Data: csvRowToString(columns)}); err != nil {
		log.Printf("Failed to send columns: %v", err)
		return err
	}

	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return err
		}

		record := make([]string, len(columns))
		for i, value := range values {
			record[i] = string(value)
		}

		if err := stream.Send(&pb.CSVDataResponse{Data: csvRowToString(record)}); err != nil {
			log.Printf("Failed to send row: %v", err)
			return err
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return err
	}

	log.Println("Completed sending all rows")
	return nil
}

func csvRowToString(row []string) string {
	return fmt.Sprintf("%s\n", row)
}

func main() {

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start the gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &server{db: db})
	log.Printf("gRPC server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
