// Package log
// Date: 2022/8/4 18:02
// Author: Amu
// Description:
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var std = &Logger{
	SugaredLogger: zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), InfoLevel)).Sugar(),
	name:          "std",
	loggers:       make(map[string]*Logger),
}

func GetLogger(options ...Option) *Logger {
	return std.GetLogger(options...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}
