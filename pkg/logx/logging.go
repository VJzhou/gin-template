package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
	"strings"
)

type (
	Config struct {
		Path     string
		Encoder  string
		Filename string
		Level    zap.LevelEnablerFunc
	}

	Encoder interface {
		Config() zapcore.Encoder
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
)

var (
	maxSize        = 200
	maxBackups     = 20
	maxAge         = 30
	zapLogger      *zap.Logger
	pool           = buffer.NewPool()
	config         *Config
	ConsoleEncoder = "console"
	JsonEncoder    = "Json"
)

func Init(conf *Config) {
	config = conf

	prefix, suffix := getFileSuffixPrefix(config.Path)
	log.Println("prefix:", prefix, "suffix:", suffix)

	infoPath := path.Join(prefix + ".info" + suffix)
	errPath := path.Join(prefix + ".err" + suffix)
	log.Println("infoPath:", infoPath, "errPath:", errPath)

	//items := []
	NewLogger()
}

func NewLogger() {
	var encoder zapcore.Encoder
	var cores []zapcore.Core

	switch config.Encoder {
	case JsonEncoder:
		encoder = NewJsonFormatter().Config()
	case ConsoleEncoder:
		encoder = NewConsoleFormatter().Config()
	default:
		encoder = NewConsoleFormatter().Config()
	}

	writer := &lumberjack.Logger{
		Filename:   v.FileName,
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   true,       // 是否压缩
		LocalTime:  true,       // 备份文件名本地/UTC时间
	}
	core := zapcore.NewCore(
		encoder,                                                                          // 编码器配置;
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), // 打印到控制台和文件
		v.Level,                                                                          // 日志级别
	)
	cores = append(cores, core)

	zapLogger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1))
	return
}

func GetEncoder() Encoder {
	var encoder Encoder
	switch config.Encoder {
	case JsonEncoder:
		encoder = NewJsonFormatter()
	case ConsoleEncoder:
		encoder = NewConsoleFormatter()
	default:
		encoder = NewConsoleFormatter()
	}
	return encoder
}

func GetLogger() *zap.Logger {
	return zapLogger
}

// getFileSuffixPrefix 文件路径切割
func getFileSuffixPrefix(fileName string) (prefix, suffix string) {
	paths, _ := path.Split(fileName)
	base := path.Base(fileName)
	suffix = path.Ext(fileName)
	prefix = strings.TrimSuffix(base, suffix)
	prefix = path.Join(paths, prefix)
	return
}

// getFilePath 自定义获取文件路径.
func getFilePath(ec *zapcore.EntryCaller) string {
	if !ec.Defined {
		return "undefined"
	}
	buf := pool.Get()
	buf.AppendString(ec.Function)
	buf.AppendByte(':')
	buf.AppendInt(int64(ec.Line))
	caller := buf.String()
	buf.Free()
	return caller
}
