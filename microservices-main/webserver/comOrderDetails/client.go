package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/krishnakashyap0704/microservices/comOrderDetails/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Initialize gRPC connection
	log.Println("Initializing gRPC connection to server at localhost:8089...")
	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer func() {
		log.Println("Closing gRPC connection...")
		conn.Close()
	}()

	client := pb.NewC_OrderServiceClient(conn)
	log.Println("gRPC client successfully created.")

	// Define the GET route for positions
	router.GET("/C_OrderService/:cod_clm_mtch_accnt", func(ctx *gin.Context) {
		accnt := ctx.Param("cod_clm_mtch_accnt")
		log.Printf("Received request for account: %s", accnt)

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Calling GetOrder on gRPC client...")
		res, err := client.GetOrder(ctxWithTimeout, &pb.OrderRequest{CodClmMtchAccnt: accnt})
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

	// Start the server
	log.Println("Client server is running on :8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
