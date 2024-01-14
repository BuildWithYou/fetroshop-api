package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerFramework Logger

type LoggerFrameworkImpl struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
}

func NewFrameworkLogger() LoggerFramework {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("logs/framework/fetroshop-api-%s.log", today)
	logger := newLogger(pathFile, "trace")
	return &LoggerFrameworkImpl{
		LogFile:    logger.LogFile,
		LogConsole: logger.LogConsole,
	}
}

func (logger *LoggerFrameworkImpl) Trace(args ...interface{}) {
	logger.LogFile.Trace(args...)
	logger.LogConsole.Trace(args...)
}

func (logger *LoggerFrameworkImpl) Debug(args ...interface{}) {
	logger.LogFile.Debug(args...)
	logger.LogConsole.Debug(args...)
}

func (logger *LoggerFrameworkImpl) Info(args ...interface{}) {
	logger.LogFile.Info(args...)
	logger.LogConsole.Info(args...)
}

func (logger *LoggerFrameworkImpl) Warning(args ...interface{}) {
	logger.LogFile.Warning(args...)
	logger.LogConsole.Warning(args...)
}

func (logger *LoggerFrameworkImpl) Error(args ...interface{}) {
	logger.LogFile.Error(args...)
	logger.LogConsole.Error(args...)
}

func (logger *LoggerFrameworkImpl) Fatal(args ...interface{}) {
	logger.LogFile.Fatal(args...)
	logger.LogConsole.Fatal(args...)
}

func (logger *LoggerFrameworkImpl) Panic(args ...interface{}) {
	logger.LogFile.Panic(args...)
	logger.LogConsole.Panic(args...)
}
