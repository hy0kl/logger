package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "@timestamp",
	LevelKey:       "x_level",
	NameKey:        "x_name",
	CallerKey:      "x_caller",
	MessageKey:     "x_msg",
	StacktraceKey:  "x_backtrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
	EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
}

type zapBuilder struct {
	builder
	zapLogger *zap.Logger
	logWriter *lumberjack.Logger
}

func (log *zapBuilder) LoadConfig(config *Config) {
	var lvl zapcore.Level
	switch config.Level {
	case "DEBUG":
		lvl = zapcore.DebugLevel
	case "INFO":
		lvl = zapcore.InfoLevel
	case "WARNING":
		lvl = zapcore.WarnLevel
	case "ERROR":
		lvl = zapcore.ErrorLevel
	case "FATAL":
		lvl = zapcore.FatalLevel
	case "PANIC":
		lvl = zapcore.PanicLevel
	default:
		_, _ = fmt.Fprintf(os.Stderr, "LoadConfig: Error: Required child <%s> for filter has unknown value %s\n", "level", config.Level)
		os.Exit(1)
	}

	log.logWriter = log.getLogWriter(config)

	var encoder = zapcore.NewConsoleEncoder(encoderConfig)
	if config.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	writeSyncer := zapcore.AddSync(log.logWriter)
	core := zapcore.NewCore(encoder, writeSyncer, lvl)

	if config.Console == true {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		colorEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewTee(core, zapcore.NewCore(colorEncoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), lvl))
	}

	log.zapLogger = zap.New(core, zap.AddCaller(), zap.Fields(
		zap.String("x_env", log.env),
		zap.String("x_name", log.name),
		zap.String("x_version", log.version),
		zap.String("x_server_ip", log.serverIp),
		zap.String("x_hostname", log.hostName),
	), zap.AddCallerSkip(2))
}

func (log *zapBuilder) LoggerX(ctx context.Context, lvl string, tag string, message string, fields ...interface{}) {
	message = fmt.Sprintf(message, fields...)

	if tag == "" {
		tag = "NoTagError"
	}

	var extFields = []interface{}{"x_tag", tag}
	fields = append(log.Build(ctx), extFields...)

	sugar := log.zapLogger.Sugar()
	switch lvl {
	case "DEBUG":
		sugar.Debugw(message, fields...)
	case "INFO":
		sugar.Infow(message, fields...)
	case "WARNING":
		sugar.Warnw(message, fields...)
	case "ERROR":
		sugar.Errorw(message, fields...)
	case "FATAL":
		sugar.Fatalw(message, fields...)
	case "PANIC":
		sugar.Panicw(message, fields...)
	}
}

func (log *zapBuilder) Sync() error {
	err := log.zapLogger.Sync()

	return err
}
