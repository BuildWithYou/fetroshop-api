package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LoggerWeb Logger

type LoggerWebImpl struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
}

func NewWebLogger(config *viper.Viper) LoggerWeb {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("logs/web/fetroshop-web-%s.log", today)
	logger := newLogger(pathFile, config.GetString("app.web.logLevel"))
	return &LoggerWebImpl{
		LogFile:    logger.LogFile,
		LogConsole: logger.LogConsole,
	}
}

func (logger *LoggerWebImpl) Trace(args ...interface{}) {
	logger.LogFile.Trace(args...)
	logger.LogConsole.Trace(args...)
}

func (logger *LoggerWebImpl) Debug(args ...interface{}) {
	logger.LogFile.Debug(args...)
	logger.LogConsole.Debug(args...)
}

func (logger *LoggerWebImpl) Info(args ...interface{}) {
	logger.LogFile.Info(args...)
	logger.LogConsole.Info(args...)
}

func (logger *LoggerWebImpl) Warning(args ...interface{}) {
	logger.LogFile.Warning(args...)
	logger.LogConsole.Warning(args...)
}

func (logger *LoggerWebImpl) Error(args ...interface{}) {
	logger.LogFile.Error(args...)
	logger.LogConsole.Error(args...)
}

func (logger *LoggerWebImpl) Fatal(args ...interface{}) {
	logger.LogFile.Fatal(args...)
	logger.LogConsole.Fatal(args...)
}

func (logger *LoggerWebImpl) Panic(args ...interface{}) {
	logger.LogFile.Panic(args...)
	logger.LogConsole.Panic(args...)
}
