package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

func init() {
	options := pfxlog.DefaultOptions().SetActiveChannels("subsystem_1", "subsystem_2")
	pfxlog.GlobalInit(logrus.InfoLevel, options)
}

func main() {
	pfxlog.Logger().Channels("subsystem_1").Infof("hello from subsystem_1")
	pfxlog.Logger().Channels("subsystem_2").Infof("hello from subsystem_2")
	pfxlog.ContextLogger("session-1234abcd").Channels("subsystem_1", "subsystem_2").Infof("hello from subsystem_1 and subsystem_2")
	pfxlog.Logger().Channels("subsystem_3").Errorf("should not log")
}
