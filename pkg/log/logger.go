package log

import (
	"os"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type logger struct {
	sugarLogger *zap.SugaredLogger
}

// NewLogger - creates a new logger
func NewLogger(cfg helpers.AppConfig) Logger {
	logLevel := getLoggerLevel(cfg)
	logWriter := zapcore.AddSync(os.Stderr)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.TimeKey = "time"
	encoderCfg.NameKey = "name"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	encoder := zapcore.NewJSONEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	newLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	sugarLogger := newLogger.Sugar()
	if err := sugarLogger.Sync(); err != nil {
		sugarLogger.Error(err)
	}
	return &logger{
		sugarLogger: sugarLogger,
	}
}

// Logger methods
func (l *logger) Log(keyvals ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
