package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/krishnakashyap0704/microservices/equOrderDetails/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Parse command-line flags
	flag.Parse()
	log.Println("Starting application...")

	// Initialize Gin router
	router := gin.Default()
	log.Println("Gin router initialized.")

	// CORS middleware
	router.Use(func(c *gin.Context) {
		log.Println("Processing CORS middleware...")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			log.Println("OPTIONS request received, returning 200")
			c.AbortWithStatus(200)
			return
		}
		c.Next()
		log.Println("CORS middleware processing complete.")
	})

	// Initialize gRPC connection
	log.Println("Attempting to connect to gRPC server at localhost:8089")
	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	log.Println("Connected to gRPC server.")
	defer func() {
		log.Println("Closing gRPC connection")
		conn.Close()
	}()

	// Create gRPC client
	client := pb.NewOrderServiceClient(conn)
	log.Println("gRPC client created.")

	// Define the GET route for positions
	router.GET("/OrderService/:ord_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("ord_clm_mtch_accnt")
		log.Printf("Received request for account: %s", accnt)

		// Set a timeout for the gRPC call
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Call the gRPC method
		log.Printf("Calling gRPC method GetOrder for account: %s", accnt)
		res, err := client.GetOrder(ctxWithTimeout, &pb.OrderRequest{OrdClmMtchAccnt: accnt})
		if err != nil {
			log.Printf("Error fetching order details for account %s: %v", accnt, err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Error fetching order details: %v", err),
			})
			return
		}

		// Process the gRPC response
		if res.OrdDtls != nil {
			log.Printf("Successfully retrieved order details for account: %s", accnt)
			ctx.JSON(http.StatusOK, gin.H{
				"equity_order_details": res.OrdDtls,
			})
		} else {
			log.Printf("No order details found for account: %s", accnt)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Equity order details not found",
			})
		}
	})

	// Start the server
	log.Println("Starting HTTP server on :8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
