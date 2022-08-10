// Package log
// Date: 2022/8/4 13:21
// Author: Amu
// Description:
package log

import "go.uber.org/zap/zapcore"

const (
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
	ErrorLevel = zapcore.ErrorLevel
	WarnLevel  = zapcore.WarnLevel
	InfoLevel  = zapcore.InfoLevel
	DebugLevel = zapcore.DebugLevel
)
