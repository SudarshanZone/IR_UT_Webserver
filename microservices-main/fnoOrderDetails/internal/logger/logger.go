package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()

	logFilename := filepath.Join("logs", "ULOG."+time.Now().Format("2006-01-02")+".log")

	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilename,
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	log.SetOutput(getLogOutput(lumberjackLogger))
	log.SetFormatter(&CustomFormatter{})
	log.SetLevel(getLogLevel())
	log.AddHook(NewContextHook())
}

func GetLogger() *logrus.Logger {
	return log
}

func getLogOutput(lumberjackLogger *lumberjack.Logger) io.Writer {
	if os.Getenv("ENV") == "production" {
		return lumberjackLogger
	}
	return io.MultiWriter(lumberjackLogger, os.Stdout)
}

func getLogLevel() logrus.Level {
	switch os.Getenv("ENV") {
	case "production":
		return logrus.WarnLevel
	case "staging":
		return logrus.InfoLevel
	default:
		return logrus.DebugLevel
	}
}

type ContextHook struct{}

func NewContextHook() *ContextHook {
	return &ContextHook{}
}

func (hook *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	if _, file, line, ok := runtime.Caller(6); ok {
		shortFile := filepath.Base(file)
		entry.Data["file"] = shortFile
		entry.Data["line"] = line
	}
	return nil
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	serviceName, _ := entry.Data["service"].(string)
	message := entry.Message

	// Format with timestamp, service name, and message
	formatted := timestamp + " " + serviceName + ": " + message + "\n"

	return []byte(formatted), nil
}

