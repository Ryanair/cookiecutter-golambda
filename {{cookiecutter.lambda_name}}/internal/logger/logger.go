package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger

	invalidRequestMessage  = "Invalid request body: %v, due to: %v"
	invalidResponseMessage = "Invalid response body: %v, due to: %v"
)

// nolint
func Initialize() {
	l, err := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	if err != nil {
		panic("error")
	}
	l = l.WithOptions(zap.AddCallerSkip(1))
	log = l.Sugar()
}

func InvalidRequest(value interface{}, err error) {
	log.Errorf(invalidRequestMessage, value, err)
}

func InvalidResponse(value interface{}, err error) {
	log.Errorf(invalidResponseMessage, value, err)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args)
}

func Infow(msg string, keyAndValues ...interface{}) {
	log.Infow(msg, keyAndValues)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args)
}

func Warnw(msg string, keyAndValues ...interface{}) {
	log.Warnw(msg, keyAndValues)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args)
}

func Errorw(msg string, keyAndValues ...interface{}) {
	log.Errorw(msg, keyAndValues)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args)
}

func Fatalw(msg string, keyAndValues ...interface{}) {
	log.Fatalw(msg, keyAndValues)
}
