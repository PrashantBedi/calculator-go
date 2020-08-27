package logger

import (
	// "log"
	"os"
	"path"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// InfoLogger *log.Logger
	Logger *zap.Logger
	lumberjackLogger *lumberjack.Logger
)

func InitializeLogger() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    // if err != nil {
	//   log.Fatal(err)
	// }
	Logger, _ = zap.NewProduction()
	// Logger, _ = zap.New(
	// 	zap.Object(),
	// 	zap.AddCaller()
	// )
	Configure("logs.txt","/Users/prashantbedi/workspace-go/calculator")
	// InfoLogger = log.New(file, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Configure(logfilename string, logdir string) {
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
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encCfg)

	return zap.New(zapcore.NewCore(encoder, output, zap.NewAtomicLevel()))
}