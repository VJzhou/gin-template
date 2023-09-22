package zapx

import (
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestConsoleEncoder(t *testing.T) {
	logger := New(os.Stdout, NewConsoleEncoder(), zap.DebugLevel)

	console := &ConsoleEncoder{
		raw:    "",
		logger: logger.l,
	}

	console.Info("hahaha")
	console.Error("error info ")
	console.Debug("debug")

	console.Infof("%s", "infof")
	console.Debugf("%s", "debugf")
	console.Errorf("%s", "errorf")
}
