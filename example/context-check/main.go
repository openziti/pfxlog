package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

func init() {
	options := pfxlog.DefaultOptions()
	options.ContextDataFielder = func(v interface{}, entry *logrus.Entry) *logrus.Entry {
		if i, ok := v.(int); ok {
			return entry.WithField("i", i)
		}
		return entry.WithFields(nil)
	}
	options.ContextDataChecker = func(v interface{}) bool {
		i, ok := v.(int)
		if ok {
			if i%2 == 0 {
				return true
			}
		}
		return false
	}
	pfxlog.GlobalInit(logrus.InfoLevel, options)
}

func main() {
	logrus.Info("starting")
	pfxlog.Logger().Enabled(2).Info("oh, wow!")
	pfxlog.Logger().Enabled(3).Warn("oh, no!")
	pfxlog.Logger().Enabled(44).Data(44).Info("show it")
}
