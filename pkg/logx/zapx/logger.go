package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type (
	Level = zapcore.Level

	Logger struct {
		l  *zap.Logger
		al *zap.AtomicLevel
	}
)

func New(out io.Writer, encoder zapcore.Encoder, level Level, opts ...zap.Option) *Logger {
	if out == nil {
		out = os.Stderr
	}
	al := zap.NewAtomicLevelAt(level)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(out),
		al,
	)
	zLogger := zap.New(core, opts...)

	return &Logger{l: zLogger, al: &al}
}

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
