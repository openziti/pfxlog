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
	pfxlog.Logger().EnableChannels("subsystem_1").Infof("hello from subsystem_1")
	pfxlog.Logger().EnableChannels("subsystem_2").Infof("hello from subsystem_2")
	pfxlog.ContextLogger("session-1234abcd").EnableChannels("subsystem_1", "subsystem_2").Infof("hello from subsystem_1 and subsystem_2")
	pfxlog.Logger().EnableChannels("subsystem_3").Errorf("should not log")
}
