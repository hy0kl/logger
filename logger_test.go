package logger

import (
	"context"
	"testing"
	"time"
)

func TestIx(t *testing.T) {
	cfgMap := map[string]string{
		"fileName":  `./app.log`,
		"console":   "true",
		"level":     "DEBUG",
		"maxSize":   "200",
		"maxAge":    "0",
		"suffixEnv": "release",
		"compress":  "true",
		"format":    "json",
	}

	// 初始化日志环境
	SetEnv("dev")
	SetName("testLogger")
	SetDepartment("r&d")
	SetVersion("api:0.0.1")
	logConfig := NewConfig().SetConfigMap(cfgMap)

	InitWithConfig(logConfig)

	defer Sync()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_time", time.Now())

	time.Sleep(1 * time.Second)
	Ix(ctx, "testTag", "我是一条测试日志.... %v, %v, %v, %v", "age", 30, "box", "Tom")
	time.Sleep(120 * time.Second)
}
