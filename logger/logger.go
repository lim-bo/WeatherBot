package logger

import (
	"context"
	"os"

	"log/slog"
)

// Custom wrap for slog logger
type SLogger struct {
	l *slog.Logger
}

const (
	LevelApi   = slog.Level(12)
	LevelBot   = slog.Level(13)
	LevelFatal = slog.Level(14)
	LevelTrace = slog.Level(-8)
	LevelInfo  = slog.LevelInfo
	LevelError = slog.LevelError
	LevelDebug = slog.LevelDebug
)

// mapping names to levels
var levelNames = map[slog.Leveler]string{
	LevelApi:   "API",
	LevelBot:   "BOT",
	LevelFatal: "FATAL",
	LevelTrace: "TRACE",
}

// Returns new slog wrap
func New() *SLogger {
	opts := slog.HandlerOptions{
		Level: LevelTrace,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exist := levelNames[level]
				if !exist {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	return &SLogger{
		l: l,
	}
}

// Logs error as fatal and shut down app
func (sl *SLogger) Fatal(ctx context.Context, err error) {
	sl.l.Log(ctx, LevelFatal, "error: "+err.Error())
	os.Exit(1)
}

// Logs at info level
func (sl *SLogger) Info(ctx context.Context, msg string) {
	sl.l.Log(ctx, LevelInfo, msg)
}

// Logs error
func (sl *SLogger) Error(ctx context.Context, err error) {
	sl.l.Log(ctx, LevelError, err.Error())
}

// Logs any message at level,
// which provided from current package
func (sl *SLogger) LogMsgAtLevel(ctx context.Context, level slog.Level, msg string) {
	sl.l.Log(ctx, level, msg)
}

// Logs attr group at level,
// which provided from current package
func (sl *SLogger) LogWithGroupAtLevel(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	sl.l.LogAttrs(ctx, level, msg, attrs...)
}
