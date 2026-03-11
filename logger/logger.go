package logger

import (
	"fmt"
	"net/url"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DebugLevel   = "debug"
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"
)

var globalLogger *zap.Logger

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

// InitLogger нь logger-ийг эхлүүлнэ.
//
// logLevel: "debug", "info", "warning", "error"
// logFile: лог файлын зам
// isProd: production горим бол stacktrace нэмэхгүй
func InitLogger(logLevel string, logFile string, isProd bool) error {
	var level zapcore.Level
	switch logLevel {
	case DebugLevel:
		level = zap.DebugLevel
	case InfoLevel:
		level = zap.InfoLevel
	case WarningLevel:
		level = zap.WarnLevel
	case ErrorLevel:
		level = zap.ErrorLevel
	default:
		return fmt.Errorf("unknown log level: %s", logLevel)
	}

	writer := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1024, // MB
		MaxBackups: 30,
		MaxAge:     90, // days
		Compress:   true,
	}

	zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{Logger: &writer}, nil
	})

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(&writer), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	if isProd {
		globalLogger = zap.New(core, zap.AddCaller())
	} else {
		globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	}

	zap.ReplaceGlobals(globalLogger)
	return nil
}

// Info нь info түвшний лог бичнэ.
func Info(message string, fields ...zap.Field) {
	globalLogger.Info(message, fields...)
}

// Debug нь debug түвшний лог бичнэ.
func Debug(message string, fields ...zap.Field) {
	globalLogger.Debug(message, fields...)
}

// Error нь error түвшний лог бичнэ.
func Error(message string, fields ...zap.Field) {
	globalLogger.Error(message, fields...)
}

// Warn нь warning түвшний лог бичнэ.
func Warn(message string, fields ...zap.Field) {
	globalLogger.Warn(message, fields...)
}

// Fatal нь fatal түвшний лог бичнэ. Програм зогсоно.
func Fatal(message string, fields ...zap.Field) {
	globalLogger.Fatal(message, fields...)
}

// NewSugar нь нэртэй SugaredLogger үүсгэнэ.
func NewSugar(name string) *zap.SugaredLogger {
	return globalLogger.Named(name).Sugar()
}

// GetLogger нь global logger-ийг буцаана.
func GetLogger() *zap.Logger {
	return globalLogger
}
