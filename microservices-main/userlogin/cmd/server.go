package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/krishnakashyap0704/microservices/userlogin/generated"
	"github.com/krishnakashyap0704/microservices/userlogin/internal/database"
	logger "github.com/krishnakashyap0704/microservices/userlogin/utils"
)

type server struct {
	pb.UnimplementedUserLoginServiceServer
	db *sql.DB
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	logger.InfoLogger.Printf("Login attempt for user: %s", req.GetUsername())
	var hashedPassword, accessFlag, isAdmin string
	query := "SELECT ops_password, ops_user_access_flag, ops_is_admin FROM ops_user_irra WHERE ops_user_id=$1"

	// Fetch the hashed password, access flag, and admin status
	err := s.db.QueryRow(query, req.GetUsername()).Scan(&hashedPassword, &accessFlag, &isAdmin)
	if err != nil {
		logger.ErrorLogger.Printf("Invalid username: %s", req.GetUsername())
		return &pb.LoginResponse{Success: false, Message: "Invalid username or password"}, nil
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.GetPassword())); err != nil {
		logger.ErrorLogger.Printf("Invalid password for user: %s", req.GetUsername())
		return &pb.LoginResponse{Success: false, Message: "Invalid username or password"}, nil
	}

	// Check if the user has access
	if accessFlag != "Y" {
		logger.InfoLogger.Printf("Access denied for user: %s", req.GetUsername())
		return &pb.LoginResponse{Success: false, Message: "Access denied. Contact your administrator."}, nil
	}

	// Log the login event
	_, err = s.db.Exec("INSERT INTO userlog (username, login_time) VALUES ($1, NOW())", req.GetUsername())
	if err != nil {
		logger.ErrorLogger.Printf("Failed to log login event for user: %s, error: %v", req.GetUsername(), err)
		return nil, fmt.Errorf("failed to log login event: %v", err)
	}
	logger.InfoLogger.Printf("Login event successfully logged for user: %s", req.GetUsername())

	// Determine user type (admin or operational)
	var userType string
	if isAdmin == "Y" {
		userType = "Admin"
	} else {
		userType = "Operational"
	}

	logger.InfoLogger.Printf("User %s logged in as %s", req.GetUsername(), userType)
	return &pb.LoginResponse{Success: true, Message: "Login successful", UserType: userType}, nil
}

func main() {

	// Initialize the logger
	logger.InitLogger()

	// Initialize the database connection
	database.Init()
	db := database.GetDB()
	defer db.Close()

	// Check if the db is nil
	if db == nil {
		logger.ErrorLogger.Fatal("database connection is nil")
		log.Fatalf("database connection is nil")
	} else {
		logger.InfoLogger.Println("Database connection is successfully initialized")
		log.Println("Database connection is successfully initialized")
	}

	// Test the database connection
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error querying database: %v", err)
		log.Fatalf("Error querying database: %v", err)
	}
	logger.InfoLogger.Printf("Database version: %s", version)
	log.Println("Database version:", version)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.ErrorLogger.Fatalf("failed to listen: %v", err)
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userLogService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterUserLoginServiceServer(s, userLogService)
	reflection.Register(s)

	logger.InfoLogger.Println("gRPC server is running on port 8089")
	log.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		logger.ErrorLogger.Fatalf("failed to serve: %v", err)
		log.Fatalf("failed to serve: %v", err)
	}
}
