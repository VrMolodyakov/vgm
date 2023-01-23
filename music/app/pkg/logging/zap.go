package logging

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l logger

type Logger interface {
	WithFields(args ...interface{})
}

type logger struct {
	*zap.Logger
}

//TODO: probably remove it
func GetLogger() *logger {
	return &l
}

func initLogger(consoleLevel string, console io.Writer, files ...io.Writer) {
	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)
	fileEncoder := zapcore.NewJSONEncoder(config)

	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(caller.TrimmedPath())
		encoder.AppendString("|")
	}

	config.EncodeTime = zapcore.TimeEncoderOfLayout("02/01 15:04:05 -0700")
	config.ConsoleSeparator = " "
	config.EncodeName = func(n string, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(n)
		enc.AppendString("|")
	}

	config.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("|")
		enc.AppendString(l.CapitalString())
		enc.AppendString("|")
	}

	consoleEncoder := zapcore.NewConsoleEncoder(config)

	cores := make([]zapcore.Core, len(files)+1)

	var zapConsoleLevel zapcore.Level
	switch {
	case consoleLevel == "info":
		zapConsoleLevel = zap.InfoLevel
	case consoleLevel == "debug":
		zapConsoleLevel = zap.DebugLevel
	case consoleLevel == "error":
		zapConsoleLevel = zap.ErrorLevel
	}

	cores[0] = zapcore.NewCore(consoleEncoder,
		zapcore.AddSync(console),
		zapConsoleLevel,
	)

	for i := range files {
		cores[i+1] = zapcore.NewCore(fileEncoder,
			zapcore.AddSync(files[i]),
			zap.DebugLevel,
		)
	}

	zap := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
	)

	l = logger{zap}
}

func (l *logger) Infow(msg string, keyAndValues ...interface{}) {
	sugar := l.Logger.Sugar()
	sugar.Infow(msg, keyAndValues...)
	l.Logger = sugar.Desugar()
}
