package zapx

import (
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestJsonEncoder(t *testing.T) {
	logger := New(os.Stdout, NewJsonEncoder(), zap.DebugLevel)

	json := &JsonEncoder{
		fields: make([]zap.Field, 0),
		val:    "",
		logger: logger.l,
	}

	json.Info("hahaha")
	json.Error("error info ")
	json.Debug("debug")

	json.Infof("%s", "infof")
	json.Debugf("%s", "debugf")
	json.Errorf("%s", "errorf")
}
