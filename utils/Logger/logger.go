package Logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

const logName = "log.txt"

func init() {
	absPath, err := filepath.Abs("../../logs")
	if err != nil {
		fmt.Print("problem in finding logs path")
	}

	logFile, err := os.OpenFile(absPath + "\\" +logName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("error opening Logger file:", err)
		os.Exit(1)
	}

	GeneralLogger = log.New(logFile, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
