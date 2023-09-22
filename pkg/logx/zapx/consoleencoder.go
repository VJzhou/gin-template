package zapx

import (
	"fmt"
	"go.uber.org/zap"
)

type ConsoleEncoder struct {
	raw    string
	logger *zap.Logger
}

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
