package logger

import (
	"fmt"
	"os"
	"runtime"
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

func getCallerFilePath() string {
	_, file, line, _ := runtime.Caller(2)
	return string(fmt.Sprint(file, ":", line))
}
