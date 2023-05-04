package logger

import (
	"context"
)

var stdBuilder = new(zapBuilder)

func SetEnv(env string) {
	stdBuilder.SetEnv(env)
}

func SetName(name string) {
	stdBuilder.SetName(name)
}

func SetVersion(version string) {
	stdBuilder.SetVersion(version)
}

func Sync() error {
	return stdBuilder.Sync()
}

func Dx(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "DEBUG", tag, message, fields...)
}

func Ix(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "INFO", tag, message, fields...)
}

func Wx(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "WARNING", tag, message, fields...)
}

func Ex(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "ERROR", tag, message, fields...)
}

func Fx(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "FATAL", tag, message, fields...)
}

func Px(ctx context.Context, tag string, message string, fields ...interface{}) {
	stdBuilder.LoggerX(ctx, "FATAL", tag, message, fields...)
}
