package ilog

import (
	"context"
	"fmt"
	"io"
)

// Level 日志级别
type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid logrus level %d", level)
}

type ILog interface {
	// SetOutput 设置日志输出
	SetOutput(output io.Writer)
	// SetLevel 设置日志级别
	SetLevel(level Level)

	Log(level Level, args ...any)
	Logln(level Level, args ...any)
	Logf(level Level, format string, args ...any)
	LogContext(ctx context.Context, level Level, args ...any)
	LogContextln(ctx context.Context, level Level, args ...any)
	LogContextf(ctx context.Context, level Level, format string, args ...any)

	Panic(args ...any)
	Panicln(args ...any)
	Panicf(format string, args ...any)
	PanicContext(ctx context.Context, args ...any)
	PanicContextln(ctx context.Context, args ...any)
	PanicContextf(ctx context.Context, format string, args ...any)

	Fatal(args ...any)
	Fatalln(args ...any)
	Fatalf(format string, args ...any)
	FatalContext(ctx context.Context, args ...any)
	FatalContextln(ctx context.Context, args ...any)
	FatalContextf(ctx context.Context, format string, args ...any)

	Error(args ...any)
	Errorln(args ...any)
	Errorf(format string, args ...any)
	ErrorContext(ctx context.Context, args ...any)
	ErrorContextln(ctx context.Context, args ...any)
	ErrorContextf(ctx context.Context, format string, args ...any)

	Warn(args ...any)
	Warnln(args ...any)
	Warnf(format string, args ...any)
	WarnContext(ctx context.Context, args ...any)
	WarnContextln(ctx context.Context, args ...any)
	WarnContextf(ctx context.Context, format string, args ...any)

	Info(args ...any)
	Infoln(args ...any)
	Infof(format string, args ...any)
	InfoContext(ctx context.Context, args ...any)
	InfoContextln(ctx context.Context, args ...any)
	InfoContextf(ctx context.Context, format string, args ...any)

	Debug(args ...any)
	Debugln(args ...any)
	Debugf(format string, args ...any)
	DebugContext(ctx context.Context, args ...any)
	DebugContextln(ctx context.Context, args ...any)
	DebugContextf(ctx context.Context, format string, args ...any)

	Trace(args ...any)
	Traceln(args ...any)
	Tracef(format string, args ...any)
	TraceContext(ctx context.Context, args ...any)
	TraceContextln(ctx context.Context, args ...any)
	TraceContextf(ctx context.Context, format string, args ...any)
}
