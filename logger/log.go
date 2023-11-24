package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func LogFileInit(typeOf string, id ...int32) {
	var folder string
	if typeOf == "main" {
		folder = "log"
	} else {
		folder = fmt.Sprintf("log/%vlog", typeOf)
	}

	if checkIfExist(folder) {
		ClearLog(folder)
	}
	MakeLogFolder(folder)

	ClearLog(folder) // Clear the existing directory if it exists

	MakeLogFolder(folder) // Create the log folder

	var file *os.File
	var err error
	if typeOf == "main" {
		file, err = os.Create(fmt.Sprintf("%v/log.txt", folder))
	} else {
		file, err = os.Create(fmt.Sprintf("%v/%v%d.txt", folder, typeOf, id[0]))
	}
	if err != nil {
		log.Fatal(err)
	}

	initLoggerMessages(file)
}

func initLoggerMessages(file *os.File) {
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ClearLog(name string) {
	if checkIfExist(name) {
		os.RemoveAll(name)
	}
}

func checkIfExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func MakeLogFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, os.ModePerm)
	}
}
