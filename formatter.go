package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type formatter struct {
	start time.Time
}

func NewFormatter() logrus.Formatter {
	return &formatter{start: time.Now()}
}

func NewFormatterStarting(start time.Time) logrus.Formatter {
	return &formatter{start: start}
}

func NewFormatterStartingToday() logrus.Formatter {
	now := time.Now()
	dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return &formatter{start: dayStart}
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
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
	trimmedFunction := ""
	if entry.Caller != nil {
		trimmedFunction = strings.TrimPrefix(entry.Caller.Function, prefix)
	}
	if context, found := entry.Data["context"]; found {
		trimmedFunction += " [" + context.(string) + "]"
	}
	message := entry.Message
	if withFields(entry.Data) {
		fields := "{"
		field := 0
		for k, v := range entry.Data {
			if k != "context" {
				if field > 0 {
					fields += " "
				}
				field++
				fields += fmt.Sprintf("%s=[%v]", k, v)
			}
		}
		fields += "} "
		message = lightCyanColor + fields + defaultFgColor + message
	}
	return []byte(fmt.Sprintf("%s %s %s: %s\n",
			blueColor+fmt.Sprintf("[%8.3f]", second)+defaultFgColor,
			level,
			cyanColor+trimmedFunction+defaultFgColor,
			message),
		),
		nil
}

func withFields(data map[string]interface{}) bool {
	if _, found := data["context"]; found {
		return len(data) > 1
	} else {
		return len(data) > 0
	}
}
