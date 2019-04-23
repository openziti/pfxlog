package pfxlog

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Formatter struct {
	start time.Time
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	second := time.Since(f.start).Seconds()
	prefix := strings.TrimPrefix(entry.Caller.Function, prefix)
	if attachment, found := entry.Data["attachment"]; found {
		prefix += " {" + attachment.(string) +"}"
	}
	return []byte(fmt.Sprintf("%s %s %s: %s\n",
			ansi.Blue+fmt.Sprintf("[%5.3f]", second)+ansi.DefaultFG,
			ansi.Yellow+fmt.Sprintf("%5s", entry.Level.String())+ansi.DefaultFG,
			ansi.Green+prefix+ansi.DefaultFG,
			entry.Message),
		),
		nil
}
