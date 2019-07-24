package pfxlog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
	"time"
)

var prefix string

func Global(level logrus.Level) {
	noJson, _ := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_NO_JSON")))
	if noJson || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(&Formatter{start: time.Now()})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
}

func SetPrefix(p string) {
	prefix = p
}

func ContextLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

// Deprecated: Use ContextLogger instead.
//
func AttachedLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}
