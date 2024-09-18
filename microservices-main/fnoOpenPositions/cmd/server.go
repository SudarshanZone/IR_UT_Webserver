//ApplicationServer\Open_Pos\cmd\main.go
/*****************************************************************************************************************/

/*	Program	    	 :- FNO_Open_Pos GRPC Service                                          						 */
/*                                                                            									 */
/*  Input            :-	Required				                                                    			 */
/*	                    FCP_CLM_MTCH_ACCNT                                                          			 */
/*                         																					     */
/*  Output            :  Open Positions                                   										 */
/*                                                                            						             */
/*  Description       : This service retrieves the values of User From Their Unique ID 							 */
/*								 			  			 														 */
/*																												 */
/*                                                                            									 */
/*   Release               :	1.0					Sudarshan Zarkar											 */
/*****************************************************************************************************************/
/* 	1.0		-			New Release                                                    							 */
/*****************************************************************************************************************/
package main

import (
	"net"

	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/logger"

	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/config"
	positions "github.com/krishnakashyap0704/microservices/fnoOpenPositions/generated"
	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/repository"
	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/service"
	"google.golang.org/grpc"
)

func main() {

	//log imported
	logger.InitLogger()
	log := logger.GetLogger()

	//Service
	serviceName := "FnoOpenPosService"

	//Config
	fileName := "config/EnvConfig.ini"

	//Load Config Manager
	cm := &config.ConfigManager{}

	//Load Db from Config Manager
	dbConfig, err := cm.LoadPostgreSQLConfig(serviceName, fileName)
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	//check Connection Status for DB
	connectionStatus := cm.GetDatabaseConnection(serviceName, *dbConfig)
	if connectionStatus != 0 {
		log.Fatalf("Failed to connect to the database")
	}

	//Access DB
	db := cm.GetDB(serviceName)

	//Repository
	repo := &repository.FnoPositionRepository{Db: db}

	//Service
	srv := &service.FnoPositionService{Repo: *repo}

	//Start gRPC server for FNO_Open_Position Service
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	positions.RegisterFnoPositionServiceServer(grpcServer, srv)

	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
