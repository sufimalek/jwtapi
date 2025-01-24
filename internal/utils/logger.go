package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout and a file
	logFile, err := os.OpenFile("/var/log/jwtapi/jwtapi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("Failed to open log file:", err)
	}

	// Write logs to both stdout and the file
	Log.SetOutput(logFile)

	// Only log the warning severity or above.
	Log.SetLevel(logrus.InfoLevel)
}
