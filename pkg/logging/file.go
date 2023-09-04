package logging

import (
	"fmt"
	"gin-demo/conf"
	"gin-demo/pkg/file"
	"os"
	"time"
)

var (
	LogSavePath,
	LogSaveName,
	LogFileExt,
	TimeFormat string
)

func init() {
	LogSavePath = conf.AppConfig.LogSavePath
	LogSaveName = conf.AppConfig.LogSaveName
	LogFileExt = conf.AppConfig.LogFileExt
	TimeFormat = conf.AppConfig.TimeFormat
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", conf.AppConfig.RuntimeRootPath, conf.AppConfig.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		conf.AppConfig.LogSaveName,
		time.Now().Format(conf.AppConfig.TimeFormat),
		conf.AppConfig.LogFileExt,
	)
}

func openLogFile(filename, filePath string) (*os.File, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s , err: %v", src, err)
	}

	f, err := file.Open(src+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}
	return f, nil
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
