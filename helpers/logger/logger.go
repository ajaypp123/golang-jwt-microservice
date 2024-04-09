package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	// Create a new logger instance
	Logger = log.New()

	/*
		// Log to a file
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Failed to open log file:", err)
		}
		defer file.Close()
		Logger.SetOutput(file)
	*/

	// Log to stdout by default
	Logger.SetOutput(os.Stdout)

	// Use JSON formatter for structured logging
	Logger.SetFormatter(&log.JSONFormatter{})
}
