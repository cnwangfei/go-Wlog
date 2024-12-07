# go-Wlog
log工具库

每个库对日志接口的要求各不相同，这里统一一个日志接口ilog;  
wlog实现接口ilog;  
再根据每个库的要求单独实现日志接口，例如glog实现gorm的日志接口；  
在没有日志集中管理分析工具时，可以使用此工具库的分组日志记录，以方便查询日志;


记录三部分日志：  
1. 正常日志 包含所有日志记录
2. 异常日志 只有异常日志记录
3. 以分组id为日志文件名，记录分组明细日志 只有context 且具有分组名时才会有记录


```go
package main

import (
	"context"
	"fmt"
	"github.com/cnWangFei/go-Wlog/ilog"
	"github.com/cnWangFei/go-Wlog/wlog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log := wlog.New()
	path := "C:\\log"
	// 一般日志
	file, err := rotatelogs.New(path+"/%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	writer := io.MultiWriter(os.Stdout, file)
	log.SetOutput(writer)
	// 异常日志
	errorFile, err := rotatelogs.New(path+"/error_%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	log.SetOutputForError(errorFile)

	log.SetLevel(ilog.DebugLevel)        // 设置日志级别
	log.SetReportCaller(true, "example") // 设置是否打印调用者

	// 分组日志
	log.SetGroupLog(filepath.Join(path, "details"), 7*24*time.Hour)

	// 记录一般日志
	log.Infoln("info", "这是一条一般日志")
	// 记录异常日志
	log.Errorln("error", "这是一条异常日志")
	// 记录分组日志，相同分组名的日志会打印到同一个文件记录
	ctx := context.Background()
	ctx = log.SetCtxGroupName(ctx, "test_log") // 设置日志分组名
	//log.PanicContextln(ctx, "panic < debug 显示")
	log.FatalContextln(ctx, "fatal < debug", "显示")
	log.ErrorContextln(ctx, "error < debug", "显示")
	log.WarnContextln(ctx, "warn < debug", "显示")
	log.InfoContextln(ctx, "info < debug", "显示")
	log.DebugContextln(ctx, "debug = debug", "显示")
	log.TraceContextln(ctx, "trace > debug", "不显示")

	log.Wait()
}
```

