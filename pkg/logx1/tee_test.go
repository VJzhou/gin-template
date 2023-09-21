package logx1

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"testing"
)

func TestTee(t *testing.T) {
	config := &Config{
		FilePath:   "./logs/x.log",
		Encoder:    "console",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
	}
	encoder := GetEncoder(config.Encoder)
	tees := []TeeOption{
		{
			ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				//zapcore.AddSync(getHook(config, config.GetInfoPath())),
			},
			LevelEnablerFunc: func(level Level) bool {
				return level <= zap.InfoLevel
			},
			encoder: encoder,
		},
		{
			ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				//zapcore.AddSync(getHook(config, config.GetErrPath())),
			},
			LevelEnablerFunc: func(level Level) bool {
				return level > zap.InfoLevel
			},
			encoder: encoder,
		},
	}
	logger := NewTee(tees)

	defer logger.Sync()

	logger.Info("Info tee msg1")
	logger.Warn("Warn tee 2")
	logger.Error("Error tee msg3") // 不会输出
}

func getHook(config *Config, filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     config.MaxAge,     // 文件最多保存多少天
		Compress:   true,              // 是否压缩
		LocalTime:  true,              // 备份文件名本地/UTC时间
	}
}

// https://github.com/jianghushinian/gokit/blob/main/log/zap/log.go
