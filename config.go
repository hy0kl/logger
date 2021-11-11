package logger

import (
	"strconv"
)

type Config struct {
	//日志输出文件
	FileName string
	//日志级别
	Level string
	//日志标签 多日志时使用
	Tag string
	//日志格式: json, console(默认)
	Format string
	//旧日志保留5个备份
	MaxBackups string
	//日志最大保存MB
	MaxSize string
	//最多保留30个日志 和MaxBackups参数配置1个就可以
	MaxAge string
	// gzip包 默认false
	Compress bool
	//Console输出
	Console bool
	//SuffixEnable环境变量
	SuffixEnv string
}

func (c *Config) SetConfigMap(conf map[string]string) *Config {
	for k, v := range conf {
		switch k {
		case "fileName":
			c.FileName = v
		case "level":
			c.Level = v
		case "tag":
			c.Tag = v
		case "format":
			c.Format = v
		case "maxBackups":
			c.MaxBackups = v
		case "maxSize":
			c.MaxSize = v
		case "maxAge":
			c.MaxAge = v
		case "compress":
			c.Compress = v != "false"
		case "console":
			c.Console = v == "true"
		case "suffixEnv":
			c.SuffixEnv = v
		}
	}

	return c
}

// Parse a number with K/M/G suffixes based on thousands (1000) or 2^10 (1024)
func strToNumSuffix(str string, mult int) int {
	num := 1
	if len(str) > 1 {
		switch str[len(str)-1] {
		case 'G', 'g':
			num *= mult
			fallthrough
		case 'M', 'm':
			num *= mult
			fallthrough
		case 'K', 'k':
			num *= mult
			str = str[0 : len(str)-1]
		}
	}
	parsed, _ := strconv.Atoi(str)
	return parsed * num
}
