package config

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"time"
)

// Creates a log file name with UUID-style ID
func GenerateLogFileName(prefix string) string {
	if prefix == "" {
		prefix = "req"
	}
	timestamp := time.Now().Format("02-01-2006_15-04")
	id := fmt.Sprintf("%08x", rand.Uint32()) // 8-character hex ID
	return fmt.Sprintf("%s_%s_%s.log", prefix, timestamp, id)
}

// Custom logger for specific actions
type Logger struct {
	*log.Logger
}

func NewFileWriter() io.Writer {
	// Create logs directory if it doesn't exist
	os.Mkdir("logs/actions", 0755)
	// Create log file
	logFileName := GenerateLogFileName("act")
	f, _ := os.Create(fmt.Sprintf("logs/actions/%s", logFileName))
	multiWriter := io.MultiWriter(f)

	return multiWriter
}

var actionLogger = slog.New(slog.NewTextHandler(NewFileWriter(), &slog.HandlerOptions{Level: slog.LevelInfo}))

func GetLogger() *slog.Logger {
	return actionLogger
}
