package logger

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cast"
)

type builder struct {
	env      string
	name     string
	version  string
	hostName string
	serverIp string
}

func (t *builder) SetEnv(env string) {
	t.env = env
}

func (t *builder) SetName(name string) {
	t.name = name
}

func (t *builder) SetVersion(version string) {
	t.version = version
}

func (t *builder) LoadConfig(config *Config) {
	return
}

func (t *builder) getLogWriter(config *Config) *lumberjack.Logger {
	filename := strings.Trim(config.FileName, " \r\n")
	maxBackups := strToNumSuffix(strings.Trim(config.MaxBackups, " \r\n"), 1000)
	maxSize := strToNumSuffix(strings.Trim(config.MaxSize, " \r\n"), 1024)
	maxAge, _ := strconv.Atoi(strings.Trim(config.MaxAge, " \r\n"))
	compress := config.Compress

	// MaxAge 与 MaxBackups参数配置1个就可以。MaxAge 优先
	if maxAge > 0 {
		maxBackups = 0
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,   // 日志输出文件
		MaxSize:    maxSize,    // 日志最大保存1M
		MaxBackups: maxBackups, // 旧日志保留5个备份
		MaxAge:     maxAge,     // 最多保留30天日志 和MaxBackups参数配置1个就可以
		Compress:   compress,   // 自导打 gzip包 默认false
	}

	return lumberJackLogger
}

func (t *builder) Build(ctx context.Context) (expand []interface{}) {
	traceId := cast.ToString(ctx.Value("trace_id"))

	if traceId == "" {
		traceId = uuid.New().String()
	}

	var duration string
	if startValue := ctx.Value("request_time"); startValue != nil {
		if start, ok := startValue.(time.Time); ok {
			cost := time.Now().Sub(start)
			duration = cost.String()
		}
	}
	expand = []interface{}{
		"x_trace_id", traceId,
		"x_timestamp", time.Now().UnixMilli(),
		"x_duration", duration,
	}

	return
}
