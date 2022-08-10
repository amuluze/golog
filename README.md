### golog
Golang log 工具封装

#### example
- 不进行初始化直接使用时，只会在终端打印 log 信息
```go
package main

import (
	"github.com/amuluze/golog"
)

func main() {
	log.Info("hello world.")
}
```

- 初始化，将 log 信息写入文件
```go
package main

import (
	"github.com/amuluze/golog"
	"time"
)

func main() {
	log.InitLogger(
		log.SetLogOutput("file"), // 定义日志输出方式为 file
		log.SetLogFile("./logs/std.log"),  // 定义日志写入文件路径
		log.SetLogLevel("error"), // 定义日志级别，默认 info
		log.SetLogFileRotationTime(time.Hour),  // 定义日志切割间隔，默认为 1 天
		log.SetLogFileMaxAge(time.Hour*24*7), // 定义日志保留时长，默认保留近 7 天的日志
		log.SetLogFileSuffix(".%Y%m%d%H"),  // 定义日志归档文件后缀，这里注意要与切割间隔保持一直
	)

	log.Info("hello", "good")
	log.Error("this is a error message")
}
```

- 获取多个 log 实例，将不同的日志信息写入不同的文件
```go
package main

import (
	"github.com/amuluze/golog"
	"time"
)

func main() {
	log.InitLogger(
		log.SetLogOutput("file"), // 定义日志输出方式为 file
		log.SetLogFile("./logs/std.log"),  // 定义日志写入文件路径
		log.SetLogLevel("error"), // 定义日志级别，默认 info
		log.SetLogFileRotationTime(time.Hour),  // 定义日志切割间隔，默认为 1 天
		log.SetLogFileMaxAge(time.Hour*24*7), // 定义日志保留时长，默认保留近 7 天的日志
		log.SetLogFileSuffix(".%Y%m%d%H"),  // 定义日志归档文件后缀，这里注意要与切割间隔保持一直
	)

	log.Info("hello", "good")
	log.Error("this is a error message")

	nlog := log.GetLogger(
		log.SetName("nlog"),
		log.SetLogFile("./logs/nlog.log"),
		log.SetLogLevel("info"),
		log.SetLogOutput("file"),
		log.SetLogFormat("json"),
		log.SetLogFileRotationTime(time.Hour),
		log.SetLogFileMaxAge(time.Hour*24*7),
		log.SetLogFileSuffix(".%Y%m%d%H"),
	)
	nlog.Info("this is from another logger info")
	nlog.Error("error message")
}
```