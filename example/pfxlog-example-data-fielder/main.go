package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

type contextData struct {
	name  string
	id    string
	value int
}

func init() {
	options := pfxlog.DefaultOptions()
	options.DataFielder = func(v interface{}, l *logrus.Entry) *logrus.Entry {
		cd, ok := v.(*contextData)
		if ok {
			return l.WithFields(map[string]interface{}{
				"name":  cd.name,
				"id":    cd.id,
				"value": cd.value,
			})
		} else {
			return l.WithFields(nil)
		}
	}
	pfxlog.GlobalInit(logrus.InfoLevel, options)
}

func main() {
	pfxlog.Logger().Data(&contextData{"testing", "0x33", 33}).Entry.WithField("testing", "a").Infof("oh, wow!")
}
