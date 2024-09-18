package main

import (
	"context"
	"io"
	"log"
	"net/http"

	pb "github.com/krishnakashyap0704/microservices/tradeRecord/internal/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.Default())
	r.GET("/download/:type/:sub_type", func(c *gin.Context) {
		typ := c.Param("type")
		sub_typ := c.Param("sub_type")

		log.Printf("Received download request: type=%s, sub_type=%s", typ, sub_typ)

		conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to gRPC server: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		client := pb.NewDataServiceClient(conn)
		stream, err := client.GetCSV(context.Background(), &pb.CSVDataRequest{Type: typ, SubType: sub_typ})
		if err != nil {
			log.Printf("Failed to get CSV data from gRPC server: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Disposition", "attachment; filename=data.csv")
		c.Header("Content-Type", "text/csv")

		writer := c.Writer
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				log.Println("Completed receiving data from gRPC server")
				break
			}
			if err != nil {
				log.Printf("Error receiving data from gRPC server: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			_, err = writer.Write([]byte(msg.Data))
			if err != nil {
				log.Printf("Error writing data to response: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			log.Printf("Sent data  to client")
		}

	})

	log.Println("Starting HTTP server on port 8000")
	r.Run(":8000")
}
