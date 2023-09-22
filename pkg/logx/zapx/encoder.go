package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	consoleEncoder = "console"
	jsonEncoder    = "json"
)

type Encoder interface {
	WithKey(key string) Encoder
	WithField(key, val string) Encoder
	Debug(msg string)
	Debugf(format string, v ...interface{})
	Info(msg string)
	Infof(format string, v ...interface{})
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Error(msg string)
	Errorf(format string, v ...interface{})
	Fatal(msg string)
	Fatalf(format string, v ...interface{})
}

func NewEncoder(encoderStr string, logger *zap.Logger) Encoder {
	var encoder Encoder
	switch encoderStr {
	case jsonEncoder:
		encoder = &JsonEncoder{
			fields: make([]zap.Field, 0),
			val:    "",
			logger: logger,
		}
	case consoleEncoder:
		encoder = &ConsoleEncoder{
			raw:    "",
			logger: logger,
		}
	default:
		encoder = &ConsoleEncoder{
			raw:    "",
			logger: logger,
		}
	}
	return encoder
}

func GetZapCoreEncoder(encoder string) zapcore.Encoder {
	var zapEncoder zapcore.Encoder
	switch encoder {
	case jsonEncoder:
		zapEncoder = NewJsonEncoder()
	case consoleEncoder:
		zapEncoder = NewConsoleEncoder()
	default:
		zapEncoder = NewConsoleEncoder()
	}
	return zapEncoder
}

func NewJsonEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(getFilePath(&caller))
	}

	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	return zapcore.NewJSONEncoder(config)
}

func NewConsoleEncoder() zapcore.Encoder {
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

func getFilePath(ec *zapcore.EntryCaller) string {
	if !ec.Defined {
		return "undefined"
	}
	buf := buffer.NewPool().Get()
	buf.AppendString(ec.Function)
	buf.AppendByte(':')
	buf.AppendInt(int64(ec.Line))
	caller := buf.String()
	buf.Free()
	return caller
}
