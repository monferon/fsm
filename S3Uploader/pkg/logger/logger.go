package logger

import (
	"go.uber.org/zap"
	"os"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zap.Logger
}

//
//var _ Interface = (*Logger)(nil)

//
// New -.
func New() *Logger {
	//var l zerolog.Level
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	//switch strings.ToLower(level) {
	//case "error":
	//	l = zerolog.ErrorLevel
	//case "warn":
	//	l = zerolog.WarnLevel
	//case "info":
	//	l = zerolog.InfoLevel
	//case "debug":
	//	l = zerolog.DebugLevel
	//default:
	//	l = zerolog.InfoLevel
	//}
	//
	//zerolog.SetGlobalLevel(l)
	//
	//skipFrameCount := 3
	//logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: logger,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.Debug(message, args...)
	//l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.Info(message, args...)
	//l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.Warn(message, args...)
	//l.log(message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.Error(message, args...)
	//if l.logger.GetLevel() == zerolog.DebugLevel {
	//	l.Debug(message, args...)
	//}

	//l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.Fatal(message, args...)
	//l.msg("fatal", message, args...)
	//TODO убрать exit
	os.Exit(1)
}

//func (l *Logger) log(message string, args ...interface{}) {
//	if len(args) == 0 {
//		l.logger.Info(message)
//	} else {
//		l.logger.Info(message, args...)
//	}
//}
//
//func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
//	switch msg := message.(type) {
//	case error:
//		l.log(msg.Error(), args...)
//	case string:
//		l.log(msg, args...)
//	default:
//		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
//	}
//}
