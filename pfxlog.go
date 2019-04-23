package pfxlog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"time"
)

var prefix string

func Global(level logrus.Level) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
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

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

func AttachedLogger(attached string) *logrus.Entry {
	return logrus.StandardLogger().WithField("attachment", attached)
}
