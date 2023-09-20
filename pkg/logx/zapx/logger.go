package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger struct {
	zapLogger *zap.Logger
}

func New(config *Config) (*Logger, error) {
	zapLogger := GetZapLogger(config)
	return &Logger{
		zapLogger: zapLogger,
	}, nil
}

type LogFile struct {
	Filename string
	Level    zap.LevelEnablerFunc
}

func GetLogFiles(conf *Config) []LogFile {
	infoPath := conf.GetInfoPath()
	errPath := conf.GetErrPath()
	return []LogFile{
		{Filename: infoPath, Level: func(level zapcore.Level) bool {
			return level <= zap.InfoLevel
		}},
		{Filename: errPath, Level: func(level zapcore.Level) bool {
			return level > zap.InfoLevel
		}},
	}
}

func GetZapLogger(conf *Config) *zap.Logger {
	cores := make([]zapcore.Core, 0)
	zapEncoder := GetZapEncoder(conf.Encoder)
	logFiles := GetLogFiles(conf)
	for _, v := range logFiles {
		cores = append(cores, GetZapCore(&v, conf, zapEncoder))
	}
	return zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1))
}

func GetZapCore(logFile *LogFile, config *Config, zapEncoder zapcore.Encoder) zapcore.Core {
	writer := &lumberjack.Logger{
		Filename:   logFile.Filename,
		MaxSize:    config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     config.MaxAge,     // 文件最多保存多少天
		Compress:   true,              // 是否压缩
		LocalTime:  true,              // 备份文件名本地/UTC时间
	}
	return zapcore.NewCore(
		zapEncoder,                                                                       // 编码器配置;
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), // 打印到控制台和文件
		logFile.Level,                                                                    // 日志级别
	)
}

// getFilePath 自定义获取文件路径.
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
