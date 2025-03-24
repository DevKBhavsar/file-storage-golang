package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance
var Logger = logrus.New()

// InitLogger initializes the logger
func InitLogger() {
	// Set log format to JSON for better structure
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Create log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}
}
