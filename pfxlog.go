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

func Logger() *PfxBuilder {
	return &PfxBuilder{logrus.NewEntry(globalOptions.StandardLogger)}
}

func ContextLogger(context string) *PfxBuilder {
	return &PfxBuilder{globalOptions.StandardLogger.WithField("context", context)}
}

type PfxBuilder struct {
	Entry *logrus.Entry
}

func (self *PfxBuilder) Data(data interface{}) *PfxBuilder {
	if globalOptions.ContextDataFielder != nil {
		self.Entry = globalOptions.ContextDataFielder(data, self.Entry)
	}
	return self
}

func (self *PfxBuilder) Enabled(data interface{}) *PfxBuilder {
	if globalOptions.ContextDataChecker != nil && !globalOptions.ContextDataChecker(data) {
		self.Entry.Logger = globalOptions.NoLogger
	}
	return self
}

func (self *PfxBuilder) Trace(args ...interface{}) {
	self.Entry.Trace(args...)
}

func (self *PfxBuilder) Tracef(format string, args ...interface{}) {
	self.Entry.Tracef(format, args...)
}

func (self *PfxBuilder) Debug(args ...interface{}) {
	self.Entry.Debug(args...)
}

func (self *PfxBuilder) Debugf(format string, args ...interface{}) {
	self.Entry.Debugf(format, args...)
}

func (self *PfxBuilder) Print(args ...interface{}) {
	self.Entry.Print(args...)
}

func (self *PfxBuilder) Printf(format string, args ...interface{}) {
	self.Entry.Printf(format, args...)
}

func (self *PfxBuilder) Info(args ...interface{}) {
	self.Entry.Info(args...)
}

func (self *PfxBuilder) Infof(format string, args ...interface{}) {
	self.Entry.Infof(format, args...)
}

func (self *PfxBuilder) Warn(args ...interface{}) {
	self.Entry.Warn(args...)
}

func (self *PfxBuilder) Warnf(format string, args ...interface{}) {
	self.Entry.Warnf(format, args...)
}

func (self *PfxBuilder) Error(args ...interface{}) {
	self.Entry.Error(args...)
}

func (self *PfxBuilder) Errorf(format string, args ...interface{}) {
	self.Entry.Errorf(format, args...)
}

func (self *PfxBuilder) Fatal(args ...interface{}) {
	self.Entry.Fatal(args...)
}

func (self *PfxBuilder) Fatalf(format string, args ...interface{}) {
	self.Entry.Fatalf(format, args...)
}

func (self *PfxBuilder) Panic(args ...interface{}) {
	self.Entry.Panic(args...)
}

func (self *PfxBuilder) Panicf(format string, args ...interface{}) {
	self.Entry.Panicf(format, args...)
}

var globalOptions *Options
