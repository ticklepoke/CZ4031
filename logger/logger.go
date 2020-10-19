package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger - logger instance
var Logger *log.Logger

// InitlizeLogger - set log file path
func InitlizeLogger(filename string) {
	path := "logs/" + filename + ".log"
	openLogfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	Logger = log.New(openLogfile, filename+":\t", log.Ldate|log.Ltime|log.Lshortfile)
}
