package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/krishnakashyap0704/microservices/equOpenPositions/generated"
	"github.com/krishnakashyap0704/microservices/equOpenPositions/internal/database"
	logger "github.com/krishnakashyap0704/microservices/equOpenPositions/utils"

	_ "github.com/lib/pq" // Import PostgreSQL driver
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"time"

	"google.golang.org/grpc/keepalive"
)

type server struct {
	pb.UnimplementedPositionServiceServer
	db *gorm.DB
}

func (s *server) GetPosition(ctx context.Context, req *pb.PositionRequest) (*pb.PositionResponse, error) {
	fmt.Println("Read Positions", req.GetEpbClmMtchAccnt())
	logger.InfoLogger.Println("Starting equity open positions")
	// Prepare the SQL query
	query := `SELECT * FROM epb_em_pstn_book WHERE epb_clm_mtch_accnt = $1`

	// Execute the query
	//row := s.db.Raw(query, req.GetCcpClmMtchAccnt()) <------------------------

	// Prepare a variable to hold the data
	var tempPositions []pb.EquityPosition
	result := s.db.Raw(query, req.GetEpbClmMtchAccnt()).Scan(&tempPositions)

	// Scan the result into the Commoditypositions struct
	//result := row.Scan(&p1)							<------------------------

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("Equity positions not found")
			logger.ErrorLogger.Printf("Equity positions not found")
			return nil, errors.New("Equity positions not found")
		}

		fmt.Println("Error querying equity positions:", result.Error)
		logger.ErrorLogger.Printf("Error querying equity positions:", result.Error)
		return nil, result.Error
	}
	positions := make([]*pb.EquityPosition, len(tempPositions))
	for i := range tempPositions {
		positions[i] = &tempPositions[i]
	}

	response := &pb.PositionResponse{
		Equity: positions,
	}

	return response, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.ErrorLogger.Printf("Failed to listen")
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*50), // 16 MB
		grpc.MaxSendMsgSize(1024*1024*50), // 16 MB
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
			Timeout:           20 * time.Second,
			MaxConnectionAge:  10 * time.Minute,
		}),
	)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	// Initialize the database connection
	database.Init()

	// Get the GORM database instance
	db := database.GetDB()
	if db == nil {
		logger.ErrorLogger.Printf("Failed to get database instance")
		log.Fatal("Failed to get database instance")
	}

	// Setup gRPC server
	// lis, err := net.Listen("tcp", ":8089")
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	PositionService := &server{db: db} // Pass the GORM database connection to the server struct
	pb.RegisterPositionServiceServer(grpcServer, PositionService)
	reflection.Register(grpcServer)
	logger.ErrorLogger.Printf("gRPC server is running on port 8089")
	//log.Println("gRPC server is running on port 8089")

	if err := grpcServer.Serve(lis); err != nil {
		logger.ErrorLogger.Printf("Failed to serve: %v", err)
		log.Fatalf("Failed to serve: %v", err)
	}
}
