package wlog

import (
	"context"
	"fmt"
	"github.com/cnWangFei/go-Wlog/ilog"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Wlog struct {
	output    io.Writer
	Level     ilog.Level
	log       *logrus.Logger // 正常日志输出
	errOutput io.Writer      // error日志单独输出
	errLog    *logrus.Logger // error日志
	// 分组日志相关 会新建一个文件夹，把每个分组日志id相关的日志写入到一个独立的日志文件
	isWriteGroupLog bool   // 是否写分组日志
	groupLogPath    string // 分组日志路径
	groupLogMaxAge  time.Duration
	wg              sync.WaitGroup // 用于等待写分组日志协程结束
}

const Group = "groupId" // 用于区分分组日志的字段名

var (
	groupLogCh          chan int   // 处理同时写日志的协程数
	removeGroupLogLock  sync.Mutex // 为删除历史记录加锁
	reportCaller        bool       // 是否打印调用者
	callerProjectFolder string     // 调用wlog的项目文件夹名称
)

func New() *Wlog {
	w := new(Wlog)
	w.log = logrus.New()
	w.errLog = logrus.New()
	return w
}

// SetOutput 设置日志输出
func (w *Wlog) SetOutput(output io.Writer) {
	writer := io.MultiWriter(output)
	if w.output != nil {
		writer = io.MultiWriter(w.output, output)
	}
	w.output = writer
	w.log.SetOutput(writer)
}

// SetOutputForError 设置error日志输出
func (w *Wlog) SetOutputForError(output io.Writer) {
	writer := io.MultiWriter(output)
	if w.errOutput != nil {
		writer = io.MultiWriter(w.errOutput, output)
	}
	w.errOutput = writer
	w.errLog.SetOutput(writer)
}

// SetLevel 设置日志级别
func (w *Wlog) SetLevel(level ilog.Level) {
	w.Level = level
	w.log.SetLevel(logrus.TraceLevel)
	w.errLog.SetLevel(logrus.TraceLevel)
}

// SetReportCaller 设置是否打印调用者
// projectFolder 调用wlog的项目文件夹名称，用来去除敏感路径
func (w *Wlog) SetReportCaller(caller bool, projectFolder string) {
	reportCaller = caller
	callerProjectFolder = projectFolder
}

// SetCtxGroupName 设置分组名
func (w *Wlog) SetCtxGroupName(ctx context.Context, fileName string) context.Context {
	if ctx.Value(Group) != nil {
		panic(fmt.Sprintf("ctx already has group name: %s", fileName))
		return ctx
	}
	ctx = context.WithValue(ctx, Group, fileName)
	return ctx
}

// setgrouplog 设置分组日志
func (w *Wlog) SetGroupLog(path string, age time.Duration) {
	w.isWriteGroupLog = true
	w.groupLogPath = path
	w.groupLogMaxAge = age
	groupLogCh = make(chan int, 1000)
}

// Wait 等待所有协程结束
func (w *Wlog) Wait() {
	w.wg.Wait()
}

func (w *Wlog) Log(level ilog.Level, args ...any) {
	if w.Level < level {
		return
	}
	w.log.WithFields(ctxToMap(nil)).Log(logrus.Level(level), args...)
	// 日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(ctxToMap(nil)).Log(logrus.Level(level), args...)
	}
}

func (w *Wlog) Logln(level ilog.Level, args ...any) {
	if w.Level < level {
		return
	}
	w.log.WithFields(ctxToMap(nil)).Logln(logrus.Level(level), args...)
	// 日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(ctxToMap(nil)).Logln(logrus.Level(level), args...)
	}
}

func (w *Wlog) Logf(level ilog.Level, format string, args ...any) {
	if w.Level < level {
		return
	}
	w.log.WithFields(ctxToMap(nil)).Logf(logrus.Level(level), format, args...)
	// 日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(ctxToMap(nil)).Logf(logrus.Level(level), format, args...)
	}
}

func (w *Wlog) LogContext(ctx context.Context, level ilog.Level, args ...any) {
	if w.Level < level {
		return
	}
	m := ctxToMap(ctx)
	// 分组日志
	w.writeGroupLog(m, level, args...)
	// 一般日志
	w.log.WithFields(m).Log(logrus.Level(level), args...)
	// error日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(m).Log(logrus.Level(level), args...)
	}
}
func (w *Wlog) LogContextln(ctx context.Context, level ilog.Level, args ...any) {
	if w.Level < level {
		return
	}
	m := ctxToMap(ctx)
	// 分组日志
	w.writeGroupLogln(m, level, args...)
	// 一般日志
	w.log.WithFields(m).Logln(logrus.Level(level), args...)
	// error日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(m).Logln(logrus.Level(level), args...)
	}
}
func (w *Wlog) LogContextf(ctx context.Context, level ilog.Level, format string, args ...any) {
	if w.Level < level {
		return
	}
	m := ctxToMap(ctx)
	// 分组日志
	w.writeGroupLogf(m, level, format, args...)
	// 一般日志
	w.log.WithFields(m).Logf(logrus.Level(level), format, args...)
	// error日志单独输出
	if w.errOutput != nil && level <= ilog.ErrorLevel {
		w.errLog.WithFields(m).Logf(logrus.Level(level), format, args...)
	}
}

func (w *Wlog) writeGroupLog(m map[string]interface{}, level ilog.Level, args ...any) {
	if !w.isWriteGroupLog {
		return
	}
	group := getMapGroup(m)
	if group == "" {
		return
	}
	m["time"] = time.Now().Format(time.RFC3339)
	m["level"] = level
	msg := fmt.Sprint(args...)
	// 获取ctx中的内容，包括caller
	msg = fmt.Sprintf("%v\n%v", mapToStr(m), msg)
	w.writeGroupLogToChan(&group, []byte(msg))
}

func (w *Wlog) writeGroupLogln(m map[string]interface{}, level ilog.Level, args ...any) {
	if !w.isWriteGroupLog {
		return
	}
	group := getMapGroup(m)
	if group == "" {
		return
	}
	m["time"] = time.Now().Format(time.RFC3339)
	m["level"] = level
	msg := fmt.Sprintln(args...)
	// 获取ctx中的内容，包括caller
	msg = fmt.Sprintf("%v\n%v", mapToStr(m), msg)
	w.writeGroupLogToChan(&group, []byte(msg))
}

func (w *Wlog) writeGroupLogf(m map[string]interface{}, level ilog.Level, format string, args ...any) {
	if !w.isWriteGroupLog {
		return
	}
	group := getMapGroup(m)
	if group == "" {
		return
	}
	m["time"] = time.Now().Format(time.RFC3339)
	m["level"] = level
	msg := fmt.Sprintf(format, args...)
	// 获取ctx中的内容，包括caller
	msg = fmt.Sprintf("%v\n%v", mapToStr(m), msg)
	w.writeGroupLogToChan(&group, []byte(msg))
}
func (w *Wlog) writeGroupLogToChan(group *string, msg []byte) {
	groupLogCh <- 1
	w.wg.Add(1)
	// 开启协程记录日志，只要不超过LogCh，就不会卡
	go func(group *string, msg []byte) {
		defer w.wg.Done()
		err := w.writeGroupLogToFile(group, msg)
		if err != nil {
			w.Errorf("记录分组日志失败,err=%v,group=%v,log=%v", err.Error(), group, string(msg))
		}
		<-groupLogCh
	}(group, msg)
}

// writeGroupLogToFile 记录分组日志
func (w *Wlog) writeGroupLogToFile(group *string, msg []byte) (err error) {
	if w.groupLogPath == "" {
		err = fmt.Errorf("no group log path")
		return
	}
	if w.groupLogMaxAge.Seconds() == 0 {
		err = fmt.Errorf("log max age is 0")
		return
	}

	path := filepath.Join(w.groupLogPath, time.Now().Format("20060102_15"))
	// 创建目录
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}
	// 创建日志文件
	logFilePath := filepath.Join(path, fmt.Sprintf("%v.log", *group))
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 写入日志
	_, err = file.Write(msg)
	if err != nil {
		return
	}
	// 删除过期日志
	if removeGroupLogLock.TryLock() {
		defer removeGroupLogLock.Unlock()
		dir, e := os.ReadDir(w.groupLogPath)
		if e != nil {
			return e
		}

		cutoff := time.Now().Add(-1 * w.groupLogMaxAge)
		for _, v := range dir {
			if v.IsDir() {
				fl, e := v.Info()
				if e != nil {
					continue
				}
				if w.groupLogMaxAge > 0 && fl.ModTime().After(cutoff) {
					continue
				}
				err = os.RemoveAll(filepath.Join(w.groupLogPath, v.Name()))
				if err != nil {
					return
				}
			}
		}
	}
	return
}
