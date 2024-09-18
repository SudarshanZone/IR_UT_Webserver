package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/krishnakashyap0704/microservices/fnoSquareOff/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Initialize gRPC connection
	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSquareOffServiceClient(conn)

	// Struct to handle JSON input from frontend
	type SquareOffRequest struct {
		FcpDetails []*pb.FnoData `json:"FcpDetails"`
	}

	// Handler function to process square off orders
	handleSquareOff := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			var reqBody SquareOffRequest

			// Extract user ID and product type from query parameters
			FFO_USR_ID_Str := c.Query("FFO_USR_ID")
			fmt.Println("Required ID: ", FFO_USR_ID_Str)
			if FFO_USR_ID_Str == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "FFO_USR_ID is required"})
				return
			}

			// Extract user ID and product type from query parameters
			FFO_USR_ID, err := strconv.ParseInt(c.Query("FFO_USR_ID"), 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FFO_USR_ID", "details": err.Error()})
				return
			}
			FFO_PRDCT_TYP := c.Query("FFO_PRDCT_TYP")

			// Bind JSON input to struct
			if err := c.ShouldBindJSON(&reqBody); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input", "details": err.Error()})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			// Create the gRPC request
			req := &pb.SquareOffRequest{
				FFO_USR_ID:    FFO_USR_ID,
				FFO_PRDCT_TYP: FFO_PRDCT_TYP,
				FcpDetails:    reqBody.FcpDetails,
			}

			// Call the SquareOffOrder RPC method
			resp, err := client.SquareOffOrder(ctx, req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "gRPC square off order failed", "details": err.Error()})
				return
			}

			// Return the response to the client
			c.JSON(http.StatusOK, gin.H{"status": resp.Status, "message": resp.Message})
		}
	}

	// Define route for square off orders
	router.POST("/squareoff/order", handleSquareOff())

	log.Println("Client server is running on :8090")
	router.Run(":8090")
}
