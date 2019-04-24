package main

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/michaelquigley/pfxlog/example/other"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	pfxlog.Global(logrus.DebugLevel)
	pfxlog.SetPrefix("github.com/michaelquigley/")
}

func main() {
	log := pfxlog.Logger()
	log.Info("hello world.")

	notifications := make(chan int)
	for i := 0; i < 50; i++ {
		go counter(i, notifications)
	}

	for i := 0; i < 50; i++ {
		n := <-notifications
		log.Infof("%d done.", n)
	}

	log.Info("complete.")
}

func counter(number int, notify chan int) {
	log := pfxlog.ContextLogger(fmt.Sprintf("#%d", number))

	for i := 0; i < 5; i++ {
		log.Infof("visited %d.", i)
	}

	time.Sleep(1 * time.Second)

	c := &other.Component{}
	c.Hello()

	log.Info("complete")

	notify <- number
}
