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
import (
    "github.com/hy0kl/gconfig"
    "github.com/hy0kl/logger"
)

func main(){
    
        // 准备日志配置参数
        cfgMap := gconfig.GetConfStringMap("Log")
    
        // 设置环境信息
    	logger.SetEnv("gray")
    	logger.SetName("testLogger")
    	logger.SetVersion("logger-v1.0.0")
    	logConfig := NewConfig().SetConfigMap(cfgMap)
    
        // 初始化日志对象
    	logger.InitWithConfig(logConfig)    
    	defer logger.Sync()
    
        // 记录错误日志
        // 前3个参数必传：上下文, 日志标签, 日志正文
    	logger.Ex(ctx, "logTag", "我是一条日志....")
}
```

### 日志内容示例

```json
{"x_level":"info","@timestamp":"2021-11-11T18:07:04.664+0800","x_caller":"/Users/tal/yangyj/logger/logger_test.go:36","x_msg":"我是一条测试日志.... age, 30, box, Tom","x_env":"dev","x_name":"testLogger","x_version":"api:0.0.1","x_department":"r&d","x_server_ip":"10.25.190.61","x_hostname":"workdev","x_trace_id":"886ab55c-5186-42f8-8fbf-12d08b59792e","x_timestamp":1636625224664,"x_duration":"1.001359083s","x_tag":"testTag"}
```