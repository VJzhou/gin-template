package zapx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var _ Encoder = (*ConsoleEncoder)(nil)

type ConsoleEncoder struct {
	raw    string
	logger *zap.Logger
}

func NewConsoleEncoder() *ConsoleEncoder {
	return &ConsoleEncoder{}
}

func GetConsoleZapCoreEncoder() zapcore.Encoder {
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
func (ce *ConsoleEncoder) WithKey(key string) Encoder {
	ce.raw = ce.raw + "[" + key + "]    "
	return ce
}

// WithField 添加字段.
func (ce *ConsoleEncoder) WithField(key, raw string) Encoder {
	ce.raw = ce.raw + fmt.Sprintf("[%s:%s]    ", key, raw)
	return ce
}

func (ce *ConsoleEncoder) Debug(msg string) {
	ce.logger.Debug(ce.raw + msg)
}

func (ce *ConsoleEncoder) Debugf(format string, v ...interface{}) {
	ce.logger.Debug(fmt.Sprintf(ce.raw+format, v...))
}

func (ce *ConsoleEncoder) Info(msg string) {
	ce.logger.Info(ce.raw + msg)
}

func (ce *ConsoleEncoder) Infof(format string, v ...interface{}) {
	ce.logger.Info(fmt.Sprintf(ce.raw+format, v...))
}

func (ce *ConsoleEncoder) Warn(msg string) {
	ce.logger.Warn(ce.raw + msg)
}

func (ce *ConsoleEncoder) Warnf(format string, v ...interface{}) {
	ce.logger.Warn(fmt.Sprintf(ce.raw+format, v...))
}

func (ce *ConsoleEncoder) Error(msg string) {
	ce.logger.Error(ce.raw + msg)
}

func (ce *ConsoleEncoder) Errorf(format string, v ...interface{}) {
	ce.logger.Error(fmt.Sprintf(ce.raw+format, v...))
}

func (ce *ConsoleEncoder) Fatal(msg string) {
	ce.logger.Fatal(ce.raw + msg)
}

func (ce *ConsoleEncoder) Fatalf(format string, v ...interface{}) {
	ce.logger.Fatal(fmt.Sprintf(ce.raw+format, v...))
}
