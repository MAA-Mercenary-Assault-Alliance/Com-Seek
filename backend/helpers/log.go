package helpers

import (
	"com-seek/backend/helpers"
	"fmt"
	"io"
	"log"
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

func NewLogger() *Logger {
	// Create log file
	logFileName := helpers.GenerateLogFileName("act")
	f, _ := os.Create(fmt.Sprintf("logs/actions/%s", logFileName))
	multiWriter := io.MultiWriter(f)

	return &Logger{
		Logger: log.New(multiWriter, "[ACTION] ", log.LstdFlags),
	}
}

func (l *Logger) LogAction(action, userID, details string) {
	l.Printf("Action: %s | UserID: %s | Details: %s", action, userID, details)
}
