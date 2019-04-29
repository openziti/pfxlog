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
	var level string
	switch entry.Level {
	case logrus.PanicLevel:
		level = panicColor
	case logrus.FatalLevel:
		level = fatalColor
	case logrus.ErrorLevel:
		level = errorColor
	case logrus.WarnLevel:
		level = warnColor
	case logrus.InfoLevel:
		level = infoColor
	case logrus.DebugLevel:
		level = debugColor
	case logrus.TraceLevel:
		level = traceColor
	}
	prefix := strings.TrimPrefix(entry.Caller.Function, prefix)
	if context, found := entry.Data["context"]; found {
		prefix += " [" + context.(string) + "]"
	}
	return []byte(fmt.Sprintf("%s %s %s: %s\n",
			ansi.Blue+fmt.Sprintf("[%8.3f]", second)+ansi.DefaultFG,
			level,
			ansi.Cyan+prefix+ansi.DefaultFG,
			entry.Message),
		),
		nil
}

var panicColor = ansi.Red + "  PANIC" + ansi.DefaultFG
var fatalColor = ansi.Red + "  FATAL" + ansi.DefaultFG
var errorColor = ansi.Red + "  ERROR" + ansi.DefaultFG
var warnColor = ansi.Yellow + "WARNING" + ansi.DefaultFG
var infoColor = ansi.White + "   INFO" + ansi.DefaultFG
var debugColor = ansi.Blue + "  DEBUG" + ansi.DefaultFG
var traceColor = ansi.LightBlack + "  TRACE" + ansi.DefaultFG
