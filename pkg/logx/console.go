package logx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var _ Encoder = (*ConsoleFormatter)(nil)

type ConsoleFormatter struct {
	raw string
}

func NewConsoleFormatter() *ConsoleFormatter {
	return &ConsoleFormatter{}
}

// Config 自定义配置.
func (cf *ConsoleFormatter) Config() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	// 时间格式自定义
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	}
	// 打印路径自定义
	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + getFilePath(&caller) + "]")
	}
	// 级别显示自定义
	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + level.String() + "]")
	}
	return zapcore.NewConsoleEncoder(config)
}

// WithKey 添加单个键.
func (cf *ConsoleFormatter) WithKey(key string) Encoder {
	cf.raw = cf.raw + "[" + key + "]    "
	return cf
}

// WithField 添加字段.
func (cf *ConsoleFormatter) WithField(key, raw string) Encoder {
	cf.raw = cf.raw + fmt.Sprintf("[%s:%s]    ", key, raw)
	return cf
}

func (cf *ConsoleFormatter) Debug(msg string) {
	zapLogger.Debug(cf.raw + msg)
}

func (cf *ConsoleFormatter) Debugf(format string, v ...interface{}) {
	zapLogger.Debug(fmt.Sprintf(cf.raw+format, v...))
}

func (cf *ConsoleFormatter) Info(msg string) {
	zapLogger.Info(cf.raw + msg)
}

func (cf *ConsoleFormatter) Infof(format string, v ...interface{}) {
	zapLogger.Info(fmt.Sprintf(cf.raw+format, v...))
}

func (cf *ConsoleFormatter) Warn(msg string) {
	zapLogger.Warn(cf.raw + msg)
}

func (cf *ConsoleFormatter) Warnf(format string, v ...interface{}) {
	zapLogger.Warn(fmt.Sprintf(cf.raw+format, v...))
}

func (cf *ConsoleFormatter) Error(msg string) {
	zapLogger.Error(cf.raw + msg)
}

func (cf *ConsoleFormatter) Errorf(format string, v ...interface{}) {
	zapLogger.Error(fmt.Sprintf(cf.raw+format, v...))
}

func (cf *ConsoleFormatter) Fatal(msg string) {
	zapLogger.Fatal(cf.raw + msg)
}

func (cf *ConsoleFormatter) Fatalf(format string, v ...interface{}) {
	zapLogger.Fatal(fmt.Sprintf(cf.raw+format, v...))
}
