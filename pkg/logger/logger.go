package logger

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"app/pkg/conf"
)

var (
	L           *zap.Logger
	AtomicLevel = zap.NewAtomicLevel()
)

func Init(ctx context.Context, cfg *conf.LogConfig) (err error) {
	err = initLogger(cfg)
	if err != nil {
		return
	}

	L = zap.L()

	go func() {
		<-ctx.Done()
		err = L.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}

func getEncoder(format string) zapcore.Encoder {

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	if strings.ToUpper(format) == "JSON" {
		return zapcore.NewJSONEncoder(encodeConfig)
	}

	return zapcore.NewConsoleEncoder(encodeConfig)
}

// getLogWriter define the rotate log config
func getLogWriter(cfg *conf.LogConfig) zapcore.WriteSyncer {
	logRotate := &lumberjack.Logger{
		Filename:   cfg.Path,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxAge,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	syncFile := zapcore.AddSync(logRotate) // write to file
	if cfg.ConsoleEnable {
		syncConsole := zapcore.AddSync(os.Stdout) // write to std
		return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
	}

	return zapcore.AddSync(logRotate)
}

func initLogger(cfg *conf.LogConfig) (err error) {
	writeSyncer := getLogWriter(cfg)
	encoder := getEncoder(cfg.Format)

	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return err
	}
	AtomicLevel.SetLevel(level.Level())

	core := zapcore.NewCore(encoder, writeSyncer, AtomicLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)

	return
}
