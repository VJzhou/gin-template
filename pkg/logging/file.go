package logging

import (
	"fmt"
	"gin-demo/pkg/file"
	"gin-demo/pkg/setting"
	"os"
	"time"
)

var (
	LogSavePath,
	LogSaveName,
	LogFileExt,
	TimeFormat string
)

func init () {
	LogSavePath = setting.AppConfig.LogSavePath
	LogSaveName = setting.AppConfig.LogSaveName
	LogFileExt = setting.AppConfig.LogFileExt
	TimeFormat = setting.AppConfig.TimeFormat
}


func getLogFilePath () string {
	return fmt.Sprintf("%s%s", setting.AppConfig.RuntimeRootPath, setting.AppConfig.LogSavePath)
}

func getLogFileName () string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppConfig.LogSaveName,
		time.Now().Format(setting.AppConfig.TimeFormat),
		setting.AppConfig.LogFileExt,
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

	f, err := file.Open(src + filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}
	return f, nil
}

func mkDir() {
	dir, _ := os.Getwd();
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

