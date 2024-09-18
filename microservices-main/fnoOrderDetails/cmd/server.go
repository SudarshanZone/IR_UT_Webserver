/*****************************************************************************************************************/
/*	Program	    	 :- FNO_Ord_Dtls GRPC Service                                          						 */
/*                                                                            									 */
/*  Input            :-	Required				                                                    			 */
/*	                    FOD_CLM_MTCH_ACCNT                                                          			 */
/*                         																					     */
/*  Output            :  Orders Based on Unique ID                                   							 */
/*                                                                            						             */
/*  Description       : This service retrieves the Orders and Their Curent State 							     */
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

	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/config"
	pb "github.com/krishnakashyap0704/microservices/fnoOrderDetails/generated"
	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/internal/logger"
	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/internal/repository"
	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/internal/service"
	"google.golang.org/grpc"
)

func main() {

	logger.InitLogger()
	log := logger.GetLogger()

	serviceName := "main"
	fileName := "config/EnvConfig.ini"

	cm := &config.ConfigManager{}

	// Load PostgreSQL configuration
	dbConfig, err := cm.LoadPostgreSQLConfig(serviceName, fileName)
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Get database connection
	connectionStatus := cm.GetDatabaseConnection(serviceName, *dbConfig)
	if connectionStatus != 0 {
		log.Fatalf("Failed to connect to the database")
	}

	db := cm.GetDB(serviceName)

	// Initialize the repository with the database connection
	repo := repository.OrderDetailsRepository{Db: db}

	// Initialize the service with the repository
	orderDetailsService := &service.OrderDetailsService{Repo: repo}

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterOrderDetailsServiceServer(grpcServer, orderDetailsService)

	// Listen on a TCP port
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
