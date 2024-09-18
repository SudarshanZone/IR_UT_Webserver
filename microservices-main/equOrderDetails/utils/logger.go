package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// InitLogger initializes the logger and ensures the directory structure for logs is created.
func InitLogger() {
	// Get the current date for file naming
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	// Define log directory
	logDir := filepath.Join("../", "logs", "fnoSquareOff")

	// Create the necessary directories (only the fnoSquareOff directory now)
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		logrus.Fatalf("error creating log directories: %v", err)
	}

	// Define log file paths for info, error, and debug logs (date included in file names)
	infoLogPath := filepath.Join(logDir, "fnoSquareOff-"+date+"-info.log")
	errorLogPath := filepath.Join(logDir, "fnoSquareOff-"+date+"-error.log")
	debugLogPath := filepath.Join(logDir, "fnoSquareOff-"+date+"-debug.log")

	// Open the info log file
	infoFile, err := os.OpenFile(infoLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("error opening info log file: %v", err)
	}

	// Open the error log file
	errorFile, err := os.OpenFile(errorLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("error opening error log file: %v", err)
	}

	// Open the debug log file
	debugFile, err := os.OpenFile(debugLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("error opening debug log file: %v", err)
	}

	// Set Logrus output for general info logs
	logrus.SetOutput(infoFile)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Info("Service started and logging initialized")

	// Create custom loggers for info, error, and debug logs
	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(debugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
