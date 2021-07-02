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

func Logger() *Builder {
	return &Builder{logrus.NewEntry(globalOptions.StandardLogger)}
}

func ContextLogger(context string) *Builder {
	return &Builder{globalOptions.StandardLogger.WithField("context", context)}
}

type Builder struct {
	*logrus.Entry
}

func (self *Builder) Data(data interface{}) *Builder {
	if globalOptions.DataFielder != nil {
		self.Entry = globalOptions.DataFielder(data, self.Entry)
	}
	return self
}

func (self *Builder) Enabled(data interface{}) *Builder {
	if globalOptions.EnabledChecker != nil && !globalOptions.EnabledChecker(data) {
		self.Entry.Logger = globalOptions.NoLogger
	}
	return self
}

func (self *Builder) Channels(channels ...string) *Builder {
	for _, channel := range channels {
		if _, found := globalOptions.ActiveChannels[channel]; found {
			self.Entry = self.Entry.WithField("channels", channels)
			return self
		}
	}
	self.Entry.Logger = globalOptions.NoLogger
	return self
}

var globalOptions *Options
