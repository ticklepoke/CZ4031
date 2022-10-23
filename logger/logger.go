package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger - logger instance
var Logger *log.Logger

// InitializeLogger - set log file path
func InitializeLogger(filename string) {
	path := "logs/" + filename + ".log"
	openLogfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	Logger = log.New(openLogfile, filename+":\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitStdoutLogger() {
	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}
