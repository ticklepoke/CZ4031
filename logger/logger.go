package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger - logger instance
var Logger = log.New()

// InitFileLogger - logs to file
func InitFileLogger(filename string, level log.Level) {
	path := "logs/" + filename + ".log"
	openLogfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	Logger.Out = openLogfile
	Logger.Level = level
}

// InitLogger - logs to stdout
func InitStdoutLogger(level log.Level) {
	Logger.Out = os.Stdout
	Logger.Level = level
}
