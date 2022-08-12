// Package golog
// Date: 2022/8/4 18:02
// Author: Amu
// Description:
package golog

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var std = &Logger{
	SugaredLogger: zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), InfoLevel)).Sugar(),
	name:          "std",
	loggers:       make(map[string]*Logger),
}

func CreateLogger(options ...Option) {
	std.CreateLogger(options...)
}

func GetLoggerByName(name string) *Logger {
	if _, ok := std.loggers[name]; ok {
		return std.loggers[name]
	}
	return nil
}

func NewTagContext(ctx context.Context, tag string) context.Context {
	return std.NewTagContext(ctx, tag)
}

func FromTagContext(ctx context.Context) string {
	return std.FromTagContext(ctx)
}

func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return std.NewTraceIDContext(ctx, traceID)
}

func FromTraceIDContext(ctx context.Context) string {
	return std.FromTraceIDContext(ctx)
}

func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return std.NewUserIDContext(ctx, userID)
}

func FromUserIDContext(ctx context.Context) string {
	return std.FromUserIDContext(ctx)
}

func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return std.NewUserNameContext(ctx, userName)
}

func FromUserNameContext(ctx context.Context) string {
	return std.FromUserNameContext(ctx)
}

func WithContext(ctx context.Context) {
	std.SugaredLogger = std.WithContext(ctx)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Info(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(ctx context.Context, args ...interface{}) {
	WithContext(ctx)
	std.Error(args...)
}
