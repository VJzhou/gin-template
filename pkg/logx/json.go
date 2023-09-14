package logx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var _ Encoder = (*JsonFormatter)(nil)

type JsonFormatter struct {
	fields []zap.Field
	val    string
}

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{
		fields: make([]zap.Field, 0),
	}
}

func (formater *JsonFormatter) Config() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(getFilePath(caller))
	}

	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(level.String())
	}

	return zapcore.NewJSONEncoder(config)
}

func (jf *JsonFormatter) WithKey(key string) Encoder {
	jf.val = jf.val + key + " "
	return jf
}

func (jf *JsonFormatter) WithField(key, val string) Encoder {
	jf.fields = append(jf.fields, zap.String(key, val))
	return jf
}

func (jf *JsonFormatter) Debug(msg string) {
	zapLogger.Debug(jf.val+msg, jf.fields...)
}

func (jf *JsonFormatter) Debugf(format string, v ...interface{}) {
	zapLogger.Debug(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonFormatter) Info(msg string) {
	zapLogger.Info(jf.val+msg, jf.fields...)
}

func (jf *JsonFormatter) Infof(format string, v ...interface{}) {
	zapLogger.Info(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonFormatter) Warn(msg string) {
	zapLogger.Warn(jf.val+msg, jf.fields...)
}

func (jf *JsonFormatter) Warnf(format string, v ...interface{}) {
	zapLogger.Warn(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonFormatter) Error(msg string) {
	zapLogger.Error(jf.val+msg, jf.fields...)
}

func (jf *JsonFormatter) Errorf(format string, v ...interface{}) {
	zapLogger.Error(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}

func (jf *JsonFormatter) Fatal(msg string) {
	zapLogger.Fatal(jf.val+msg, jf.fields...)
}

func (jf *JsonFormatter) Fatalf(format string, v ...interface{}) {
	zapLogger.Fatal(fmt.Sprintf(jf.val+format, v...), jf.fields...)
}
