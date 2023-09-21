package logx1

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LevelEnablerFunc func(Level) bool

type TeeOption struct {
	ws []zapcore.WriteSyncer
	LevelEnablerFunc
	encoder zapcore.Encoder
}

func NewTee(tees []TeeOption, opts ...zap.Option) *Logger {
	var cores []zapcore.Core
	for _, tee := range tees {
		core := zapcore.NewCore(
			tee.encoder,
			//zapcore.AddSync(tee.Out),
			zapcore.NewMultiWriteSyncer(tee.ws...),
			zap.LevelEnablerFunc(tee.LevelEnablerFunc),
		)
		cores = append(cores, core)
	}
	return &Logger{l: zap.New(zapcore.NewTee(cores...), opts...)}
}
