package logger

import (
	"fmt"
	"os"
	"time"
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
func (lg *Logger) UseError(err error) {
	/* stackTrace := errorhelper.GetStackTrace(err)
	lg.LogFile.WithField("module", lg.module).WithField("stackTrace", stackTrace).Error(err.Error())
	lg.LogConsole.WithField("module", lg.module).WithField("stackTrace", stackTrace).Error(err.Error()) */

	lg.LogFile.WithField("module", lg.module).Error(err.Error())
	lg.LogConsole.WithField("module", lg.module).Error(err.Error())
}

func (lg *Logger) Fatal(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Fatal(args...)
	lg.LogConsole.WithField("module", lg.module).Fatal(args...)
}

func (lg *Logger) Panic(args ...interface{}) {
	lg.LogFile.WithField("module", lg.module).Panic(args...)
	lg.LogConsole.WithField("module", lg.module).Panic(args...)
}
