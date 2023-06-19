package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
)

const (
	CONSOLE = "console"
	FILE    = "file"
)

var (
	Level  = zap.DebugLevel
	Target = FILE
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "name",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "trace",
		// LineEnding:     zapcore.DPanicLevel.CapitalString(),
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func init() {
	w := zapcore.AddSync(&lumberjackv2.Logger{
		Filename:   "log/live.log", //位置
		MaxSize:    300,            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,             //保留旧文件的最大个数
		MaxAge:     7,              //保留旧文件的最大天数
	})
	var writeSyncer zapcore.WriteSyncer
	if Target == CONSOLE {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}
	if Target == FILE {
		writeSyncer = zapcore.NewMultiWriteSyncer(w)
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(NewEncoderConfig()),
		writeSyncer,
		Level,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}
