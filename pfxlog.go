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
	Entry *logrus.Entry
}

func (self *Builder) Data(data interface{}) *Builder {
	if globalOptions.ContextDataFielder != nil {
		self.Entry = globalOptions.ContextDataFielder(data, self.Entry)
	}
	return self
}

func (self *Builder) Enabled(data interface{}) *Builder {
	if globalOptions.ContextDataChecker != nil && !globalOptions.ContextDataChecker(data) {
		self.Entry.Logger = globalOptions.NoLogger
	}
	return self
}

func (self *Builder) Trace(args ...interface{}) {
	self.Entry.Trace(args...)
}

func (self *Builder) Tracef(format string, args ...interface{}) {
	self.Entry.Tracef(format, args...)
}

func (self *Builder) Debug(args ...interface{}) {
	self.Entry.Debug(args...)
}

func (self *Builder) Debugf(format string, args ...interface{}) {
	self.Entry.Debugf(format, args...)
}

func (self *Builder) Print(args ...interface{}) {
	self.Entry.Print(args...)
}

func (self *Builder) Printf(format string, args ...interface{}) {
	self.Entry.Printf(format, args...)
}

func (self *Builder) Info(args ...interface{}) {
	self.Entry.Info(args...)
}

func (self *Builder) Infof(format string, args ...interface{}) {
	self.Entry.Infof(format, args...)
}

func (self *Builder) Warn(args ...interface{}) {
	self.Entry.Warn(args...)
}

func (self *Builder) Warnf(format string, args ...interface{}) {
	self.Entry.Warnf(format, args...)
}

func (self *Builder) Error(args ...interface{}) {
	self.Entry.Error(args...)
}

func (self *Builder) Errorf(format string, args ...interface{}) {
	self.Entry.Errorf(format, args...)
}

func (self *Builder) Fatal(args ...interface{}) {
	self.Entry.Fatal(args...)
}

func (self *Builder) Fatalf(format string, args ...interface{}) {
	self.Entry.Fatalf(format, args...)
}

func (self *Builder) Panic(args ...interface{}) {
	self.Entry.Panic(args...)
}

func (self *Builder) Panicf(format string, args ...interface{}) {
	self.Entry.Panicf(format, args...)
}

var globalOptions *Options
