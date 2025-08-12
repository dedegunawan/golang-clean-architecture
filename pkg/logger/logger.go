package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct{ *zap.SugaredLogger }

func New(level string) *Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	lvl := new(zapcore.Level)
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		*lvl = zapcore.InfoLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(*lvl)
	l, _ := cfg.Build()
	return &Logger{l.Sugar()}
}

func (l *Logger) Sync() { _ = l.SugaredLogger.Sync() }
