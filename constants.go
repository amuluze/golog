// Package golog
// Date: 2022/8/4 13:21
// Author: Amu
// Description:
package golog

import "go.uber.org/zap/zapcore"

const (
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
	ErrorLevel = zapcore.ErrorLevel
	WarnLevel  = zapcore.WarnLevel
	InfoLevel  = zapcore.InfoLevel
	DebugLevel = zapcore.DebugLevel
)

// Define key
const (
	TraceIDKey  = "trace_id"
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	TagKey      = "tag"
)

type (
	traceIDKey  struct{}
	userIDKey   struct{}
	userNameKey struct{}
	tagKey      struct{}
)
