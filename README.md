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
	golog.Info("hello world.")
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
	golog.InitLogger(
		golog.SetLogOutput("file"), // 定义日志输出方式为 file
		golog.SetLogFile("./logs/std.log"),  // 定义日志写入文件路径
		golog.SetLogLevel("error"), // 定义日志级别，默认 info
		golog.SetLogFileRotationTime(time.Hour),  // 定义日志切割间隔，默认为 1 天
		golog.SetLogFileMaxAge(time.Hour*24*7), // 定义日志保留时长，默认保留近 7 天的日志
		golog.SetLogFileSuffix(".%Y%m%d%H"),  // 定义日志归档文件后缀，这里注意要与切割间隔保持一直
	)

	golog.Info("hello", "good")
	golog.Error("this is a error message")
}
```

- 获取多个 log 实例，将不同的日志信息写入不同的文件
```go
package main

import (
	"context"
	"github.com/amuluze/golog"
	"time"
)

func InitLog() {
	golog.CreateLogger(
		golog.SetName("nlog"),
		golog.SetLogFile("./logs/nlog.log"),
		golog.SetLogLevel("info"),
		golog.SetLogOutput("file"),
		golog.SetLogFormat("json"),
		golog.SetLogFileRotationTime(time.Hour),
		golog.SetLogFileMaxAge(time.Hour*24*7),
		golog.SetLogFileSuffix(".%Y%m%d%H"),
	)

	golog.CreateLogger(
		golog.SetName("mlog"),
		golog.SetLogFile("./logs/mlog.log"),
		golog.SetLogLevel("info"),
		golog.SetLogOutput("file"),
		golog.SetLogFormat("text"),
		golog.SetLogFileRotationTime(time.Hour),
		golog.SetLogFileMaxAge(time.Hour*24*7),
		golog.SetLogFileSuffix(".%Y%m%d%H"),
	)
}

func main() {
	InitLog()
	nLogger := golog.GetLoggerByName("nlog")
	nLogger.Info("test info level log")

	ctx := nLogger.NewTraceIDContext(context.Background(), "123456")
	ctx = nLogger.NewTagContext(ctx, "__main__")
	nLogger.Infof(ctx, "test log with context")

	mLogger := golog.GetLoggerByName("mlog")
	mLogger.Error("test error level log")
}
```