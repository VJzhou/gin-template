package logx

import (
	"fmt"
	"gin-template/pkg/logx/zapx"
)

func Sync() {
	_ = zapx.GetLogger().Sync()
}

func Debug(msg string) {
	zapx.GetLogger().Debug(msg)
}

func Debugf(format string, v ...interface{}) {
	zapx.GetLogger().Debug(fmt.Sprintf(format, v...))
}

func Info(msg string) {
	zapx.GetLogger().Info(msg)
}

func Infof(format string, v ...interface{}) {
	zapx.GetLogger().Info(fmt.Sprintf(format, v...))
}

func Warn(msg string) {
	zapx.GetLogger().Warn(msg)
}

func Warnf(format string, v ...interface{}) {
	zapx.GetLogger().Warn(fmt.Sprintf(format, v...))
}

func Error(msg string) {
	zapx.GetLogger().Error(msg)
}

func Errorf(format string, v ...interface{}) {
	zapx.GetLogger().Error(fmt.Sprintf(format, v...))
}

func Fatal(msg string) {
	zapx.GetLogger().Fatal(msg)
}

func Fatalf(format string, v ...interface{}) {
	zapx.GetLogger().Fatal(fmt.Sprintf(format, v...))
}
