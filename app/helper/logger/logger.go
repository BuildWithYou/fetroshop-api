package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO: need to sanitize logged data

type Logger struct {
	LogFile    *logrus.Logger
	LogConsole *logrus.Logger
	module     string
}

const webBasePath = "logs/web/fetroshop-web"
const cmsBasePath = "logs/cms/fetroshop-cms"

func newLogger(module string, pathFile string, levelStr string) *Logger {
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
		module:     module,
	}
}

var WebLogger, CmsLogger *Logger

func NewWebLogger(config *viper.Viper) *Logger {
	if WebLogger == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", webBasePath, today)
		WebLogger = newLogger("web", pathFile, config.GetString("app.web.logLevel"))
	}
	return WebLogger
}

func NewCmsLogger(config *viper.Viper) *Logger {
	if CmsLogger == nil {
		today := time.Now().Format("2006-01-02")
		pathFile := fmt.Sprintf("%s-%s.log", cmsBasePath, today)
		CmsLogger = newLogger("cms", pathFile, config.GetString("app.cms.logLevel"))
	}
	return CmsLogger
}

func (lg *Logger) WebLoggerResetOutput() {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("%s-%s.log", webBasePath, today)
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	lg.LogFile.SetOutput(file)
}

func (lg *Logger) CmsLoggerResetOutput() {
	today := time.Now().Format("2006-01-02")
	pathFile := fmt.Sprintf("%s-%s.log", cmsBasePath, today)
	file, _ := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	lg.LogFile.SetOutput(file)
}

func (lg *Logger) Trace(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Trace(args...)
	lg.LogConsole.WithField("module", lg.module).Trace(args...)
}

func (lg *Logger) Debug(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Debug(args...)
	lg.LogConsole.WithField("module", lg.module).Debug(args...)
}

func (lg *Logger) Info(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Info(args...)
	lg.LogConsole.WithField("module", lg.module).Info(args...)
}

func (lg *Logger) Warning(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Warning(args...)
	lg.LogConsole.WithField("module", lg.module).Warning(args...)
}

func (lg *Logger) Error(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Error(args...)
	lg.LogConsole.WithField("module", lg.module).Error(args...)
}

func (lg *Logger) Fatal(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Fatal(args...)
	lg.LogConsole.WithField("module", lg.module).Fatal(args...)
}

func (lg *Logger) Panic(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Panic(args...)
	lg.LogConsole.WithField("module", lg.module).Panic(args...)
}
