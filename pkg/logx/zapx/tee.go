package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LevelEnablerFunc func(Level) bool

type TeeOption struct {
	Ws []zapcore.WriteSyncer
	LevelEnablerFunc
	Encoder zapcore.Encoder
}

func NewTee(tees []TeeOption, opts ...zap.Option) *Logger {
	var cores []zapcore.Core
	for _, tee := range tees {
		core := zapcore.NewCore(
			tee.Encoder,
			zapcore.NewMultiWriteSyncer(tee.Ws...),
			zap.LevelEnablerFunc(tee.LevelEnablerFunc),
		)
		cores = append(cores, core)
	}
	return &Logger{l: zap.New(zapcore.NewTee(cores...), opts...)}
}
