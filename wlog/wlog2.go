package wlog

import (
	"context"
	"github.com/cnWangFei/go-Wlog/ilog"
)

func (w *Wlog) Panic(args ...any) {
	w.Log(ilog.PanicLevel, args...)
}

func (w *Wlog) Panicln(args ...any) {
	w.Logln(ilog.PanicLevel, args...)
}

func (w *Wlog) Panicf(format string, args ...any) {
	w.Logf(ilog.PanicLevel, format, args...)
}

func (w *Wlog) PanicContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.PanicLevel, args...)
}

func (w *Wlog) PanicContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.PanicLevel, args...)
}

func (w *Wlog) PanicContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.PanicLevel, format, args...)
}

func (w *Wlog) Fatal(args ...any) {
	w.Log(ilog.FatalLevel, args...)
}

func (w *Wlog) Fatalln(args ...any) {
	w.Logln(ilog.FatalLevel, args...)
}

func (w *Wlog) Fatalf(format string, args ...any) {
	w.Logf(ilog.FatalLevel, format, args...)
}

func (w *Wlog) FatalContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.FatalLevel, args...)
}

func (w *Wlog) FatalContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.FatalLevel, args...)
}

func (w *Wlog) FatalContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.FatalLevel, format, args...)
}

func (w *Wlog) Error(args ...any) {
	w.Log(ilog.ErrorLevel, args...)
}

func (w *Wlog) Errorln(args ...any) {
	w.Logln(ilog.ErrorLevel, args...)
}

func (w *Wlog) Errorf(format string, args ...any) {
	w.Logf(ilog.ErrorLevel, format, args...)
}

func (w *Wlog) ErrorContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.ErrorLevel, args...)
}

func (w *Wlog) ErrorContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.ErrorLevel, args...)
}

func (w *Wlog) ErrorContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.ErrorLevel, format, args...)
}

func (w *Wlog) Warn(args ...any) {
	w.Log(ilog.WarnLevel, args...)
}

func (w *Wlog) Warnln(args ...any) {
	w.Logln(ilog.WarnLevel, args...)
}

func (w *Wlog) Warnf(format string, args ...any) {
	w.Logf(ilog.WarnLevel, format, args...)
}

func (w *Wlog) WarnContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.WarnLevel, args...)
}

func (w *Wlog) WarnContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.WarnLevel, args...)
}

func (w *Wlog) WarnContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.WarnLevel, format, args...)
}

func (w *Wlog) Info(args ...any) {
	w.Log(ilog.InfoLevel, args...)
}

func (w *Wlog) Infoln(args ...any) {
	w.Logln(ilog.InfoLevel, args...)
}

func (w *Wlog) Infof(format string, args ...any) {
	w.Logf(ilog.InfoLevel, format, args...)
}

func (w *Wlog) InfoContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.InfoLevel, args...)
}

func (w *Wlog) InfoContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.InfoLevel, args...)
}

func (w *Wlog) InfoContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.InfoLevel, format, args...)
}

func (w *Wlog) Debug(args ...any) {
	w.Log(ilog.DebugLevel, args...)
}

func (w *Wlog) Debugln(args ...any) {
	w.Logln(ilog.DebugLevel, args...)
}

func (w *Wlog) Debugf(format string, args ...any) {
	w.Logf(ilog.DebugLevel, format, args...)
}

func (w *Wlog) DebugContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.DebugLevel, args...)
}

func (w *Wlog) DebugContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.DebugLevel, args...)
}

func (w *Wlog) DebugContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.DebugLevel, format, args...)
}

func (w *Wlog) Trace(args ...any) {
	w.Log(ilog.TraceLevel, args...)
}

func (w *Wlog) Traceln(args ...any) {
	w.Logln(ilog.TraceLevel, args...)
}

func (w *Wlog) Tracef(format string, args ...any) {
	w.Logf(ilog.TraceLevel, format, args...)
}

func (w *Wlog) TraceContext(ctx context.Context, args ...any) {
	w.LogContext(ctx, ilog.TraceLevel, args...)
}

func (w *Wlog) TraceContextln(ctx context.Context, args ...any) {
	w.LogContextln(ctx, ilog.TraceLevel, args...)
}

func (w *Wlog) TraceContextf(ctx context.Context, format string, args ...any) {
	w.LogContextf(ctx, ilog.TraceLevel, format, args...)
}
