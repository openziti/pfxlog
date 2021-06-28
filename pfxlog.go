package pfxlog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

var prefix string

func Global(level logrus.Level, options *Options) {
	noJson, _ := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_NO_JSON")))
	if noJson || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(NewFormatterWithOptions(options))
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: options.JsonTimestampFormat})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
}

func ContextLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}
