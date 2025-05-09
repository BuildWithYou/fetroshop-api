package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

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

func (lg *Logger) logFileTemplate(useStackTrace bool) *logrus.Entry {
	entry := lg.LogFile.WithField("module", lg.module)
	if useStackTrace {
		stackTrace := getCallerFilePath()
		entry = entry.WithField("stackTrace", stackTrace)
	}
	return entry
}

func (lg *Logger) logConsoleTemplate(useStackTrace bool) *logrus.Entry {
	entry := lg.LogConsole.WithField("module", lg.module)
	if useStackTrace {
		stackTrace := getCallerFilePath()
		entry = entry.WithField("stackTrace", stackTrace)
	}
	return entry
}

func (lg *Logger) Trace(args ...interface{}) {
	lg.logFileTemplate(true).Trace(args...)
	lg.logConsoleTemplate(true).Trace(args...)
}

func (lg *Logger) Debug(args ...interface{}) {
	lg.logFileTemplate(true).Debug(args...)
	lg.logConsoleTemplate(true).Debug(args...)
}

func (lg *Logger) Info(args ...interface{}) {
	lg.logFileTemplate(false).Info(args...)
	lg.logConsoleTemplate(false).Info(args...)
}

func (lg *Logger) Warning(args ...interface{}) {
	lg.logFileTemplate(true).Warning(args...)
	lg.logConsoleTemplate(true).Warning(args...)
}

func (lg *Logger) Error(args ...interface{}) {
	lg.logFileTemplate(true).Error(args...)
	lg.logConsoleTemplate(true).Error(args...)
}
func (lg *Logger) UseError(err error) {
	lg.logFileTemplate(true).Error(err.Error())
	lg.logConsoleTemplate(true).Error(err.Error())
}

func (lg *Logger) Fatal(args ...interface{}) {
	errorArg := append([]interface{}{"Panic Error: "}, args...)
	lg.logConsoleTemplate(true).Error(errorArg...)
	lg.logFileTemplate(true).Fatal(args...)
}

func (lg *Logger) Panic(args ...interface{}) {
	errorArg := append([]interface{}{"Panic Error: "}, args...)
	lg.logConsoleTemplate(true).Error(errorArg...)
	lg.logFileTemplate(true).Panic(args...)
}
