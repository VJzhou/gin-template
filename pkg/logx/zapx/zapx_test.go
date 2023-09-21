package zapx

import (
	"gin-demo/pkg/logx"
	"testing"
)

func TestNewLogger(t *testing.T) {
	config := getConfig(consoleEncoder)
	logger, _ := New(config)

	defer Sync()
	logx.Debug("debug msg")
	logx.Debugf("debugf %s", "fly")
	logx.Info("info msg")
	logx.Infof("infof %d", 10)
	logx.Warn("warn msg")
	logx.Warnf("warnf %v", true)
	logx.Error("err msg")
	logx.Errorf("errorf %v", []int{1, 2, 3})
	logx.Fatal("fatal msg")
}

func TestFatalf(t *testing.T) {

	config := getConfig(consoleEncoder)
	zapLogger, _ := New(config)

	defer logx.Sync()
	logx.Fatalf("fatalf %v", map[string]interface{}{"name": "master"})
}

func getConfig(encoder string) *Config {
	return &Config{
		FilePath:   "../../logs/fly.log",
		Encoder:    encoder,
		MaxSize:    200,
		MaxBackups: 1,
		MaxAge:     1,
	}
}
