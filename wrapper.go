package logger

import (
	"fmt"
	"os"
)

func init() {
	stdBuilder.serverIp, _ = Extract("")
	stdBuilder.hostName, _ = os.Hostname()
}

func DefaultConf() *Config {
	config := new(Config)
	//存储路径
	config.FileName = "./logs/app.log"
	//日志级别
	config.Level = "INFO"
	//日志标签 多日志时使用
	config.Tag = "default"
	//日志格式
	config.Format = ""
	//旧日志保留5个备份
	config.MaxBackups = "5"
	//日志最大保存MB
	config.MaxSize = "100"
	//最多保留30天日志 和MaxBackups参数配置1个就可以
	config.MaxAge = "0"
	// gzip包 默认false
	config.Compress = false
	//Console输出
	config.Console = false

	return config
}

// InitWithConfig 自定义config Init
func InitWithConfig(config *Config) {
	if config.FileName == "" {
		_, _ = fmt.Fprintf(os.Stderr, "InitLoggerConfig: Error: config could not found logpath %s\n", config.FileName)
		os.Exit(1)
	}

	stdBuilder.LoadConfig(config)
}
