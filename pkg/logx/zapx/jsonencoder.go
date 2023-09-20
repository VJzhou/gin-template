package zapx

import (
	"fmt"
	"gin-demo/pkg/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var _ Encoder = (*JsonEncoder)(nil)

type JsonEncoder struct {
	fields []zap.Field
	val    string
}

func NewJsonEncoder() *JsonEncoder {
	return &JsonEncoder{
		fields: make([]zap.Field, 0),
	}
}

func GetJsonZapCoreEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(getFilePath(&caller))
	}

	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	return zapcore.NewJSONEncoder(config)
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
	zapLogger.Debug(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Debugf(format string, v ...interface{}) {
	logx.zapLogger.Debug(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Info(msg string) {
	logx.zapLogger.Info(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Infof(format string, v ...interface{}) {
	logx.zapLogger.Info(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Warn(msg string) {
	logx.zapLogger.Warn(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Warnf(format string, v ...interface{}) {
	logx.zapLogger.Warn(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Error(msg string) {
	logx.zapLogger.Error(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Errorf(format string, v ...interface{}) {
	logx.zapLogger.Error(fmt.Sprintf(je.val+format, v...), je.fields...)
}

func (je *JsonEncoder) Fatal(msg string) {
	logx.zapLogger.Fatal(je.val+msg, je.fields...)
}

func (je *JsonEncoder) Fatalf(format string, v ...interface{}) {
	logx.zapLogger.Fatal(fmt.Sprintf(je.val+format, v...), je.fields...)
}
