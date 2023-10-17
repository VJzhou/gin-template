package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestTee(t *testing.T) {
	config := &Config{
		FilePath:   "./logs/x.log",
		Encoder:    "json",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
	}
	encoder := GetZapCoreEncoder(config.Encoder)
	tees := []TeeOption{
		{
			Ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(GetHook(config, config.GetInfoPath())),
			},
			LevelEnablerFunc: func(level Level) bool {
				return level <= zap.InfoLevel
			},
			Encoder: encoder,
		},
		{
			Ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(GetHook(config, config.GetErrPath())),
			},
			LevelEnablerFunc: func(level Level) bool {
				return level > zap.InfoLevel
			},
			Encoder: encoder,
		},
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开发模式
	development := zap.Development()
	// 二次封装
	skip := zap.AddCallerSkip(1)

	logger := NewTee(tees, caller, development, skip)

	defer logger.Sync()

	logger.Info("Info tee msg1")
	logger.Warn("Warn tee 2")
	logger.Error("Error tee msg3") // 不会输出
}
