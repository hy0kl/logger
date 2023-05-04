# logger

日志作为整个代码行为的记录，是程序执行逻辑和异常最直接的反馈，支持标准输出和高性能磁盘写入,多样化的配置方式，使用起来简单方便，5个日志级别满足项目中各种需求。

## 安装

```shell script
go get github.com/hy0kl/logger
```

## 配置示例
```ini
[Log]
;日志输出文件
fileName = /home/logs/app.log
;日志级别
level = DEBUG
;日志最大保存MB
maxSize = 200
;旧日志保留5个备份
maxBackups = 5
;最多保留30天日志 和MaxBackups参数配置1个就可以
MaxAge=0
;控制台输出
console = true
;SuffixEnable环境变量
suffixEnv = dev
;启用日志压缩
compress = true
```

### 初始化

```go
package main

import (
	"context"
	"time"
	
    "github.com/hy0kl/logger"
)

func main(){
    cfgMap := map[string]string{
        "fileName": `./app.log`,
        "console":  "true",
        "level":    "DEBUG",
        "maxSize":  "200",
        "maxAge":   "0",
        "format":   "json",
    }
        
    // 设置环境信息
    logger.SetEnv("gray")
    logger.SetName("testLogger")
    logger.SetVersion("logger-v1.0.0")
    logConfig := logger.DefaultConf().SetConfigMap(cfgMap)

    // 初始化日志对象
    logger.InitWithConfig(logConfig)    
    defer logger.Sync()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_time", time.Now())
	
    // 记录错误日志
    // 前3个参数必传：上下文, 日志标签, 日志正文
    logger.Ix(ctx, "testTag", "我是一条测试日志.... %v, %v, %v, %v", "age", 30, "box", "Tom")
}
```

### 日志内容示例

```json
{"x_level":"info","@timestamp":"2023-05-04T10:53:48.054+0800","x_caller":"/Users/tal/yangyj/logger/logger_test.go:36","x_msg":"我是一条测试日志.... age, 30, box, Tom","x_env":"dev","x_name":"testLogger","x_version":"api:0.0.1","x_server_ip":"10.29.116.85","x_hostname":"workdev","x_trace_id":"01f29462-8619-4c64-8af1-dc4f315a4de7","x_timestamp":1683168828054,"x_duration":"1.001063083s","x_tag":"testTag"}
```