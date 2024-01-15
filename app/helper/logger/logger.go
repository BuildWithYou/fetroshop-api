package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Logger struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
}

const frameworkBasePath = "logs/framework/fetroshop-api"
const webBasePath = "logs/web/fetroshop-web"
const cmsBasePath = "logs/cms/fetroshop-cms"

func newLogger(pathFile string, levelStr string) *Logger {
	var level logrus.Level

	switch levelStr {
	case "trace":
		{
			level = logrus.TraceLevel
		}
	case "debug":
		{
			level = logrus.DebugLevel
		}
	case "info":
		{
			level = logrus.InfoLevel
		}
	case "warning":
		{
			level = logrus.WarnLevel
		}
	case "error":
		{
			level = logrus.ErrorLevel
		}
	case "fatal":
		{
			level = logrus.FatalLevel
		}
	case "panic":
		{
			level = logrus.PanicLevel
		}
	default:
		{
			level = logrus.InfoLevel
		}
	}

	logFile := logrus.New()
	logFile.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logFile.SetOutput(file)
	logFile.SetLevel(level)

	logConsole := logrus.New()
	logConsole.SetFormatter(&logrus.JSONFormatter{})
	logConsole.SetOutput(os.Stdout)
	logConsole.SetLevel(level)

	return &Logger{
		LogFile:    logFile,
		LogConsole: logConsole,
	}
}

var FrameworkLogger, WebLogger, CmsLogger *Logger

func NewFrameworkLogger() *Logger {
	if FrameworkLogger == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", frameworkBasePath, today)
		FrameworkLogger = newLogger(pathFile, "trace")
	}
	return FrameworkLogger
}

func NewWebLogger(config *viper.Viper) *Logger {
	if WebLogger == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", webBasePath, today)
		WebLogger = newLogger(pathFile, config.GetString("app.web.logLevel"))
	}
	return WebLogger
}

func NewCmsLogger(config *viper.Viper) *Logger {
	if CmsLogger == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", cmsBasePath, today)
		CmsLogger = newLogger(pathFile, config.GetString("app.cms.logLevel"))
	}
	return CmsLogger
}

func (logger *Logger) FrameworkLoggerResetOutput() {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("%s-%s.log", frameworkBasePath, today)
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.LogFile.SetOutput(file)
}

func (logger *Logger) WebLoggerResetOutput() {
	logger.FrameworkLoggerResetOutput()
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("%s-%s.log", webBasePath, today)
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.LogFile.SetOutput(file)
}

func (logger *Logger) CmsLoggerResetOutput() {
	logger.FrameworkLoggerResetOutput()
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("%s-%s.log", cmsBasePath, today)
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.LogFile.SetOutput(file)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.LogFile.Trace(args...)
	logger.LogConsole.Trace(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.LogFile.Debug(args...)
	logger.LogConsole.Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.LogFile.Info(args...)
	logger.LogConsole.Info(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.LogFile.Warning(args...)
	logger.LogConsole.Warning(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.LogFile.Error(args...)
	logger.LogConsole.Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.LogFile.Fatal(args...)
	logger.LogConsole.Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.LogFile.Panic(args...)
	logger.LogConsole.Panic(args...)
}
