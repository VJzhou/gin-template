package logx1

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

func GetEncoder(encoder string) zapcore.Encoder {
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
