package logx

import "testing"

func TestNewLogger(t *testing.T) {

	config := &Config{
		Path:    "../../logs/fly.log",
		Encoder: ConsoleEncoder,
	}

	Init(config)
	defer Sync()
	Debug("debug msg")
	Debugf("debugf %s", "fly")
	Info("info msg")
	Infof("infof %d", 10)
	Warn("warn msg")
	Warnf("warnf %v", true)
	Error("err msg")
	Errorf("errorf %v", []int{1, 2, 3})
	Fatal("fatal msg")
}

func TestFatalf(t *testing.T) {

	config := &Config{
		Path:    "../../logs/fly.log",
		Encoder: JsonEncoder,
	}
	Init(config)
	defer Sync()
	Fatalf("fatalf %v", map[string]interface{}{"name": "master"})
}
