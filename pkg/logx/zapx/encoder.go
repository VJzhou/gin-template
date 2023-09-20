package zapx

import "go.uber.org/zap/zapcore"

type Encoder interface {
	//GetEncoder() zapcore.Encoder
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

var (
	consoleEncoder = "console"
	jsonEncoder    = "json"
)

func GetZapEncoder(encoder string) zapcore.Encoder {
	var zapEncoder zapcore.Encoder
	switch encoder {
	case jsonEncoder:
		zapEncoder = GetJsonZapCoreEncoder()
	case consoleEncoder:
		zapEncoder = GetConsoleZapCoreEncoder()
	default:
		zapEncoder = GetConsoleZapCoreEncoder()
	}
	return zapEncoder
}

//func GetEncoder(encoder string) Encoder {
//	var localEncoder Encoder
//	switch encoder {
//	case jsonEncoder:
//		localEncoder = NewJsonEncoder()
//	case consoleEncoder:
//		localEncoder = NewConsoleEncoder()
//	default:
//		localEncoder = NewConsoleEncoder()
//	}
//	return localEncoder
//}
