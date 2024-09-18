package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/krishnakashyap0704/microservices/userlogin/generated"
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

	client := pb.NewUserLoginServiceClient(conn)

	router.POST("/login", func(c *gin.Context) {
		var formData struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&formData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.Login(ctx, &pb.LoginRequest{
			Username: formData.Username,
			Password: formData.Password,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": "gRPC login failed", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{"success": resp.Success, "message": resp.Message, "userType": resp.UserType})
	})

	log.Println("Client server is running on :8090")
	router.Run(":8090")
}
