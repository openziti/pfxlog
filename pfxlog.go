package pfxlog

import (
	"github.com/sirupsen/logrus"
)

var prefix string

func Global(level logrus.Level) {
	logrus.SetFormatter(&Formatter{})
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
	return logrus.StandardLogger().WithField("attached", attached)
}