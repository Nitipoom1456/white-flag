package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

func F(key string, value any) Field {
	return zap.Any(key, value)
}

type Logger struct {
	Zap *zap.Logger
}

var global = &Logger{Zap: zap.NewNop()}

func NewLogger() (*Logger, error) {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)

	log := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.DPanicLevel),
	)

	return &Logger{Zap: log}, nil
}

func SetDefault(l *Logger) { global = l }

func Sync() { _ = global.Zap.Sync() }

func Debug(ctx context.Context, msg string, fields ...Field) {
	global.Zap.Debug(msg, merge(ctx, fields)...)
}

func Info(ctx context.Context, msg string, fields ...Field) {
	global.Zap.Info(msg, merge(ctx, fields)...)
}

func Warn(ctx context.Context, msg string, fields ...Field) {
	global.Zap.Warn(msg, merge(ctx, fields)...)
}

func Error(ctx context.Context, msg string, err error, fields ...Field) {
	global.Zap.Error(msg, append(merge(ctx, fields), zap.Error(err))...)
}

type ctxKey struct{}

func With(ctx context.Context, fields ...Field) context.Context {
	return context.WithValue(ctx, ctxKey{}, merge(ctx, fields))
}

func fromContext(ctx context.Context) []Field {
	if ctx == nil {
		return nil
	}
	fields, _ := ctx.Value(ctxKey{}).([]Field)
	return fields
}

func merge(ctx context.Context, fields []Field) []Field {
	carried := fromContext(ctx)
	if len(carried) == 0 {
		return fields
	}
	out := make([]Field, 0, len(carried)+len(fields))
	out = append(out, carried...)
	return append(out, fields...)
}
