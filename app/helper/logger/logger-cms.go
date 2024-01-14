package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LoggerCms Logger

type LoggerCmsImpl struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
}

func NewCmsLogger(config *viper.Viper) LoggerCms {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("logs/cms/fetroshop-cms-%s.log", today)
	var logger = newLogger(pathFile, config.GetString("app.cms.logLevel"))
	return &LoggerCmsImpl{
		LogFile:    logger.LogFile,
		LogConsole: logger.LogConsole,
	}
}

func (logger *LoggerCmsImpl) Trace(args ...interface{}) {
	logger.LogFile.Trace(args...)
	logger.LogConsole.Trace(args...)
}

func (logger *LoggerCmsImpl) Debug(args ...interface{}) {
	logger.LogFile.Debug(args...)
	logger.LogConsole.Debug(args...)
}

func (logger *LoggerCmsImpl) Info(args ...interface{}) {
	logger.LogFile.Info(args...)
	logger.LogConsole.Info(args...)
}

func (logger *LoggerCmsImpl) Warning(args ...interface{}) {
	logger.LogFile.Warning(args...)
	logger.LogConsole.Warning(args...)
}

func (logger *LoggerCmsImpl) Error(args ...interface{}) {
	logger.LogFile.Error(args...)
	logger.LogConsole.Error(args...)
}

func (logger *LoggerCmsImpl) Fatal(args ...interface{}) {
	logger.LogFile.Fatal(args...)
	logger.LogConsole.Fatal(args...)
}

func (logger *LoggerCmsImpl) Panic(args ...interface{}) {
	logger.LogFile.Panic(args...)
	logger.LogConsole.Panic(args...)
}
