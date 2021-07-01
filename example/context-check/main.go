package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

func init() {
	options := pfxlog.DefaultOptions()
	options.ContextDataFielder = func(v interface{}, logger *logrus.Logger) *logrus.Entry {
		if i, ok := v.(int); ok {
			return logger.WithField("i", i)
		}
		return logger.WithFields(nil)
	}
	options.ContextChecker = func(v interface{}) bool {
		i, ok := v.(int)
		if ok {
			if i % 2 == 0 {
				return true
			}
		}
		return false
	}
	pfxlog.GlobalInit(logrus.InfoLevel, options)
}

func main() {
	logrus.Info("starting")
	pfxlog.ContextCheck(2).Info("oh, wow!")
	pfxlog.ContextCheck(3).Warn("oh, no!")
	pfxlog.ContextCheckData(44).Info("show it")
}

