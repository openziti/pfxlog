package pfxlog

import (
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

var prefix string

func Global(level logrus.Level) {
	noJson, _ := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_NO_JSON")))
	if noJson || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(NewFormatter())
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
}

func SetPrefix(p string) {
	prefix = p
}

func SetDefaultNoColor() {
	useColorVar := strings.ToLower(os.Getenv("PFXLOG_USE_COLOR"))
	useColor := false
	if useColorVar != "" {
		useColor, _ = strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_USE_COLOR")))
	}
	if !useColor {
		blueColor = ""
		cyanColor = ""
		defaultFgColor = ""
		lightCyanColor = ""
		panicColor = "  PANIC"
		fatalColor = "  FATAL"
		errorColor = "  ERROR"
		warnColor = "WARNING"
		infoColor = "   INFO"
		debugColor = "  DEBUG"
		traceColor = "  TRACE"
	}
}

func ContextLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

var panicColor = ansi.Red + "  PANIC" + ansi.DefaultFG
var fatalColor = ansi.Red + "  FATAL" + ansi.DefaultFG
var errorColor = ansi.Red + "  ERROR" + ansi.DefaultFG
var warnColor = ansi.Yellow + "WARNING" + ansi.DefaultFG
var infoColor = ansi.White + "   INFO" + ansi.DefaultFG
var debugColor = ansi.Blue + "  DEBUG" + ansi.DefaultFG
var traceColor = ansi.LightBlack + "  TRACE" + ansi.DefaultFG

var blueColor = ansi.Blue
var cyanColor = ansi.Cyan
var defaultFgColor = ansi.DefaultFG
var lightCyanColor = ansi.LightCyan
