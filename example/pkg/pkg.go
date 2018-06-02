package pkg

import "github.com/michaelquigley/pfxlog"

func Hello() {
	log := pfxlog.Logger()
	log.Info("Hello")
}