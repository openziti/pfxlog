package other

import "github.com/michaelquigley/pfxlog"

type Component struct{}

func (c *Component) Hello() {
	pfxlog.Logger().Infof("oh, wow!")
}