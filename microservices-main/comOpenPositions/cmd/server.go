package main

import (
	"context"
	"errors"

	// "fmt"
	"log"
	"net"

	pb "github.com/krishnakashyap0704/microservices/comOpenPositions/generated"
	"github.com/krishnakashyap0704/microservices/comOpenPositions/internal/database"
	logger "github.com/krishnakashyap0704/microservices/comOpenPositions/utils"

	_ "github.com/lib/pq" // Import PostgreSQL driver
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCCPServiceServer
	db *gorm.DB
}

func (s *server) GetCCPData(ctx context.Context, req *pb.CCPRequest) (*pb.CCPResponse, error) {
	logger.InfoLogger.Println("Displaying Commodity Open Positions")

	logger.InfoLogger.Printf("Received request for account: %s", req.GetCcpClmMtchAccnt())

	// Prepare the SQL query
	logger.InfoLogger.Printf("Fetching Processing from ccp_cod_spn_cntrct_pstn")
	query := `
		SELECT ccp_clm_mtch_accnt, ccp_xchng_cd, CONCAT(ccp_prdct_typ, '_', ccp_undrlyng, '_', ccp_expry_dt) AS ccp_undrlyng, 
			ccp_prdct_typ, ccp_indstk, ccp_undrlyng, ccp_expry_dt, ccp_exer_typ, ccp_strk_prc, ccp_opt_typ, ccp_ibuy_qty, 
			ccp_ibuy_ord_val, ccp_isell_qty, ccp_isell_ord_val, ccp_exbuy_qty, ccp_exbuy_ord_val, ccp_exsell_qty, 
			ccp_exsell_ord_val, ccp_buy_exctd_qty, ccp_sell_exctd_qty, ccp_opnpstn_flw, ccp_opnpstn_qty, ccp_opnpstn_val, 
			ccp_exrc_qty, ccp_asgnd_qty, ccp_opt_premium, ccp_mtm_opn_val, ccp_imtm_opn_val, ccp_extrmloss_mrgn_extra, 
			ccp_addnl_mrgn, ccp_spcl_mrgn, ccp_tndr_mrgn, ccp_dlvry_mrgn, ccp_extrm_min_loss_mrgn, ccp_mtm_flg, 
			ccp_extrm_loss_mrgn, ccp_flat_val_mrgn, ccp_trg_prc, ccp_min_trg_prc, ccp_devolmnt_mrgn, ccp_mtmsq_ordcnt, 
			ccp_avg_prc 
		FROM ccp_cod_spn_cntrct_pstn 
		WHERE ccp_clm_mtch_accnt = $1 AND ccp_opnpstn_flw <> 'N'`

	logger.InfoLogger.Printf("Executing query: %s", query)
	logger.InfoLogger.Printf("With parameter: %s", req.GetCcpClmMtchAccnt())

	// Execute the query
	var tempPositions []pb.Commoditypositions
	result := s.db.Raw(query, req.GetCcpClmMtchAccnt()).Scan(&tempPositions)

	// Handle potential errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.ErrorLogger.Println("Commodity positions not found")
			return nil, errors.New("commodity positions not found")
		}
		logger.ErrorLogger.Printf("Error querying commodity positions: %v", result.Error)
		return nil, result.Error
	}

	// Prepare the response
	positions := make([]*pb.Commoditypositions, len(tempPositions))
	for i := range tempPositions {
		positions[i] = &tempPositions[i]
	}

	response := &pb.CCPResponse{
		Commo: positions,
	}
	logger.InfoLogger.Printf("response: %d", response)
	logger.InfoLogger.Printf("Returning %d positions for account: %s", len(positions), req.GetCcpClmMtchAccnt())
	return response, nil
}

func main() {
	// Initialize the database connection
	//log.Println("Initializing database connection...")
	logger.InitLogger()
	database.Init()

	// Get the GORM database instance
	db := database.GetDB()

	if db == nil {
		logger.ErrorLogger.Printf("Database connection is nil")
	} else {
		logger.InfoLogger.Println("Database connection is successfully initialized")
	}

	// Setup gRPC server
	log.Println("Setting up gRPC server...")

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	CCPService := &server{db: db}
	pb.RegisterCCPServiceServer(s, CCPService)
	reflection.Register(s)
	logger.InfoLogger.Println("gRPC server is running on port 8089")

	if err := s.Serve(lis); err != nil {
		logger.ErrorLogger.Fatalf("Failed to serve gRPC server: %v", err)
		log.Fatalf("Failed to serve: %v", err)
	}
}
