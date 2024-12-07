package glog

import (
	"context"
	"fmt"
	"github.com/cnWangFei/go-Wlog/ilog"
	"github.com/cnWangFei/go-Wlog/wlog"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type Glog struct {
	wlog *wlog.Wlog
}

func New(wllog *wlog.Wlog) *Glog {
	return &Glog{wlog: wllog}
}
func (g *Glog) LogMode(level gormLog.LogLevel) gormLog.Interface {
	g.wlog.SetLevel(ilog.Level(level))
	return g
}
func (g *Glog) Info(ctx context.Context, format string, args ...interface{}) {
	g.wlog.InfoContextf(ctx, format, args...)
}
func (g *Glog) Warn(ctx context.Context, format string, args ...interface{}) {
	g.wlog.WarnContextf(ctx, format, args...)
}
func (g *Glog) Error(ctx context.Context, format string, args ...interface{}) {
	g.wlog.ErrorContextf(ctx, format, args...)
}
func (g *Glog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.wlog.Level <= ilog.FatalLevel {
		return
	}
	// 执行时间
	elapsed := time.Since(begin)
	elapsedStr := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

	sql, rows := fc()

	ctx = context.WithValue(ctx, "elapsed", elapsedStr) // 执行时间
	ctx = context.WithValue(ctx, "rows", rows)          // 影响行数

	g.wlog.DebugContextln(ctx, sql)
}
