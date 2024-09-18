package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	pb "github.com/SudarshanZone/Fno_Ord_Dtls/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger = logrus.New()

func main() {
	flag.Parse()
	router := gin.Default()

	// Initialize logrus
	logger.Out = os.Stdout
	logger.SetLevel(logrus.InfoLevel)

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Initialize gRPC connection with a goroutine for graceful shutdown
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderDetailsServiceClient(conn)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		logger.Info("Shutting down gracefully...")
		conn.Close()
		os.Exit(0)
	}()

	router.GET("/getOrderDetails/:OrderID", func(c *gin.Context) {
		OrderID := c.Param("OrderID")

		// Log incoming request
		logger.WithFields(logrus.Fields{
			"endpoint": "/getOrderDetails",
			"OrderID":  OrderID,
		}).Info("Received request")

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Create request object for gRPC
		req := &pb.OrderDetailsRequest{OrderId: OrderID}

		// Call gRPC service
		resp, err := client.GetOrderDetails(ctx, req)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"endpoint": "/getOrderDetails",
				"OrderID":  OrderID,
				"error":    err.Error(),
			}).Error("Error fetching order details")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
			return
		}

		// Log successful response
		logger.WithFields(logrus.Fields{
			"endpoint": "/getOrderDetails",
			"OrderID":  OrderID,
		}).Info("Successfully fetched order details")

		c.JSON(http.StatusOK, gin.H{"orderDetails": resp.OrderDetails})
	})

	logger.Info("Client server is running on :8090")
	if err := router.Run(":8090"); err != nil {
		logger.Fatalf("Failed to run server: %v", err)
	}

	// Wait for all goroutines to finish before exiting
	wg.Wait()
}
