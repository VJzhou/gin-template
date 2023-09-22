package zapx

import (
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New(os.Stdout, NewJsonEncoder(), zap.InfoLevel)
	logger.Info("hahaha")
}
