package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = ".log"
	TimeFormat = "20060102"
)

func getLogFilePath () string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath () string {
	prefixPath := getLogFilePath()
	filename := fmt.Sprintf("%s%s%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, filename)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to Openfile %v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd();
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

