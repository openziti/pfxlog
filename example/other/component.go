package other

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

type Component struct{}

func (c *Component) Hello() {
	logrus.Warnf("this is #%d", 6)
	pfxlog.Logger().Infof("oh, wow!")
	pfxlog.Logger().Errorf("uh...")
}