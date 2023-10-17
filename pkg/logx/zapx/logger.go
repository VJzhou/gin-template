package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var logger *Logger

type (
	Level = zapcore.Level

	Logger struct {
		l  *zap.Logger
		al *zap.AtomicLevel
	}
)

func SetLogger(l *Logger) {
	logger = l
}

func GetLogger() *Logger {
	return logger
}

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

func GetHook(config *Config, filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     config.MaxAge,     // 文件最多保存多少天
		Compress:   true,              // 是否压缩
		LocalTime:  true,              // 备份文件名本地/UTC时间
	}
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
