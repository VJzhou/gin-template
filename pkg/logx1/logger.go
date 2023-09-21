package logx1

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Level = zapcore.Level

type Logger struct {
	l  *zap.Logger
	al *zap.AtomicLevel
}

func New(out io.Writer, level Level, encoder zapcore.Encoder, opts ...zap.Option) *Logger {
	if out == nil {
		out = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(out),
		al,
	)
	return &Logger{l: zap.New(core, opts...), al: &al}
}

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

//func (l *Logger) WithKey(key string) Encoder {
//	l.val = l.val + key + " "
//	return l
//}
//
//func (l *Logger) WithField(key, val string) Encoder {
//	l.fields = append(l.fields, zap.String(key, val))
//	return l
//}
//
//func (l *Logger) Debug(msg string) {
//	zapLogger.Debug(l.val+msg, l.fields...)
//}
//
//func (l *Logger) Debugf(format string, v ...interface{}) {
//	logx.zapLogger.Debug(fmt.Sprintf(l.val+format, v...), l.fields...)
//}
//
//func (l *Logger) Info(msg string) {
//	logx.zapLogger.Info(l.val+msg, l.fields...)
//}
//
//func (l *Logger) Infof(format string, v ...interface{}) {
//	logx.zapLogger.Info(fmt.Sprintf(l.val+format, v...), l.fields...)
//}
//
//func (l *Logger) Warn(msg string) {
//	logx.zapLogger.Warn(l.val+msg, l.fields...)
//}
//
//func (l *Logger) Warnf(format string, v ...interface{}) {
//	logx.zapLogger.Warn(fmt.Sprintf(l.val+format, v...), l.fields...)
//}
//
//func (l *Logger) Error(msg string) {
//	logx.zapLogger.Error(l.val+msg, l.fields...)
//}
//
//func (l *Logger) Errorf(format string, v ...interface{}) {
//	logx.zapLogger.Error(fmt.Sprintf(l.val+format, v...), l.fields...)
//}
//
//func (l *Logger) Fatal(msg string) {
//	logx.zapLogger.Fatal(l.val+msg, l.fields...)
//}
//
//func (l *Logger) Fatalf(format string, v ...interface{}) {
//	logx.zapLogger.Fatal(fmt.Sprintf(l.val+format, v...), l.fields...)
//}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}
