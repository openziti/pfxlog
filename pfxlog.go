package pfxlog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func GlobalInit(level logrus.Level, options *Options) {
	if defaultEnv("PFXLOG_NO_JSON", false) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(NewFormatter(options))
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: options.JsonTimestampFormat})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	globalOptions = options
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(globalOptions.StandardLogger)
}

func ContextLogger(context string) *logrus.Entry {
	return globalOptions.StandardLogger.WithField("context", context)
}

func ContextDataLogger(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextDataFielder != nil {
		return globalOptions.ContextDataFielder(contextData, logrus.StandardLogger())
	} else {
		return globalOptions.StandardLogger.WithFields(nil)
	}
}

func ContextCheck(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextDataChecker != nil && globalOptions.ContextDataChecker(contextData) {
		return globalOptions.StandardLogger.WithFields(nil)
	} else {
		return &logrus.Entry{Logger: globalOptions.NoLogger}
	}
}

func ContextCheckData(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextDataChecker != nil && globalOptions.ContextDataChecker(contextData) {
		return ContextDataLogger(contextData)
	} else {
		return &logrus.Entry{Logger: globalOptions.NoLogger}
	}
}

var globalOptions *Options
