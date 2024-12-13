package wlog_test

import (
	"context"
	"github.com/cnWangFei/go-Wlog/ilog"
	"github.com/cnWangFei/go-Wlog/wlog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	log := wlog.New()
	path := "C:\\log"
	// 分组日志
	log.SetGroupLog(filepath.Join(path, "details"), 7*24*time.Hour)
	// 一般日志
	file, err := rotatelogs.New(path+"/%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		t.Fatal(err)
		return
	}
	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)
	// 异常日志
	errorFile, err := rotatelogs.New(path+"/error_%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.SetOutputForError(errorFile)

	log.SetLevel(ilog.DebugLevel)
	//log.SetReportCaller(true, "go-wlog")
	//log.Panicln("panic < debug 显示")
	log.Fatalln("fatal < debug 显示")
	log.Errorln("error < debug 显示")
	log.Warnln("warn < debug 显示")
	log.Infoln("info < debug 显示")
	log.Debugln("debug = debug 显示")
	log.Traceln("trace > debug 不显示")

	log.Wait()
}
func TestLogContext(t *testing.T) {
	log := wlog.New()
	path := "C:\\log"
	// 分组日志
	log.SetGroupLog(filepath.Join(path, "details"), 7*24*time.Hour)
	// 一般日志
	file, err := rotatelogs.New(path+"/%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		t.Fatal(err)
		return
	}
	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)
	// 异常日志
	errorFile, err := rotatelogs.New(path+"/error_%Y%m%d%H%M%S.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.SetOutputForError(errorFile)

	log.SetLevel(ilog.DebugLevel)
	log.SetReportCaller(true, "go-wlog")
	ctx := context.Background()
	ctx = log.SetCtxGroupName(ctx, "test_log")
	//log.PanicContextln(ctx, "panic < debug 显示")
	log.FatalContextln(ctx, "fatal < debug", "显示")
	log.ErrorContextln(ctx, "error < debug", "显示")
	log.WarnContextln(ctx, "warn < debug", "显示")
	log.InfoContextln(ctx, "info < debug", "显示")
	log.DebugContextln(ctx, "debug = debug", "显示")
	log.TraceContextln(ctx, "trace > debug", "不显示")

	ctx = context.WithValue(ctx, "log1", "test log1")
	ctx = context.WithValue(ctx, "log2", "test log2")
	ctx = context.WithValue(ctx, "log3", "test log3")
	ctx = context.WithValue(ctx, "log4", "test log4")
	log.InfoContextln(ctx, "")

	log.Wait()
}
