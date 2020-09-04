package logger

import (
	"os"
	"path"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
	lumberjackLogger *lumberjack.Logger
)

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
	configure("logs.txt","/Users/prashantbedi/workspace-go/calculator")
}

func configure(logfilename string, logdir string) {
	writers := []zapcore.WriteSyncer{os.Stdout}
	if logdir != "" {
		writers = append(writers, newRollingFile(logfilename, logdir))
	}

	Logger = newZapLogger(zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(Logger)
	Logger.Info("logging configured",
		zap.Bool("fileLogging", logdir != ""),
		zap.String("logDirectory", logdir),
		zap.String("logFilename", logfilename))
}

func newRollingFile(logfilename string, logdir string) zapcore.WriteSyncer {
	if err := os.MkdirAll(logdir, 0744); err != nil {
		Logger.Error("failed create log directory",
			zap.Error(err),
			zap.String("path", logdir))
		return nil
	}

	lumberjackLogger = &lumberjack.Logger{
		Filename:   path.Join(logdir, logfilename),
		MaxSize:    5,  //megabytes
		MaxAge:     10, //days
		MaxBackups: 5,  //files
	}
	return zapcore.AddSync(lumberjackLogger)
}

func newZapLogger(output zapcore.WriteSyncer) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "logger_name",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encCfg)
	var core = zapcore.NewTee(zapcore.NewCore(encoder, output, zap.NewAtomicLevel()))
	return zap.New(core, zap.AddCaller())
}