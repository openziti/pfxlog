package pfxlog

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	StartTimestamp time.Time
	AbsoluteTime   bool
	TrimPrefix     string

	PanicLabel   string
	FatalLabel   string
	ErrorLabel   string
	WarningLabel string
	InfoLabel    string
	DebugLabel   string
	TraceLabel   string

	TimestampColor string
	FunctionColor  string
	FieldsColor    string
	DefaultFgColor string

	PrettyTimestampFormat string
	JsonTimestampFormat   string

	ContextDataFielder func(data interface{}, logger *logrus.Logger) *logrus.Entry
}

func DefaultOptions() *Options {
	options := &Options{
		StartTimestamp:        time.Now(),
		AbsoluteTime:          false,
		PrettyTimestampFormat: "2006-01-02 15:04:05.000",
		JsonTimestampFormat:   "2006-01-02T15:04:05.000Z",
	}
	if color, err := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_USE_COLOR"))); err == nil && color {
		return options.Color()
	} else {
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "pfxlog: error parsing 'PFXLOG_USE_COLOR' (%v)", err)
		}
		return options.NoColor()
	}
}

func (options *Options) Starting(t time.Time) *Options {
	options.StartTimestamp = t
	return options
}

func (options *Options) StartingToday() *Options {
	now := time.Now()
	options.StartTimestamp = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return options
}

func (options *Options) SetAbsoluteTime() *Options {
	options.AbsoluteTime = true
	return options
}

func (options *Options) SetTrimPrefix(prefix string) *Options {
	options.TrimPrefix = prefix
	return options
}

func (options *Options) Color() *Options {
	options.PanicLabel = ansi.Red + "  PANIC" + ansi.DefaultFG
	options.FatalLabel = ansi.Red + "  FATAL" + ansi.DefaultFG
	options.ErrorLabel = ansi.Red + "  ERROR" + ansi.DefaultFG
	options.WarningLabel = ansi.Yellow + "WARNING" + ansi.DefaultFG
	options.InfoLabel = ansi.White + "   INFO" + ansi.DefaultFG
	options.DebugLabel = ansi.Blue + "  DEBUG" + ansi.DefaultFG
	options.TraceLabel = ansi.LightBlack + "  TRACE" + ansi.DefaultFG

	options.TimestampColor = ansi.Blue
	options.FunctionColor = ansi.Cyan
	options.FieldsColor = ansi.LightCyan
	options.DefaultFgColor = ansi.DefaultFG

	return options
}

func (options *Options) NoColor() *Options {
	options.PanicLabel = "  PANIC"
	options.FatalLabel = "  FATAL"
	options.ErrorLabel = "  ERROR"
	options.WarningLabel = "WARNING"
	options.InfoLabel = "   INFO"
	options.DebugLabel = "  DEBUG"
	options.TraceLabel = "  TRACE"

	options.TimestampColor = ""
	options.FunctionColor = ""
	options.FieldsColor = ""
	options.DefaultFgColor = ""

	return options
}
