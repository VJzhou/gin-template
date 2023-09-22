package zapx

import (
	"fmt"
	"go.uber.org/zap"
)

type JsonEncoder struct {
	fields []zap.Field
	val    string
	logger *zap.Logger
}

func (je *JsonEncoder) WithKey(key string) Encoder {
	je.val = je.val + key + " "
	return je
}

func (je *JsonEncoder) WithField(key, val string) Encoder {
	je.fields = append(je.fields, zap.String(key, val))
	return je
}

func (je *JsonEncoder) Debug(msg string) {
	je.logger.Debug(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Debugf(format string, v ...interface{}) {
	je.logger.Debug(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Info(msg string) {
	je.logger.Info(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Infof(format string, v ...interface{}) {
	je.logger.Info(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Warn(msg string) {
	je.logger.Warn(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Warnf(format string, v ...interface{}) {
	je.logger.Warn(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Error(msg string) {
	je.logger.Error(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Errorf(format string, v ...interface{}) {
	je.logger.Error(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Fatal(msg string) {
	je.logger.Fatal(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Fatalf(format string, v ...interface{}) {
	je.logger.Fatal(fmt.Sprintf(je.val+format, v...), je.fields...)
}
