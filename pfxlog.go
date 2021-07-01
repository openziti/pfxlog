package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

func GlobalInit(level logrus.Level, options *Options) {
	noJson, err := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_NO_JSON")))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "pfxlog: error parsing 'PFXLOG_NO_JSON' (%v)", err)
	}
	if (err == nil && noJson) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(NewFormatter(options))
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: options.JsonTimestampFormat})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

func ContextLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}
