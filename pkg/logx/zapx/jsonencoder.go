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

func (formater *JsonEncoder) Config() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(logx.getFilePath(&caller))
	}

	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	return zapcore.NewJSONEncoder(config)
}

func (jf *JsonEncoder) WithKey(key string) Encoder {
	jf.val = jf.val + key + " "
	return jf
}

func (jf *JsonEncoder) WithField(key, val string) Encoder {
	jf.fields = append(jf.fields, zap.String(key, val))
	return jf
}

func (jf *JsonEncoder) Debug(msg string) {
	logx.zapLogger.Debug(jf.val+msg, jf.fields...)
}

func (jf *JsonEncoder) Debugf(format string, v ...interface{}) {
	logx.zapLogger.Debug(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonEncoder) Info(msg string) {
	logx.zapLogger.Info(jf.val+msg, jf.fields...)
}

func (jf *JsonEncoder) Infof(format string, v ...interface{}) {
	logx.zapLogger.Info(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonEncoder) Warn(msg string) {
	logx.zapLogger.Warn(jf.val+msg, jf.fields...)
}

func (jf *JsonEncoder) Warnf(format string, v ...interface{}) {
	logx.zapLogger.Warn(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonEncoder) Error(msg string) {
	logx.zapLogger.Error(jf.val+msg, jf.fields...)
}

func (jf *JsonEncoder) Errorf(format string, v ...interface{}) {
	logx.zapLogger.Error(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonEncoder) Fatal(msg string) {
	logx.zapLogger.Fatal(jf.val+msg, jf.fields...)
}

func (jf *JsonEncoder) Fatalf(format string, v ...interface{}) {
	logx.zapLogger.Fatal(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}
